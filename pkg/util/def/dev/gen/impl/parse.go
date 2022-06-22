//go:build !nodev

package impl

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
	"golang.org/x/tools/go/packages"

	"github.com/wrmsr/bane/pkg/util/check"
)

var dontFixRetract modfile.VersionFixer = func(_, vers string) (string, error) {
	return vers, nil
}

func findDirWithFile(cd, fn string) (string, error) {
	for {
		if _, err := os.Stat(filepath.Join(cd, fn)); err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return "", err
			}
		} else {
			return cd, nil
		}

		nd := filepath.Dir(cd)
		if nd == cd {
			return "", fmt.Errorf("file %s not found from dir %s", fn, cd)
		}
		cd = nd
	}
}

func ParsePackages(cd string) []*packages.Package {
	rd := check.Must(findDirWithFile(cd, "go.mod"))
	if !strings.HasPrefix(cd, rd) {
		panic(fmt.Errorf("can't find path %s from root %s", cd, rd))
	}
	do := cd[len(rd)+1:]

	mf := filepath.Join(rd, "go.mod")
	mc := check.Must(ioutil.ReadFile(mf))
	mo := check.Must(modfile.Parse(mf, mc, dontFixRetract))
	mp := mo.Module.Mod.Path

	cfg := &packages.Config{
		Mode: packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
		ParseFile: func(fset *token.FileSet, filename string, src []byte) (*ast.File, error) {
			const mode = parser.AllErrors | parser.ParseComments
			for _, suf := range []string{
				"_gen.go", "_xgen.go",
			} {
				if strings.HasSuffix(filepath.Base(filename), suf) {
					src = []byte{}
				}
			}
			return parser.ParseFile(fset, filename, src, mode)
		},
	}

	pn := fmt.Sprintf("%s/%s", mp, do)
	return check.Must(packages.Load(cfg, pn))
}
