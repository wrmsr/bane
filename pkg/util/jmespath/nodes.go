package jmespath

import (
	"fmt"

	opt "github.com/wrmsr/bane/pkg/util/optional"
)

// TODO: //go:generate go run github.com/wrmsr/bane/pkg/util/trees/dev/gen

//

type Node interface {
	isNode()
}

type node struct{}

func (n node) isNode() {}

//

type Op interface {
	Node
	isOp()
}

type op struct {
	node
}

func (o op) isOp() {}

//

type Leaf interface {
	Node
	isLeaf()
}

type leaf struct {
	node
}

func (l leaf) isLeaf() {}

//

type And struct {
	op
	Left  Node
	Right Node
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
	op
	Op    CmpOp
	Left  Node
	Right Node
}

//

type CreateArray struct {
	node
	Items []Node
}

//

type CreateObject struct {
	node
	Fields map[string]Node
}

//

type Current struct {
	leaf
}

//

type ExprRef struct {
	node
	Expr Node
}

//

type FlattenArray struct {
	leaf
}

//

type FlattenObject struct {
	leaf
}

//

type Call struct {
	node
	Name string
	Args []Node
}

//

type Index struct {
	leaf
	Value int
}

//

type JsonLiteral struct {
	leaf
	Text string
}

//

type Negate struct {
	node
	Item Node
}

//

type Or struct {
	op
	Left  Node
	Right Node
}

//

type Target interface {
	isTarget()
}

type NumberTarget int
type NameTarget string

func (t NumberTarget) isTarget() {}
func (t NameTarget) isTarget()   {}

type Parameter struct {
	leaf
	Target Target
}

//

type Project struct {
	node
	Child Node
}

//

type Property struct {
	leaf
	Name string
}

//

type Selection struct {
	Node
	Child Node
}

//

type Sequence struct {
	node
	Items []Node
}

//

type Slice struct {
	leaf
	Start opt.Optional[int]
	Stop  opt.Optional[int]
	Step  opt.Optional[int]
}

//

type String struct {
	leaf
	Value string
}
