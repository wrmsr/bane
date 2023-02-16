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

	ov := check.Must1(s.Reduce("foo"))
	fmt.Println(check.Must1(ju.MarshalPretty(check.Must1(msh.Marshal(&ov)))))
}

func findValueSpecExpr(vs *ast.ValueSpec, n string) ast.Expr {
	for i, vn := range vs.Names {
		if vn.Name == n {
			return vs.Values[i]
		}
	}
	return nil
}

func TestScopeFile(t *testing.T) {
	src := `
package foo

const foo = 10
const bar = foo
var a = []any{foo, bar}
var m = map[any]any{"F": foo, "B": bar}
var baz = Brax{F: foo, B: bar}
`

	const mode = parser.AllErrors
	var fset token.FileSet
	fil := check.Must1(parser.ParseFile(&fset, "", src, mode))
	//_ = ast.Fprint(os.Stdout, nil, fil, nil)

	s := &Scope{
		m: make(map[string]any),
	}
	for n, obj := range fil.Scope.Objects {
		switch obj.Kind {
		case ast.Con, ast.Var:
			s.m[n] = findValueSpecExpr(obj.Decl.(*ast.ValueSpec), obj.Name)
		}
	}

	v := check.Must1(s.Reduce(&ast.Ident{Name: "foo"}))
	fmt.Println(check.Must1(ju.MarshalPretty(check.Must1(msh.Marshal(&v)))))
}
