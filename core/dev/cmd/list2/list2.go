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
 - !!! DEP GRAPHS - go mod update IN DEP ORDER
 - gobin


https://github.com/golang/go/issues/50745
*/
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/wrmsr/bane/core/check"
	ctr "github.com/wrmsr/bane/core/container"
	fnu "github.com/wrmsr/bane/core/funcs"
	osu "github.com/wrmsr/bane/core/os"
	rtu "github.com/wrmsr/bane/core/runtime"
	"github.com/wrmsr/bane/core/workers"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	cwd := check.Must1(os.Getwd())

	check.Ok(check.Must1(osu.Exists(filepath.Join(cwd, "go.mod"))))

	m := ctr.NewLockedMutMap[string, []*Package](ctr.NewMutStdMap[string, []*Package](nil))

	var fns []func()
	for _, d := range []string{
		".",
		"sql",
		"x",
	} {
		d := d

		fns = append(fns, func() {
			cmd := exec.CommandContext(
				ctx,
				"go", "list", "-json", "./...",
			)
			cmd.Dir = filepath.Join(cwd, d)

			var outb bytes.Buffer
			cmd.Stdout = &outb
			cmd.Stderr = os.Stderr
			check.Must(cmd.Run())

			var pkgs []*Package
			jd := json.NewDecoder(bytes.NewReader(outb.Bytes()))
			for jd.More() {
				var pkg *Package
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
}
