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
	pkgs := impl.ParsePackages(check.Must1(os.Getwd()))

	for _, pkg := range pkgs {
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
		s := impl.NewFileGen(ps).Gen()
		fmt.Println(s)
	}
}
