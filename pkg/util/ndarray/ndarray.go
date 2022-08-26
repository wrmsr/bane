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

	l := v.Len()
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

func (a NdArray[T]) Transpose(axes ...int) NdArray[T] {
	check.Equal(len(axes), a.v.Order())

	sh := a.v.Shape()
	nsh := NewMutDims(len(axes))
	for i, axis := range axes {
		nsh.Set(i, sh.Get(axis))
	}

	b := New[T](Shape{nsh.Decay()})

	asl := make([]Dim, len(axes))
	bsl := make([]Dim, len(axes))

	var rec func(i int)
	rec = func(i int) {
		if i < len(axes) {
			n := sh.Get(i)
			for j := Dim(0); j < n; j++ {
				asl[i] = j
				bsl[axes[i]] = j
				rec(i + 1)
			}
		} else {
			*b.At(asl...) = a.Get(bsl...)
		}
	}

	rec(0)
	return b
}
