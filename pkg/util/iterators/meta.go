package iterators

import (
	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/slices"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type mapIterator[F, T any] struct {
	it Iterator[F]
	fn func(f F) T
}

func Map[F, T any](it Iterable[F], fn func(f F) T) Iterable[T] {
	check.NotNil(it)
	return Factory(func() Iterator[T] {
		return &mapIterator[F, T]{it: it.Iterate(), fn: fn}
	}, it)

}

var _ Iterator[string] = &mapIterator[int, string]{}

func (i *mapIterator[F, T]) Iterate() Iterator[T] {
	return i
}

func (i *mapIterator[F, T]) HasNext() bool {
	return i.it.HasNext()
}

func (i *mapIterator[F, T]) Next() T {
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
	check.NotNil(it)
	return Factory(func() Iterator[T] {
		return &filterIterator[T]{it: it.Iterate(), fn: fn}
	}, it)
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
			panic(bt.IteratorExhaustedError{})
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
	return Factory(func() Iterator[bt.Kv[int, T]] {
		return &enumerateIterator[T]{i: it.Iterate()}
	}, it)
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

type zipIterator[L, R any] struct {
	l Iterator[L]
	r Iterator[R]
}

func Zip[L, R any](l Iterable[L], r Iterable[R]) Iterable[bt.Pair[L, R]] {
	return Factory(func() Iterator[bt.Pair[L, R]] {
		return &zipIterator[L, R]{l: l.Iterate(), r: r.Iterate()}
	}, bt.PairOf(l, r))
}

var _ Iterator[bt.Pair[int, string]] = &zipIterator[int, string]{}

func (i *zipIterator[L, R]) Iterate() Iterator[bt.Pair[L, R]] {
	return i
}

func (i *zipIterator[L, R]) HasNext() bool {
	return i.l.HasNext() && i.r.HasNext()
}

func (i *zipIterator[L, R]) Next() bt.Pair[L, R] {
	return bt.PairOf(i.l.Next(), i.r.Next())
}

//

func Kvs[K comparable, V any](it Iterable[bt.Pair[K, V]]) Iterable[bt.Kv[K, V]] {
	return Map(it, func(p bt.Pair[K, V]) bt.Kv[K, V] {
		return bt.KvOf(p.L, p.R)
	})
}

//

type chunkSharedIterator[T any] struct {
	i Iterator[T]
	n int
	s []T
}

func ChunkShared[T any](it Iterable[T], n int) Iterable[[]T] {
	return Factory(func() Iterator[[]T] {
		return &chunkSharedIterator[T]{i: it.Iterate(), n: n}
	}, it)
}

var _ Iterator[[]any] = &chunkSharedIterator[any]{}

func (i *chunkSharedIterator[T]) Iterate() Iterator[[]T] {
	return i
}

func (i chunkSharedIterator[T]) HasNext() bool {
	return i.i.HasNext()
}

func (i *chunkSharedIterator[T]) Next() []T {
	s := i.s
	if s != nil {
		s = s[:0]
	}
	n := 0
	for n < i.n && i.i.HasNext() {
		s = append(s, i.i.Next())
		n++
	}
	i.s = s
	return s
}

//

func Chunk[T any](it Iterable[T], n int) Iterable[[]T] {
	return Map(ChunkShared(it, n), func(s []T) []T { return slices.Clone(s) })
}

//

type flattenIterator[T any] struct {
	it Iterator[Iterable[T]]
	ci Iterator[T]
}

func Flatten[T any](it Iterable[Iterable[T]]) Iterable[T] {
	return Factory(func() Iterator[T] {
		return &flattenIterator[T]{
			it: it.Iterate(),
		}
	}, it)
}

var _ Iterator[int] = &flattenIterator[int]{}

func (i *flattenIterator[T]) Iterate() Iterator[T] {
	return i
}

func (i *flattenIterator[T]) HasNext() bool {
	if i.ci != nil {
		if i.ci.HasNext() {
			return true
		}
		i.ci = nil
	}

	for {
		if !i.it.HasNext() {
			return false
		}

		ci := i.it.Next().Iterate()
		if ci.HasNext() {
			i.ci = ci
			return true
		}
	}
}

func (i *flattenIterator[T]) Next() T {
	if !i.HasNext() {
		panic(bt.IteratorExhaustedError{})
	}
	return i.ci.Next()
}

//

func FlatMap[F, T any](it Iterable[F], fn func(f F) Iterable[T]) Iterable[T] {
	return Flatten(Map(it, fn))
}

func FlatSliceMap[F, T any](it Iterable[F], fn func(f F) []T) Iterable[T] {
	return Flatten(Map(Map(it, fn), func(s []T) Iterable[T] { return OfSlice(s) }))
}
