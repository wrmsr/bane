package iterators

type AnyIterable interface {
	AnyIterate() Iterator[any]
}

func AsAny[T any](it Iterable[T]) Iterable[any] {
	return Map(it, func(e T) any { return e })
}
