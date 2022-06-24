package builder

import (
	"fmt"

	iou "github.com/wrmsr/bane/pkg/util/io"
)

func Render(w iou.DiscardStringWriter, n Node) {
	switch n := n.(type) {

	// expr

	case *Literal:
		w.WriteString(n.String)

	case *Identifier:
		w.WriteString(n.Name)

	case *InfixExpr:
		for i, a := range n.Args {
			if i > 0 {
				w.WriteString(" ")
				w.WriteString(n.Op.String())
				w.WriteString(" ")
			}
			Render(w, a)
		}

	case *UnaryExpr:
		w.WriteString(n.Op.String())
		w.WriteString(" ")
		Render(w, n.Arg)

	// raw

	case *Raw:
		w.WriteString(n.Raw)

	// relation

	case *Table:
		Render(w, n.Name)

		if n.Alias != nil {
			w.WriteString(" AS ")
			Render(w, n.Alias)
		}

	// stmt

	case *Select:
		w.WriteString("SELECT ")

		for i, si := range n.Items {
			if i > 0 {
				w.WriteString(", ")
			}
			Render(w, si)
		}

		if n.From != nil {
			w.WriteString(" FROM ")
			Render(w, n.From)
		}

		if n.Where != nil {
			w.WriteString(" WHERE ")
			Render(w, n.Where)
		}

	case *SelectItem:
		Render(w, n.Expr)

		if n.Label != nil {
			w.WriteString(" AS ")
			Render(w, n.Label)
		}

	default:
		panic(fmt.Errorf("unhandled node type: %T", n))
	}
}
