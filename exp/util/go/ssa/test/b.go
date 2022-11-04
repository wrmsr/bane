package test

func B(x, y int) int {
	for i := 0; i < y; i++ {
		x += A(x, y) * 2
	}
	return x
}
