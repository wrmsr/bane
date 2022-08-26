package tensor

import (
	"fmt"
	"testing"

	nd "github.com/wrmsr/bane/pkg/util/ndarray"
)

func TestNdTd(t *testing.T) {
	// [[[[[[0.]]],,,   [[[3.]]],,,   [[[6.]]]],,,,  [[[[1.]]],,,   [[[4.]]],,,   [[[7.]]]],,,,  [[[[2.]]],,,   [[[5.]]],,,   [[[8.]]]]]]
	txg := nd.Maker[float32]{
		Shape: nd.ShapeOf(3, 3, 1, 1, 1), // cin, oy, ox, H, W,
		Data: []float32{
			0, 3, 6,
			1, 4, 7,
			2, 5, 8,
		},
	}.Make()

	fmt.Println(txg)

	// [[[[10.]],,  [[13.]],,  [[16.]]],,, [[[11.]],,  [[14.]],,  [[17.]]],,, [[[12.]],,  [[15.]],,  [[18.]]]]
	twg := nd.Maker[float32]{
		Shape: nd.ShapeOf(3, 3, 1, 1), // cout, cin, H, W
		Data: []float32{
			10, 13, 16,
			11, 14, 17,
			12, 15, 18,
		},
	}.Make()

	fmt.Println(twg)

	// [[[[ 45.  48.  51.]],,  [[162. 174. 186.]],,  [[279. 300. 321.]]]]
	want := nd.Maker[float32]{
		Shape: nd.ShapeOf(3, 1, 3),
		Data: []float32{
			45, 48, 51,
			162, 174, 186,
			279, 300, 321,
		},
	}

	fmt.Println(want)

	// np.tensordot(txg, twg, ((0, 3, 4), (1, 2, 3)))

	/*
		In [164]: a = np.asarray([0, 3, 6, 1, 4, 7, 2, 5, 8]).reshape(3,3)
		In [165]: b = np.asarray([10, 13, 16, 11, 14, 17, 12, 15, 18]).reshape(3,3)
		In [166]: np.tensordot(a.reshape(3,3,1,1,1), b.reshape(3,3,1,1), ((0, 3, 4), (1, 2, 3)))
		Out[166]:
		array([[[ 45,  48,  51]],
		       [[162, 174, 186]],
		       [[279, 300, 321]]])
		In [167]: a.transpose() @ b.transpose()
		Out[167]:
		array([[ 45,  48,  51],
		       [162, 174, 186],
		       [279, 300, 321]])
	*/

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
