package gogen

import (
	"fmt"

	iou "github.com/wrmsr/bane/pkg/util/io"
)

func renderImport(w *iou.IndentWriter, n Import) {
	if n.Alias.Present() {
		Render(w, n.Alias.Value())
		w.WriteString(" ")
	}
	w.WriteString("\"")
	w.WriteString(n.Spec)
	w.WriteString("\"")
	w.WriteString("\n")
}

func Render(w *iou.IndentWriter, n Node) {
	switch n := n.(type) {

	// base

	case Type:
		Render(w, n.Name)

	// decl

	case Import:
		w.WriteString("import ")
		renderImport(w, n)
		w.WriteString("\n")

	case Imports:
		w.WriteString("import (\n")
		w.Indent()
		for _, i := range n.Imports {
			renderImport(w, i)
		}
		w.Dedent()
		w.WriteString(")\n")

	case Param:
		Render(w, n.Name)
		w.WriteString(" ")
		Render(w, n.Type)

	case Func:
		w.WriteString("func ")
		Render(w, n.Name)
		w.WriteString("(")
		for i, p := range n.Params {
			if i > 0 {
				w.WriteString(", ")
			}
			Render(w, p)
		}
		w.WriteString(")")
		if n.Type.Present() {
			Render(w, n.Type.Value())
		}
		Render(w, n.Body)

	// expr

	case Lit:
		w.WriteString(n.String)

	case Ident:
		w.WriteString(n.Name)

	case InfixExpr:
		for i, a := range n.Args {
			if i > 0 {
				w.WriteString(" ")
				w.WriteString(n.Op.String())
				w.WriteString(" ")
			}
			Render(w, a)
		}

	case UnaryExpr:
		w.WriteString(n.Op.String())
		w.WriteString(" ")
		Render(w, n.Arg)

	// raw

	// stmt

	case ExprStmt:
		Render(w, n.Expr)
		w.WriteString("\n")

	case Block:
		w.WriteString("{\n")
		w.Indent()
		for _, s := range n.Body {
			Render(w, s)
		}
		w.Dedent()
		w.WriteString("}")

	case If:
		w.WriteString("if ")
		Render(w, n.Cond)
		w.WriteString(" ")
		Render(w, n.Then)
		if !n.Else.Empty() {
			w.WriteString("else ")
			Render(w, n.Else)
		}
		w.WriteString("\n\n")

	//

	default:
		panic(fmt.Errorf("unhandled node type: %T", n))
	}
}
