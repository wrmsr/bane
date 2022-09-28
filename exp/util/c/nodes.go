package c

//

type Node interface {
	isNode()
}

type node struct{}

func (n node) isNode() {}
