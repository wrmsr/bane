package tg

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/dev/paths"
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

func dumpObj(o any) {
	var rec func(any, string)
	rec = func(o any, p string) {
		switch o := o.(type) {
		case *Tensor:
			rec(o.data, p)
		case *LazyBuffer:
			rec(o.op, p)
		case *LazyOp:
			fmt.Printf("%s%v\n", p, o.op)
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

	xt := NewTensor(MakeLoadBuffer(BufferOf(sh, bt.RangeTo[float32](9.).Slice()), Shape{9}), true).Reshape(sh)
	l1t := NewTensor(MakeLoadBuffer(BufferOf(sh, bt.RangeOf[float32](10., 19., 1.).Slice()), Shape{9}), true).Reshape(sh)
	l2t := NewTensor(MakeLoadBuffer(BufferOf(sh, bt.RangeOf[float32](20., 29., 1.).Slice()), Shape{9}), true).Reshape(sh)

	zt := xt.Dot(l1t).Relu().Dot(l2t).LogSoftmax()
	//zt := xt.Dot(l1t).Relu().Dot(l2t)
	fmt.Println("====")

	dumpObj(zt)

	//fmt.Println(zt.Mean(nil, false).Data().Realize())
	fmt.Println(zt.Mean(nil, false).Data().Realize())

	// garbage

	//scc := func(out, y *Tensor) *Tensor {
	//	// return out.mul(y).mean()
	//}
	//
	//y := NewTensor(MakeLoadBuffer(BufferOf(sh, bt.RangeTo[float32](3.).Slice()), Shape{3}), true)

	//zt.Mean(nil, false).Backward()
	//zg := zt.grad.Data()
	//zgt := zg.Realize()
	//fmt.Println(zgt)
}

func TestBobNet2(t *testing.T) {
	rj := func(n string) []byte {
		p := filepath.Join(
			paths.FindProjectRoot(),
			fmt.Sprintf("../../geohot/tinygrad/%s.json", n),
		)

		return check.Must1(os.ReadFile(p))
	}

	rm := func(n string) [][]float32 {
		b := rj(n)

		var m [][]float32
		check.Must(json.Unmarshal(b, &m))
		return m
	}

	l1m := rm("l1")
	l2m := rm("l2")

	fmt.Println(len(l1m))
	fmt.Println(len(l2m))
}
