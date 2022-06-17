package reflect

import (
	"fmt"
	"testing"
)

//

type NodeVisitor interface {
	VisitNode(n *Node) any
	VisitNum(n *Num) any
	VisitBinOp(n *BinOp) any
	VisitAdd(n *Add) any
}

func VisitNode(n NodePtr, v NodeVisitor) any {
	switch n := n.(type) {
	case *Node:
		return v.VisitNode(n)
	case *Num:
		return v.VisitNum(n)
	case *BinOp:
		return v.VisitBinOp(n)
	case *Add:
		return v.VisitAdd(n)
	}
	panic(fmt.Errorf("unhandled node type: %T", n))
}

//

type NodeSuper struct {
	super NodeVisitor
}

func (v NodeSuper) VisitNode(n *Node) any {
	panic("unreachable")
}

var _ NodeVisitor = NodeSuper{}

//

type NilNodeSuper struct {
	NodeSuper
}

func (v NilNodeSuper) VisitNode(n *Node) any {
	return nil
}

//

type Node struct{}

type NodePtr interface {
	isNodePtr()
}

func (n *Node) isNodePtr() {}

//

type Num struct {
	Node
	v int
}

func NewNum(v int) *Num {
	return &Num{
		v: v,
	}
}

func (v NodeSuper) VisitNum(n *Num) any {
	return v.super.VisitNode(&n.Node)
}

//

type BinOp struct {
	Node
	l, r NodePtr
}

func (v NodeSuper) VisitBinOp(n *BinOp) any {
	return v.super.VisitNode(&n.Node)
}

//

type Add struct {
	BinOp
}

func NewAdd(l, r NodePtr) *Add {
	return &Add{
		BinOp: BinOp{
			l: l,
			r: r,
		},
	}
}

func (v NodeSuper) VisitAdd(n *Add) any {
	return v.super.VisitBinOp(&n.BinOp)
}

//

type NumFinder struct {
	NilNodeSuper
}

func NewNumFinder() *NumFinder {
	nf := &NumFinder{}
	nf.NilNodeSuper.super = nf
	return nf
}

//

func TestTree(t *testing.T) {
	tree := NewAdd(
		NewAdd(
			NewNum(1),
			NewNum(2)),
		NewNum(3))

	fmt.Printf("%+v\n", tree)

	VisitNode(tree, NewNumFinder())
}
