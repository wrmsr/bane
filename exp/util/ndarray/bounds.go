package ndarray

import (
	"github.com/wrmsr/bane/pkg/util/check"
	opt "github.com/wrmsr/bane/pkg/util/optional"
)

type Bound struct {
	Start int
	Stop  int
	Step  int
}

func AsBound(o ...any) opt.Optional[Bound] {
	if len(o) < 1 {
		return opt.None[Bound]()
	}

	if len(o) > 1 {
		check.Condition(len(o) < 4)
		var step int
		if len(o) > 2 {
			step = o[2].(int)
		}
		return opt.Just(Bound{
			Start: o[0].(int),
			Stop:  o[1].(int),
			Step:  step,
		})
	}

	switch o := check.Single(o).(type) {
	case Bound:
		return opt.Just(o)
	case int:
		return opt.Just(Bound{Start: o})
	}

	panic(o)
}
