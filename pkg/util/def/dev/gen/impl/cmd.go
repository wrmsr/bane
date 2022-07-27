package impl

import (
	"context"
	"time"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/def"
)

func Run(cwd string) string {
	pps := ParsePackages(cwd)
	pkg := check.Single(pps.Pkgs)

	pkgDefs := make([]def.PackageDef, 0)
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
	}

	ps := def.NewPackageSpec(pkg.ID, pkgDefs)
	s := NewFileGen(pps.Mod, pkg, ps).Gen()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	return FormatCode(ctx, s)
}
