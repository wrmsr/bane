package iterators

import bt "github.com/wrmsr/bane/pkg/util/types"

//

type transformIterator[F, T any] struct {
	it Iterator[F]
	fn func(f F) T
}

func Transform[F, T any](it Iterable[F], fn func(f F) T) Iterable[T] {
	return factory(func() Iterator[T] { return &transformIterator[F, T]{it: it.Iterate(), fn: fn} })

}

var _ Iterator[string] = &transformIterator[int, string]{}

func (i *transformIterator[F, T]) Iterate() Iterator[T] {
	return i
}

func (i *transformIterator[F, T]) HasNext() bool {
	return i.it.HasNext()
}

func (i *transformIterator[F, T]) Next() T {
	return i.fn(i.it.Next())
}

//

type filterIterator[T any] struct {
	it Iterator[T]
	fn func(f T) bool

	p bool
	v T
}

func Filter[T any](it Iterable[T], fn func(v T) bool) Iterable[T] {
	return factory(func() Iterator[T] { return &filterIterator[T]{it: it.Iterate(), fn: fn} })
}

var _ Iterator[int] = &filterIterator[int]{}

func (i *filterIterator[T]) Iterate() Iterator[T] {
	return i
}

func (i *filterIterator[T]) HasNext() bool {
	if !i.p {
		for i.it.HasNext() {
			v := i.it.Next()
			if i.fn(v) {
				i.p = true
				i.v = v
				break
			}
		}
	}
	return i.p
}

func (i *filterIterator[T]) Next() T {
	if !i.p {
		if !i.HasNext() {
			panic(IteratorExhaustedError{})
		}
	}
	v := i.v
	var z T
	i.v = z
	i.p = false
	return v
}

//

type enumerateIterator[T any] struct {
	i Iterator[T]
	c int
}

func Enumerate[T any](it Iterable[T]) Iterable[bt.Kv[int, T]] {
	return factory(func() Iterator[bt.Kv[int, T]] { return &enumerateIterator[T]{i: it.Iterate()} })
}

var _ Iterator[bt.Kv[int, any]] = &enumerateIterator[any]{}

func (i *enumerateIterator[T]) Iterate() Iterator[bt.Kv[int, T]] {
	return i
}

func (i enumerateIterator[T]) HasNext() bool {
	return i.i.HasNext()
}

func (i *enumerateIterator[T]) Next() bt.Kv[int, T] {
	r := bt.Kv[int, T]{i.c, i.i.Next()}
	i.c++
	return r
}

//

type chunkIterator[T any] struct {
	i Iterator[T]
	n int
	s []T
}

func Chunk[T any](i Iterable[T], n int) Iterable[[]T] {
	return factory(func() Iterator[[]T] { return &chunkIterator[T]{i: i.Iterate(), n: n} })
}

var _ Iterator[[]any] = &chunkIterator[any]{}

func (i *chunkIterator[T]) Iterate() Iterator[[]T] {
	return i
}

func (i chunkIterator[T]) HasNext() bool {
	return i.i.HasNext()
}

func (i *chunkIterator[T]) Next() []T {
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
