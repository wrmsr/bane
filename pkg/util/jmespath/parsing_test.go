package jmespath

import (
	"fmt"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/wrmsr/bane/pkg/util/jmespath/parser"
)

func TestParsing(t *testing.T) {
	is := antlr.NewInputStream(testExpr)
	lexer := parser.NewJmespathLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewJmespathParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.SingleExpression()

	v := &parseVisitor{}
	root := tree.Accept(v).(Node)
	fmt.Println(root)
}
