package ndarray

import (
	"github.com/wrmsr/bane/pkg/util/check"
	opt "github.com/wrmsr/bane/pkg/util/optional"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type Bound struct {
	Start, Stop, Step opt.Optional[Dim]
}

func (b Bound) Range(l Dim) DimRange {
	ci := func(o opt.Optional[Dim], d Dim) Dim {
		if !o.Present() {
			return d
		}
		i := o.Value()
		if i < 0 {
			i += l
			if i < 0 {
				i = 0
			}
		}
		if i > l {
			return l
		}
		return i
	}

	r := DimRange{
		Start: ci(b.Start, 0),
		Step:  b.Step.Or(1),
	}

	if b.Start.Present() && !b.Stop.Present() {
		if r.Step > 0 {
			r.Stop = r.Start + 1
		} else {
			r.Stop = r.Start - 1
		}
	} else {
		r.Stop = ci(b.Stop, l)
	}

	return r
}

//

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

		if o, ok := o[0].(DimRange); ok {
			return Bound{
				Start: opt.Just(o.Start),
				Stop:  opt.Just(o.Stop),
				Step:  opt.Just(o.Step),
			}
		}

		if o, ok := o[0].(bt.Range[Dim]); ok {
			return Bound{
				Start: opt.Just(o.Start),
				Stop:  opt.Just(o.Stop),
				Step:  opt.Just(o.Step),
			}
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
