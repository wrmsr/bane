package openblas

import (
	"fmt"
	"testing"
)

func TestOpenblas(t *testing.T) {
	x := []float32{
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10,
	}
	r := Ssum(len(x), x[:], 1)
	fmt.Println(r)
}
