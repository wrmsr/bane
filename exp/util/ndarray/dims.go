package ndarray

import (
	"fmt"
	"strconv"
	"strings"
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
		for i := 0; i < ds._l-_dimsWidth; i++ {
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

		sb.WriteString(strconv.FormatInt(int64(d), 10))
	}

	sb.WriteRune(']')
	return sb.String()
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

func (ds *MutDims) Decay() Dims { return ds._ds }

func (ds *MutDims) Len() int      { return ds._ds.Len() }
func (ds *MutDims) Get(i int) Dim { return ds._ds.Get(i) }

func (ds *MutDims) Set(i int, d Dim) {
	if i < 0 || i >= ds._ds.Len() {
		panic(fmt.Errorf("index out of bounds: %d", i))
	}

	if i < _dimsWidth {
		ds._ds._a[i] = d
	} else {
		ds._ds._s[i-_dimsWidth] = d
	}
}
