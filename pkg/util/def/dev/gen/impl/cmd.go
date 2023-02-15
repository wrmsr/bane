//go:build !nodev

package impl

import (
	"go/format"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/def"
)

func Run(cwd string) string {
	pps := ParsePackages(cwd)
	pkg := check.Single(pps.Pkgs)

	var pkgDefs []def.PackageDef
	for _, fil := range pkg.Syntax {
		//_ = ast.Fprint(os.Stdout, pkg.Fset, fil, nil)
		//_ = printer.Fprint(os.Stdout, pkg.Fset, fil)

		astDefs := FindPkgDefCalls(fil, pkg.TypesInfo)
		for _, d := range astDefs {
			//_ = ast.Fprint(os.Stdout, pkg.Fset, d, nil)
			//_ = printer.Fprint(os.Stdout, pkg.Fset, d)

			rd := ReifyDef(d, pkg.TypesInfo)
			pkgDefs = append(pkgDefs, rd.(def.PackageDef))
		}

		astFuncs := FindPkgDefFuncs(fil, pkg.TypesInfo)
		for _, f := range astFuncs {
			rd := ReifyFunc(f, pkg.TypesInfo)
			pkgDefs = append(pkgDefs, rd.(def.PackageDef))
		}
	}

	ps := def.NewPackageSpec(pkg.ID, pkgDefs)
	s := NewFileGen(pps.Mod, pkg, ps).Gen()

	return string(check.Must1(format.Source([]byte(s))))
}
