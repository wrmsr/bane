package ndarray

func Unnest2[T any](n [][]T) NdArray[T] {
	l0, l1 := len(n), len(n[0])
	s := make([]T, l0*l1)
	for i0, n0 := range n {
		for i1, v := range n0 {
			s[(i0*l1)+i1] = v
		}
	}

	return Maker[T]{
		Shape: ShapeOf(Dim(l0), Dim(l1)),
		Data:  s,
	}.Make()
}

func Unnest3[T any](n [][][]T) NdArray[T] {
	l0, l1, l2 := len(n), len(n[0]), len(n[0][0])
	s := make([]T, l0*l1*l2)
	for i0, n0 := range n {
		for i1, n1 := range n0 {
			for i2, v := range n1 {
				s[(i0*l1*l2)+(i1*l2)+i2] = v
			}
		}
	}

	return Maker[T]{
		Shape: ShapeOf(Dim(l0), Dim(l1), Dim(l2)),
		Data:  s,
	}.Make()
}
