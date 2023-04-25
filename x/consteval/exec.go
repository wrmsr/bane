package consteval

import (
	"go/ast"
	"go/constant"
	"go/token"

	"github.com/wrmsr/bane/core/check"
	bt "github.com/wrmsr/bane/core/types"
)

type executor struct {
	sc *Scope

	synthFuncs map[string]struct{}
}

func (ex *executor) makeChild() *executor {
	ret := *ex
	ex.sc = ex.sc.makeChild()
	return &ret
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

func (ex *executor) evalExpr(n ast.Expr) effect {
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
					v := ex.evalExpr(kv.Value).expr()
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
					v := ex.evalExpr(ne).expr()
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
					k := ex.evalExpr(kve.Key).expr()
					if k.Err != nil {
						return k
					}
					v := ex.evalExpr(kve.Value).expr()
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
			if ex.synthFuncs != nil {
				if _, ok := ex.synthFuncs[ii.Name]; ok {
					var a []Value
					if len(n.Args) > 0 {
						a = make([]Value, len(n.Args))
						for i, ca := range n.Args {
							cae := ex.evalExpr(ca)
							if cae.Err != nil || cae.flow != noFlow {
								return cae
							}
							a[i] = cae.mustVal()
						}
					}
					return valEffect(Call{
						N: ii.Name,
						A: a,
					})
				}
			}

			if len(n.Args) > 0 {
				panic(n)
			}
			fd := ex.sc.lookup(ii.Name)
			if err, ok := fd.(error); ok {
				return errEffect(err)
			}
			rv := ex.makeChild().execStmts(fd.(*ast.FuncDecl).Body.List)
			if rv.flow != returnFlow {
				panic(rv)
			}
			return effect{Result: rv.Result}
		}

	case *ast.Ident:
		iv := ex.sc.lookup(n.Name)

		if ix, ok := iv.(bt.Result[Value]); ok {
			return effect{Result: ix}
		}

		if ie, ok := iv.(ast.Expr); ok {
			rv := ex.evalExpr(ie).expr()
			if rv.Err != nil {
				return rv
			}
			ex.sc.assign(n.Name, rv.Result)
			return rv
		}

		panic("???")

	//

	case *ast.UnaryExpr:
		x := ex.evalExpr(n.X).expr()
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
		x := ex.evalExpr(n.X).expr()
		if x.Err != nil {
			return x
		}
		y := ex.evalExpr(n.Y).expr()
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

func (ex *executor) execStmts(sts []ast.Stmt) effect {
	for _, st := range sts {
		e := ex.execStmt(st)
		if e.Err != nil || e.flow != noFlow {
			return e
		}
	}
	return effect{}
}

func (ex *executor) execStmt(st ast.Stmt) effect {
	switch st := st.(type) {

	case *ast.ReturnStmt:
		vn := check.Single(st.Results)
		ve := ex.evalExpr(vn)
		if ve.Err != nil || ve.flow != noFlow {
			return ve
		}
		return effect{Result: ve.Result, flow: returnFlow}

	case *ast.AssignStmt:
		tn := check.Single(st.Lhs).(*ast.Ident).Name
		vn := check.Single(st.Rhs)
		vv := ex.evalExpr(vn).expr()
		if vv.Err != nil {
			return vv
		}
		ex.sc.assign(tn, vv.Result)

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
			e := ex.makeChild().execStmts(st.Body.List)
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
		cv := ex.evalExpr(st.Cond).expr()
		if cv.Err != nil {
			return cv
		}
		cbv := cv.mustVal().(Basic)
		check.Equal(cbv.K, BoolBasic)
		cb := stringBool(cbv.S)
		if cb {
			e := ex.makeChild().execStmts(st.Body.List)
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
