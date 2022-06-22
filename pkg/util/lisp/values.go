package lisp

import (
	"fmt"
	"strconv"
	"strings"

	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type Value interface {
	IsIdentity() bool

	String() string

	isValue()
}

type Number interface {
	Value
	Kind() NumKind
	AsInt() Int
	AsFloat() Float
	AsComplex() Complex
}

//

type Cons struct {
	Car, Cdr Value
}

//

type (
	Bool    bool
	Int     int64
	Float   float64
	Complex complex128
	Char    rune
	String  string
	Atom    string
)

var _ Value = bt.Zero[Bool]()
var _ Value = bt.Zero[Int]()
var _ Value = bt.Zero[Float]()
var _ Value = bt.Zero[Complex]()
var _ Value = bt.Zero[Char]()
var _ Value = bt.Zero[String]()
var _ Value = bt.Zero[Atom]()
var _ Value = bt.Zero[*Cons]()

func (v Bool) isValue()    {}
func (v Int) isValue()     {}
func (v Float) isValue()   {}
func (v Complex) isValue() {}
func (v Char) isValue()    {}
func (v String) isValue()  {}
func (v Atom) isValue()    {}
func (v *Cons) isValue()   {}

func (Bool) IsIdentity() bool    { return true }
func (Int) IsIdentity() bool     { return true }
func (Float) IsIdentity() bool   { return true }
func (Complex) IsIdentity() bool { return true }
func (Char) IsIdentity() bool    { return true }
func (String) IsIdentity() bool  { return true }
func (Atom) IsIdentity() bool    { return false }
func (*Cons) IsIdentity() bool   { return false }

func (v Bool) String() string {
	if v {
		return "#t"
	} else {
		return "#f"
	}
}

func (v Int) String() string {
	return strconv.Itoa(int(v))
}

func (v Float) String() string {
	vv := strconv.FormatFloat(float64(v), 'g', -1, 64)
	vp := strings.Split(vv, "e")
	if strings.ContainsRune(vp[0], '.') {
		return vv
	}
	vp[0] += ".0"
	return strings.Join(vp, "e")
}

func (v Complex) String() string {
	if im := imag(v); im >= 0 {
		return fmt.Sprintf("%g+%gi", real(v), im)
	} else {
		return fmt.Sprintf("%g-%gi", real(v), -im)
	}
}

func (v Char) String() string {
	switch v {
	case ' ':
		return `#\space`
	case '\n':
		return `#\newline`
	case '\b':
		return `#\backspace`
	case '\t':
		return `#\tab`
	case '\f':
		return `#\page`
	case '\r':
		return `#\return`
	case 0x7f:
		return `#\rubout`
	default:
		return `#\` + string(v)
	}
}

func (v String) String() string {
	return strconv.Quote(string(v))
}

func (v Atom) String() string {
	return string(v)
}

func (v *Cons) String() string {
	var ok bool
	var vv *Cons
	var rb []string
	if vv = v; vv == nil {
		return "()"
	}

	for vv != nil {
		d := vv.Cdr
		s := AsString(vv.Car)
		vv, ok = AsCons(vv.Cdr)
		if rb = append(rb, s); !ok {
			rb = append(rb, ".", AsString(d))
		}
	}

	return fmt.Sprintf("(%s)", strings.Join(rb, " "))
}

//

func AsValue(v any) Value {
	if v == nil {
		return nil
	}
	if v, ok := v.(Value); ok {
		return v
	}

	switch v := v.(type) {

	case bool:
		return Bool(v)

	case int:
		return Int(v)
	case int8:
		return Int(v)
	case int16:
		return Int(v)
	case int32:
		return Int(v)
	case int64:
		return Int(v)
	case uint8:
		return Int(v)
	case uint16:
		return Int(v)
	case uint32:
		return Int(v)
	case uint64:
		return Int(v)

	case float32:
		return Float(v)
	case float64:
		return Float(v)

	case complex64:
		return Complex(v)
	case complex128:
		return Complex(v)

	case string:
		return String(v)
	}

	panic(fmt.Errorf("can't make value: %T %v", v, v))
}

func AsCons(v Value) (*Cons, bool) {
	r, ok := v.(*Cons)
	return r, ok || v == nil
}

func AsString(v Value) string {
	if v == nil {
		return "()"
	} else {
		return v.String()
	}
}

func MakeList(vs ...Value) *Cons {
	var p, q *Cons
	for _, v := range vs {
		AppendValue(&p, &q, v)
	}
	return p
}

func AppendValue(p **Cons, q **Cons, v Value) {
	if *p == nil {
		*p = new(Cons)
		*q, (*p).Car = *p, v
	} else {
		r := new(Cons)
		r.Car, (*q).Cdr, *q = v, r, r
	}
}
