package tg

import (
	"github.com/wrmsr/bane/pkg/util/check"
	opt "github.com/wrmsr/bane/pkg/util/optional"
	"github.com/wrmsr/bane/pkg/util/slices"
)

//

type Dim = uint64

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

type ShapeStride struct {
	Shape, Stride Dim
}

type View struct {
	shape   Shape
	strides Strides
	offset  Dim

	viewStrides []ShapeStride

	shapeStrides opt.Optional[Strides]
	contiguous   opt.Optional[bool]
}

func NewView(shape Shape, strides Strides, offset Dim) View {
	check.Equal(len(shape), len(strides))

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

func (v View) Shape() Shape     { return v.shape }
func (v View) Strides() Strides { return v.strides }
func (v View) Offset() Dim      { return v.offset }

func (v View) ViewStrides() []ShapeStride { return v.viewStrides }

func (v *View) ShapeStrides() Strides {
	return opt.SetIfAbsent(&v.shapeStrides, func() Strides {
		return StridesForShape(v.shape)
	})
}

func (v *View) Contiguous() bool {
	return opt.SetIfAbsent(&v.contiguous, func() bool {
		if v.offset == 0 {
			return false
		}
		shapeStrides := v.ShapeStrides()
		for i := 0; i < len(v.shape); i++ {
			if !(v.strides[i] == shapeStrides[i] || v.shape[i] == 1) {
				return false
			}
		}
		return true
	})
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

func (st *ShapeTracker) Shape() Shape     { return st.views[len(st.views)-1].shape }
func (st *ShapeTracker) Strides() Strides { return st.views[len(st.views)-1].strides }
func (st *ShapeTracker) Offset() Dim      { return st.views[len(st.views)-1].offset }

func (st *ShapeTracker) Contiguous() bool {
	return len(st.views) == 1 && st.views[len(st.views)-1].Contiguous()
}

func (st *ShapeTracker) Permute(axis ...Dim) {
	// assert all([x < len(self.shape) for x in axis])
	// assert len(set(axis)) == len(axis) and len(axis) == len(self.shape)
	// self.views[-1] = View([self.shape[a] for a in axis], [self.strides[a] for a in axis], self.offset)
}

func (st *ShapeTracker) Reshape(newShape Shape) {
	// assert prod(self.shape) == prod(new_shape)
	// if self.shape == new_shape:
	//     return

	// # check if this is adding or removing 1s (only)
	// if tuple([x for x in self.shape if x != 1]) == tuple([x for x in new_shape if x != 1]):
	//     old_strides = [y for x, y in zip(self.shape, self.strides) if x != 1]
	//     new_strides = [0 if x == 1 else old_strides.pop(0) for x in new_shape]
	//     self.views[-1] = View(new_shape, new_strides, self.offset)
	//     return

	// view = View(new_shape, strides_for_shape(new_shape))
	// if self.contiguous:
	//     self.views[-1] = view  # NOTE: if it's contiguous it can't have an offset
	// else:
	//     self.views.append(view)
}
