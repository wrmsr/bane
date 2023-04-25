package iterators

import (
	bt "github.com/wrmsr/bane/core/types"
)

type emptyIterator[T any] struct{}

func Empty[T any]() bt.Iterable[T] {
	return emptyIterator[T]{}
}

var _ bt.Iterator[int] = emptyIterator[int]{}

func (i emptyIterator[T]) Iterate() bt.Iterator[T] { return i }
func (i emptyIterator[T]) HasNext() bool           { return false }
func (i emptyIterator[T]) Next() T                 { panic(bt.IteratorExhaustedError{}) }
