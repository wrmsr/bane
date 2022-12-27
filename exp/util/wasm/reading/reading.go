package reading

import (
	"encoding/binary"
	"fmt"

	"github.com/wrmsr/bane/exp/util/wasm/exprs"
	"github.com/wrmsr/bane/exp/util/wasm/instrs"
	"github.com/wrmsr/bane/exp/util/wasm/reading/consts"
	"github.com/wrmsr/bane/exp/util/wasm/reading/leb128"
	"github.com/wrmsr/bane/pkg/util/check"
)

/*
https://en.wikipedia.org/wiki/LEB128
https://en.wikipedia.org/wiki/LEB128#Signed_LEB128

numtype
0x7F ⇒ i32
| 0x7E ⇒ i64
| 0x7D ⇒ f32
| 0x7C ⇒ f64

vectype
0x7B -> v128

reftype
0x70 ⇒ funcref
| 0x6F ⇒ externref

functype ::= 0x60 rt1: resulttype rt2: resulttype ⇒ rt1 → rt2

limits ::= 0x00 𝑛:u32 ⇒ {min 𝑛, max 𝜖}
| 0x01 𝑛:u32 𝑚:u32 ⇒ {min 𝑛, max 𝑚}

globaltype ::= 𝑡:valtype 𝑚:mut ⇒ 𝑚 𝑡
mut ::= 0x00 ⇒ const
| 0x01 ⇒ var

control
blocktype ::= 0x40 ⇒ 𝜖
| 𝑡:valtype ⇒ 𝑡
| 𝑥:s33 ⇒ 𝑥 (if 𝑥 ≥ 0)
instr ::= 0x00 ⇒ unreachable
| 0x01 ⇒ nop
| 0x02 bt:blocktype (in:instr) * 0x0B ⇒ block bt in* end
| 0x03 bt:blocktype (in:instr) * 0x0B ⇒ loop bt in* end
| 0x04 bt:blocktype (in:instr) * 0x0B ⇒ if bt in* else 𝜖 end
| 0x04 bt:blocktype (in1:instr) * 0x05 (in2:instr) * 0x0B ⇒ if bt in* 1 else in* 2 end
| 0x0C 𝑙:labelidx ⇒ br 𝑙
| 0x0D 𝑙:labelidx ⇒ br_if 𝑙
| 0x0E 𝑙 * :vec(labelidx) 𝑙𝑁 :labelidx ⇒ br_table 𝑙 * 𝑙𝑁
| 0x0F ⇒ return
| 0x10 𝑥:funcidx ⇒ call 𝑥
| 0x11 𝑦:typeidx 𝑥:tableidx ⇒ call_indirect 𝑥 �

reference
instr ::= . . .
| 0xD0 𝑡:reftype ⇒ ref.null 𝑡
| 0xD1 ⇒ ref.is_null
| 0xD2 𝑥:funcidx ⇒ ref.func �

parametric
instr ::= . . .
| 0x1A ⇒ drop
| 0x1B ⇒ select
| 0x1C 𝑡 * :vec(valtype) ⇒ select 𝑡 *

variable
instr ::= . . .
| 0x20 𝑥:localidx ⇒ local.get 𝑥
| 0x21 𝑥:localidx ⇒ local.set 𝑥
| 0x22 𝑥:localidx ⇒ local.tee 𝑥
| 0x23 𝑥:globalidx ⇒ global.get 𝑥
| 0x24 𝑥:globalidx ⇒ global.set 𝑥

table
instr ::= . . .
| 0x25 𝑥:tableidx ⇒ table.get 𝑥
| 0x26 𝑥:tableidx ⇒ table.set 𝑥
| 0xFC 12:u32 𝑦:elemidx 𝑥:tableidx ⇒ table.init 𝑥 𝑦
| 0xFC 13:u32 𝑥:elemidx ⇒ elem.drop 𝑥
| 0xFC 14:u32 𝑥:tableidx 𝑦:tableidx ⇒ table.copy 𝑥 𝑦
| 0xFC 15:u32 𝑥:tableidx ⇒ table.grow 𝑥
| 0xFC 16:u32 𝑥:tableidx ⇒ table.size 𝑥
| 0xFC 17:u32 𝑥:tableidx ⇒ table.fill �

memory
memarg ::= 𝑎:u32 𝑜:u32 ⇒ {align 𝑎, offset 𝑜}
instr ::= . . .
| 0x28 𝑚:memarg ⇒ i32.load 𝑚
| 0x29 𝑚:memarg ⇒ i64.load 𝑚
| 0x2A 𝑚:memarg ⇒ f32.load 𝑚
| 0x2B 𝑚:memarg ⇒ f64.load 𝑚
| 0x2C 𝑚:memarg ⇒ i32.load8_s 𝑚
| 0x2D 𝑚:memarg ⇒ i32.load8_u 𝑚
| 0x2E 𝑚:memarg ⇒ i32.load16_s 𝑚
| 0x2F 𝑚:memarg ⇒ i32.load16_u 𝑚
| 0x30 𝑚:memarg ⇒ i64.load8_s 𝑚
| 0x31 𝑚:memarg ⇒ i64.load8_u 𝑚
| 0x32 𝑚:memarg ⇒ i64.load16_s 𝑚
| 0x33 𝑚:memarg ⇒ i64.load16_u 𝑚
| 0x34 𝑚:memarg ⇒ i64.load32_s 𝑚
| 0x35 𝑚:memarg ⇒ i64.load32_u 𝑚
| 0x36 𝑚:memarg ⇒ i32.store 𝑚
| 0x37 𝑚:memarg ⇒ i64.store 𝑚
| 0x38 𝑚:memarg ⇒ f32.store 𝑚
| 0x39 𝑚:memarg ⇒ f64.store 𝑚
| 0x3A 𝑚:memarg ⇒ i32.store8 𝑚
| 0x3B 𝑚:memarg ⇒ i32.store16 𝑚
| 0x3C 𝑚:memarg ⇒ i64.store8 𝑚
| 0x3D 𝑚:memarg ⇒ i64.store16 𝑚
| 0x3E 𝑚:memarg ⇒ i64.store32 𝑚
| 0x3F 0x00 ⇒ memory.size
| 0x40 0x00 ⇒ memory.grow
| 0xFC 8:u32 𝑥:dataidx 0x00 ⇒ memory.init 𝑥
| 0xFC 9:u32 𝑥:dataidx ⇒ data.drop 𝑥
| 0xFC 10:u32 0x00 0x00 ⇒ memory.copy
| 0xFC 11:u32 0x00 ⇒ memory.fill

instr ::= . . .
| 0x41 𝑛:i32 ⇒ i32.const 𝑛
| 0x42 𝑛:i64 ⇒ i64.const 𝑛
| 0x43 𝑧:f32 ⇒ f32.const 𝑧
| 0x44 𝑧:f64 ⇒ f64.const �
| 0x45 ⇒ i32.eqz
| 0x46 ⇒ i32.eq
| 0x47 ⇒ i32.ne
| 0x48 ⇒ i32.lt_s
| 0x49 ⇒ i32.lt_u
| 0x4A ⇒ i32.gt_s
| 0x4B ⇒ i32.gt_u
| 0x4C ⇒ i32.le_s
| 0x4D ⇒ i32.le_u
| 0x4E ⇒ i32.ge_s
| 0x4F ⇒ i32.ge_u
| 0x50 ⇒ i64.eqz
| 0x51 ⇒ i64.eq
| 0x52 ⇒ i64.ne
| 0x53 ⇒ i64.lt_s
| 0x54 ⇒ i64.lt_u
| 0x55 ⇒ i64.gt_s
| 0x56 ⇒ i64.gt_u
| 0x57 ⇒ i64.le_s
| 0x58 ⇒ i64.le_u
| 0x59 ⇒ i64.ge_s
| 0x5A ⇒ i64.ge_u
| 0x5B ⇒ f32.eq
| 0x5C ⇒ f32.ne
| 0x5D ⇒ f32.lt
| 0x5E ⇒ f32.gt
| 0x5F ⇒ f32.le
| 0x60 ⇒ f32.ge
| 0x61 ⇒ f64.eq
| 0x62 ⇒ f64.ne
| 0x63 ⇒ f64.lt
| 0x64 ⇒ f64.gt
| 0x65 ⇒ f64.le
| 0x66 ⇒ f64.ge
| 0x67 ⇒ i32.clz
| 0x68 ⇒ i32.ctz
| 0x69 ⇒ i32.popcnt
| 0x6A ⇒ i32.add
| 0x6B ⇒ i32.sub
| 0x6C ⇒ i32.mul
| 0x6D ⇒ i32.div_s
| 0x6E ⇒ i32.div_u
| 0x6F ⇒ i32.rem_s
| 0x70 ⇒ i32.rem_u
| 0x71 ⇒ i32.and
| 0x72 ⇒ i32.or
| 0x73 ⇒ i32.xor
| 0x74 ⇒ i32.shl
| 0x75 ⇒ i32.shr_s
| 0x76 ⇒ i32.shr_u
| 0x77 ⇒ i32.rotl
| 0x78 ⇒ i32.rotr
| 0x79 ⇒ i64.clz
| 0x7A ⇒ i64.ctz
| 0x7B ⇒ i64.popcnt
| 0x7C ⇒ i64.add
| 0x7D ⇒ i64.sub
| 0x7E ⇒ i64.mul
| 0x7F ⇒ i64.div_s
| 0x80 ⇒ i64.div_u
| 0x81 ⇒ i64.rem_s
| 0x82 ⇒ i64.rem_u
| 0x83 ⇒ i64.and
| 0x84 ⇒ i64.or
| 0x85 ⇒ i64.xor
| 0x86 ⇒ i64.shl
| 0x87 ⇒ i64.shr_s
| 0x88 ⇒ i64.shr_u
| 0x89 ⇒ i64.rotl
| 0x8A ⇒ i64.rotr
| 0x8B ⇒ f32.abs
| 0x8C ⇒ f32.neg
| 0x8D ⇒ f32.ceil
| 0x8E ⇒ f32.floor
| 0x8F ⇒ f32.trunc
| 0x90 ⇒ f32.nearest
| 0x91 ⇒ f32.sqrt
| 0x92 ⇒ f32.add
| 0x93 ⇒ f32.sub
| 0x94 ⇒ f32.mul
| 0x95 ⇒ f32.div
| 0x96 ⇒ f32.min
| 0x97 ⇒ f32.max
| 0x98 ⇒ f32.copysign
| 0x99 ⇒ f64.abs
| 0x9A ⇒ f64.neg
| 0x9B ⇒ f64.ceil
| 0x9C ⇒ f64.floor
| 0x9D ⇒ f64.trunc
| 0x9E ⇒ f64.nearest
| 0x9F ⇒ f64.sqrt
| 0xA0 ⇒ f64.add
| 0xA1 ⇒ f64.sub
| 0xA2 ⇒ f64.mul
| 0xA3 ⇒ f64.div
| 0xA4 ⇒ f64.min
| 0xA5 ⇒ f64.max
| 0xA6 ⇒ f64.copysign
| 0xA7 ⇒ i32.wrap_i64
| 0xA8 ⇒ i32.trunc_f32_s
| 0xA9 ⇒ i32.trunc_f32_u
| 0xAA ⇒ i32.trunc_f64_s
| 0xAB ⇒ i32.trunc_f64_u
| 0xAC ⇒ i64.extend_i32_s
| 0xAD ⇒ i64.extend_i32_u
| 0xAE ⇒ i64.trunc_f32_s
| 0xAF ⇒ i64.trunc_f32_u
| 0xB0 ⇒ i64.trunc_f64_s
| 0xB1 ⇒ i64.trunc_f64_u
| 0xB2 ⇒ f32.convert_i32_s
| 0xB3 ⇒ f32.convert_i32_u
| 0xB4 ⇒ f32.convert_i64_s
| 0xB5 ⇒ f32.convert_i64_u
| 0xB6 ⇒ f32.demote_f64
| 0xB7 ⇒ f64.convert_i32_s
| 0xB8 ⇒ f64.convert_i32_u
| 0xB9 ⇒ f64.convert_i64_s
| 0xBA ⇒ f64.convert_i64_u
| 0xBB ⇒ f64.promote_f32
| 0xBC ⇒ i32.reinterpret_f32
| 0xBD ⇒ i64.reinterpret_f64
| 0xBE ⇒ f32.reinterpret_i32
| 0xBF ⇒ f64.reinterpret_i64
| 0xC0 ⇒ i32.extend8_s
| 0xC1 ⇒ i32.extend16_s
| 0xC2 ⇒ i64.extend8_s
| 0xC3 ⇒ i64.extend16_s
| 0xC4 ⇒ i64.extend32_s

instr ::= . . .
| 0xFC 0:u32 ⇒ i32.trunc_sat_f32_s
| 0xFC 1:u32 ⇒ i32.trunc_sat_f32_u
| 0xFC 2:u32 ⇒ i32.trunc_sat_f64_s
| 0xFC 3:u32 ⇒ i32.trunc_sat_f64_u
| 0xFC 4:u32 ⇒ i64.trunc_sat_f32_s
| 0xFC 5:u32 ⇒ i64.trunc_sat_f32_u
| 0xFC 6:u32 ⇒ i64.trunc_sat_f64_s
| 0xFC 7:u32 ⇒ i64.trunc_sat_f64_u

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

type ModuleReader struct {
	r *ByteReader
}

func (r *ModuleReader) readByte() byte {
	return check.Must1(r.r.ReadByte())
}

func (r *ModuleReader) readI64() int64 {
	return leb128.DecodeI64(r.readByte)
}

func (r *ModuleReader) readU64() uint64 {
	return leb128.DecodeU64(r.readByte)
}

func (r *ModuleReader) readI32le() int32 {
	var v int32
	check.Must(binary.Read(r.r, binary.LittleEndian, &v))
	return v
}

func (r *ModuleReader) readString() string {
	l := r.readU64()
	b := make([]byte, l)
	check.Must1(r.r.Read(b))
	return string(b)
}

func (r *ModuleReader) ReadModule() {
	magic := r.readI32le()
	check.Equal(magic, consts.Magic)

	version := r.readI32le()
	check.Equal(version, consts.Version)

	for r.r.Len() > 0 {
		r.readSection()
	}

	check.Equal(r.r.Len(), 0)
}

func (r *ModuleReader) readSection() {
	secTy := r.readByte()
	secSz := r.readU64()

	switch secTy {

	case consts.TypeSection:
		numSigs := int(r.readU64())
		for i := 0; i < numSigs; i++ {
			ty := r.readByte()
			check.Equal(ty, consts.FuncType)

			numParam := int(r.readU64())
			for j := 0; j < numParam; j++ {
				pty := r.readI64()
				_ = pty
			}

			numResult := int(r.readU64())
			for j := 0; j < numResult; j++ {
				rty := r.readI64()
				_ = rty
			}
		}

	case consts.ImportSection:
		numImps := int(r.readU64())
		for i := 0; i < numImps; i++ {
			modName := r.readString()
			_ = modName
			fldName := r.readString()
			_ = fldName

			k := r.readByte()
			switch k {
			case consts.FuncImport:
				idx := r.readU64()
				_ = idx
			default:
				panic(k)
			}
		}

	case consts.FunctionSection:
		numFuncs := int(r.readU64())
		for i := 0; i < numFuncs; i++ {
			sigIdx := r.readU64()
			_ = sigIdx
		}

	case consts.TableSection:
		//numTables := int(r.readU64())
		//for i := 0; i < numTables; i++ {
		//}
		check.Must(r.r.Skip(int64(secSz)))

	case consts.CodeSection:
		numFuncBodies := int(r.readU64())

		for i := 0; i < numFuncBodies; i++ {
			bodySize := int(r.readU64())

			startPos := r.r.Tell()
			endPos := startPos + int64(bodySize)

			numLocals := int(r.readU64())
			for j := 0; j < numLocals; j++ {
				numTys := int(r.readU64())
				_ = numTys
				ty := r.readI64()
				_ = ty
			}

			var es []exprs.Expr
			for r.r.Tell() < endPos {
				var opPfx int
				op := int(r.readByte())

				if op == consts.MathPrefix || op == consts.SimdPrefix {
					opPfx = op
					op = int(r.readByte())
				}

				inst := instrs.ByOp(int8(opPfx), int8(op))

				switch inst.I() {

				case instrs.Block:
					ty := r.readI64()
					_ = ty

				case
					instrs.LocalGet:
					idx := r.readU64()
					es = append(es, exprs.LocalGet{I: int(idx)})

				case
					instrs.LocalSet,
					instrs.GlobalGet,
					instrs.GlobalSet:
					idx := r.readU64()
					_ = idx

				case instrs.Call:
					idx := r.readU64()
					_ = idx

				case instrs.Const_I32:
					v := int32(r.readI64())
					_ = v

				case instrs.Const_I64:
					v := r.readI64()
					_ = v

				case
					instrs.Eqz_I32,
					instrs.Eq_I32,
					instrs.Ne_I32,
					instrs.LtS_I32,
					instrs.LtU_I32,
					instrs.GtS_I32,
					instrs.GtU_I32,
					instrs.LeS_I32,
					instrs.LeU_I32,
					instrs.GeS_I32,
					instrs.GeU_I32,

					instrs.Eqz_I64,
					instrs.Eq_I64,
					instrs.Ne_I64,
					instrs.LtS_I64,
					instrs.LtU_I64,
					instrs.GtS_I64,
					instrs.GtU_I64,
					instrs.LeS_I64,
					instrs.LeU_I64,
					instrs.GeS_I64,
					instrs.GeU_I64,

					instrs.Eq_F32,
					instrs.Ne_F32,
					instrs.Lt_F32,
					instrs.Gt_F32,
					instrs.Le_F32,
					instrs.Ge_F32,

					instrs.Eq_F64,
					instrs.Ne_F64,
					instrs.Lt_F64,
					instrs.Gt_F64,
					instrs.Le_F64,
					instrs.Ge_F64:
					//

				default:
					panic(inst)
				}
			}
		}

	default:
		fmt.Printf("unhandled section: %v %v\n", secTy, secSz)
		check.Must(r.r.Skip(int64(secSz)))

	}
}
