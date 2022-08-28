package tensor

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	nd "github.com/wrmsr/bane/pkg/util/ndarray"
	"github.com/wrmsr/bane/pkg/util/slices"
)

type AxisPair struct {
	A, B int
}

/*
In [267]: np.tensordot(matrange(2,2,3),matrange(2,2),(0,1))
Out[267]:
array([[[ 6, 18],
        [ 7, 23],
        [ 8, 28]],
       [[ 9, 33],
        [10, 38],
        [11, 43]]])

In [268]: np.tensordot(matrange(2,2),matrange(2,2,3),(0,1))
Out[268]:
array([[[ 6,  8, 10],
        [18, 20, 22]],
       [[ 9, 13, 17],
        [33, 37, 41]]])

In [269]: np.tensordot(matrange(2,2,3),matrange(2,2),(1,0))
Out[269]:
array([[[ 6,  9],
        [ 8, 13],
        [10, 17]],
       [[18, 33],
        [20, 37],
        [22, 41]]])

In [270]: np.tensordot(matrange(2,2),matrange(2,2,3),(1,0))
Out[270]:
array([[[ 6,  7,  8],
        [ 9, 10, 11]],
       [[18, 23, 28],
        [33, 38, 43]]])
*/

func NdTd(a, b nd.NdArray[float32], axes ...AxisPair) nd.NdArray[float32] {
	ash := a.View().Shape()
	bsh := b.View().Shape()

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
	nsh := nd.Shape{nshm.Decay()}

	c := nd.New[float32](nsh)

	ax := make([]any, ash.Order())
	bx := make([]any, bsh.Order())
	cx := make([]nd.Dim, nsh.Order())

	sum := func() {
		aq := a.Slice(ax...).Squeeze()
		bq := b.Slice(bx...).Squeeze()

		fmt.Println(aq)
		fmt.Println(bq)

		l := aq.View().Shape().Get(0)
		check.Equal(l, bq.View().Shape().Get(0))

		var x float32
		for j := nd.Dim(0); j < l; j++ {
			ae := aq.Get(j)
			be := bq.Get(j)
			x += ae * be
		}
		*c.At(cx...) = x
	}

	var brec func(int, int)
	brec = func(i, o int) {
		if i >= bsh.Order() {
			sum()
		} else if bmsk&(1<<i) != 0 {
			brec(i+1, o)
		} else {
			n := bsh.Get(i)
			for j := nd.Dim(0); j < n; j++ {
				bx[i] = j
				cx[o] = j
				brec(i+1, o+1)
			}
		}
	}

	var arec func(int, int)
	arec = func(i, o int) {
		if i >= ash.Order() {
			brec(0, o)
		} else if amsk&(1<<i) != 0 {
			arec(i+1, o)
		} else {
			n := ash.Get(i)
			for j := nd.Dim(0); j < n; j++ {
				ax[i] = j
				cx[o] = j
				arec(i+1, o+1)
			}
		}
	}

	arec(0, 0)
	return c
}

func NdTd2(a, b nd.NdArray[float32], axes_a, axes_b []int) nd.NdArray[float32] {
	na := len(axes_a)
	nb := len(axes_b)

	as_ := a.View().Shape()
	nda := as_.Order()

	bs := b.View().Shape()
	ndb := bs.Order()

	equal := true
	if na != nb {
		equal = false
	} else {
		for k := 0; k < na; k++ {
			if as_.Get(axes_a[k]) != bs.Get(axes_b[k]) {
				equal = false
				break
			}
			if axes_a[k] < 0 {
				axes_a[k] += nda
			}
			if axes_b[k] < 0 {
				axes_b[k] += ndb
			}
		}
	}

	if !equal {
		panic("shape-mismatch for sum")
	}

	// Move the axes to sum over to the end of "a" and to the front of "b"
	var notin []int
	for k := 0; k < nda; k++ {
		if !slices.Contains(axes_a, k) {
			notin = append(notin, k)
		}
	}
	newaxes_a := slices.Join(notin, axes_a)

	N2 := nd.Dim(1)
	for _, axis := range axes_a {
		N2 *= as_.Get(axis)
	}
	newshape_a := []nd.Dim{1, N2}
	for _, ax := range notin {
		newshape_a[0] *= as_.Get(ax)
	}
	var olda []nd.Dim
	for _, axis := range notin {
		olda = append(olda, as_.Get(axis))
	}

	notin = nil
	for k := 0; k < ndb; k++ {
		if !slices.Contains(axes_b, k) {
			notin = append(notin, k)
		}
	}
	newaxes_b := slices.Join(axes_b, notin)
	N2 = 1
	for _, axis := range axes_b {
		N2 *= bs.Get(axis)
	}
	newshape_b := []nd.Dim{N2, 1}
	for _, ax := range notin {
		newshape_b[1] *= bs.Get(ax)
	}
	var oldb []nd.Dim
	for _, axis := range notin {
		oldb = append(oldb, bs.Get(axis))
	}

	at_ := a.Transpose(nd.DimsOf(nd.IntDims(newaxes_a...)...))
	bt_ := b.Transpose(nd.DimsOf(nd.IntDims(newaxes_b...)...))

	fmt.Println(at_)
	fmt.Println(bt_)

	at := at_.Reshape(nd.ShapeOf(newshape_a...))
	bt := bt_.Reshape(nd.ShapeOf(newshape_b...))

	res := NdDot(at, bt)

	nsh := nd.ShapeOf(slices.Join(olda, oldb)...)
	return res.Reshape(nsh)
}

func TestNdTranspose(t *testing.T) {
	a := nd.OfRange[float32](nd.ShapeOf(2, 3, 5))
	at := a.Transpose(nd.DimsOf(2, 0, 1))
	fmt.Println(at)
}

func NdDot(a, b nd.NdArray[float32]) nd.NdArray[float32] {
	/*
		Dot product of two arrays. Specifically,
		- If both a and b are 1-D arrays, it is inner product of vectors (without complex conjugation).
		- If both a and b are 2-D arrays, it is matrix multiplication, but using matmul or a @ b is preferred.
		- If either a or b is 0-D (scalar), it is equivalent to multiply and using numpy.multiply(a, b) or a * b is
		  preferred.
		- If a is an N-D array and b is a 1-D array, it is a sum product over the last axis of a and b.
		- If a is an N-D array and b is an M-D array (where M>=2), it is a sum product over the last axis of a and the
		  second-to-last axis of b:
			dot(a, b)[i,j,k,m] = sum(a[i,j,:] * b[k,:,m])
	*/
	panic("nyi")
}

func TestNdTd2(t *testing.T) {
	fmt.Println(NdTd(
		nd.OfRange[float32](nd.ShapeOf(2, 3, 5)),
		nd.OfRange[float32](nd.ShapeOf(3, 2, 4)),
		AxisPair{0, 1},
	))
}

func TestNdTd3(t *testing.T) {
	fmt.Println(NdTd(
		nd.OfRange[float32](nd.ShapeOf(2, 3, 5)),
		nd.OfRange[float32](nd.ShapeOf(3, 2, 4)),
		AxisPair{0, 1},
		AxisPair{1, 0},
	))
}

func TestNdTdX(t *testing.T) {
	fmt.Println(NdTd2(
		nd.OfRange[float32](nd.ShapeOf(2, 3, 5)),
		nd.OfRange[float32](nd.ShapeOf(3, 2, 4)),
		[]int{0, 1},
		[]int{1, 0},
	))
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
