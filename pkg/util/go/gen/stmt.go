package gen

import (
	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/slices"
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

func AssignOf(var_, value any) Assign {
	return Assign{
		Var:   ExprOf(var_),
		Value: ExprOf(value),
	}
}

//

type Blank struct {
	stmt
}

//

type Block struct {
	Body []Stmt

	stmt
}

func NewBlock(body ...Stmt) *Block {
	return &Block{
		Body: body,
	}
}

func BlockOf(body ...Stmt) Block {
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

func ConstOf(name Ident, value Expr) Const {
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

func ConstsOf(consts []Const) Consts {
	return Consts{
		Consts: consts,
	}
}

//

type ExprStmt struct {
	Expr Expr

	stmt
}

func ExprStmtOf(expr Expr) ExprStmt {
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

func IfElseOf(cond Expr, then, else_ Block) If {
	return If{
		Cond: check.NotNil(cond).(Expr),
		Then: then,
		Else: else_,
	}
}

func IfOf(cond Expr, then Block) If {
	return IfElseOf(cond, then, BlockOf())
}

//

type ShortVar struct {
	Name  Ident
	Value Expr

	stmt
}

func ShortVarOf(name Ident, value Expr) ShortVar {
	return ShortVar{
		Name:  name,
		Value: value,
	}
}

//

type Var struct {
	Name  Ident
	Type  Type
	Value Expr

	stmt
}

func VarOf(name any, args ...any) Var {
	v := Var{
		Name:  IdentOf(name),
		Type:  slices.TryShiftAs[Type](&args),
		Value: slices.TryShiftAs[Expr](&args),
	}
	check.EmptySlice(args)
	return v
}

//

type Vars struct {
	Vars []Var

	stmt
}

func VarsOf(vars ...Var) Vars {
	return Vars{
		Vars: vars,
	}
}
