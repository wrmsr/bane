//go:build !nodev

package main

import (
	"context"
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
	fnu "github.com/wrmsr/bane/pkg/util/funcs"
	ju "github.com/wrmsr/bane/pkg/util/json"
	"github.com/wrmsr/bane/pkg/util/maps"
	osu "github.com/wrmsr/bane/pkg/util/os"
	rtu "github.com/wrmsr/bane/pkg/util/runtime"
	"github.com/wrmsr/bane/pkg/util/slices"
	"github.com/wrmsr/bane/pkg/util/workers"
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

	var dirs []string
	for _, arg := range os.Args[1:] {
		check.Must(filepath.Walk(arg, func(path string, info fs.FileInfo, err error) error {
			if info.IsDir() {
				dirs = append(dirs, path)
			}
			return nil
		}))
	}

	fns := make([]func(), len(dirs))
	m := ctr.NewLockedMutMap[string, []string](ctr.NewMutStdMap[string, []string](nil))
	for i, dir := range dirs {
		dir := dir
		fns[i] = func() {
			cfg := &packages.Config{
				Mode: packages.NeedImports | packages.NeedDeps,
			}

			pn := fmt.Sprintf("%s/%s", mp, dir)
			pkgs := check.Must1(packages.Load(cfg, pn))

			for _, pkg := range pkgs {
				s := maps.Keys(pkg.Imports)
				slices.Sort(s)
				m.Put(pkg.ID, s)
			}
		}
	}

	wfns := make([]func(context.Context) (any, error), len(fns))
	for i, fn := range fns {
		fn := fn
		wfns[i] = func(context.Context) (any, error) {
			return nil, fnu.Recover(fn)
		}
	}
	_ = check.Must1(workers.DoParallelErrGroup(context.Background(), rtu.MaxProcs(), wfns))

	o := ctr.MapValues[string, []string, rtu.SortedNames](fnu.Bind1x1x1(rtu.SortNames, mp), m)
	fmt.Println(check.Must1(ju.MarshalIndentString(o, "", "  ")))
}
