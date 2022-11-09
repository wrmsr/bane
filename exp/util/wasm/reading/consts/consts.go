package consts

/*
https://en.wikipedia.org/wiki/LEB128
https://en.wikipedia.org/wiki/LEB128#Signed_LEB128
*/

const (
	// numtype

	I32 = 0x7F
	I64 = 0x7E
	F32 = 0x7D
	F64 = 0x7C

	// vectype

	V128 = 0x7B

	// reftype

	FuncRef   = 0x70
	ExternRef = 0x6F

	FuncType = 0x60

	MinLimit = 0x00
	MaxLimit = 0x01

	ConstGlobalType = 0x00
	VarGlobalType   = 0x01

	// control

	BlockType    = 0x40
	Unreachable  = 0x00
	Nop          = 0x01
	Block        = 0x02
	Loop         = 0x03
	If           = 0x04
	Else         = 0x05
	End          = 0x0B
	Br           = 0x0C
	BrIf         = 0x0D
	BrTable      = 0x0E
	Return       = 0x0F
	Call         = 0x10
	CallIndirect = 0x11

	// reference

	RefNull   = 0xD0
	RefIsNull = 0xD1
	RefFunc   = 0xD2

	// parametric

	Drop       = 0x1A
	Select     = 0x1B
	SelectType = 0x1C

	// variable

	LocalGet  = 0x20
	LocalSet  = 0x21
	LocalTee  = 0x22
	GlobalGet = 0x23
	GlobalSet = 0x24

	// table

	TableGet  = 0x25
	TableSet  = 0x26
	TableInit = 0xFC_0C
	ElemDrop  = 0xFC_0D
	TableCopy = 0xFC_0E
	TableGrow = 0xFC_0F
	TableSize = 0xFC_10
	TableFill = 0xFC_11

	// memory

	Load_I32    = 0x28
	Load_I64    = 0x29
	Load_F32    = 0x2A
	Load_F64    = 0x2B
	Load8S_I32  = 0x2C
	Load8U_I32  = 0x2D
	Load16S_I32 = 0x2E
	Load16U_I32 = 0x2F
	Load8S_I64  = 0x30
	Load8U_I64  = 0x31
	Load16S_I64 = 0x32
	Load16U_I64 = 0x33
	Load32S_I64 = 0x34
	Load32U_I64 = 0x35
	Store_I32   = 0x36
	Store_I64   = 0x37
	Store_F32   = 0x38
	Store_F64   = 0x39
	Store8_I32  = 0x3A
	Store16_I32 = 0x3B
	Store8_I64  = 0x3C
	Store16_I64 = 0x3D
	Store32_I64 = 0x3E
	MemorySize  = 0x3F_00
	MemoryGrow  = 0x40_00
	MemoryInit  = 0xFC_08
	DataDrop    = 0xFC_09
	MemoryCopy  = 0xFC_0A
	MemoryFill  = 0xFC_0B

	// numeric

	Const_I32 = 0x41
	Const_I64 = 0x42
	Const_F32 = 0x43
	Const_F64 = 0x44

	Eqz_I32 = 0x45
	Eq_I32  = 0x46
	Ne_I32  = 0x47
	LtS_I32 = 0x48
	LtU_I32 = 0x49
	GtS_I32 = 0x4A
	GtU_I32 = 0x4B
	LeS_I32 = 0x4C
	LeU_I32 = 0x4D
	GeS_I32 = 0x4E
	GeU_I32 = 0x4F

	Eqz_I64 = 0x50
	Eq_I64  = 0x51
	Ne_I64  = 0x52
	LtS_I64 = 0x53
	LtU_I64 = 0x54
	GtS_I64 = 0x55
	GtU_I64 = 0x56
	LeS_I64 = 0x57
	LeU_I64 = 0x58
	GeS_I64 = 0x59
	GeU_I64 = 0x5A

	Eq_F32 = 0x5B
	Ne_F32 = 0x5C
	Lt_F32 = 0x5D
	Gt_F32 = 0x5E
	Le_F32 = 0x5F
	Ge_F32 = 0x60

	Eq_F64 = 0x61
	Ne_F64 = 0x62
	Lt_F64 = 0x63
	Gt_F64 = 0x64
	Le_F64 = 0x65
	Ge_F64 = 0x66

	Clz_I32    = 0x67
	Ctz_I32    = 0x68
	Popcnt_I32 = 0x69
	Add_I32    = 0x6A
	Sub_I32    = 0x6B
	Mul_I32    = 0x6C
	DivS_I32   = 0x6D
	DivU_I32   = 0x6E
	RemS_I32   = 0x6F
	RemU_I32   = 0x70
	And_I32    = 0x71
	Or_I32     = 0x72
	Xor_I32    = 0x73
	Shl_I32    = 0x74
	ShrS_I32   = 0x75
	ShrU_I32   = 0x76
	Rotl_I32   = 0x77
	Rotr_I32   = 0x78
	Clz_I64    = 0x79
	Ctz_I64    = 0x7A
	Popcnt_I64 = 0x7B
	Add_I64    = 0x7C
	Sub_I64    = 0x7D
	Mul_I64    = 0x7E
	DivS_I64   = 0x7F
	DivU_I64   = 0x80
	RemS_I64   = 0x81
	RemU_I64   = 0x82
	And_I64    = 0x83
	Or_I64     = 0x84
	Xor_I64    = 0x85
	Shl_I64    = 0x86
	ShrS_I64   = 0x87
	ShrU_I64   = 0x88
	Rotl_I64   = 0x89
	Rotr_I64   = 0x8A

	Abs_F32      = 0x8B
	Neg_F32      = 0x8C
	Ceil_F32     = 0x8D
	Floor_F32    = 0x8E
	Trunc_F32    = 0x8F
	Nearest_F32  = 0x90
	Sqrt_F32     = 0x91
	Add_F32      = 0x92
	Sub_F32      = 0x93
	Mul_F32      = 0x94
	Div_F32      = 0x95
	Min_F32      = 0x96
	Max_F32      = 0x97
	Copysign_F32 = 0x98

	Abs_F64      = 0x99
	Neg_F64      = 0x9A
	Ceil_F64     = 0x9B
	Floor_F64    = 0x9C
	Trunc_F64    = 0x9D
	Nearest_F64  = 0x9E
	Sqrt_F64     = 0x9F
	Add_F64      = 0xA0
	Sub_F64      = 0xA1
	Mul_F64      = 0xA2
	Div_F64      = 0xA3
	Min_F64      = 0xA4
	Max_F64      = 0xA5
	Copysign_F64 = 0xA6

	WrapI64_I32        = 0xA7
	TruncF32S_I32      = 0xA8
	TruncF32U_I32      = 0xA9
	TruncF64S_I32      = 0xAA
	TruncF64U_I32      = 0xAB
	ExtendI32S_I64     = 0xAC
	ExtendI32U_I64     = 0xAD
	TruncF32S_I64      = 0xAE
	TruncF32U_I64      = 0xAF
	TruncF64S_I64      = 0xB0
	TruncF64U_I64      = 0xB1
	ConvertI32S_F32    = 0xB2
	ConvertI32U_F32    = 0xB3
	ConvertI64S_F32    = 0xB4
	ConvertI64U_F32    = 0xB5
	DemoteF64_F32      = 0xB6
	ConvertI32S_F64    = 0xB7
	ConvertI32U_F64    = 0xB8
	ConvertI64S_F64    = 0xB9
	ConvertI64U_F64    = 0xBA
	PromoteF32_F64     = 0xBB
	ReinterpretF32_I32 = 0xBC
	ReinterpretF64_I64 = 0xBD
	ReinterpretI32_F32 = 0xBE
	ReinterpretI64_F64 = 0xBF

	Extend8S_I32  = 0xC0
	Extend16S_I32 = 0xC1
	Extend8S_I64  = 0xC2
	Extend16S_I64 = 0xC3
	Extend32S_I64 = 0xC4

	TruncSatF32S_I32 = 0xFC_00
	TruncSatF32U_I32 = 0xFC_01
	TruncSatF64S_I32 = 0xFC_02
	TruncSatF64U_I32 = 0xFC_03
	TruncSatF32S_I64 = 0xFC_04
	TruncSatF32U_I64 = 0xFC_05
	TruncSatF64S_I64 = 0xFC_06
	TruncSatF64U_I64 = 0xFC_07

	// vector

	Load_V128        = 0xFD_00
	Load8x8S_V128    = 0xFD_01
	Load8x8U_V128    = 0xFD_02
	Load16x4S_V128   = 0xFD_03
	Load16x4U_V128   = 0xFD_04
	Load32x2S_V128   = 0xFD_05
	Load32x2U_V128   = 0xFD_06
	Load8Splat_V128  = 0xFD_07
	Load16Splat_V128 = 0xFD_08
	Load32Splat_V128 = 0xFD_09
	Load64Splat_V128 = 0xFD_0A
	Load32Zero_V128  = 0xFD_5C
	Load64Zero_V128  = 0xFD_5D
	Store_V128       = 0xFD_0B
	Load8Lane_V128   = 0xFD_54
	Load16Lane_V128  = 0xFD_55
	Load32Lane_V128  = 0xFD_56
	Load64Lane_V128  = 0xFD_67
	Store8Lane_V128  = 0xFD_58
	Store16Lane_V128 = 0xFD_59
	Store32Lane_V128 = 0xFD_5A
	Store64Lane_V128 = 0xFD_5B

	Const_V128 = 0xFD_0C

	Shuffle_I8X16 = 0xFD_0D

	ExtractLaneS_I8X16 = 0xFD_15
	ExtractLaneU_I8X16 = 0xFD_16
	ReplaceLane_I8X16  = 0xFD_17
	ExtractLaneS_I16X8 = 0xFD_18
	ExtractLaneU_I16X8 = 0xFD_19
	ReplaceLane_I16X8  = 0xFD_1A
	ExtractLane_I32X4  = 0xFD_1B
	ReplaceLane_I32X4  = 0xFD_1C
	ExtractLane_I64X2  = 0xFD_1D
	ReplaceLane_I64X2  = 0xFD_1E
	ExtractLane_F32X4  = 0xFD_1F
	ReplaceLane_F32X4  = 0xFD_20
	ExtractLane_F64X2  = 0xFD_21
	ReplaceLane_F64X2  = 0xFD_22

	Swizzle_I8X16 = 0xFD_0E
	Splat_I8X16   = 0xFD_0F
	Splat_I16X8   = 0xFD_10
	Splat_I32X4   = 0xFD_11
	Splat_I64X2   = 0xFD_12
	Splat_F32X4   = 0xFD_13
	Splat_F64X2   = 0xFD_14

	Eq_I8X16  = 0xFD_23
	Ne_I8X16  = 0xFD_24
	LtS_I8X16 = 0xFD_25
	LtU_I8X16 = 0xFD_26
	GtS_I8X16 = 0xFD_27
	GtU_I8X16 = 0xFD_28
	LeS_I8X16 = 0xFD_29
	LeU_I8X16 = 0xFD_2A
	GeS_I8X16 = 0xFD_2B
	GeU_I8X16 = 0xFD_2C

	Eq_I16X8  = 0xFD_2D
	Ne_I16X8  = 0xFD_2E
	LtS_I16X8 = 0xFD_2F
	LtU_I16X8 = 0xFD_30
	GtS_I16X8 = 0xFD_31
	GtU_I16X8 = 0xFD_32
	LeS_I16X8 = 0xFD_33
	LeU_I16X8 = 0xFD_34
	GeS_I16X8 = 0xFD_35
	GeU_I16X8 = 0xFD_36

	Eq_I32X4  = 0xFD_37
	Ne_I32X4  = 0xFD_38
	LtS_I32X4 = 0xFD_39
	LtU_I32X4 = 0xFD_3A
	GtS_I32X4 = 0xFD_3B
	GtU_I32X4 = 0xFD_3C
	LeS_I32X4 = 0xFD_3D
	LeU_I32X4 = 0xFD_3E
	GeS_I32X4 = 0xFD_3F
	GeU_I32X4 = 0xFD_40

	Eq_I64X2  = 0xFD_D6
	Ne_I64X2  = 0xFD_D7
	LtS_I64X2 = 0xFD_D8
	GtS_I64X2 = 0xFD_D9
	LeS_I64X2 = 0xFD_DA
	GeS_I64X2 = 0xFD_DB

	Eq_F32X4 = 0xFD_41
	Ne_F32X4 = 0xFD_42
	Lt_F32X4 = 0xFD_43
	Gt_F32X4 = 0xFD_44
	Le_F32X4 = 0xFD_45
	Ge_F32X4 = 0xFD_46

	Eq_F64X2 = 0xFD_47
	Ne_F64X2 = 0xFD_48
	Lt_F64X2 = 0xFD_49
	Gt_F64X2 = 0xFD_4A
	Le_F64X2 = 0xFD_4B
	Ge_F64X2 = 0xFD_4C

	Not_v128       = 0xFD_4D
	And_v128       = 0xFD_4E
	Andnot_v128    = 0xFD_4F
	Or_v128        = 0xFD_50
	Xor_v128       = 0xFD_51
	Bitselect_v128 = 0xFD_52
	AnyTrue_v128   = 0xFD_53

	Abs_I8X16          = 0xFD_60
	Neg_I8X16          = 0xFD_61
	Popcnt_I8X16       = 0xFD_62
	AllTrue_I8X16      = 0xFD_63
	Bitmask_I8X16      = 0xFD_64
	NarrowI16X8S_I8X16 = 0xFD_65
	NarrowI16X8U_I8X16 = 0xFD_66
	Shl_I8X16          = 0xFD_6B
	ShrS_I8X16         = 0xFD_6C
	ShrU_I8X16         = 0xFD_6D
	Add_I8X16          = 0xFD_6E
	AddSatS_I8X16      = 0xFD_6F
	AddSatU_I8X16      = 0xFD_70
	Sub_I8X16          = 0xFD_71
	SubSatS_I8X16      = 0xFD_72
	SubSatU_I8X16      = 0xFD_73
	MinS_I8X16         = 0xFD_76
	MinU_I8X16         = 0xFD_77
	MaxS_I8X16         = 0xFD_78
	MaxU_I8X16         = 0xFD_79
	AvgrU_I8X16        = 0xFD_7B

	ExtaddPairwiseI8X16S_I16X8 = 0xFD_7C
	ExtaddPairwiseI8X16U_I16X8 = 0xFD_7D
	Abs_I16X8                  = 0xFD_80
	Neg_I16X8                  = 0xFD_81
	Q15mulrSatS_I16X8          = 0xFD_82
	AllTrue_I16X8              = 0xFD_83
	Bitmask_I16X8              = 0xFD_84
	NarrowI32X4S_I16X8         = 0xFD_85
	NarrowI32X4U_I16X8         = 0xFD_86
	ExtendLowI8X16S_I16X8      = 0xFD_87
	ExtendHighI8X16S_I16X8     = 0xFD_88
	ExtendLowI8X16U_I16X8      = 0xFD_89
	ExtendHighI8X16U_I16X8     = 0xFD_8A
	Shl_I16X8                  = 0xFD_8B
	ShrS_I16X8                 = 0xFD_8C
	ShrU_I16X8                 = 0xFD_8D
	Add_I16X8                  = 0xFD_8E
	AddSatS_I16X8              = 0xFD_8F
	AddSatU_I16X8              = 0xFD_90
	Sub_I16X8                  = 0xFD_91
	SubSatS_I16X8              = 0xFD_92
	SubSatU_I16X8              = 0xFD_93
	Mul_I16X8                  = 0xFD_95
	MinS_I16X8                 = 0xFD_96
	MinU_I16X8                 = 0xFD_97
	MaxS_I16X8                 = 0xFD_98
	MaxU_I16X8                 = 0xFD_99
	AvgrU_I16X8                = 0xFD_9B
	ExtmulLowI8X16S_I16X8      = 0xFD_9C
	ExtmulHighI8X16S_I16X8     = 0xFD_9D
	ExtmulLowI8X16U_I16X8      = 0xFD_9E
	ExtmulHighI8X16U_I16X8     = 0xFD_9F

	ExtaddPairwiseI16X8S_I32X4 = 0xFD_7E
	ExtaddPairwiseI16X8U_I32X4 = 0xFD_7F
	Abs_I32X4                  = 0xFD_A0
	Neg_I32X4                  = 0xFD_A1
	AllTrue_I32X4              = 0xFD_A3
	Bitmask_I32X4              = 0xFD_A4
	ExtendLowI16X8S_I32X4      = 0xFD_A7
	ExtendHighI16X8S_I32X4     = 0xFD_A8
	ExtendLowI16X8U_I32X4      = 0xFD_A9
	ExtendHighI16X8U_I32X4     = 0xFD_AA
	Shl_I32X4                  = 0xFD_AB
	ShrS_I32X4                 = 0xFD_AC
	ShrU_I32X4                 = 0xFD_AD
	Add_I32X4                  = 0xFD_AE
	Sub_I32X4                  = 0xFD_B1
	Mul_I32X4                  = 0xFD_B5
	MinS_I32X4                 = 0xFD_B6
	MinU_I32X4                 = 0xFD_B7
	MaxS_I32X4                 = 0xFD_B8
	MaxU_I32X4                 = 0xFD_B9
	DotI16X8S_I32X4            = 0xFD_BA
	ExtmulLowI16X8S_I32X4      = 0xFD_BC
	ExtmulHighI16X8S_I32X4     = 0xFD_BD
	ExtmulLowI16X8U_I32X4      = 0xFD_BE
	ExtmulHighI16X8U_I32X4     = 0xFD_BF

	Abs_I64X2              = 0xFD_C0
	Neg_I64X2              = 0xFD_C1
	AllTrue_I64X2          = 0xFD_C3
	Bitmask_I64X2          = 0xFD_C4
	ExtendLowI32X4S_I64X2  = 0xFD_C7
	ExtendHighI32X4S_I64X2 = 0xFD_C8
	ExtendLowI32X4U_I64X2  = 0xFD_C9
	ExtendHighI32X4U_I64X2 = 0xFD_CA
	Shl_I64X2              = 0xFD_CB
	ShrS_I64X2             = 0xFD_CC
	ShrU_I64X2             = 0xFD_CD
	Add_I64X2              = 0xFD_CE
	Sub_I64X2              = 0xFD_D1
	Mul_I64X2              = 0xFD_D5
	ExtmulLowI32X4S_I64X2  = 0xFD_DC
	ExtmulHighI32X4S_I64X2 = 0xFD_DD
	ExtmulLowI32X4U_I64X2  = 0xFD_DE
	ExtmulHighI32X4U_I64X2 = 0xFD_DF

	Ceil_F32X4    = 0xFD_67
	Floor_F32X4   = 0xFD_68
	Trunc_F32X4   = 0xFD_69
	Nearest_F32X4 = 0xFD_6A
	Abs_F32X4     = 0xFD_E0
	Neg_F32X4     = 0xFD_E1
	Sqrt_F32X4    = 0xFD_E3
	Add_F32X4     = 0xFD_E4
	Sub_F32X4     = 0xFD_E5
	Mul_F32X4     = 0xFD_E6
	Div_F32X4     = 0xFD_E7
	Min_F32X4     = 0xFD_E8
	Max_F32X4     = 0xFD_E9
	Pmin_F32X4    = 0xFD_EA
	Pmax_F32X4    = 0xFD_EB

	Ceil_F64X2    = 0xFD_74
	Floor_F64X2   = 0xFD_75
	Trunc_F64X2   = 0xFD_7A
	Nearest_F64X2 = 0xFD_94
	Abs_F64X2     = 0xFD_EC
	Neg_F64X2     = 0xFD_ED
	Sqrt_F64X2    = 0xFD_EF
	Add_F64X2     = 0xFD_F0
	Sub_F64X2     = 0xFD_F1
	Mul_F64X2     = 0xFD_F2
	Div_F64X2     = 0xFD_F3
	Min_F64X2     = 0xFD_F4
	Max_F64X2     = 0xFD_F5
	Pmin_F64X2    = 0xFD_F6
	Pmax_F64X2    = 0xFD_F7

	TruncSatF32X4S_I32X4      = 0xFD_F8
	TruncSatF32X4U_I32X4      = 0xFD_F9
	ConvertI32X4S_F32X4       = 0xFD_FA
	ConvertI32X4U_F32X4       = 0xFD_FB
	TruncSatF64X2S_Zero_I32X4 = 0xFD_FC
	TruncSatF64X2U_Zero_I32X4 = 0xFD_FD
	ConvertLowI32X4S_F64X2    = 0xFD_FE
	ConvertLowI32X4U_F64X2    = 0xFD_FF
	DemoteF64X2Zero_F32X4     = 0xFD_5E
	PromoteLowF32X4_F64X2     = 0xFD_5F

	CustomSection    = 0
	TypeSection      = 1
	ImportSection    = 2
	FunctionSection  = 3
	TableSection     = 4
	MemorySection    = 5
	GlobalSection    = 6
	ExportSection    = 7
	StartSection     = 8
	ElementSection   = 9
	CodeSection      = 10
	DataSection      = 11
	DataCountSection = 12

	FuncImport   = 0x00
	TableImport  = 0x01
	MemImport    = 0x02
	GlobalImport = 0x03

	FuncExport   = 0x00
	TableExport  = 0x01
	MemExport    = 0x01
	GlobalExport = 0x03

	/*

	   elemsec ::= seg* :section9(vec(elem)) ⇒ seg*
	   elem ::=
	   0:u32 𝑒:expr 𝑦 * :vec(funcidx) ⇒ {type funcref, init ((ref.func 𝑦) end) * , mode active {table 0, offset 𝑒}}
	   | 1:u32 et : elemkind 𝑦 * :vec(funcidx) ⇒ {type et, init ((ref.func 𝑦) end) * , mode passive}
	   | 2:u32 𝑥:tableidx 𝑒:expr et : elemkind 𝑦 * :vec(funcidx) ⇒ {type et, init ((ref.func 𝑦) end) * , mode active {table 𝑥, offset 𝑒}}
	   | 3:u32 et : elemkind 𝑦 * :vec(funcidx) ⇒ {type et, init ((ref.func 𝑦) end) * , mode declarative}
	   | 4:u32 𝑒:expr el * :vec(expr) ⇒ {type funcref, init el * , mode active {table 0, offset 𝑒}}
	   | 5:u32 et : reftype el * :vec(expr) ⇒ {type 𝑒𝑡, init el * , mode passive}
	   | 6:u32 𝑥:tableidx 𝑒:expr et : reftype el * :vec(expr) ⇒ {type 𝑒𝑡, init el * , mode active {table 𝑥, offset 𝑒}}
	   | 7:u32 et : reftype el * :vec(expr) ⇒ {type 𝑒𝑡, init el * , mode declarative}
	   elemkind ::= 0x00 ⇒ funcref

	   codesec ::= code* :section10(vec(code)) ⇒ code*
	   code ::= size:u32 code:func ⇒ code (if size = ||func||)
	   func ::= (𝑡 * ) * :vec(locals) 𝑒:expr ⇒ concat((𝑡 * ) * ), 𝑒 (if |concat((𝑡 * ) * )| < 2 32)
	   locals ::= 𝑛:u32 𝑡:valtype ⇒ 𝑡 𝑛

	   datasec ::= seg* :section11(vec(data)) ⇒ seg*
	   data ::=
	   0:u32 𝑒:expr 𝑏 * :vec(byte) ⇒ {init 𝑏 * , mode active {memory 0, offset 𝑒}}
	   | 1:u32 𝑏 * :vec(byte) ⇒ {init 𝑏 * , mode passive}
	   | 2:u32 𝑥:memidx 𝑒:expr 𝑏 * :vec(byte) ⇒ {init 𝑏 * , mode active {memory 𝑥, offset 𝑒}}

	   datacountsec ::= n ? :section12(u32) ⇒ n ?

	   magic ::= 0x00 0x61 0x73 0x6D
	   version ::= 0x01 0x00 0x00 0x00
	   module ::= magic
	   version
	   customsec*
	   functype* : typesec
	   customsec*
	   import* : importsec
	   customsec*
	   typeidx 𝑛: funcsec
	   customsec*
	   table* : tablesec
	   customsec*
	   mem* : memsec
	   customsec*
	   global * : globalsec
	   customsec*
	   export* : exportsec
	   customsec*
	   start ? : startsec
	   customsec*
	   elem* : elemsec
	   customsec*
	   𝑚? : datacountsec
	   customsec*
	   code𝑛 : codesec
	   customsec*
	   data𝑚: datasec
	   customsec* ⇒ { types functype* ,
	   	funcs func𝑛,
	   	tables table* ,
	   	mems mem* ,
	   	globals global * ,
	   	elems elem* ,
	   	datas data𝑚,
	   	start start ? ,
	   	imports import* ,
	   	exports export* }
	   (if 𝑚? ̸= 𝜖 ∨ dataidx(code𝑛 ) = ∅)
	   func𝑛 [𝑖] = {type typeidx 𝑛 [𝑖], locals 𝑡 * 𝑖 , body 𝑒𝑖}

	*/

)
