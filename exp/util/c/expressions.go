package c

import "fmt"

//

type Expression interface {
	Node
	isExpression()
}

type expression struct {
	node
}

func (e expression) isExpression() {}

//

type BinaryOp int8

const (
	InvalidBinaryOp BinaryOp = iota

	AddOp
	SubOp
	MulOp
	DivOp
	ModOp

	LtOp
	LteOp
	EqOp
	NeOp
	GtOp
	GteOp

	BitAndOp
	BitOrOp
	BitXorOp

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
	case ModOp:
		return "%"

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

	case BitAndOp:
		return "&"
	case BitOrOp:
		return "|"
	case BitXorOp:
		return "^"

	case AndOp:
		return "&&"
	case OrOp:
		return "||"

	}
	panic(fmt.Errorf("invalid binary op: %d", o))
}

type Binary struct {
	expression
	Op BinaryOp
	L  Expression
	R  Expression
}

//

type Call struct {
	expression
	Fn   Expression
	Args []Expression
}
