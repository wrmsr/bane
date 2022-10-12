package impl

import (
	"go/ast"

	"golang.org/x/tools/go/packages"

	"github.com/wrmsr/bane/pkg/util/check"
)

type FuncDecl struct {
	Name string
	Decl *ast.FuncDecl
	File *ast.File
}

type Resolver struct {
	pkg *packages.Package

	funcDecls map[string]*FuncDecl
}

func nameType(e ast.Expr) string {
	switch e := e.(type) {
	case *ast.Ident:
		return e.Name
	case *ast.StarExpr:
		return nameType(e.X)
	case *ast.IndexExpr:
		return nameType(e.X)
	default:
		panic(e)
	}
}

func nameFuncDecl(fd *ast.FuncDecl) string {
	n := fd.Name.Name
	if fd.Recv != nil && len(fd.Recv.List) > 0 {
		rt := check.Single(fd.Recv.List).Type
		rn := nameType(rt)
		n = rn + "." + n
	}
	return n
}

func nameFuncRef(e any) string {
	switch e := e.(type) {
	case string:
		return e
	case *ast.Ident:
		return e.Name
	default:
		panic(e)
	}
}

func newResolver(pkg *packages.Package) *Resolver {
	r := &Resolver{
		funcDecls: make(map[string]*FuncDecl),
	}

	for _, file := range pkg.Syntax {
		for _, decl := range file.Decls {
			switch decl := decl.(type) {
			case *ast.FuncDecl:
				n := nameFuncDecl(decl)
				if _, ok := r.funcDecls[n]; ok {
					panic(n)
				}
				r.funcDecls[n] = &FuncDecl{
					Name: n,
					Decl: decl,
					File: file,
				}
			}
		}
	}

	return r
}

func (r *Resolver) FuncDecl(n string) *FuncDecl {
	return r.funcDecls[n]
}
