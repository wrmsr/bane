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
	left  Node
	right Node
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
		return "="
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
	case "=":
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
	o     CmpOp
	left  Node
	right Node
}

//

type CreateArray struct {
	node
	items []Node
}

//

type CreateObject struct {
	node
	fields map[string]Node
}

//

type Current struct {
	leaf
}

//

type ExprRef struct {
	node
	expr Node
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
	name string
	args []Node
}

//

type Index struct {
	leaf
	value int
}

//

type JsonLiteral struct {
	leaf
	text string
}

//

type Negate struct {
	node
	item Node
}

//

type Or struct {
	op
	left  Node
	right Node
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
	target Target
}

//

type Project struct {
	node
	child Node
}

//

type Property struct {
	leaf
	name string
}

//

type Selection struct {
	Node
	child Node
}

//

type Sequence struct {
	node
	items []Node
}

//

type Slice struct {
	leaf
	start opt.Optional[int]
	stop  opt.Optional[int]
	step  opt.Optional[int]
}

//

type String struct {
	leaf
	value Node
}
