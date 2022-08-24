package ndarray

import "fmt"

type Strides struct {
	Dims
}

func StridesOf(s ...Dim) Strides {
	return Strides{DimsOf(s...)}
}

func CalcStrides(shape Shape) Strides {
	st := make([]Dim, shape.Len())
	st[shape.Len()-1] = 1
	for i := shape.Len() - 2; i >= 0; i-- {
		st[i] = st[i+1] * shape.Get(i+1)
	}
	return Strides{Dims{st}}
}

func (st Strides) Offset(idxs []Dim) Dim {
	if len(idxs) != st.Len() {
		panic(fmt.Errorf("dim mismatch: got %d need %d", len(idxs), st.Len()))
	}
	o := Dim(1)
	for i, idx := range idxs {
		o += idx * st.Get(i)
	}
	return o
}
