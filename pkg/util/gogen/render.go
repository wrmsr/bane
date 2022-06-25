package gogen

import (
	"fmt"
	"strconv"

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
		w.WriteString(")\n\n")

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
		w.WriteString("\n")

	case StmtDecl:
		Render(w, n.Stmt)
		w.WriteString("\n")

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

	case Addr:
		w.WriteString("&")
		Render(w, n.Value)

	case Deref:
		w.WriteString("*")
		Render(w, n.Value)

	// raw

	case Raw:
		w.WriteString(n.Raw)

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

	case Var:
		w.WriteString("var ")
		Render(w, n.Name)
		w.WriteString(" ")
		Render(w, n.Type)

	case ShortVar:
		Render(w, n.Name)
		w.WriteString(" := ")
		Render(w, n.Value)

	// type

	case NameType:
		Render(w, n.Name)

	case Slice:
		w.WriteString("[]")
		Render(w, n.Elem)

	case Array:
		w.WriteString("[")
		if n.Len.Present() {
			w.WriteString(strconv.Itoa(n.Len.Value()))
		} else {
			w.WriteString("...")
		}
		w.WriteString("]")
		Render(w, n.Elem)

	case Ptr:
		w.WriteString("*")
		Render(w, n.Elem)

	case Map:
		w.WriteString("map[")
		Render(w, n.Key)
		w.WriteString("]")
		Render(w, n.Value)

	//

	default:
		panic(fmt.Errorf("unhandled node type: %T", n))
	}
}
