package iterators

import bt "github.com/wrmsr/bane/pkg/utils/types"

type enumeratingIterator[T any] struct {
	i Iterator[T]
	c int
}

func Enumerate[T any](it Iterator[T]) Iterator[bt.Kv[int, T]] {
	return enumeratingIterator[T]{i: it}
}

var _ Iterator[bt.Kv[int, any]] = enumeratingIterator[any]{}

func (i enumeratingIterator[T]) HasNext() bool {
	return i.i.HasNext()
}

func (i enumeratingIterator[T]) Next() bt.Kv[int, T] {
	r := bt.Kv[int, T]{i.c, i.i.Next()}
	i.c++
	return r
}
