package ndarray

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/def"
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
	return v.sh.Equals(o.sh) && v.st.Equals(o.st) && v.o == o.o
}

func (v View) String() string {
	var sb strings.Builder
	sb.WriteRune('{')
	sb.WriteString(v.sh.String())
	sb.WriteRune(' ')
	sb.WriteString(v.st.String())
	sb.WriteRune(' ')
	sb.WriteString(strconv.FormatInt(int64(v.o), 10))
	sb.WriteRune('}')
	return sb.String()
}

func (v View) Shape() Shape     { return v.sh }
func (v View) Strides() Strides { return v.st }
func (v View) Offset() Dim      { return v.o }

func (v View) DataSize() Dim { return v.sh.Get(0) * v.st.Get(0) }

var _ = def.Inline(
	View.Index,
)

func (v View) Index(idxs Dims) Dim {
	return v.st.Offset(idxs) + v.o
}

// FIXME: test code, remove

var _ = def.WithInline(foo)

func foo(v View) Dim {
	return v.Index(DimsOf(2, 2))
}

//

//func (v View) Contiguous() bool {
//	st := 0
//	for i := v.sh.Len()-1; i > 0; i-- {
//		if v.
//
//	}
//}

func (v View) Linear() bool {
	return v.st.Equals(CalcStrides(v.sh))
}

//

func (v View) Slice(bs ...any) View {
	nd := v.sh.Len()
	for _, b := range bs {
		if IsSqueezedRange(b) {
			nd--
		}
	}

	nsh := NewMutDims(nd)
	nst := NewMutDims(nd)
	no := v.o

	p := nd - 1
	for i := v.sh.Len() - 1; i >= 0; i-- {
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

		no += r.Start * (v.st.Get(i) * r.Step)

		if r.Squeeze {
			check.Condition(r.Scalar().Present())
		} else {
			nsh.Set(p, (r.Stop-r.Start+rnd)/r.Step)
			nst.Set(p, v.st.Get(i)*r.Step)
			p--
		}
	}

	return View{
		sh: Shape{nsh.Decay()},
		st: Strides{nst.Decay()},
		o:  no,
	}
}

//

func (v View) Squeeze() View {
	if v.sh.Len() < 2 {
		return v
	}

	nsd := v.sh.NumScalarDims()
	if nsd < 1 {
		return v
	}

	nd := v.sh.Len()
	nsh := NewMutDims(nd - nsd)
	nst := NewMutDims(nd - nsd)

	i := 0
	for j := 0; j < nd; j++ {
		if v.sh.Get(j) == 1 {
			continue
		}
		nst.Set(i, v.st.Get(j))
		nsh.Set(i, v.sh.Get(j))
		i++
	}

	return View{
		sh: Shape{nsh.Decay()},
		st: Strides{nst.Decay()},
		o:  v.o,
	}
}
