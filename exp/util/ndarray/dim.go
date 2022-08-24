package ndarray

import bt "github.com/wrmsr/bane/pkg/util/types"

//

type Dim int64

func AsDim(o any) (bt.Optional[Dim], bool) {
	if o == nil {
		return bt.None[Dim](), true
	}
	switch o := o.(type) {
	case bt.Optional[Dim]:
		return o, true
	case Dim:
		return bt.Just(o), true
	case int:
		return bt.Just(Dim(o)), true
	case int8:
		return bt.Just(Dim(o)), true
	case int16:
		return bt.Just(Dim(o)), true
	case int32:
		return bt.Just(Dim(o)), true
	case int64:
		return bt.Just(Dim(o)), true
	case uint:
		return bt.Just(Dim(o)), true
	case uint8:
		return bt.Just(Dim(o)), true
	case uint16:
		return bt.Just(Dim(o)), true
	case uint32:
		return bt.Just(Dim(o)), true
	}
	return bt.None[Dim](), false
}
