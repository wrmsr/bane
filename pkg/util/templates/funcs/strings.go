package funcs

import (
	"strings"
	"text/template"

	stru "github.com/wrmsr/bane/pkg/util/strings"
)

var _ = Std.AddAll(template.FuncMap{
	"indent": Indent,

	"dedent": stru.Dedent,
})

func Indent(n int, s, v string) string {
	pad := strings.Repeat(s, n)
	return pad + strings.Replace(v, "\n", "\n"+pad, -1)
}
