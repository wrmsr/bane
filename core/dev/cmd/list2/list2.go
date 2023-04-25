//
/*
TODO:
  - subcmds
   - list
    - formats:
     - json output
     - dirs    (. ./sql ./x)
     - submods (sql x)
     - pkgs    (core/foo core/foo/bar sql/foo)
     - srcs    (core/foo/foo.go core/foo/bar/bar.go sql/foo/sql.go)
  - foreach
   - find-style '%'
   - parallel
 - flags
  - 'all' - include submodules
  - *ignore* .goignore - honor by default in all output modes
  - *honor* .gitignore by default
 - !!! DEP GRAPHS - go mod update IN DEP ORDER
 - gobin


https://github.com/golang/go/issues/50745
*/
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/wrmsr/bane/core/check"
	ctr "github.com/wrmsr/bane/core/container"
	"github.com/wrmsr/bane/core/dev/tool"
	fnu "github.com/wrmsr/bane/core/funcs"
	ju "github.com/wrmsr/bane/core/json"
	osu "github.com/wrmsr/bane/core/os"
	rtu "github.com/wrmsr/bane/core/runtime"
	"github.com/wrmsr/bane/core/workers"
)

//

type Prj struct {
	mods []*Mod

	modsByDir map[string]*Mod
}

type Mod struct {
	dir  string
	pkgs []*tool.ListPackage
}

//

func execIn(ctx context.Context, wd string, exe string, args ...string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, exe, args...)
	cmd.Dir = wd

	var outb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return nil, err
	}
	return outb.Bytes(), nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	cwd := check.Must1(os.Getwd())

	check.Ok(check.Must1(osu.Exists(filepath.Join(cwd, "go.mod"))))

	m := ctr.NewLockedMutMap[string, []*tool.ListPackage](ctr.NewMutStdMap[string, []*tool.ListPackage](nil))

	var ds []string
	check.Must(filepath.Walk(cwd, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() && info.Name() == "go.mod" {
			ds = append(ds, info.Name())
		}
		return nil
	}))

	var fns []func()
	for _, d := range []string{
		".",
		"sql",
		"x",
	} {

		d := d

		fns = append(fns, func() {
			o := check.Must1(execIn(ctx, filepath.Join(cwd, d), "go", "list", "-json", "./..."))

			var pkgs []*tool.ListPackage
			jd := json.NewDecoder(bytes.NewReader(o))
			for jd.More() {
				var pkg *tool.ListPackage
				check.Must(jd.Decode(&pkg))
				pkgs = append(pkgs, pkg)
			}

			m.Put(d, pkgs)
		})
	}

	wfns := make([]func(context.Context) (any, error), len(fns))
	for i, fn := range fns {
		fn := fn
		wfns[i] = func(context.Context) (any, error) {
			return nil, fnu.Recover(fn)
		}
	}
	_ = check.Must1(workers.DoParallelErrGroup(context.Background(), rtu.MaxProcs(), wfns))

	fmt.Println(check.Must1(ju.MarshalPretty(m)))
}
