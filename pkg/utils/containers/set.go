package containers

import bit "github.com/wrmsr/bane/pkg/utils/iterators"

//

type Set[T comparable] interface {
	Contains(T) bool

	Iterate() bit.Iterator[T]
}

type MutSet[T comparable] interface {
	Set[T]

	Add(value T)
	TryAdd(value T) bool
	Remove(value T)
	TryRemove(value T) bool
}

//

type setImpl[T comparable] struct {
	m map[T]struct{}
}

func newSetImpl[T comparable]() setImpl[T] {
	return setImpl[T]{
		m: make(map[T]struct{}),
	}
}

func NewSet[T comparable]() Set[T] {
	return newSetImpl[T]()
}

var _ Set[int] = setImpl[int]{}

func (s setImpl[T]) Contains(t T) bool {
	_, ok := s.m[t]
	return ok
}

func (s setImpl[T]) Iterate() bit.Iterator[T] {
	panic("nyi")
}

//

type mutSetImpl[T comparable] struct {
	setImpl[T]
}

func NewMutSet[T comparable]() MutSet[T] {
	return mutSetImpl[T]{newSetImpl[T]()}
}

var _ MutSet[int] = mutSetImpl[int]{}

func (m mutSetImpl[T]) Add(value T) {
	m.m[value] = struct{}{}
}

func (m mutSetImpl[T]) TryAdd(value T) bool {
	if m.Contains(value) {
		return false
	}
	m.m[value] = struct{}{}
	return true
}

func (m mutSetImpl[T]) Remove(value T) {
	delete(m.m, value)
}

func (m mutSetImpl[T]) TryRemove(value T) bool {
	if m.Contains(value) {
		return false
	}
	delete(m.m, value)
	return true
}
