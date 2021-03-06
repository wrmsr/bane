package container

import (
	its "github.com/wrmsr/bane/pkg/util/iterators"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type StdSet[T comparable] struct {
	m map[T]struct{}
}

func NewStdSet[T comparable](it bt.Iterable[T]) StdSet[T] {
	s := StdSet[T]{
		m: make(map[T]struct{}),
	}
	if it != nil {
		for it := it.Iterate(); it.HasNext(); {
			s.m[it.Next()] = struct{}{}
		}
	}
	return s
}

func NewStdSetOf[T comparable](vs ...T) StdSet[T] {
	return NewStdSet(its.Of(vs...))
}

var _ Set[int] = StdSet[int]{}

func (s StdSet[T]) Len() int {
	return len(s.m)
}

func (s StdSet[T]) Contains(t T) bool {
	_, ok := s.m[t]
	return ok
}

func (s StdSet[T]) Iterate() bt.Iterator[T] {
	return its.Map(
		its.OfMap(s.m),
		func(kv bt.Kv[T, struct{}]) T {
			return kv.K
		},
	).Iterate()
}

func (s StdSet[T]) ForEach(fn func(T) bool) bool {
	for v := range s.m {
		if !fn(v) {
			return false
		}
	}
	return true
}

//

type MutStdSet[T comparable] struct {
	s StdSet[T]
}

func NewMutStdSet[T comparable](it bt.Iterable[T]) MutStdSet[T] {
	return MutStdSet[T]{s: NewStdSet(it)}
}

func NewMutStdSetOf[T comparable](vs ...T) MutStdSet[T] {
	return NewMutStdSet(its.Of(vs...))
}

var _ MutSet[int] = MutStdSet[int]{}

func (s MutStdSet[T]) isMutable() {}

func (s MutStdSet[T]) Len() int                       { return s.s.Len() }
func (s MutStdSet[T]) Contains(v T) bool              { return s.s.Contains(v) }
func (s MutStdSet[T]) Iterate() bt.Iterator[T]        { return s.s.Iterate() }
func (s MutStdSet[T]) ForEach(fn func(v T) bool) bool { return s.s.ForEach(fn) }

func (s MutStdSet[T]) Add(value T) {
	s.s.m[value] = struct{}{}
}

func (s MutStdSet[T]) TryAdd(value T) bool {
	if s.Contains(value) {
		return false
	}
	s.s.m[value] = struct{}{}
	return true
}

func (s MutStdSet[T]) Remove(value T) {
	delete(s.s.m, value)
}

func (s MutStdSet[T]) TryRemove(value T) bool {
	if s.Contains(value) {
		return false
	}
	delete(s.s.m, value)
	return true
}

var _ Decay[Set[int]] = &MutStdSet[int]{}

func (s *MutStdSet[T]) Decay() Set[T] { return s.s }
