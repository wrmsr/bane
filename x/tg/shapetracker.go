package tg

import (
	"github.com/wrmsr/bane/core/check"
	"github.com/wrmsr/bane/core/maps"
	"github.com/wrmsr/bane/core/slices"
	bt "github.com/wrmsr/bane/core/types"
)

//

type View struct {
	shape   Shape
	strides Strides
	offset  Dim

	shapeStrides []ShapeStride

	baseStrides bt.Optional[Strides]
	contiguous  bt.Optional[bool]
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
	return bt.SetIfAbsent(&v.baseStrides, func() Strides {
		return StridesForShape(v.shape)
	})
}

func (v *View) Contiguous() bool {
	return bt.SetIfAbsent(&v.contiguous, func() bool {
		if v.offset != 0 {
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

func (st *ShapeTracker) Clone() *ShapeTracker {
	return &ShapeTracker{
		views: slices.Clone(st.views),
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

func (st *ShapeTracker) Flip(axes ...Dim) {
	s := make([]Dim, len(st.Shape()))
	for i := range s {
		if slices.Contains(axes, Dim(i)) {
			s[i] = -1
		} else {
			s[i] = 1
		}
	}
	st.Stride(s...)
}

func (st *ShapeTracker) MovementOp(op Op, arg any) {
	switch op {
	case ReshapeOp:
		st.Reshape(arg.(Shape))
	case ExpandOp:
		st.Expand(arg.(Shape))
	case PermuteOp:
		st.Permute(arg.([]Dim)...)
	case FlipOp:
		st.Flip(arg.([]Dim)...)
	default:
		panic(op)
	}
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
				newStrides = append(newStrides, 0)
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

func (st *ShapeTracker) Expand(newShape Shape) {
	oldShape := st.Shape()
	for i, x := range oldShape {
		y := newShape[i]
		check.Condition(x == y || x == 1)
	}
	newStrides := make(Strides, len(oldShape))
	for i, s := range st.Strides() {
		x := oldShape[i]
		y := newShape[i]
		newStrides[i] = bt.Choose(x == y, s, 0)

	}
	st.views = append(st.views, NewView(newShape, newStrides, st.Offset()))
}

func (st *ShapeTracker) Stride(mul ...Dim) {
	shape := st.Shape()
	strides := st.Strides()
	newStrides := slices.ZipMap2(bt.Mul[Dim], strides, mul)
	newShape := slices.ZipMap2(func(s, m Dim) Dim { return (s + (bt.Abs(m) - 1)) / bt.Abs(m) }, shape, mul)
	newOffset := bt.Sum[Dim](
		slices.ZipMap3(func(s, z, m Dim) Dim { return bt.Choose(m < 0, (s-1)*z, 0) }, shape, strides, mul)...,
	) + st.Offset()
	st.views[len(st.views)-1] = NewView(newShape, newStrides, newOffset)
}
