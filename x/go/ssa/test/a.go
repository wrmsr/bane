package test

func A(x, y int) int {
	if x > 0 {
		x += 1
	}
	return -x + y + 1
}
