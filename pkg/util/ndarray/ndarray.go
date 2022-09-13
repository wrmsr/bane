package ndarray

import (
	"fmt"

	"github.com/wrmsr/bane/pkg/util/check"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type NdArray[T any] struct {
	v View
	d []T
}

//

type Maker[T any] struct {
	Shape   Shape
	Strides Strides
	Offset  Dim
	Data    []T
	Default bt.Optional[T]
}

func (m Maker[T]) Make() NdArray[T] {
	shape := m.Shape
	data := m.Data

	if shape.Len() < 1 && m.Strides.Len() < 1 && data != nil {
		shape = ShapeOf(Dim(len(data)))
	}

	v := ViewOf(shape, m.Strides, m.Offset)

	l := v.DataSize()
	if data == nil {
		data = make([]T, l+v.o)
		if m.Default.Present() {
			d := m.Default.Value()
			for i := range data {
				data[i] = d
			}
		}
	}

	if Dim(len(data))-m.Offset < l {
		panic(fmt.Errorf("size mismatch"))
	}

	return NdArray[T]{
		v: v,
		d: data,
	}
}

func New[T any](shape Shape) NdArray[T] {
	return Maker[T]{Shape: shape}.Make()
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
		n := nsh.Get(i)
		for j := Dim(0); j < n; j++ {
			asl[axes.Get(i)] = j
			bsl[i] = j
			if i < axes.Len()-1 {
				rec(i + 1)
			} else {
				*b.At(bsl...) = a.Get(asl...)
			}
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

func (a NdArray[T]) MoveAxis(src, dst int) NdArray[T] {
	if src == dst {
		return a
	}

	o := a.v.sh.Order()
	axes := NewMutDims(o)
	var p Dim
	for i := 0; i < o; i++ {
		if i == src {
			if src < dst {
				p++
			}
			axes.Set(i, p)
			p++
		} else if i == dst {
			axes.Set(i, Dim(src))
		} else {
			axes.Set(i, p)
			p++
		}
	}

	return a.Transpose(axes.Decay())
}

func (a NdArray[T]) Reshape(nsh Shape) NdArray[T] {
	if dd, ok := nsh.Find(-1); ok {
		x := Dim(1)
		for i := 0; i < nsh.Len(); i++ {
			if i != dd {
				x *= nsh.Get(i)
			}
		}

		z := a.View().Shape().Prod()
		if z%x != 0 {
			panic(fmt.Errorf("reshape error: %d %% %d != 0", z, x))
		}

		mnsh := nsh.Mutate()
		mnsh.Set(dd, z/x)
		nsh = Shape{mnsh.Decay()}
	}

	l := a.v.sh.Prod()
	check.Equal(l, nsh.Prod())

	// FIXME: everything lol (contiguous)
	s := make([]T, l)

	o := a.v.sh.Order()
	sl := make([]Dim, a.v.sh.Order())
	p := 0

	var rec func(int)
	rec = func(i int) {
		n := a.v.sh.Get(i)
		for j := Dim(0); j < n; j++ {
			sl[i] = j
			if i < o-1 {
				rec(i + 1)
			} else {
				s[p] = a.Get(sl...)
				p++
			}
		}
	}
	rec(0)

	return Maker[T]{
		Shape: nsh,
		Data:  s,
	}.Make()
}

func (a NdArray[T]) Assign(src NdArray[T]) {
	check.Equal(a.v.sh.Order(), src.v.sh.Order())

	o := a.v.sh.Order()
	sl := make([]Dim, o)

	var rec func(int)
	rec = func(i int) {
		n := bt.Min(a.v.sh.Get(i), src.v.sh.Get(i))
		for j := Dim(0); j < n; j++ {
			sl[i] = j
			if i < o-1 {
				rec(i + 1)
			} else {
				*a.At(sl...) = src.Get(sl...)
			}
		}
	}
	rec(0)
}

func (a NdArray[T]) Flat() []T {
	if a.v.Linear() {
		return a.d[a.v.o:]
	}

	sh := a.v.sh
	o := sh.Order()
	s := make([]T, sh.Prod())
	p := 0
	sl := make([]Dim, o)
	var rec func(int)
	rec = func(i int) {
		l := sh.Get(i)
		for j := Dim(0); j < l; j++ {
			sl[i] = j
			if i < o-1 {
				rec(i + 1)
			} else {
				s[p] = a.Get(sl...)
				p++
			}
		}
	}
	rec(0)
	return s
}

//

func Map[F, T any](fn func(F) T, f NdArray[F]) NdArray[T] {
	t := New[T](f.v.sh)
	o := f.v.sh.Order()
	sl := make([]Dim, o)
	var rec func(int)
	rec = func(i int) {
		l := f.v.sh.Get(i)
		for j := Dim(0); j < l; j++ {
			sl[i] = j
			if i < o-1 {
				rec(i + 1)
			} else {
				*t.At(sl...) = fn(f.Get(sl...))
			}
		}
	}
	rec(0)
	return t
}
