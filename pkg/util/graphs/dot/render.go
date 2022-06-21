package dot

import (
	"bufio"
	"html"
)

func write(w *bufio.Writer, s string) {
	_, err := w.WriteString(s)
	if err != nil {
		panic(err)
	}
}

func Render(w *bufio.Writer, o any) {
	switch o := o.(type) {
	case Raw:
		write(w, o.Raw)

	case Text:
		write(w, html.EscapeString(o.Text))

	case Cell:
		write(w, "<td>")
		Render(w, o.Value)
		write(w, "</td>")

	case Row:
		write(w, "<tr>")
		for _, c := range o.Cells {
			Render(w, c)
		}
		write(w, "</tr>")

	case Table:
		write(w, "<table>")
		for _, r := range o.Rows {
			Render(w, r)
		}
		write(w, "</table>")

	case Id:
		write(w, "\"")
		write(w, o.Id)
		write(w, "\"")

	case Attrs:
		if o.Attrs == nil {
			return
		}
		write(w, "[")
		i := 0
		for k, v := range o.Attrs {
			if i > 0 {
				write(w, ", ")
			}
			write(w, k)
			write(w, "=<")
			Render(w, v)
			write(w, ">")
			i++
		}
		write(w, "]")

	case RawStmt:
		write(w, o.Raw)
		write(w, "\n")

	case Edge:
		Render(w, o.Left)
		write(w, " -> ")
		Render(w, o.Right)
		if len(o.Attrs.Attrs) > 0 {
			write(w, " ")
			Render(w, o.Attrs)
		}
		write(w, ";\n")

	case Node:
		Render(w, o.Id)
		if len(o.Attrs.Attrs) > 0 {
			write(w, " ")
			Render(w, o.Attrs)
		}
		write(w, ";\n")

	case Graph:
		write(w, "digraph ")
		Render(w, o.Id)
		write(w, " {\n")
		for _, s := range o.Stmts {
			Render(w, s)
		}
		write(w, "}\n")

	default:
		panic(typeError(o))
	}
}
