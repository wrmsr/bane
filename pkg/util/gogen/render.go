package gogen

import (
	"fmt"
	"strconv"
	"strings"

	iou "github.com/wrmsr/bane/pkg/util/io"
)

func renderFunc(w *iou.IndentWriter, n Func) {
	w.WriteString("func")

	if n.Receiver.Present() {
		w.WriteString(" (")
		Render(w, n.Receiver.Value())
		w.WriteString(")")
	}

	if n.Name.Present() {
		w.WriteString(" ")
		Render(w, n.Name.Value())
	}

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

	if n.Body.Present() {
		w.WriteString(" ")
		Render(w, n.Body.Value())
	}
}

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

func renderConst(w *iou.IndentWriter, n Const) {
	Render(w, n.Name)
	w.WriteString(" ")
	Render(w, n.Value)
	w.WriteString("\n")
}

func renderVar(w *iou.IndentWriter, n Var) {
	Render(w, n.Name)
	w.WriteString(" ")
	Render(w, n.Type)
	w.WriteString("\n")
}

func Render(w *iou.IndentWriter, n Node) {
	switch n := n.(type) {

	// base

	// decl

	case Func:
		renderFunc(w, n)
		w.WriteString("\n")

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
		if n.Name.Present() {
			Render(w, n.Name.Value())
			w.WriteString(" ")
		}
		Render(w, n.Type)

	case Struct:
		w.WriteString("type ")
		Render(w, n.Name)
		w.WriteString(" struct {\n")
		w.Indent()
		for _, f := range n.Fields {
			Render(w, f)
		}
		w.Dedent()
		w.WriteString("}\n")

	case StructField:
		Render(w, n.Name)
		w.WriteString(" ")
		Render(w, n.Type)
		w.WriteString("\n")

	case StmtDecl:
		Render(w, n.Stmt)
		w.WriteString("\n")

	// expr

	case Addr:
		w.WriteString("&")
		Render(w, n.Value)

	case Call:
		Render(w, n.Func)
		w.WriteString("(")
		for i, a := range n.Args {
			if i > 0 {
				w.WriteString(", ")
			}
			Render(w, a)
		}
		w.WriteString(")")

	case Deref:
		w.WriteString("*")
		Render(w, n.Value)

	case Field:
		Render(w, n.Value)
		for _, a := range n.Names {
			w.WriteString(".")
			Render(w, a)
		}

	case FuncExpr:
		renderFunc(w, n.Func)

	case Ident:
		w.WriteString(n.Name)

	case Index:
		Render(w, n.Value)
		w.WriteString("[")
		Render(w, n.Index)
		w.WriteString("]")

	case InfixExpr:
		for i, a := range n.Args {
			if i > 0 {
				w.WriteString(" ")
				w.WriteString(n.Op.String())
				w.WriteString(" ")
			}
			Render(w, a)
		}

	case Lit:
		w.WriteString(n.String)

	case Paren:
		w.WriteString("(")
		Render(w, n.Value)
		w.WriteString(")")

	case TypeAssert:
		Render(w, n.Value)
		w.WriteString(".(")
		Render(w, n.Type)
		w.WriteString(")")

	case UnaryExpr:
		w.WriteString(n.Op.String())
		w.WriteString(" ")
		Render(w, n.Arg)

	// raw

	case Raw:
		w.WriteString(n.Raw)

	// stmt

	case Assign:
		Render(w, n.Var)
		w.WriteString(" = ")
		Render(w, n.Value)
		w.WriteString("\n")

	case Block:
		w.WriteString("{\n")
		w.Indent()
		for _, s := range n.Body {
			Render(w, s)
		}
		w.Dedent()
		w.WriteString("}")

	case Const:
		w.WriteString("const ")
		renderConst(w, n)

	case Consts:
		w.WriteString("const (\n")
		w.Indent()
		for _, c := range n.Consts {
			renderConst(w, c)
		}
		w.Dedent()
		w.WriteString(")\n\n")

	case ExprStmt:
		Render(w, n.Expr)
		w.WriteString("\n")

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

	case ShortVar:
		Render(w, n.Name)
		w.WriteString(" := ")
		Render(w, n.Value)
		w.WriteString("\n")

	case Var:
		w.WriteString("var ")
		renderVar(w, n)

	case Vars:
		w.WriteString("var (\n")
		w.Indent()
		for _, v := range n.Vars {
			renderVar(w, v)
		}
		w.Dedent()
		w.WriteString(")\n\n")

	// type

	case Array:
		w.WriteString("[")
		if n.Len.Present() {
			w.WriteString(strconv.Itoa(n.Len.Value()))
		} else {
			w.WriteString("...")
		}
		w.WriteString("]")
		Render(w, n.Elem)

	case FuncType:
		renderFunc(w, n.Func)

	case Map:
		w.WriteString("map[")
		Render(w, n.Key)
		w.WriteString("]")
		Render(w, n.Value)

	case NameType:
		Render(w, n.Name)

	case Ptr:
		w.WriteString("*")
		Render(w, n.Elem)

	case Slice:
		w.WriteString("[]")
		Render(w, n.Elem)

	//

	default:
		panic(fmt.Errorf("unhandled node type: %T", n))
	}
}

func RenderString(n Node) string {
	var sb strings.Builder
	Render(iou.NewIndentWriter(iou.NewDiscardStringWriter(&sb), "\t"), n)
	return sb.String()
}
