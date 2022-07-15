package tensor

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/platform"
	"github.com/wrmsr/bane/pkg/util/slices"
	bt "github.com/wrmsr/bane/pkg/util/types"
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

func TestBufferTensor(t *testing.T) {
	tns := NewHeapTensor(
		NewBaseTensor(Sz(4), Shape{4, 4, 4}, RowMajor),
	)

	fmt.Println(tns)

	fmt.Println(tns.strides.Offset(1, 3))
	fmt.Println(tns.strides.Offset(2))
	fmt.Println(tns.strides.Offset(2, 3))
	fmt.Println(tns.strides.Offset(2, 3, 0))
	fmt.Println(tns.strides.Offset(2, 3, 1))
	fmt.Println(tns.strides.Offset(3, 3, 3))

	for i := int64(0); i < int64(tns.Sz()); i += int64(tns.width) {
		b := make([]byte, tns.width)
		platform.NativeEndian().PutUint32(b, math.Float32bits(float32(i)))
		copy(tns.buf[i:i+int64(tns.width)], b)
	}

	fmt.Println(tns.buf)

	for i := Dim(0); i < tns.shape[0]; i++ {
		for j := Dim(0); j < tns.shape[1]; j++ {
			for k := Dim(0); k < tns.shape[2]; k++ {
				ofs := tns.strides.Offset(i, j, k)
				f := math.Float32frombits(platform.NativeEndian().Uint32(tns.buf[ofs : ofs+tns.width]))
				fmt.Printf("%2d %2d %2d %3d %12f\n", i, j, k, ofs, f)
			}
		}
	}
}

//

type Float32Tensor struct {
	BaseTensor

	s []float32

	noGrad bool
}

func NewFloat32Tensor(sh Shape, s []float32) *Float32Tensor {
	if s == nil {
		s = make([]float32, sh.Dim())
	} else {
		check.Condition(Dim(len(s)) == sh.Dim())
	}

	return &Float32Tensor{
		BaseTensor: NewBaseTensor(Sz(4), sh, RowMajor),

		s: s,
	}
}

func (t *Float32Tensor) Add(o *Float32Tensor) {
	check.Condition(t.shape.Equals(o.shape))
}

func (t *Float32Tensor) Clone() *Float32Tensor {
	return NewFloat32Tensor(t.shape, slices.Clone(t.s))
}

func (t *Float32Tensor) Dot(o *Float32Tensor) *Float32Tensor {
	bs := bt.Prod(t.shape[0:-2]...)
	groups := bt.Prod(o.shape[0:-2]...)
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
	m_init := RandFloat32Tensor(Shape{1, 3})

	x := x_init.Clone()
	W := W_init.Clone()
	m := m_init.Clone()
	m.noGrad = true

	out := x.Dot(W)
	fmt.Println(out)
}
