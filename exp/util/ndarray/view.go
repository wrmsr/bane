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
	if strides.Len() < 1 {
		strides = CalcStrides(shape)
	}
	if strides.Len() != shape.Len() {
		panic(fmt.Errorf("strides mismatch"))
	}

	return View{
		sh: shape,
		st: strides,
		o:  offset,
	}
}

func (v View) Equals(o View) bool {
	if v.sh.Len() != o.sh.Len() || v.o != o.o {
		return false
	}
	for i := 0; i < v.sh.Len(); i++ {
		if v.sh.Get(i) != o.sh.Get(i) || v.st.Get(i) != o.st.Get(i) {
			return false
		}
	}
	return true
}

func (v View) Shape() Shape     { return v.sh }
func (v View) Strides() Strides { return v.st }
func (v View) Offset() Dim      { return v.o }

func (v View) Len() Dim { return v.sh.Get(0) * v.st.Get(0) }

func (v View) Index(idxs ...Dim) Dim {
	return v.st.Offset(idxs) + v.o
}

//

func (v View) Slice(bs ...any) View {
	nd := v.sh.Len()

	nsh := make([]Dim, nd)
	nst := make([]Dim, nd)
	no := v.o

	for i := nd - 1; i >= 0; i-- {
		var r Range
		if i < len(bs) {
			r = CalcRange(bs[i], v.sh.Get(i))
			check.Condition(r.Step > 0)
		} else {
			r = Range{
				Start: 0,
				Stop:  v.sh.Get(i),
				Step:  1,
			}
		}

		var rnd Dim
		if r.Step < 0 {
			rnd = r.Step + 1
		} else {
			rnd = r.Step - 1
		}

		nsh[i] = (r.Stop - r.Start + rnd) / r.Step
		nst[i] = v.st.Get(i) * r.Step
		no += r.Start * (v.st.Get(i) * r.Step)
	}

	return View{
		sh: ShapeOf(nsh...),
		st: StridesOf(nst...),
		o:  no,
	}
}

//

func (v View) Squeeze() View {
	nsd := v.sh.NumScalarDims()
	if nsd < 1 {
		return v
	}

	nd := v.sh.Len()
	nsh := make([]Dim, nd-nsd)
	nst := make([]Dim, nd-nsd)

	i := 0
	for j := 0; j < nd; j++ {
		if v.sh.Get(j) == 1 {
			continue
		}
		nst[i] = v.st.Get(j)
		nsh[i] = v.sh.Get(j)
		i++
	}

	return View{
		sh: ShapeOf(nsh...),
		st: StridesOf(nst...),
		o:  v.o,
	}
}
