package ndarray

import (
	"fmt"
)

type NdArray[T any] struct {
	sh Shape
	st Strides
	o  Dim
	s  []T
}

func (a NdArray[T]) Shape() Shape     { return a.sh }
func (a NdArray[T]) Strides() Strides { return a.st }
func (a NdArray[T]) Offset() Dim      { return a.o }
func (a NdArray[T]) Data() []T        { return a.s }

func New[T any](shape Shape) NdArray[T] {
	return Of[T](shape, nil, 0, nil)
}

func Of[T any](
	shape Shape,
	strides Strides,
	offset Dim,
	data []T,
) NdArray[T] {
	if len(strides) < 1 {
		strides = calcStrides(shape)
	}
	if len(strides) != len(shape) {
		panic(fmt.Errorf("strides mismatch"))
	}

	sz := strides[0] * shape[0]
	if data == nil {
		data = make([]T, sz)
	}
	if len(data)-offset < sz {
		panic(fmt.Errorf("size mismatch"))
	}

	return NdArray[T]{
		sh: shape,
		st: strides,
		o:  offset,
		s:  data,
	}
}

func (a NdArray[T]) At(idxs ...Dim) *T {
	return &a.s[a.st.Offset(idxs)+a.o]
}

func (a NdArray[T]) Get(idxs ...Dim) T {
	return *a.At(idxs...)
}

func (a NdArray[T]) Slice(bs ...any) NdArray[T] {
	nd := len(a.sh)
	if len(bs) != nd {
		panic(fmt.Errorf("slice dimension mismatch"))
	}

	nsh := make(Shape, nd)
	nst := make(Strides, nd)
	no := a.o

	for i := nd - 1; i >= 0; i-- {
		r := CalcRange(bs[i], a.sh[i])
		if r.Step == 1 {
			nsh[i] = r.Stop - r.Start
			nst[i] = a.st[i]
			no += r.Start * a.st[i]
		} else {
			panic("nyi")
		}
	}

	return NdArray[T]{
		sh: nsh,
		st: nst,
		o:  no,
		s:  a.s,
	}
}

func (a NdArray[T]) Squeeze() NdArray[T] {
	nsd := a.sh.NumScalarDims()
	if nsd < 1 {
		return a
	}

	nd := len(a.sh)
	nsh := make(Shape, nd-nsd)
	nst := make(Strides, nd-nsd)

	i := 0
	for j := 0; j < nd; j++ {
		if a.sh[j] == 1 {
			continue
		}
		nst[i] = a.st[j]
		nsh[i] = a.sh[j]
		i++
	}

	return NdArray[T]{
		sh: nsh,
		st: nst,
		o:  a.o,
		s:  a.s,
	}
}
