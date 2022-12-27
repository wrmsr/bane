package instrs

import wt "github.com/wrmsr/bane/exp/util/wasm/types"

type Instr int16

type Def struct {
	I Instr

	Class Class
	Name  string
	Name2 string
	OpPfx uint8
	Op    uint8

	T wt.Prim
	A wt.Prim
	B wt.Prim
	C wt.Prim

	Ma Access
	Mz int
}
