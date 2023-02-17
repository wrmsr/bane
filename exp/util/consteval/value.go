package consteval

import (
	"go/ast"
	"go/token"
)

//

type TypeName string

//

type Value interface {
	isValue()
}

type value struct{}

func (v value) isValue() {}

//

type BasicKind int8

const (
	IntBasic BasicKind = iota
	FloatBasic
	ImagBasic
	CharBasic
	StringBasic
)

func (k BasicKind) String() string {
	switch k {
	case IntBasic:
		return "INT"
	case FloatBasic:
		return "FLOAT"
	case ImagBasic:
		return "IMAG"
	case CharBasic:
		return "CHAR"
	case StringBasic:
		return "STRING"
	}
	panic(k)
}

func (k BasicKind) Ast() token.Token {
	switch k {
	case IntBasic:
		return token.INT
	case FloatBasic:
		return token.FLOAT
	case ImagBasic:
		return token.IMAG
	case CharBasic:
		return token.CHAR
	case StringBasic:
		return token.STRING
	}
	panic(k)
}

func basicKindFromAst(tok token.Token) BasicKind {
	switch tok {
	case token.INT:
		return IntBasic
	case token.FLOAT:
		return FloatBasic
	case token.IMAG:
		return ImagBasic
	case token.CHAR:
		return CharBasic
	case token.STRING:
		return StringBasic
	}
	panic(tok)
}

type Basic struct {
	value
	K BasicKind
	S string
}

//

type Struct struct {
	value
	T TypeName
	M map[string]Value
}

//

type Type struct {
	value
	T TypeName
}

//

type Array struct {
	value
	T TypeName
	S []Value
}

//

type MapEntry struct {
	K,
	V Value
}

type Map struct {
	value
	KT,
	VT TypeName
	S []MapEntry
}

//

type Dynamic struct {
	value
	a ast.Node
}
