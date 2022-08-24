package ndarray

import (
	"fmt"

	"github.com/wrmsr/bane/pkg/util/check"
)

//

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
	if len(strides.s) < 1 {
		strides = CalcStrides(shape)
	}
	if len(strides.s) != len(shape.s) {
		panic(fmt.Errorf("strides mismatch"))
	}

	return View{
		sh: shape,
		st: strides,
		o:  offset,
	}
}

func (v View) Equals(o View) bool {
	if len(v.sh.s) != len(o.sh.s) || v.o != o.o {
		return false
	}
	for i := range v.sh.s {
		if v.sh.s[i] != o.sh.s[i] || v.st.s[i] != o.st.s[i] {
			return false
		}
	}
	return true
}

func (v View) Shape() Shape     { return v.sh }
func (v View) Strides() Strides { return v.st }
func (v View) Offset() Dim      { return v.o }

func (v View) Len() Dim { return v.sh.s[0] * v.st.s[0] }

func (v View) Index(idxs ...Dim) Dim {
	return v.st.Offset(idxs) + v.o
}

//

func (v View) Slice(bs ...any) View {
	nd := len(v.sh.s)
	if len(bs) != nd {
		panic(fmt.Errorf("slice dimension mismatch"))
	}

	nsh := Shape{s: make([]Dim, nd)}
	nst := Strides{s: make([]Dim, nd)}
	no := v.o

	for i := nd - 1; i >= 0; i-- {
		r := CalcRange(bs[i], v.sh.s[i])
		check.Condition(r.Step > 0)

		var rnd Dim
		if r.Step < 0 {
			rnd = r.Step + 1
		} else {
			rnd = r.Step - 1
		}

		nsh.s[i] = (r.Stop - r.Start + rnd) / r.Step
		nst.s[i] = v.st.s[i] * r.Step
		no += r.Start * (v.st.s[i] * r.Step)
	}

	return View{
		sh: nsh,
		st: nst,
		o:  no,
	}
}

//

func (v View) Squeeze() View {
	nsd := v.sh.NumScalarDims()
	if nsd < 1 {
		return v
	}

	nd := len(v.sh.s)
	nsh := Shape{s: make([]Dim, nd-nsd)}
	nst := Strides{s: make([]Dim, nd-nsd)}

	i := 0
	for j := 0; j < nd; j++ {
		if v.sh.s[j] == 1 {
			continue
		}
		nst.s[i] = v.st.s[j]
		nsh.s[i] = v.sh.s[j]
		i++
	}

	return View{
		sh: nsh,
		st: nst,
		o:  v.o,
	}
}
