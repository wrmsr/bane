//go:build !nodev

package main

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
	"golang.org/x/tools/go/packages"

	"github.com/wrmsr/bane/core/check"
	ctr "github.com/wrmsr/bane/core/container"
	fnu "github.com/wrmsr/bane/core/funcs"
	ju "github.com/wrmsr/bane/core/json"
	"github.com/wrmsr/bane/core/maps"
	osu "github.com/wrmsr/bane/core/os"
	rtu "github.com/wrmsr/bane/core/runtime"
	"github.com/wrmsr/bane/core/slices"
	bt "github.com/wrmsr/bane/core/types"
	"github.com/wrmsr/bane/core/workers"
)

var dontFixRetract modfile.VersionFixer = func(_, vers string) (string, error) {
	return vers, nil
}

type Output struct {
	Builtin  map[string][]string `json:"builtin,omitempty"`
	External map[string][]string `json:"external,omitempty"`

	Internal ctr.Map[string, rtu.SortedNames] `json:"internal"`
}

func main() {
	cd := check.Must1(os.Getwd())
	rd := check.Must1(osu.FindParentDirWithFile(cd, "go.mod"))
	if !strings.HasPrefix(cd, rd) {
		panic(fmt.Errorf("can't find path %s from root %s", cd, rd))
	}

	mf := filepath.Join(rd, "go.mod")
	mc := check.Must1(os.ReadFile(mf))
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

	i := ctr.MapValues[string, []string, rtu.SortedNames](fnu.Bind1x1x1(rtu.SortNames, mp), m)

	bsm := make(map[string]maps.Set[string])
	esm := make(map[string]maps.Set[string])
	i.ForEach(func(kv bt.Kv[string, rtu.SortedNames]) bool {
		for _, b := range kv.V.Builtin {
			maps.ComputeDefault(bsm, b, maps.MakeSet[string]).Add(kv.K)
		}
		for _, e := range kv.V.External {
			maps.ComputeDefault(esm, e, maps.MakeSet[string]).Add(kv.K)
		}
		return true
	})

	gsm := func(s maps.Set[string]) []string {
		return slices.Sort(s.Slice())
	}
	o := Output{
		Builtin:  maps.MapValues(gsm, bsm),
		External: maps.MapValues(gsm, esm),
		Internal: i,
	}

	fmt.Println(check.Must1(ju.MarshalIndentString(o, "", "  ")))
}
