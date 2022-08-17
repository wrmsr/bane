package gen

import (
	"github.com/wrmsr/bane/pkg/util/check"
	opt "github.com/wrmsr/bane/pkg/util/optional"
)

//

type Stmt interface {
	Node
	isStmt()
}

type stmt struct {
	node
}

func (s stmt) isStmt() {}

//

type Assign struct {
	Var   Expr
	Value Expr

	stmt
}

func NewAssign(var_ Expr, value Expr) Assign {
	return Assign{
		Var:   var_,
		Value: value,
	}
}

//

type Blank struct {
	stmt
}

func NewBlank() Blank {
	return Blank{}
}

//

type Block struct {
	Body []Stmt

	stmt
}

func BlockOf(body ...Stmt) *Block {
	return &Block{
		Body: body,
	}
}

func NewBlock(body ...Stmt) Block {
	return Block{
		Body: body,
	}
}

func (b Block) Empty() bool {
	return len(b.Body) < 1
}

//

type Const struct {
	Name  Ident
	Value Expr

	stmt
}

func NewConst(name Ident, value Expr) Const {
	return Const{
		Name:  name,
		Value: value,
	}
}

//

type Consts struct {
	Consts []Const

	stmt
}

func NewConsts(consts []Const) Consts {
	return Consts{
		Consts: consts,
	}
}

//

type ExprStmt struct {
	Expr Expr

	stmt
}

func NewExprStmt(expr Expr) ExprStmt {
	return ExprStmt{
		Expr: check.NotNil(expr).(Expr),
	}
}

//

type If struct {
	Cond Expr
	Then Block
	Else Block

	stmt
}

func NewIfElse(cond Expr, then, else_ Block) If {
	return If{
		Cond: check.NotNil(cond).(Expr),
		Then: then,
		Else: else_,
	}
}

func NewIf(cond Expr, then Block) If {
	return NewIfElse(cond, then, NewBlock())
}

//

type ShortVar struct {
	Name  Ident
	Value Expr

	stmt
}

func NewShortVar(name Ident, value Expr) ShortVar {
	return ShortVar{
		Name:  name,
		Value: value,
	}
}

//

type Var struct {
	Name  Ident
	Type  opt.Optional[Type]
	Value opt.Optional[Expr]

	stmt
}

func NewVar(name Ident, type_ opt.Optional[Type], value opt.Optional[Expr]) Var {
	return Var{
		Name:  name,
		Type:  type_,
		Value: value,
	}
}

//

type Vars struct {
	Vars []Var

	stmt
}

func NewVars(vars []Var) Vars {
	return Vars{
		Vars: vars,
	}
}
