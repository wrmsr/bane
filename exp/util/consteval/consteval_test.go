//
/*
https://github.com/golang/tools/tree/master/go/ssa/interp

TODO:
 - array
 - map
 - break apart basic - parse all but float
 - ** call **
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
	ju "github.com/wrmsr/bane/pkg/util/json"
	msh "github.com/wrmsr/bane/pkg/util/marshal"
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

type Struct struct {
	value
	N string
	M map[string]Value
}

//

type Type struct {
	value
	N string
}

//

var _ = msh.RegisterTo[Value](
	msh.SetImplOf[Basic](),
	msh.SetImplOf[Type](),
	msh.SetImplOf[Struct](),
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

func mkValue(node ast.Node) Value {
	switch node := node.(type) {
	case *ast.BasicLit:
		return Basic{
			K: basicKindFromAst(node.Kind),
			S: node.Value,
		}
	case *ast.CompositeLit:
		n := node.Type.(*ast.Ident).Name
		var m map[string]Value
		if len(node.Elts) > 0 {
			m = make(map[string]Value, len(node.Elts))
			for _, e := range node.Elts {
				kv := e.(*ast.KeyValueExpr)
				k := kv.Key.(*ast.Ident).Name
				v := mkValue(kv.Value)
				m[k] = v
			}
		}
		return Struct{
			N: n,
			M: m,
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
					N: tn,
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
	} {
		a := check.Must1(parser.ParseExpr(s))

		_ = ast.Fprint(os.Stdout, nil, a, nil)

		v := mkValue(a)
		fmt.Println(v)

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
