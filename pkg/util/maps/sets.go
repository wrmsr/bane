package maps

import "reflect"

//

type Set[T comparable] map[T]struct{}

func MakeSet[T comparable]() Set[T] {
	return make(map[T]struct{})
}

func MakeSetOf[T comparable](vs []T) Set[T] {
	s := MakeSet[T]()
	s.AddAll(vs)
	return s
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) AddAll(vs []T) {
	for _, v := range vs {
		s[v] = struct{}{}
	}
}

func (s Set[T]) Remove(v T) {
	delete(s, v)
}

func (s Set[T]) RemoveAll(vs []T) {
	for _, v := range vs {
		delete(s, v)
	}
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Slice() []T {
	r := make([]T, len(s))
	i := 0
	for v := range s {
		r[i] = v
		i++
	}
	return r
}

//

type TypeSet 

func NewTypeSet(s []reflect.Type) map[reflect.Type]struct{} {
	m := make(map[reflect.Type]struct{}, len(s))
	for _, ty := range s {
		m[ty] = struct{}{}
	}
	return m
}
