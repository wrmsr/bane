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

func (sc *Scope) assign(n string, v any) bool {
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

func (sc *Scope) Eval(o any) any {
	if v, ok := o.(Value); ok {
		return v
	}

	if err, ok := o.(error); ok {
		return err
	}

	if n, ok := o.(ast.Expr); ok {
		return sc.evalExpr(n)
	}

	panic(o)
}

func (sc *Scope) makeChild() *Scope {
	return &Scope{
		p: sc,
		m: make(map[string]any),
	}
}

func (sc *Scope) evalExpr(n ast.Expr) any {
	switch n := n.(type) {

	case *ast.BasicLit:
		return Basic{
			K: basicKindFromAst(n.Kind),
			S: n.Value,
		}

	case *ast.CompositeLit:
		switch tn := n.Type.(type) {

		case *ast.Ident:
			var m map[string]Value
			if len(n.Elts) > 0 {
				m = make(map[string]Value, len(n.Elts))
				for _, ne := range n.Elts {
					kv := ne.(*ast.KeyValueExpr)
					k := kv.Key.(*ast.Ident).Name
					v := sc.evalExpr(kv.Value)
					if _, ok := v.(error); ok {
						return v
					}
					m[k] = v.(Value)
				}
			}
			return Struct{
				T: TypeName(tn.Name),
				M: m,
			}

		case *ast.ArrayType:
			t := tn.Elt.(*ast.Ident).Name
			var s []Value
			if len(n.Elts) > 0 {
				s = make([]Value, len(n.Elts))
				for i, ne := range n.Elts {
					sv := sc.evalExpr(ne)
					if _, ok := sv.(error); ok {
						return sv
					}
					s[i] = sv.(Value)
				}
			}
			return Array{
				T: TypeName(t),
				S: s,
			}

		case *ast.MapType:
			kt := tn.Key.(*ast.Ident).Name
			vt := tn.Value.(*ast.Ident).Name
			var s []MapEntry
			if len(n.Elts) > 0 {
				s = make([]MapEntry, len(n.Elts))
				for i, ne := range n.Elts {
					kve := ne.(*ast.KeyValueExpr)
					k := sc.evalExpr(kve.Key)
					if _, ok := k.(error); ok {
						return k
					}
					v := sc.evalExpr(kve.Value)
					if _, ok := v.(error); ok {
						return v
					}
					s[i] = MapEntry{
						K: k.(Value),
						V: v.(Value),
					}
				}
			}
			return Map{
				KT: TypeName(kt),
				VT: TypeName(vt),
				S:  s,
			}
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
				}

			}
			break
		}

		if ii, ok := n.Fun.(*ast.Ident); ok {
			fd := sc.lookup(ii.Name)
			if _, ok := fd.(error); ok {
				return fd
			}
			return sc.makeChild().execStmts(fd.(*ast.FuncDecl).Body.List)
		}

	case *ast.Ident:
		iv := sc.lookup(n.Name)

		if _, ok := iv.(error); ok {
			return iv
		}
		if _, ok := iv.(Value); ok {
			return iv
		}

		if ie, ok := iv.(ast.Expr); ok {
			rv := sc.evalExpr(ie)
			sc.assign(n.Name, rv)
			return rv
		}

	//

	case *ast.UnaryExpr:
		x := sc.evalExpr(n.X)
		if _, ok := x.(error); ok {
			return x
		}
		xb := x.(Basic)
		xc := constant.MakeFromLiteral(xb.S, xb.K.Ast(), 0)
		zc := constant.UnaryOp(n.Op, xc, 0) // FIXME: prec?
		return Basic{
			K: xb.K,
			S: zc.ExactString(),
		}

	case *ast.BinaryExpr:
		x := sc.evalExpr(n.X)
		if _, ok := x.(error); ok {
			return x
		}
		y := sc.evalExpr(n.Y)
		if _, ok := y.(error); ok {
			return y
		}
		xb := x.(Basic)
		yb := y.(Basic)
		if xb.K != yb.K {
			return TypeError{S: []string{xb.K.String(), yb.K.String()}}
		}
		xc := constant.MakeFromLiteral(xb.S, xb.K.Ast(), 0)
		yc := constant.MakeFromLiteral(yb.S, yb.K.Ast(), 0)
		zc := constant.BinaryOp(xc, n.Op, yc)
		return Basic{
			K: xb.K,
			S: zc.ExactString(),
		}

	}
	return Dynamic{}
}

func (sc *Scope) execStmts(sts []ast.Stmt) any {
	check.NotNil(sc.m)
	for _, st := range sts {
		switch st := st.(type) {

		case *ast.ReturnStmt:
			vn := check.Single(st.Results)
			return sc.evalExpr(vn)

		case *ast.AssignStmt:
			tn := check.Single(st.Lhs).(*ast.Ident).Name
			vn := check.Single(st.Rhs)
			vv := sc.evalExpr(vn)
			if _, ok := vv.(error); ok {
				return vv
			}
			sc.assign(tn, vv)

		case *ast.ForStmt:
			check.Nil(st.Init)
			check.Nil(st.Cond)
			check.Nil(st.Post)
			return sc.makeChild().execStmts(st.Body.List) // FIXME: lol...

		case *ast.IfStmt:
			check.Nil(st.Init)
			cv := sc.evalExpr(st.Cond)
			if _, ok := cv.(error); ok {
				return cv
			}
			check.Nil(st.Else)
			panic(st)

		default:
			panic(st)
		}
	}
	panic(sc)
}
