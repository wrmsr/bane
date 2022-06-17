package reflect

import (
	"errors"
	"fmt"
	"testing"
)

//

type NodeVisitor[C, R any] interface {
	VisitNode(n *Node, c C) R
	VisitNum(n *Num, c C) R
	VisitBinOp(n *BinOp, c C) R
	VisitAdd(n *Add, c C) R
}

func VisitNode[C, R any](n NodePtr, v NodeVisitor[C, R], c C) R {
	if n == nil {
		panic(errors.New("cannot visit nil node"))
	}
	switch n := n.(type) {
	case *Node:
		return v.VisitNode(n, c)
	case *Num:
		return v.VisitNum(n, c)
	case *BinOp:
		return v.VisitBinOp(n, c)
	case *Add:
		return v.VisitAdd(n, c)
	}
	panic(fmt.Errorf("unhandled node type: %T", v))
}

//

type DefaultNodeVisitor[C, R any] struct{}

var _ NodeVisitor[any, any] = DefaultNodeVisitor[any, any]{}

//

type Node struct{}

type NodePtr interface {
	isNodePtr()
}

func (n *Node) isNodePtr() {}

func (d DefaultNodeVisitor[C, R]) VisitNode(n *Node, c C) R {
	return nil
}

//

type Num struct {
	Node
}

func (d DefaultNodeVisitor[C, R]) VisitNum(n *Num, c C) R {
	return d.VisitNode(&n.Node, c)
}

//

type BinOp struct {
	Node
}

func (d DefaultNodeVisitor[C, R]) VisitBinOp(n *BinOp, c C) R {
	return d.VisitNode(&n.Node, c)
}

//

type Add struct {
	BinOp
}

func (d DefaultNodeVisitor[C, R]) VisitAdd(n *Add, c C) R {
	return d.VisitBinOp(&n.BinOp, c)
}

//

func TestVisit(t *testing.T) {

}
