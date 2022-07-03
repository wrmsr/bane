package lisp

import (
	"fmt"
	"math"
)

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

//

type Number interface {
	Value
	Kind() NumKind
	AsInt() Int
	AsFloat() Float
	AsComplex() Complex
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

//

func (v Int) Kind() NumKind     { return NumInt }
func (v Float) Kind() NumKind   { return NumFloat }
func (v Complex) Kind() NumKind { return NumComplex }

func (v Int) AsInt() Int     { return v }
func (v Float) AsInt() Int   { return Int(v) }
func (v Complex) AsInt() Int { return Int(v.AsRealNumber()) }

func (v Int) AsFloat() Float     { return Float(v) }
func (v Float) AsFloat() Float   { return v }
func (v Complex) AsFloat() Float { return Float(v.AsRealNumber()) }

func (v Int) AsComplex() Complex     { return Complex(complex(float64(v), 0)) }
func (v Float) AsComplex() Complex   { return Complex(complex(v, 0)) }
func (v Complex) AsComplex() Complex { return v }

func (v Complex) Magnitude() Float {
	return Float(math.Hypot(real(v), imag(v)))
}

func (v Complex) AsRealNumber() float64 {
	cv := complex128(v)
	if imag(cv) != 0 {
		panic("cannot convert complex numbers with non-zero imaginary part into real numbers")
	}
	return real(cv)
}

//

func NumberAdd(a Value, b Value) Value {
	switch x, y, vt := AsNumbers(a, b); vt {
	case NumInt:
		return x.AsInt() + y.AsInt()
	case NumFloat:
		return x.AsFloat() + y.AsFloat()
	case NumComplex:
		return x.AsComplex() + y.AsComplex()
	}
	panic("unreachable")
}

func NumberLt(a Value, b Value) bool {
	switch x, y, vt := AsNumbers(a, b); vt {
	case NumInt:
		return x.AsInt() < y.AsInt()
	case NumFloat:
		return x.AsFloat() < y.AsFloat()
	case NumComplex:
		panic("complex numbers can only be compared for equality")
	}
	panic("unreachable")
}

func NumberGt(a Value, b Value) bool {
	switch x, y, vt := AsNumbers(a, b); vt {
	case NumInt:
		return x.AsInt() > y.AsInt()
	case NumFloat:
		return x.AsFloat() > y.AsFloat()
	case NumComplex:
		panic("complex numbers can only be compared for equality")
	}
	panic("unreachable")
}
