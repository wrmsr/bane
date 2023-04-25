//
/*
memarg ::= {offset u32 , align u32}

ww ::=
    8
  | 16
  | 32
  | 64

instr ::= ...
  | inn.load memarg
  | fnn.load memarg
  | v128.load memarg
  | inn.store memarg
  | fnn.store memarg
  | v128.store memarg
  | inn.load8_sx memarg
  | inn.load16_sx memarg
  | i64.load32_sx memarg
  | inn.store8 memarg
  | inn.store16 memarg
  | i64.store32 memarg
  | v128.load8x8_sx memarg
  | v128.load16x4_sx memarg
  | v128.load32x2_sx memarg
  | v128.load32_zero memarg
  | v128.load64_zero memarg
  | v128.loadww_splat memarg
  | v128.loadww_lane memarg laneidx
  | v128.storeww_lane memarg laneidx
  | memory.size
  | memory.grow
  | memory.fill
  | memory.copy
  | memory.init dataidx
  | data.drop dataidx
*/
package exprs

//

type SetLocal struct {
	expr
	N string
	V Expr
}

var _ Expr = SetLocal{}
