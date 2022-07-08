package tensor

import (
	"fmt"
	"testing"
)

type ByteWidth int64

type Shape []int64

type Orientation int8

const (
	RowMajor Orientation = iota
	ColumnMajor
)

func (o Orientation) BuildStrides(shape Shape, width ByteWidth) Strides {
	switch o {
	case RowMajor:
		rem := int64(width)
		for _, v := range shape {
			rem *= v
		}

		if rem == 0 {
			strides := make(Strides, len(shape))
			for i := range strides {
				strides[i] = int64(width)
			}
			return strides
		}

		strides := make(Strides, len(shape))
		for i, v := range shape {
			rem /= v
			strides[i] = rem
		}
		return strides

	case ColumnMajor:
		total := int64(width)
		for _, v := range shape {
			if v == 0 {
				strides := make(Strides, len(shape))
				for i := range strides {
					strides[i] = total
				}
				return strides
			}
		}

		strides := make(Strides, len(shape))
		for i, v := range shape {
			strides[i] = total
			total *= v
		}
		return strides

	}
	panic("unknown orientation")
}

type Index int64

type Offset int64

func (s Strides) Offset(idxs []Index, width ByteWidth) Offset {
	var offset int64
	for i, idx := range idxs {
		offset += int64(idx) * s[i]
	}
	return Offset(offset / int64(width))
}

type Strides []int64

type BaseTensor struct {
	width  ByteWidth
	shape  Shape
	orient Orientation

	strides Strides
}

func NewBaseTensor(
	width ByteWidth,
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

func (t BaseTensor) Offset(idxs ...Index) Offset {
	return t.strides.Offset(idxs, t.width)
}

//type Float32Tensor struct {
//	buf []float32
//}

func TestTensor(t *testing.T) {
	tns := NewBaseTensor(
		ByteWidth(4),
		Shape{4, 4, 4},
		RowMajor,
	)

	fmt.Println(tns)

	fmt.Println(tns.Offset(1, 3))
	fmt.Println(tns.Offset(2))
	fmt.Println(tns.Offset(2, 3))
	fmt.Println(tns.Offset(2, 3, 0))
	fmt.Println(tns.Offset(2, 3, 1))
}
