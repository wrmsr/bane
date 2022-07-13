package iterators

type sliceIterator[T any] struct {
	s []T
	i int
	x int
}

func OfSlice[T any](s []T) Iterable[T] {
	return Factory[T](func() Iterator[T] {
		return &sliceIterator[T]{s: s, x: 1}
	}, s)
}

func Of[T any](vs ...T) Iterable[T] {
	return &sliceIterator[T]{s: vs, x: 1}
}

func OfSliceRange[T any](s []T, start, stop, step int) Iterable[T] {
	return Factory[T](func() Iterator[T] {
		return &sliceIterator[T]{s: s, x: step}
	}, s)
}

var _ Iterator[any] = &sliceIterator[any]{}

func (i *sliceIterator[T]) Iterate() Iterator[T] {
	return i
}

func (i sliceIterator[T]) HasNext() bool {
	return (i.i + i.x) <= len(i.s)
}

func (i *sliceIterator[T]) Next() T {
	if i.i >= len(i.s) {
		panic(IteratorExhaustedError{})
	}
	v := i.s[i.i]
	i.i += i.x
	return v
}
