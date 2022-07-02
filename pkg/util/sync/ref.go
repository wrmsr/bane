package sync

import (
	"sync/atomic"

	opt "github.com/wrmsr/bane/pkg/util/optional"
)

type Ref[T any] struct {
	v opt.Optional[T]
	c int32

	cbs []func(r *Ref[T], v T)
}

func NewWeakRef[T any](v T) *Ref[T] {
	return &Ref[T]{v: opt.Just(v), c: 1}
}

func (r *Ref[T]) Get() opt.Optional[T] {
	return r.v
}

func (r *Ref[T]) AddCallback(cb func(r *Ref[T], v T)) {
	r.cbs = append(r.cbs, cb)
}

func (r *Ref[T]) Acquire() {
	atomic.AddInt32(&r.c, 1)
}

func (r *Ref[T]) Release() {
	c := atomic.AddInt32(&r.c, -1)
	if c < 0 {
		panic("negative refs")
	}

	if c == 0 {
		if !r.v.Present() {
			panic("broken ref")
		}

		v := r.v.Value()
		for _, cb := range r.cbs {
			cb(r, v)
		}

		if c < 0 {
			panic("negative refs")
		}

		if c == 0 {
			r.v = opt.None[T]()
		}
	}
}
