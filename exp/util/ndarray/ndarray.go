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
	if len(bs) > len(a.sh) {
		panic(fmt.Errorf("slice dimension mismatch"))
	}

	nd := len(a.sh)
	rs := make([]DimRange, len(bs))
	for i, rb := range bs {
		b := AsBound(rb)
		r := b.Range(a.sh[i])
		if r.Scalar().Present() {
			nd--
		}
		rs[i] = r
	}

	nsh := make(Shape, nd)
	nst := make(Strides, nd)
	no := a.o

	j := 0
	for i, r := range rs {
		if s := r.Scalar(); s.Present() {
			no += s.Value() * a.st[i]
		} else {
			if r.Step == 1 {
				nsh[j] = r.Stop - r.Step
				nst[j] = a.st[i]
				no += r.Start * a.st[i]
			} else {
				panic("nyi")
			}
			j++
		}
	}

	do := len(a.sh) - nd
	copy(nsh[j:], a.sh[do:])
	copy(nst[j:], a.st[do:])

	return NdArray[T]{
		sh: nsh,
		st: nst,
		o:  no,
		s:  a.s,
	}
}
