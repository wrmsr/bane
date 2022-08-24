package iterators

import (
	"context"

	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type chanIterator[T any] struct {
	ch   <-chan T
	done <-chan struct{}
	nxt  bt.Optional[T]
}

func OfChan[T any](ch <-chan T) bt.Iterable[T] {
	return Factory[T](func() bt.Iterator[T] {
		return &chanIterator[T]{ch: ch}
	}, ch)
}

func OfChanDone[T any](ch <-chan T, done <-chan struct{}) bt.Iterable[T] {
	return Factory[T](func() bt.Iterator[T] {
		return &chanIterator[T]{ch: ch, done: done}
	}, ch)
}

func OfChanContext[T any](ctx context.Context, ch <-chan T) bt.Iterable[T] {
	return OfChanDone[T](ch, ctx.Done())
}

var _ bt.Iterator[int] = &chanIterator[int]{}

func (i *chanIterator[T]) Iterate() bt.Iterator[T] {
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
		i.nxt = bt.Just(v)
		return true
	}
}

func (i *chanIterator[T]) Next() T {
	if !i.HasNext() {
		panic(bt.IteratorExhaustedError{})
	}
	v := i.nxt.Value()
	i.nxt = bt.None[T]()
	return v
}

//

func ToChanDone[T any](it bt.Iterable[T], done <-chan struct{}) <-chan T {
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
