package iterators

type chunkingIterator[T any] struct {
	i Iterator[T]
	n int
	s []T
}

func Chunk[T any](i Iterator[T], n int) Iterator[[]T] {
	return chunkingIterator[T]{
		i: i,
		n: n,
	}
}

var _ Iterator[[]any] = chunkingIterator[any]{}

func (i chunkingIterator[T]) HasNext() bool {
	return i.i.HasNext()
}

func (i chunkingIterator[T]) Next() []T {
	s := i.s
	if s != nil {
		s = s[:0]
	}
	for i.i.HasNext() {
		s = append(s, i.i.Next())
	}
	i.s = s
	return s
}
