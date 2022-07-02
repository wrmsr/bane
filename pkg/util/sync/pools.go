package sync

import (
	"sync"
	"sync/atomic"
)

type DrainPoolStats struct {
	New, Put, Get int64
}

type DrainPool[T any] struct {
	o sync.Pool

	s DrainPoolStats
}

func NewDrainPool[T any](fn func() T) *DrainPool[T] {
	p := &DrainPool[T]{}
	p.o.New = func() any {
		atomic.AddInt64(&p.s.New, 1)
		return fn()
	}
	return p
}

func (p *DrainPool[T]) Put(x T) {
	atomic.AddInt64(&p.s.Put, 1)
	p.o.Put(x)
}

func (p *DrainPool[T]) Get() T {
	atomic.AddInt64(&p.s.Get, 1)
	return p.o.Get()
}

func (p *DrainPool[T]) Stats() DrainPoolStats {
	return p.s
}

func (p *DrainPool[T]) Drain() int {
	i := 0
	for c := atomic.LoadInt64(&p.s.New); c == atomic.LoadInt64(&p.s.New); i++ {
		p.o.Get()
	}
	return i
}
