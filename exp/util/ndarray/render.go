package ndarray

import (
	"fmt"
	"io"
	"strings"
)

func (a NdArray[T]) Render(w io.Writer, fn func(io.Writer, T)) {
	indent := func(n int) {
		for ; n > 0; n-- {
			_, _ = w.Write([]byte(" "))
		}
	}

	var rec func(int, Dim)
	rec = func(n int, o Dim) {
		st := a.v.st.Get(n)

		indent(n)
		_, _ = w.Write([]byte("["))

		if n < a.v.sh.Len()-1 {
			_, _ = w.Write([]byte("\n"))

			for i := Dim(0); i < a.v.sh.Get(n); i++ {
				if i > 0 {
					_, _ = w.Write([]byte(",\n"))
				}

				rec(n+1, o+(i*st))
			}

			_, _ = w.Write([]byte("\n"))
			indent(n)

		} else {
			_, _ = w.Write([]byte(" "))

			for i := Dim(0); i < a.v.sh.Get(n); i++ {
				if i > 0 {
					_, _ = w.Write([]byte(", "))
				}

				e := a.d[a.v.o+o+(i*st)]
				if fn != nil {
					fn(w, e)
				} else {
					_, _ = fmt.Fprintf(w, "%v", e)
				}
			}

			_, _ = w.Write([]byte(" "))
		}

		_, _ = w.Write([]byte("]"))
	}

	rec(0, 0)
}

func (a NdArray[T]) String() string {
	var sb strings.Builder
	a.Render(&sb, func(w io.Writer, e T) {
		_, _ = fmt.Fprintf(w, "%6v", e)
	})
	return sb.String()
}
