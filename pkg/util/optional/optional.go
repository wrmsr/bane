package optional

import (
	its "github.com/wrmsr/bane/pkg/util/iterators"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type NoValueError struct{}

func (e NoValueError) Error() string {
	return "no value"
}

//

type Optional[T any] struct {
	v T
	p bool
}

func None[T any]() Optional[T] {
	return Optional[T]{}
}

func Just[T any](v T) Optional[T] {
	return Optional[T]{v: v, p: true}
}

//

func (o Optional[T]) Present() bool {
	return o.p
}

func (o Optional[T]) IfPresent(fn func()) {
	if o.p {
		fn()
	}
}

func (o Optional[T]) IfAbsent(fn func()) {
	if !o.p {
		fn()
	}
}

func (o Optional[T]) Value() T {
	if !o.p {
		panic(NoValueError{})
	}
	return o.v
}

func (o Optional[T]) OrDefault(d T) T {
	if !o.p {
		return d
	}
	return o.v
}

func (o Optional[T]) OrZero() T {
	return o.OrDefault(bt.Zero[T]())
}

//

type Interface interface {
	Present() bool
	Interface() any
	Replace(any) Interface
}

var _ Interface = Optional[int]{}

func (o Optional[T]) Interface() any {
	return o.v
}

func (o Optional[T]) Replace(v any) Interface {
	return Just(v.(T))
}

//

var _ its.Iterable[any] = Optional[any]{}

func (o Optional[T]) Iterate() its.Iterator[T] {
	if o.Present() {
		return its.Of[T](o.Value()).Iterate()
	}
	return its.Of[T]().Iterate()
}

var _ its.Traversable[any] = Optional[any]{}

func (o Optional[T]) ForEach(fn func(v T) bool) bool {
	if o.Present() {
		if !fn(o.Value()) {
			return false
		}
	}
	return true
}

//

func SetIfAbsent[T any](o *Optional[T], fn func() T) T {
	if o.Present() {
		return o.Value()
	}
	r := fn()
	*o = Just(r)
	return r
}

//

func HashEq[T any](he bt.HashEqImpl[T]) bt.HashEqImpl[Optional[T]] {
	return bt.HashEqOf[Optional[T]](
		func(v Optional[T]) uintptr {
			if v.Present() {
				return he.Hash(v.Value())
			}
			return uintptr(0)
		},
		func(l, r Optional[T]) bool {
			if l.Present() && !r.Present() {
				return he.Eq(l.Value(), r.Value())
			}
			return !(l.Present() || !r.Present())
		},
	)
}
