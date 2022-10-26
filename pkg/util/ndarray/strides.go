package ndarray

import "github.com/wrmsr/bane/pkg/util/def"

type Strides struct {
	Dims
}

func StridesOf(s ...Dim) Strides {
	return Strides{DimsOf(s...)}
}

func (st Strides) Equals(o Strides) bool {
	return st.Dims.Equals(o.Dims)
}

func CalcStrides(shape Shape) Strides {
	st := NewMutDims(shape.Len())
	st.Set(shape.Len()-1, 1)

	for i := shape.Len() - 2; i >= 0; i-- {
		st.Set(i, st.Get(i+1)*shape.Get(i+1))
	}

	return Strides{st.Decay()}
}

func (st Strides) Offset(idxs Dims) Dim {
	st.CheckEqualLen(idxs)

	return st._offset(idxs)
}

var _ = def.Inline(Strides._offset)

func (st Strides) _offset(idxs Dims) Dim {
	l := st.Len()

	o := Dim(0)
	for i := 0; i < _dimsWidth; i++ {
		o += idxs._a[i] * st._a[i]
	}
	for i := 0; i < l-_dimsWidth; i++ {
		o += idxs._s[i] * st._s[i]
	}

	return o
}
