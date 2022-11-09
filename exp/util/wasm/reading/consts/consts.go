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

	/*
	   laneidx ::= 𝑙:byte ⇒ 𝑙
	   instr ::= . . .
	   | 0xFD 0:u32 𝑚:memarg ⇒ v128.load 𝑚
	   | 0xFD 1:u32 𝑚:memarg ⇒ v128.load8x8_s 𝑚
	   | 0xFD 2:u32 𝑚:memarg ⇒ v128.load8x8_u 𝑚
	   | 0xFD 3:u32 𝑚:memarg ⇒ v128.load16x4_s 𝑚
	   | 0xFD 4:u32 𝑚:memarg ⇒ v128.load16x4_u 𝑚
	   | 0xFD 5:u32 𝑚:memarg ⇒ v128.load32x2_s 𝑚
	   | 0xFD 6:u32 𝑚:memarg ⇒ v128.load32x2_u 𝑚
	   | 0xFD 7:u32 𝑚:memarg ⇒ v128.load8_splat 𝑚
	   | 0xFD 8:u32 𝑚:memarg ⇒ v128.load16_splat 𝑚
	   | 0xFD 9:u32 𝑚:memarg ⇒ v128.load32_splat 𝑚
	   | 0xFD 10:u32 𝑚:memarg ⇒ v128.load64_splat 𝑚
	   | 0xFD 92:u32 𝑚:memarg ⇒ v128.load32_zero 𝑚
	   | 0xFD 93:u32 𝑚:memarg ⇒ v128.load64_zero 𝑚
	   | 0xFD 11:u32 𝑚:memarg ⇒ v128.store 𝑚
	   | 0xFD 84:u32 𝑚:memarg 𝑙:laneidx ⇒ v128.load8_lane 𝑚 𝑙
	   | 0xFD 85:u32 𝑚:memarg 𝑙:laneidx ⇒ v128.load16_lane 𝑚 𝑙
	   | 0xFD 86:u32 𝑚:memarg 𝑙:laneidx ⇒ v128.load32_lane 𝑚 𝑙
	   | 0xFD 87:u32 𝑚:memarg 𝑙:laneidx ⇒ v128.load64_lane 𝑚 𝑙
	   | 0xFD 88:u32 𝑚:memarg 𝑙:laneidx ⇒ v128.store8_lane 𝑚 𝑙
	   | 0xFD 89:u32 𝑚:memarg 𝑙:laneidx ⇒ v128.store16_lane 𝑚 𝑙
	   | 0xFD 90:u32 𝑚:memarg 𝑙:laneidx ⇒ v128.store32_lane 𝑚 𝑙
	   | 0xFD 91:u32 𝑚:memarg 𝑙:laneidx ⇒ v128.store64_lane 𝑚 𝑙

	   instr ::= . . .
	   | 0xFD 12:u32 (𝑏:byte) 16 ⇒ v128.const bytes−1 i128(𝑏0 . . . 𝑏15)

	   instr ::= . . .
	   | 0xFD 13:u32 (𝑙:laneidx) 16 ⇒ i8x16.shuffle 𝑙 16

	   instr ::= . . .
	   | 0xFD 21:u32 𝑙:laneidx ⇒ i8x16.extract_lane_s 𝑙
	   | 0xFD 22:u32 𝑙:laneidx ⇒ i8x16.extract_lane_u 𝑙
	   | 0xFD 23:u32 𝑙:laneidx ⇒ i8x16.replace_lane 𝑙
	   | 0xFD 24:u32 𝑙:laneidx ⇒ i16x8.extract_lane_s 𝑙
	   | 0xFD 25:u32 𝑙:laneidx ⇒ i16x8.extract_lane_u 𝑙
	   | 0xFD 26:u32 𝑙:laneidx ⇒ i16x8.replace_lane 𝑙
	   | 0xFD 27:u32 𝑙:laneidx ⇒ i32x4.extract_lane 𝑙
	   | 0xFD 28:u32 𝑙:laneidx ⇒ i32x4.replace_lane 𝑙
	   | 0xFD 29:u32 𝑙:laneidx ⇒ i64x2.extract_lane 𝑙
	   | 0xFD 30:u32 𝑙:laneidx ⇒ i64x2.replace_lane 𝑙
	   | 0xFD 31:u32 𝑙:laneidx ⇒ f32x4.extract_lane 𝑙
	   | 0xFD 32:u32 𝑙:laneidx ⇒ f32x4.replace_lane 𝑙
	   | 0xFD 33:u32 𝑙:laneidx ⇒ f64x2.extract_lane 𝑙
	   | 0xFD 34:u32 𝑙:laneidx ⇒ f64x2.replace_lane 𝑙

	   instr ::= . . .
	   | 0xFD 14:u32 ⇒ i8x16.swizzle
	   | 0xFD 15:u32 ⇒ i8x16.splat
	   | 0xFD 16:u32 ⇒ i16x8.splat
	   | 0xFD 17:u32 ⇒ i32x4.splat
	   | 0xFD 18:u32 ⇒ i64x2.splat
	   | 0xFD 19:u32 ⇒ f32x4.splat
	   | 0xFD 20:u32 ⇒ f64x2.splat
	   | 0xFD 35:u32 ⇒ i8x16.eq
	   | 0xFD 36:u32 ⇒ i8x16.ne
	   | 0xFD 37:u32 ⇒ i8x16.lt_s
	   | 0xFD 38:u32 ⇒ i8x16.lt_u
	   | 0xFD 39:u32 ⇒ i8x16.gt_s
	   | 0xFD 40:u32 ⇒ i8x16.gt_u
	   | 0xFD 41:u32 ⇒ i8x16.le_s
	   | 0xFD 42:u32 ⇒ i8x16.le_u
	   | 0xFD 43:u32 ⇒ i8x16.ge_s
	   | 0xFD 44:u32 ⇒ i8x16.ge_u
	   | 0xFD 45:u32 ⇒ i16x8.eq
	   | 0xFD 46:u32 ⇒ i16x8.ne
	   | 0xFD 47:u32 ⇒ i16x8.lt_s
	   | 0xFD 48:u32 ⇒ i16x8.lt_u
	   | 0xFD 49:u32 ⇒ i16x8.gt_s
	   | 0xFD 50:u32 ⇒ i16x8.gt_u
	   | 0xFD 51:u32 ⇒ i16x8.le_s
	   | 0xFD 52:u32 ⇒ i16x8.le_u
	   | 0xFD 53:u32 ⇒ i16x8.ge_s
	   | 0xFD 54:u32 ⇒ i16x8.ge_u
	   | 0xFD 55:u32 ⇒ i32x4.eq
	   | 0xFD 56:u32 ⇒ i32x4.ne
	   | 0xFD 57:u32 ⇒ i32x4.lt_s
	   | 0xFD 58:u32 ⇒ i32x4.lt_u
	   | 0xFD 59:u32 ⇒ i32x4.gt_s
	   | 0xFD 60:u32 ⇒ i32x4.gt_u
	   | 0xFD 61:u32 ⇒ i32x4.le_s
	   | 0xFD 62:u32 ⇒ i32x4.le_u
	   | 0xFD 63:u32 ⇒ i32x4.ge_s
	   | 0xFD 64:u32 ⇒ i32x4.ge_u
	   | 0xFD 214:u32 ⇒ i64x2.eq
	   | 0xFD 215:u32 ⇒ i64x2.ne
	   | 0xFD 216:u32 ⇒ i64x2.lt_s
	   | 0xFD 217:u32 ⇒ i64x2.gt_s
	   | 0xFD 218:u32 ⇒ i64x2.le_s
	   | 0xFD 219:u32 ⇒ i64x2.ge_s
	   | 0xFD 65:u32 ⇒ f32x4.eq
	   | 0xFD 66:u32 ⇒ f32x4.ne
	   | 0xFD 67:u32 ⇒ f32x4.lt
	   | 0xFD 68:u32 ⇒ f32x4.gt
	   | 0xFD 69:u32 ⇒ f32x4.le
	   | 0xFD 70:u32 ⇒ f32x4.ge
	   | 0xFD 71:u32 ⇒ f64x2.eq
	   | 0xFD 72:u32 ⇒ f64x2.ne
	   | 0xFD 73:u32 ⇒ f64x2.lt
	   | 0xFD 74:u32 ⇒ f64x2.gt
	   | 0xFD 75:u32 ⇒ f64x2.le
	   | 0xFD 76:u32 ⇒ f64x2.ge
	   | 0xFD 77:u32 ⇒ v128.not
	   | 0xFD 78:u32 ⇒ v128.and
	   | 0xFD 79:u32 ⇒ v128.andnot
	   | 0xFD 80:u32 ⇒ v128.or
	   | 0xFD 81:u32 ⇒ v128.xor
	   | 0xFD 82:u32 ⇒ v128.bitselect
	   | 0xFD 83:u32 ⇒ v128.any_true
	   | 0xFD 96:u32 ⇒ i8x16.abs
	   | 0xFD 97:u32 ⇒ i8x16.neg
	   | 0xFD 98:u32 ⇒ i8x16.popcnt
	   | 0xFD 99:u32 ⇒ i8x16.all_true
	   | 0xFD 100:u32 ⇒ i8x16.bitmask
	   | 0xFD 101:u32 ⇒ i8x16.narrow_i16x8_s
	   | 0xFD 102:u32 ⇒ i8x16.narrow_i16x8_u
	   | 0xFD 107:u32 ⇒ i8x16.shl
	   | 0xFD 108:u32 ⇒ i8x16.shr_s
	   | 0xFD 109:u32 ⇒ i8x16.shr_u
	   | 0xFD 110:u32 ⇒ i8x16.add
	   | 0xFD 111:u32 ⇒ i8x16.add_sat_s
	   | 0xFD 112:u32 ⇒ i8x16.add_sat_u
	   | 0xFD 113:u32 ⇒ i8x16.sub
	   | 0xFD 114:u32 ⇒ i8x16.sub_sat_s
	   | 0xFD 115:u32 ⇒ i8x16.sub_sat_u
	   | 0xFD 118:u32 ⇒ i8x16.min_s
	   | 0xFD 119:u32 ⇒ i8x16.min_u
	   | 0xFD 120:u32 ⇒ i8x16.max_s
	   | 0xFD 121:u32 ⇒ i8x16.max_u
	   | 0xFD 123:u32 ⇒ i8x16.avgr_u
	   | 0xFD 124:u32 ⇒ i16x8.extadd_pairwise_i8x16_s
	   | 0xFD 125:u32 ⇒ i16x8.extadd_pairwise_i8x16_u
	   | 0xFD 128:u32 ⇒ i16x8.abs
	   | 0xFD 129:u32 ⇒ i16x8.neg
	   | 0xFD 130:u32 ⇒ i16x8.q15mulr_sat_s
	   | 0xFD 131:u32 ⇒ i16x8.all_true
	   | 0xFD 132:u32 ⇒ i16x8.bitmask
	   | 0xFD 133:u32 ⇒ i16x8.narrow_i32x4_s
	   | 0xFD 134:u32 ⇒ i16x8.narrow_i32x4_u
	   | 0xFD 135:u32 ⇒ i16x8.extend_low_i8x16_s
	   | 0xFD 136:u32 ⇒ i16x8.extend_high_i8x16_s
	   | 0xFD 137:u32 ⇒ i16x8.extend_low_i8x16_u
	   | 0xFD 138:u32 ⇒ i16x8.extend_high_i8x16_u
	   | 0xFD 139:u32 ⇒ i16x8.shl
	   | 0xFD 140:u32 ⇒ i16x8.shr_s
	   | 0xFD 141:u32 ⇒ i16x8.shr_u
	   | 0xFD 142:u32 ⇒ i16x8.add
	   | 0xFD 143:u32 ⇒ i16x8.add_sat_s
	   | 0xFD 144:u32 ⇒ i16x8.add_sat_u
	   | 0xFD 145:u32 ⇒ i16x8.sub
	   | 0xFD 146:u32 ⇒ i16x8.sub_sat_s
	   | 0xFD 147:u32 ⇒ i16x8.sub_sat_u
	   | 0xFD 149:u32 ⇒ i16x8.mul
	   | 0xFD 150:u32 ⇒ i16x8.min_s
	   | 0xFD 151:u32 ⇒ i16x8.min_u
	   | 0xFD 152:u32 ⇒ i16x8.max_s
	   | 0xFD 153:u32 ⇒ i16x8.max_u
	   | 0xFD 155:u32 ⇒ i16x8.avgr_u
	   | 0xFD 156:u32 ⇒ i16x8.extmul_low_i8x16_s
	   | 0xFD 157:u32 ⇒ i16x8.extmul_high_i8x16_s
	   | 0xFD 158:u32 ⇒ i16x8.extmul_low_i8x16_u
	   | 0xFD 159:u32 ⇒ i16x8.extmul_high_i8x16_u
	   | 0xFD 126:u32 ⇒ i32x4.extadd_pairwise_i16x8_s
	   | 0xFD 127:u32 ⇒ i32x4.extadd_pairwise_i16x8_u
	   | 0xFD 160:u32 ⇒ i32x4.abs
	   | 0xFD 161:u32 ⇒ i32x4.neg
	   | 0xFD 163:u32 ⇒ i32x4.all_true
	   | 0xFD 164:u32 ⇒ i32x4.bitmask
	   | 0xFD 167:u32 ⇒ i32x4.extend_low_i16x8_s
	   | 0xFD 168:u32 ⇒ i32x4.extend_high_i16x8_s
	   | 0xFD 169:u32 ⇒ i32x4.extend_low_i16x8_u
	   | 0xFD 170:u32 ⇒ i32x4.extend_high_i16x8_u
	   | 0xFD 171:u32 ⇒ i32x4.shl
	   | 0xFD 172:u32 ⇒ i32x4.shr_s
	   | 0xFD 173:u32 ⇒ i32x4.shr_u
	   | 0xFD 174:u32 ⇒ i32x4.add
	   | 0xFD 177:u32 ⇒ i32x4.sub
	   | 0xFD 181:u32 ⇒ i32x4.mul
	   | 0xFD 182:u32 ⇒ i32x4.min_s
	   | 0xFD 183:u32 ⇒ i32x4.min_u
	   | 0xFD 184:u32 ⇒ i32x4.max_s
	   | 0xFD 185:u32 ⇒ i32x4.max_u
	   | 0xFD 186:u32 ⇒ i32x4.dot_i16x8_s
	   | 0xFD 188:u32 ⇒ i32x4.extmul_low_i16x8_s
	   | 0xFD 189:u32 ⇒ i32x4.extmul_high_i16x8_s
	   | 0xFD 190:u32 ⇒ i32x4.extmul_low_i16x8_u
	   | 0xFD 191:u32 ⇒ i32x4.extmul_high_i16x8_u
	   | 0xFD 192:u32 ⇒ i64x2.abs
	   | 0xFD 193:u32 ⇒ i64x2.neg
	   | 0xFD 195:u32 ⇒ i64x2.all_true
	   | 0xFD 196:u32 ⇒ i64x2.bitmask
	   | 0xFD 199:u32 ⇒ i64x2.extend_low_i32x4_s
	   | 0xFD 200:u32 ⇒ i64x2.extend_high_i32x4_s
	   | 0xFD 201:u32 ⇒ i64x2.extend_low_i32x4_u
	   | 0xFD 202:u32 ⇒ i64x2.extend_high_i32x4_u
	   | 0xFD 203:u32 ⇒ i64x2.shl
	   | 0xFD 204:u32 ⇒ i64x2.shr_s
	   | 0xFD 205:u32 ⇒ i64x2.shr_u
	   | 0xFD 206:u32 ⇒ i64x2.add
	   | 0xFD 209:u32 ⇒ i64x2.sub
	   | 0xFD 213:u32 ⇒ i64x2.mul
	   | 0xFD 220:u32 ⇒ i64x2.extmul_low_i32x4_s
	   | 0xFD 221:u32 ⇒ i64x2.extmul_high_i32x4_s
	   | 0xFD 222:u32 ⇒ i64x2.extmul_low_i32x4_u
	   | 0xFD 223:u32 ⇒ i64x2.extmul_high_i32x4_u
	   | 0xFD 103:u32 ⇒ f32x4.ceil
	   | 0xFD 104:u32 ⇒ f32x4.floor
	   | 0xFD 105:u32 ⇒ f32x4.trunc
	   | 0xFD 106:u32 ⇒ f32x4.nearest
	   | 0xFD 224:u32 ⇒ f32x4.abs
	   | 0xFD 225:u32 ⇒ f32x4.neg
	   | 0xFD 227:u32 ⇒ f32x4.sqrt
	   | 0xFD 228:u32 ⇒ f32x4.add
	   | 0xFD 229:u32 ⇒ f32x4.sub
	   | 0xFD 230:u32 ⇒ f32x4.mul
	   | 0xFD 231:u32 ⇒ f32x4.div
	   | 0xFD 232:u32 ⇒ f32x4.min
	   | 0xFD 233:u32 ⇒ f32x4.max
	   | 0xFD 234:u32 ⇒ f32x4.pmin
	   | 0xFD 235:u32 ⇒ f32x4.pmax
	   | 0xFD 116:u32 ⇒ f64x2.ceil
	   | 0xFD 117:u32 ⇒ f64x2.floor
	   | 0xFD 122:u32 ⇒ f64x2.trunc
	   | 0xFD 148:u32 ⇒ f64x2.nearest
	   | 0xFD 236:u32 ⇒ f64x2.abs
	   | 0xFD 237:u32 ⇒ f64x2.neg
	   | 0xFD 239:u32 ⇒ f64x2.sqrt
	   | 0xFD 240:u32 ⇒ f64x2.add
	   | 0xFD 241:u32 ⇒ f64x2.sub
	   | 0xFD 242:u32 ⇒ f64x2.mul
	   | 0xFD 243:u32 ⇒ f64x2.div
	   | 0xFD 244:u32 ⇒ f64x2.min
	   | 0xFD 245:u32 ⇒ f64x2.max
	   | 0xFD 246:u32 ⇒ f64x2.pmin
	   | 0xFD 247:u32 ⇒ f64x2.pmax
	   | 0xFD 248:u32 ⇒ i32x4.trunc_sat_f32x4_s
	   | 0xFD 249:u32 ⇒ i32x4.trunc_sat_f32x4_u
	   | 0xFD 250:u32 ⇒ f32x4.convert_i32x4_s
	   | 0xFD 251:u32 ⇒ f32x4.convert_i32x4_u
	   | 0xFD 252:u32 ⇒ i32x4.trunc_sat_f64x2_s_zero
	   | 0xFD 253:u32 ⇒ i32x4.trunc_sat_f64x2_u_zero
	   | 0xFD 254:u32 ⇒ f64x2.convert_low_i32x4_s
	   | 0xFD 255:u32 ⇒ f64x2.convert_low_i32x4_u
	   | 0xFD 94:u32 ⇒ f32x4.demote_f64x2_zero
	   | 0xFD 95:u32 ⇒ f64x2.promote_low_f32x4

	   expr ::= (in:instr) * 0x0B ⇒ in* end

	   e.
	   typeidx ::= 𝑥:u32 ⇒ 𝑥
	   funcidx ::= 𝑥:u32 ⇒ 𝑥
	   tableidx ::= 𝑥:u32 ⇒ 𝑥
	   memidx ::= 𝑥:u32 ⇒ 𝑥
	   globalidx ::= 𝑥:u32 ⇒ 𝑥
	   elemidx ::= 𝑥:u32 ⇒ 𝑥
	   dataidx ::= 𝑥:u32 ⇒ 𝑥
	   localidx ::= 𝑥:u32 ⇒ 𝑥
	   labelidx ::= 𝑙:u32 ⇒ �

	   .
	   section𝑁 (B) ::= 𝑁:byte size:u32 cont:B ⇒ cont (if size = ||B||)
	   | 𝜖 ⇒ 𝜖

	   Id Section
	   0 custom section
	   1 type section
	   2 import section
	   3 function section
	   4 table section
	   5 memory section
	   6 global section
	   7 export section
	   8 start section
	   9 element section
	   10 code section
	   11 data section
	   12 data count section

	   typesec ::= ft* : section1(vec(functype)) ⇒ ft*

	   importsec ::= im* :section2(vec(import)) ⇒ im*
	   import ::= mod:name nm:name 𝑑:importdesc ⇒ {module mod, name nm, desc 𝑑}
	   importdesc ::= 0x00 𝑥:typeidx ⇒ func 𝑥
	   | 0x01 tt:tabletype ⇒ table tt
	   | 0x02 mt:memtype ⇒ mem mt
	   | 0x03 gt:globaltype ⇒ global gt

	   funcsec ::= 𝑥 * :section3(vec(typeidx)) ⇒ �

	   tablesec ::= tab* :section4(vec(table)) ⇒ tab*
	   table ::= tt:tabletype ⇒ {type tt}

	   memsec ::= mem* :section5(vec(mem)) ⇒ mem*
	   mem ::= mt:memtype ⇒ {type mt}

	   globalsec ::= glob* :section6(vec(global)) ⇒ glob*
	   global ::= gt:globaltype 𝑒:expr ⇒ {type gt, init 𝑒}

	   exportsec ::= ex * :section7(vec(export)) ⇒ ex *
	   export ::= nm:name 𝑑:exportdesc ⇒ {name nm, desc 𝑑}
	   exportdesc ::= 0x00 𝑥:funcidx ⇒ func 𝑥
	   | 0x01 𝑥:tableidx ⇒ table 𝑥
	   | 0x02 𝑥:memidx ⇒ mem 𝑥
	   | 0x03 𝑥:globalidx ⇒ global 𝑥

	   startsec ::= st ? :section8(start) ⇒ st ?
	   start ::= 𝑥:funcidx ⇒ {func 𝑥}

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
