package iterators

//

type IteratorExhaustedError struct{}

func (e IteratorExhaustedError) Error() string {
	return "iterator exhausted"
}

//

type Iterator[T any] interface {
	Iterable[T]

	HasNext() bool
	Next() T
}

type Iterable[T any] interface {
	Iterate() Iterator[T]
}

//

type factoryIterable[T any] struct {
	fn  func() Iterator[T]
	src any
}

func Factory[T any](fn func() Iterator[T], src any) Iterable[T] {
	return factoryIterable[T]{fn: fn, src: src}
}

var _ Iterable[int] = factoryIterable[int]{}

func (f factoryIterable[T]) Iterate() Iterator[T] {
	return f.fn()
}
