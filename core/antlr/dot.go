package antlr

import (
	"fmt"
	"reflect"

	antlr "github.com/wrmsr/bane/core/antlr/runtime"

	"github.com/wrmsr/bane/core/check"
	"github.com/wrmsr/bane/core/graphs/dot"
)

func Dot(obj any) dot.Graph {
	g := dot.Graph{
		//Stmts: []dot.Stmt{
		//	dot.AsRawStmt("rankdir=LR;"),
		//},
	}

	idx := BuildIndex(obj)
	ns := make(map[any]string, len(idx))
	for k, v := range idx {
		ns[k] = fmt.Sprintf("_%d", v)
	}

	var rec func(any, any)
	rec = func(obj, parent any) {
		if o, ok := obj.(antlr.ParserRuleContext); ok {
			n := check.NotEmptyStr(ns[o])
			g.Stmts = append(g.Stmts, dot.NewNode(n).SetAttr("label", []any{
				[]any{reflect.TypeOf(o).String()},
				[]any{n},
				[]any{fmt.Sprintf("%d %d", o.GetStart().GetStop(), o.GetStop().GetStop())},
			}))
			if parent != nil {
				g.Stmts = append(g.Stmts, dot.NewEdge(ns[parent], n))
			}
			for _, c := range o.GetChildren() {
				rec(c, o)
			}
			return
		}

		if _, ok := obj.(antlr.TerminalNode); ok {
			return
		}

		panic(obj)
	}
	rec(obj, nil)

	return g
}
