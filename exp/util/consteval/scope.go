//
/*
https://github.com/golang/tools/tree/master/go/ssa/interp
https://github.com/llvm/llvm-project/tree/main/clang/lib/AST/Interp

TODO:
 - func
 - addr - *including* var refs, like &foo[2]
 - func Const[T any](v T) { return v } ?
  - if consteval then _no_ as running funcs lol
*/
package consteval

import (
	"fmt"
	"go/ast"
)

//

type IdentError struct {
	N string
}

func (e IdentError) Error() string {
	return fmt.Sprintf("unresolved identifier: %s", e.N)
}

//

type Scope struct {
	m map[string]any // Value | ast.Node | error
	//p *Scope
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
	s := &Scope{
		m: make(map[string]any),
	}
	for n, obj := range a.Objects {
		switch obj.Kind {
		case ast.Con, ast.Var:
			s.m[n] = findValueSpecExpr(obj.Decl.(*ast.ValueSpec), obj.Name)
		}
	}
	return s
}

func (e *Scope) Reduce(o any) (Value, error) {
	if v, ok := o.(Value); ok {
		return v, nil
	}

	if err, ok := o.(error); ok {
		return nil, err
	}

	if n, ok := o.(ast.Node); ok {
		return e.reduceAst(n)
	}

	panic(o)
}

func (e *Scope) reduceAst(n ast.Node) (Value, error) {
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
					v, err := e.reduceAst(kv.Value)
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
					sv, err := e.reduceAst(ne)
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
					k, err := e.reduceAst(kve.Key)
					if err != nil {
						return nil, err
					}
					v, err := e.reduceAst(kve.Value)
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

	case *ast.Ident:
		iv := e.m[n.Name]
		if iv == nil {
			return nil, IdentError{N: n.Name}
		}

		if v2, ok := iv.(Value); ok {
			return v2, nil
		}

		rv, err := e.Reduce(iv)
		if err != nil {
			e.m[n.Name] = err
			return nil, err
		}
		v2 := rv.(Value)
		e.m[n.Name] = v2
		return v2, nil

	}
	return Dynamic{}, nil
}
