//
/*
https://github.com/golang/tools/tree/master/go/ssa/interp
https://github.com/llvm/llvm-project/tree/main/clang/lib/AST/Interp

IMMED:
 - toplevel instantiations

TODO:
 - type-specific default values..
 - more lazy - mirror ast.Scope re inheritance
 - togglable memoization - consteval func scopes
 - func
 - addr - *including* var refs, like &foo[2]
 - func Const[T any](v T) { return v } ?
  - if consteval then _no_ as running funcs lol
 - intrinsics?
  - math
  - strings
  - reflect - fields, etc
  - io? encoding?
 - views interop - frozen slices
  - inlineable? lol fuckin go
 - Call values - tack type names onto args? optionally?
*/
package consteval

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/wrmsr/bane/core/check"
	ju "github.com/wrmsr/bane/core/json"
	msh "github.com/wrmsr/bane/core/marshal"
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

//func TestEval(t *testing.T) {
//	s := &Scope{
//		m: map[string]bt.Result[Value]{
//			"foo": bt.Ok[Value](Basic{K: IntBasic, S: "420"}),
//			"bar": &ast.Ident{Name: "foo"},
//		},
//	}
//
//	ov := s.Eval("foo").(Value)
//	fmt.Println(check.Must1(ju.MarshalPretty(check.Must1(msh.Marshal(&ov)))))
//}

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

func ljunk() int {
	// var i int // FIXME:
	i := 0
	for {
		i = i + 1
		if i > 3 {
			// return i
			break
		}
	}
	return 0
}
var ljunkv = ljunk()
var ahem = Bluh[int, float](420)
`

	const mode = parser.AllErrors
	var fset token.FileSet
	fil := check.Must1(parser.ParseFile(&fset, "", src, mode))
	//_ = ast.Fprint(os.Stdout, nil, fil, nil)

	s := ScopeFromAst(fil.Scope)

	sfs := make(map[string]struct{})
	sfs["Bluh"] = struct{}{}

	for _, n := range []string{
		//"baz",
		//"bar2",
		"junkv",
		"ljunkv",
		"ahem",
		//"n", // FIXME: lol
	} {
		v := (&executor{sc: s, synthFuncs: sfs}).evalExpr(&ast.Ident{Name: n})
		fmt.Println(check.Must1(ju.MarshalPretty(check.Must1(msh.Marshal(&v.Val)))))
	}
}
