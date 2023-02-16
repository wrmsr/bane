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

func (e *Scope) Reduce(o any) (Value, error) {
	if s, ok := o.(string); ok {
		o = Ident{N: s}
	}

	if err, ok := o.(error); ok {
		return nil, err
	}

	if a, ok := o.(ast.Node); ok {
		v := ValueFromAst(a)
		return e.Reduce(v)
	}

	if v, ok := o.(Value); ok {
		switch v := v.(type) {

		case Ident:
			mv := e.m[v.N]
			if mv == nil {
				return nil, IdentError{N: v.N}
			}
			v2, err := e.Reduce(mv)
			if err != nil {
				e.m[v.N] = err
				return nil, err
			}
			e.m[v.N] = v2
			return v2, nil

		default:
			return v, nil
		}
	}

	panic(o)
}
