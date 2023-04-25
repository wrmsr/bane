package exprs

import (
	"fmt"
	"strings"

	iou "github.com/wrmsr/bane/pkg/util/io"
)

//

type Node interface {
	isNode()

	Render(sb *strings.Builder)
}

type node struct{}

func (n node) isNode() {}

//

type Const struct {
	node
	V any
}

var _ Node = Const{}

func NewConst(v any) Const {
	return Const{V: v}
}

//

type Var struct {
	node
	N string
}

var _ Node = Var{}

func NewVar(n string) Var {
	return Var{N: n}
}

//

type Op struct {
	node
	N  string
	Vs []Node
}

var _ Node = Op{}

func NewOp(n string, vs ...Node) Op {
	return Op{N: n, Vs: vs}
}

//

func (n Const) Render(sb *strings.Builder) {
	iou.Discard(fmt.Fprintf(sb, "%v", n.V))
}

func (n Var) Render(sb *strings.Builder) {
	iou.Discard(fmt.Fprintf(sb, "$%s", n.N))
}

func (n Op) Render(sb *strings.Builder) {
	for i, v := range n.Vs {
		if i > 0 {
			iou.Discard(sb.WriteString(n.N))
		}
		v.Render(sb)
	}
}
