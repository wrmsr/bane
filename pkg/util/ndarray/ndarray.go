package ndarray

import (
	"fmt"

	"github.com/wrmsr/bane/pkg/util/check"
)

//

type NdArray[T any] struct {
	v View
	d []T
}

//

func New[T any](shape Shape) NdArray[T] {
	return Of[T](shape, Strides{}, 0, nil)
}

func Of[T any](
	shape Shape,
	strides Strides,
	offset Dim,
	data []T,
) NdArray[T] {
	v := ViewOf(shape, strides, offset)

	l := v.DataSize()
	if data == nil {
		data = make([]T, l+v.o)
	}
	if Dim(len(data))-offset < l {
		panic(fmt.Errorf("size mismatch"))
	}

	return NdArray[T]{
		v: v,
		d: data,
	}
}

type Maker[T any] struct {
	Shape   Shape
	Strides Strides
	Offset  Dim
	Data    []T
}

func (m Maker[T]) Make() NdArray[T] {
	return Of[T](
		m.Shape,
		m.Strides,
		m.Offset,
		m.Data,
	)
}

//

func (a NdArray[T]) View() View { return a.v }
func (a NdArray[T]) Data() []T  { return a.d }

func (a NdArray[T]) At(idxs ...Dim) *T {
	return &a.d[a.v.Index(idxs...)]
}

func (a NdArray[T]) Get(idxs ...Dim) T {
	return *a.At(idxs...)
}

//

func (a NdArray[T]) Slice(bs ...any) NdArray[T] {
	nv := a.v.Slice(bs...)
	if nv.Equals(a.v) {
		return a
	}

	return NdArray[T]{
		v: nv,
		d: a.d,
	}
}

func (a NdArray[T]) Squeeze() NdArray[T] {
	nv := a.v.Squeeze()
	if nv.Equals(a.v) {
		return a
	}

	return NdArray[T]{
		v: nv,
		d: a.d,
	}
}

func (a NdArray[T]) Transpose(axes Dims) NdArray[T] {
	check.Equal(axes.Len(), a.v.sh.Order())

	sh := a.v.Shape()
	nsh := NewMutDims(axes.Len())
	for i := 0; i < axes.Len(); i++ {
		nsh.Set(i, sh.Get(int(axes.Get(i))))
	}

	b := New[T](Shape{nsh.Decay()})

	asl := make([]Dim, axes.Len())
	bsl := make([]Dim, axes.Len())

	var rec func(i int)
	rec = func(i int) {
		if i < axes.Len() {
			n := nsh.Get(i)
			for j := Dim(0); j < n; j++ {
				asl[axes.Get(i)] = j
				bsl[i] = j
				rec(i + 1)
			}
		} else {
			*b.At(bsl...) = a.Get(asl...)
		}
	}

	rec(0)
	return b
}

func (a NdArray[T]) SwapAxes(x, y Dim) NdArray[T] {
	axes := MutDimsTo(a.v.sh.Order())
	axes.Set(int(x), y)
	axes.Set(int(y), x)
	return a.Transpose(axes.Decay())
}

func (a NdArray[T]) Reshape(nsh Shape) NdArray[T] {
	l := a.v.sh.Prod()
	check.Equal(l, nsh.Prod())

	// FIXME: everything lol (contiguous)
	s := make([]T, l)

	o := a.v.sh.Order()
	sl := make([]Dim, a.v.sh.Order())
	p := 0

	var rec func(int)
	rec = func(i int) {
		if i < o {
			n := a.v.sh.Get(i)
			for j := Dim(0); j < n; j++ {
				sl[i] = j
				rec(i + 1)
			}
		} else {
			s[p] = a.Get(sl...)
			p++
		}
	}
	rec(0)

	return Maker[T]{
		Shape: nsh,
		Data:  s,
	}.Make()
}
