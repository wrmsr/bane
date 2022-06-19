package iterators

type SliceIterator[T any] struct {
	s []T
	i int
}

func IterateSlice[T any](s []T) Iterator[T] {
	return &SliceIterator[T]{s: s}
}

func Of[T any](vs ...T) CanIterate[T] {
	return &SliceIterator[T]{s: vs}
}

var _ Iterator[any] = &SliceIterator[any]{}

func (i *SliceIterator[T]) Iterate() Iterator[T] {
	return i
}

func (i SliceIterator[T]) HasNext() bool {
	return i.i < len(i.s)
}

func (i *SliceIterator[T]) Next() T {
	if i.i >= len(i.s) {
		panic(IteratorExhaustedError{})
	}
	v := i.s[i.i]
	i.i++
	return v
}
