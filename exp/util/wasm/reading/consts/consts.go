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

	X = 0xFC

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

	Const_I32          = 0x41
	Const_I64          = 0x42
	Const_F32          = 0x43
	Const_F64          = 0x44
	Eqz_I32            = 0x45
	Eq_I32             = 0x46
	Ne_I32             = 0x47
	LtS_I32            = 0x48
	LtU_I32            = 0x49
	GtS_I32            = 0x4A
	GtU_I32            = 0x4B
	LeS_I32            = 0x4C
	LeU_I32            = 0x4D
	GeS_I32            = 0x4E
	GeU_I32            = 0x4F
	Eqz_I64            = 0x50
	Eq_I64             = 0x51
	Ne_I64             = 0x52
	LtS_I64            = 0x53
	LtU_I64            = 0x54
	GtS_I64            = 0x55
	GtU_I64            = 0x56
	LeS_I64            = 0x57
	LeU_I64            = 0x58
	GeS_I64            = 0x59
	GeU_I64            = 0x5A
	Eq_F32             = 0x5B
	Ne_F32             = 0x5C
	Lt_F32             = 0x5D
	Gt_F32             = 0x5E
	Le_F32             = 0x5F
	Ge_F32             = 0x60
	Eq_F64             = 0x61
	Ne_F64             = 0x62
	Lt_F64             = 0x63
	Gt_F64             = 0x64
	Le_F64             = 0x65
	Ge_F64             = 0x66
	Clz_I32            = 0x67
	Ctz_I32            = 0x68
	Popcnt_I32         = 0x69
	Add_I32            = 0x6A
	Sub_I32            = 0x6B
	Mul_I32            = 0x6C
	DivS_I32           = 0x6D
	DivU_I32           = 0x6E
	RemS_I32           = 0x6F
	RemU_I32           = 0x70
	And_I32            = 0x71
	Or_I32             = 0x72
	Xor_I32            = 0x73
	Shl_I32            = 0x74
	ShrS_I32           = 0x75
	ShrU_I32           = 0x76
	Rotl_I32           = 0x77
	Rotr_I32           = 0x78
	Clz_I64            = 0x79
	Ctz_I64            = 0x7A
	Popcnt_I64         = 0x7B
	Add_I64            = 0x7C
	Sub_I64            = 0x7D
	Mul_I64            = 0x7E
	DivS_I64           = 0x7F
	DivU_I64           = 0x80
	RemS_I64           = 0x81
	RemU_I64           = 0x82
	And_I64            = 0x83
	Or_I64             = 0x84
	Xor_I64            = 0x85
	Shl_I64            = 0x86
	ShrS_I64           = 0x87
	ShrU_I64           = 0x88
	Rotl_I64           = 0x89
	Rotr_I64           = 0x8A
	Abs_F32            = 0x8B
	Neg_F32            = 0x8C
	Ceil_F32           = 0x8D
	Floor_F32          = 0x8E
	Trunc_F32          = 0x8F
	Nearest_F32        = 0x90
	Sqrt_F32           = 0x91
	Add_F32            = 0x92
	Sub_F32            = 0x93
	Mul_F32            = 0x94
	Div_F32            = 0x95
	Min_F32            = 0x96
	Max_F32            = 0x97
	Copysign_F32       = 0x98
	Abs_F64            = 0x99
	Neg_F64            = 0x9A
	Ceil_F64           = 0x9B
	Floor_F64          = 0x9C
	Trunc_F64          = 0x9D
	Nearest_F64        = 0x9E
	Sqrt_F64           = 0x9F
	Add_F64            = 0xA0
	Sub_F64            = 0xA1
	Mul_F64            = 0xA2
	Div_F64            = 0xA3
	Min_F64            = 0xA4
	Max_F64            = 0xA5
	Copysign_F64       = 0xA6
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
	Extend8S_I32       = 0xC0
	Extend16S_I32      = 0xC1
	Extend8S_I64       = 0xC2
	Extend16S_I64      = 0xC3
	Extend32S_I64      = 0xC4

	TruncSatF32S_I32 = 0xFC_00
	TruncSatF32U_I32 = 0xFC_01
	TruncSatF64S_I32 = 0xFC_02
	TruncSatF64U_I32 = 0xFC_03
	TruncSatF32S_I64 = 0xFC_04
	TruncSatF32U_I64 = 0xFC_05
	TruncSatF64S_I64 = 0xFC_06
	TruncSatF64U_I64 = 0xFC_07

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
	Eq_I8X16      = 0xFD_23
	Ne_I8X16      = 0xFD_24
	LtS_I8X16     = 0xFD_25
	LtU_I8X16     = 0xFD_26
	GtS_I8X16     = 0xFD_27
	GtU_I8X16     = 0xFD_28
	LeS_I8X16     = 0xFD_29
	LeU_I8X16     = 0xFD_2A
	GeS_I8X16     = 0xFD_2B
	GeU_I8X16     = 0xFD_2C
	Eq_I16X8      = 0xFD_2D
	Ne_I16X8      = 0xFD_2E
	LtS_I16X8     = 0xFD_2F
	LtU_I16X8     = 0xFD_30

	// FIXME:
	GtS_I16X8                   = 0xFD_49
	GtU_I16X8                   = 0xFD_50
	LeS_I16X8                   = 0xFD_51
	LeU_I16X8                   = 0xFD_52
	GeS_I16X8                   = 0xFD_53
	GeU_I16X8                   = 0xFD_54
	Eq_I32X4                    = 0xFD_55
	Ne_I32X4                    = 0xFD_56
	LtS_I32X4                   = 0xFD_57
	LtU_I32X4                   = 0xFD_58
	GtS_I32X4                   = 0xFD_59
	GtU_I32X4                   = 0xFD_60
	LeS_I32X4                   = 0xFD_61
	LeU_I32X4                   = 0xFD_62
	GeS_I32X4                   = 0xFD_63
	GeU_I32X4                   = 0xFD_64
	Eq_I64X2                    = 0xFD_214
	Ne_I64X2                    = 0xFD_215
	LtS_I64X2                   = 0xFD_216
	GtS_I64X2                   = 0xFD_217
	LeS_I64X2                   = 0xFD_218
	GeS_I64X2                   = 0xFD_219
	Eq_F32X4                    = 0xFD_65
	Ne_F32X4                    = 0xFD_66
	Lt_F32X4                    = 0xFD_67
	Gt_F32X4                    = 0xFD_68
	Le_F32X4                    = 0xFD_69
	Ge_F32X4                    = 0xFD_70
	Eq_F64X2                    = 0xFD_71
	Ne_F64X2                    = 0xFD_72
	Lt_F64X2                    = 0xFD_73
	Gt_F64X2                    = 0xFD_74
	Le_F64X2                    = 0xFD_75
	Ge_F64X2                    = 0xFD_76
	Not_v128                    = 0xFD_77
	And_v128                    = 0xFD_78
	Andnot_v128                 = 0xFD_79
	Or_v128                     = 0xFD_80
	Xor_v128                    = 0xFD_81
	Bitselect_v128              = 0xFD_82
	AnyTrue_v128                = 0xFD_83
	Abs_I8X16                   = 0xFD_96
	Neg_I8X16                   = 0xFD_97
	Popcnt_I8X16                = 0xFD_98
	AllTrue_I8X16               = 0xFD_99
	Bitmask_I8X16               = 0xFD_100
	Narrow_I16X8S_I8X16         = 0xFD_101
	Narrow_I16X8U_I8X16         = 0xFD_102
	Shl_I8X16                   = 0xFD_107
	ShrS_I8X16                  = 0xFD_108
	ShrU_I8X16                  = 0xFD_109
	Add_I8X16                   = 0xFD_110
	AddSatS_I8X16               = 0xFD_111
	AddSatU_I8X16               = 0xFD_112
	Sub_I8X16                   = 0xFD_113
	SubSatS_I8X16               = 0xFD_114
	SubSatU_I8X16               = 0xFD_115
	MinS_I8X16                  = 0xFD_118
	MinU_I8X16                  = 0xFD_119
	MaxS_I8X16                  = 0xFD_120
	MaxU_I8X16                  = 0xFD_121
	AvgrU_I8X16                 = 0xFD_123
	ExtaddPairwise_I8X16S_I16X8 = 0xFD_124
	ExtaddPairwise_I8X16U_I16X8 = 0xFD_125
	Abs_I16X8                   = 0xFD_128
	Neg_I16X8                   = 0xFD_129
	Q15mulrSatS_I16X8           = 0xFD_130
	AllTrue_I16X8               = 0xFD_131
	Bitmask_I16X8               = 0xFD_132
	Narrow_I32X4S_I16X8         = 0xFD_133
	Narrow_I32X4U_I16X8         = 0xFD_134
	ExtendLow_I8X16S_I16X8      = 0xFD_135
	ExtendHigh_I8X16S_I16X8     = 0xFD_136
	ExtendLow_I8X16U_I16X8      = 0xFD_137
	ExtendHigh_I8X16U_I16X8     = 0xFD_138
	Shl_I16X8                   = 0xFD_139
	ShrS_I16X8                  = 0xFD_140
	ShrU_I16X8                  = 0xFD_141
	Add_I16X8                   = 0xFD_142
	AddSatS_I16X8               = 0xFD_143
	AddSatU_I16X8               = 0xFD_144
	Sub_I16X8                   = 0xFD_145
	SubSatS_I16X8               = 0xFD_146
	SubSatU_I16X8               = 0xFD_147
	Mul_I16X8                   = 0xFD_149
	MinS_I16X8                  = 0xFD_150
	MinU_I16X8                  = 0xFD_151
	MaxS_I16X8                  = 0xFD_152
	MaxU_I16X8                  = 0xFD_153
	AvgrU_I16X8                 = 0xFD_155
	ExtmulLow_I8X16S_I16X8      = 0xFD_156
	ExtmulHigh_I8X16S_I16X8     = 0xFD_157
	ExtmulLow_I8X16U_I16X8      = 0xFD_158
	ExtmulHigh_I8X16U_I16X8     = 0xFD_159
	ExtaddPairwise_I16X8S_I32X4 = 0xFD_126
	ExtaddPairwise_I16X8U_I32X4 = 0xFD_127
	Abs_I32X4                   = 0xFD_160
	Neg_I32X4                   = 0xFD_161
	AllTrue_I32X4               = 0xFD_163
	Bitmask_I32X4               = 0xFD_164
	ExtendLow_I16X8S_I32X4      = 0xFD_167
	ExtendHigh_I16X8S_I32X4     = 0xFD_168
	ExtendLow_I16X8U_I32X4      = 0xFD_169
	ExtendHigh_I16X8U_I32X4     = 0xFD_170
	Shl_I32X4                   = 0xFD_171
	ShrS_I32X4                  = 0xFD_172
	ShrU_I32X4                  = 0xFD_173
	Add_I32X4                   = 0xFD_174
	Sub_I32X4                   = 0xFD_177
	Mul_I32X4                   = 0xFD_181
	MinS_I32X4                  = 0xFD_182
	MinU_I32X4                  = 0xFD_183
	MaxS_I32X4                  = 0xFD_184
	MaxU_I32X4                  = 0xFD_185
	Dot_I16X8S_I32X4            = 0xFD_186
	ExtmulLow_I16X8S_I32X4      = 0xFD_188
	ExtmulHigh_I16X8S_I32X4     = 0xFD_189
	ExtmulLow_I16X8U_I32X4      = 0xFD_190
	ExtmulHigh_I16X8U_I32X4     = 0xFD_191
	Abs_I64X2                   = 0xFD_192
	Neg_I64X2                   = 0xFD_193
	AllTrue_I64X2               = 0xFD_195
	Bitmask_I64X2               = 0xFD_196
	ExtendLow_I32X4S_I64X2      = 0xFD_199
	ExtendHigh_I32X4S_I64X2     = 0xFD_200
	ExtendLow_I32X4U_I64X2      = 0xFD_201
	ExtendHigh_I32X4U_I64X2     = 0xFD_202
	Shl_I64X2                   = 0xFD_203
	ShrS_I64X2                  = 0xFD_204
	ShrU_I64X2                  = 0xFD_205
	Add_I64X2                   = 0xFD_206
	Sub_I64X2                   = 0xFD_209
	Mul_I64X2                   = 0xFD_213
	ExtmulLow_I32X4S_I64X2      = 0xFD_220
	ExtmulHigh_I32X4S_I64X2     = 0xFD_221
	ExtmulLow_I32X4U_I64X2      = 0xFD_222
	ExtmulHigh_I32X4U_I64X2     = 0xFD_223
	Ceil_F32X4                  = 0xFD_103
	Floor_F32X4                 = 0xFD_104
	Trunc_F32X4                 = 0xFD_105
	Nearest_F32X4               = 0xFD_106
	Abs_F32X4                   = 0xFD_224
	Neg_F32X4                   = 0xFD_225
	Sqrt_F32X4                  = 0xFD_227
	Add_F32X4                   = 0xFD_228
	Sub_F32X4                   = 0xFD_229
	Mul_F32X4                   = 0xFD_230
	Div_F32X4                   = 0xFD_231
	Min_F32X4                   = 0xFD_232
	Max_F32X4                   = 0xFD_233
	Pmin_F32X4                  = 0xFD_234
	Pmax_F32X4                  = 0xFD_235
	Ceil_F64X2                  = 0xFD_116
	Floor_F64X2                 = 0xFD_117
	Trunc_F64X2                 = 0xFD_122
	Nearest_F64X2               = 0xFD_148
	Abs_F64X2                   = 0xFD_236
	Neg_F64X2                   = 0xFD_237
	Sqrt_F64X2                  = 0xFD_239
	Add_F64X2                   = 0xFD_240
	Sub_F64X2                   = 0xFD_241
	Mul_F64X2                   = 0xFD_242
	Div_F64X2                   = 0xFD_243
	Min_F64X2                   = 0xFD_244
	Max_F64X2                   = 0xFD_245
	Pmin_F64X2                  = 0xFD_246
	Pmax_F64X2                  = 0xFD_247
	TruncSat_F32X4S_I32X4       = 0xFD_248
	TruncSat_F32X4U_I32X4       = 0xFD_249
	Convert_I32X4S_F32X4        = 0xFD_250
	Convert_I32X4U_F32X4        = 0xFD_251
	TruncSat_F64X2S_Zero_I32X4  = 0xFD_252
	TruncSat_F64X2U_Zero_I32X4  = 0xFD_253
	ConvertLow_I32X4S_F64X2     = 0xFD_254
	ConvertLow_I32X4U_F64X2     = 0xFD_255
	Demote_F64X2_Zero_F32X4     = 0xFD_94
	PromoteLow_F32X4_F64X2      = 0xFD_95

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
