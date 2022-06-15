package iterators

type IteratorExhaustedError struct{}

func (e IteratorExhaustedError) Error() string {
	return "iterator exhausted"
}

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}
