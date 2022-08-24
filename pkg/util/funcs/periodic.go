package funcs

import (
	"time"

	bt "github.com/wrmsr/bane/pkg/util/types"
)

type PeriodicFn[T any] struct {
	fn func() T
	iv time.Duration

	clock func() time.Time

	dl time.Time
	v  bt.Optional[T]
}

func NewPeriodicFn[T any](fn func() T, iv time.Duration) PeriodicFn[T] {
	return PeriodicFn[T]{fn: fn, iv: iv, clock: time.Now}
}

func (pf PeriodicFn[T]) Call() T {
	var now time.Time

	rc := !pf.v.Present() || pf.dl.IsZero()
	if !rc {
		now = pf.clock()
		rc = now.After(pf.dl)
	}

	if rc {
		if now.IsZero() {
			now = pf.clock()
		}

		pf.v = bt.Just(pf.fn())
		pf.dl = now.Add(pf.iv)
	}

	return pf.v.Value()
}
