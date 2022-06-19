package iterators

//

type IteratorExhaustedError struct{}

func (e IteratorExhaustedError) Error() string {
	return "iterator exhausted"
}

//

type Iterator[T any] interface {
	CanIterate[T]

	HasNext() bool
	Next() T
}

type CanIterate[T any] interface {
	Iterate() Iterator[T]
}
