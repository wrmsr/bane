package tg

import (
	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/maps"
	"github.com/wrmsr/bane/pkg/util/slices"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

type Tensor struct {
	data *LazyBuffer

	grad         *Tensor
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
	xsh := x.Shape()
	ysh := y.Shape()
	nDims := bt.Max(len(xsh), len(ysh))

	if len(xsh) != nDims {
		x = x.Reshape(slices.Join(xsh, slices.Repeat([]Dim{1}, nDims-len(xsh))))
	}
	if len(ysh) != nDims {
		y = y.Reshape(slices.Join(ysh, slices.Repeat([]Dim{1}, nDims-len(ysh))))
	}

	shapeRet := make(Shape, len(xsh))
	for i, sx := range xsh {
		sy := ysh[i]
		shapeRet[i] = bt.Max(sx, sy)
	}
	if !xsh.Equals(shapeRet) {
		x = x.Expand(shapeRet)
	}
	if !ysh.Equals(shapeRet) {
		y = y.Expand(shapeRet)
	}

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
	return out.Mul(NewTensor(MakeLoadConstBuffer(c, nil), false))
}

func (t *Tensor) Expand(shape Shape) *Tensor {
	return Apply(ReshapeFunc{Shape: shape}, []*Tensor{t})
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

	// self.grad = Tensor.ones(*self.shape, device=self.device, requires_grad=False)

	t.grad = NewTensor(
		MakeLoadBuffer(
			BufferOf(
				t.Shape(),
				slices.Repeat([]float32{1.}, int(t.Shape().Dim())),
			),
			t.Shape(),
		),
		false,
	)

	ps := slices.Reverse(t.deepWalk())
	for _, t0 := range ps {
		var hg bool
		for _, p := range t0.ctx.parents {
			if p.requiresGrad {
				hg = true
				break
			}
		}
		if !hg {
			continue
		}

		check.Condition(t0.grad != nil)

		ograds := t0.ctx.fn.Backward(t0.ctx, t0.grad.data)
		grads := make([]*Tensor, len(ograds))
		for i, g := range ograds {
			if g == nil {
				continue
			}
			grads[i] = NewTensor(g, false)
		}

		for i, p := range t0.ctx.parents {
			g := grads[i]
			if g == nil || !p.requiresGrad {
				continue
			}
			check.Condition(g.Shape().Equals(p.Shape()))
			if p.grad == nil {
				p.grad = g
			} else {
				p.grad = p.grad.Add(g)
			}
		}
	}
}
