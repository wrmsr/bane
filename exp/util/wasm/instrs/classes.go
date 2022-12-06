package instrs

type Class int8

const (
	Control Class = iota + 1
	Memory
	Numeric
	Parametric
	Reference
	Table
	Variable
	Vector
)

func (c Class) String() string {
	switch c {
	case 0:
		return "?"
	case Control:
		return "control"
	case Memory:
		return "memory"
	case Numeric:
		return "numeric"
	case Parametric:
		return "parametric"
	case Reference:
		return "reference"
	case Table:
		return "table"
	case Variable:
		return "variable"
	case Vector:
		return "vector"
	}
	panic(c)
}

func (o Opcode) Class() Class {
	switch o {

	case
		Op_Unreachable,
		Op_Nop,
		Op_Block,
		Op_Loop,
		Op_If,
		Op_Else,
		Op_End,
		Op_Br,
		Op_BrIf,
		Op_BrTable,
		Op_Return,
		Op_Call,
		Op_CallIndirect:

		return Control

	case
		Op_RefNull,
		Op_RefIsNull,
		Op_RefFunc:

		return Reference

	case
		Op_Drop,
		Op_Select,
		Op_SelectType:

		return Parametric

	case
		Op_LocalGet,
		Op_LocalSet,
		Op_LocalTee,
		Op_GlobalGet,
		Op_GlobalSet:

		return Variable

	case
		Op_TableGet,
		Op_TableSet,
		Op_TableInit,
		Op_ElemDrop,
		Op_TableCopy,
		Op_TableGrow,
		Op_TableSize,
		Op_TableFill:

		return Table

	case
		Op_Load_I32,
		Op_Load_I64,
		Op_Load_F32,
		Op_Load_F64,
		Op_Load8S_I32,
		Op_Load8U_I32,
		Op_Load16S_I32,
		Op_Load16U_I32,
		Op_Load8S_I64,
		Op_Load8U_I64,
		Op_Load16S_I64,
		Op_Load16U_I64,
		Op_Load32S_I64,
		Op_Load32U_I64,
		Op_Store_I32,
		Op_Store_I64,
		Op_Store_F32,
		Op_Store_F64,
		Op_Store8_I32,
		Op_Store16_I32,
		Op_Store8_I64,
		Op_Store16_I64,
		Op_Store32_I64,
		Op_MemorySize,
		Op_MemoryGrow,
		Op_MemoryInit,
		Op_DataDrop,
		Op_MemoryCopy,
		Op_MemoryFill:

		return Memory

	case
		Op_Const_I32,
		Op_Const_I64,
		Op_Const_F32,
		Op_Const_F64,

		Op_Eqz_I32,
		Op_Eq_I32,
		Op_Ne_I32,
		Op_LtS_I32,
		Op_LtU_I32,
		Op_GtS_I32,
		Op_GtU_I32,
		Op_LeS_I32,
		Op_LeU_I32,
		Op_GeS_I32,
		Op_GeU_I32,

		Op_Eqz_I64,
		Op_Eq_I64,
		Op_Ne_I64,
		Op_LtS_I64,
		Op_LtU_I64,
		Op_GtS_I64,
		Op_GtU_I64,
		Op_LeS_I64,
		Op_LeU_I64,
		Op_GeS_I64,
		Op_GeU_I64,

		Op_Eq_F32,
		Op_Ne_F32,
		Op_Lt_F32,
		Op_Gt_F32,
		Op_Le_F32,
		Op_Ge_F32,

		Op_Eq_F64,
		Op_Ne_F64,
		Op_Lt_F64,
		Op_Gt_F64,
		Op_Le_F64,
		Op_Ge_F64,

		Op_Clz_I32,
		Op_Ctz_I32,
		Op_Popcnt_I32,
		Op_Add_I32,
		Op_Sub_I32,
		Op_Mul_I32,
		Op_DivS_I32,
		Op_DivU_I32,
		Op_RemS_I32,
		Op_RemU_I32,
		Op_And_I32,
		Op_Or_I32,
		Op_Xor_I32,
		Op_Shl_I32,
		Op_ShrS_I32,
		Op_ShrU_I32,
		Op_Rotl_I32,
		Op_Rotr_I32,
		Op_Clz_I64,
		Op_Ctz_I64,
		Op_Popcnt_I64,
		Op_Add_I64,
		Op_Sub_I64,
		Op_Mul_I64,
		Op_DivS_I64,
		Op_DivU_I64,
		Op_RemS_I64,
		Op_RemU_I64,
		Op_And_I64,
		Op_Or_I64,
		Op_Xor_I64,
		Op_Shl_I64,
		Op_ShrS_I64,
		Op_ShrU_I64,
		Op_Rotl_I64,
		Op_Rotr_I64,

		Op_Abs_F32,
		Op_Neg_F32,
		Op_Ceil_F32,
		Op_Floor_F32,
		Op_Trunc_F32,
		Op_Nearest_F32,
		Op_Sqrt_F32,
		Op_Add_F32,
		Op_Sub_F32,
		Op_Mul_F32,
		Op_Div_F32,
		Op_Min_F32,
		Op_Max_F32,
		Op_Copysign_F32,

		Op_Abs_F64,
		Op_Neg_F64,
		Op_Ceil_F64,
		Op_Floor_F64,
		Op_Trunc_F64,
		Op_Nearest_F64,
		Op_Sqrt_F64,
		Op_Add_F64,
		Op_Sub_F64,
		Op_Mul_F64,
		Op_Div_F64,
		Op_Min_F64,
		Op_Max_F64,
		Op_Copysign_F64,

		Op_WrapI64_I32,
		Op_TruncF32S_I32,
		Op_TruncF32U_I32,
		Op_TruncF64S_I32,
		Op_TruncF64U_I32,
		Op_ExtendI32S_I64,
		Op_ExtendI32U_I64,
		Op_TruncF32S_I64,
		Op_TruncF32U_I64,
		Op_TruncF64S_I64,
		Op_TruncF64U_I64,
		Op_ConvertI32S_F32,
		Op_ConvertI32U_F32,
		Op_ConvertI64S_F32,
		Op_ConvertI64U_F32,
		Op_DemoteF64_F32,
		Op_ConvertI32S_F64,
		Op_ConvertI32U_F64,
		Op_ConvertI64S_F64,
		Op_ConvertI64U_F64,
		Op_PromoteF32_F64,
		Op_ReinterpretF32_I32,
		Op_ReinterpretF64_I64,
		Op_ReinterpretI32_F32,
		Op_ReinterpretI64_F64,

		Op_Extend8S_I32,
		Op_Extend16S_I32,
		Op_Extend8S_I64,
		Op_Extend16S_I64,
		Op_Extend32S_I64,

		Op_TruncSatF32S_I32,
		Op_TruncSatF32U_I32,
		Op_TruncSatF64S_I32,
		Op_TruncSatF64U_I32,
		Op_TruncSatF32S_I64,
		Op_TruncSatF32U_I64,
		Op_TruncSatF64S_I64,
		Op_TruncSatF64U_I64:

		return Numeric

	case
		Op_Load_V128,
		Op_Load8x8S_V128,
		Op_Load8x8U_V128,
		Op_Load16x4S_V128,
		Op_Load16x4U_V128,
		Op_Load32x2S_V128,
		Op_Load32x2U_V128,
		Op_Load8Splat_V128,
		Op_Load16Splat_V128,
		Op_Load32Splat_V128,
		Op_Load64Splat_V128,
		Op_Load32Zero_V128,
		Op_Load64Zero_V128,
		Op_Store_V128,
		Op_Load8Lane_V128,
		Op_Load16Lane_V128,
		Op_Load32Lane_V128,
		Op_Load64Lane_V128,
		Op_Store8Lane_V128,
		Op_Store16Lane_V128,
		Op_Store32Lane_V128,
		Op_Store64Lane_V128,

		Op_Const_V128,

		Op_Shuffle_I8X16,

		Op_ExtractLaneS_I8X16,
		Op_ExtractLaneU_I8X16,
		Op_ReplaceLane_I8X16,
		Op_ExtractLaneS_I16X8,
		Op_ExtractLaneU_I16X8,
		Op_ReplaceLane_I16X8,
		Op_ExtractLane_I32X4,
		Op_ReplaceLane_I32X4,
		Op_ExtractLane_I64X2,
		Op_ReplaceLane_I64X2,
		Op_ExtractLane_F32X4,
		Op_ReplaceLane_F32X4,
		Op_ExtractLane_F64X2,
		Op_ReplaceLane_F64X2,

		Op_Swizzle_I8X16,
		Op_Splat_I8X16,
		Op_Splat_I16X8,
		Op_Splat_I32X4,
		Op_Splat_I64X2,
		Op_Splat_F32X4,
		Op_Splat_F64X2,

		Op_Eq_I8X16,
		Op_Ne_I8X16,
		Op_LtS_I8X16,
		Op_LtU_I8X16,
		Op_GtS_I8X16,
		Op_GtU_I8X16,
		Op_LeS_I8X16,
		Op_LeU_I8X16,
		Op_GeS_I8X16,
		Op_GeU_I8X16,

		Op_Eq_I16X8,
		Op_Ne_I16X8,
		Op_LtS_I16X8,
		Op_LtU_I16X8,
		Op_GtS_I16X8,
		Op_GtU_I16X8,
		Op_LeS_I16X8,
		Op_LeU_I16X8,
		Op_GeS_I16X8,
		Op_GeU_I16X8,

		Op_Eq_I32X4,
		Op_Ne_I32X4,
		Op_LtS_I32X4,
		Op_LtU_I32X4,
		Op_GtS_I32X4,
		Op_GtU_I32X4,
		Op_LeS_I32X4,
		Op_LeU_I32X4,
		Op_GeS_I32X4,
		Op_GeU_I32X4,

		Op_Eq_I64X2,
		Op_Ne_I64X2,
		Op_LtS_I64X2,
		Op_GtS_I64X2,
		Op_LeS_I64X2,
		Op_GeS_I64X2,

		Op_Eq_F32X4,
		Op_Ne_F32X4,
		Op_Lt_F32X4,
		Op_Gt_F32X4,
		Op_Le_F32X4,
		Op_Ge_F32X4,

		Op_Eq_F64X2,
		Op_Ne_F64X2,
		Op_Lt_F64X2,
		Op_Gt_F64X2,
		Op_Le_F64X2,
		Op_Ge_F64X2,

		Op_Not_v128,
		Op_And_v128,
		Op_Andnot_v128,
		Op_Or_v128,
		Op_Xor_v128,
		Op_Bitselect_v128,
		Op_AnyTrue_v128,

		Op_Abs_I8X16,
		Op_Neg_I8X16,
		Op_Popcnt_I8X16,
		Op_AllTrue_I8X16,
		Op_Bitmask_I8X16,
		Op_NarrowI16X8S_I8X16,
		Op_NarrowI16X8U_I8X16,
		Op_Shl_I8X16,
		Op_ShrS_I8X16,
		Op_ShrU_I8X16,
		Op_Add_I8X16,
		Op_AddSatS_I8X16,
		Op_AddSatU_I8X16,
		Op_Sub_I8X16,
		Op_SubSatS_I8X16,
		Op_SubSatU_I8X16,
		Op_MinS_I8X16,
		Op_MinU_I8X16,
		Op_MaxS_I8X16,
		Op_MaxU_I8X16,
		Op_AvgrU_I8X16,

		Op_ExtaddPairwiseI8X16S_I16X8,
		Op_ExtaddPairwiseI8X16U_I16X8,
		Op_Abs_I16X8,
		Op_Neg_I16X8,
		Op_Q15mulrSatS_I16X8,
		Op_AllTrue_I16X8,
		Op_Bitmask_I16X8,
		Op_NarrowI32X4S_I16X8,
		Op_NarrowI32X4U_I16X8,
		Op_ExtendLowI8X16S_I16X8,
		Op_ExtendHighI8X16S_I16X8,
		Op_ExtendLowI8X16U_I16X8,
		Op_ExtendHighI8X16U_I16X8,
		Op_Shl_I16X8,
		Op_ShrS_I16X8,
		Op_ShrU_I16X8,
		Op_Add_I16X8,
		Op_AddSatS_I16X8,
		Op_AddSatU_I16X8,
		Op_Sub_I16X8,
		Op_SubSatS_I16X8,
		Op_SubSatU_I16X8,
		Op_Mul_I16X8,
		Op_MinS_I16X8,
		Op_MinU_I16X8,
		Op_MaxS_I16X8,
		Op_MaxU_I16X8,
		Op_AvgrU_I16X8,
		Op_ExtmulLowI8X16S_I16X8,
		Op_ExtmulHighI8X16S_I16X8,
		Op_ExtmulLowI8X16U_I16X8,
		Op_ExtmulHighI8X16U_I16X8,

		Op_ExtaddPairwiseI16X8S_I32X4,
		Op_ExtaddPairwiseI16X8U_I32X4,
		Op_Abs_I32X4,
		Op_Neg_I32X4,
		Op_AllTrue_I32X4,
		Op_Bitmask_I32X4,
		Op_ExtendLowI16X8S_I32X4,
		Op_ExtendHighI16X8S_I32X4,
		Op_ExtendLowI16X8U_I32X4,
		Op_ExtendHighI16X8U_I32X4,
		Op_Shl_I32X4,
		Op_ShrS_I32X4,
		Op_ShrU_I32X4,
		Op_Add_I32X4,
		Op_Sub_I32X4,
		Op_Mul_I32X4,
		Op_MinS_I32X4,
		Op_MinU_I32X4,
		Op_MaxS_I32X4,
		Op_MaxU_I32X4,
		Op_DotI16X8S_I32X4,
		Op_ExtmulLowI16X8S_I32X4,
		Op_ExtmulHighI16X8S_I32X4,
		Op_ExtmulLowI16X8U_I32X4,
		Op_ExtmulHighI16X8U_I32X4,

		Op_Abs_I64X2,
		Op_Neg_I64X2,
		Op_AllTrue_I64X2,
		Op_Bitmask_I64X2,
		Op_ExtendLowI32X4S_I64X2,
		Op_ExtendHighI32X4S_I64X2,
		Op_ExtendLowI32X4U_I64X2,
		Op_ExtendHighI32X4U_I64X2,
		Op_Shl_I64X2,
		Op_ShrS_I64X2,
		Op_ShrU_I64X2,
		Op_Add_I64X2,
		Op_Sub_I64X2,
		Op_Mul_I64X2,
		Op_ExtmulLowI32X4S_I64X2,
		Op_ExtmulHighI32X4S_I64X2,
		Op_ExtmulLowI32X4U_I64X2,
		Op_ExtmulHighI32X4U_I64X2,

		Op_Ceil_F32X4,
		Op_Floor_F32X4,
		Op_Trunc_F32X4,
		Op_Nearest_F32X4,
		Op_Abs_F32X4,
		Op_Neg_F32X4,
		Op_Sqrt_F32X4,
		Op_Add_F32X4,
		Op_Sub_F32X4,
		Op_Mul_F32X4,
		Op_Div_F32X4,
		Op_Min_F32X4,
		Op_Max_F32X4,
		Op_Pmin_F32X4,
		Op_Pmax_F32X4,

		Op_Ceil_F64X2,
		Op_Floor_F64X2,
		Op_Trunc_F64X2,
		Op_Nearest_F64X2,
		Op_Abs_F64X2,
		Op_Neg_F64X2,
		Op_Sqrt_F64X2,
		Op_Add_F64X2,
		Op_Sub_F64X2,
		Op_Mul_F64X2,
		Op_Div_F64X2,
		Op_Min_F64X2,
		Op_Max_F64X2,
		Op_Pmin_F64X2,
		Op_Pmax_F64X2,

		Op_TruncSatF32X4S_I32X4,
		Op_TruncSatF32X4U_I32X4,
		Op_ConvertI32X4S_F32X4,
		Op_ConvertI32X4U_F32X4,
		Op_TruncSatF64X2S_Zero_I32X4,
		Op_TruncSatF64X2U_Zero_I32X4,
		Op_ConvertLowI32X4S_F64X2,
		Op_ConvertLowI32X4U_F64X2,
		Op_DemoteF64X2Zero_F32X4,
		Op_PromoteLowF32X4_F64X2:

		return Vector

	}
	panic(o)
}
