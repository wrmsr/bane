package gen

import (
	"fmt"
	"strconv"
	"strings"

	iou "github.com/wrmsr/bane/core/io"
)

func renderFunc(w *iou.IndentWriter, n Func) {
	w.WriteString("func")

	if n.Receiver != nil {
		w.WriteString(" (")
		Render(w, *n.Receiver)
		w.WriteString(")")
	}

	if n.Name != nil {
		w.WriteString(" ")
		Render(w, *n.Name)
	}

	w.WriteString("(")
	for i, p := range n.Params {
		if i > 0 {
			w.WriteString(", ")
		}
		Render(w, p)
	}
	w.WriteString(")")

	if n.Type != nil {
		w.WriteString(" ")
		Render(w, n.Type)
	}

	if n.Body != nil {
		w.WriteString(" ")
		Render(w, *n.Body)
	}
}

func renderImport(w *iou.IndentWriter, n Import) {
	if n.Spec == "" {
		w.WriteString("\n")
		return
	}
	if n.Alias != nil {
		Render(w, *n.Alias)
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
	if n.Type != nil {
		w.WriteString(" ")
		Render(w, n.Type)
	}
	if n.Value != nil {
		w.WriteString(" = ")
		Render(w, n.Value)
	}
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

	case Imports:
		w.WriteString("import (\n")
		w.Indent()
		for _, i := range n.Imports {
			renderImport(w, i)
		}
		w.Dedent()
		w.WriteString(")\n")

	case Package:
		w.WriteString("package ")
		Render(w, n.Name)
		w.WriteString("\n")

	case Param:
		if n.Name != nil {
			Render(w, *n.Name)
			w.WriteString(" ")
		}
		Render(w, n.Type)

	case Struct:
		w.WriteString("type ")
		Render(w, n.Name)
		w.WriteString(" struct {")
		w.WriteString("\n")
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

	// expr

	case Addr:
		w.WriteString("&")
		Render(w, n.Value)

	case Binary:
		for i, a := range n.Args {
			if i > 0 {
				w.WriteString(" ")
				w.WriteString(n.Op.String())
				w.WriteString(" ")
			}
			Render(w, a)
		}

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

	case FuncExpr:
		renderFunc(w, n.Func)

	case Ident:
		w.WriteString(n.Name)

	case Index:
		Render(w, n.Value)
		w.WriteString("[")
		Render(w, n.Index)
		w.WriteString("]")

	case Lit:
		w.WriteString(n.String)

	case Paren:
		w.WriteString("(")
		Render(w, n.Value)
		w.WriteString(")")

	case Select:
		Render(w, n.Value)
		for _, a := range n.Names {
			w.WriteString(".")
			Render(w, a)
		}

	case TypeAssert:
		Render(w, n.Value)
		w.WriteString(".(")
		Render(w, n.Type)
		w.WriteString(")")

	case Unary:
		if n.Op.IsPrefix() {
			w.WriteString(n.Op.String())
		}
		Render(w, n.Arg)
		if !n.Op.IsPrefix() {
			w.WriteString(n.Op.String())
		}

	// raw

	case Raw:
		w.WriteString(n.Raw)

	// stmt

	case Assign:
		Render(w, n.Var)
		w.WriteString(" = ")
		Render(w, n.Value)
		w.WriteString("\n")

	case Blank:
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
		w.WriteString(")\n")

	case ExprStmt:
		Render(w, n.Expr)
		w.WriteString("\n")

	case Goto:
		w.WriteString("goto ")
		Render(w, n.Name)
		w.WriteString("\n")

	case If:
		w.WriteString("if ")
		Render(w, n.Cond)
		w.WriteString(" ")
		Render(w, n.Then)
		if !n.Else.Empty() {
			w.WriteString(" else ")
			Render(w, n.Else)
		}
		w.WriteString("\n")

	case Label:
		Render(w, n.Name)
		w.WriteString(":\n")

	case Return:
		w.WriteString("return")
		if n.Expr != nil {
			w.WriteString(" ")
			Render(w, n.Expr)
		}
		w.WriteString("\n")

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
		w.WriteString(")\n")

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
