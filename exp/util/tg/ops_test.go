package tg

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/dev/paths"
	nd "github.com/wrmsr/bane/pkg/util/ndarray"
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

func TestMnistData(t *testing.T) {
	/*
		# In the labels file, the number of items is 32-bit big-endian integer at offset 4.
		# The labels are one byte each, binary 0..9, starting at offset 8.

		# In the images file, the number of items is 32-bit big-endian integer at offset 4.
		# The image height is 32-bit big-endian integer at offset 8.
		# The image width is 32-bit big-endian integer at offset 0xc.
		# The images are height*width bytes each, starting at offset 0x10.
	*/
	p := filepath.Join(
		paths.FindProjectRoot(),
		"../../geohot/tinygrad/datasets/mnist/train-images-idx3-ubyte.gz",
	)
	gb := check.Must1(os.ReadFile(p))
	gr := check.Must1(gzip.NewReader(bytes.NewReader(gb)))
	b := check.Must1(io.ReadAll(gr))[0x10:]
	fmt.Println(len(b))

	ndb := nd.Maker[byte]{Data: b}.Make().Reshape(nd.ShapeOf(-1, 28*28))
	f := nd.Map(func(b byte) float32 { return float32(b) }, ndb)
	fmt.Println(f.View())
}
