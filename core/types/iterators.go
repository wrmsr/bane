package types

//

type Iterator[T any] interface {
	Iterable[T]

	HasNext() bool
	Next() T
}

type Iterable[T any] interface {
	Iterate() Iterator[T]
}

type IteratorExhaustedError struct{}

//

type AnyIterable interface {
	AnyIterate() Iterator[any]
}

type anyIterator[T any] struct {
	it Iterator[T]
}

var _ Iterator[any] = &anyIterator[int]{}

func (i *anyIterator[T]) Iterate() Iterator[any] { return i }
func (i *anyIterator[T]) HasNext() bool          { return i.it.HasNext() }
func (i *anyIterator[T]) Next() any              { return i.it.Next() }
