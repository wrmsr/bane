//
/*
blocktype ::=
    typeidx
  | valtype?

instr ::= ...
  | nop
  | unreachable
  | block blocktype instr* end
  | loop blocktype instr* end
  | if blocktype instr* else instr* end
  | br labelidx
  | br_if labelidx
  | br_table vec(labelidx) labelidx
  | return
  | call funcidx
  | call_indirect tableidx typeidx
*/
package exprs

//

type Block struct {
	expr
	Id string
	S  []Expr
}

var _ Expr = Block{}

//

type Nop struct {
	expr
}

var _ Expr = Nop{}

//

type Select struct {
	expr
	C Expr
	T Expr
	F Expr
}

var _ Expr = Select{}

//

type Unreachable struct {
	expr
}

var _ Expr = Unreachable{}
