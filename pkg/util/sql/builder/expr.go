package builder

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

type BinaryOp int8

const (
	InvalidBinaryOp BinaryOp = iota

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

func (o BinaryOp) String() string {
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
	panic(fmt.Errorf("invalid binary op: %d", o))
}

func (o BinaryOp) IsBoolean() bool {
	switch o {
	case
		AndOp,
		OrOp:
		return true
	}
	return false
}

type BinaryExpr struct {
	expr
	Op   BinaryOp
	Args []Expr
}

func NewBinaryExpr(op BinaryOp, args ...Expr) BinaryExpr {
	return BinaryExpr{
		Op:   op,
		Args: check.NotEmptySlice(args),
	}
}

func NewBinaryExprOrSelf(op BinaryOp, args ...Expr) Expr {
	if len(args) == 1 {
		return args[0]
	}
	return NewBinaryExpr(
		op,
		args...,
	)
}

func Add(args ...Expr) Expr { return NewBinaryExprOrSelf(AddOp, args...) }
func Sub(args ...Expr) Expr { return NewBinaryExprOrSelf(SubOp, args...) }
func Mul(args ...Expr) Expr { return NewBinaryExprOrSelf(MulOp, args...) }
func Div(args ...Expr) Expr { return NewBinaryExprOrSelf(DivOp, args...) }

func Lt(args ...Expr) Expr  { return NewBinaryExprOrSelf(LtOp, args...) }
func Lte(args ...Expr) Expr { return NewBinaryExprOrSelf(LteOp, args...) }
func Eq(args ...Expr) Expr  { return NewBinaryExprOrSelf(EqOp, args...) }
func Ne(args ...Expr) Expr  { return NewBinaryExprOrSelf(NeOp, args...) }
func Gt(args ...Expr) Expr  { return NewBinaryExprOrSelf(GtOp, args...) }
func Gte(args ...Expr) Expr { return NewBinaryExprOrSelf(GteOp, args...) }

func And(args ...Expr) Expr { return NewBinaryExprOrSelf(AndOp, args...) }
func Or(args ...Expr) Expr  { return NewBinaryExprOrSelf(OrOp, args...) }

//

type Lit struct {
	expr
	String string
}

func NewLiteral(s string) Lit {
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
