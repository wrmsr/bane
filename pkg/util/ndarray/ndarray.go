package ndarray

import (
	"fmt"
)

type NdArray[T any] struct {
	v View
	d []T
}

func (a NdArray[T]) View() View { return a.v }
func (a NdArray[T]) Data() []T  { return a.d }

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

func (a NdArray[T]) At(idxs ...Dim) *T {
	return &a.d[a.v.Index(idxs...)]
}

func (a NdArray[T]) Get(idxs ...Dim) T {
	return *a.At(idxs...)
}

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
