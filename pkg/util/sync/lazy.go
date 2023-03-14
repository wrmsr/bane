package sync

import "sync"

type Lazy[T any] struct {
	once sync.Once

	v T
}

func (va *Lazy[T]) Set(v T) bool {
	ret := false
	va.once.Do(func() {
		va.v = v
		ret = true
	})
	return ret
}

func (va *Lazy[T]) Get(fn func() T) T {
	va.once.Do(func() {
		va.v = fn()
	})
	return va.v
}
