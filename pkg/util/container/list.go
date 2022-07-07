package container

import (
	its "github.com/wrmsr/bane/pkg/util/iterators"
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
	Mutable
	//Decay[List[T]]

	Append(v T)
	Delete(i int)
}

//

type listImpl[T any] struct {
	s []T
}

func newListImpl[T any](it its.Iterable[T]) listImpl[T] {
	s := listImpl[T]{}
	if it != nil {
		for it := it.Iterate(); it.HasNext(); {
			s.s = append(s.s, it.Next())
		}
	}
	return s
}

func NewList[T any](it its.Iterable[T]) List[T] {
	if it, ok := it.(List[T]); ok {
		return it
	}
	return newListImpl(it)
}

func NewListOf[T any](vs ...T) List[T] {
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
	return its.OfSlice(l.s).Iterate()
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

func NewMutList[T any](it its.Iterable[T]) MutList[T] {
	return &mutListImpl[T]{newListImpl(it)}
}

func NewMutListOf[T any](vs ...T) MutList[T] {
	return NewMutList(its.Of(vs...))
}

func WrapSlice[T any](s []T) MutList[T] {
	return &mutListImpl[T]{listImpl[T]{s}}
}

var _ MutList[int] = &mutListImpl[int]{}

func (l *mutListImpl[T]) isMutable()     {}
func (l *mutListImpl[T]) Decay() List[T] { return l.listImpl }

func (l *mutListImpl[T]) Append(v T) {
	l.s = append(l.s, v)
}

func (l *mutListImpl[T]) Delete(i int) {
	l.s = append(l.s[:i], l.s[i+1:]...)
}
