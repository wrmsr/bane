package python

import "fmt"

//

type Node interface {
	isNode()
}

type node struct{}

func (n node) isNode() {}

//

type And struct {
	node
	Children []Node
}

//

type Or struct {
	node
	Children []Node
}

//

type Not struct {
	node
	Child Node
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

type Comparison struct {
	Op    CmpOp
	Left  Node
	Right Node
}
