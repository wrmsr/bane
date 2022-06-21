//go:build !nodev

package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"golang.org/x/mod/modfile"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/types/typeutil"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/def"
	rtu "github.com/wrmsr/bane/pkg/util/runtime"
)

var defsPkg = func() string {
	return reflect.TypeOf(def.PackageSpec{}).PkgPath()
}()

func findDirWithFile(cd, fn string) (string, error) {
	for {
		if _, err := os.Stat(filepath.Join(cd, fn)); err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return "", err
			}
		} else {
			return cd, nil
		}

		nd := filepath.Dir(cd)
		if nd == cd {
			return "", fmt.Errorf("file %s not found from dir %s", fn, cd)
		}
		cd = nd
	}
}

var dontFixRetract modfile.VersionFixer = func(_, vers string) (string, error) {
	return vers, nil
}

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

func findPkgDefCalls(fil *ast.File, ti *types.Info) []*ast.CallExpr {
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

func reifyIdentStr(node ast.Node) string {
	if node, ok := node.(*ast.BasicLit); ok {
		if node.Kind == token.STRING {
			s := node.Value
			if s[0:1] != "\"" || s[len(s)-1:] != "\"" {
				panic(fmt.Errorf("illegal ident str: %s", s))
			}
			s = s[1 : len(s)-1]
			if strings.Index(s, "/") >= 0 {
				panic(fmt.Errorf("illegal ident str: %s", s))
			}
			return s
		}
	}
	panic(fmt.Errorf("illegal ident str: %v", node))
}

func reifyDef(node ast.Node, ti *types.Info) any {
	call := node.(*ast.CallExpr)
	fn, _ := typeutil.Callee(ti, call).(*types.Func)
	check.NotNil(fn)

	pn := rtu.ParseName(fn.FullName())
	switch pn.Obj {

	case "Struct":
		opts := make([]def.StructOpt, len(call.Args)-1)
		for i, arg := range call.Args[1:] {
			opts[i] = reifyDef(arg, ti).(def.StructOpt)
		}
		return def.StructDef{
			Name: reifyIdentStr(call.Args[0]),
			Opts: opts,
		}

	case "Field":
		opts := make([]def.FieldOpt, len(call.Args)-1)
		for i, arg := range call.Args[1:] {
			opts[i] = reifyDef(arg, ti).(def.FieldOpt)
		}
		return def.FieldDef{
			Name: reifyIdentStr(call.Args[0]),
			Opts: opts,
		}

	case "Type":
		idx := call.Fun.(*ast.IndexExpr)
		panic(idx)

	case "Default":
		panic(call)

	case "Init":
		panic(call)

	}

	panic(fmt.Errorf("unhandled type: %T", node))
}

func main() {
	cd := check.Must(os.Getwd())
	rd := check.Must(findDirWithFile(cd, "go.mod"))
	if !strings.HasPrefix(cd, rd) {
		panic(fmt.Errorf("can't find path %s from root %s", cd, rd))
	}
	do := cd[len(rd)+1:]

	mf := filepath.Join(rd, "go.mod")
	mc := check.Must(ioutil.ReadFile(mf))
	mo := check.Must(modfile.Parse(mf, mc, dontFixRetract))
	mp := mo.Module.Mod.Path

	cfg := &packages.Config{
		Mode: packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
		ParseFile: func(fset *token.FileSet, filename string, src []byte) (*ast.File, error) {
			const mode = parser.AllErrors | parser.ParseComments
			if strings.HasSuffix(filepath.Base(filename), "_gen.go") {
				src = []byte{}
			}
			return parser.ParseFile(fset, filename, src, mode)
		},
	}

	pn := fmt.Sprintf("%s/%s", mp, do)
	pkgs := check.Must(packages.Load(cfg, pn))

	for _, pkg := range pkgs {
		for _, fil := range pkg.Syntax {
			//_ = ast.Fprint(os.Stdout, pkg.Fset, fil, nil)
			//_ = printer.Fprint(os.Stdout, pkg.Fset, fil)

			defs := findPkgDefCalls(fil, pkg.TypesInfo)
			for _, d := range defs {
				//_ = ast.Fprint(os.Stdout, pkg.Fset, d, nil)
				//_ = printer.Fprint(os.Stdout, pkg.Fset, d)

				rd := reifyDef(d, pkg.TypesInfo)
				fmt.Printf("%+v\n", rd)
			}
		}
	}
}
