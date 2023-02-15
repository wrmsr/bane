//
/*
https://github.com/golang/tools/tree/master/go/ssa/interp
*/
package consteval

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

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

type Type struct {
	value
	N string
}

//

type Struct struct {
	value
	N string
	M map[string]Value
}

//

func mkValue(node ast.Node) Value {
	switch node := node.(type) {
	case *ast.BasicLit:
		return Basic{
			K: basicKindFromAst(node.Kind),
			S: node.Value,
		}
	}
	panic(node)
}

func TestStuff(t *testing.T) {
	a := check.Must1(parser.ParseExpr("\"foo\""))

	_ = ast.Fprint(os.Stdout, nil, a, nil)

	v := mkValue(a)
	fmt.Println(v)

	//ast.Inspect(a, func(node ast.Node) bool {
	//	switch node := node.(type) {
	//	case *ast.BasicLit:
	//
	//	default:
	//		panic(node)
	//	}
	//	return true
	//})
}
