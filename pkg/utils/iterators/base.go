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
