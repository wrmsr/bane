package ndarray

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/def"
)

//

const _dimsWidth = 6

type Dims struct {
	_a [_dimsWidth]Dim
	_s []Dim
	_l int
}

func DimsOf(s ...Dim) Dims {
	ds := Dims{
		_l: len(s),
	}

	n := _dimsWidth
	if n > len(s) {
		n = len(s)
	}

	copy(ds._a[:], s[:n])

	o := len(s) - _dimsWidth
	if o > 0 {
		ds._s = make([]Dim, o)
		copy(ds._s, s[_dimsWidth:])
	}

	return ds
}

var _ = def.Inline(
	Dims.CheckEqualLen,
	Dims.Len,
	Dims.Get,
)

func (ds Dims) Equals(o Dims) bool {
	if ds._l != o._l {
		return false
	}

	n := _dimsWidth
	if n > ds._l {
		n = ds._l
	}

	for i := 0; i < n; i++ {
		if ds._a[i] != o._a[i] {
			return false
		}
	}

	if ds._l > _dimsWidth {
		for i := 0; i < len(ds._s); i++ {
			if ds._s[i] != o._s[i] {
				return false
			}
		}

	}

	return true
}

func (ds Dims) String() string {
	var sb strings.Builder
	sb.WriteRune('[')

	for i := 0; i < ds._l; i++ {
		if i > 0 {
			sb.WriteRune(' ')
		}

		var d Dim
		if i < _dimsWidth {
			d = ds._a[i]
		} else {
			d = ds._s[i-_dimsWidth]
		}

		sb.WriteString(strconv.FormatInt(d, 10))
	}

	sb.WriteRune(']')
	return sb.String()
}

func (ds Dims) CheckEqualLen(o Dims) {
	if ds._l != o._l {
		panic(fmt.Errorf("dim mismatch: got %d need %d", ds._l, o._l))
	}
}

func (ds Dims) Slice() []Dim {
	s := make([]Dim, ds._l)
	copy(s[:ds._l], ds._a[:ds._l])
	if ds._l > _dimsWidth {
		copy(s[_dimsWidth:], ds._s)
	}
	return s
}

func (ds Dims) Len() int {
	return ds._l
}

func (ds Dims) Get(i int) Dim {
	if i < 0 || i >= ds._l {
		panic(fmt.Errorf("index out of bounds: %d", i))
	}

	if i < _dimsWidth {
		return ds._a[i]
	} else {
		return ds._s[i-_dimsWidth]
	}
}

func (ds *Dims) _set(i int, d Dim) {
	if i < _dimsWidth {
		ds._a[i] = d
	} else {
		ds._s[i-_dimsWidth] = d
	}
}

func (ds Dims) Find(d Dim) (int, bool) {
	for i := 0; i < ds._l; i++ {
		if ds.Get(i) == d {
			return i, true
		}
	}
	return -1, false
}

func (ds Dims) Contains(d Dim) bool {
	_, ok := ds.Find(d)
	return ok
}

func (ds Dims) Sum() Dim {
	var r Dim
	for i := 0; i < ds._l; i++ {
		r += ds.Get(i)
	}
	return r
}

func (ds Dims) Prod() Dim {
	var r Dim = 1
	for i := 0; i < ds._l; i++ {
		r *= ds.Get(i)
	}
	return r
}

//

type MutDims struct {
	_ds Dims
}

func NewMutDims(l int) MutDims {
	ds := Dims{
		_l: l,
	}

	if l > _dimsWidth {
		ds._s = make([]Dim, l-_dimsWidth)
	}

	return MutDims{ds}
}

func MutDimsTo(l int) MutDims {
	m := NewMutDims(l)
	for i := 0; i < l; i++ {
		m.Set(i, Dim(i))
	}
	return m
}

func DimsTo(l int) Dims {
	md := MutDimsTo(l)
	return md.Decay()
}

func (ds Dims) Mutate() MutDims {
	var s []Dim
	if ds._s != nil {
		s = make([]Dim, len(ds._s))
		copy(s, ds._s)
	}

	return MutDims{
		_ds: Dims{
			_a: ds._a,
			_s: s,
			_l: ds._l,
		},
	}
}

func (ds *MutDims) Decay() Dims { return ds._ds }

func (ds *MutDims) Len() int      { return ds._ds.Len() }
func (ds *MutDims) Get(i int) Dim { return ds._ds.Get(i) }

func (ds *MutDims) Set(i int, d Dim) {
	if i < 0 || i >= ds._ds.Len() {
		panic(fmt.Errorf("index out of bounds: %d", i))
	}

	ds._ds._set(i, d)
}

//

func AsDims(s ...any) Dims {
	m := NewMutDims(len(s))
	for i, d := range s {
		m.Set(i, check.Ok1(AsDim(d)).Value())
	}
	return m.Decay()
}
