package iterators

type IteratorExhaustedError struct{}

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}
