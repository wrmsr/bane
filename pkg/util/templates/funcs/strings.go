package funcs

import (
	"strings"
	"text/template"
)

var _ = Std.AddAll(template.FuncMap{
	"indent": Indent,
})

func Indent(n int, s, v string) string {
	pad := strings.Repeat(s, n)
	return pad + strings.Replace(v, "\n", "\n"+pad, -1)
}
