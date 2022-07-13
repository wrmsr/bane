package types

type Iterator[T any] interface {
	Iterable[T]

	HasNext() bool
	Next() T
}

type Iterable[T any] interface {
	Iterate() Iterator[T]
}

type IteratorExhaustedError struct{}
