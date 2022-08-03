package tg

import (
	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/maps"
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
		return Apply(AddFunc{}, []*Tensor{x, y})
	}, t, y)
}

func (t *Tensor) Mul(y *Tensor) *Tensor {
	return BroadcastedTensor(func(x, y *Tensor) *Tensor {
		return Apply(MulFunc{}, []*Tensor{x, y})
	}, t, y)
}

func (t *Tensor) Reshape(shape Shape) *Tensor {
	return Apply(ReshapeFunc{Shape: shape}, []*Tensor{t})
}

func canonicalizeReduceAxis(shape Shape, axis []int) ([]int, Shape) {
	if len(axis) < 1 {
		axis = bt.RangeTo(len(shape)).Slice()
	}
	for i, x := range axis {
		if x < 0 {
			axis[i] += len(shape)
		}
	}
	var newShape Shape
	for i := 0; i < len(shape); i++ {
		if !slices.Contains(axis, i) {
			newShape = append(newShape, shape[i])
		}
	}
	if len(newShape) < 1 {
		newShape = scalarShape
	}
	return axis, newShape
}

func (t *Tensor) Sum(axis []int, keepDim bool) *Tensor {
	axis, outShape := canonicalizeReduceAxis(t.Shape(), axis)
	ret := Apply(SumFunc{Axis: axis}, []*Tensor{t})
	if keepDim || ret.Shape().Equals(outShape) {
		return ret
	}
	return ret.Reshape(outShape)
}

func (t *Tensor) Mean(axis []int, keepDim bool) *Tensor {
	out := t.Sum(axis, keepDim)
	c := float32(bt.Prod[Dim](out.Shape()...)) / float32(bt.Prod[Dim](t.Shape()...))
	return out.Mul(NewTensor(MakeConstBuffer(c), false))
}

var scalarShape = Shape{1}

func (t *Tensor) deepWalk() []*Tensor {
	seen := maps.MakeSet[*Tensor]()
	var ret []*Tensor
	var rec func(*Tensor)
	rec = func(node *Tensor) {
		seen.Add(node)
		if node.ctx != nil {
			for _, p := range node.ctx.parents {
				if !seen.Contains(p) {
					rec(p)
				}
			}
			ret = append(ret, node)
		}
	}
	rec(t)
	return ret
}
func (t *Tensor) Backward() {
	check.Condition(t.Shape().Equals(scalarShape))
	ps := slices.Reverse(t.deepWalk())
	_ = ps
	panic("nyi")
}
