package tensor

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	nd "github.com/wrmsr/bane/pkg/util/ndarray"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

type AxisPair struct {
	A, B int
}

func NdTd(a, b nd.NdArray[float32], axes ...AxisPair) nd.NdArray[float32] {
	ash, bsh := a.View().Shape(), b.View().Shape()

	var amsk, bmsk int64
	for _, x := range axes {
		check.Equal(ash.Get(x.A), bsh.Get(x.B))
		amsk |= int64(1) << x.A
		bmsk |= int64(1) << x.B
	}

	nshm := nd.NewMutDims(a.View().Shape().Order() + b.View().Shape().Order() - 2*len(axes))
	p := 0
	for i := 0; i < ash.Len(); i++ {
		if amsk&(1<<i) == 0 {
			nshm.Set(p, ash.Get(i))
			p++
		}
	}
	for i := 0; i < bsh.Len(); i++ {
		if bmsk&(1<<i) == 0 {
			nshm.Set(p, bsh.Get(i))
			p++
		}
	}
	check.Equal(p, nshm.Len())
	nsh := nshm.Decay()

	fmt.Println(nsh)
	panic(nsh)
}

func NdRange(sh nd.Shape) nd.NdArray[float32] {
	return nd.Maker[float32]{
		Shape: sh,
		Data:  bt.RangeTo[float32](float32(sh.Size())).Slice(),
	}.Make()
}

func TestNdTd2(t *testing.T) {
	NdTd(
		NdRange(nd.ShapeOf(2, 3, 5, 1)),
		NdRange(nd.ShapeOf(3, 2, 4)),
		AxisPair{0, 1},
	)
}

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

	// def matrange(*sh): return np.arange(math.prod(sh)).reshape(*sh
	// https://stackoverflow.com/questions/41870228/understanding-tensordot

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
