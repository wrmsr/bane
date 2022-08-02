package tg

import (
	"github.com/wrmsr/bane/pkg/util/slices"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type Func interface {
	Forward(bs []*LazyBuffer) *LazyBuffer
	Backward(g *LazyBuffer) []*LazyBuffer

	isFunc()
}

//

type FuncContext struct {
	fn      Func
	parents []*Tensor

	needsInputGrad []bool
	requiresGrad   bool

	savedTensors []*Tensor
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
	ret := NewTensor(fn.Forward(slices.Map((*Tensor).Data, parents)), ctx.requiresGrad)
	if ctx.requiresGrad {
		ret.ctx = ctx
	}
	return ret
}
