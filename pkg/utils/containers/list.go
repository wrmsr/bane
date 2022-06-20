package containers

import (
	its "github.com/wrmsr/bane/pkg/utils/iterators"
)

//

type List[T any] interface {
	Len() int
	Get(i int) T

	its.Iterable[T]
	its.Traversable[T]
}

type MutList[T any] interface {
	List[T]

	Append(v T)
	Delete(i int)
}

//

type listImpl[T any] struct {
	s []T
}

func newListImpl[T comparable](it its.Iterable[T]) listImpl[T] {
	s := listImpl[T]{}
	for i := it.Iterate(); i.HasNext(); {
		s.s = append(s.s, i.Next())
	}
	return s
}

func NewList[T comparable](it its.Iterable[T]) List[T] {
	return newListImpl(it)
}

func NewListOf[T comparable](vs ...T) List[T] {
	return NewList(its.Of(vs...))
}

var _ List[int] = listImpl[int]{}

func (l listImpl[T]) Len() int {
	return len(l.s)
}

func (l listImpl[T]) Get(i int) T {
	return l.s[i]
}

func (l listImpl[T]) Iterate() its.Iterator[T] {
	return its.Slice(l.s).Iterate()
}

func (l listImpl[T]) ForEach(fn func(v T) bool) bool {
	for _, v := range l.s {
		if !fn(v) {
			return false
		}
	}
	return true
}

//

type mutListImpl[T any] struct {
	listImpl[T]
}

func NewMutList[T comparable](it its.Iterable[T]) MutList[T] {
	return &mutListImpl[T]{newListImpl(it)}
}

func NewMutListOf[T comparable](vs ...T) MutList[T] {
	return NewMutList(its.Of(vs...))
}

var _ MutList[int] = &mutListImpl[int]{}

func (l *mutListImpl[T]) Append(v T) {
	l.s = append(l.s, v)
}

func (l *mutListImpl[T]) Delete(i int) {
	l.s = append(l.s[:i], l.s[i+1:]...)
}
