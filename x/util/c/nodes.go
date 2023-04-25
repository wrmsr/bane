package c

//

type Node interface {
	isNode()
}

type node struct{}

func (n node) isNode() {}

//

type TranslationUnit struct {
	node
	Ds []Declaration
}

//

type Identifier struct {
	expression
	S string
}

//

type Constant struct {
	expression
	S string
}

//

type StringLiteral struct {
	expression
	S []string
}
