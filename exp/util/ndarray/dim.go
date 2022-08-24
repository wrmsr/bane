package ndarray

import (
	"github.com/wrmsr/bane/pkg/util/check"
	opt "github.com/wrmsr/bane/pkg/util/optional"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type Dim = int

//

type DimRange bt.Range[Dim]

func (r DimRange) Check(l Dim) DimRange {
	check.NotEqual(r.Step, 0)
	if r.Step > 0 {
		check.Condition(r.Start >= 0)
		check.Condition(r.Stop < l)
		check.Condition(r.Stop >= r.Start)
	} else {
		check.Condition(r.Start < l)
		check.Condition(r.Stop >= 0)
		check.Condition(r.Stop <= r.Start)
	}
	return r
}

func (r DimRange) Scalar() opt.Optional[Dim] {
	if r.Step > 0 {
		if r.Stop == r.Start+1 {
			return opt.Just(r.Start)
		}
	} else {
		if r.Start == r.Stop+1 {
			return opt.Just(r.Stop)
		}
	}
	return opt.None[Dim]()
}
