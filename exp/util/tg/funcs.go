package tg

import (
	"sort"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/slices"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type Func interface {
	Forward(ctx *FuncContext, bs []*LazyBuffer) *LazyBuffer
	Backward(ctx *FuncContext, g *LazyBuffer) []*LazyBuffer
}

//

type FuncContext struct {
	fn      Func
	parents []*Tensor

	needsInputGrad []bool
	requiresGrad   bool

	savedBuffers []*LazyBuffer

	inputShape bt.Optional[Shape]
	inputOrder bt.Optional[[]Dim]
	convArgs   bt.Optional[ConvArgs]
}

func NewFuncContext(fn Func, parents []*Tensor) *FuncContext {
	needsInputGrad := slices.Map((*Tensor).RequiresGrad, parents)
	return &FuncContext{
		fn:      fn,
		parents: parents,

		needsInputGrad: needsInputGrad,
		requiresGrad:   slices.Any(needsInputGrad, bt.Identity[bool]()),
	}
}

func (ctx *FuncContext) saveForBackward(bs ...*LazyBuffer) {
	ctx.savedBuffers = append(ctx.savedBuffers, bs...)
}

func Apply(fn Func, parents []*Tensor) *Tensor {
	ctx := NewFuncContext(fn, parents)
	ret := NewTensor(fn.Forward(ctx, slices.Map((*Tensor).Data, parents)), ctx.requiresGrad)
	if ctx.requiresGrad {
		ret.ctx = ctx
	}
	return ret
}

//

type AddFunc struct{}

func (a AddFunc) Forward(ctx *FuncContext, bs []*LazyBuffer) *LazyBuffer {
	check.Condition(len(bs) == 2)
	return bs[0].BinaryOp(AddOp, bs[1])
}

func (a AddFunc) Backward(ctx *FuncContext, g *LazyBuffer) []*LazyBuffer {
	return []*LazyBuffer{
		bt.Choose(ctx.needsInputGrad[0], g, nil),
		bt.Choose(ctx.needsInputGrad[1], g, nil),
	}
}

//

type SubFunc struct{}

func (a SubFunc) Forward(ctx *FuncContext, bs []*LazyBuffer) *LazyBuffer {
	check.Condition(len(bs) == 2)
	return bs[0].BinaryOp(SubOp, bs[1])
}

func (a SubFunc) Backward(ctx *FuncContext, g *LazyBuffer) []*LazyBuffer {
	ret := []*LazyBuffer{
		bt.Choose(ctx.needsInputGrad[0], g, nil),
		nil,
	}
	if ctx.needsInputGrad[1] {
		ret[1] = g.UnaryOp(NegOp)
	}
	return ret
}

//

type MulFunc struct{}

func (a MulFunc) Forward(ctx *FuncContext, bs []*LazyBuffer) *LazyBuffer {
	check.Condition(len(bs) == 2)
	ctx.saveForBackward(bs...)
	return bs[0].BinaryOp(MulOp, bs[1])
}

func (a MulFunc) Backward(ctx *FuncContext, g *LazyBuffer) []*LazyBuffer {
	var gradX *LazyBuffer
	if ctx.needsInputGrad[0] {
		gradX = ctx.savedBuffers[1].BinaryOp(MulOp, g)
	}
	var gradY *LazyBuffer
	if ctx.needsInputGrad[1] {
		gradY = ctx.savedBuffers[0].BinaryOp(MulOp, g)
	}
	return []*LazyBuffer{gradX, gradY}
}

//

type PowFunc struct{}

func (a PowFunc) Forward(ctx *FuncContext, bs []*LazyBuffer) *LazyBuffer {
	x, y := slices.Unpack2(bs)
	ret := x.BinaryOp(PowOp, y)
	ctx.saveForBackward(x, y, ret)
	return ret
}

func (a PowFunc) Backward(ctx *FuncContext, g *LazyBuffer) []*LazyBuffer {
	x, y, powxy := slices.Unpack3(ctx.savedBuffers)
	var gradx, grady *LazyBuffer
	if ctx.needsInputGrad[0] {
		tmp := y.BinaryOp(MulOp, powxy.BinaryOp(DivOp, x))
		gradx = g.BinaryOp(MulOp, tmp)
	}
	if ctx.needsInputGrad[1] {
		tmp := x.UnaryOp(LogOp).BinaryOp(MulOp, powxy)
		grady = g.BinaryOp(MulOp, tmp)
	}
	return []*LazyBuffer{gradx, grady}
}

//

type SumFunc struct {
	Axis []int
}

func reduceShape(shape Shape, axis []int) Shape {
	newShape := make(Shape, len(shape))
	for i := 0; i < len(shape); i++ {
		if slices.Contains(axis, i) {
			newShape[i] = 1
		} else {
			newShape[i] = shape[i]
		}
	}
	return newShape
}

func (f SumFunc) Forward(ctx *FuncContext, bs []*LazyBuffer) *LazyBuffer {
	input := check.Single(bs)
	ctx.inputShape = bt.Just(input.Shape())
	return input.ReduceOp(SumOp, reduceShape(input.Shape(), f.Axis))
}

func (f SumFunc) Backward(ctx *FuncContext, g *LazyBuffer) []*LazyBuffer {
	return []*LazyBuffer{g.MovementOp(ExpandOp, ctx.inputShape.Value())}
}

//

type ReshapeFunc struct {
	Shape Shape
}

func (f ReshapeFunc) Forward(ctx *FuncContext, bs []*LazyBuffer) *LazyBuffer {
	input := check.Single(bs)
	oldShape := input.Shape()
	ctx.inputShape = bt.Just(oldShape)
	newShape := make(Shape, len(f.Shape))
	for i, s := range f.Shape {
		if s == -1 {
			newShape[i] = (bt.Prod[Dim](oldShape...) * -1) / bt.Prod[Dim](f.Shape...)
		} else {
			newShape[i] = s
		}
	}
	return input.MovementOp(ReshapeOp, newShape)
}

func (f ReshapeFunc) Backward(ctx *FuncContext, g *LazyBuffer) []*LazyBuffer {
	return []*LazyBuffer{g.MovementOp(ReshapeOp, ctx.inputShape.Value())}
}

//

type ExpandFunc struct {
	Shape Shape
}

func (f ExpandFunc) Forward(ctx *FuncContext, bs []*LazyBuffer) *LazyBuffer {
	input := check.Single(bs)
	ctx.inputShape = bt.Just(input.Shape())
	return input.MovementOp(ExpandOp, f.Shape)
}

func (f ExpandFunc) Backward(ctx *FuncContext, g *LazyBuffer) []*LazyBuffer {
	return []*LazyBuffer{g.ReduceOp(SumOp, ctx.inputShape.Value())}
}

//

type ReluFunc struct{}

func (f ReluFunc) Forward(ctx *FuncContext, bs []*LazyBuffer) *LazyBuffer {
	input := check.Single(bs)
	ctx.saveForBackward(input)
	return input.UnaryOp(ReluOp)
}

func (f ReluFunc) Backward(ctx *FuncContext, g *LazyBuffer) []*LazyBuffer {
	return []*LazyBuffer{check.Single(ctx.savedBuffers).UnaryOp(SignOp).UnaryOp(ReluOp).BinaryOp(MulOp, g)}
}

//

type PermuteFunc struct {
	Order []Dim
}

func (f PermuteFunc) Forward(ctx *FuncContext, bs []*LazyBuffer) *LazyBuffer {
	input := check.Single(bs)
	ctx.inputOrder = bt.Just(f.Order)
	return input.MovementOp(PermuteOp, f.Order)
}

func (f PermuteFunc) Backward(ctx *FuncContext, g *LazyBuffer) []*LazyBuffer {
	ds := ctx.inputOrder.Value()
	a := bt.RangeTo(Dim(len(ds))).Slice()
	sort.Slice(a, func(i, j int) bool {
		return ds[i] < ds[j]
	})
	return []*LazyBuffer{g.MovementOp(PermuteOp, a)}
}

//

type Conv2dFunc struct {
	co ConvOpts
}

func (f Conv2dFunc) Forward(ctx *FuncContext, bs []*LazyBuffer) *LazyBuffer {
	x, w := slices.Unpack2(bs)
	ca := BuildConvArgs(x.Shape(), w.Shape(), f.co)
	ctx.convArgs = bt.Just(ca)
	ctx.saveForBackward(x, w)
	return x.ProcessingOp(ConvOp, w, ca)
}

func (f Conv2dFunc) Backward(ctx *FuncContext, g *LazyBuffer) []*LazyBuffer {
	x, w := check.Pair(ctx.savedBuffers)
	ca := ctx.convArgs.Value()

	var dx, dw *LazyBuffer

	// compute derivative of inputs using ProcessingOps.CONV (this is a transposed conv)
	if ctx.needsInputGrad[0] {
		xt := g
		// unstride. NOTE: this is really memory intensive for big strides.
		if ca.sx > 1 || ca.sy > 1 {
			xt = xt.MovementOp(
				ReshapeOp,
				[]Dim{
					g.Shape()[0],
					g.Shape()[1],
					g.Shape()[2],
					1,
					g.Shape()[3],
					1,
				},
			)

			xt = xt.MovementOp(
				SliceOp,
				[]SliceBound{
					{0, xt.Shape()[0]},
					{0, xt.Shape()[1]},
					{0, xt.Shape()[2]},
					{0, ca.sy},
					{0, xt.Shape()[4]},
					{0, ca.sx},
				},
			)

			xt = xt.MovementOp(
				ReshapeOp,
				[]Dim{
					xt.Shape()[0],
					xt.Shape()[1],
					xt.Shape()[2] * ca.sy,
					xt.Shape()[4] * ca.sx,
				},
			)
		}

		wt := w.MovementOp(
			ReshapeOp,
			Shape{
				ca.groups,
				ca.rcout,
				ca.cin,
				ca.h,
				ca.w,
			},
		).MovementOp(
			PermuteOp,
			[]Dim{
				0,
				2,
				1,
				3,
				4,
			},
		)

		wt = wt.MovementOp(
			ReshapeOp,
			Shape{
				ca.groups * ca.cin,
				ca.rcout,
				ca.h,
				ca.w,
			},
		).MovementOp(
			FlipOp,
			[]Dim{2, 3},
		)

		py, px := (ca.h-1)*ca.dy-ca.py, (ca.w-1)*ca.dx-ca.px
		cdx := BuildConvArgs(
			xt.Shape(),
			wt.Shape(),
			ConvOpts{
				OutShape: x.Shape(),
				Dilation: []Dim{ca.dy, ca.dx},
				Padding:  []Dim{py, px},
				Groups:   ca.groups,
			},
		)

		dx = xt.ProcessingOp(ConvOp, wt, cdx)
	}

	// compute derivative of weights using ProcessingOps.CONV
	if ctx.needsInputGrad[1] {
		xdw := x.MovementOp(
			ReshapeOp,
			Shape{
				ca.bs,
				ca.groups,
				ca.cin,
				ca.iy,
				ca.ix,
			},
		).MovementOp(
			PermuteOp,
			[]Dim{
				2,
				1,
				0,
				3,
				4,
			},
		)

		xdw = xdw.MovementOp(
			ReshapeOp,
			Shape{
				ca.cin,
				ca.groups * ca.bs,
				ca.iy,
				ca.ix,
			},
		)

		grad_output_dw := g.MovementOp(PermuteOp, []Dim{1, 0, 2, 3})

		cdw := BuildConvArgs(
			xdw.Shape(),
			grad_output_dw.Shape(),
			ConvOpts{
				OutShape: []Dim{
					w.Shape()[1],
					w.Shape()[0],
					w.Shape()[2],
					w.Shape()[3],
				},
				Padding:  []Dim{ca.py, ca.px},
				Stride:   []Dim{ca.dy, ca.dx},
				Dilation: []Dim{ca.sy, ca.sx},
				Groups:   ca.groups,
			},
		)

		dw = xdw.ProcessingOp(ConvOp, grad_output_dw, cdw).MovementOp(PermuteOp, []Dim{1, 0, 2, 3})
	}

	return []*LazyBuffer{dx, dw}
}

//

type MaxFunc struct {
	axis []int
}

func (f MaxFunc) Forward(ctx *FuncContext, bs []*LazyBuffer) *LazyBuffer {
	input := check.Single(bs)
	ret := input.ReduceOp(MaxOp, reduceShape(input.Shape(), f.axis))
	ctx.saveForBackward(input, ret)
	return ret
}

func (f MaxFunc) Backward(ctx *FuncContext, g *LazyBuffer) []*LazyBuffer {
	input, ret := check.Pair(ctx.savedBuffers)

	// 1s in locations where the max was chosen (can be two locations)
	max_is_1s := input.BinaryOp(CmpEqOp, ret.MovementOp(ExpandOp, input.Shape()))

	// sum of locations, averaged
	div := max_is_1s.ReduceOp(SumOp, g.Shape())
	div = div.MovementOp(ExpandOp, input.Shape())

	max_is_amount := max_is_1s.BinaryOp(DivOp, div)
	grad_output_expanded := g.MovementOp(ExpandOp, input.Shape())

	return []*LazyBuffer{max_is_amount.BinaryOp(MulOp, grad_output_expanded)}
}

//

type ExpFunc struct{}

func (f ExpFunc) Forward(ctx *FuncContext, bs []*LazyBuffer) *LazyBuffer {
	input := check.Single(bs)
	ret := input.UnaryOp(ExpOp)
	ctx.saveForBackward(ret)
	return ret
}

func (f ExpFunc) Backward(ctx *FuncContext, g *LazyBuffer) []*LazyBuffer {
	return []*LazyBuffer{ctx.savedBuffers[0].BinaryOp(MulOp, g)}
}

//

type LogFunc struct{}

func (f LogFunc) Forward(ctx *FuncContext, bs []*LazyBuffer) *LazyBuffer {
	input := check.Single(bs)
	ctx.saveForBackward(input)
	return input.UnaryOp(LogOp)
}

func (f LogFunc) Backward(ctx *FuncContext, g *LazyBuffer) []*LazyBuffer {
	return []*LazyBuffer{g.BinaryOp(DivOp, ctx.savedBuffers[0])}
}
