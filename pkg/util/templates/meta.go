package templates

import (
	"bytes"
	"text/template"

	tmplfn "github.com/wrmsr/bane/pkg/util/templates/funcs"
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

func WithStd(tmpl *template.Template) *template.Template {
	tmpl = tmpl.Funcs(template.FuncMap(tmplfn.Std))
	tmpl = WithInclude(tmpl)
	return tmpl
}
