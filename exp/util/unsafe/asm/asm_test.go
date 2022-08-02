package asm

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
