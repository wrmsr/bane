package instrs

import wt "github.com/wrmsr/bane/exp/util/wasm/types"

type Instr struct {
	N string
	O Opcode

	T wt.Type
	A wt.Type
	B wt.Type
	C wt.Type

	M int
}

func _add(i Instr) Instr {
	return i
}

var (
	// control

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

	// reference

	// RefNull = _add(Instr{N: "ref_null", O: Op_RefNull})
	// RefIsNull = _add(Instr{N: "ref_is_null", O: Op_RefIsNull})
	// RefFunc = _add(Instr{N: "ref_func", O: Op_RefFunc})

	// parametric

	Drop       = _add(Instr{N: "drop", O: Op_Drop})
	Select     = _add(Instr{N: "select", O: Op_Select, C: wt.I32{}})
	SelectType = _add(Instr{N: "select_type", O: Op_SelectType, C: wt.I32{}})

	// variable

	LocalGet  = _add(Instr{N: "local_get", O: Op_LocalGet})
	LocalSet  = _add(Instr{N: "local_set", O: Op_LocalSet})
	LocalTee  = _add(Instr{N: "local_tee", O: Op_LocalTee})
	GlobalGet = _add(Instr{N: "global_get", O: Op_GlobalGet})
	GlobalSet = _add(Instr{N: "global_set", O: Op_GlobalSet})

	// table

	TableGet  = _add(Instr{N: "table_get", O: Op_TableGet, A: wt.I32{}})
	TableSet  = _add(Instr{N: "table_set", O: Op_TableSet, A: wt.I32{}})
	TableInit = _add(Instr{N: "table_init", O: Op_TableInit, A: wt.I32{}, B: wt.I32{}, C: wt.I32{}})
	ElemDrop  = _add(Instr{N: "elem_drop", O: Op_ElemDrop})
	TableCopy = _add(Instr{N: "table_copy", O: Op_TableCopy, A: wt.I32{}, B: wt.I32{}, C: wt.I32{}})
	TableGrow = _add(Instr{N: "table_grow", O: Op_TableGrow, B: wt.I32{}})
	TableSize = _add(Instr{N: "table_size", O: Op_TableSize})
	TableFill = _add(Instr{N: "table_fill", O: Op_TableFill, A: wt.I32{}, C: wt.I32{}})

	// memory

	Load_I32 = _add(Instr{N: "load_i32", O: Op_Load_I32, T: wt.I32{}, A: wt.I32{}, M: 4})
	Load_I64 = _add(Instr{N: "load_i64", O: Op_Load_I64, T: wt.I64{}, A: wt.I32{}, M: 8})
	// Load_F32 = _add(Instr{N: "load_f32", O: Op_Load_F32})
	// Load_F64 = _add(Instr{N: "load_f64", O: Op_Load_F64})
	// Load8S_I32 = _add(Instr{N: "load8s_i32", O: Op_Load8S_I32})
	// Load8U_I32 = _add(Instr{N: "load8u_i32", O: Op_Load8U_I32})
	// Load16S_I32 = _add(Instr{N: "load16s_i32", O: Op_Load16S_I32})
	// Load16U_I32 = _add(Instr{N: "load16u_i32", O: Op_Load16U_I32})
	// Load8S_I64 = _add(Instr{N: "load8s_i64", O: Op_Load8S_I64})
	// Load8U_I64 = _add(Instr{N: "load8u_i64", O: Op_Load8U_I64})
	// Load16S_I64 = _add(Instr{N: "load16s_i64", O: Op_Load16S_I64})
	// Load16U_I64 = _add(Instr{N: "load16u_i64", O: Op_Load16U_I64})
	// Load32S_I64 = _add(Instr{N: "load32s_i64", O: Op_Load32S_I64})
	// Load32U_I64 = _add(Instr{N: "load32u_i64", O: Op_Load32U_I64})
	// Store_I32 = _add(Instr{N: "store_i32", O: Op_Store_I32})
	// Store_I64 = _add(Instr{N: "store_i64", O: Op_Store_I64})
	// Store_F32 = _add(Instr{N: "store_f32", O: Op_Store_F32})
	// Store_F64 = _add(Instr{N: "store_f64", O: Op_Store_F64})
	// Store8_I32 = _add(Instr{N: "store8_i32", O: Op_Store8_I32})
	// Store16_I32 = _add(Instr{N: "store16_i32", O: Op_Store16_I32})
	// Store8_I64 = _add(Instr{N: "store8_i64", O: Op_Store8_I64})
	// Store16_I64 = _add(Instr{N: "store16_i64", O: Op_Store16_I64})
	// Store32_I64 = _add(Instr{N: "store32_i64", O: Op_Store32_I64})
	// MemorySize = _add(Instr{N: "memory_size", O: Op_MemorySize})
	// MemoryGrow = _add(Instr{N: "memory_grow", O: Op_MemoryGrow})
	// MemoryInit = _add(Instr{N: "memory_init", O: Op_MemoryInit})
	// DataDrop = _add(Instr{N: "data_drop", O: Op_DataDrop})
	// MemoryCopy = _add(Instr{N: "memory_copy", O: Op_MemoryCopy})
	// MemoryFill = _add(Instr{N: "memory_fill", O: Op_MemoryFill})

)
