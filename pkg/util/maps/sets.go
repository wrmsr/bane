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

func (s Set[T]) Update(o Set[T]) {
	for v := range o {
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

func (s Set[T]) Equals(o Set[T]) bool {
	if len(s) != len(o) {
		return false
	}
	for v := range o {
		if !s.Contains(v) {
			return false
		}
	}
	return true
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

type TypeSet map[reflect.Type]struct{}

func NewTypeSet(s []reflect.Type) TypeSet {
	m := make(map[reflect.Type]struct{}, len(s))
	for _, ty := range s {
		m[ty] = struct{}{}
	}
	return m
}
