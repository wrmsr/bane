//go:build !nodev

package impl

import (
	"fmt"
	"go/ast"
	"go/types"
	"reflect"

	"golang.org/x/tools/go/types/typeutil"

	"github.com/wrmsr/bane/x/annotate"

	rtu "github.com/wrmsr/bane/core/runtime"
)

var annotatePkg = func() string {
	return reflect.TypeOf((*annotate.Annotator)(nil)).Elem().PkgPath()
}()

func findAnnotateCalls(root ast.Node, ti *types.Info) []*ast.CallExpr {
	var ret []*ast.CallExpr
	ast.Inspect(root, func(node ast.Node) bool {
		switch node := node.(type) {
		case *ast.CallExpr:
			fn, _ := typeutil.Callee(ti, node).(*types.Func)
			if fn != nil {
				pn := rtu.ParseName(fn.FullName())
				if pn.Pkg == annotatePkg {
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
			ret = append(ret, findAnnotateCalls(decl, ti)...)
		default:
			bad := findAnnotateCalls(decl, ti)
			if len(bad) > 0 {
				panic(fmt.Errorf("illegal annotations: %v", bad))
			}
		}
	}
	return ret
}
