package python

import (
	"fmt"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/wrmsr/bane/exp/util/python/parser"
	antlru "github.com/wrmsr/bane/pkg/util/antlr"
)

func TestParser(t *testing.T) {
	for _, testExpr := range []string{
		`{'descr': '<f4', 'fortran_order': False, 'shape': (3, 3), }`,

		`420`,
		`None`,
		`"abc"`,
		`'def'`,
		`1+2`,
		`1 + 2`,
		` 1 + 2 `,
		`1+2+3`,
		`1+2/3+4`,
		`+1`,
		`2**2`,
		`1<<2`,
		`[1,2]`,
		`1+(2,3)`,
		`{'descr':'<f4','fortran_order':False,'shape':(3,3),}`,
		`{'descr': '<f4', 'fortran_order': False, 'shape': (3, 3), }`,
		`[1, *[2, 3]]`,
		`x`,
		`x[0]`,
		`x[0:1]`,
		`x[:1]`,
		`x[:1:]`,
		`x[:1:-1]`,
		`x[::-1]`,
	} {
		fmt.Println(testExpr)

		is := antlr.NewInputStream(testExpr)
		lexer := parser.NewExprLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewExprParser(stream)
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
		p.AddErrorListener(antlru.PanicErrorListener{})
		p.BuildParseTrees = true
		tree := p.SingleExpr()

		//antlru.Dump(tree, "")

		v := parseVisitor{}
		n := tree.Accept(&v).(Node)
		fmt.Printf("%#v\n", n)

		fmt.Println(RenderString(n))

		fmt.Println()
	}
}
