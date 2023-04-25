package tensor

import (
	"testing"

	tu "github.com/wrmsr/bane/core/dev/testing"
)

func TestDumbMatmul(t *testing.T) {
	a := [][]float32{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}

	b := [][]float32{
		{10, 11},
		{13, 14},
		{16, 17},
	}

	c := [][]float32{
		{0, 0},
		{0, 0},
		{0, 0},
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			var o float32
			for k := 0; k < 3; k++ {
				o += a[i][k] * b[k][j]
			}
			c[i][j] = o
		}
	}

	tu.AssertDeepEqual(t, c, [][]float32{
		{45, 48},
		{162, 174},
		{279, 300},
	})
}
