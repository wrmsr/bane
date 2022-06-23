package templates

import (
	"bytes"
	"strings"
	"text/template"
)

func WithInclude(tmpl *template.Template) *template.Template {
	return tmpl.Funcs(template.FuncMap{
		"include": func(name string, data interface{}) (string, error) {
			buf := bytes.NewBuffer(nil)
			if err := tmpl.ExecuteTemplate(buf, name, data); err != nil {
				return "", err
			}
			return buf.String(), nil
		},
	})
}

func Indent(n int, s, v string) string {
	pad := strings.Repeat(s, n)
	return pad + strings.Replace(v, "\n", "\n"+pad, -1)
}

func IntRange(start, end int) []int {
	n := end - start
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = start + i
	}
	return r
}
