package rendering

import (
	"io"
	"strings"
)

type Renderer interface {
	Render(w io.Writer)
}

func RenderString(r Renderer) string {
	var bs strings.Builder
	r.Render(&bs)
	return bs.String()
}
