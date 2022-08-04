package tg

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/slices"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

func TestOps1(t *testing.T) {
	//a := -.5
	//b := 20.
	//rand.Seed(0)

	for _, sh := range []Shape{
		{9, 1},
		{3, 3},
	} {
		xt := NewTensor(MakeLoadBuffer(BufferOf(sh, bt.RangeTo[float32](9.).Slice()), sh), true)
		yt := NewTensor(MakeLoadBuffer(BufferOf(sh, bt.RangeOf[float32](10., 19., 1.).Slice()), sh), true)

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

		xt := NewTensor(MakeLoadBuffer(xs, sh), true)

		zt := xt.Relu()
		fmt.Println(zt)

		fmt.Println(zt.Data().Realize())

		zt.Mean(nil, false).Backward()

		zg := zt.grad.Data()

		zgt := zg.Realize()
		fmt.Println(zgt)
	}
}

func TestBobNet(t *testing.T) {
	sh := Shape{3, 3}

	xt := NewTensor(MakeLoadBuffer(BufferOf(sh, bt.RangeTo[float32](9.).Slice()), sh), true)
	l1t := NewTensor(MakeLoadBuffer(BufferOf(sh, bt.RangeOf[float32](10., 19., 1.).Slice()), sh), true)
	l2t := NewTensor(MakeLoadBuffer(BufferOf(sh, bt.RangeOf[float32](20., 29., 1.).Slice()), sh), true)

	zt := xt.Dot(l1t).Relu().Dot(l2t).LogSoftmax()
}
