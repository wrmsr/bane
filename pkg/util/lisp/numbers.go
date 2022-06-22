package lisp

import "fmt"

//

type NumKind int8

const (
	NumInt NumKind = iota
	NumFloat
	NumComplex
)

func (k NumKind) Coerce(vt NumKind) NumKind {
	if vt > k {
		return vt
	} else {
		return k
	}
}

func AsNumber(v Value) Number {
	if r, ok := v.(Number); ok {
		return r
	}
	panic(fmt.Sprintf("object is not a number: %s", v))
}

func AsNumbers(v1 Value, v2 Value) (Number, Number, NumKind) {
	r1, r2 := AsNumber(v1), AsNumber(v2)
	return r1, r2, r1.Kind().Coerce(r2.Kind())
}

func NumberAdd(a Value, b Value) Value {
	switch x, y, vt := AsNumbers(a, b); vt {
	case NumInt:
		return x.AsInt() + y.AsInt()
	case NumFloat:
		return x.AsFloat() + y.AsFloat()
	case NumComplex:
		return x.AsComplex() + y.AsComplex()
	default:
		panic("unreachable")
	}
}
