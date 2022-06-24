package dot

import (
	"html"

	iou "github.com/wrmsr/bane/pkg/util/io"
)

func Render(w iou.DiscardStringWriter, o any) {
	switch o := o.(type) {
	case Raw:
		w.WriteString(o.Raw)

	case Text:
		w.WriteString(html.EscapeString(o.Text))

	case Cell:
		w.WriteString("<td>")
		Render(w, o.Value)
		w.WriteString("</td>")

	case Row:
		w.WriteString("<tr>")
		for _, c := range o.Cells {
			Render(w, c)
		}
		w.WriteString("</tr>")

	case Table:
		w.WriteString("<table>")
		for _, r := range o.Rows {
			Render(w, r)
		}
		w.WriteString("</table>")

	case Id:
		w.WriteString("\"")
		w.WriteString(o.Id)
		w.WriteString("\"")

	case Attrs:
		if o.Attrs == nil {
			return
		}
		w.WriteString("[")
		i := 0
		for k, v := range o.Attrs {
			if i > 0 {
				w.WriteString(", ")
			}
			w.WriteString(k)
			w.WriteString("=<")
			Render(w, v)
			w.WriteString(">")
			i++
		}
		w.WriteString("]")

	case RawStmt:
		w.WriteString(o.Raw)
		w.WriteString("\n")

	case Edge:
		Render(w, o.Left)
		w.WriteString(" -> ")
		Render(w, o.Right)
		if len(o.Attrs.Attrs) > 0 {
			w.WriteString(" ")
			Render(w, o.Attrs)
		}
		w.WriteString(";\n")

	case Node:
		Render(w, o.Id)
		if len(o.Attrs.Attrs) > 0 {
			w.WriteString(" ")
			Render(w, o.Attrs)
		}
		w.WriteString(";\n")

	case Graph:
		w.WriteString("digraph ")
		Render(w, o.Id)
		w.WriteString(" {\n")
		for _, s := range o.Stmts {
			Render(w, s)
		}
		w.WriteString("}\n")

	default:
		panic(typeError(o))
	}
}
