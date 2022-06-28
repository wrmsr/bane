//go:build !nodev

package main

import (
	"fmt"
	"os"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/def"
	"github.com/wrmsr/bane/pkg/util/def/dev/gen/impl"
)

func main() {
	pkgs := impl.ParsePackages(check.Must(os.Getwd()))

	for _, pkg := range pkgs {
		for _, fil := range pkg.Syntax {
			//_ = ast.Fprint(os.Stdout, pkg.Fset, fil, nil)
			//_ = printer.Fprint(os.Stdout, pkg.Fset, fil)

			defs := impl.FindPkgDefCalls(fil, pkg.TypesInfo)
			rdefs := make([]def.PackageDef, len(defs))
			for i, d := range defs {
				//_ = ast.Fprint(os.Stdout, pkg.Fset, d, nil)
				//_ = printer.Fprint(os.Stdout, pkg.Fset, d)

				rd := impl.ReifyDef(d, pkg.TypesInfo)
				rdefs[i] = rd.(def.PackageDef)
			}

			pspec := def.NewPackageSpec("?", rdefs)

			s := impl.NewFileGen(pspec).Gen()
			fmt.Println(s)
		}
	}
}
