package instrs

type Opcode int

const (
	// control

	Op_Unreachable  Opcode = 0x00
	Op_Nop                 = 0x01
	Op_Block               = 0x02
	Op_Loop                = 0x03
	Op_If                  = 0x04
	Op_Else                = 0x05
	Op_End                 = 0x0B
	Op_Br                  = 0x0C
	Op_BrIf                = 0x0D
	Op_BrTable             = 0x0E
	Op_Return              = 0x0F
	Op_Call                = 0x10
	Op_CallIndirect        = 0x11

	// reference

	Op_RefNull   = 0xD0
	Op_RefIsNull = 0xD1
	Op_RefFunc   = 0xD2

	// parametric

	Op_Drop       = 0x1A
	Op_Select     = 0x1B
	Op_SelectType = 0x1C

	// variable

	Op_LocalGet  = 0x20
	Op_LocalSet  = 0x21
	Op_LocalTee  = 0x22
	Op_GlobalGet = 0x23
	Op_GlobalSet = 0x24

	// table

	Op_TableGet  = 0x25
	Op_TableSet  = 0x26
	Op_TableInit = 0xFC_0C
	Op_ElemDrop  = 0xFC_0D
	Op_TableCopy = 0xFC_0E
	Op_TableGrow = 0xFC_0F
	Op_TableSize = 0xFC_10
	Op_TableFill = 0xFC_11

	// memory

	Op_Load_I32    = 0x28
	Op_Load_I64    = 0x29
	Op_Load_F32    = 0x2A
	Op_Load_F64    = 0x2B
	Op_Load8S_I32  = 0x2C
	Op_Load8U_I32  = 0x2D
	Op_Load16S_I32 = 0x2E
	Op_Load16U_I32 = 0x2F
	Op_Load8S_I64  = 0x30
	Op_Load8U_I64  = 0x31
	Op_Load16S_I64 = 0x32
	Op_Load16U_I64 = 0x33
	Op_Load32S_I64 = 0x34
	Op_Load32U_I64 = 0x35
	Op_Store_I32   = 0x36
	Op_Store_I64   = 0x37
	Op_Store_F32   = 0x38
	Op_Store_F64   = 0x39
	Op_Store8_I32  = 0x3A
	Op_Store16_I32 = 0x3B
	Op_Store8_I64  = 0x3C
	Op_Store16_I64 = 0x3D
	Op_Store32_I64 = 0x3E
	Op_MemorySize  = 0x3F_00
	Op_MemoryGrow  = 0x40_00
	Op_MemoryInit  = 0xFC_08
	Op_DataDrop    = 0xFC_09
	Op_MemoryCopy  = 0xFC_0A
	Op_MemoryFill  = 0xFC_0B

	// numeric

	Op_Const_I32 = 0x41
	Op_Const_I64 = 0x42
	Op_Const_F32 = 0x43
	Op_Const_F64 = 0x44

	Op_Eqz_I32 = 0x45
	Op_Eq_I32  = 0x46
	Op_Ne_I32  = 0x47
	Op_LtS_I32 = 0x48
	Op_LtU_I32 = 0x49
	Op_GtS_I32 = 0x4A
	Op_GtU_I32 = 0x4B
	Op_LeS_I32 = 0x4C
	Op_LeU_I32 = 0x4D
	Op_GeS_I32 = 0x4E
	Op_GeU_I32 = 0x4F

	Op_Eqz_I64 = 0x50
	Op_Eq_I64  = 0x51
	Op_Ne_I64  = 0x52
	Op_LtS_I64 = 0x53
	Op_LtU_I64 = 0x54
	Op_GtS_I64 = 0x55
	Op_GtU_I64 = 0x56
	Op_LeS_I64 = 0x57
	Op_LeU_I64 = 0x58
	Op_GeS_I64 = 0x59
	Op_GeU_I64 = 0x5A

	Op_Eq_F32 = 0x5B
	Op_Ne_F32 = 0x5C
	Op_Lt_F32 = 0x5D
	Op_Gt_F32 = 0x5E
	Op_Le_F32 = 0x5F
	Op_Ge_F32 = 0x60

	Op_Eq_F64 = 0x61
	Op_Ne_F64 = 0x62
	Op_Lt_F64 = 0x63
	Op_Gt_F64 = 0x64
	Op_Le_F64 = 0x65
	Op_Ge_F64 = 0x66

	Op_Clz_I32    = 0x67
	Op_Ctz_I32    = 0x68
	Op_Popcnt_I32 = 0x69
	Op_Add_I32    = 0x6A
	Op_Sub_I32    = 0x6B
	Op_Mul_I32    = 0x6C
	Op_DivS_I32   = 0x6D
	Op_DivU_I32   = 0x6E
	Op_RemS_I32   = 0x6F
	Op_RemU_I32   = 0x70
	Op_And_I32    = 0x71
	Op_Or_I32     = 0x72
	Op_Xor_I32    = 0x73
	Op_Shl_I32    = 0x74
	Op_ShrS_I32   = 0x75
	Op_ShrU_I32   = 0x76
	Op_Rotl_I32   = 0x77
	Op_Rotr_I32   = 0x78
	Op_Clz_I64    = 0x79
	Op_Ctz_I64    = 0x7A
	Op_Popcnt_I64 = 0x7B
	Op_Add_I64    = 0x7C
	Op_Sub_I64    = 0x7D
	Op_Mul_I64    = 0x7E
	Op_DivS_I64   = 0x7F
	Op_DivU_I64   = 0x80
	Op_RemS_I64   = 0x81
	Op_RemU_I64   = 0x82
	Op_And_I64    = 0x83
	Op_Or_I64     = 0x84
	Op_Xor_I64    = 0x85
	Op_Shl_I64    = 0x86
	Op_ShrS_I64   = 0x87
	Op_ShrU_I64   = 0x88
	Op_Rotl_I64   = 0x89
	Op_Rotr_I64   = 0x8A

	Op_Abs_F32      = 0x8B
	Op_Neg_F32      = 0x8C
	Op_Ceil_F32     = 0x8D
	Op_Floor_F32    = 0x8E
	Op_Trunc_F32    = 0x8F
	Op_Nearest_F32  = 0x90
	Op_Sqrt_F32     = 0x91
	Op_Add_F32      = 0x92
	Op_Sub_F32      = 0x93
	Op_Mul_F32      = 0x94
	Op_Div_F32      = 0x95
	Op_Min_F32      = 0x96
	Op_Max_F32      = 0x97
	Op_Copysign_F32 = 0x98

	Op_Abs_F64      = 0x99
	Op_Neg_F64      = 0x9A
	Op_Ceil_F64     = 0x9B
	Op_Floor_F64    = 0x9C
	Op_Trunc_F64    = 0x9D
	Op_Nearest_F64  = 0x9E
	Op_Sqrt_F64     = 0x9F
	Op_Add_F64      = 0xA0
	Op_Sub_F64      = 0xA1
	Op_Mul_F64      = 0xA2
	Op_Div_F64      = 0xA3
	Op_Min_F64      = 0xA4
	Op_Max_F64      = 0xA5
	Op_Copysign_F64 = 0xA6

	Op_WrapI64_I32        = 0xA7
	Op_TruncF32S_I32      = 0xA8
	Op_TruncF32U_I32      = 0xA9
	Op_TruncF64S_I32      = 0xAA
	Op_TruncF64U_I32      = 0xAB
	Op_ExtendI32S_I64     = 0xAC
	Op_ExtendI32U_I64     = 0xAD
	Op_TruncF32S_I64      = 0xAE
	Op_TruncF32U_I64      = 0xAF
	Op_TruncF64S_I64      = 0xB0
	Op_TruncF64U_I64      = 0xB1
	Op_ConvertI32S_F32    = 0xB2
	Op_ConvertI32U_F32    = 0xB3
	Op_ConvertI64S_F32    = 0xB4
	Op_ConvertI64U_F32    = 0xB5
	Op_DemoteF64_F32      = 0xB6
	Op_ConvertI32S_F64    = 0xB7
	Op_ConvertI32U_F64    = 0xB8
	Op_ConvertI64S_F64    = 0xB9
	Op_ConvertI64U_F64    = 0xBA
	Op_PromoteF32_F64     = 0xBB
	Op_ReinterpretF32_I32 = 0xBC
	Op_ReinterpretF64_I64 = 0xBD
	Op_ReinterpretI32_F32 = 0xBE
	Op_ReinterpretI64_F64 = 0xBF

	Op_Extend8S_I32  = 0xC0
	Op_Extend16S_I32 = 0xC1
	Op_Extend8S_I64  = 0xC2
	Op_Extend16S_I64 = 0xC3
	Op_Extend32S_I64 = 0xC4

	Op_TruncSatF32S_I32 = 0xFC_00
	Op_TruncSatF32U_I32 = 0xFC_01
	Op_TruncSatF64S_I32 = 0xFC_02
	Op_TruncSatF64U_I32 = 0xFC_03
	Op_TruncSatF32S_I64 = 0xFC_04
	Op_TruncSatF32U_I64 = 0xFC_05
	Op_TruncSatF64S_I64 = 0xFC_06
	Op_TruncSatF64U_I64 = 0xFC_07

	// vector

	Op_Load_V128        = 0xFD_00
	Op_Load8x8S_V128    = 0xFD_01
	Op_Load8x8U_V128    = 0xFD_02
	Op_Load16x4S_V128   = 0xFD_03
	Op_Load16x4U_V128   = 0xFD_04
	Op_Load32x2S_V128   = 0xFD_05
	Op_Load32x2U_V128   = 0xFD_06
	Op_Load8Splat_V128  = 0xFD_07
	Op_Load16Splat_V128 = 0xFD_08
	Op_Load32Splat_V128 = 0xFD_09
	Op_Load64Splat_V128 = 0xFD_0A
	Op_Load32Zero_V128  = 0xFD_5C
	Op_Load64Zero_V128  = 0xFD_5D
	Op_Store_V128       = 0xFD_0B
	Op_Load8Lane_V128   = 0xFD_54
	Op_Load16Lane_V128  = 0xFD_55
	Op_Load32Lane_V128  = 0xFD_56
	Op_Load64Lane_V128  = 0xFD_57
	Op_Store8Lane_V128  = 0xFD_58
	Op_Store16Lane_V128 = 0xFD_59
	Op_Store32Lane_V128 = 0xFD_5A
	Op_Store64Lane_V128 = 0xFD_5B

	Op_Const_V128 = 0xFD_0C

	Op_Shuffle_I8X16 = 0xFD_0D

	Op_ExtractLaneS_I8X16 = 0xFD_15
	Op_ExtractLaneU_I8X16 = 0xFD_16
	Op_ReplaceLane_I8X16  = 0xFD_17
	Op_ExtractLaneS_I16X8 = 0xFD_18
	Op_ExtractLaneU_I16X8 = 0xFD_19
	Op_ReplaceLane_I16X8  = 0xFD_1A
	Op_ExtractLane_I32X4  = 0xFD_1B
	Op_ReplaceLane_I32X4  = 0xFD_1C
	Op_ExtractLane_I64X2  = 0xFD_1D
	Op_ReplaceLane_I64X2  = 0xFD_1E
	Op_ExtractLane_F32X4  = 0xFD_1F
	Op_ReplaceLane_F32X4  = 0xFD_20
	Op_ExtractLane_F64X2  = 0xFD_21
	Op_ReplaceLane_F64X2  = 0xFD_22

	Op_Swizzle_I8X16 = 0xFD_0E
	Op_Splat_I8X16   = 0xFD_0F
	Op_Splat_I16X8   = 0xFD_10
	Op_Splat_I32X4   = 0xFD_11
	Op_Splat_I64X2   = 0xFD_12
	Op_Splat_F32X4   = 0xFD_13
	Op_Splat_F64X2   = 0xFD_14

	Op_Eq_I8X16  = 0xFD_23
	Op_Ne_I8X16  = 0xFD_24
	Op_LtS_I8X16 = 0xFD_25
	Op_LtU_I8X16 = 0xFD_26
	Op_GtS_I8X16 = 0xFD_27
	Op_GtU_I8X16 = 0xFD_28
	Op_LeS_I8X16 = 0xFD_29
	Op_LeU_I8X16 = 0xFD_2A
	Op_GeS_I8X16 = 0xFD_2B
	Op_GeU_I8X16 = 0xFD_2C

	Op_Eq_I16X8  = 0xFD_2D
	Op_Ne_I16X8  = 0xFD_2E
	Op_LtS_I16X8 = 0xFD_2F
	Op_LtU_I16X8 = 0xFD_30
	Op_GtS_I16X8 = 0xFD_31
	Op_GtU_I16X8 = 0xFD_32
	Op_LeS_I16X8 = 0xFD_33
	Op_LeU_I16X8 = 0xFD_34
	Op_GeS_I16X8 = 0xFD_35
	Op_GeU_I16X8 = 0xFD_36

	Op_Eq_I32X4  = 0xFD_37
	Op_Ne_I32X4  = 0xFD_38
	Op_LtS_I32X4 = 0xFD_39
	Op_LtU_I32X4 = 0xFD_3A
	Op_GtS_I32X4 = 0xFD_3B
	Op_GtU_I32X4 = 0xFD_3C
	Op_LeS_I32X4 = 0xFD_3D
	Op_LeU_I32X4 = 0xFD_3E
	Op_GeS_I32X4 = 0xFD_3F
	Op_GeU_I32X4 = 0xFD_40

	Op_Eq_I64X2  = 0xFD_D6
	Op_Ne_I64X2  = 0xFD_D7
	Op_LtS_I64X2 = 0xFD_D8
	Op_GtS_I64X2 = 0xFD_D9
	Op_LeS_I64X2 = 0xFD_DA
	Op_GeS_I64X2 = 0xFD_DB

	Op_Eq_F32X4 = 0xFD_41
	Op_Ne_F32X4 = 0xFD_42
	Op_Lt_F32X4 = 0xFD_43
	Op_Gt_F32X4 = 0xFD_44
	Op_Le_F32X4 = 0xFD_45
	Op_Ge_F32X4 = 0xFD_46

	Op_Eq_F64X2 = 0xFD_47
	Op_Ne_F64X2 = 0xFD_48
	Op_Lt_F64X2 = 0xFD_49
	Op_Gt_F64X2 = 0xFD_4A
	Op_Le_F64X2 = 0xFD_4B
	Op_Ge_F64X2 = 0xFD_4C

	Op_Not_v128       = 0xFD_4D
	Op_And_v128       = 0xFD_4E
	Op_Andnot_v128    = 0xFD_4F
	Op_Or_v128        = 0xFD_50
	Op_Xor_v128       = 0xFD_51
	Op_Bitselect_v128 = 0xFD_52
	Op_AnyTrue_v128   = 0xFD_53

	Op_Abs_I8X16          = 0xFD_60
	Op_Neg_I8X16          = 0xFD_61
	Op_Popcnt_I8X16       = 0xFD_62
	Op_AllTrue_I8X16      = 0xFD_63
	Op_Bitmask_I8X16      = 0xFD_64
	Op_NarrowI16X8S_I8X16 = 0xFD_65
	Op_NarrowI16X8U_I8X16 = 0xFD_66
	Op_Shl_I8X16          = 0xFD_6B
	Op_ShrS_I8X16         = 0xFD_6C
	Op_ShrU_I8X16         = 0xFD_6D
	Op_Add_I8X16          = 0xFD_6E
	Op_AddSatS_I8X16      = 0xFD_6F
	Op_AddSatU_I8X16      = 0xFD_70
	Op_Sub_I8X16          = 0xFD_71
	Op_SubSatS_I8X16      = 0xFD_72
	Op_SubSatU_I8X16      = 0xFD_73
	Op_MinS_I8X16         = 0xFD_76
	Op_MinU_I8X16         = 0xFD_77
	Op_MaxS_I8X16         = 0xFD_78
	Op_MaxU_I8X16         = 0xFD_79
	Op_AvgrU_I8X16        = 0xFD_7B

	Op_ExtaddPairwiseI8X16S_I16X8 = 0xFD_7C
	Op_ExtaddPairwiseI8X16U_I16X8 = 0xFD_7D
	Op_Abs_I16X8                  = 0xFD_80
	Op_Neg_I16X8                  = 0xFD_81
	Op_Q15mulrSatS_I16X8          = 0xFD_82
	Op_AllTrue_I16X8              = 0xFD_83
	Op_Bitmask_I16X8              = 0xFD_84
	Op_NarrowI32X4S_I16X8         = 0xFD_85
	Op_NarrowI32X4U_I16X8         = 0xFD_86
	Op_ExtendLowI8X16S_I16X8      = 0xFD_87
	Op_ExtendHighI8X16S_I16X8     = 0xFD_88
	Op_ExtendLowI8X16U_I16X8      = 0xFD_89
	Op_ExtendHighI8X16U_I16X8     = 0xFD_8A
	Op_Shl_I16X8                  = 0xFD_8B
	Op_ShrS_I16X8                 = 0xFD_8C
	Op_ShrU_I16X8                 = 0xFD_8D
	Op_Add_I16X8                  = 0xFD_8E
	Op_AddSatS_I16X8              = 0xFD_8F
	Op_AddSatU_I16X8              = 0xFD_90
	Op_Sub_I16X8                  = 0xFD_91
	Op_SubSatS_I16X8              = 0xFD_92
	Op_SubSatU_I16X8              = 0xFD_93
	Op_Mul_I16X8                  = 0xFD_95
	Op_MinS_I16X8                 = 0xFD_96
	Op_MinU_I16X8                 = 0xFD_97
	Op_MaxS_I16X8                 = 0xFD_98
	Op_MaxU_I16X8                 = 0xFD_99
	Op_AvgrU_I16X8                = 0xFD_9B
	Op_ExtmulLowI8X16S_I16X8      = 0xFD_9C
	Op_ExtmulHighI8X16S_I16X8     = 0xFD_9D
	Op_ExtmulLowI8X16U_I16X8      = 0xFD_9E
	Op_ExtmulHighI8X16U_I16X8     = 0xFD_9F

	Op_ExtaddPairwiseI16X8S_I32X4 = 0xFD_7E
	Op_ExtaddPairwiseI16X8U_I32X4 = 0xFD_7F
	Op_Abs_I32X4                  = 0xFD_A0
	Op_Neg_I32X4                  = 0xFD_A1
	Op_AllTrue_I32X4              = 0xFD_A3
	Op_Bitmask_I32X4              = 0xFD_A4
	Op_ExtendLowI16X8S_I32X4      = 0xFD_A7
	Op_ExtendHighI16X8S_I32X4     = 0xFD_A8
	Op_ExtendLowI16X8U_I32X4      = 0xFD_A9
	Op_ExtendHighI16X8U_I32X4     = 0xFD_AA
	Op_Shl_I32X4                  = 0xFD_AB
	Op_ShrS_I32X4                 = 0xFD_AC
	Op_ShrU_I32X4                 = 0xFD_AD
	Op_Add_I32X4                  = 0xFD_AE
	Op_Sub_I32X4                  = 0xFD_B1
	Op_Mul_I32X4                  = 0xFD_B5
	Op_MinS_I32X4                 = 0xFD_B6
	Op_MinU_I32X4                 = 0xFD_B7
	Op_MaxS_I32X4                 = 0xFD_B8
	Op_MaxU_I32X4                 = 0xFD_B9
	Op_DotI16X8S_I32X4            = 0xFD_BA
	Op_ExtmulLowI16X8S_I32X4      = 0xFD_BC
	Op_ExtmulHighI16X8S_I32X4     = 0xFD_BD
	Op_ExtmulLowI16X8U_I32X4      = 0xFD_BE
	Op_ExtmulHighI16X8U_I32X4     = 0xFD_BF

	Op_Abs_I64X2              = 0xFD_C0
	Op_Neg_I64X2              = 0xFD_C1
	Op_AllTrue_I64X2          = 0xFD_C3
	Op_Bitmask_I64X2          = 0xFD_C4
	Op_ExtendLowI32X4S_I64X2  = 0xFD_C7
	Op_ExtendHighI32X4S_I64X2 = 0xFD_C8
	Op_ExtendLowI32X4U_I64X2  = 0xFD_C9
	Op_ExtendHighI32X4U_I64X2 = 0xFD_CA
	Op_Shl_I64X2              = 0xFD_CB
	Op_ShrS_I64X2             = 0xFD_CC
	Op_ShrU_I64X2             = 0xFD_CD
	Op_Add_I64X2              = 0xFD_CE
	Op_Sub_I64X2              = 0xFD_D1
	Op_Mul_I64X2              = 0xFD_D5
	Op_ExtmulLowI32X4S_I64X2  = 0xFD_DC
	Op_ExtmulHighI32X4S_I64X2 = 0xFD_DD
	Op_ExtmulLowI32X4U_I64X2  = 0xFD_DE
	Op_ExtmulHighI32X4U_I64X2 = 0xFD_DF

	Op_Ceil_F32X4    = 0xFD_67
	Op_Floor_F32X4   = 0xFD_68
	Op_Trunc_F32X4   = 0xFD_69
	Op_Nearest_F32X4 = 0xFD_6A
	Op_Abs_F32X4     = 0xFD_E0
	Op_Neg_F32X4     = 0xFD_E1
	Op_Sqrt_F32X4    = 0xFD_E3
	Op_Add_F32X4     = 0xFD_E4
	Op_Sub_F32X4     = 0xFD_E5
	Op_Mul_F32X4     = 0xFD_E6
	Op_Div_F32X4     = 0xFD_E7
	Op_Min_F32X4     = 0xFD_E8
	Op_Max_F32X4     = 0xFD_E9
	Op_Pmin_F32X4    = 0xFD_EA
	Op_Pmax_F32X4    = 0xFD_EB

	Op_Ceil_F64X2    = 0xFD_74
	Op_Floor_F64X2   = 0xFD_75
	Op_Trunc_F64X2   = 0xFD_7A
	Op_Nearest_F64X2 = 0xFD_94
	Op_Abs_F64X2     = 0xFD_EC
	Op_Neg_F64X2     = 0xFD_ED
	Op_Sqrt_F64X2    = 0xFD_EF
	Op_Add_F64X2     = 0xFD_F0
	Op_Sub_F64X2     = 0xFD_F1
	Op_Mul_F64X2     = 0xFD_F2
	Op_Div_F64X2     = 0xFD_F3
	Op_Min_F64X2     = 0xFD_F4
	Op_Max_F64X2     = 0xFD_F5
	Op_Pmin_F64X2    = 0xFD_F6
	Op_Pmax_F64X2    = 0xFD_F7

	Op_TruncSatF32X4S_I32X4      = 0xFD_F8
	Op_TruncSatF32X4U_I32X4      = 0xFD_F9
	Op_ConvertI32X4S_F32X4       = 0xFD_FA
	Op_ConvertI32X4U_F32X4       = 0xFD_FB
	Op_TruncSatF64X2S_Zero_I32X4 = 0xFD_FC
	Op_TruncSatF64X2U_Zero_I32X4 = 0xFD_FD
	Op_ConvertLowI32X4S_F64X2    = 0xFD_FE
	Op_ConvertLowI32X4U_F64X2    = 0xFD_FF
	Op_DemoteF64X2Zero_F32X4     = 0xFD_5E
	Op_PromoteLowF32X4_F64X2     = 0xFD_5F
)
