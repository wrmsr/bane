package asm

import (
	"fmt"
	"math"
	"testing"
)

type Op int16

/*
100003f54: 80 34 80 52  mov     w0, #420
100003f58: c0 03 5f d6  ret
*/

const (
	OpInvalid Op = iota
	OpAdd
	OpMov
	OpRet
)

type Inst struct {
	Op Op
}

//go:noinline
func someFloat() float64 {
	return 420.69
}

func TestFloatBits(t *testing.T) {
	f := someFloat()
	b := math.Float64bits(f)
	fmt.Println(b)
}
