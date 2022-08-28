package ndarray

import (
	"github.com/wrmsr/bane/pkg/util/check"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type Range struct {
	Start, Stop, Step Dim

	Squeeze bool
}

func (r Range) Range() bt.Range[Dim] {
	return bt.Range[Dim]{
		Start: r.Start,
		Stop:  r.Stop,
		Step:  r.Step,
	}
}

func (r Range) Check(l Dim) Range {
	r.Range().CheckNormal(l)
	return r
}

func (r Range) Scalar() bt.Optional[Dim] {
	if r.Step > 0 {
		if r.Stop == r.Start+1 {
			return bt.Just(r.Start)
		}
	} else {
		if r.Start == r.Stop+1 {
			return bt.Just(r.Stop)
		}
	}
	return bt.None[Dim]()
}

//

func IsSqueezedRange(o any) bool {
	if o, ok := o.(Range); ok {
		return o.Squeeze
	}

	if o, ok := AsDim(o); ok {
		return o.Present()
	}

	return false
}

func CalcRange(o any, l Dim) Range {
	if o == nil {
		return Range{
			Start: 0,
			Stop:  l,
			Step:  1,
		}
	}

	if o, ok := o.(Range); ok {
		return o
	}

	if o, ok := o.(bt.Range[Dim]); ok {
		return Range{
			Start: o.Start,
			Stop:  o.Stop,
			Step:  o.Step,
		}
	}

	clamp := func(o any, d Dim) Dim {
		v := check.Ok1(AsDim(o))
		if !v.Present() {
			return d
		}

		i := v.Value()
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

	if o, ok := AsDim(o); ok {
		return Range{
			Start: clamp(o, 0),
			Stop:  clamp(o.Map(func(d Dim) Dim { return d + 1 }), l),
			Step:  1,

			Squeeze: true,
		}
	}

	var al int
	switch o := o.(type) {
	case []Dim:
		al = len(o)
	case []any:
		al = len(o)
	default:
		panic(o)
	}

	check.Condition(al < 4)

	ga := func(i int) any {
		switch o := o.(type) {
		case []Dim:
			return o[i]
		case []any:
			return o[i]
		default:
			panic(o)
		}
	}

	aclamp := func(i int, d Dim) Dim {
		if i >= al {
			return d
		}
		return clamp(ga(i), d)
	}

	r := Range{
		Step: 1,
	}

	if al > 2 {
		if astep := ga(2); astep != nil {
			ostep := check.Ok1(AsDim(astep))
			if ostep.Present() {
				r.Step = ostep.Value()
			}
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
