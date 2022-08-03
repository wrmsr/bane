package tg

import (
	"github.com/wrmsr/bane/pkg/util/slices"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

type Tensor struct {
	data *LazyBuffer

	grad         *LazyBuffer
	requiresGrad bool

	ctx *FuncContext
}

func NewTensor(data *LazyBuffer, requiresGrad bool) *Tensor {
	return &Tensor{
		data: data,

		requiresGrad: requiresGrad,
	}
}

func (t *Tensor) Data() *LazyBuffer  { return t.data }
func (t *Tensor) RequiresGrad() bool { return t.requiresGrad }

func (t *Tensor) Shape() Shape {
	return t.data.st.Shape()
}

func BroadcastedTensor(fn func(x, y *Tensor) *Tensor, x, y *Tensor) *Tensor {
	// nDims := bt.Max(len(x.Shape()), len(y.Shape()))

	// if len(x.shape) != n_dims {
	// 	x = x.reshape(list(x.shape) + [1] * (n_dims - len(x.shape)))
	// }
	// if len(y.shape) != n_dims {
	// 	y = y.reshape(list(y.shape) + [1] * (n_dims - len(y.shape)))
	// }

	// shapeRet = tuple([max(sx, sy) for sx, sy in zip(x.shape, y.shape)])
	// if x.shape != shapeRet {
	//	x = x.expand(shape_ret)
	// }
	// if y.shape != shapeRet {
	//	y = y.expand(shape_ret)
	// }

	return fn(x, y)
}

func (t *Tensor) Add(y *Tensor) *Tensor {
	return BroadcastedTensor(func(x, y *Tensor) *Tensor {
		ps := []*Tensor{x, y}

		fn := AddFunc{}
		ctx := NewFuncContext(fn, ps)

		bs := slices.Map((*Tensor).Data, ps)
		z := NewTensor(
			fn.Forward(bs),
			ctx.requiresGrad,
		)

		if ctx.requiresGrad {
			z.ctx = ctx
		}

		return z
	}, t, y)
}

func canonicalizeReduceAxis(shape Shape, axis []int) ([]int, Shape) {
	if len(axis) < 1 {
		axis = bt.RangeTo(len(shape)).Slice()
	}
	// axis = tuple([x if x >= 0 else x + len(self.shape) for x in axis])
	// shape = [self.shape[i] for i in range(len(self.shape)) if i not in axis]
	// shape = [1] if shape == [] else shape
	// return axis, shape
	panic("nyi")
}

func (t *Tensor) Sum(axis []int, keepDim bool) *Tensor {
	axis, outShape := canonicalizeReduceAxis(t.Shape(), axis)
	/*
	   ret := self._sum(axis=axis)
	   return ret if keepDim or ret.shape == outShape else ret.reshape(shape=outShape)
	*/
	_ = outShape
	panic("nyi")
}

func (t *Tensor) Mean(axis []int, keepDim bool) *Tensor {
	out := t.Sum(axis, keepDim)
	/*
		return out * (prod(out.shape) / prod(self.shape))
	*/
	_ = out
	panic("nyi")
}
