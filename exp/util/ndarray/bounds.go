package ndarray

import (
	"github.com/wrmsr/bane/pkg/util/check"
	opt "github.com/wrmsr/bane/pkg/util/optional"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

type Bound struct {
	Start, Stop, Step opt.Optional[Dim]
}

func AsBound(o ...any) Bound {
	if len(o) < 1 {
		return Bound{}
	}

	cd := func(o any) opt.Optional[Dim] {
		if o == nil {
			return opt.None[Dim]()
		}

		switch o := o.(type) {
		case opt.Optional[Dim]:
			return o

		case Dim:
			return opt.Just[Dim](o)
		}

		panic(o)
	}

	if len(o) == 1 {
		if o, ok := o[0].(Bound); ok {
			return o
		}
	}

	a := func() bt.AnySlice {
		if len(o) == 1 {
			if o, ok := o[0].([]opt.Optional[Dim]); ok {
				return bt.AnySliceOf(o)
			}

			if o, ok := o[0].([]Dim); ok {
				return bt.AnySliceOf(o)
			}

			if o, ok := o[0].([]any); ok {
				return bt.AnySliceOf(o)
			}
		}

		return bt.AnySliceOf(o)
	}()

	check.Condition(a.Len() < 4)

	cdi := func(i int) opt.Optional[Dim] {
		if i >= a.Len() {
			return opt.None[Dim]()
		}

		return cd(a.Get(i))
	}

	return Bound{
		Start: cdi(0),
		Stop:  cdi(1),
		Step:  cdi(2),
	}
}
