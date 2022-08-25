package ndarray

type Shape struct {
	Dims
}

func ShapeOf(s ...Dim) Shape {
	return Shape{DimsOf(s...)}
}

func (sh Shape) Equals(o Shape) bool {
	return sh.Dims.Equals(o.Dims)
}

func (sh Shape) NumScalarDims() int {
	i := 0
	for n := 0; n < sh.Len(); n++ {
		d := sh.Get(n)
		if d == 1 {
			i++
		}
	}
	return i
}
