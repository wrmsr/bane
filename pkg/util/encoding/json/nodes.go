package json

import bt "github.com/wrmsr/bane/pkg/util/types"

//

type Node interface {
	isNode()
}

type node struct{}

func (n node) isNode() {}

//

type Pair = bt.Kv[string, Node]

type Object struct {
	node
	Pairs []Pair
}

//

type Array struct {
	node
	Values []Node
}

//

type String struct {
	node
	S string
}

//

type Number struct {
	node
	S string
}

//

type True struct {
	node
}

type False struct {
	node
}

type Null struct {
	node
}
