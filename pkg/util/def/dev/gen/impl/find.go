//go:build !nodev

package impl

import (
	"fmt"
	"go/ast"
	"go/types"
	"reflect"
	"strings"

	"golang.org/x/tools/go/types/typeutil"

	"github.com/wrmsr/bane/pkg/util/def"
	rtu "github.com/wrmsr/bane/pkg/util/runtime"
)

var defsPkg = func() string {
	return reflect.TypeOf(def.PackageSpec{}).PkgPath()
}()

func findDefCalls(root ast.Node, ti *types.Info) []*ast.CallExpr {
	var ret []*ast.CallExpr
	ast.Inspect(root, func(node ast.Node) bool {
		switch node := node.(type) {
		case *ast.CallExpr:
			fn, _ := typeutil.Callee(ti, node).(*types.Func)
			if fn != nil {
				pn := rtu.ParseName(fn.FullName())
				if pn.Pkg == defsPkg {
					ret = append(ret, node)
					return false
				}
			}
		}
		return true
	})
	return ret
}

func FindPkgDefCalls(fil *ast.File, ti *types.Info) []*ast.CallExpr {
	var ret []*ast.CallExpr
	for _, decl := range fil.Decls {
		switch decl := decl.(type) {
		case *ast.GenDecl:
			ret = append(ret, findDefCalls(decl, ti)...)
		default:
			bad := findDefCalls(decl, ti)
			if len(bad) > 0 {
				panic(fmt.Errorf("illegal defs: %v", bad))
			}
		}
	}
	return ret
}

func FindPkgDefFuncs(fil *ast.File, ti *types.Info) []*ast.FuncDecl {
	var ret []*ast.FuncDecl
	for _, decl := range fil.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			if strings.HasPrefix(decl.Name.Name, "_def_") {
				ret = append(ret, decl)
			}
		}
	}
	return ret
}
