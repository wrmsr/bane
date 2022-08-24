package ndarray

type Shape []Dim

func (sh Shape) NumScalarDims() int {
	i := 0
	for _, d := range sh {
		if d == 1 {
			i++
		}
	}
	return i
}
