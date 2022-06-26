package gogen

import "github.com/wrmsr/bane/pkg/util/check"

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

type Block struct {
	stmt
	Body []Stmt
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
	stmt
	Name  Ident
	Value Expr
}

func NewConst(name Ident, value Expr) Const {
	return Const{
		Name:  name,
		Value: value,
	}
}

//

type Consts struct {
	stmt
	Consts []Const
}

func NewConsts(consts []Const) Consts {
	return Consts{
		Consts: consts,
	}
}

//

type ExprStmt struct {
	stmt
	Expr Expr
}

func NewExprStmt(expr Expr) ExprStmt {
	return ExprStmt{
		Expr: check.NotNil(expr).(Expr),
	}
}

//

type If struct {
	stmt
	Cond Expr
	Then Block
	Else Block
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
	stmt
	Name  Ident
	Value Expr
}

func NewShortVar(name Ident, value Expr) ShortVar {
	return ShortVar{
		Name:  name,
		Value: value,
	}
}

//

type Var struct {
	stmt
	Name Ident
	Type Type
}

func NewVar(name Ident, type_ Type) Var {
	return Var{
		Name: name,
		Type: type_,
	}
}

//

type Vars struct {
	stmt
	Vars []Var
}

func NewVars(vars []Var) Vars {
	return Vars{
		Vars: vars,
	}
}
