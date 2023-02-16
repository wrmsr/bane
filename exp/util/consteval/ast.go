//
/*
https://github.com/golang/tools/tree/master/go/ssa/interp

TODO:
 - ** call **
 - func Const[T any](v T) { return v } ?
*/
package consteval

import "go/ast"

func ValueFromAst(node ast.Node) Value {
	switch node := node.(type) {

	case *ast.BasicLit:
		return Basic{
			K: basicKindFromAst(node.Kind),
			S: node.Value,
		}

	case *ast.CompositeLit:
		switch tn := node.Type.(type) {

		case *ast.Ident:
			var m map[string]Value
			if len(node.Elts) > 0 {
				m = make(map[string]Value, len(node.Elts))
				for _, e := range node.Elts {
					kv := e.(*ast.KeyValueExpr)
					k := kv.Key.(*ast.Ident).Name
					v := ValueFromAst(kv.Value)
					m[k] = v
				}
			}
			return Struct{
				T: TypeName(tn.Name),
				M: m,
			}

		case *ast.ArrayType:
			t := tn.Elt.(*ast.Ident).Name
			var s []Value
			if len(node.Elts) > 0 {
				s = make([]Value, len(node.Elts))
				for i, e := range node.Elts {
					s[i] = ValueFromAst(e)
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
			if len(node.Elts) > 0 {
				s = make([]MapEntry, len(node.Elts))
				for i, e := range node.Elts {
					kve := e.(*ast.KeyValueExpr)
					k := ValueFromAst(kve.Key)
					v := ValueFromAst(kve.Value)
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
			}
		}

	case *ast.CallExpr:
		if ie, ok := node.Fun.(*ast.IndexExpr); ok {
			fn := ie.X.(*ast.Ident).Name
			switch fn {

			case "Type":
				if len(node.Args) > 0 {
					panic(node)
				}
				tn := ie.Index.(*ast.Ident).Name
				return Type{
					T: TypeName(tn),
				}

			}
			break
		}

	case *ast.Ident:
		return Ident{
			N: node.Name,
		}

	}
	return Dynamic{}
}
