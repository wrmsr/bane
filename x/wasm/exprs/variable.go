//
/*
instr ::= ...
  | local.get localidx
  | local.set localidx
  | local.tee localidx
  | global.get globalidx
  | global.set globalidx
*/
package exprs

//

type LocalGet struct {
	expr
	I int
}

var _ Expr = LocalGet{}
