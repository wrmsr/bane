package cache

import (
	"reflect"
	"time"

	ctr "github.com/wrmsr/bane/pkg/util/container"
	its "github.com/wrmsr/bane/pkg/util/iterators"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type Stats struct {
	Seq int

	Size   int
	Weight Weight

	Hits   int
	Misses int

	MaxSizeEver   int
	MaxWeightEver Weight
}

type list struct {
	next, prev *link
}

func (l *list) set(next, prev *link) {
	l.next = next
	l.prev = prev
}

func unlink(l *link, self, next, prev *list) {
	if next.prev != l || prev.next != l {
		panic("broken list")
	}
	prev.next = self.next
	next.prev = self.prev
	self.next = l
	self.prev = l
}

type link struct {
	seq int

	ins list
	lru list
	lfu list

	key    any
	value  any
	weight Weight

	written  int64
	accessed int64

	hits     int
	unlinked bool
}

type cache struct {
	root *link
}

type Cache[K comparable] struct {
	cfg Config

	cache

	m map[K]*link

	stats Stats

	eaaMillis int64
	eawMillis int64
}

func NewCache[K comparable](cfg Config) *Cache[K] {
	if _, ok := cfg.Eviction.(LFU); ok {
		cfg.TrackFrequency = true
	}

	root := &link{
		seq: 0,
	}

	root.ins.set(root, root)
	root.lru.set(root, root)
	if cfg.TrackFrequency {
		root.lfu.set(root, root)
	}

	return &Cache[K]{
		cfg: cfg,

		cache: cache{
			root: root,
		},

		m: make(map[K]*link),

		eaaMillis: int64(cfg.ExpireAfterAccess / time.Millisecond),
		eawMillis: int64(cfg.ExpireAfterWrite / time.Millisecond),
	}
}

func (c *Cache[K]) unlink(l *link) {
	if l == c.root {
		panic("can't unlink root")
	}
	if l.unlinked {
		return
	}

	unlink(l, &l.ins, &l.ins.next.ins, &l.ins.prev.ins)
	unlink(l, &l.lru, &l.lru.next.lru, &l.lru.prev.lru)
	if c.cfg.TrackFrequency {
		unlink(l, &l.lfu, &l.lfu.next.lfu, &l.lfu.prev.lfu)
	}

	if c.cfg.OnRemove != nil {
		c.cfg.OnRemove(l.key, l.value)
	}

	l.key = nil
	l.value = nil
	c.stats.Size--
	if c.cfg.Weigher != nil {
		c.stats.Weight -= l.weight
	}
	l.unlinked = true
}

func (c *Cache[K]) kill(l *link) {
	if l == c.root {
		panic("can't kill root")
	}

	cacheLink := c.m[l.key.(K)]
	if cacheLink == l {
		delete(c.m, l.key.(K))
	}

	c.unlink(l)
}

func (c *Cache[K]) now() int64 {
	return c.cfg.Clock().UnixMilli()
}

func (c *Cache[K]) reap() {
	var now int64

	if c.eawMillis > 0 {
		now = c.now()
		deadline := now - c.eawMillis
		for c.root.ins.next != c.root {
			l := c.root.ins.next
			if l.written > deadline {
				break
			}
			c.kill(l)
		}
	}

	if c.eaaMillis > 0 {
		if now == 0 {
			now = c.now()
		}
		deadline := now - c.eaaMillis
		for c.root.lru.next != c.root {
			l := c.root.lru.next
			if l.accessed > deadline {
				break
			}
			c.kill(l)
		}
	}
}

func (c *Cache[K]) Reap() {
	if c.cfg.Lock != nil {
		c.cfg.Lock.Lock()
		defer c.cfg.Lock.Unlock()
	}

	c.reap()
}

func (c *Cache[K]) getLink(k K) *link {
	l := c.m[k]
	if l == nil {
		return nil
	}
	if l.unlinked {
		panic("unlinked")
	}
	return l
}

func (c *Cache[K]) TryGet(k K) (any, bool) {
	if c.cfg.Lock != nil {
		c.cfg.Lock.Lock()
		defer c.cfg.Lock.Unlock()
	}

	c.reap()

	l := c.getLink(k)
	if l == nil {
		c.stats.Misses++
		return nil, false
	}

	if l.lru.next != c.root {
		unlink(l, &l.lru, &l.lru.next.lru, &l.lru.prev.lru)

		lruLast := c.root.lru.prev
		lruLast.lru.next = l
		c.root.lru.prev = l
		l.lru.prev = lruLast
		l.lru.next = c.root
	}

	if c.cfg.TrackFrequency {
		lfuPos := l.lfu.prev
		for lfuPos != c.root && lfuPos.hits <= l.hits {
			lfuPos = lfuPos.lfu.prev
		}

		if l.lfu.prev != lfuPos {
			unlink(l, &l.lfu, &l.lfu.next.lfu, &l.lfu.prev.lfu)

			lfuLast := lfuPos.lfu.prev
			lfuLast.lfu.next = l
			lfuPos.lfu.prev = l
			l.lfu.prev = lfuLast
			l.lfu.next = lfuPos
		}
	}

	l.accessed = c.now()
	l.hits++
	c.stats.Hits++
	return l.value, true
}

func (c *Cache[K]) Get(k K) any {
	v, _ := c.TryGet(k)
	return v
}

func (c *Cache[K]) full() bool {
	if c.cfg.MaxSize > 0 && c.stats.Size >= c.cfg.MaxSize {
		return true
	}
	if c.cfg.Weigher != nil && c.stats.Weight >= c.cfg.MaxWeight {
		return true
	}
	return false
}

func (c *Cache[K]) Clear() {
	if c.cfg.Lock != nil {
		c.cfg.Lock.Lock()
		defer c.cfg.Lock.Unlock()
	}

	c.m = make(map[K]*link)
	for {
		l := c.root.ins.prev
		if l == c.root {
			break
		}
		if l.unlinked {
			panic("already unlinked")
		}
		c.unlink(l)
	}
}

func (c *Cache[K]) Put(k K, v any) bool {
	var weight Weight
	if c.cfg.Weigher != nil {
		weight = c.cfg.Weigher(v)
	}

	if c.cfg.Lock != nil {
		c.cfg.Lock.Lock()
		defer c.cfg.Lock.Unlock()
	}
	c.reap()

	if c.cfg.Weigher != nil && c.cfg.MaxWeight > 0 && weight > c.cfg.MaxWeight {
		if c.cfg.OnOverweight != nil {
			c.cfg.OnOverweight(k, v)
		}
		return false
	}

	if l := c.getLink(k); l != nil {
		c.unlink(l)
	}

	for c.full() {
		t := c.cfg.Eviction.target(&c.cache)
		if t == c.root {
			panic("can't evict root")
		}
		c.kill(t)
	}

	c.stats.Seq++
	now := c.now()

	l := &link{
		seq: c.stats.Seq,

		key:   k,
		value: v,

		written:  now,
		accessed: now,

		weight: weight,
	}

	insLast := c.root.ins.prev
	insLast.ins.next = l
	c.root.ins.prev = l
	l.ins.prev = insLast
	l.ins.next = c.root

	lruLast := c.root.lru.prev
	lruLast.lru.next = l
	c.root.lru.prev = l
	l.lru.prev = lruLast
	l.lru.next = c.root

	if c.cfg.TrackFrequency {
		lfuLast := c.root.lfu.prev
		lfuLast.lfu.next = l
		c.root.lfu.prev = l
		l.lfu.prev = lfuLast
		l.lfu.next = c.root
	}

	c.stats.Size++
	c.stats.Weight += weight

	if c.stats.Size > c.stats.MaxSizeEver {
		c.stats.MaxSizeEver = c.stats.Size
	}
	if c.stats.Weight > c.stats.MaxWeightEver {
		c.stats.MaxWeightEver = c.stats.Weight
	}

	c.m[k] = l

	return true
}

func (c *Cache[K]) Delete(k K) bool {
	if c.cfg.Lock != nil {
		c.cfg.Lock.Lock()
		defer c.cfg.Lock.Unlock()
	}
	c.reap()

	l := c.getLink(k)
	if l == nil {
		return false
	}

	delete(c.m, k)
	c.unlink(l)

	return true
}

func (c *Cache[K]) Len() int {
	if c.cfg.Lock != nil {
		c.cfg.Lock.Lock()
		defer c.cfg.Lock.Unlock()
	}
	c.reap()

	return c.stats.Size
}

func (c *Cache[K]) Contains(k K) bool {
	if c.cfg.Lock != nil {
		c.cfg.Lock.Lock()
		defer c.cfg.Lock.Unlock()
	}
	c.reap()

	return c.getLink(k) != nil
}

func (c *Cache[K]) ForEach(fn func(bt.Kv[K, any]) bool) bool {
	if c.cfg.Lock != nil {
		c.cfg.Lock.Lock()
		defer c.cfg.Lock.Unlock()
	}
	c.reap()

	for l := c.root.ins.prev; l != c.root; {
		if !fn(bt.KvOf(l.key.(K), l.value)) {
			return false
		}
		n := l.ins.prev
		if l == n {
			panic("cycle")
		}
		l = n
	}
	return true
}

//

var _ ctr.Map[int, any] = &Cache[int]{}

func (c *Cache[K]) Iterate() its.Iterator[bt.Kv[K, any]] {
	return its.OfSlice(its.SeqForEach[bt.Kv[K, any]](c)).Iterate()
}

func (c *Cache[K]) ReflectTypeArgs() []reflect.Type {
	return []reflect.Type{rfl.TypeOf[K]()}
}
