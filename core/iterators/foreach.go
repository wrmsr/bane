package iterators

import (
	bt "github.com/wrmsr/bane/core/types"
)

//

func ForEach[T any](it bt.Iterable[T], fn func(v T) bool) bool {
	for it := it.Iterate(); it.HasNext(); {
		if !fn(it.Next()) {
			return false
		}
	}
	return true
}

func ForAll[T any](tv bt.Traversable[T], fn func(v T)) {
	tv.ForEach(func(v T) bool {
		fn(v)
		return true
	})
}

func SeqForEach[T any](t bt.Traversable[T]) []T {
	var s []T
	t.ForEach(func(v T) bool {
		s = append(s, v)
		return true
	})
	return s
}

//

type traversableIterable[T any] struct {
	bt.Iterable[T]
}

func AsTraversable[T any](it bt.Iterable[T]) bt.Traversable[T] {
	return traversableIterable[T]{it}
}

var _ bt.Traversable[any] = traversableIterable[any]{}

func (t traversableIterable[T]) ForEach(fn func(v T) bool) bool {
	return ForEach[T](t, fn)
}

//

type fnTraversable[T any] struct {
	fn func(fn func(v T) bool) bool
}

func TraversableOf[T any](fn func(fn func(v T) bool) bool) bt.Traversable[T] {
	return fnTraversable[T]{fn: fn}
}

var _ bt.Traversable[any] = fnTraversable[any]{}

func (f fnTraversable[T]) ForEach(fn func(v T) bool) bool {
	return f.fn(fn)
}
