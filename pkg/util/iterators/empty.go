package iterators

type emptyIterator[T any] struct{}

func Empty[T any]() Iterator[T] {
	return emptyIterator[T]{}
}

var _ Iterator[int] = emptyIterator[int]{}

func (i emptyIterator[T]) Iterate() Iterator[T] { return i }
func (i emptyIterator[T]) HasNext() bool        { return false }
func (i emptyIterator[T]) Next() T              { panic(bt.IteratorExhaustedError{}) }
