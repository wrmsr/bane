package templates

import (
	"fmt"
	"strings"
	"testing"
	"text/template"

	"github.com/wrmsr/bane/pkg/util/check"
)

func TestInclude(t *testing.T) {
	tmpl := WithStd(template.New("?"))

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
