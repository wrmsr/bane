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

func TestEval(t *testing.T) {
	s := &Scope{
		m: map[string]any{
			"foo": Basic{K: IntBasic, S: "420"},
			"bar": Ident{N: "foo"},
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
`

	const mode = parser.AllErrors
	var fset token.FileSet
	fil := check.Must1(parser.ParseFile(&fset, "", src, mode))
	_ = ast.Fprint(os.Stdout, nil, fil, nil)

	s := &Scope{
		m: make(map[string]any),
	}
	for n, obj := range fil.Scope.Objects {
		av := findValueSpecExpr(obj.Decl.(*ast.ValueSpec), obj.Name)
		v := ValueFromAst(av)
		s.m[n] = v
	}

	fmt.Println(s.Reduce("bar"))
}
