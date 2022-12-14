package main

import "github.com/wrmsr/bane/exp/util/wasm/instrs"

type Prim int8

const (
	I32 Prim = iota + 1
	I64
	F32
	F64

	V128
)

type Def struct {
	I instrs.Instr

	Class instrs.Class
	Name  string
	OpPfx int8
	Op    int8

	T Prim
	A Prim
	B Prim
	C Prim

	Ma instrs.Access
	Mz int
}

//

type Instr = instrs.Instr

const (
	Control = instrs.Control
)

//
