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

func (e *expr) isExpr() {}

//

type Literal struct {
	expr
	String string
}

func NewLiteral(s string) *Literal {
	return &Literal{
		String: check.NotZero(s),
	}
}

//

type Identifier struct {
	expr
	Name string
}

func NewIdentifier(name string) *Identifier {
	return &Identifier{
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

type InfixExpr struct {
	expr
	Op   InfixOp
	Args []Expr
}

func NewInfixExpr(op InfixOp, args ...Expr) *InfixExpr {
	return &InfixExpr{
		Op:   op,
		Args: check.NotEmptySlice(args),
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

func NewUnaryExpr(op UnaryOp, arg Expr) *UnaryExpr {
	return &UnaryExpr{
		Op:  op,
		Arg: check.NotNil(arg).(Expr),
	}
}
