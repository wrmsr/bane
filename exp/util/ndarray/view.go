package ndarray

import (
	"fmt"

	"github.com/wrmsr/bane/pkg/util/check"
)

type View struct {
	sh Shape
	st Strides
	o  Dim
}

func ViewOf(
	shape Shape,
	strides Strides,
	offset Dim,
) View {
	if len(strides) < 1 {
		strides = calcStrides(shape)
	}
	if len(strides) != len(shape) {
		panic(fmt.Errorf("strides mismatch"))
	}

	return View{
		sh: shape,
		st: strides,
		o:  offset,
	}
}

func (v View) Equals(o View) bool {
	if len(v.sh) != len(o.sh) || v.o != o.o {
		return false
	}
	for i := range v.sh {
		if v.sh[i] != o.sh[i] || v.st[i] != o.st[i] {
			return false
		}
	}
	return true
}

func (v View) Shape() Shape     { return v.sh }
func (v View) Strides() Strides { return v.st }
func (v View) Offset() Dim      { return v.o }

func (v View) Index(idxs ...Dim) Dim {
	return v.st.Offset(idxs) + v.o
}

func (v View) Slice(bs ...any) View {
	nd := len(v.sh)
	if len(bs) != nd {
		panic(fmt.Errorf("slice dimension mismatch"))
	}

	nsh := make(Shape, nd)
	nst := make(Strides, nd)
	no := v.o

	for i := nd - 1; i >= 0; i-- {
		r := CalcRange(bs[i], v.sh[i])
		check.Condition(r.Step > 0)

		rnd := 0
		if r.Step < 0 {
			rnd = r.Step + 1
		} else {
			rnd = r.Step - 1
		}

		nsh[i] = (r.Stop - r.Start + rnd) / r.Step
		nst[i] = v.st[i] * r.Step
		no += r.Start * (v.st[i] * r.Step)
	}

	return View{
		sh: nsh,
		st: nst,
		o:  no,
	}
}

func (v View) Squeeze() View {
	nsd := v.sh.NumScalarDims()
	if nsd < 1 {
		return v
	}

	nd := len(v.sh)
	nsh := make(Shape, nd-nsd)
	nst := make(Strides, nd-nsd)

	i := 0
	for j := 0; j < nd; j++ {
		if v.sh[j] == 1 {
			continue
		}
		nst[i] = v.st[j]
		nsh[i] = v.sh[j]
		i++
	}

	return View{
		sh: nsh,
		st: nst,
		o:  v.o,
	}
}
