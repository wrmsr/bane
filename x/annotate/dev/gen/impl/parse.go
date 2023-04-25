//go:build !nodev

package impl

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
	"golang.org/x/tools/go/packages"

	"github.com/wrmsr/bane/core/check"
	osu "github.com/wrmsr/bane/core/os"
)

var dontFixRetract modfile.VersionFixer = func(_, vers string) (string, error) {
	return vers, nil
}

type ParsedPackages struct {
	Mod  *modfile.File
	Pkgs []*packages.Package
}

func ParsePackages(cd string) ParsedPackages {
	rd := check.Must1(osu.FindParentDirWithFile(cd, "go.mod"))
	if !strings.HasPrefix(cd, rd) {
		panic(fmt.Errorf("can't find path %s from root %s", cd, rd))
	}
	do := cd[len(rd)+1:]

	mf := filepath.Join(rd, "go.mod")
	mc := check.Must1(os.ReadFile(mf))
	mo := check.Must1(modfile.Parse(mf, mc, dontFixRetract))
	mp := mo.Module.Mod.Path

	cfg := &packages.Config{
		Mode: packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
		ParseFile: func(fset *token.FileSet, filename string, src []byte) (*ast.File, error) {
			const mode = parser.AllErrors | parser.ParseComments
			for _, suf := range []string{
				"_gen.go",
			} {
				if strings.HasSuffix(filepath.Base(filename), suf) {
					src = []byte{}
				}
			}
			return parser.ParseFile(fset, filename, src, mode)
		},
	}

	pn := fmt.Sprintf("%s/%s", mp, do)
	return ParsedPackages{
		Mod:  mo,
		Pkgs: check.Must1(packages.Load(cfg, pn)),
	}
}
