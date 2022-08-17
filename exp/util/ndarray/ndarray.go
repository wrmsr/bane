package ndarray

import "fmt"

type NdArray[T any] struct {
	sh []int
	st []int
	o  int
	s  []T
}

func (a NdArray[T]) Shape() []int  { return a.sh }
func (a NdArray[T]) Stride() []int { return a.st }
func (a NdArray[T]) Offset() int   { return a.o }
func (a NdArray[T]) Data() []T     { return a.s }

func calcStrides(shape []int) []int {
	st := make([]int, len(shape))
	st[len(shape)-1] = 1
	for i := len(shape) - 2; i >= 0; i-- {
		st[i] = st[i+1] * shape[i+1]
	}
	return st
}

func NewNdArray[T any](shape []int) NdArray[T] {
	return NdArrayOf[T](shape, nil, 0, nil)
}

func NdArrayOf[T any](
	shape []int,
	strides []int,
	offset int,
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

func calcOffset(idxs []int, strides []int) int {
	if len(idxs) != len(strides) {
		panic(fmt.Errorf("dim mismatch: got %d need %d", len(idxs), len(strides)))
	}

	o := 0
	for i, idx := range idxs {
		o += i * strides[idx]
	}

	return o
}

func (a NdArray[T]) At(idxs ...int) *T {
	return &a.s[calcOffset(idxs, a.st)+a.o]
}

func (a NdArray[T]) Get(idxs ...int) T {
	return *a.At(idxs...)
}

type Bound struct {
	Start  int
	Stop   int
	Stride int
}

var Z = Bound{}

func (a NdArray[T]) Slice(bounds ...Bound) NdArray[T] {
	if len(bounds) > len(a.sh) {
		panic(fmt.Errorf("slice dimension mismatch"))
	}

	panic("nyi")
}
