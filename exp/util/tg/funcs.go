package tg

import (
	"github.com/wrmsr/bane/pkg/util/check"
	opt "github.com/wrmsr/bane/pkg/util/optional"
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

	inputShape opt.Optional[Shape]
	inputOrder opt.Optional[[]Dim]
	convArgs   opt.Optional[ConvArgs]
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
	ctx.inputShape = opt.Just(input.Shape())
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
	ctx.inputShape = opt.Just(oldShape)
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

type ReluFunc struct{}

func (f ReluFunc) Forward(ctx *FuncContext, bs []*LazyBuffer) *LazyBuffer {
	input := check.Single(bs)
	ctx.saveForBackward(input)
	return input.UnaryOp(ReluOp)
}

func (f ReluFunc) Backward(ctx *FuncContext, g *LazyBuffer) []*LazyBuffer {
	return []*LazyBuffer{ctx.savedBuffers[0].UnaryOp(SignOp).UnaryOp(ReluOp).BinaryOp(MulOp, g)}
}

//

type PermuteFunc struct {
	Order []Dim
}

func (f PermuteFunc) Forward(ctx *FuncContext, bs []*LazyBuffer) *LazyBuffer {
	input := check.Single(bs)
	ctx.inputOrder = opt.Just(f.Order)
	return input.MovementOp(PermuteOp, f.Order)
}

func (f PermuteFunc) Backward(ctx *FuncContext, g *LazyBuffer) []*LazyBuffer {
	panic("implement me")
}

//

type Conv2dFunc struct {
	co ConvOpts
}

func (f Conv2dFunc) Forward(ctx *FuncContext, bs []*LazyBuffer) *LazyBuffer {
	x, w := slices.Unpack2(bs)
	ca := BuildConvArgs(x.Shape(), w.Shape(), f.co)
	ctx.convArgs = opt.Just(ca)
	ctx.saveForBackward(x, w)
	return x.ProcessingOp(ConvOp, w, ca)
}

func (f Conv2dFunc) Backward(ctx *FuncContext, g *LazyBuffer) []*LazyBuffer {
	panic("implement me")
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
	/*
	   input, ret = ctx.saved_tensors

	   # 1s in locations where the max was chosen (can be two locations)
	   max_is_1s = input.binary_op(BinaryOps.CMPEQ, ret.movement_op(MovementOps.EXPAND, input.shape))

	   # sum of locations, averaged
	   div = max_is_1s.reduce_op(ReduceOps.SUM, grad_output.shape)
	   div = div.movement_op(MovementOps.EXPAND, input.shape)
	   max_is_amount = max_is_1s.binary_op(BinaryOps.DIV, div)

	   grad_output_expanded = grad_output.movement_op(MovementOps.EXPAND, input.shape)
	   return max_is_amount.binary_op(BinaryOps.MUL, grad_output_expanded)
	*/
	panic("implement me")
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
