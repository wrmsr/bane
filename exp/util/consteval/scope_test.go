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

//func TestSimple(t *testing.T) {
//	for _, s := range []string{
//		"5",
//		"\"foo\"",
//		"Foo{}",
//		"Foo{X: 420}",
//		"Type[Foo]()",
//		"[]any{1,\"foo\"}",
//		"map[any]any{1:\"foo\", \"bar\": 420.0}",
//		"foo",
//	} {
//		a := check.Must1(parser.ParseExpr(s))
//
//		//_ = ast.Fprint(os.Stdout, nil, a, nil)
//
//		v := ValueFromAst(a)
//
//		fmt.Println(check.Must1(ju.MarshalPretty(check.Must1(msh.Marshal(&v)))))
//
//		//ast.Inspect(a, func(node ast.Node) bool {
//		//	switch node := node.(type) {
//		//	case *ast.BasicLit:
//		//
//		//	default:
//		//		panic(node)
//		//	}
//		//	return true
//		//})
//	}
//}

//func TestBigStuff(t *testing.T) {
//	src := `
//package foo
//
//const foo = 10
//
//func bar() int {
//	return 420
//}
//`
//
//	const mode = parser.AllErrors
//	var fset token.FileSet
//	a := check.Must1(parser.ParseFile(&fset, "", src, mode))
//	_ = ast.Fprint(os.Stdout, nil, a, nil)
//}

func TestEval(t *testing.T) {
	s := &Scope{
		m: map[string]any{
			"foo": Basic{K: IntBasic, S: "420"},
			"bar": &ast.Ident{Name: "foo"},
		},
	}

	ov := s.Eval("foo").(Value)
	fmt.Println(check.Must1(ju.MarshalPretty(check.Must1(msh.Marshal(&ov)))))
}

func TestScopeFile(t *testing.T) {
	src := `
package foo

const foo = 10
const bar = foo
var n any = nil
var a = []any{foo, bar}
var m = map[any]any{"F": foo, "B": bar, "A": a}
var baz = Brax{F: foo, B: bar, A: a, M: m}
const bar2 = 2 * bar

func junk() int {
	x := foo
	x = x + 1
	return 4 + x
}
var junkv = junk()
`

	const mode = parser.AllErrors
	var fset token.FileSet
	fil := check.Must1(parser.ParseFile(&fset, "", src, mode))
	//_ = ast.Fprint(os.Stdout, nil, fil, nil)

	s := ScopeFromAst(fil.Scope)

	for _, n := range []string{
		"baz",
		//"n", // FIXME: lol
		"bar2",
		"junkv",
	} {
		v := s.Eval(&ast.Ident{Name: n}).(Value)
		fmt.Println(check.Must1(ju.MarshalPretty(check.Must1(msh.Marshal(&v)))))
	}
}
