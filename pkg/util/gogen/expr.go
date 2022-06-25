package gogen

import (
	"fmt"

	"github.com/wrmsr/bane/pkg/util/check"
)

//

type Expr interface {
	Node
	isExpr()
}

type expr struct {
	node
}

func (e expr) isExpr() {}

//

type Lit struct {
	expr
	String string
}

func NewLit(s string) Lit {
	return Lit{
		String: check.NotZero(s),
	}
}

//

type Ident struct {
	expr
	Name string
}

func NewIdent(name string) Ident {
	return Ident{
		Name: check.NotZero(name),
	}
}

//

type InfixOp int8

const (
	InvalidInfixOp InfixOp = iota

	AddOp
	SubOp
	MulOp
	DivOp

	LtOp
	LteOp
	EqOp
	NeOp
	GtOp
	GteOp

	AndOp
	OrOp
)

func (o InfixOp) String() string {
	switch o {

	case AddOp:
		return "+"
	case SubOp:
		return "-"
	case MulOp:
		return "*"
	case DivOp:
		return "/"

	case LtOp:
		return "<"
	case LteOp:
		return "<="
	case EqOp:
		return "="
	case NeOp:
		return "!="
	case GtOp:
		return ">"
	case GteOp:
		return ">="

	case AndOp:
		return "AND"
	case OrOp:
		return "OR"

	}
	panic(fmt.Errorf("invalid infix op: %d", o))
}

func (o InfixOp) IsBoolean() bool {
	switch o {
	case
		AndOp,
		OrOp:
		return true
	}
	return false
}

type InfixExpr struct {
	expr
	Op   InfixOp
	Args []Expr
}

func NewInfixExpr(op InfixOp, args ...Expr) InfixExpr {
	return InfixExpr{
		Op:   op,
		Args: check.NotEmptySlice(args),
	}
}

func NewInfixExprOrSelf(op InfixOp, args ...Expr) Expr {
	if len(args) == 1 {
		return args[0]
	}
	return NewInfixExpr(
		op,
		args...,
	)
}

func Add(args ...Expr) Expr { return NewInfixExprOrSelf(AddOp, args...) }
func Sub(args ...Expr) Expr { return NewInfixExprOrSelf(SubOp, args...) }
func Mul(args ...Expr) Expr { return NewInfixExprOrSelf(MulOp, args...) }
func Div(args ...Expr) Expr { return NewInfixExprOrSelf(DivOp, args...) }

func Lt(args ...Expr) Expr  { return NewInfixExprOrSelf(LtOp, args...) }
func Lte(args ...Expr) Expr { return NewInfixExprOrSelf(LteOp, args...) }
func Eq(args ...Expr) Expr  { return NewInfixExprOrSelf(EqOp, args...) }
func Ne(args ...Expr) Expr  { return NewInfixExprOrSelf(NeOp, args...) }
func Gt(args ...Expr) Expr  { return NewInfixExprOrSelf(GtOp, args...) }
func Gte(args ...Expr) Expr { return NewInfixExprOrSelf(GteOp, args...) }

func And(args ...Expr) Expr { return NewInfixExprOrSelf(AndOp, args...) }
func Or(args ...Expr) Expr  { return NewInfixExprOrSelf(OrOp, args...) }

//

type UnaryOp int8

const (
	InvalidUnaryOp UnaryOp = iota

	NotOp
)

func (o UnaryOp) String() string {
	switch o {

	case NotOp:
		return "NOT"

	}
	panic(fmt.Errorf("invalid unary op: %d", o))
}

type UnaryExpr struct {
	expr
	Op  UnaryOp
	Arg Expr
}

func NewUnaryExpr(op UnaryOp, arg Expr) UnaryExpr {
	return UnaryExpr{
		Op:  op,
		Arg: check.NotNil(arg).(Expr),
	}
}

func Not(arg Expr) Expr { return NewUnaryExpr(NotOp, arg) }
