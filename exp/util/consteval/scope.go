//
/*
https://github.com/golang/tools/tree/master/go/ssa/interp
https://github.com/llvm/llvm-project/tree/main/clang/lib/AST/Interp

TODO:
 - more lazy - mirror ast.Scope re inheritance
 - togglable memoization - consteval func scopes
 - func
 - addr - *including* var refs, like &foo[2]
 - func Const[T any](v T) { return v } ?
  - if consteval then _no_ as running funcs lol
 - intrinsics?
  - math
  - strings
  - reflect - fields, etc
  - io? encoding?
 - views interop - frozen slices
  - inlineable? lol fuckin go
*/
package consteval

import (
	"go/ast"
	"go/constant"

	"github.com/wrmsr/bane/pkg/util/check"
)

//

type Lookup interface {
	Lookup(n string) any // Value | ast.Node | error
}

//

type Scope struct {
	p *Scope
	a *ast.Scope
	m map[string]any // Value | error
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
		m: make(map[string]any),
	}
}

func (sc *Scope) lookup(n string) any {
	if sc.m != nil {
		if v := sc.m[n]; v != nil {
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

func (sc *Scope) Reduce(o any) (Value, error) {
	if v, ok := o.(Value); ok {
		return v, nil
	}

	if err, ok := o.(error); ok {
		return nil, err
	}

	if n, ok := o.(ast.Expr); ok {
		return sc.evalExpr(n)
	}

	panic(o)
}

func (sc *Scope) evalExpr(n ast.Expr) (Value, error) {
	switch n := n.(type) {

	case *ast.BasicLit:
		return Basic{
			K: basicKindFromAst(n.Kind),
			S: n.Value,
		}, nil

	case *ast.CompositeLit:
		switch tn := n.Type.(type) {

		case *ast.Ident:
			var m map[string]Value
			if len(n.Elts) > 0 {
				m = make(map[string]Value, len(n.Elts))
				for _, ne := range n.Elts {
					kv := ne.(*ast.KeyValueExpr)
					k := kv.Key.(*ast.Ident).Name
					v, err := sc.evalExpr(kv.Value)
					if err != nil {
						return nil, err
					}
					m[k] = v
				}
			}
			return Struct{
				T: TypeName(tn.Name),
				M: m,
			}, nil

		case *ast.ArrayType:
			t := tn.Elt.(*ast.Ident).Name
			var s []Value
			if len(n.Elts) > 0 {
				s = make([]Value, len(n.Elts))
				for i, ne := range n.Elts {
					sv, err := sc.evalExpr(ne)
					if err != nil {
						return nil, err
					}
					s[i] = sv
				}
			}
			return Array{
				T: TypeName(t),
				S: s,
			}, nil

		case *ast.MapType:
			kt := tn.Key.(*ast.Ident).Name
			vt := tn.Value.(*ast.Ident).Name
			var s []MapEntry
			if len(n.Elts) > 0 {
				s = make([]MapEntry, len(n.Elts))
				for i, ne := range n.Elts {
					kve := ne.(*ast.KeyValueExpr)
					k, err := sc.evalExpr(kve.Key)
					if err != nil {
						return nil, err
					}
					v, err := sc.evalExpr(kve.Value)
					if err != nil {
						return nil, err
					}
					s[i] = MapEntry{
						K: k,
						V: v,
					}
				}
			}
			return Map{
				KT: TypeName(kt),
				VT: TypeName(vt),
				S:  s,
			}, nil
		}

	case *ast.CallExpr:
		if ie, ok := n.Fun.(*ast.IndexExpr); ok {
			fn := ie.X.(*ast.Ident).Name
			switch fn {

			case "Type":
				if len(n.Args) > 0 {
					panic(n)
				}
				tn := ie.Index.(*ast.Ident).Name
				return Type{
					T: TypeName(tn),
				}, nil

			}
			break
		}

		if ii, ok := n.Fun.(*ast.Ident); ok {
			fd := sc.lookup(ii.Name)
			if fe, ok := fd.(error); ok {
				return nil, fe
			}
			s2 := &Scope{
				p: sc,
			}
			return s2.execStmts(fd.(*ast.FuncDecl).Body.List)
		}

	case *ast.Ident:
		iv := sc.lookup(n.Name)

		if le, ok := iv.(error); ok {
			return nil, le
		}

		if v2, ok := iv.(Value); ok {
			return v2, nil
		}

		if ie, ok := iv.(ast.Expr); ok {
			rv, err := sc.evalExpr(ie)
			if err != nil {
				if sc.m != nil {
					sc.m[n.Name] = err
				}
				return nil, err
			}
			v2 := rv.(Value)
			if sc.m != nil {
				sc.m[n.Name] = v2
			}
			return v2, nil
		}

	//

	case *ast.UnaryExpr:
		x, err := sc.evalExpr(n.X)
		if err != nil {
			return nil, err
		}
		xb := x.(Basic)
		xc := constant.MakeFromLiteral(xb.S, xb.K.Ast(), 0)
		zc := constant.UnaryOp(n.Op, xc, 0) // FIXME: prec?
		zv := Basic{
			K: xb.K,
			S: zc.ExactString(),
		}
		return zv, nil

	case *ast.BinaryExpr:
		x, err := sc.evalExpr(n.X)
		if err != nil {
			return nil, err
		}
		y, err := sc.evalExpr(n.Y)
		if err != nil {
			return nil, err
		}
		xb := x.(Basic)
		yb := y.(Basic)
		if xb.K != yb.K {
			return nil, TypeError{S: []string{xb.K.String(), yb.K.String()}}
		}
		xc := constant.MakeFromLiteral(xb.S, xb.K.Ast(), 0)
		yc := constant.MakeFromLiteral(yb.S, yb.K.Ast(), 0)
		zc := constant.BinaryOp(xc, n.Op, yc)
		zv := Basic{
			K: xb.K,
			S: zc.ExactString(),
		}
		return zv, nil

	}
	return Dynamic{}, nil
}

func (sc *Scope) execStmts(sts []ast.Stmt) (Value, error) {
	for _, st := range sts {
		switch st := st.(type) {

		case *ast.ReturnStmt:
			vn := check.Single(st.Results)
			return sc.evalExpr(vn)

		default:
			panic(st)
		}
	}
	panic(sc)
}
