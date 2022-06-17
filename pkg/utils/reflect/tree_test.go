package reflect

import (
	"fmt"
	"testing"
)

//

type Node struct{}

type NodePtr interface {
	Super() NodePtr
	Children() []NodePtr

	isNodePtr()
}

var _ NodePtr = &Node{}

func (n Node) Super() NodePtr      { return nil }
func (n Node) Children() []NodePtr { return nil }

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

func (n Num) Super() NodePtr { return &n.Node }

//

type BinOp struct {
	Node
	l, r NodePtr
}

func (n BinOp) Super() NodePtr      { return &n.Node }
func (n BinOp) Children() []NodePtr { return []NodePtr{n.l, n.r} }

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

func (n Add) Super() NodePtr { return &n.BinOp }

//

func FindNums(n NodePtr) {
	switch n := n.(type) {
	case *Num:
		fmt.Println(n.v)
	}

	for _, c := range n.Children() {
		FindNums(c)
	}
}

//

func TestTree(t *testing.T) {
	tree := NewAdd(
		NewAdd(
			NewNum(1),
			NewNum(2)),
		NewNum(3))

	fmt.Printf("%+v\n", tree)

	FindNums(tree)
}
