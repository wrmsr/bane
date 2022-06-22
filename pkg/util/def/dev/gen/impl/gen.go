//go:build !nodev

package impl

import (
	"embed"
	"strings"
	"text/template"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/def"
	"github.com/wrmsr/bane/pkg/util/maps"
	"github.com/wrmsr/bane/pkg/util/slices"
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
	tmpl := template.New("?")
	tmplu.WithInclude(tmpl)
	check.Must(tmpl.ParseFS(genTplEmbed, "*.tpl"))
	return tmpl.Lookup("gen.tpl")
}()

//

type fileGenFieldData struct {
	Struct *fileGenStructData
	Spec   *def.FieldSpec
}

type fileGenStructData struct {
	Pkg  *fileGenPkgData
	Spec *def.StructSpec

	Fields []*fileGenFieldData
}

type fileGenPkgData struct {
	Ns   string
	Spec *def.PackageSpec

	Structs []*fileGenStructData
}

func (fg *FileGen) Gen() string {
	pd := &fileGenPkgData{
		Ns:   "bane",
		Spec: fg.pkg,
	}

	pd.Structs = slices.Map(func(ss *def.StructSpec) *fileGenStructData {
		sd := &fileGenStructData{
			Pkg:  pd,
			Spec: ss,
		}

		sd.Fields = slices.Map(func(fs *def.FieldSpec) *fileGenFieldData {
			return &fileGenFieldData{
				Struct: sd,
				Spec:   fs,
			}
		}, maps.Values(ss.Fields()))

		return sd
	}, maps.Values(fg.pkg.Structs()))

	var sb strings.Builder
	check.NoErr(genTpl.Execute(&sb, &pd))
	return strings.ReplaceAll(sb.String(), "    ", "\t")
}
