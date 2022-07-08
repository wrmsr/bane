package tensor

import (
	"fmt"
	"testing"
)

//

type Sz int64
type Dim int64

//

type Shape []Dim

func (sh Shape) Dim() Dim {
	r := sh[0]
	for _, c := range sh[1:] {
		r *= c
	}
	return r
}

//

type Orientation int8

const (
	RowMajor Orientation = iota
	ColumnMajor
)

func (o Orientation) BuildStrides(shape Shape, width Sz) Strides {
	switch o {
	case RowMajor:
		rem := int64(width)
		for _, v := range shape {
			rem *= int64(v)
		}

		if rem == 0 {
			strides := make(Strides, len(shape))
			for i := range strides {
				strides[i] = width
			}
			return strides
		}

		strides := make(Strides, len(shape))
		for i, v := range shape {
			rem /= int64(v)
			strides[i] = Sz(rem)
		}
		return strides

	case ColumnMajor:
		total := int64(width)
		for _, v := range shape {
			if v == 0 {
				strides := make(Strides, len(shape))
				for i := range strides {
					strides[i] = Sz(total)
				}
				return strides
			}
		}

		strides := make(Strides, len(shape))
		for i, v := range shape {
			strides[i] = Sz(total)
			total *= int64(v)
		}
		return strides

	}
	panic("unknown orientation")
}

//

type Strides []Sz

func (st Strides) Offset(idxs ...Dim) Sz {
	var ofs Sz
	for i, idx := range idxs {
		ofs += Sz(idx) * st[i]
	}
	return ofs
}

//

type BaseTensor struct {
	width  Sz
	shape  Shape
	orient Orientation

	strides Strides
}

func NewBaseTensor(
	width Sz,
	shape Shape,
	orient Orientation,
) BaseTensor {
	return BaseTensor{
		width:  width,
		shape:  shape,
		orient: orient,

		strides: orient.BuildStrides(shape, width),
	}
}

//

type HeapTensor struct {
	BaseTensor

	buf []byte
}

func NewHeapTensor(bt BaseTensor) HeapTensor {
	return HeapTensor{
		BaseTensor: bt,

		buf: make([]byte, bt.strides[0]*bt.width),
	}
}

//

func TestTensor(t *testing.T) {
	tns := NewHeapTensor(
		NewBaseTensor(
			Sz(4),
			Shape{4, 4, 4},
			RowMajor,
		),
	)

	fmt.Println(tns)

	fmt.Println(tns.strides.Offset(1, 3))
	fmt.Println(tns.strides.Offset(2))
	fmt.Println(tns.strides.Offset(2, 3))
	fmt.Println(tns.strides.Offset(2, 3, 0))
	fmt.Println(tns.strides.Offset(2, 3, 1))
	fmt.Println(tns.strides.Offset(3, 3, 3))
}
