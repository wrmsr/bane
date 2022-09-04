package python

import (
	"fmt"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/wrmsr/bane/exp/util/python/parser"
)

func TestParser(t *testing.T) {
	//testExpr := `{'descr': '<f4', 'fortran_order': False, 'shape': (3, 3), }`
	testExpr := `420`

	is := antlr.NewInputStream(testExpr)
	lexer := parser.NewExprLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewExprParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Expr()

	fmt.Println(tree)

	v := parseVisitor{}
	n := tree.Accept(&v).(Node)
	fmt.Printf("%+v\n", n)
}
