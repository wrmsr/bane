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

var instrsByOpcode = make(map[Opcode]Instr, 512)

func _add(i Instr) Instr {
	if _, ok := instrsByOpcode[i.O]; ok {
		panic(i.O)
	}
	instrsByOpcode[i.O] = i
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

	Load_I32    = _add(Instr{N: "load_i32", O: Op_Load_I32, T: wt.I32{}, A: wt.I32{}, M: 4})
	Load_I64    = _add(Instr{N: "load_i64", O: Op_Load_I64, T: wt.I64{}, A: wt.I32{}, M: 8})
	Load_F32    = _add(Instr{N: "load_f32", O: Op_Load_F32, T: wt.F32{}, A: wt.I32{}, M: 4})
	Load_F64    = _add(Instr{N: "load_f64", O: Op_Load_F64, T: wt.F64{}, A: wt.I32{}, M: 8})
	Load8S_I32  = _add(Instr{N: "load8s_i32", O: Op_Load8S_I32, T: wt.I32{}, A: wt.I32{}, M: 1})
	Load8U_I32  = _add(Instr{N: "load8u_i32", O: Op_Load8U_I32, T: wt.I32{}, A: wt.I32{}, M: 1})
	Load16S_I32 = _add(Instr{N: "load16s_i32", O: Op_Load16S_I32, T: wt.I32{}, A: wt.I32{}, M: 2})
	Load16U_I32 = _add(Instr{N: "load16u_i32", O: Op_Load16U_I32, T: wt.I32{}, A: wt.I32{}, M: 2})
	Load8S_I64  = _add(Instr{N: "load8s_i64", O: Op_Load8S_I64, T: wt.I64{}, A: wt.I32{}, M: 1})
	Load8U_I64  = _add(Instr{N: "load8u_i64", O: Op_Load8U_I64, T: wt.I64{}, A: wt.I32{}, M: 1})
	Load16S_I64 = _add(Instr{N: "load16s_i64", O: Op_Load16S_I64, T: wt.I64{}, A: wt.I32{}, M: 2})
	Load16U_I64 = _add(Instr{N: "load16u_i64", O: Op_Load16U_I64, T: wt.I64{}, A: wt.I32{}, M: 2})
	Load32S_I64 = _add(Instr{N: "load32s_i64", O: Op_Load32S_I64, T: wt.I64{}, A: wt.I32{}, M: 4})
	Load32U_I64 = _add(Instr{N: "load32u_i64", O: Op_Load32U_I64, T: wt.I64{}, A: wt.I32{}, M: 4})
	Store_I32   = _add(Instr{N: "store_i32", O: Op_Store_I32, A: wt.I32{}, B: wt.I32{}, M: 4})
	Store_I64   = _add(Instr{N: "store_i64", O: Op_Store_I64, A: wt.I32{}, B: wt.I64{}, M: 8})
	Store_F32   = _add(Instr{N: "store_f32", O: Op_Store_F32, A: wt.I32{}, B: wt.F32{}, M: 4})
	Store_F64   = _add(Instr{N: "store_f64", O: Op_Store_F64, A: wt.I32{}, B: wt.F64{}, M: 8})
	Store8_I32  = _add(Instr{N: "store8_i32", O: Op_Store8_I32, A: wt.I32{}, B: wt.I32{}, M: 1})
	Store16_I32 = _add(Instr{N: "store16_i32", O: Op_Store16_I32, A: wt.I32{}, B: wt.I32{}, M: 2})
	Store8_I64  = _add(Instr{N: "store8_i64", O: Op_Store8_I64, A: wt.I32{}, B: wt.I64{}, M: 1})
	Store16_I64 = _add(Instr{N: "store16_i64", O: Op_Store16_I64, A: wt.I32{}, B: wt.I64{}, M: 2})
	Store32_I64 = _add(Instr{N: "store32_i64", O: Op_Store32_I64, A: wt.I32{}, B: wt.I64{}, M: 4})
	MemorySize  = _add(Instr{N: "memory_size", O: Op_MemorySize, T: wt.I32{}})
	MemoryGrow  = _add(Instr{N: "memory_grow", O: Op_MemoryGrow, T: wt.I32{}, A: wt.I32{}})
	MemoryInit  = _add(Instr{N: "memory_init", O: Op_MemoryInit, A: wt.I32{}, B: wt.I32{}, C: wt.I32{}})
	DataDrop    = _add(Instr{N: "data_drop", O: Op_DataDrop})
	MemoryCopy  = _add(Instr{N: "memory_copy", O: Op_MemoryCopy, A: wt.I32{}, B: wt.I32{}, C: wt.I32{}})
	MemoryFill  = _add(Instr{N: "memory_fill", O: Op_MemoryFill, A: wt.I32{}, B: wt.I32{}, C: wt.I32{}})

	// numeric

	Const_I32 = _add(Instr{N: "const_i32", O: Op_Const_I32, T: wt.I32{}})
	Const_I64 = _add(Instr{N: "const_i64", O: Op_Const_I64, T: wt.F32{}})
	Const_F32 = _add(Instr{N: "const_f32", O: Op_Const_F32, T: wt.I64{}})
	Const_F64 = _add(Instr{N: "const_f64", O: Op_Const_F64, T: wt.F64{}})

	Eqz_I32 = _add(Instr{N: "eqz_i32", O: Op_Eqz_I32, T: wt.I32{}, A: wt.I32{}})
	Eq_I32  = _add(Instr{N: "eq_i32", O: Op_Eq_I32, T: wt.I32{}, A: wt.I32{}, B: wt.I32{}})
	Ne_I32  = _add(Instr{N: "ne_i32", O: Op_Ne_I32, T: wt.I32{}, A: wt.I32{}, B: wt.I32{}})
	LtS_I32 = _add(Instr{N: "lts_i32", O: Op_LtS_I32, T: wt.I32{}, A: wt.I32{}, B: wt.I32{}})
	LtU_I32 = _add(Instr{N: "ltu_i32", O: Op_LtU_I32, T: wt.I32{}, A: wt.I32{}, B: wt.I32{}})
	GtS_I32 = _add(Instr{N: "gts_i32", O: Op_GtS_I32, T: wt.I32{}, A: wt.I32{}, B: wt.I32{}})
	GtU_I32 = _add(Instr{N: "gtu_i32", O: Op_GtU_I32, T: wt.I32{}, A: wt.I32{}, B: wt.I32{}})
	LeS_I32 = _add(Instr{N: "les_i32", O: Op_LeS_I32, T: wt.I32{}, A: wt.I32{}, B: wt.I32{}})
	LeU_I32 = _add(Instr{N: "leu_i32", O: Op_LeU_I32, T: wt.I32{}, A: wt.I32{}, B: wt.I32{}})
	GeS_I32 = _add(Instr{N: "ges_i32", O: Op_GeS_I32, T: wt.I32{}, A: wt.I32{}, B: wt.I32{}})
	GeU_I32 = _add(Instr{N: "geu_i32", O: Op_GeU_I32, T: wt.I32{}, A: wt.I32{}, B: wt.I32{}})

	Eqz_I64 = _add(Instr{N: "eqz_i64", O: Op_Eqz_I64, T: wt.I32{}, A: wt.I64{}})
	Eq_I64  = _add(Instr{N: "eq_i64", O: Op_Eq_I64, T: wt.I32{}, A: wt.I64{}, B: wt.I64{}})
	Ne_I64  = _add(Instr{N: "ne_i64", O: Op_Ne_I64, T: wt.I32{}, A: wt.I64{}, B: wt.I64{}})
	LtS_I64 = _add(Instr{N: "lts_i64", O: Op_LtS_I64, T: wt.I32{}, A: wt.I64{}, B: wt.I64{}})
	LtU_I64 = _add(Instr{N: "ltu_i64", O: Op_LtU_I64, T: wt.I32{}, A: wt.I64{}, B: wt.I64{}})
	GtS_I64 = _add(Instr{N: "gts_i64", O: Op_GtS_I64, T: wt.I32{}, A: wt.I64{}, B: wt.I64{}})
	GtU_I64 = _add(Instr{N: "gtu_i64", O: Op_GtU_I64, T: wt.I32{}, A: wt.I64{}, B: wt.I64{}})
	LeS_I64 = _add(Instr{N: "les_i64", O: Op_LeS_I64, T: wt.I32{}, A: wt.I64{}, B: wt.I64{}})
	LeU_I64 = _add(Instr{N: "leu_i64", O: Op_LeU_I64, T: wt.I32{}, A: wt.I64{}, B: wt.I64{}})
	GeS_I64 = _add(Instr{N: "ges_i64", O: Op_GeS_I64, T: wt.I32{}, A: wt.I64{}, B: wt.I64{}})
	GeU_I64 = _add(Instr{N: "geu_i64", O: Op_GeU_I64, T: wt.I32{}, A: wt.I64{}, B: wt.I64{}})

	Eq_F32 = _add(Instr{N: "eq_f32", O: Op_Eq_F32, T: wt.I32{}, A: wt.F32{}})
	Ne_F32 = _add(Instr{N: "ne_f32", O: Op_Ne_F32, T: wt.I32{}, A: wt.F32{}, B: wt.F32{}})
	Lt_F32 = _add(Instr{N: "lt_f32", O: Op_Lt_F32, T: wt.I32{}, A: wt.F32{}, B: wt.F32{}})
	Gt_F32 = _add(Instr{N: "gt_f32", O: Op_Gt_F32, T: wt.I32{}, A: wt.F32{}, B: wt.F32{}})
	Le_F32 = _add(Instr{N: "le_f32", O: Op_Le_F32, T: wt.I32{}, A: wt.F32{}, B: wt.F32{}})
	Ge_F32 = _add(Instr{N: "ge_f32", O: Op_Ge_F32, T: wt.I32{}, A: wt.F32{}, B: wt.F32{}})

	Eq_F64 = _add(Instr{N: "eq_f64", O: Op_Eq_F64, T: wt.I64{}, A: wt.F64{}})
	Ne_F64 = _add(Instr{N: "ne_f64", O: Op_Ne_F64, T: wt.I64{}, A: wt.F64{}, B: wt.F64{}})
	Lt_F64 = _add(Instr{N: "lt_f64", O: Op_Lt_F64, T: wt.I64{}, A: wt.F64{}, B: wt.F64{}})
	Gt_F64 = _add(Instr{N: "gt_f64", O: Op_Gt_F64, T: wt.I64{}, A: wt.F64{}, B: wt.F64{}})
	Le_F64 = _add(Instr{N: "le_f64", O: Op_Le_F64, T: wt.I64{}, A: wt.F64{}, B: wt.F64{}})
	Ge_F64 = _add(Instr{N: "ge_f64", O: Op_Ge_F64, T: wt.I64{}, A: wt.F64{}, B: wt.F64{}})

	// Clz_I32 = _add(Instr{N: "clz_i32", O: Op_Clz_I32})
	// Ctz_I32 = _add(Instr{N: "ctz_i32", O: Op_Ctz_I32})
	// Popcnt_I32 = _add(Instr{N: "popcnt_i32", O: Op_Popcnt_I32})
	// Add_I32 = _add(Instr{N: "add_i32", O: Op_Add_I32})
	// Sub_I32 = _add(Instr{N: "sub_i32", O: Op_Sub_I32})
	// Mul_I32 = _add(Instr{N: "mul_i32", O: Op_Mul_I32})
	// DivS_I32 = _add(Instr{N: "divs_i32", O: Op_DivS_I32})
	// DivU_I32 = _add(Instr{N: "divu_i32", O: Op_DivU_I32})
	// RemS_I32 = _add(Instr{N: "rems_i32", O: Op_RemS_I32})
	// RemU_I32 = _add(Instr{N: "remu_i32", O: Op_RemU_I32})
	// And_I32 = _add(Instr{N: "and_i32", O: Op_And_I32})
	// Or_I32 = _add(Instr{N: "or_i32", O: Op_Or_I32})
	// Xor_I32 = _add(Instr{N: "xor_i32", O: Op_Xor_I32})
	// Shl_I32 = _add(Instr{N: "shl_i32", O: Op_Shl_I32})
	// ShrS_I32 = _add(Instr{N: "shrs_i32", O: Op_ShrS_I32})
	// ShrU_I32 = _add(Instr{N: "shru_i32", O: Op_ShrU_I32})
	// Rotl_I32 = _add(Instr{N: "rotl_i32", O: Op_Rotl_I32})
	// Rotr_I32 = _add(Instr{N: "rotr_i32", O: Op_Rotr_I32})
	// Clz_I64 = _add(Instr{N: "clz_i64", O: Op_Clz_I64})
	// Ctz_I64 = _add(Instr{N: "ctz_i64", O: Op_Ctz_I64})
	// Popcnt_I64 = _add(Instr{N: "popcnt_i64", O: Op_Popcnt_I64})
	// Add_I64 = _add(Instr{N: "add_i64", O: Op_Add_I64})
	// Sub_I64 = _add(Instr{N: "sub_i64", O: Op_Sub_I64})
	// Mul_I64 = _add(Instr{N: "mul_i64", O: Op_Mul_I64})
	// DivS_I64 = _add(Instr{N: "divs_i64", O: Op_DivS_I64})
	// DivU_I64 = _add(Instr{N: "divu_i64", O: Op_DivU_I64})
	// RemS_I64 = _add(Instr{N: "rems_i64", O: Op_RemS_I64})
	// RemU_I64 = _add(Instr{N: "remu_i64", O: Op_RemU_I64})
	// And_I64 = _add(Instr{N: "and_i64", O: Op_And_I64})
	// Or_I64 = _add(Instr{N: "or_i64", O: Op_Or_I64})
	// Xor_I64 = _add(Instr{N: "xor_i64", O: Op_Xor_I64})
	// Shl_I64 = _add(Instr{N: "shl_i64", O: Op_Shl_I64})
	// ShrS_I64 = _add(Instr{N: "shrs_i64", O: Op_ShrS_I64})
	// ShrU_I64 = _add(Instr{N: "shru_i64", O: Op_ShrU_I64})
	// Rotl_I64 = _add(Instr{N: "rotl_i64", O: Op_Rotl_I64})
	// Rotr_I64 = _add(Instr{N: "rotr_i64", O: Op_Rotr_I64})

	// Abs_F32 = _add(Instr{N: "abs_f32", O: Op_Abs_F32})
	// Neg_F32 = _add(Instr{N: "neg_f32", O: Op_Neg_F32})
	// Ceil_F32 = _add(Instr{N: "ceil_f32", O: Op_Ceil_F32})
	// Floor_F32 = _add(Instr{N: "floor_f32", O: Op_Floor_F32})
	// Trunc_F32 = _add(Instr{N: "trunc_f32", O: Op_Trunc_F32})
	// Nearest_F32 = _add(Instr{N: "nearest_f32", O: Op_Nearest_F32})
	// Sqrt_F32 = _add(Instr{N: "sqrt_f32", O: Op_Sqrt_F32})
	// Add_F32 = _add(Instr{N: "add_f32", O: Op_Add_F32})
	// Sub_F32 = _add(Instr{N: "sub_f32", O: Op_Sub_F32})
	// Mul_F32 = _add(Instr{N: "mul_f32", O: Op_Mul_F32})
	// Div_F32 = _add(Instr{N: "div_f32", O: Op_Div_F32})
	// Min_F32 = _add(Instr{N: "min_f32", O: Op_Min_F32})
	// Max_F32 = _add(Instr{N: "max_f32", O: Op_Max_F32})
	// Copysign_F32 = _add(Instr{N: "copysign_f32", O: Op_Copysign_F32})

	// Abs_F64 = _add(Instr{N: "abs_f64", O: Op_Abs_F64})
	// Neg_F64 = _add(Instr{N: "neg_f64", O: Op_Neg_F64})
	// Ceil_F64 = _add(Instr{N: "ceil_f64", O: Op_Ceil_F64})
	// Floor_F64 = _add(Instr{N: "floor_f64", O: Op_Floor_F64})
	// Trunc_F64 = _add(Instr{N: "trunc_f64", O: Op_Trunc_F64})
	// Nearest_F64 = _add(Instr{N: "nearest_f64", O: Op_Nearest_F64})
	// Sqrt_F64 = _add(Instr{N: "sqrt_f64", O: Op_Sqrt_F64})
	// Add_F64 = _add(Instr{N: "add_f64", O: Op_Add_F64})
	// Sub_F64 = _add(Instr{N: "sub_f64", O: Op_Sub_F64})
	// Mul_F64 = _add(Instr{N: "mul_f64", O: Op_Mul_F64})
	// Div_F64 = _add(Instr{N: "div_f64", O: Op_Div_F64})
	// Min_F64 = _add(Instr{N: "min_f64", O: Op_Min_F64})
	// Max_F64 = _add(Instr{N: "max_f64", O: Op_Max_F64})
	// Copysign_F64 = _add(Instr{N: "copysign_f64", O: Op_Copysign_F64})

	// WrapI64_I32 = _add(Instr{N: "wrapi64_i32", O: Op_WrapI64_I32})
	// TruncF32S_I32 = _add(Instr{N: "truncf32s_i32", O: Op_TruncF32S_I32})
	// TruncF32U_I32 = _add(Instr{N: "truncf32u_i32", O: Op_TruncF32U_I32})
	// TruncF64S_I32 = _add(Instr{N: "truncf64s_i32", O: Op_TruncF64S_I32})
	// TruncF64U_I32 = _add(Instr{N: "truncf64u_i32", O: Op_TruncF64U_I32})
	// ExtendI32S_I64 = _add(Instr{N: "extendi32s_i64", O: Op_ExtendI32S_I64})
	// ExtendI32U_I64 = _add(Instr{N: "extendi32u_i64", O: Op_ExtendI32U_I64})
	// TruncF32S_I64 = _add(Instr{N: "truncf32s_i64", O: Op_TruncF32S_I64})
	// TruncF32U_I64 = _add(Instr{N: "truncf32u_i64", O: Op_TruncF32U_I64})
	// TruncF64S_I64 = _add(Instr{N: "truncf64s_i64", O: Op_TruncF64S_I64})
	// TruncF64U_I64 = _add(Instr{N: "truncf64u_i64", O: Op_TruncF64U_I64})
	// ConvertI32S_F32 = _add(Instr{N: "converti32s_f32", O: Op_ConvertI32S_F32})
	// ConvertI32U_F32 = _add(Instr{N: "converti32u_f32", O: Op_ConvertI32U_F32})
	// ConvertI64S_F32 = _add(Instr{N: "converti64s_f32", O: Op_ConvertI64S_F32})
	// ConvertI64U_F32 = _add(Instr{N: "converti64u_f32", O: Op_ConvertI64U_F32})
	// DemoteF64_F32 = _add(Instr{N: "demotef64_f32", O: Op_DemoteF64_F32})
	// ConvertI32S_F64 = _add(Instr{N: "converti32s_f64", O: Op_ConvertI32S_F64})
	// ConvertI32U_F64 = _add(Instr{N: "converti32u_f64", O: Op_ConvertI32U_F64})
	// ConvertI64S_F64 = _add(Instr{N: "converti64s_f64", O: Op_ConvertI64S_F64})
	// ConvertI64U_F64 = _add(Instr{N: "converti64u_f64", O: Op_ConvertI64U_F64})
	// PromoteF32_F64 = _add(Instr{N: "promotef32_f64", O: Op_PromoteF32_F64})
	// ReinterpretF32_I32 = _add(Instr{N: "reinterpretf32_i32", O: Op_ReinterpretF32_I32})
	// ReinterpretF64_I64 = _add(Instr{N: "reinterpretf64_i64", O: Op_ReinterpretF64_I64})
	// ReinterpretI32_F32 = _add(Instr{N: "reinterpreti32_f32", O: Op_ReinterpretI32_F32})
	// ReinterpretI64_F64 = _add(Instr{N: "reinterpreti64_f64", O: Op_ReinterpretI64_F64})

	// Extend8S_I32 = _add(Instr{N: "extend8s_i32", O: Op_Extend8S_I32})
	// Extend16S_I32 = _add(Instr{N: "extend16s_i32", O: Op_Extend16S_I32})
	// Extend8S_I64 = _add(Instr{N: "extend8s_i64", O: Op_Extend8S_I64})
	// Extend16S_I64 = _add(Instr{N: "extend16s_i64", O: Op_Extend16S_I64})
	// Extend32S_I64 = _add(Instr{N: "extend32s_i64", O: Op_Extend32S_I64})

	// TruncSatF32S_I32 = _add(Instr{N: "truncsatf32s_i32", O: Op_TruncSatF32S_I32})
	// TruncSatF32U_I32 = _add(Instr{N: "truncsatf32u_i32", O: Op_TruncSatF32U_I32})
	// TruncSatF64S_I32 = _add(Instr{N: "truncsatf64s_i32", O: Op_TruncSatF64S_I32})
	// TruncSatF64U_I32 = _add(Instr{N: "truncsatf64u_i32", O: Op_TruncSatF64U_I32})
	// TruncSatF32S_I64 = _add(Instr{N: "truncsatf32s_i64", O: Op_TruncSatF32S_I64})
	// TruncSatF32U_I64 = _add(Instr{N: "truncsatf32u_i64", O: Op_TruncSatF32U_I64})
	// TruncSatF64S_I64 = _add(Instr{N: "truncsatf64s_i64", O: Op_TruncSatF64S_I64})
	// TruncSatF64U_I64 = _add(Instr{N: "truncsatf64u_i64", O: Op_TruncSatF64U_I64})

	// vector

	// Load_V128 = _add(Instr{N: "load_v128", O: Op_Load_V128})
	// Load8x8S_V128 = _add(Instr{N: "load8x8s_v128", O: Op_Load8x8S_V128})
	// Load8x8U_V128 = _add(Instr{N: "load8x8u_v128", O: Op_Load8x8U_V128})
	// Load16x4S_V128 = _add(Instr{N: "load16x4s_v128", O: Op_Load16x4S_V128})
	// Load16x4U_V128 = _add(Instr{N: "load16x4u_v128", O: Op_Load16x4U_V128})
	// Load32x2S_V128 = _add(Instr{N: "load32x2s_v128", O: Op_Load32x2S_V128})
	// Load32x2U_V128 = _add(Instr{N: "load32x2u_v128", O: Op_Load32x2U_V128})
	// Load8Splat_V128 = _add(Instr{N: "load8splat_v128", O: Op_Load8Splat_V128})
	// Load16Splat_V128 = _add(Instr{N: "load16splat_v128", O: Op_Load16Splat_V128})
	// Load32Splat_V128 = _add(Instr{N: "load32splat_v128", O: Op_Load32Splat_V128})
	// Load64Splat_V128 = _add(Instr{N: "load64splat_v128", O: Op_Load64Splat_V128})
	// Load32Zero_V128 = _add(Instr{N: "load32zero_v128", O: Op_Load32Zero_V128})
	// Load64Zero_V128 = _add(Instr{N: "load64zero_v128", O: Op_Load64Zero_V128})
	// Store_V128 = _add(Instr{N: "store_v128", O: Op_Store_V128})
	// Load8Lane_V128 = _add(Instr{N: "load8lane_v128", O: Op_Load8Lane_V128})
	// Load16Lane_V128 = _add(Instr{N: "load16lane_v128", O: Op_Load16Lane_V128})
	// Load32Lane_V128 = _add(Instr{N: "load32lane_v128", O: Op_Load32Lane_V128})
	// Load64Lane_V128 = _add(Instr{N: "load64lane_v128", O: Op_Load64Lane_V128})
	// Store8Lane_V128 = _add(Instr{N: "store8lane_v128", O: Op_Store8Lane_V128})
	// Store16Lane_V128 = _add(Instr{N: "store16lane_v128", O: Op_Store16Lane_V128})
	// Store32Lane_V128 = _add(Instr{N: "store32lane_v128", O: Op_Store32Lane_V128})
	// Store64Lane_V128 = _add(Instr{N: "store64lane_v128", O: Op_Store64Lane_V128})

	// Const_V128 = _add(Instr{N: "const_v128", O: Op_Const_V128})

	// Shuffle_I8X16 = _add(Instr{N: "shuffle_i8x16", O: Op_Shuffle_I8X16})

	// ExtractLaneS_I8X16 = _add(Instr{N: "extractlanes_i8x16", O: Op_ExtractLaneS_I8X16})
	// ExtractLaneU_I8X16 = _add(Instr{N: "extractlaneu_i8x16", O: Op_ExtractLaneU_I8X16})
	// ReplaceLane_I8X16 = _add(Instr{N: "replacelane_i8x16", O: Op_ReplaceLane_I8X16})
	// ExtractLaneS_I16X8 = _add(Instr{N: "extractlanes_i16x8", O: Op_ExtractLaneS_I16X8})
	// ExtractLaneU_I16X8 = _add(Instr{N: "extractlaneu_i16x8", O: Op_ExtractLaneU_I16X8})
	// ReplaceLane_I16X8 = _add(Instr{N: "replacelane_i16x8", O: Op_ReplaceLane_I16X8})
	// ExtractLane_I32X4 = _add(Instr{N: "extractlane_i32x4", O: Op_ExtractLane_I32X4})
	// ReplaceLane_I32X4 = _add(Instr{N: "replacelane_i32x4", O: Op_ReplaceLane_I32X4})
	// ExtractLane_I64X2 = _add(Instr{N: "extractlane_i64x2", O: Op_ExtractLane_I64X2})
	// ReplaceLane_I64X2 = _add(Instr{N: "replacelane_i64x2", O: Op_ReplaceLane_I64X2})
	// ExtractLane_F32X4 = _add(Instr{N: "extractlane_f32x4", O: Op_ExtractLane_F32X4})
	// ReplaceLane_F32X4 = _add(Instr{N: "replacelane_f32x4", O: Op_ReplaceLane_F32X4})
	// ExtractLane_F64X2 = _add(Instr{N: "extractlane_f64x2", O: Op_ExtractLane_F64X2})
	// ReplaceLane_F64X2 = _add(Instr{N: "replacelane_f64x2", O: Op_ReplaceLane_F64X2})

	// Swizzle_I8X16 = _add(Instr{N: "swizzle_i8x16", O: Op_Swizzle_I8X16})
	// Splat_I8X16 = _add(Instr{N: "splat_i8x16", O: Op_Splat_I8X16})
	// Splat_I16X8 = _add(Instr{N: "splat_i16x8", O: Op_Splat_I16X8})
	// Splat_I32X4 = _add(Instr{N: "splat_i32x4", O: Op_Splat_I32X4})
	// Splat_I64X2 = _add(Instr{N: "splat_i64x2", O: Op_Splat_I64X2})
	// Splat_F32X4 = _add(Instr{N: "splat_f32x4", O: Op_Splat_F32X4})
	// Splat_F64X2 = _add(Instr{N: "splat_f64x2", O: Op_Splat_F64X2})

	// Eq_I8X16 = _add(Instr{N: "eq_i8x16", O: Op_Eq_I8X16})
	// Ne_I8X16 = _add(Instr{N: "ne_i8x16", O: Op_Ne_I8X16})
	// LtS_I8X16 = _add(Instr{N: "lts_i8x16", O: Op_LtS_I8X16})
	// LtU_I8X16 = _add(Instr{N: "ltu_i8x16", O: Op_LtU_I8X16})
	// GtS_I8X16 = _add(Instr{N: "gts_i8x16", O: Op_GtS_I8X16})
	// GtU_I8X16 = _add(Instr{N: "gtu_i8x16", O: Op_GtU_I8X16})
	// LeS_I8X16 = _add(Instr{N: "les_i8x16", O: Op_LeS_I8X16})
	// LeU_I8X16 = _add(Instr{N: "leu_i8x16", O: Op_LeU_I8X16})
	// GeS_I8X16 = _add(Instr{N: "ges_i8x16", O: Op_GeS_I8X16})
	// GeU_I8X16 = _add(Instr{N: "geu_i8x16", O: Op_GeU_I8X16})

	// Eq_I16X8 = _add(Instr{N: "eq_i16x8", O: Op_Eq_I16X8})
	// Ne_I16X8 = _add(Instr{N: "ne_i16x8", O: Op_Ne_I16X8})
	// LtS_I16X8 = _add(Instr{N: "lts_i16x8", O: Op_LtS_I16X8})
	// LtU_I16X8 = _add(Instr{N: "ltu_i16x8", O: Op_LtU_I16X8})
	// GtS_I16X8 = _add(Instr{N: "gts_i16x8", O: Op_GtS_I16X8})
	// GtU_I16X8 = _add(Instr{N: "gtu_i16x8", O: Op_GtU_I16X8})
	// LeS_I16X8 = _add(Instr{N: "les_i16x8", O: Op_LeS_I16X8})
	// LeU_I16X8 = _add(Instr{N: "leu_i16x8", O: Op_LeU_I16X8})
	// GeS_I16X8 = _add(Instr{N: "ges_i16x8", O: Op_GeS_I16X8})
	// GeU_I16X8 = _add(Instr{N: "geu_i16x8", O: Op_GeU_I16X8})

	// Eq_I32X4 = _add(Instr{N: "eq_i32x4", O: Op_Eq_I32X4})
	// Ne_I32X4 = _add(Instr{N: "ne_i32x4", O: Op_Ne_I32X4})
	// LtS_I32X4 = _add(Instr{N: "lts_i32x4", O: Op_LtS_I32X4})
	// LtU_I32X4 = _add(Instr{N: "ltu_i32x4", O: Op_LtU_I32X4})
	// GtS_I32X4 = _add(Instr{N: "gts_i32x4", O: Op_GtS_I32X4})
	// GtU_I32X4 = _add(Instr{N: "gtu_i32x4", O: Op_GtU_I32X4})
	// LeS_I32X4 = _add(Instr{N: "les_i32x4", O: Op_LeS_I32X4})
	// LeU_I32X4 = _add(Instr{N: "leu_i32x4", O: Op_LeU_I32X4})
	// GeS_I32X4 = _add(Instr{N: "ges_i32x4", O: Op_GeS_I32X4})
	// GeU_I32X4 = _add(Instr{N: "geu_i32x4", O: Op_GeU_I32X4})

	// Eq_I64X2 = _add(Instr{N: "eq_i64x2", O: Op_Eq_I64X2})
	// Ne_I64X2 = _add(Instr{N: "ne_i64x2", O: Op_Ne_I64X2})
	// LtS_I64X2 = _add(Instr{N: "lts_i64x2", O: Op_LtS_I64X2})
	// GtS_I64X2 = _add(Instr{N: "gts_i64x2", O: Op_GtS_I64X2})
	// LeS_I64X2 = _add(Instr{N: "les_i64x2", O: Op_LeS_I64X2})
	// GeS_I64X2 = _add(Instr{N: "ges_i64x2", O: Op_GeS_I64X2})

	// Eq_F32X4 = _add(Instr{N: "eq_f32x4", O: Op_Eq_F32X4})
	// Ne_F32X4 = _add(Instr{N: "ne_f32x4", O: Op_Ne_F32X4})
	// Lt_F32X4 = _add(Instr{N: "lt_f32x4", O: Op_Lt_F32X4})
	// Gt_F32X4 = _add(Instr{N: "gt_f32x4", O: Op_Gt_F32X4})
	// Le_F32X4 = _add(Instr{N: "le_f32x4", O: Op_Le_F32X4})
	// Ge_F32X4 = _add(Instr{N: "ge_f32x4", O: Op_Ge_F32X4})

	// Eq_F64X2 = _add(Instr{N: "eq_f64x2", O: Op_Eq_F64X2})
	// Ne_F64X2 = _add(Instr{N: "ne_f64x2", O: Op_Ne_F64X2})
	// Lt_F64X2 = _add(Instr{N: "lt_f64x2", O: Op_Lt_F64X2})
	// Gt_F64X2 = _add(Instr{N: "gt_f64x2", O: Op_Gt_F64X2})
	// Le_F64X2 = _add(Instr{N: "le_f64x2", O: Op_Le_F64X2})
	// Ge_F64X2 = _add(Instr{N: "ge_f64x2", O: Op_Ge_F64X2})

	// Not_v128 = _add(Instr{N: "not_v128", O: Op_Not_v128})
	// And_v128 = _add(Instr{N: "and_v128", O: Op_And_v128})
	// Andnot_v128 = _add(Instr{N: "andnot_v128", O: Op_Andnot_v128})
	// Or_v128 = _add(Instr{N: "or_v128", O: Op_Or_v128})
	// Xor_v128 = _add(Instr{N: "xor_v128", O: Op_Xor_v128})
	// Bitselect_v128 = _add(Instr{N: "bitselect_v128", O: Op_Bitselect_v128})
	// Op_AnyTrue_v128   = 0xFD_53

	// Abs_I8X16 = _add(Instr{N: "abs_i8x16", O: Op_Abs_I8X16})
	// Neg_I8X16 = _add(Instr{N: "neg_i8x16", O: Op_Neg_I8X16})
	// Popcnt_I8X16 = _add(Instr{N: "popcnt_i8x16", O: Op_Popcnt_I8X16})
	// AllTrue_I8X16 = _add(Instr{N: "alltrue_i8x16", O: Op_AllTrue_I8X16})
	// Bitmask_I8X16 = _add(Instr{N: "bitmask_i8x16", O: Op_Bitmask_I8X16})
	// NarrowI16X8S_I8X16 = _add(Instr{N: "narrowi16x8s_i8x16", O: Op_NarrowI16X8S_I8X16})
	// NarrowI16X8U_I8X16 = _add(Instr{N: "narrowi16x8u_i8x16", O: Op_NarrowI16X8U_I8X16})
	// Shl_I8X16 = _add(Instr{N: "shl_i8x16", O: Op_Shl_I8X16})
	// ShrS_I8X16 = _add(Instr{N: "shrs_i8x16", O: Op_ShrS_I8X16})
	// ShrU_I8X16 = _add(Instr{N: "shru_i8x16", O: Op_ShrU_I8X16})
	// Add_I8X16 = _add(Instr{N: "add_i8x16", O: Op_Add_I8X16})
	// AddSatS_I8X16 = _add(Instr{N: "addsats_i8x16", O: Op_AddSatS_I8X16})
	// AddSatU_I8X16 = _add(Instr{N: "addsatu_i8x16", O: Op_AddSatU_I8X16})
	// Sub_I8X16 = _add(Instr{N: "sub_i8x16", O: Op_Sub_I8X16})
	// SubSatS_I8X16 = _add(Instr{N: "subsats_i8x16", O: Op_SubSatS_I8X16})
	// SubSatU_I8X16 = _add(Instr{N: "subsatu_i8x16", O: Op_SubSatU_I8X16})
	// MinS_I8X16 = _add(Instr{N: "mins_i8x16", O: Op_MinS_I8X16})
	// MinU_I8X16 = _add(Instr{N: "minu_i8x16", O: Op_MinU_I8X16})
	// MaxS_I8X16 = _add(Instr{N: "maxs_i8x16", O: Op_MaxS_I8X16})
	// MaxU_I8X16 = _add(Instr{N: "maxu_i8x16", O: Op_MaxU_I8X16})
	// AvgrU_I8X16 = _add(Instr{N: "avgru_i8x16", O: Op_AvgrU_I8X16})

	// ExtaddPairwiseI8X16S_I16X8 = _add(Instr{N: "extaddpairwisei8x16s_i16x8", O: Op_ExtaddPairwiseI8X16S_I16X8})
	// ExtaddPairwiseI8X16U_I16X8 = _add(Instr{N: "extaddpairwisei8x16u_i16x8", O: Op_ExtaddPairwiseI8X16U_I16X8})
	// Abs_I16X8 = _add(Instr{N: "abs_i16x8", O: Op_Abs_I16X8})
	// Neg_I16X8 = _add(Instr{N: "neg_i16x8", O: Op_Neg_I16X8})
	// Q15mulrSatS_I16X8 = _add(Instr{N: "q15mulrsats_i16x8", O: Op_Q15mulrSatS_I16X8})
	// AllTrue_I16X8 = _add(Instr{N: "alltrue_i16x8", O: Op_AllTrue_I16X8})
	// Bitmask_I16X8 = _add(Instr{N: "bitmask_i16x8", O: Op_Bitmask_I16X8})
	// NarrowI32X4S_I16X8 = _add(Instr{N: "narrowi32x4s_i16x8", O: Op_NarrowI32X4S_I16X8})
	// NarrowI32X4U_I16X8 = _add(Instr{N: "narrowi32x4u_i16x8", O: Op_NarrowI32X4U_I16X8})
	// ExtendLowI8X16S_I16X8 = _add(Instr{N: "extendlowi8x16s_i16x8", O: Op_ExtendLowI8X16S_I16X8})
	// ExtendHighI8X16S_I16X8 = _add(Instr{N: "extendhighi8x16s_i16x8", O: Op_ExtendHighI8X16S_I16X8})
	// ExtendLowI8X16U_I16X8 = _add(Instr{N: "extendlowi8x16u_i16x8", O: Op_ExtendLowI8X16U_I16X8})
	// ExtendHighI8X16U_I16X8 = _add(Instr{N: "extendhighi8x16u_i16x8", O: Op_ExtendHighI8X16U_I16X8})
	// Shl_I16X8 = _add(Instr{N: "shl_i16x8", O: Op_Shl_I16X8})
	// ShrS_I16X8 = _add(Instr{N: "shrs_i16x8", O: Op_ShrS_I16X8})
	// ShrU_I16X8 = _add(Instr{N: "shru_i16x8", O: Op_ShrU_I16X8})
	// Add_I16X8 = _add(Instr{N: "add_i16x8", O: Op_Add_I16X8})
	// AddSatS_I16X8 = _add(Instr{N: "addsats_i16x8", O: Op_AddSatS_I16X8})
	// AddSatU_I16X8 = _add(Instr{N: "addsatu_i16x8", O: Op_AddSatU_I16X8})
	// Sub_I16X8 = _add(Instr{N: "sub_i16x8", O: Op_Sub_I16X8})
	// SubSatS_I16X8 = _add(Instr{N: "subsats_i16x8", O: Op_SubSatS_I16X8})
	// SubSatU_I16X8 = _add(Instr{N: "subsatu_i16x8", O: Op_SubSatU_I16X8})
	// Mul_I16X8 = _add(Instr{N: "mul_i16x8", O: Op_Mul_I16X8})
	// MinS_I16X8 = _add(Instr{N: "mins_i16x8", O: Op_MinS_I16X8})
	// MinU_I16X8 = _add(Instr{N: "minu_i16x8", O: Op_MinU_I16X8})
	// MaxS_I16X8 = _add(Instr{N: "maxs_i16x8", O: Op_MaxS_I16X8})
	// MaxU_I16X8 = _add(Instr{N: "maxu_i16x8", O: Op_MaxU_I16X8})
	// AvgrU_I16X8 = _add(Instr{N: "avgru_i16x8", O: Op_AvgrU_I16X8})
	// ExtmulLowI8X16S_I16X8 = _add(Instr{N: "extmullowi8x16s_i16x8", O: Op_ExtmulLowI8X16S_I16X8})
	// ExtmulHighI8X16S_I16X8 = _add(Instr{N: "extmulhighi8x16s_i16x8", O: Op_ExtmulHighI8X16S_I16X8})
	// ExtmulLowI8X16U_I16X8 = _add(Instr{N: "extmullowi8x16u_i16x8", O: Op_ExtmulLowI8X16U_I16X8})
	// ExtmulHighI8X16U_I16X8 = _add(Instr{N: "extmulhighi8x16u_i16x8", O: Op_ExtmulHighI8X16U_I16X8})

	// ExtaddPairwiseI16X8S_I32X4 = _add(Instr{N: "extaddpairwisei16x8s_i32x4", O: Op_ExtaddPairwiseI16X8S_I32X4})
	// ExtaddPairwiseI16X8U_I32X4 = _add(Instr{N: "extaddpairwisei16x8u_i32x4", O: Op_ExtaddPairwiseI16X8U_I32X4})
	// Abs_I32X4 = _add(Instr{N: "abs_i32x4", O: Op_Abs_I32X4})
	// Neg_I32X4 = _add(Instr{N: "neg_i32x4", O: Op_Neg_I32X4})
	// AllTrue_I32X4 = _add(Instr{N: "alltrue_i32x4", O: Op_AllTrue_I32X4})
	// Bitmask_I32X4 = _add(Instr{N: "bitmask_i32x4", O: Op_Bitmask_I32X4})
	// ExtendLowI16X8S_I32X4 = _add(Instr{N: "extendlowi16x8s_i32x4", O: Op_ExtendLowI16X8S_I32X4})
	// ExtendHighI16X8S_I32X4 = _add(Instr{N: "extendhighi16x8s_i32x4", O: Op_ExtendHighI16X8S_I32X4})
	// ExtendLowI16X8U_I32X4 = _add(Instr{N: "extendlowi16x8u_i32x4", O: Op_ExtendLowI16X8U_I32X4})
	// ExtendHighI16X8U_I32X4 = _add(Instr{N: "extendhighi16x8u_i32x4", O: Op_ExtendHighI16X8U_I32X4})
	// Shl_I32X4 = _add(Instr{N: "shl_i32x4", O: Op_Shl_I32X4})
	// ShrS_I32X4 = _add(Instr{N: "shrs_i32x4", O: Op_ShrS_I32X4})
	// ShrU_I32X4 = _add(Instr{N: "shru_i32x4", O: Op_ShrU_I32X4})
	// Add_I32X4 = _add(Instr{N: "add_i32x4", O: Op_Add_I32X4})
	// Sub_I32X4 = _add(Instr{N: "sub_i32x4", O: Op_Sub_I32X4})
	// Mul_I32X4 = _add(Instr{N: "mul_i32x4", O: Op_Mul_I32X4})
	// MinS_I32X4 = _add(Instr{N: "mins_i32x4", O: Op_MinS_I32X4})
	// MinU_I32X4 = _add(Instr{N: "minu_i32x4", O: Op_MinU_I32X4})
	// MaxS_I32X4 = _add(Instr{N: "maxs_i32x4", O: Op_MaxS_I32X4})
	// MaxU_I32X4 = _add(Instr{N: "maxu_i32x4", O: Op_MaxU_I32X4})
	// DotI16X8S_I32X4 = _add(Instr{N: "doti16x8s_i32x4", O: Op_DotI16X8S_I32X4})
	// ExtmulLowI16X8S_I32X4 = _add(Instr{N: "extmullowi16x8s_i32x4", O: Op_ExtmulLowI16X8S_I32X4})
	// ExtmulHighI16X8S_I32X4 = _add(Instr{N: "extmulhighi16x8s_i32x4", O: Op_ExtmulHighI16X8S_I32X4})
	// ExtmulLowI16X8U_I32X4 = _add(Instr{N: "extmullowi16x8u_i32x4", O: Op_ExtmulLowI16X8U_I32X4})
	// ExtmulHighI16X8U_I32X4 = _add(Instr{N: "extmulhighi16x8u_i32x4", O: Op_ExtmulHighI16X8U_I32X4})

	// Abs_I64X2 = _add(Instr{N: "abs_i64x2", O: Op_Abs_I64X2})
	// Neg_I64X2 = _add(Instr{N: "neg_i64x2", O: Op_Neg_I64X2})
	// AllTrue_I64X2 = _add(Instr{N: "alltrue_i64x2", O: Op_AllTrue_I64X2})
	// Bitmask_I64X2 = _add(Instr{N: "bitmask_i64x2", O: Op_Bitmask_I64X2})
	// ExtendLowI32X4S_I64X2 = _add(Instr{N: "extendlowi32x4s_i64x2", O: Op_ExtendLowI32X4S_I64X2})
	// ExtendHighI32X4S_I64X2 = _add(Instr{N: "extendhighi32x4s_i64x2", O: Op_ExtendHighI32X4S_I64X2})
	// ExtendLowI32X4U_I64X2 = _add(Instr{N: "extendlowi32x4u_i64x2", O: Op_ExtendLowI32X4U_I64X2})
	// ExtendHighI32X4U_I64X2 = _add(Instr{N: "extendhighi32x4u_i64x2", O: Op_ExtendHighI32X4U_I64X2})
	// Shl_I64X2 = _add(Instr{N: "shl_i64x2", O: Op_Shl_I64X2})
	// ShrS_I64X2 = _add(Instr{N: "shrs_i64x2", O: Op_ShrS_I64X2})
	// ShrU_I64X2 = _add(Instr{N: "shru_i64x2", O: Op_ShrU_I64X2})
	// Add_I64X2 = _add(Instr{N: "add_i64x2", O: Op_Add_I64X2})
	// Sub_I64X2 = _add(Instr{N: "sub_i64x2", O: Op_Sub_I64X2})
	// Mul_I64X2 = _add(Instr{N: "mul_i64x2", O: Op_Mul_I64X2})
	// ExtmulLowI32X4S_I64X2 = _add(Instr{N: "extmullowi32x4s_i64x2", O: Op_ExtmulLowI32X4S_I64X2})
	// ExtmulHighI32X4S_I64X2 = _add(Instr{N: "extmulhighi32x4s_i64x2", O: Op_ExtmulHighI32X4S_I64X2})
	// ExtmulLowI32X4U_I64X2 = _add(Instr{N: "extmullowi32x4u_i64x2", O: Op_ExtmulLowI32X4U_I64X2})
	// ExtmulHighI32X4U_I64X2 = _add(Instr{N: "extmulhighi32x4u_i64x2", O: Op_ExtmulHighI32X4U_I64X2})

	// Ceil_F32X4 = _add(Instr{N: "ceil_f32x4", O: Op_Ceil_F32X4})
	// Floor_F32X4 = _add(Instr{N: "floor_f32x4", O: Op_Floor_F32X4})
	// Trunc_F32X4 = _add(Instr{N: "trunc_f32x4", O: Op_Trunc_F32X4})
	// Nearest_F32X4 = _add(Instr{N: "nearest_f32x4", O: Op_Nearest_F32X4})
	// Abs_F32X4 = _add(Instr{N: "abs_f32x4", O: Op_Abs_F32X4})
	// Neg_F32X4 = _add(Instr{N: "neg_f32x4", O: Op_Neg_F32X4})
	// Sqrt_F32X4 = _add(Instr{N: "sqrt_f32x4", O: Op_Sqrt_F32X4})
	// Add_F32X4 = _add(Instr{N: "add_f32x4", O: Op_Add_F32X4})
	// Sub_F32X4 = _add(Instr{N: "sub_f32x4", O: Op_Sub_F32X4})
	// Mul_F32X4 = _add(Instr{N: "mul_f32x4", O: Op_Mul_F32X4})
	// Div_F32X4 = _add(Instr{N: "div_f32x4", O: Op_Div_F32X4})
	// Min_F32X4 = _add(Instr{N: "min_f32x4", O: Op_Min_F32X4})
	// Max_F32X4 = _add(Instr{N: "max_f32x4", O: Op_Max_F32X4})
	// Pmin_F32X4 = _add(Instr{N: "pmin_f32x4", O: Op_Pmin_F32X4})
	// Pmax_F32X4 = _add(Instr{N: "pmax_f32x4", O: Op_Pmax_F32X4})

	// Ceil_F64X2 = _add(Instr{N: "ceil_f64x2", O: Op_Ceil_F64X2})
	// Floor_F64X2 = _add(Instr{N: "floor_f64x2", O: Op_Floor_F64X2})
	// Trunc_F64X2 = _add(Instr{N: "trunc_f64x2", O: Op_Trunc_F64X2})
	// Nearest_F64X2 = _add(Instr{N: "nearest_f64x2", O: Op_Nearest_F64X2})
	// Abs_F64X2 = _add(Instr{N: "abs_f64x2", O: Op_Abs_F64X2})
	// Neg_F64X2 = _add(Instr{N: "neg_f64x2", O: Op_Neg_F64X2})
	// Sqrt_F64X2 = _add(Instr{N: "sqrt_f64x2", O: Op_Sqrt_F64X2})
	// Add_F64X2 = _add(Instr{N: "add_f64x2", O: Op_Add_F64X2})
	// Sub_F64X2 = _add(Instr{N: "sub_f64x2", O: Op_Sub_F64X2})
	// Mul_F64X2 = _add(Instr{N: "mul_f64x2", O: Op_Mul_F64X2})
	// Div_F64X2 = _add(Instr{N: "div_f64x2", O: Op_Div_F64X2})
	// Min_F64X2 = _add(Instr{N: "min_f64x2", O: Op_Min_F64X2})
	// Max_F64X2 = _add(Instr{N: "max_f64x2", O: Op_Max_F64X2})
	// Pmin_F64X2 = _add(Instr{N: "pmin_f64x2", O: Op_Pmin_F64X2})
	// Pmax_F64X2 = _add(Instr{N: "pmax_f64x2", O: Op_Pmax_F64X2})

	// TruncSatF32X4S_I32X4 = _add(Instr{N: "truncsatf32x4s_i32x4", O: Op_TruncSatF32X4S_I32X4})
	// TruncSatF32X4U_I32X4 = _add(Instr{N: "truncsatf32x4u_i32x4", O: Op_TruncSatF32X4U_I32X4})
	// ConvertI32X4S_F32X4 = _add(Instr{N: "converti32x4s_f32x4", O: Op_ConvertI32X4S_F32X4})
	// ConvertI32X4U_F32X4 = _add(Instr{N: "converti32x4u_f32x4", O: Op_ConvertI32X4U_F32X4})
	// TruncSatF64X2S_Zero_I32X4 = _add(Instr{N: "truncsatf64x2s_zero_i32x4", O: Op_TruncSatF64X2S_Zero_I32X4})
	// TruncSatF64X2U_Zero_I32X4 = _add(Instr{N: "truncsatf64x2u_zero_i32x4", O: Op_TruncSatF64X2U_Zero_I32X4})
	// ConvertLowI32X4S_F64X2 = _add(Instr{N: "convertlowi32x4s_f64x2", O: Op_ConvertLowI32X4S_F64X2})
	// ConvertLowI32X4U_F64X2 = _add(Instr{N: "convertlowi32x4u_f64x2", O: Op_ConvertLowI32X4U_F64X2})
	// DemoteF64X2Zero_F32X4 = _add(Instr{N: "demotef64x2zero_f32x4", O: Op_DemoteF64X2Zero_F32X4})
	// PromoteLowF32X4_F64X2 = _add(Instr{N: "promotelowf32x4_f64x2", O: Op_PromoteLowF32X4_F64X2})
)
