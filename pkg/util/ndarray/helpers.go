package ndarray

import bt "github.com/wrmsr/bane/pkg/util/types"

func OfRange[T bt.Rational](sh Shape) NdArray[T] {
	return Maker[T]{
		Shape: sh,
		Data:  bt.RangeTo[T](T(sh.Size())).Slice(),
	}.Make()
}
