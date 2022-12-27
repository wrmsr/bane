//go:build !nodev

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
	Name2 string
	OpPfx uint8
	Op    uint8

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
	Control    = instrs.Control
	Memory     = instrs.Memory
	Numeric    = instrs.Numeric
	Parametric = instrs.Parametric
	Reference  = instrs.Reference
	Table      = instrs.Table
	Variable   = instrs.Variable
	Vector     = instrs.Vector
)

const (
	Load  = instrs.Load
	Store = instrs.Store
)

//
