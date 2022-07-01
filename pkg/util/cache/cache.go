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

	root  *link
	stats Stats
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

		root: root,
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
	unlink(l, &l.lru, &l.ins.next.lru, &l.ins.prev.lru)
	if c.cfg.TrackFrequency {
		unlink(l, &l.lfu, &l.ins.next.lfu, &l.ins.prev.lfu)
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

}

/*
   def _kill(self, link: Link) -> None:
       if link is self._root:
           raise RuntimeError

       key = link.key
       if self._weak_keys:
           if key is not None:
               key = key()
           if key is None:
               key = SKIP

       if key is not SKIP:
           cache_key = id(key) if self._identity_keys else key
           cache_link = self._cache.get(cache_key)
           if cache_link is link:
               del self._cache[cache_key]

       self._unlink(link)

   def _reap(self) -> None:
       if self._weak_dead is not None:
           while True:
               try:
                   link = self._weak_dead.popleft()
               except IndexError:
                   break
               self._kill(link)

       clock = None

       if self._expire_after_write is not None:
           clock = self._clock()
           deadline = clock - self._expire_after_write

           while self._root.ins_next is not self._root:
               link = self._root.ins_next
               if link.written > deadline:
                   break
               self._kill(link)

       if self._expire_after_access is not None:
           if clock is None:
               clock = self._clock()
           deadline = clock - self._expire_after_access

           while self._root.lru_next is not self._root:
               link = self._root.lru_next
               if link.accessed > deadline:
                   break
               self._kill(link)

   def reap(self) -> None:
       with self._lock():
           self._reap()

   def _get_link(self, key: K) -> ta.Tuple[Link, V]:
       cache_key = id(key) if self._identity_keys else key

       link = self._cache[cache_key]
       if link.unlinked:
           raise Exception

       def fail():
           try:
               del self._cache[cache_key]
           except KeyError:
               pass
           self._unlink(link)
           raise KeyError(key)

       if self._identity_keys:
           link_key = link.key
           if self._weak_keys:
               link_key = link_key()
               if link_key is None:
                   fail()
           if key is not link_key:
               fail()

       value = link.value
       if self._weak_values:
           if value is not None:
               value = value()
           if value is None:
               fail()

       return link, value

   def __getitem__(self, key: K) -> V:
       with self._lock():
           self._reap()

           try:
               link, value = self._get_link(key)
           except KeyError:
               self._misses += 1
               raise KeyError(key)

           if link.lru_next is not self._root:
               link.lru_prev.lru_next = link.lru_next
               link.lru_next.lru_prev = link.lru_prev

               lru_last = self._root.lru_prev
               lru_last.lru_next = self._root.lru_prev = link
               link.lru_prev = lru_last
               link.lru_next = self._root

           if self._track_frequency:
               lfu_pos = link.lfu_prev
               while lfu_pos is not self._root and lfu_pos.hits <= link.hits:
                   lfu_pos = lfu_pos.lfu_prev

               if link.lfu_prev is not lfu_pos:
                   link.lfu_prev.lfu_next = link.lfu_next
                   link.lfu_next.lfu_prev = link.lfu_prev

                   lfu_last = lfu_pos.lfu_prev
                   lfu_last.lfu_next = lfu_pos.lfu_prev = link
                   link.lfu_prev = lfu_last
                   link.lfu_next = lfu_pos

           link.accessed = self._clock()
           link.hits += 1
           self._hits += 1
           return value

   @staticmethod
   def _weak_die(dead_ref: weakref.ref, link: Link, key_ref: weakref.ref) -> None:
       dead = dead_ref()
       if dead is not None:
           dead.append(link)

   @property
   def _full(self) -> bool:
       if self._max_size is not None and self._size >= self._max_size:
           return True
       if self._max_weight is not None and self._weight >= self._max_weight:
           return True
       return False

   def clear(self) -> None:
       with self._lock():
           self._cache.clear()
           while True:
               link = self._root.ins_prev
               if link is self._root:
                   break
               if link.unlinked:
                   raise TypeError
               self._unlink(link)

   def __setitem__(self, key: K, value: V) -> None:
       weight = self._weigher(value)

       with self._lock():
           self._reap()

           if self._max_weight is not None and weight > self._max_weight:
               if self._raise_overweight:
                   raise OverweightException
               else:
                   return

           try:
               existing_link, existing_value = self._get_link(key)
           except KeyError:
               pass
           else:
               self._unlink(existing_link)

           while self._full:
               self._eviction(self)

           link = CacheImpl.Link()
           self._seq += 1
           link.seq = self._seq
           link.key = weakref.ref(key, functools.partial(CacheImpl._weak_die, self._weak_dead_ref, link)) if self._weak_keys else key  # noqa
           link.value = weakref.ref(value, functools.partial(CacheImpl._weak_die, self._weak_dead_ref, link)) if self._weak_values else value  # noqa
           link.weight = weight
           link.written = link.accessed = self._clock()
           link.hits = 0
           link.unlinked = False

           ins_last = self._root.ins_prev
           ins_last.ins_next = self._root.ins_prev = link
           link.ins_prev = ins_last
           link.ins_next = self._root

           lru_last = self._root.lru_prev
           lru_last.lru_next = self._root.lru_prev = link
           link.lru_prev = lru_last
           link.lru_next = self._root

           if self._track_frequency:
               lfu_last = self._root.lfu_prev
               lfu_last.lfu_next = self._root.lfu_prev = link
               link.lfu_prev = lfu_last
               link.lfu_next = self._root

           self._weight += weight
           self._size += 1
           self._max_size_ever = max(self._size, self._max_size_ever)
           self._max_weight_ever = max(self._weight, self._max_weight_ever)

           cache_key = id(key) if self._identity_keys else key
           self._cache[cache_key] = link

   def __delitem__(self, key: K) -> None:
       with self._lock():
           self._reap()

           link, value = self._get_link(key)

           cache_key = id(key) if self._identity_keys else key
           del self._cache[cache_key]

           self._unlink(link)

   def __len__(self) -> int:
       with self._lock():
           self._reap()

           return self._size

   def __contains__(self, key: K) -> bool:
       with self._lock():
           self._reap()

           try:
               self._get_link(key)
           except KeyError:
               return False
           else:
               return True

   def __iter__(self) -> ta.Iterator[K]:
       with self._lock():
           self._reap()

           link = self._root.ins_prev
           while link is not self._root:
               yield link.key
               next = link.ins_prev
               if next is link:
                   raise ValueError
               link = next

*/
