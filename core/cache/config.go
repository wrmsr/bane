package cache

import (
	"sync"
	"time"
)

//

type Weight = float32

//

type Eviction interface {
	target(c *cache) *link
}

type LRU struct{}

func (e LRU) target(c *cache) *link { return c.root.lru.next }

type LRI struct{}

func (e LRI) target(c *cache) *link { return c.root.ins.next }

type LFU struct{}

func (e LFU) target(c *cache) *link { return c.root.lfu.prev }

//

type OverweightError struct {
	Key, Value any
}

func (e OverweightError) Error() string {
	return "cache value overweight"
}

func PanicOnOverweight(k, v any) { panic(OverweightError{k, v}) }

//

type Config struct {
	MaxSize int

	Weigher   func(v any) Weight
	MaxWeight Weight

	ExpireAfterAccess time.Duration
	ExpireAfterWrite  time.Duration

	OnAdd        func(k, v any)
	OnRemove     func(k, v any)
	OnOverweight func(k, v any)

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
