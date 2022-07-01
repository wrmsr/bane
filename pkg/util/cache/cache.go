package cache

import (
	"sync"
	"time"
)

//

type Weight = float32

//

type Eviction interface {
	evict(c *Cache)
}

type LRU struct{}

func (e LRU) evict(c *Cache) { c.kill(c.root.lru.next) }

type LRI struct{}

func (e LRI) evict(c *Cache) { c.kill(c.root.ins.next) }

type LFU struct{}

func (e LFU) evict(c *Cache) { c.kill(c.root.lfu.prev) }

//

type Config struct {
	MaxSize int

	Weigher   func(v any) Weight
	MaxWeight Weight

	ExpireAfterAccess time.Duration
	ExpireAfterWrite  time.Duration

	OnAdd        func(k, v any)
	OnRemove     func(k, v any)
	OnOverweight func()

	Clock func() time.Time

	Lock sync.Locker

	Eviction Eviction

	TrackFrequency bool
}

const DefaultMaxSize = 256

func DefaultConfig() Config {
	return Config{
		MaxSize: DefaultMaxSize,

		Clock: time.Now,

		Eviction: LRU{},
	}
}

//

type Stats struct {
	Seq int

	Size   int
	Weight Weight

	Hits   int
	Misses int

	MaxSizeEver   int
	MaxWeightEver float64
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

	writtenMilli  int64
	accessedMilli int64

	hits     int
	unlinked bool
}

type Cache struct {
	cfg Config

	m map[any]*link

	root  *link
	stats Stats

	eaaMillis int64
	eawMillis int64
}

func NewCache(cfg Config) *Cache {
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

	return &Cache{
		cfg: cfg,

		m: make(map[any]*link),

		root: root,

		eaaMillis: int64(cfg.ExpireAfterAccess / time.Millisecond),
		eawMillis: int64(cfg.ExpireAfterWrite / time.Millisecond),
	}
}

func (c *Cache) unlink(l *link) {
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

func (c *Cache) kill(l *link) {
	if l == c.root {
		panic("can't kill root")
	}

	cacheLink := c.m[l.key]
	if cacheLink == l {
		delete(c.m, l.key)
	}

	c.unlink(l)
}

func (c *Cache) now() int64 {
	c.cfg.Clock().UnixMilli()
}

func (c *Cache) reap() {
	var now int64

	if c.eawMillis > 0 {
		now = c.now()
		deadline := now - c.eawMillis
		for c.root.ins.next != c.root {
			l := c.root.ins.next
			if l.writtenMilli > deadline {
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
			if l.accessedMilli > deadline {
				break
			}
			c.kill(l)
		}
	}
}

func (c *Cache) Reap() {
	if c.cfg.Lock != nil {
		c.cfg.Lock.Lock()
		defer c.cfg.Lock.Unlock()
	}

	c.reap()
}

func (c *Cache) getLink(k any) *link {
	l := c.m[k]
	if l == nil {
		return nil
	}
	if l.unlinked {
		panic("unlinked")
	}
	return l
}

func (c *Cache) Get(k any) (any, bool) {
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

	l.accessedMilli = c.now()
	l.hits++
	c.stats.Hits++
	return l.value, true
}

func (c *Cache) full() bool {
	if c.cfg.MaxSize > 0 && c.stats.Size >= c.cfg.MaxSize {
		return true
	}
	if c.cfg.Weigher != nil && c.stats.Weight >= c.cfg.MaxWeight {
		return true
	}
	return false
}

func (c *Cache) Clear() {
	if c.cfg.Lock != nil {
		c.cfg.Lock.Lock()
		defer c.cfg.Lock.Unlock()
	}

	c.m = make(map[any]*link)
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

/*
def __setitem__(self, key: K, value: V) -> None:
   weight = c._weigher(value)

   with c._lock():
	   c._reap()

	   if c._max_weight is not None and weight > c._max_weight:
		   if c._raise_overweight:
			   raise OverweightException
		   else:
			   return

	   try:
		   existing_link, existing_value = c._get_link(key)
	   except KeyError:
		   pass
	   else:
		   c._unlink(existing_link)

	   while c._full:
		   c._eviction(self)

	   link = CacheImpl.Link()
	   c._seq += 1
	   link.seq = c._seq
	   link.key = weakref.ref(key, functools.partial(CacheImpl._weak_die, c._weak_dead_ref, link)) if c._weak_keys else key  # noqa
	   link.value = weakref.ref(value, functools.partial(CacheImpl._weak_die, c._weak_dead_ref, link)) if c._weak_values else value  # noqa
	   link.weight = weight
	   link.written = link.accessed = c._clock()
	   link.hits = 0
	   link.unlinked = False

	   ins_last = c._root.ins_prev
	   ins_last.ins_next = c._root.ins_prev = link
	   link.ins_prev = ins_last
	   link.ins_next = c._root

	   lru_last = c._root.lru_prev
	   lru_last.lru_next = c._root.lru_prev = link
	   link.lru_prev = lru_last
	   link.lru_next = c._root

	   if c._track_frequency:
		   lfu_last = c._root.lfu_prev
		   lfu_last.lfu_next = c._root.lfu_prev = link
		   link.lfu_prev = lfu_last
		   link.lfu_next = c._root

	   c._weight += weight
	   c._size += 1
	   c._max_size_ever = max(c._size, c._max_size_ever)
	   c._max_weight_ever = max(c._weight, c._max_weight_ever)

	   cache_key = id(key) if c._identity_keys else key
	   c.m[cache_key] = link

def __delitem__(self, key: K) -> None:
   with c._lock():
	   c._reap()

	   link, value = c._get_link(key)

	   cache_key = id(key) if c._identity_keys else key
	   del c.m[cache_key]

	   c._unlink(link)

def __len__(self) -> int:
   with c._lock():
	   c._reap()

	   return c._size

def __contains__(self, key: K) -> bool:
   with c._lock():
	   c._reap()

	   try:
		   c._get_link(key)
	   except KeyError:
		   return False
	   else:
		   return True

def __iter__(self) -> ta.Iterator[K]:
   with c._lock():
	   c._reap()

	   link = c._root.ins_prev
	   while link is not c._root:
		   yield link.key
		   next = link.ins_prev
		   if next is link:
			   raise ValueError
		   link = next

*/
