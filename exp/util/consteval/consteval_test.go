//
/*
https://github.com/golang/tools/tree/master/go/ssa/interp

TODO:
 - break apart basic - parse all but float
 - ** call **
*/
package consteval

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	ju "github.com/wrmsr/bane/pkg/util/json"
	msh "github.com/wrmsr/bane/pkg/util/marshal"
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

var _ = msh.RegisterTo[Value](
	msh.SetImplOf[Basic](),
	msh.SetImplOf[Type](),
	msh.SetImplOf[Struct](),
	msh.SetImplOf[Array](),
	msh.SetImplOf[Map](),
)

var _ = msh.RegisterTo[BasicKind](
	msh.SetStringerEnumTypes(
		IntBasic,
		FloatBasic,
		ImagBasic,
		CharBasic,
		StringBasic,
	),
)

//

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
			panic(ie)
		}

	}
	panic(node)
}

func TestStuff(t *testing.T) {
	for _, s := range []string{
		"5",
		"\"foo\"",
		"Foo{}",
		"Foo{X: 420}",
		"Type[Foo]()",
		"[]any{1,\"foo\"}",
		"map[any]any{1:\"foo\", \"bar\": 420.0}",
	} {
		a := check.Must1(parser.ParseExpr(s))

		//_ = ast.Fprint(os.Stdout, nil, a, nil)

		v := ValueFromAst(a)

		fmt.Println(check.Must1(ju.MarshalPretty(check.Must1(msh.Marshal(&v)))))

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
}
