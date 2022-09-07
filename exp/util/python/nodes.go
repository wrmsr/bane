package python

import (
	"fmt"
)

//

type Node interface {
	isNode()
}

type node struct{}

func (n node) isNode() {}

//

type And struct {
	Children []Node

	node
}

//

type Or struct {
	Children []Node

	node
}

//

type Not struct {
	Child Node

	node
}

//

type CmpOp int8

const (
	CmpInvalid CmpOp = iota
	CmpEq
	CmpNe
	CmpGt
	CmpGe
	CmpLt
	CmpLe
)

func (o CmpOp) String() string {
	switch o {
	case CmpEq:
		return "=="
	case CmpNe:
		return "!="
	case CmpGt:
		return ">"
	case CmpGe:
		return ">="
	case CmpLt:
		return "<"
	case CmpLe:
		return "<="
	}
	panic(fmt.Errorf("unhandled CmpOp: %d", o))
}

func ParseCmpOp(s string) CmpOp {
	switch s {
	case "==":
		return CmpEq
	case "!=":
		return CmpNe
	case ">":
		return CmpGt
	case ">=":
		return CmpGe
	case "<":
		return CmpLt
	case "<=":
		return CmpLe
	}
	panic(fmt.Errorf("unhandled CmpOp: %v", s))
}

type Cmp struct {
	Left   Node
	Ops    []CmpOp
	Rights []Node

	node
}

//

type ArithOp int8

const (
	ArithInvalid ArithOp = iota

	ArithAdd
	ArithSub
	ArithMul
	ArithDiv
	ArithMod
	ArithPow
	ArithFloorDiv
	ArithMatMul

	ArithAnd
	ArithOr
	ArithShl
	ArithShr
)

func (o ArithOp) String() string {
	switch o {

	case ArithAdd:
		return "+"
	case ArithSub:
		return "-"
	case ArithMul:
		return "*"
	case ArithDiv:
		return "/"
	case ArithMod:
		return "%"
	case ArithPow:
		return "**"
	case ArithFloorDiv:
		return "//"
	case ArithMatMul:
		return "@"

	case ArithAnd:
		return "&"
	case ArithOr:
		return "|"
	case ArithShl:
		return "<<"
	case ArithShr:
		return ">>"

	}
	panic(fmt.Errorf("invalid ArithOp: %d", o))
}

func ParseArithOp(s string) ArithOp {
	switch s {

	case "+":
		return ArithAdd
	case "-":
		return ArithSub
	case "*":
		return ArithMul
	case "/":
		return ArithDiv
	case "%":
		return ArithMod
	case "**":
		return ArithPow
	case "//":
		return ArithFloorDiv
	case "@":
		return ArithMatMul

	case "&":
		return ArithAnd
	case "|":
		return ArithOr
	case "<<":
		return ArithShl
	case ">>":
		return ArithShr

	}
	panic(fmt.Errorf("unhandled ArithOp: %v", s))
}

type Arith struct {
	Op    ArithOp
	Left  Node
	Right Node

	node
}

//

type UnaryOp int8

const (
	UnaryInvalid UnaryOp = iota

	UnaryPlus
	UnaryMinus
)

func (o UnaryOp) String() string {
	switch o {
	case UnaryPlus:
		return "+"
	case UnaryMinus:
		return "-"
	}
	panic(fmt.Errorf("invalid UnaryOp: %d", o))
}

func ParseUnaryOp(s string) UnaryOp {
	switch s {
	case "+":
		return UnaryPlus
	case "-":
		return UnaryMinus
	}
	panic(fmt.Errorf("unhandled UnaryOp: %v", s))
}

type Unary struct {
	Op    UnaryOp
	Child Node

	node
}

//

type Parens struct {
	Children []Node

	node
}

//

type Brackets struct {
	Children []Node

	node
}

//

type DictItem struct {
	IsStars bool

	K Node
	V Node
}

type Dict struct {
	Items []DictItem

	node
}

//

type Set struct {
	Children []Node

	node
}

//

type Slice struct {
	Start Node
	Stop  Node
	Step  Node

	node
}

//

type Subscript struct {
	Child Node
	Items []Node

	node
}

//

type Star struct {
	Child Node

	node
}

//

type Attr struct {
	Child Node
	Attr  string

	node
}

//

type String struct {
	S []string

	node
}

//

type Number struct {
	S string

	node
}

//

type True struct {
	node
}

type False struct {
	node
}

type None struct {
	node
}

type Ellipsis struct {
	node
}
