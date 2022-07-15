package iterators

import (
	"context"

	opt "github.com/wrmsr/bane/pkg/util/optional"
)

//

type chanIterator[T any] struct {
	ch   <-chan T
	done <-chan struct{}
	nxt  opt.Optional[T]
}

func OfChan[T any](ch <-chan T) Iterable[T] {
	return Factory[T](func() Iterator[T] {
		return &chanIterator[T]{ch: ch}
	}, ch)
}

func OfChanDone[T any](ch <-chan T, done <-chan struct{}) Iterable[T] {
	return Factory[T](func() Iterator[T] {
		return &chanIterator[T]{ch: ch, done: done}
	}, ch)
}

func OfChanContext[T any](ctx context.Context, ch <-chan T) Iterable[T] {
	return OfChanDone[T](ch, ctx.Done())
}

var _ Iterator[int] = &chanIterator[int]{}

func (i *chanIterator[T]) Iterate() Iterator[T] {
	return i
}

func (i *chanIterator[T]) HasNext() bool {
	if i.nxt.Present() {
		return true
	}

	select {
	case <-i.done:
		return false
	case v, ok := <-i.ch:
		if !ok {
			return false
		}
		i.nxt = opt.Just(v)
		return true
	}
}

func (i *chanIterator[T]) Next() T {
	if !i.HasNext() {
		panic(bt.IteratorExhaustedError{})
	}
	v := i.nxt.Value()
	i.nxt = opt.None[T]()
	return v
}

//

func ToChanDone[T any](it Iterable[T], done <-chan struct{}) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for it := it.Iterate(); it.HasNext(); {
			select {
			case <-done:
				return
			case ch <- it.Next():
			}
		}
	}()
	return ch
}
