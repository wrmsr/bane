package ndarray

import "fmt"

type Strides []Dim

func calcStrides(shape Shape) Strides {
	st := make(Strides, len(shape))
	st[len(shape)-1] = 1
	for i := len(shape) - 2; i >= 0; i-- {
		st[i] = st[i+1] * shape[i+1]
	}
	return st
}

func (st Strides) Offset(idxs []Dim) Dim {
	if len(idxs) != len(st) {
		panic(fmt.Errorf("dim mismatch: got %d need %d", len(idxs), len(st)))
	}

	o := 0
	for i, idx := range idxs {
		o += i * st[idx]
	}

	return o
}
