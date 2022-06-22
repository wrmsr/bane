//go:build !nodev

package impl

import (
	"embed"
	"strings"
	"text/template"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/def"
	tmplu "github.com/wrmsr/bane/pkg/util/templates"
)

type FileGen struct {
	pkg *def.PackageSpec
}

func NewFileGen(pkg *def.PackageSpec) *FileGen {
	return &FileGen{
		pkg: pkg,
	}
}

//go:embed gen.tpl
var genTplEmbed embed.FS

var genTpl = func() *template.Template {
	tmpl := check.Must(template.ParseFS(genTplEmbed, "*.tpl"))
	return tmplu.WithInclude(tmpl)
}()

type fileGenData struct {
	Ns  string
	Pkg *def.PackageSpec
}

func (fg *FileGen) Gen() string {
	data := fileGenData{
		Ns:  "bane",
		Pkg: fg.pkg,
	}

	var sb strings.Builder
	check.NoErr(genTpl.Execute(&sb, &data))
	return strings.ReplaceAll(sb.String(), "    ", "\t")
}
