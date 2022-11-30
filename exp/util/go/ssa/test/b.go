package test

import (
	"math"
)

func B(x, y int) int {
	math.Jn(x, 0)
	for i := 0; i < y; i++ {
		x += A(x, y) * 2
	}
	return x
}
