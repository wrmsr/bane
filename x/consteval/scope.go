package consteval

import (
	"go/ast"

	bt "github.com/wrmsr/bane/core/types"
)

type Lookup interface {
	Lookup(n string) any // Value | ast.Node | error
}

type Scope struct {
	p *Scope
	a *ast.Scope
	m map[string]bt.Result[Value]
}

func findValueSpecExpr(vs *ast.ValueSpec, n string) ast.Expr {
	for i, vn := range vs.Names {
		if vn.Name == n {
			return vs.Values[i]
		}
	}
	return nil
}

func ScopeFromAst(a *ast.Scope) *Scope {
	return &Scope{
		a: a,
		m: make(map[string]bt.Result[Value]),
	}
}

func (sc *Scope) lookup(n string) any {
	if sc.m != nil {
		if v, ok := sc.m[n]; ok {
			return v
		}
	}

	if sc.a != nil {
		if obj := sc.a.Lookup(n); obj != nil {
			switch obj.Kind {

			case ast.Con, ast.Var:
				return findValueSpecExpr(obj.Decl.(*ast.ValueSpec), obj.Name)

			case ast.Fun:
				return obj.Decl.(*ast.FuncDecl)

			}
			panic(obj)
		}
	}

	if sc.p != nil {
		return sc.p.lookup(n)
	}

	return NameError{N: n}
}

func (sc *Scope) assign(n string, v bt.Result[Value]) bool {
	if sc.m != nil {
		if _, ok := sc.m[n]; ok {
			sc.m[n] = v
			return true
		}
	}

	if sc.p != nil {
		if sc.p.assign(n, v) {
			return true
		}
	}

	if sc.m != nil {
		sc.m[n] = v
	}

	return false
}

func (sc *Scope) makeChild() *Scope {
	return &Scope{
		p: sc,
		m: make(map[string]bt.Result[Value]),
	}
}
