package ast

import (
	"fmt"
	"go/ast"
	"io"
)

func Dump(w io.Writer, node ast.Node) {
	var rec func(*TreeNode, int)
	rec = func(tn *TreeNode, lvl int) {
		for i := 0; i < lvl; i++ {
			_, _ = w.Write([]byte("  "))
		}
		_, _ = fmt.Fprintf(w, "%#v\n", tn.Node)
		for _, c := range tn.Children {
			rec(c, lvl+1)
		}
	}
	rec(MakeTree(node), 0)
}
