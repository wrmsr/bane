package container

import (
	its "github.com/wrmsr/bane/pkg/util/iterators"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type setImpl[T comparable] struct {
	m map[T]struct{}
}

func newSetImpl[T comparable](it its.Iterable[T]) setImpl[T] {
	s := setImpl[T]{
		m: make(map[T]struct{}),
	}
	if it != nil {
		for it := it.Iterate(); it.HasNext(); {
			s.m[it.Next()] = struct{}{}
		}
	}
	return s
}

func NewSet[T comparable](it its.Iterable[T]) Set[T] {
	if it, ok := it.(Set[T]); ok {
		return it
	}
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
	return its.Map(
		its.OfMap(s.m),
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

func (s mutSetImpl[T]) isMutable()    {}
func (s mutSetImpl[T]) Decay() Set[T] { return s.setImpl }

func (s mutSetImpl[T]) Add(value T) {
	s.m[value] = struct{}{}
}

func (s mutSetImpl[T]) TryAdd(value T) bool {
	if s.Contains(value) {
		return false
	}
	s.m[value] = struct{}{}
	return true
}

func (s mutSetImpl[T]) Remove(value T) {
	delete(s.m, value)
}

func (s mutSetImpl[T]) TryRemove(value T) bool {
	if s.Contains(value) {
		return false
	}
	delete(s.m, value)
	return true
}
