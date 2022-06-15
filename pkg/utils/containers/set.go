package containers

import bit "github.com/wrmsr/bane/pkg/utils/iterators"

//

type Set[T comparable] interface {
	Contains(T) bool
	Iterate() bit.Iterator[T]
}

type MutSet[T comparable] interface {
	Set[T]

	Add(value T) bool
	Remove(value T) bool
}

//

type SetImpl[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() Set[T] {
	return SetImpl[T]{
		m: make(map[T]struct{}),
	}
}

var _ Set[any] = SetImpl[any]{}

func (s SetImpl[T]) Contains(t T) bool {
	_, ok := s.m[t]
	return ok
}

func (s SetImpl[T]) Iterate() bit.Iterator[T] {
	panic("nyi")
}
