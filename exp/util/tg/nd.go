package tg

import (
	"math"

	"github.com/wrmsr/bane/pkg/util/check"
	nd "github.com/wrmsr/bane/pkg/util/ndarray"
	"github.com/wrmsr/bane/pkg/util/slices"
)

func NdDotSep(a, b nd.NdArray[float32]) nd.NdArray[float32] {
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

	check.Equal(a.View().Shape().Order(), 2)
	check.Equal(b.View().Shape().Order(), 2)
	check.Equal(a.View().Shape().Get(1), b.View().Shape().Get(0))

	z := a.View().Shape().Get(1)
	h := a.View().Shape().Get(0)
	w := b.View().Shape().Get(1)
	c := nd.New[float32](nd.ShapeOf(h, w))

	for i := nd.Dim(0); i < h; i++ {
		for j := nd.Dim(0); j < w; j++ {
			var o float32
			for k := nd.Dim(0); k < z; k++ {
				av := a.Get(i, k)
				bv := b.Get(k, j)
				o += av * bv
			}
			*c.At(i, j) = o
		}
	}

	return c
}

func NdDotFma(a, b nd.NdArray[float32]) nd.NdArray[float32] {
	check.Equal(a.View().Shape().Order(), 2)
	check.Equal(b.View().Shape().Order(), 2)
	check.Equal(a.View().Shape().Get(1), b.View().Shape().Get(0))

	z := a.View().Shape().Get(1)
	h := a.View().Shape().Get(0)
	w := b.View().Shape().Get(1)
	c := nd.New[float32](nd.ShapeOf(h, w))

	for i := nd.Dim(0); i < h; i++ {
		for j := nd.Dim(0); j < w; j++ {
			var o float64
			for k := nd.Dim(0); k < z; k++ {
				av := a.Get(i, k)
				bv := b.Get(k, j)
				o = math.FMA(float64(av), float64(bv), o)
			}
			*c.At(i, j) = float32(o)
		}
	}

	return c
}

var NdDot = NdDotSep

func NdTensorDot(a, b nd.NdArray[float32], axes_a, axes_b nd.Dims) nd.NdArray[float32] {
	mutaxes_a := axes_a.Mutate()
	mutaxes_b := axes_b.Mutate()

	na := mutaxes_a.Len()
	nb := mutaxes_b.Len()

	as_ := a.View().Shape()
	nda := nd.Dim(as_.Order())

	bs := b.View().Shape()
	ndb := nd.Dim(bs.Order())

	equal := true
	if na != nb {
		equal = false
	} else {
		for k := 0; k < na; k++ {
			if as_.Get(int(mutaxes_a.Get(k))) != bs.Get(int(mutaxes_b.Get(k))) {
				equal = false
				break
			}
			if mutaxes_a.Get(k) < 0 {
				mutaxes_a.Set(k, mutaxes_a.Get(k)+nda)
			}
			if mutaxes_b.Get(k) < 0 {
				mutaxes_b.Set(k, mutaxes_b.Get(k)+ndb)
			}
		}
	}

	if !equal {
		panic("shape-mismatch for sum")
	}

	axes_a = mutaxes_a.Decay()
	axes_b = mutaxes_b.Decay()

	// Move the axes to sum over to the end of "a" and to the front of "b"
	var notin []nd.Dim
	for k := nd.Dim(0); k < nda; k++ {
		if !axes_a.Contains(k) {
			notin = append(notin, k)
		}
	}
	newaxes_a := slices.Join(notin, axes_a.Slice())

	N2 := nd.Dim(1)
	for i := 0; i < axes_a.Len(); i++ {
		axis := axes_a.Get(i)
		N2 *= as_.Get(int(axis))
	}
	newshape_a := []nd.Dim{1, N2}
	for _, ax := range notin {
		newshape_a[0] *= as_.Get(int(ax))
	}
	var olda []nd.Dim
	for _, axis := range notin {
		olda = append(olda, as_.Get(int(axis)))
	}

	notin = nil
	for k := nd.Dim(0); k < ndb; k++ {
		if !axes_b.Contains(k) {
			notin = append(notin, k)
		}
	}
	newaxes_b := slices.Join(axes_b.Slice(), notin)
	N2 = 1
	for i := 0; i < axes_b.Len(); i++ {
		axis := axes_b.Get(i)
		N2 *= bs.Get(int(axis))
	}
	newshape_b := []nd.Dim{N2, 1}
	for _, ax := range notin {
		newshape_b[1] *= bs.Get(int(ax))
	}
	var oldb []nd.Dim
	for _, axis := range notin {
		oldb = append(oldb, bs.Get(int(axis)))
	}

	at_ := a.Transpose(nd.DimsOf(newaxes_a...))
	bt_ := b.Transpose(nd.DimsOf(newaxes_b...))

	at := at_.Reshape(nd.ShapeOf(newshape_a...))
	bt := bt_.Reshape(nd.ShapeOf(newshape_b...))

	res := NdDot(at, bt)

	nsh := nd.ShapeOf(slices.Join(olda, oldb)...)
	return res.Reshape(nsh)
}
