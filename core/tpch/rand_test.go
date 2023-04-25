package tpch

import (
	"fmt"
	"testing"
)

func TestRandomAlphaNumeric(t *testing.T) {
	r := newRandomAlphaNumeric(420, 25, 100)
	for i := 0; i < 100; i++ {
		for j := 0; j < 8; j++ {
			fmt.Println(r.nextValue())
		}
		r.rowFinished()
		fmt.Println()
	}
}
