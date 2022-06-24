package iterators

type sliceIterator[T any] struct {
	s []T
	i int
}

func OfSlice[T any](s []T) Iterable[T] {
	return Factory[T](func() Iterator[T] {
		return &sliceIterator[T]{s: s}
	}, s)
}

func Of[T any](vs ...T) Iterable[T] {
	return &sliceIterator[T]{s: vs}
}

var _ Iterator[any] = &sliceIterator[any]{}

func (i *sliceIterator[T]) Iterate() Iterator[T] {
	return i
}

func (i sliceIterator[T]) HasNext() bool {
	return i.i < len(i.s)
}

func (i *sliceIterator[T]) Next() T {
	if i.i >= len(i.s) {
		panic(IteratorExhaustedError{})
	}
	v := i.s[i.i]
	i.i++
	return v
}
