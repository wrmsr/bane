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

func TestStuff(t *testing.T) {
	for _, s := range []string{
		"5",
		"\"foo\"",
		"Foo{}",
		"Foo{X: 420}",
		"Type[Foo]()",
		"[]any{1,\"foo\"}",
		"map[any]any{1:\"foo\", \"bar\": 420.0}",
		"foo",
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

func TestBigStuff(t *testing.T) {
	src := `
package foo

const foo = 10

func bar() int {
	return 420
}
`

	const mode = parser.AllErrors
	var fset token.FileSet
	a := check.Must1(parser.ParseFile(&fset, "", src, mode))
	_ = ast.Fprint(os.Stdout, nil, a, nil)
}
