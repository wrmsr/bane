package containers

import (
	its "github.com/wrmsr/bane/pkg/utils/iterators"
	bt "github.com/wrmsr/bane/pkg/utils/types"
)

//

type Set[T comparable] interface {
	Len() int
	Contains(T) bool

	its.Iterable[T]
	its.Traversable[T]
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

func newSetImpl[T comparable](it its.Iterable[T]) setImpl[T] {
	s := setImpl[T]{
		m: make(map[T]struct{}),
	}
	for i := it.Iterate(); i.HasNext(); {
		s.m[i.Next()] = struct{}{}
	}
	return s
}

func NewSet[T comparable](it its.Iterable[T]) Set[T] {
	return newSetImpl(it)
}

func NewSetOf[T comparable](vs ...T) Set[T] {
	return NewSet(its.Of(vs...))
}

var _ Set[int] = setImpl[int]{}

func (s setImpl[T]) Len() int {
	return len(s.m)
}

func (s setImpl[T]) Contains(t T) bool {
	_, ok := s.m[t]
	return ok
}

func (s setImpl[T]) Iterate() its.Iterator[T] {
	return its.Transform(
		its.Map(s.m),
		func(kv bt.Kv[T, struct{}]) T {
			return kv.K
		},
	).Iterate()
}

func (s setImpl[T]) ForEach(fn func(T) bool) bool {
	for v := range s.m {
		if !fn(v) {
			return false
		}
	}
	return true
}

//

type mutSetImpl[T comparable] struct {
	setImpl[T]
}

func NewMutSet[T comparable](it its.Iterable[T]) MutSet[T] {
	return mutSetImpl[T]{newSetImpl(it)}
}

func NewMutSetOf[T comparable](vs ...T) MutSet[T] {
	return NewMutSet(its.Of(vs...))
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
