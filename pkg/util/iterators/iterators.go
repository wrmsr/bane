package iterators

import (
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type factoryIterable[T any] struct {
	fn  func() bt.Iterator[T]
	src any
}

func Factory[T any](fn func() bt.Iterator[T], src any) bt.Iterable[T] {
	return factoryIterable[T]{fn: fn, src: src}
}

var _ bt.Iterable[int] = factoryIterable[int]{}

func (f factoryIterable[T]) Iterate() bt.Iterator[T] {
	return f.fn()
}
