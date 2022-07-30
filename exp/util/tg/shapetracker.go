package tg

import (
	"github.com/wrmsr/bane/pkg/util/slices"
)

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

func StridesForShape(shape Shape) Strides {
	strides := make(Strides, len(shape))
	strides[len(shape)-1] = 1
	for i := len(shape) - 2; i >= 0; i-- {
		strides[i] = strides[i+1] * shape[i+1]
	}
	return strides
}

func ViewFromShape(shape Shape) View {
	if len(shape) < 1 {
		shape = Shape{1}
	}
	return NewView(shape, StridesForShape(shape), 0)
}

//

type ShapeTracker struct {
	views []View
}

func NewShapeTracker(shape Shape) *ShapeTracker {
	return &ShapeTracker{
		views: []View{ViewFromShape(shape)},
	}
}
