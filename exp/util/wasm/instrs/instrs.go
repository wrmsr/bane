package instrs

import wt "github.com/wrmsr/bane/exp/util/wasm/types"

type Instr struct {
	N string
	O Opcode

	T wt.Type
	A wt.Type
	B wt.Type

	Z int
}

func _add(i Instr) Instr {
	return i
}

var (
	Unreachable  = _add(Instr{N: "unreachable", O: Op_Unreachable})
	Nop          = _add(Instr{N: "nop", O: Op_Nop})
	Block        = _add(Instr{N: "block", O: Op_Block})
	Loop         = _add(Instr{N: "loop", O: Op_Loop})
	If           = _add(Instr{N: "if", O: Op_If})
	Else         = _add(Instr{N: "else", O: Op_Else})
	End          = _add(Instr{N: "end", O: Op_End})
	Br           = _add(Instr{N: "br", O: Op_Br})
	BrIf         = _add(Instr{N: "br_if", O: Op_BrIf, A: wt.I32{}})
	BrTable      = _add(Instr{N: "br_table", O: Op_BrTable, A: wt.I32{}})
	Return       = _add(Instr{N: "return", O: Op_Return})
	Call         = _add(Instr{N: "call", O: Op_Call})
	CallIndirect = _add(Instr{N: "call_indirect", O: Op_CallIndirect})
)
