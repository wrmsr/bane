package ndarray

type Shape struct {
	s []Dim
}

func (sh Shape) NumScalarDims() int {
	i := 0
	for _, d := range sh.s {
		if d == 1 {
			i++
		}
	}
	return i
}
