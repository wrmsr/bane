package main

import (
	"fmt"

	"github.com/wrmsr/bane/exp/util/blas/openblas"
)

func main() {
	x := make([]float32, 1024*1024)
	for i := range x {
		x[i] = float32(i)
	}
	r := openblas.Ssum(len(x), x, 1)
	fmt.Println(r)
}
