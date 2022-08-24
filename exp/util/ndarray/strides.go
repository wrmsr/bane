package ndarray

import "fmt"

type Strides struct {
	s []Dim
}

func CalcStrides(shape Shape) Strides {
	st := Strides{s: make([]Dim, len(shape.s))}
	st.s[len(shape.s)-1] = 1
	for i := len(shape.s) - 2; i >= 0; i-- {
		st.s[i] = st.s[i+1] * shape.s[i+1]
	}
	return st
}

func (st Strides) Offset(idxs []Dim) Dim {
	if len(idxs) != len(st.s) {
		panic(fmt.Errorf("dim mismatch: got %d need %d", len(idxs), len(st.s)))
	}
	o := 0
	for i, idx := range idxs {
		o += idx * st.s[i]
	}
	return o
}
