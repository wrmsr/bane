//go:build !nodev

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/wrmsr/bane/exp/util/blas/openblas"
)

func main() {
	m := 4096
	n := 4096
	k := 4096
	a := make([]float32, m*k)
	for i := range a {
		a[i] = float32(i)
	}
	b := make([]float32, k*n)
	for i := range b {
		b[i] = float32(i)
	}
	c := make([]float32, m*n)
	for {
		fmt.Println(os.Getpid())
		bufio.NewScanner(os.Stdin).Scan()
		openblas.Sgemm(
			openblas.Colmajor,
			openblas.Notrans,
			openblas.Notrans,
			m,
			n,
			k,
			1,
			a,
			m,
			b,
			k,
			1,
			c,
			m,
		)
		s := openblas.Sasum(len(c), c, 1)
		fmt.Println(s)
	}
}
