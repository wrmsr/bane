/*
Copyright (c) 2014 Caleb Spare

# MIT License

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
documentation files (the "Software"), to deal in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit
persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the
Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
package parse

import "fmt"

//

type Node interface {
	Position() *Pos
	String() string // A non-recursive string representation
	Children() []Node

	//copy() Node

	//doChildren(func(Node) bool) bool
	//editChildren(func(Node) Node)
}

func (p *Pos) Position() *Pos { return p }

//

type BoolNode struct {
	*Pos
	Val bool
}

func (n *BoolNode) String() string {
	if n.Val {
		return "true"
	}
	return "false"
}

func (n *BoolNode) Children() []Node { return nil }

//

type CharacterNode struct {
	*Pos
	Val  rune
	Text string
}

func (n *CharacterNode) String() string   { return fmt.Sprintf("char(%q)", n.Val) }
func (n *CharacterNode) Children() []Node { return nil }

//

type CommentNode struct {
	*Pos
	Text string
}

func (n *CommentNode) String() string   { return fmt.Sprintf("comment(%q)", n.Text) }
func (n *CommentNode) Children() []Node { return nil }

//

type DerefNode struct {
	*Pos
	Node Node
}

func (n *DerefNode) String() string   { return "deref" }
func (n *DerefNode) Children() []Node { return []Node{n.Node} }

//

type KeywordNode struct {
	*Pos
	Val string
}

func (n *KeywordNode) String() string   { return fmt.Sprintf("keyword(%s)", n.Val) }
func (n *KeywordNode) Children() []Node { return nil }

//

type ListNode struct {
	*Pos
	Nodes []Node
}

func (n *ListNode) String() string {
	return fmt.Sprintf("list(length=%d)", countSemantic(n.Nodes))
}
func (n *ListNode) Children() []Node { return n.Nodes }

//

type MapNode struct {
	*Pos
	Namespace string // empty unless the map has a namespace: #:ns{:x 1}
	Nodes     []Node
}

func (n *MapNode) String() string {
	var ns string
	if n.Namespace != "" {
		ns = fmt.Sprintf("ns=%s, ", n.Namespace)
	}
	semanticNodes := countSemantic(n.Nodes)
	return fmt.Sprintf("map(%slength=%d)", ns, semanticNodes/2)
}
func (n *MapNode) Children() []Node { return n.Nodes }

//

type MetadataNode struct {
	*Pos
	Node Node
}

func (n *MetadataNode) String() string   { return "metadata" }
func (n *MetadataNode) Children() []Node { return []Node{n.Node} }

//

type NewlineNode struct {
	*Pos
}

func (n *NewlineNode) String() string   { return "newline" }
func (n *NewlineNode) Children() []Node { return nil }

//

type NilNode struct {
	*Pos
}

func (n *NilNode) String() string   { return "nil" }
func (n *NilNode) Children() []Node { return nil }

//

type NumberNode struct {
	*Pos
	Val string
}

func (n *NumberNode) String() string   { return fmt.Sprintf("num(%s)", n.Val) }
func (n *NumberNode) Children() []Node { return nil }

//

type SymbolNode struct {
	*Pos
	Val string
}

func (n *SymbolNode) String() string   { return "sym(" + n.Val + ")" }
func (n *SymbolNode) Children() []Node { return nil }

//

type QuoteNode struct {
	*Pos
	Node Node
}

func (n *QuoteNode) String() string   { return "quote" }
func (n *QuoteNode) Children() []Node { return []Node{n.Node} }

//

type StringNode struct {
	*Pos
	Val string
}

func (n *StringNode) String() string   { return fmt.Sprintf("string(%q)", n.Val) }
func (n *StringNode) Children() []Node { return nil }

//

type SyntaxQuoteNode struct {
	*Pos
	Node Node
}

func (n *SyntaxQuoteNode) String() string   { return "syntax quote" }
func (n *SyntaxQuoteNode) Children() []Node { return []Node{n.Node} }

//

type UnquoteNode struct {
	*Pos
	Node Node
}

func (n *UnquoteNode) String() string   { return "unquote" }
func (n *UnquoteNode) Children() []Node { return []Node{n.Node} }

//

type UnquoteSpliceNode struct {
	*Pos
	Node Node
}

func (n *UnquoteSpliceNode) String() string   { return "unquote splice" }
func (n *UnquoteSpliceNode) Children() []Node { return []Node{n.Node} }

//

type VectorNode struct {
	*Pos
	Nodes []Node
}

func (n *VectorNode) String() string {
	return fmt.Sprintf("vector(length=%d)", countSemantic(n.Nodes))
}
func (n *VectorNode) Children() []Node { return n.Nodes }

//

type FnLiteralNode struct {
	*Pos
	Nodes []Node
}

func (n *FnLiteralNode) String() string {
	return fmt.Sprintf("lambda(length=%d)", countSemantic(n.Nodes))
}
func (n *FnLiteralNode) Children() []Node { return n.Nodes }

//

type ReaderCondNode struct {
	*Pos
	Nodes []Node
}

func (n *ReaderCondNode) String() string {
	return fmt.Sprintf("reader-cond(length=%d)", len(n.Nodes))
}
func (n *ReaderCondNode) Children() []Node { return n.Nodes }

//

type ReaderCondSpliceNode struct {
	*Pos
	Nodes []Node
}

func (n *ReaderCondSpliceNode) String() string {
	return fmt.Sprintf("reader-cond-splice(length=%d)", len(n.Nodes))
}
func (n *ReaderCondSpliceNode) Children() []Node { return n.Nodes }

//

type ReaderDiscardNode struct {
	*Pos
	Node Node
}

func (n *ReaderDiscardNode) String() string   { return "discard" }
func (n *ReaderDiscardNode) Children() []Node { return []Node{n.Node} }

//

type ReaderEvalNode struct {
	*Pos
	Node Node
}

func (n *ReaderEvalNode) String() string   { return "eval" }
func (n *ReaderEvalNode) Children() []Node { return []Node{n.Node} }

//

type RegexNode struct {
	*Pos
	Val string
}

func (n *RegexNode) String() string   { return fmt.Sprintf("regex(%q)", n.Val) }
func (n *RegexNode) Children() []Node { return nil }

//

type SetNode struct {
	*Pos
	Nodes []Node
}

func (n *SetNode) String() string {
	return fmt.Sprintf("set(length=%d)", countSemantic(n.Nodes))
}
func (n *SetNode) Children() []Node { return n.Nodes }

//

type VarQuoteNode struct {
	*Pos
	Val string
}

func (n *VarQuoteNode) String() string   { return fmt.Sprintf("varquote(%s)", n.Val) }
func (n *VarQuoteNode) Children() []Node { return nil }

//

type TagNode struct {
	*Pos
	Val string
}

func (n *TagNode) String() string   { return fmt.Sprintf("tag(%s)", n.Val) }
func (n *TagNode) Children() []Node { return nil }

//

func isSemantic(n Node) bool {
	switch n.(type) {
	case *CommentNode, *NewlineNode:
		return false
	}
	return true
}

func countSemantic(nodes []Node) int {
	count := 0
	for _, node := range nodes {
		if isSemantic(node) {
			count++
		}
	}
	return count
}

func panicf(format string, args ...any) {
	panic(fmt.Sprintf(format, args...))
}
