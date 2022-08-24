package ndarray

type Shape struct {
	s []Dim
}

func ShapeOf(s ...Dim) Shape {
	return Shape{s: s}
}

func (sh Shape) Len() int      { return len(sh.s) }
func (sh Shape) Get(i int) Dim { return sh.s[i] }

func (sh Shape) NumScalarDims() int {
	i := 0
	for _, d := range sh.s {
		if d == 1 {
			i++
		}
	}
	return i
}
