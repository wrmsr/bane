//go:build !nodev

package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
	"golang.org/x/tools/go/packages"

	"github.com/wrmsr/bane/pkg/util/check"
	ctr "github.com/wrmsr/bane/pkg/util/container"
	ju "github.com/wrmsr/bane/pkg/util/json"
	"github.com/wrmsr/bane/pkg/util/maps"
	osu "github.com/wrmsr/bane/pkg/util/os"
	"github.com/wrmsr/bane/pkg/util/slices"
)

var dontFixRetract modfile.VersionFixer = func(_, vers string) (string, error) {
	return vers, nil
}

func main() {
	cd := check.Must1(os.Getwd())
	rd := check.Must1(osu.FindParentDirWithFile(cd, "go.mod"))
	if !strings.HasPrefix(cd, rd) {
		panic(fmt.Errorf("can't find path %s from root %s", cd, rd))
	}

	mf := filepath.Join(rd, "go.mod")
	mc := check.Must1(ioutil.ReadFile(mf))
	mo := check.Must1(modfile.Parse(mf, mc, dontFixRetract))
	mp := mo.Module.Mod.Path

	m := ctr.NewLockedMutMap[string, []string](ctr.NewMutMap[string, []string](nil))

	for _, arg := range os.Args[1:] {
		check.Must(filepath.Walk(arg, func(path string, info fs.FileInfo, err error) error {
			if !info.IsDir() {
				return nil
			}

			cfg := &packages.Config{
				Mode: packages.NeedImports | packages.NeedDeps,
			}

			pn := fmt.Sprintf("%s/%s", mp, path)
			pkgs := check.Must1(packages.Load(cfg, pn))

			for _, pkg := range pkgs {
				s := maps.Keys(pkg.Imports)
				slices.Sort(s)
				m.Put(pkg.ID, s)
			}

			return nil
		}))
	}

	fmt.Println(check.Must1(ju.MarshalIndentString(m, "", "  ")))
}
