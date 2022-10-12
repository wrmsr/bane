package ast

import (
	"go/ast"

	"golang.org/x/tools/go/ast/astutil"
)

type TreeNode struct {
	Node     ast.Node
	Parent   *TreeNode
	Children []*TreeNode
}

func MakeTree(root ast.Node) *TreeNode {
	m := map[ast.Node]*TreeNode{}
	astutil.Apply(root, func(cursor *astutil.Cursor) bool {
		n := cursor.Node()
		tn := &TreeNode{
			Node: n,
		}
		if len(m) > 0 {
			pn := cursor.Parent()
			ptn := m[pn]
			if ptn == nil {
				panic(pn)
			}
			tn.Parent = ptn
			ptn.Children = append(ptn.Children, tn)
		}
		m[n] = tn
		return true
	}, nil)
	return m[root]
}
