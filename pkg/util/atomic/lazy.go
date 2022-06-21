package atomic

import (
	"sync"
	"sync/atomic"
)

type box[T any] struct {
	v T
}

type Lazy[T any] struct {
	Fn func() T

	o sync.Once
	v atomic.Value
}

func NewLazy[T any](fn func() T) *Lazy[T] {
	return &Lazy[T]{
		Fn: fn,
	}
}

func (l *Lazy[T]) Peek() (T, bool) {
	if b := l.v.Load(); b != nil {
		return b.(box[T]).v, true
	}
	var z T
	return z, false
}

func (l *Lazy[T]) Get() T {
	l.o.Do(func() {
		if _, ok := l.Peek(); !ok {
			l.Set(l.Fn())
		}
	})
	v, ok := l.Peek()
	if !ok {
		panic("unreachable")
	}
	return v
}

func (l *Lazy[T]) Set(v T) T {
	l.v.Store(box[T]{v})
	return v
}
