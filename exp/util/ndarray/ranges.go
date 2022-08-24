package ndarray

import (
	"github.com/wrmsr/bane/pkg/util/check"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

func CalcRange(o any, l Dim) DimRange {
	if o == nil {
		return DimRange{
			Start: 0,
			Stop:  l,
			Step:  1,
		}
	}

	if o, ok := o.(DimRange); ok {
		return o
	}

	if o, ok := o.(bt.Range[Dim]); ok {
		return DimRange(o)
	}

	clamp := func(o any, d Dim) Dim {
		if o == nil {
			return d
		}
		i := o.(Dim)
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

	if o, ok := o.(Dim); ok {
		return DimRange{
			Start: clamp(o, 0),
			Stop:  clamp(o+1, l),
			Step:  1,
		}
	}

	a := func() bt.AnySlice {
		if o, ok := o.([]Dim); ok {
			return bt.AnySliceOf(o)
		}

		if o, ok := o.([]any); ok {
			return bt.AnySliceOf(o)
		}

		panic(o)
	}()

	check.Condition(a.Len() < 4)

	aclamp := func(i int, d Dim) Dim {
		if i >= a.Len() {
			return d
		}
		return clamp(a.Get(i), d)
	}

	r := DimRange{
		Step: 1,
	}

	if a.Len() > 2 {
		if astep := a.Get(2); astep != nil {
			r.Step = astep.(Dim)
		}
	}

	check.Condition(r.Step != 0)
	if r.Step > 0 {
		r.Start = aclamp(0, 0)
		r.Stop = aclamp(1, l)
	} else {
		r.Start = aclamp(0, l)
		r.Stop = aclamp(1, 0)
	}

	return r.Check(l)
}
