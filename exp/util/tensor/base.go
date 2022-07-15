package tensor

import "github.com/wrmsr/bane/pkg/util/slices"

//

type Dim = int64

//

type Shape []Dim

func (sh Shape) Dim() Dim {
	r := sh[0]
	for _, c := range sh[1:] {
		r *= c
	}
	return r
}

func (sh Shape) At(i int) Dim {
	if i < 0 {
		return sh[len(sh)+i]
	}
	return sh[i]
}

func (sh Shape) Equals(o Shape) bool {
	return slices.Equal(sh, o)
}

//

func BuildStrides(shape Shape, width Dim) Strides {
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
		strides[i] = Dim(rem)
	}
	return strides
}

//

type Strides []Dim

func (st Strides) Offset(idxs ...Dim) Dim {
	var ofs Dim
	for i, idx := range idxs {
		ofs += idx * st[i]
	}
	return ofs
}

//

type ShapeStride struct {
	Shape, Stride Dim
}

type View struct {
	shape   Shape
	strides Strides
	offset  Dim

	viewStrides []ShapeStride
}

func NewView(shape Shape, strides Strides, offset Dim) View {
	vs := []ShapeStride{{shape[0], strides[1]}}
	for i := 1; i < len(shape); i++ {
		if (strides[i] != 0 && vs[len(vs)-1].Stride/strides[i] == shape[i]) || (strides[i] == 0 && vs[len(vs)-1].Stride == 0) {
			vs[len(vs)-1] = ShapeStride{vs[len(vs)-1].Shape * shape[i], strides[i]}
		} else {
			vs = append(vs, ShapeStride{shape[i], strides[i]})
		}
	}
	return View{
		shape:       shape,
		strides:     strides,
		offset:      offset,
		viewStrides: vs,
	}
}

//

type BaseTensor struct {
	width Dim
	shape Shape

	strides Strides
}

func NewBaseTensor(
	shape Shape,
) BaseTensor {
	return BaseTensor{
		shape: shape,

		strides: BuildStrides(shape, 1),
	}
}

func (t BaseTensor) Sz() Dim {
	return t.strides[0] * t.width
}
