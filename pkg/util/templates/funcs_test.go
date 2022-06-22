package templates

import (
	"fmt"
	"strings"
	"testing"
	"text/template"

	"github.com/wrmsr/bane/pkg/util/check"
)

func TestInclude(t *testing.T) {
	tmpl := template.New("?")

	WithInclude(tmpl)

	tmpl.Funcs(template.FuncMap{
		"indent": Indent,
	})

	check.Must(tmpl.Parse(`
{{define "foo"}}
hi im foo
{{end}}

{{include "foo" . | indent 1 "	"}}
`))

	var sb strings.Builder
	check.NoErr(tmpl.Execute(&sb, nil))
	fmt.Println(sb.String())
}
