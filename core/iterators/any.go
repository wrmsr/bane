package iterators

import (
	bt "github.com/wrmsr/bane/core/types"
)

type AnyIterable interface {
	AnyIterate() bt.Iterator[any]
}

func AsAny[T any](it bt.Iterable[T]) bt.Iterable[any] {
	return Map(it, func(e T) any { return e })
}
