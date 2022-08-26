package tensor

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	nd "github.com/wrmsr/bane/pkg/util/ndarray"
)

func NdMatMul(a, b nd.NdArray[float32]) nd.NdArray[float32] {
	check.Equal(a.View().Rank(), 2)
	check.Equal(b.View().Rank(), 2)
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
	a := nd.Of[float32](
		nd.ShapeOf(3, 3),
		nd.Strides{},
		0,
		[]float32{
			0, 1, 2,
			3, 4, 5,
			6, 7, 8,
		},
	)

	b := nd.Of[float32](
		nd.ShapeOf(3, 2),
		nd.Strides{},
		0,
		[]float32{
			10, 11,
			13, 14,
			16, 17,
		},
	)

	c := NdMatMul(a, b)
	fmt.Println(c)
}

func TestNdTd(t *testing.T) {
	// [[[[[[0.]]],,,   [[[3.]]],,,   [[[6.]]]],,,,  [[[[1.]]],,,   [[[4.]]],,,   [[[7.]]]],,,,  [[[[2.]]],,,   [[[5.]]],,,   [[[8.]]]]]]
	txg := nd.Of[float32](
		nd.ShapeOf(3, 3, 1, 1, 1), // cin, oy, ox, H, W,
		nd.Strides{},
		0,
		[]float32{0, 3, 6, 1, 4, 7, 2, 5, 8},
	)

	fmt.Println(txg)

	// [[[[10.]],,  [[13.]],,  [[16.]]],,, [[[11.]],,  [[14.]],,  [[17.]]],,, [[[12.]],,  [[15.]],,  [[18.]]]]
	twg := nd.Of[float32](
		nd.ShapeOf(3, 3, 1, 1), // cout, cin, H, W
		nd.Strides{},
		0,
		[]float32{10, 13, 16, 11, 14, 17, 12, 15, 18},
	)

	fmt.Println(twg)

	// [[[[ 45.  48.  51.]],,  [[162. 174. 186.]],,  [[279. 300. 321.]]]]
	want := nd.Of[float32](
		nd.ShapeOf(3, 1, 3),
		nd.Strides{},
		0,
		[]float32{45, 48, 51, 162, 174, 186, 279, 300, 321},
	)

	fmt.Println(want)

	// np.tensordot(txg, twg, ((0, 3, 4), (1, 2, 3)))

	// (1, 4, 5)
	// (1, 2, 3)
	//                    0    1    2    3    4    5
	// 0 = {tuple: 6} (   1, 784,  69,   1,   1,   1 )
	// 1 = {tuple: 4} ( 128, 784,   1,   1)
	// 2 = {tuple: 4} (   1,  69,   1, 128)
	/*
		H, W = 1, 1
		groups = 1
		rcout = 128
		cin = 784
		oy, ox = 69, 1
		iy, ix = 69, 1
		sy, sx = 1, 1
		bs = 1
		cout = 128
		out_shape = (1, 128, 69, 1)
		py, py_ = 0, 0
		px, px_ = 0, 0
		dy, dx = 1, 1
	*/
}
