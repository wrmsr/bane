package tensor

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/wrmsr/bane/core/check"
	"github.com/wrmsr/bane/core/slices"
	bt "github.com/wrmsr/bane/core/types"
)

//

type BufferTensor struct {
	BaseTensor

	buf []byte
}

func NewHeapTensor(bt BaseTensor) BufferTensor {
	return BufferTensor{
		BaseTensor: bt,

		buf: make([]byte, bt.Sz()),
	}
}

//

//func TestBufferTensor(t *testing.T) {
//	tns := NewHeapTensor(
//		NewBaseTensor(Sz(4), Shape{4, 4, 4}, RowMajor),
//	)
//
//	fmt.Println(tns)
//
//	fmt.Println(tns.strides.Offset(1, 3))
//	fmt.Println(tns.strides.Offset(2))
//	fmt.Println(tns.strides.Offset(2, 3))
//	fmt.Println(tns.strides.Offset(2, 3, 0))
//	fmt.Println(tns.strides.Offset(2, 3, 1))
//	fmt.Println(tns.strides.Offset(3, 3, 3))
//
//	for i := int64(0); i < int64(tns.Sz()); i += int64(tns.width) {
//		b := make([]byte, tns.width)
//		platform.NativeEndian().PutUint32(b, math.Float32bits(float32(i)))
//		copy(tns.buf[i:i+int64(tns.width)], b)
//	}
//
//	fmt.Println(tns.buf)
//
//	for i := Dim(0); i < tns.shape[0]; i++ {
//		for j := Dim(0); j < tns.shape[1]; j++ {
//			for k := Dim(0); k < tns.shape[2]; k++ {
//				ofs := tns.strides.Offset(i, j, k)
//				f := math.Float32frombits(platform.NativeEndian().Uint32(tns.buf[ofs : ofs+tns.width]))
//				fmt.Printf("%2d %2d %2d %3d %12f\n", i, j, k, ofs, f)
//			}
//		}
//	}
//}

//

type Float32Tensor struct {
	BaseTensor

	s []float32
}

func NewFloat32Tensor(sh Shape, s []float32) *Float32Tensor {
	if s == nil {
		s = make([]float32, sh.Dim())
	} else {
		check.Condition(Dim(len(s)) == sh.Dim())
	}

	return &Float32Tensor{
		BaseTensor: NewBaseTensor(sh),

		s: s,
	}
}

//func (t *Float32Tensor) Add(o *Float32Tensor) {
//	check.Condition(t.shape.Equals(o.shape))
//}

func (t *Float32Tensor) Clone() *Float32Tensor {
	return NewFloat32Tensor(t.shape, slices.Clone(t.s))
}

//

type Op interface {
	isOp()
}

//

type LoadOp struct{ t *Float32Tensor }

func (o LoadOp) isOp() {}

//

type ReshapeOp struct{ sh Shape }

func (o ReshapeOp) isOp() {}

//

type LazyTensor struct {
	BaseTensor

	data *LazyTensor
	op   Op

	noGrad bool
}

func (t LazyTensor) NoGrad() bool { return t.noGrad }

type FunctionContext struct {
	parents []LazyTensor
	noGrad  bool

	saved []LazyTensor
}

func NewFunctionContext(parents ...LazyTensor) FunctionContext {
	return FunctionContext{
		parents: parents,
		noGrad:  slices.All(parents, LazyTensor.NoGrad),
	}
}

func Permute(x LazyTensor, order []int) (LazyTensor, FunctionContext) {
	panic("no")
}

func (t *Float32Tensor) Dot(o *Float32Tensor) *Float32Tensor {
	bs := bt.Prod[Dim](t.shape[0 : len(t.shape)-2]...)
	groups := bt.Prod[Dim](o.shape[0 : len(t.shape)-2]...)
	cin, cout := o.shape.At(-2), o.shape.At(-1)
	outShapeT := slices.Join(t.shape[0:len(t.shape)-2], []Dim{cout, -1})
	fmt.Println(bs)
	fmt.Println(groups)
	fmt.Println(cin)
	fmt.Println(outShapeT)

	var order Shape
	if len(t.shape) > 1 {
		order = slices.Join(bt.RangeTo[Dim](Dim(len(t.shape))-2).Slice(), []Dim{Dim(len(t.shape)) - 1, Dim(len(t.shape)) - 2})
	} else {
		order, outShapeT = Shape{0}, Shape{cout}
	}

	worder := slices.Join(bt.RangeTo[Dim](Dim(len(o.shape))-2).Slice(), []Dim{Dim(len(o.shape)) - 1, Dim(len(o.shape)) - 2})

	/*
	   cx = x.transpose(order=order).reshape(shape=(bs//groups, groups*cin, -1, 1))
	   cw = w.transpose(order=worder).reshape(shape=(groups*cout, cin, 1, 1))
	   return cx.conv2d(cw, groups=groups).reshape(shape=out_shape_t).transpose(order=order)
	*/

	fmt.Println(order)
	fmt.Println(worder)
	panic("foo")
}

func RandFloat32Tensor(sh Shape) *Float32Tensor {
	s := make([]float32, sh.Dim())
	for i := range s {
		s[i] = rand.Float32()
	}
	return NewFloat32Tensor(sh, s)
}

func TestFloat32Tensor(t *testing.T) {
	//t := RandFloat32Tensor()

	x_init := RandFloat32Tensor(Shape{1, 3})
	//U_init := RandFloat32Tensor(Shape{3, 3})
	//V_init := RandFloat32Tensor(Shape{3, 3})
	W_init := RandFloat32Tensor(Shape{3, 3})
	//m_init := RandFloat32Tensor(Shape{1, 3})

	x := x_init.Clone()
	W := W_init.Clone()
	//m := m_init.Clone()

	out := x.Dot(W)
	fmt.Println(out)
}
