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
	"go/token"

	"github.com/wrmsr/bane/pkg/util/check"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type Lookup interface {
	Lookup(n string) any // Value | ast.Node | error
}

//

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
		m: make(map[string]bt.Result[Value]),
	}
}

// FIXME: lol
func boolString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func stringBool(s string) bool {
	switch s {
	case "true":
		return true
	case "false":
		return false
	}
	panic(s)
}

//

func (sc *Scope) evalExpr(n ast.Expr) effect {
	switch n := n.(type) {

	case *ast.BasicLit:
		return valEffect(Basic{
			K: basicKindFromAst(n.Kind),
			S: n.Value,
		})

	case *ast.CompositeLit:
		switch tn := n.Type.(type) {

		case *ast.Ident:
			var m map[string]Value
			if len(n.Elts) > 0 {
				m = make(map[string]Value, len(n.Elts))
				for _, ne := range n.Elts {
					kv := ne.(*ast.KeyValueExpr)
					k := kv.Key.(*ast.Ident).Name
					v := sc.evalExpr(kv.Value).expr()
					if v.Err != nil {
						return v
					}
					m[k] = v.mustVal()
				}
			}
			return valEffect(Struct{
				T: TypeName(tn.Name),
				M: m,
			})

		case *ast.ArrayType:
			t := tn.Elt.(*ast.Ident).Name
			var s []Value
			if len(n.Elts) > 0 {
				s = make([]Value, len(n.Elts))
				for i, ne := range n.Elts {
					v := sc.evalExpr(ne).expr()
					if v.Err != nil {
						return v
					}
					s[i] = v.mustVal()
				}
			}
			return valEffect(Array{
				T: TypeName(t),
				S: s,
			})

		case *ast.MapType:
			kt := tn.Key.(*ast.Ident).Name
			vt := tn.Value.(*ast.Ident).Name
			var s []MapEntry
			if len(n.Elts) > 0 {
				s = make([]MapEntry, len(n.Elts))
				for i, ne := range n.Elts {
					kve := ne.(*ast.KeyValueExpr)
					k := sc.evalExpr(kve.Key).expr()
					if k.Err != nil {
						return k
					}
					v := sc.evalExpr(kve.Value).expr()
					if v.Err != nil {
						return v
					}
					s[i] = MapEntry{
						K: k.mustVal(),
						V: v.mustVal(),
					}
				}
			}
			return valEffect(Map{
				KT: TypeName(kt),
				VT: TypeName(vt),
				S:  s,
			})
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
				return valEffect(Type{
					T: TypeName(tn),
				})

			}
			break
		}

		if ii, ok := n.Fun.(*ast.Ident); ok {
			fd := sc.lookup(ii.Name)
			if err, ok := fd.(error); ok {
				return errEffect(err)
			}
			rv := sc.makeChild().execStmts(fd.(*ast.FuncDecl).Body.List)
			if rv.flow != returnFlow {
				panic(rv)
			}
			return effect{Result: rv.Result}
		}

	case *ast.Ident:
		iv := sc.lookup(n.Name)

		if ix, ok := iv.(bt.Result[Value]); ok {
			return effect{Result: ix}
		}

		if ie, ok := iv.(ast.Expr); ok {
			rv := sc.evalExpr(ie).expr()
			if rv.Err != nil {
				return rv
			}
			sc.assign(n.Name, rv.Result)
			return rv
		}

		panic("???")

	//

	case *ast.UnaryExpr:
		x := sc.evalExpr(n.X).expr()
		if x.Err != nil {
			return x
		}
		xb := x.mustVal().(Basic)
		xc := constant.MakeFromLiteral(xb.S, xb.K.Ast(), 0)
		zc := constant.UnaryOp(n.Op, xc, 0) // FIXME: prec?
		return valEffect(Basic{
			K: xb.K,
			S: zc.ExactString(),
		})

	case *ast.BinaryExpr:
		x := sc.evalExpr(n.X).expr()
		if x.Err != nil {
			return x
		}
		y := sc.evalExpr(n.Y).expr()
		if y.Err != nil {
			return y
		}
		xb := x.mustVal().(Basic)
		yb := y.mustVal().(Basic)
		if xb.K != yb.K {
			return errEffect(TypeError{S: []string{xb.K.String(), yb.K.String()}})
		}
		xc := constant.MakeFromLiteral(xb.S, xb.K.Ast(), 0)
		yc := constant.MakeFromLiteral(yb.S, yb.K.Ast(), 0)
		switch n.Op {
		case token.EQL,
			token.NEQ,
			token.LSS,
			token.LEQ,
			token.GTR,
			token.GEQ:
			b := constant.Compare(xc, n.Op, yc)
			return valEffect(Basic{
				K: BoolBasic,
				S: boolString(b),
			})
		default:
			zc := constant.BinaryOp(xc, n.Op, yc)
			return valEffect(Basic{
				K: xb.K,
				S: zc.ExactString(),
			})
		}

	}
	return valEffect(Dynamic{})
}

func (sc *Scope) execStmts(sts []ast.Stmt) effect {
	for _, st := range sts {
		e := sc.execStmt(st)
		if e.Err != nil || e.flow != noFlow {
			return e
		}
	}
	return effect{}
}

func (sc *Scope) execStmt(st ast.Stmt) effect {
	switch st := st.(type) {

	case *ast.ReturnStmt:
		vn := check.Single(st.Results)
		ve := sc.evalExpr(vn)
		if ve.Err != nil || ve.flow != noFlow {
			return ve
		}
		return effect{Result: ve.Result, flow: returnFlow}

	case *ast.AssignStmt:
		tn := check.Single(st.Lhs).(*ast.Ident).Name
		vn := check.Single(st.Rhs)
		vv := sc.evalExpr(vn).expr()
		if vv.Err != nil {
			return vv
		}
		sc.assign(tn, vv.Result)

	case *ast.BranchStmt:
		// BREAK, CONTINUE, GOTO, FALLTHROUGH
		if st.Label != nil {
			panic(st)
		}
		switch st.Tok {
		case token.BREAK:
			return effect{flow: breakFlow}
		default:
			panic(st)
		}

	case *ast.ForStmt:
		check.Nil(st.Init)
		check.Nil(st.Cond)
		check.Nil(st.Post)
		for {
			e := sc.makeChild().execStmts(st.Body.List)
			if e.Err != nil {
				return e
			}
			switch e.flow {
			case noFlow:
			case breakFlow:
				return effect{}
			case returnFlow:
				return e
			default:
				panic(e)
			}
		}

	case *ast.IfStmt:
		check.Nil(st.Init)
		cv := sc.evalExpr(st.Cond).expr()
		if cv.Err != nil {
			return cv
		}
		cbv := cv.mustVal().(Basic)
		check.Equal(cbv.K, BoolBasic)
		cb := stringBool(cbv.S)
		if cb {
			e := sc.makeChild().execStmts(st.Body.List)
			if e.Err != nil || e.flow != noFlow {
				return e
			}
		} else if st.Else != nil {
			panic(st)
		}

	default:
		panic(st)
	}

	return effect{}
}
