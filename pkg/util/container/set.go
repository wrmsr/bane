package container

import (
	"reflect"

	its "github.com/wrmsr/bane/pkg/util/iterators"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type Set[T any] interface {
	Len() int
	Contains(v T) bool

	its.Iterable[T]
	its.Traversable[T]

	rfl.TypeArgsReflector
}

type MutSet[T any] interface {
	Set[T]
	Mutable

	Add(v T)
	TryAdd(v T) bool
	Remove(v T)
	TryRemove(v T) bool
}

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

func (s setImpl[T]) ReflectTypeArgs() []reflect.Type {
	return []reflect.Type{rfl.TypeOf[T]()}
}

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

func (m mutSetImpl[T]) isMutable() {}

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
