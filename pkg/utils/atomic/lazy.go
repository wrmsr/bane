package atomic

import (
	"sync"
	"sync/atomic"
)

type box[T any] struct {
	v T
}

type Lazy[T any] struct {
	fn func() T

	o sync.Once
	v atomic.Value
}

func NewLazy[T any](fn func() T) *Lazy[T] {
	return &Lazy[T]{
		fn: fn,
	}
}

func (l *Lazy[T]) Peek() T {
	if b := l.v.Load(); b != nil {
		return b.(box[T]).v
	}
	var z T
	return z
}

func (l *Lazy[T]) Get() T {
	l.o.Do(func() {
		l.Set(l.fn())
	})
	return l.Peek()
}

func (l *Lazy[T]) Set(v T) T {
	l.v.Store(box[T]{v})
	return v
}
