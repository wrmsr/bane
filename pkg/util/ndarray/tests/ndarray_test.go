package tests

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
	nd "github.com/wrmsr/bane/pkg/util/ndarray"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

func TestNdArray(t *testing.T) {
	d := nd.Dim(4)
	a := nd.Maker[int]{
		Shape: nd.ShapeOf(d, d, d),
		Data:  bt.RangeTo(int(d * d * d)).Slice(),
	}.Make()
	fmt.Println(a)

	n := 0
	for i := nd.Dim(0); i < d; i++ {
		for j := nd.Dim(0); j < d; j++ {
			for k := nd.Dim(0); k < d; k++ {
				tu.AssertEqual(t, a.Get(i, j, k), n)
				n++
			}
		}
	}

	//idx := []Dim{1, 2, 0}
	//fmt.Println(a.Get(idx...))
	//fmt.Println("====")

	//*a.At(idx...) = 420
	//fmt.Println(a)
	//fmt.Println(a.Get(idx...))
	//fmt.Println("====")

	for _, sl := range [][]any{
		{1},
		{nil, 1},
		{nil, nil, 1},
		{[]any{1, nil}},
		{[]any{1, nil}, 1},
		{[]any{1, nil}, nil, 1},
		{[]any{nil, nil, 2}},
		{1, nil, []any{nil, nil, 2}},
	} {
		fmt.Println("========")
		fmt.Println(sl)
		fmt.Println("====")

		sa := a.Slice(sl...)
		fmt.Println(sa.View())
		fmt.Println(sa)
		fmt.Println("====")

		sqa := sa.Squeeze()
		fmt.Println(sqa.View())
		fmt.Println(sqa)
	}
}

func TestScalar(t *testing.T) {
	sc := nd.New[int](nd.ShapeOf(1))
	*sc.At(0) = 420
	fmt.Println(sc)
}

func BenchmarkSliceView(b *testing.B) {
	v := nd.ViewOf(
		nd.ShapeOf(3, 3, 3),
		nd.Strides{},
		0,
	)

	o := []any{1, nil, []any{1, nil, 2}}
	//o := []any{1}

	fmt.Println(b.N)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		sv := v.Slice(o...)
		//sv := v.Slice(1, nil, []any{1, nil, 2})

		if sv.Len() != 9 {
			panic("oops")
		}
	}
}

func NdMatMul(a, b nd.NdArray[float32]) nd.NdArray[float32] {
	check.Equal(a.View().Order(), 2)
	check.Equal(b.View().Order(), 2)
	check.Equal(a.View().Shape().Get(1), b.View().Shape().Get(0))

	h := a.View().Shape().Get(0)
	w := b.View().Shape().Get(1)
	c := nd.New[float32](nd.ShapeOf(h, w))

	for i := nd.Dim(0); i < h; i++ {
		for j := nd.Dim(0); j < w; j++ {
			var o float32
			for k := nd.Dim(0); k < h; k++ {
				o += a.Get(i, k) * b.Get(k, j)
			}
			*c.At(i, j) = o
		}
	}

	return c
}

func TestNdMatMul(t *testing.T) {
	a := nd.Maker[float32]{
		Shape: nd.ShapeOf(3, 3),
		Data: []float32{
			0, 1, 2,
			3, 4, 5,
			6, 7, 8,
		},
	}.Make()

	b := nd.Maker[float32]{
		Shape: nd.ShapeOf(3, 2),
		Data: []float32{
			10, 11,
			13, 14,
			16, 17,
		},
	}.Make()

	c := NdMatMul(a, b)
	fmt.Println(c)
}

func TestTranspose(t *testing.T) {
	a := nd.Maker[float32]{
		Shape: nd.ShapeOf(3, 3, 3),
		Data:  bt.RangeTo[float32](27).Slice(),
	}.Make()
	fmt.Println(a)

	for _, axes := range [][]int{
		{0, 1, 2},
		{0, 2, 1},
		{1, 0, 2},
		{1, 2, 0},
		{2, 0, 1},
		{2, 1, 0},
	} {
		fmt.Println("========")
		fmt.Println(axes)
		fmt.Println("====")
		b := a.Transpose(axes...)
		fmt.Println(b)
	}

	fmt.Println("========")
	fmt.Println("0<->1")
	fmt.Println("====")
	fmt.Println(a.SwapAxes(0, 1))
}
