package tg

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/core/check"
	"github.com/wrmsr/bane/core/slices"
	bt "github.com/wrmsr/bane/core/types"
)

func TestOps1(t *testing.T) {
	//a := -.5
	//b := 20.
	//rand.Seed(0)

	for _, sh := range []Shape{
		{9, 1},
		{3, 3},
	} {
		xt := NewTensor(MakeLoadBuffer(BufferOf(sh, bt.RangeTo[float32](9.).Slice())), true)
		yt := NewTensor(MakeLoadBuffer(BufferOf(sh, bt.RangeOf[float32](10., 19., 1.).Slice())), true)

		zt := xt.Add(yt)
		fmt.Println(zt)

		fmt.Println(zt.Data().Realize())

		zt.Mean(nil, false).Backward()

		zg := zt.grad.Data()

		zgt := zg.Realize()
		fmt.Println(zgt)
	}
}

func TestOpsRelu(t *testing.T) {
	for _, sh := range []Shape{
		{9, 1},
		{3, 3},
	} {
		xs := BufferOf(sh, slices.Map(func(v float32) float32 { return (v / 9) - .5 }, bt.RangeTo[float32](9.).Slice()))

		xt := NewTensor(MakeLoadBuffer(xs), true)

		zt := xt.Relu()
		fmt.Println(zt)

		fmt.Println(zt.Data().Realize())

		zt.Mean(nil, false).Backward()

		zg := zt.grad.Data()

		zgt := zg.Realize()
		fmt.Println(zgt)
	}
}

func dumpObj(o any) {
	var rec func(any, string)
	rec = func(o any, p string) {
		switch o := o.(type) {
		case *Tensor:
			fmt.Printf("%stn @%p\n", p, o)
			rec(o.data, "  "+p)
		case *LazyBuffer:
			fmt.Printf("%slb @%p\n", p, o)
			rec(o.op, "  "+p)
		case *LazyOp:
			fmt.Printf("%s%v @%p\n", p, o.op, o)
			for _, c := range o.srcs {
				rec(c, "  "+p)
			}
		default:
			panic(o)
		}
	}
	rec(o, "")
}

func TestBobNet(t *testing.T) {
	sh := Shape{3, 3}

	xt := NewTensor(MakeLoadBuffer(BufferOf(sh, bt.RangeTo[float32](9.).Slice())), true).Reshape(sh)
	l1t := NewTensor(MakeLoadBuffer(BufferOf(sh, bt.RangeOf[float32](10., 19., 1.).Slice())), true).Reshape(sh)
	l2t := NewTensor(MakeLoadBuffer(BufferOf(sh, bt.RangeOf[float32](20., 29., 1.).Slice())), true).Reshape(sh)

	zt := xt.Dot(l1t).Relu().Dot(l2t).LogSoftmax()
	//zt := xt.Dot(l1t).Relu().Dot(l2t)
	fmt.Println("====")

	dumpObj(zt)

	//fmt.Println(zt.Mean(nil, false).Data().Realize())
	//fmt.Println(zt.Mean(nil, false).Data().Realize())

	// garbage

	scc := func(out *Tensor, y_ []int) *Tensor {
		check.Equal(Dim(len(y_)), out.Shape()[0])
		check.Equal(out.Shape()[0], Dim(len(y_)))
		num_classes := out.Shape()[len(out.Shape())-1]
		yb := NewBuffer(out.Shape())
		for i := Dim(0); i < Dim(len(y_)); i++ {
			yb.set(float32(-1*num_classes), i, Dim(y_[i]))
		}
		y := NewTensor(MakeLoadBuffer(yb), false)
		return out.Mul(y).Mean(nil, false)
	}

	y := make([]int, 3)
	for i := 0; i < 3; i++ {
		y[i] = i % 3
	}

	z := scc(
		zt,
		y,
	)

	fmt.Println(z.Data().Realize().Nd())
	z.Backward()

	//y := NewTensor(MakeLoadBuffer(BufferOf(sh, bt.RangeTo[float32](3.).Slice()), Shape{3}), true)

	//zt.Mean(nil, false).Backward()
	//zg := zt.grad.Data()
	//zgt := zg.Realize()
	//fmt.Println(zgt)
}
