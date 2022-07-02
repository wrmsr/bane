package sync

import (
	"sync"
	"sync/atomic"
)

//

type Pool[T any] interface {
	Get() T
	Put(x T)
}

//

type WrappedSyncPool[T any] struct {
	o *sync.Pool
}

func WrapSyncPool[T any](o *sync.Pool) WrappedSyncPool[T] {
	return WrappedSyncPool[T]{o: o}
}

var _ Pool[int] = WrappedSyncPool[int]{}

func (p WrappedSyncPool[T]) Get() T {
	return p.o.Get().(T)
}

func (p WrappedSyncPool[T]) Put(x T) {
	p.o.Put(x)
}

//

type SyncPool[T any] struct {
	o sync.Pool
}

func NewSyncPool[T any](fn func() T) *SyncPool[T] {
	return &SyncPool[T]{o: sync.Pool{New: func() any { return fn() }}}
}

var _ Pool[int] = &SyncPool[int]{}

func (p *SyncPool[T]) Get() T {
	return p.o.Get().(T)
}

func (p *SyncPool[T]) Put(x T) {
	p.o.Put(x)
}

//

type StubPool[T any] struct {
	New func() T
}

var _ Pool[int] = StubPool[int]{}

func (p StubPool[T]) Get() T {
	return p.New()
}

func (p StubPool[T]) Put(x T) {}

//

type DrainPoolStats struct {
	New, Put, Get, Drain int64
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

var _ Pool[int] = &DrainPool[int]{}

func (p *DrainPool[T]) Put(x T) {
	atomic.AddInt64(&p.s.Put, 1)
	p.o.Put(x)
}

func (p *DrainPool[T]) Get() T {
	atomic.AddInt64(&p.s.Get, 1)
	return p.o.Get().(T)
}

func (p *DrainPool[T]) Stats() DrainPoolStats {
	return p.s
}

func (p *DrainPool[T]) Drain() int {
	i := 0
	for c := atomic.LoadInt64(&p.s.New); c == atomic.LoadInt64(&p.s.New); i++ {
		p.o.Get()
	}
	atomic.AddInt64(&p.s.Drain, int64(i))
	return i
}

//

type trackingPoolNode struct {
	k uintptr
	v any
}

func (n *trackingPoolNode) get() {
}

func (n *trackingPoolNode) put() {
}

type TrackingPool[T any] struct {
	o Pool[T]
	k func(T) uintptr

	mtx sync.Mutex

	n sync.Pool
	m map[uintptr]*trackingPoolNode
}

func NewTrackingPool[T any](o Pool[T], k func(T) uintptr) *TrackingPool[T] {
	p := &TrackingPool[T]{
		o: o,
		k: k,

		m: make(map[uintptr]*trackingPoolNode),
	}

	p.n.New = func() any {
		return p.put(o.Get(), true)
	}

	return p
}

func (p *TrackingPool[T]) put(v T, isNew bool) *trackingPoolNode {
	k := p.k(v)

	p.mtx.Lock()
	defer p.mtx.Unlock()

	if e := p.m[k]; e != nil {
		if isNew {
			panic("key collision")
		}

		p.n.Put(e)
		e.put()
		return e
	}

	r := &trackingPoolNode{
		v: v,
		k: k,
	}

	p.m[k] = r
	r.put()
	return r
}

var _ Pool[int] = &TrackingPool[int]{}

func (p *TrackingPool[T]) Get() T {
	n := p.n.Get().(*trackingPoolNode)
	n.get()
	return n.v.(T)
}

func (p *TrackingPool[T]) Put(v T) {
	p.put(v, false)
}
