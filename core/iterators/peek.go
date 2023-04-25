package iterators

import bt "github.com/wrmsr/bane/core/types"

//

type PeekIterable[T any] interface {
	Iterate() PeekIterator[T]
}

type peekFactoryIterable[T any] struct {
	fn  func() PeekIterator[T]
	src any
}

var _ PeekIterable[any] = peekFactoryIterable[any]{}

func (f peekFactoryIterable[T]) Iterate() PeekIterator[T] {
	return f.fn()
}

//

type PeekIterator[T any] interface {
	bt.Iterator[T]
	Peek() T
}

type peekIterator[T any] struct {
	it bt.Iterator[T]
	pv T
	pb bool
}

func Peek[T any](it bt.Iterable[T]) PeekIterable[T] {
	return peekFactoryIterable[T]{
		fn: func() PeekIterator[T] {
			return &peekIterator[T]{it: it.Iterate()}
		},
		src: it,
	}
}

var _ PeekIterator[any] = &peekIterator[any]{}

func (i *peekIterator[T]) Iterate() bt.Iterator[T] {
	return i
}

func (i *peekIterator[T]) HasNext() bool {
	return i.pb || i.it.HasNext()
}

func (i *peekIterator[T]) Next() T {
	if i.pb {
		v := i.pv
		i.pb = false
		var z T
		i.pv = z
		return v
	}
	return i.it.Next()
}

func (i *peekIterator[T]) Peek() T {
	if !i.pb {
		i.pv = i.it.Next()
		i.pb = true
	}
	return i.pv
}
