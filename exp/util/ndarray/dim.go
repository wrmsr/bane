package ndarray

import (
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type Dim = int

//

type DimRange bt.Range[Dim]

func (r DimRange) Range() bt.Range[Dim] {
	return bt.Range[Dim](r)
}

func (r DimRange) Check(l Dim) DimRange {
	r.Range().CheckNormal(l)
	return r
}

func (r DimRange) Scalar() bt.Optional[Dim] {
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
