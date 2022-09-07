/*
 */
package python

import (
	"fmt"
	"strings"

	iou "github.com/wrmsr/bane/pkg/util/io"
)

func Render(w *iou.IndentWriter, n Node) {
	switch n := n.(type) {

	case And:
		for i, c := range n.Children {
			if i > 0 {
				w.WriteString(" and ")
			}
			Render(w, c)
		}

	case Arith:
		Render(w, n.Left)
		w.WriteString(" ")
		w.WriteString(n.Op.String())
		w.WriteString(" ")
		Render(w, n.Right)

	case Attr:
		w.WriteString("(")
		w.WriteString(").")
		w.WriteString(n.Attr)

	case Brackets:
		w.WriteString("[")
		for i, c := range n.Children {
			if i > 0 {
				w.WriteString(", ")
			}
			Render(w, c)
		}
		w.WriteString("]")

	case Cmp:
		Render(w, n.Left)
		for _, c := range n.Items {
			w.WriteString(" ")
			w.WriteString(c.Op.String())
			w.WriteString(" ")
			Render(w, c.Right)
		}

	case Dict:
		w.WriteString("{")
		for i, c := range n.Items {
			if i > 0 {
				w.WriteString(", ")
			}
			if c.IsStars {
				w.WriteString("**")
				Render(w, c.K)
			} else {
				Render(w, c.K)
				w.WriteString(": ")
				Render(w, c.V)
			}
		}
		w.WriteString("}")

	case Ellipsis:
		w.WriteString("...")

	case False:
		w.WriteString("False")

	case None:
		w.WriteString("None")

	case Not:
		w.WriteString("(not ")
		Render(w, n.Child)
		w.WriteString(")")

	case Number:
		w.WriteString(n.S)

	case Or:
		for i, c := range n.Children {
			if i > 0 {
				w.WriteString(" or ")
			}
			Render(w, c)
		}

	case Parens:
		w.WriteString("(")
		for i, c := range n.Children {
			if i > 0 {
				w.WriteString(", ")
			}
			Render(w, c)
		}
		w.WriteString(")")

	case Set:
		w.WriteString("{")
		for i, c := range n.Children {
			if i > 0 {
				w.WriteString(", ")
			}
			Render(w, c)
		}
		w.WriteString("}")

	case Slice:
		if n.Start != nil {
			Render(w, n.Start)
		}
		w.WriteString(":")
		if n.Stop != nil {
			Render(w, n.Stop)
		}
		if n.Stop != nil || n.Step != nil {
			w.WriteString(":")
		}
		if n.Step != nil {
			Render(w, n.Step)
		}

	case String:
		for _, s := range n.S {
			w.WriteString(s)
		}

	case Subscript:
		Render(w, n.Child)
		w.WriteString("[")
		for i, s := range n.Items {
			if i > 0 {
				w.WriteString(", ")
			}
			Render(w, s)
		}
		w.WriteString("]")

	case Star:
		w.WriteString("*")
		Render(w, n.Child)

	case True:
		w.WriteString("True")

	case Unary:
		w.WriteString(n.Op.String())
		Render(w, n.Child)

	default:
		panic(fmt.Errorf("unhandled node type: %T", n))
	}
}

func RenderString(n Node) string {
	var sb strings.Builder
	Render(iou.NewIndentWriter(iou.NewDiscardStringWriter(&sb), "\t"), n)
	return sb.String()
}
