package iterators

import bt "github.com/wrmsr/bane/pkg/util/types"

type sliceIterator[T any] struct {
	s []T
	r bt.Range[int]
}

func OfSlice[T any](s []T) Iterable[T] {
	return Factory[T](func() Iterator[T] {
		return &sliceIterator[T]{s: s, r: bt.RangeTo(len(s))}
	}, s)
}

func Of[T any](vs ...T) Iterable[T] {
	return &sliceIterator[T]{s: vs, r: bt.RangeTo(len(vs))}
}

func OfSliceRange[T any](s []T, r bt.Range[int]) Iterable[T] {
	return Factory[T](func() Iterator[T] {
		return &sliceIterator[T]{s: s, r: r}
	}, s)
}

var _ Iterator[any] = &sliceIterator[any]{}

func (i *sliceIterator[T]) Iterate() Iterator[T] {
	return i
}

func (i sliceIterator[T]) HasNext() bool {
	return i.r.HasNext()
}

func (i *sliceIterator[T]) Next() T {
	c := i.r.Start
	n, ok := i.r.Next()
	if !ok {
		panic(bt.IteratorExhaustedError{})
	}
	v := i.s[c]
	i.r = n
	return v
}
