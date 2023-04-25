package match

import "reflect"

func Run(p Pattern, v any, cs *Captures) Match {
	if p.Previous() != nil {
		m := Run(p.Previous(), v, cs)
		if !m.Present() {
			return m
		}
		cs = m.Captures()
	}

	switch p := p.(type) {

	case CapturePattern:
		return PresentMatch{
			value:    v,
			captures: cs.AddAll(MakeCaptures(p.capture, v)),
		}

	case EqualsPattern:
		return PresentMatch{
			value:    v,
			captures: cs,
		}.Filter(func(a any) bool {
			return reflect.DeepEqual(a, p.value)
		})

	case FilterPattern:
		return PresentMatch{
			value:    v,
			captures: cs,
		}.Filter(p.fn)

	case TypePattern:
		return PresentMatch{
			value:    v,
			captures: cs,
		}.Filter(func(a any) bool {
			return reflect.TypeOf(a) == p.rtype
		})

	case WithPattern:
		pv := p.propPat.prop.fn(v)
		return Run(p.propPat.pat, pv, cs)

	}
	panic(p)
}
