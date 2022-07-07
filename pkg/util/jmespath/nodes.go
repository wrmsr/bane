package jmespath

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

type Cmp struct {
	op
	c     CmpOp
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

type ExpressionRef struct {
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

/*
//

class FunctionCall(Node):
    name: str
    args: ta.Sequence[Node]


class Index(Leaf):
    value: int


class JsonLiteral(Leaf):
    text: str


class Negate(Node):
    item: Node


class Or(Operator):
    left: Node
    right: Node


class Parameter(Leaf):
    class Target(dc.Enum, sealed=True):
        pass

    class NumberTarget(Target):
        value: int

    class NameTarget(Target):
        value: str

    target: Target


class Project(Node):
    child: Node


class Property(Leaf):
    name: str


class Selection(Node):
    child: Node


class Sequence(Node):
    items: ta.Sequence[Node]


class Slice(Leaf):
    start: ta.Optional[int]
    stop: ta.Optional[int]
    step: ta.Optional[int]


class String(Leaf):
    value: Node
*/
