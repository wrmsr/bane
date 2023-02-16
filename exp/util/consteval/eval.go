package consteval

import "fmt"

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
	p *Scope
}

/*
func (e *Scope) Eval(v Value) (Value, error) {
	switch v := v.(type) {

	case Basic:
		return v, nil

	case Ident:
		if v2, ok := e.m[v.N]; ok {
			return v2.(Value), nil
		}
		return nil, IdentError{N: v.N}

	}
	panic(v)
}
*/

func (e *Scope) Reduce(n string) (Value, error) {
	s, ok := e.m[n]
	if !ok {
		return nil, IdentError{N: n}
	}

	if i, ok := s.(Ident); ok {
		v, err := e.Reduce(i.N)
		if err != nil {
			e.m[n] = err
			return nil, err
		}
		e.m[n] = v
		return v, nil
	}

	if v, ok := s.(Value); ok {
		return v, nil
	}

	panic(s)
}
