//
/*
nn, mm ::=
    32
  | 64

sx ::=
    u
  | s

instr ::= ...
  | inn.const inn
  | fnn.const f nn
  | inn.iunop
  | fnn.funop
  | inn.ibinop
  | fnn.fbinop
  | inn.itestop
  | inn.irelop
  | fnn.frelop
  | inn.extend8_s
  | inn.extend16_s
  | i64.extend32_s
  | i32.wrap_i64
  | i64.extend_i32_sx
  | inn.trunc_fmm_sx
  | inn.trunc_sat_fmm_sx
  | f32.demote_f64
  | f64.promote_f32
  | fnn.convert_imm_sx
  | inn.reinterpret_fnn
  | fnn.reinterpret_inn

iunop ::=
    clz
  | ctz
  | popcnt

ibinop ::=
    add
  | sub
  | mul
  | div_sx
  | rem_sx
  | and
  | or
  | xor
  | shl
  | shr_sx
  | rotl
  | rotr

funop ::=
    abs
  | neg
  | sqrt
  | ceil
  | floor
  | trunc
  | nearest

fbinop ::=
    add
  | sub
  | mul
  | div
  | min
  | max
  | copysign

itestop ::=
    eqz

irelop ::=
    eq
  | ne
  | lt_sx
  | gt_sx
  | le_sx
  | ge_sx

frelop ::=
    eq
  | ne
  | lt
  | gt
  | le
  | ge

//

unop ::=
    iunop
  | funop
  | extendN_s

binop ::=
    ibinop
  | fbinop

testop ::=
    itestop

relop ::=
    irelop
  | frelop

cvtop ::=
    wrap
  | extend
  | trunc
  | trunc_sat
  | convert
  | demote
  | promote
  | reinterpret
*/
package exprs
