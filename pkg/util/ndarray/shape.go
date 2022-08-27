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

func (sh Shape) Order() int { return sh.Len() }

func (sh Shape) Size() Dim {
	sz := Dim(1)
	for i := 0; i < sh._l; i++ {
		sz *= sh.Get(i)
	}
	return sz
}

func (sh Shape) IsScalar() bool {
	return sh._l == 1 && sh._a[0] == 1
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
