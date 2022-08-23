package ndarray

import (
	opt "github.com/wrmsr/bane/pkg/util/optional"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type Dim = int

//

type DimRange bt.Range[Dim]

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
