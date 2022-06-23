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
	ps *def.PackageSpec
}

func NewFileGen(pkg *def.PackageSpec) *FileGen {
	return &FileGen{
		ps: pkg,
	}
}

//

type fileGenFieldData struct {
	Struct *fileGenStructData
	Spec   *def.FieldSpec
}

func (fg *FileGen) makeFieldData(sd *fileGenStructData, fs *def.FieldSpec) *fileGenFieldData {
	return &fileGenFieldData{
		Struct: sd,
		Spec:   fs,
	}
}

//

type fileGenStructData struct {
	Pkg  *fileGenPkgData
	Spec *def.StructSpec

	Fields []*fileGenFieldData
}

func (fg *FileGen) makeStructData(pd *fileGenPkgData, ss *def.StructSpec) *fileGenStructData {
	sd := &fileGenStructData{
		Pkg:  pd,
		Spec: ss,
	}

	sd.Fields = slices.Map(func(fs *def.FieldSpec) *fileGenFieldData {
		return fg.makeFieldData(sd, fs)
	}, maps.Values(ss.Fields()))

	return sd
}

//

type fileGenPkgData struct {
	Ns   string
	Spec *def.PackageSpec

	Structs []*fileGenStructData
}

func (fg *FileGen) makePkgData() *fileGenPkgData {
	pd := &fileGenPkgData{
		Ns:   "bane",
		Spec: fg.ps,
	}

	pd.Structs = slices.Map(func(ss *def.StructSpec) *fileGenStructData {
		return fg.makeStructData(pd, ss)
	}, maps.Values(fg.ps.Structs()))

	return pd
}

//

//go:embed gen.tpl
var genTmplEmbed embed.FS

var genTmpl = func() *template.Template {
	tmpl := template.New("?")

	tmplu.WithInclude(tmpl)

	tmpl.Funcs(template.FuncMap{
		"indent":   tmplu.Indent,
		"intRange": tmplu.IntRange,
	})

	check.Must(tmpl.ParseFS(genTmplEmbed, "*.tpl"))
	return tmpl.Lookup("gen.tpl")
}()

//

func (fg *FileGen) Gen() string {
	pd := fg.makePkgData()

	tmpl := check.Must(genTmpl.Parse(`{{include "package" .}}`))

	var sb strings.Builder
	check.NoErr(tmpl.Execute(&sb, &pd))
	s := sb.String()

	//s = strings.ReplaceAll(s, "    ", "\t")
	return s
}
