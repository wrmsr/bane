//go:build !nodev

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/def"
	"github.com/wrmsr/bane/pkg/util/def/dev/gen/impl"
)

func main() {
	fs := flag.NewFlagSet("gen", flag.ExitOnError)

	var noWrite bool
	fs.BoolVar(&noWrite, "p", false, "do not write file, just print")

	check.Must(fs.Parse(os.Args[1:]))
	if fs.NArg() > 0 {
		panic(errors.New("unexpected args"))
	}

	cwd := check.Must1(os.Getwd())
	pps := impl.ParsePackages(cwd)
	pkg := check.Single(pps.Pkgs)

	pkgDefs := make([]def.PackageDef, 0)
	for _, fil := range pkg.Syntax {
		//_ = ast.Fprint(os.Stdout, pkg.Fset, fil, nil)
		//_ = printer.Fprint(os.Stdout, pkg.Fset, fil)

		astDefs := impl.FindPkgDefCalls(fil, pkg.TypesInfo)
		for _, d := range astDefs {
			//_ = ast.Fprint(os.Stdout, pkg.Fset, d, nil)
			//_ = printer.Fprint(os.Stdout, pkg.Fset, d)

			rd := impl.ReifyDef(d, pkg.TypesInfo)
			pkgDefs = append(pkgDefs, rd.(def.PackageDef))
		}
	}

	ps := def.NewPackageSpec(pkg.ID, pkgDefs)
	s := impl.NewFileGen(pps.Mod, pkg, ps).Gen()

	if noWrite {
		fmt.Println(s)

	} else {
		fp := filepath.Join(cwd, "def_gen.go")
		check.Must(ioutil.WriteFile(fp, []byte(s), 0644))

		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
			defer cancel()

			cmd := exec.CommandContext(ctx, "go", "run", "cmd/gofmt", "-s", "-w", fp)
			check.Must(cmd.Run())
		}()
	}
}
