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

var _ Func = AddFunc{}

func (a AddFunc) Forward(ctx *FuncContext, bs []*LazyBuffer) *LazyBuffer {
	check.Condition(len(bs) == 2)
	return bs[0].BinaryOp(AddOp, bs[1])
}

func (a AddFunc) Backward(ctx *FuncContext, g *LazyBuffer) []*LazyBuffer {
	panic("nyi")
}

//

type MulFunc struct{}

var _ Func = MulFunc{}

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

var _ Func = SumFunc{}

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
	panic("implement me")
}

//

type ReshapeFunc struct {
	Shape Shape
}

var _ Func = ReshapeFunc{}

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
	panic("implement me")
}
