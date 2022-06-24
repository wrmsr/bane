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

	// raw

	case *Raw:
		w.WriteString(n.Raw)

	// relation

	case *Table:
		Render(w, n.Identifier)

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

		if n.Identifier != nil {
			w.WriteString("AS ")
			Render(w, n.Identifier)
		}

	default:
		panic(fmt.Errorf("unhandled node type: %T", n))
	}
}
