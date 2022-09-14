package ndarray

import "fmt"

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
	l := st.Len()
	if idxs.Len() != l {
		panic(fmt.Errorf("dim mismatch: got %d need %d", idxs.Len(), l))
	}

	o := Dim(0)
	for i := 0; i < _dimsWidth; i++ {
		o += idxs._a[i] * st._a[i]
	}
	for i := _dimsWidth; i < l; i++ {
		o += idxs._s[i] * st._s[i]
	}

	return o
}
