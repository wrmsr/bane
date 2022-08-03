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

	savedTensors []*Tensor

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

func (s SumFunc) Forward(ctx *FuncContext, bs []*LazyBuffer) *LazyBuffer {
	input := check.Single(bs)
	ctx.inputShape = opt.Just(input.Shape())
	return input.ReduceOp(SumOp, reduceShape(input.Shape(), s.Axis))
}

func (s SumFunc) Backward(ctx *FuncContext, g *LazyBuffer) []*LazyBuffer {
	panic("implement me")
}
