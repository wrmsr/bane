package consteval

import (
	"go/ast"
	"go/constant"

	"github.com/wrmsr/bane/pkg/util/check"
)

type exprEvalContext struct {
	lookup   func(name string) any // Value | error | ast.Node
	evalExpr func(node ast.Expr) (Value, error)
}

func evalExpr(ctx *exprEvalContext, n ast.Expr) (Value, error) {
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
					v, err := ctx.evalExpr(kv.Value)
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
					sv, err := ctx.evalExpr(ne)
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
					k, err := ctx.evalExpr(kve.Key)
					if err != nil {
						return nil, err
					}
					v, err := ctx.evalExpr(kve.Value)
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
				check.Equal(len(n.Args), 0)
				tn := ie.Index.(*ast.Ident).Name
				return Type{
					T: TypeName(tn),
				}, nil

			}
			break
		}

		//if ii, ok := n.Fun.(*ast.Ident); ok {
		//	check.Equal(len(n.Args), 0)
		//	fd, err := sc.lookupAst(ii.Name)
		//	if err != nil {
		//		return nil, err
		//	}
		//	fe := newFuncEval(fd.(*ast.FuncDecl), sc)
		//	return fe.eval()
		//}

	case *ast.Ident:
		iv := sc.m[n.Name]
		if iv == nil {
			l, err := sc.lookupAst(n.Name)
			if err != nil {
				return nil, err
			}
			iv = l
		}

		if v2, ok := iv.(Value); ok {
			return v2, nil
		}

		if en, ok := iv.(ast.Expr); ok {
			rv, err := ctx.evalExpr(en)
			if err != nil {
				sc.m[n.Name] = err
				return nil, err
			}
			v2 := rv.(Value)
			sc.m[n.Name] = v2
			return v2, nil
		}

	//

	case *ast.UnaryExpr:
		x, err := ctx.evalExpr(n.X)
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
		x, err := ctx.evalExpr(n.X)
		if err != nil {
			return nil, err
		}
		y, err := ctx.evalExpr(n.Y)
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
