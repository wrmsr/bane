package tg

import (
	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/maps"
	opt "github.com/wrmsr/bane/pkg/util/optional"
	"github.com/wrmsr/bane/pkg/util/slices"
	bt "github.com/wrmsr/bane/pkg/util/types"
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

func ToShapeStrides(shape Shape, strides Strides) []ShapeStride {
	ss := []ShapeStride{{shape[0], strides[0]}}
	for i := 1; i < len(shape); i++ {
		if (strides[i] != 0 && ss[len(ss)-1].Stride/strides[i] == shape[i]) || (strides[i] == 0 && ss[len(ss)-1].Stride == 0) {
			ss[len(ss)-1] = ShapeStride{ss[len(ss)-1].Shape * shape[i], strides[i]}
		} else {
			ss = append(ss, ShapeStride{shape[i], strides[i]})
		}
	}
	return ss
}

type View struct {
	shape   Shape
	strides Strides
	offset  Dim

	shapeStrides []ShapeStride

	baseStrides opt.Optional[Strides]
	contiguous  opt.Optional[bool]
}

func NewView(shape Shape, strides Strides, offset Dim) View {
	check.Equal(len(shape), len(strides))

	return View{
		shape:   shape,
		strides: strides,
		offset:  offset,

		shapeStrides: ToShapeStrides(shape, strides),
	}
}

func (v View) Shape() Shape     { return v.shape }
func (v View) Strides() Strides { return v.strides }
func (v View) Offset() Dim      { return v.offset }

func (v View) ShapeStrides() []ShapeStride { return v.shapeStrides }

func (v *View) BaseStrides() Strides {
	return opt.SetIfAbsent(&v.baseStrides, func() Strides {
		return StridesForShape(v.shape)
	})
}

func (v *View) Contiguous() bool {
	return opt.SetIfAbsent(&v.contiguous, func() bool {
		if v.offset == 0 {
			return false
		}
		shapeStrides := v.BaseStrides()
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
	shape := st.Shape()
	strides := st.Strides()

	check.Condition(slices.All(axis, func(x Dim) bool { return x < Dim(len(shape)) }))
	check.Condition(len(maps.MakeSetOf(axis)) == len(axis))
	check.Condition(len(axis) == len(shape))

	st.views[len(st.views)-1] = NewView(
		slices.Map(func(a Dim) Dim { return shape[a] }, axis),
		slices.Map(func(a Dim) Dim { return strides[a] }, axis),
		st.Offset(),
	)
}

func (st *ShapeTracker) Reshape(newShape Shape) {
	shape := st.Shape()
	strides := st.Strides()
	check.Condition(bt.Prod[Dim](shape...) == bt.Prod[Dim](newShape...))
	if slices.Equal(shape, newShape) {
		return
	}

	// Check if this is adding or removing 1s (only)
	if slices.Equal(
		slices.Filter(func(x Dim) bool { return x != 1 }, shape),
		slices.Filter(func(x Dim) bool { return x != 1 }, newShape)) {
		var oldStrides Strides
		for i, x := range shape {
			if x != 1 {
				oldStrides = append(oldStrides, strides[i])
			}
		}

		var newStrides Strides
		j := 0
		for _, x := range newShape {
			if x == 1 {
				newStrides = append(newStrides, x)
			} else {
				newStrides = append(newStrides, oldStrides[j])
				j++
			}
		}

		st.views[len(st.views)-1] = NewView(newShape, newStrides, st.Offset())
		return
	}

	newView := NewView(newShape, StridesForShape(newShape), 0)
	if st.Contiguous() {
		st.views[len(st.views)-1] = newView // NOTE: If it's contiguous it can't have an offset
	} else {
		st.views = append(st.views, newView)
	}
}
