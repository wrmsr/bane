package tg

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
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

//func TestBobNet2(t *testing.T) {
//	rj := func(n string) []byte {
//		p := filepath.Join(
//			paths.FindProjectRoot(),
//			fmt.Sprintf("../../geohot/tinygrad/%s.json", n),
//		)
//
//		return check.Must1(os.ReadFile(p))
//	}
//
//	rm := func(n string) [][]float32 {
//		b := rj(n)
//
//		var m [][]float32
//		check.Must(json.Unmarshal(b, &m))
//		return m
//	}
//
//	l1m := rm("l1")
//	l2m := rm("l2")
//
//	m2t := func(m [][]float32) nd.NdArray[float32] {
//
//	}
//
//	xt := NewTensor(MakeLoadBuffer(BufferOf(sh, bt.RangeTo[float32](9.).Slice())), true).Reshape(sh)
//	l1t := NewTensor(MakeLoadBuffer(BufferOf(sh, bt.RangeOf[float32](10., 19., 1.).Slice())), true).Reshape(sh)
//	l2t := NewTensor(MakeLoadBuffer(BufferOf(sh, bt.RangeOf[float32](20., 29., 1.).Slice())), true).Reshape(sh)
//
//	zt := xt.Dot(l1t).Relu().Dot(l2t).LogSoftmax()
//	//zt := xt.Dot(l1t).Relu().Dot(l2t)
//	fmt.Println("====")
//
//	dumpObj(zt)
//
//	//fmt.Println(zt.Mean(nil, false).Data().Realize())
//	//fmt.Println(zt.Mean(nil, false).Data().Realize())
//
//	// garbage
//
//	scc := func(out *Tensor, y_ []int) *Tensor {
//		check.Equal(Dim(len(y_)), out.Shape()[0])
//		check.Equal(out.Shape()[0], Dim(len(y_)))
//		num_classes := out.Shape()[len(out.Shape())-1]
//		yb := NewBuffer(out.Shape())
//		for i := Dim(0); i < Dim(len(y_)); i++ {
//			yb.set(float32(-1*num_classes), i, Dim(y_[i]))
//		}
//		y := NewTensor(MakeLoadBuffer(yb), false)
//		return out.Mul(y).Mean(nil, false)
//	}
//
//	y := make([]int, 3)
//	for i := 0; i < 3; i++ {
//		y[i] = i % 3
//	}
//
//	z := scc(
//		zt,
//		y,
//	)
//
//	fmt.Println(z.Data().Realize().Nd())
//	z.Backward()
//
//}

//

type Float32GrayscaleImage struct {
	a nd.NdArray[float32]
}

var _ image.Image = Float32GrayscaleImage{}

func (i Float32GrayscaleImage) ColorModel() color.Model {
	return color.GrayModel
}

func (i Float32GrayscaleImage) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: int(i.a.View().Shape().Get(1)),
			Y: int(i.a.View().Shape().Get(0)),
		},
	}
}

func (i Float32GrayscaleImage) At(x, y int) color.Color {
	v := i.a.Get(nd.Dim(y), nd.Dim(x))
	return color.Gray{uint8(v)}
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

	ndb := nd.Maker[byte]{Data: b}.Make().Reshape(nd.ShapeOf(-1, 28, 28))
	f := nd.Map(func(b byte) float32 { return float32(b) }, ndb)

	i := Float32GrayscaleImage{f.Slice(0)}
	wb := bytes.NewBuffer(nil)
	check.Must(jpeg.Encode(wb, i, nil))

	tf := check.Must1(ioutil.TempFile("", "foo"))
	//check.Must1(os.CreateTemp("", "foo.jpg"))

	_ = check.Must1(tf.Write(check.Must1(io.ReadAll(wb))))

	nn := tf.Name() + ".jpg"
	check.Must(os.Rename(tf.Name(), nn))

	check.Must(exec.Command("open", nn).Run())
}
