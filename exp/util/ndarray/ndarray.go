package ndarray

import (
	"fmt"

	"github.com/wrmsr/bane/pkg/util/check"
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

func (a NdArray[T]) Slice(bounds ...any) NdArray[T] {
	if len(bounds) > len(a.sh) {
		panic(fmt.Errorf("slice dimension mismatch"))
	}

	var o Dim
	for i, rb := range bounds {
		sh := a.sh[i]

		b := AsBound(rb).OrFn(func() Bound { return Bound{Stop: a.sh[i]} })
		fmt.Println(b)

		check.Between(b.Start, 0, sh)
		check.Between(b.Stop, 0, sh)
		check.Condition(b.Start <= b.Stop)
		check.Equal(b.Step, 0)
	}

	_ = o

	panic("nyi")
}
