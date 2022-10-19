//
/*
ishape ::=
    i8x16
  | i16x8
  | i32x4
  | i64x2

fshape ::=
    f32x4
  | f64x2

shape ::=
    ishape
  | fshape

half ::=
    low
  | high

laneidx ::=
    u8

instr ::= ...
  | v128.const i128
  | v128.vvunop
  | v128.vvbinop
  | v128.vvternop
  | v128.vvtestop
  | i8x16.shuffle laneidx 16
  | i8x16.swizzle
  | shape.splat
  | i8x16.extract_lane_sx laneidx
  | i16x8.extract_lane_sx laneidx
  | i32x4.extract_lane laneidx
  | i64x2.extract_lane laneidx
  | fshape.extract_lane laneidx
  | shape.replace_lane laneidx
  | i8x16.virelop
  | i16x8.virelop
  | i32x4.virelop
  | i64x2.eq
  | i64x2.ne
  | i64x2.lt_s
  | i64x2.gt_s
  | i64x2.le_s
  | i64x2.ge_s
  | fshape.vfrelop
  | ishape.viunop
  | i8x16.popcnt
  | i16x8.q15mulr_sat_s
  | i32x4.dot_i16x8_s
  | fshape.vfunop
  | ishape.vitestop
  | ishape.bitmask
  | i8x16.narrow_i16x8_sx
  | i16x8.narrow_i32x4_sx
  | i16x8.extend_half _i8x16_sx
  | i32x4.extend_half _i16x8_sx
  | i64x2.extend_half _i32x4_sx
  | ishape.vishiftop
  | ishape.vibinop
  | i8x16.viminmaxop
  | i16x8.viminmaxop
  | i32x4.viminmaxop
  | i8x16.visatbinop
  | i16x8.visatbinop
  | i16x8.mul
  | i32x4.mul
  | i64x2.mul
  | i8x16.avgr_u
  | i16x8.avgr_u
  | i16x8.extmul_half _i8x16_sx
  | i32x4.extmul_half _i16x8_sx
  | i64x2.extmul_half _i32x4_sx
  | i16x8.extadd_pairwise_i8x16_sx
  | i32x4.extadd_pairwise_i16x8_sx
  | fshape.vfbinop
  | i32x4.trunc_sat_f32x4_sx
  | i32x4.trunc_sat_f64x2_sx_zero
  | f32x4.convert_i32x4_sx
  | f32x4.demote_f64x2_zero
  | f64x2.convert_low_i32x4_sx
  | f64x2.promote_low_f32x4

vvunop ::=
    not

vvbinop ::=
    and
  | andnot
  | or
  | xor

vvternop ::=
    bitselect

vvtestop ::=
    any_true

vitestop ::=
    all_true

virelop ::=
    eq
  | ne
  | lt_sx
  | gt_sx
  | le_sx
  | ge_sx

vfrelop ::=
    eq
  | ne
  | lt
  | gt
  | le
  | ge

viunop ::=
    abs
  | neg

vibinop ::=
    add
  | sub

viminmaxop ::=
    min_sx
  | max_sx

visatbinop ::=
    add_sat_sx
  | sub_sat_sx

vishiftop ::=
    shl
  | shr_sx

vfunop ::= abs | neg | sqrt | ceil | floor | trunc | nearest

vfbinop ::= add | sub | mul | div | min | max | pmin | pmax

//

vunop ::= viunop | vfunop | popcnt

vbinop ::= vibinop | vfbinop
| viminmaxop | visatbinop
| mul | avgr_u | q15mulr_sat_s

vtestop ::= vitestop

vrelop ::= virelop | vfrelop

vcvtop ::= extend | trunc_sat | convert | demote | promote
*/
package exprs
