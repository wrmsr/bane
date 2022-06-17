package jmespath

import (
	"fmt"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/wrmsr/bane/pkg/utils/jmespath/parser"
)

type TreeShapeListener struct {
	*parser.BaseJmespathListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (l *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

const testExpr = "people[?age > `20`].[name, age]"

func TestParser1(t *testing.T) {
	is := antlr.NewInputStream(testExpr)
	lexer := parser.NewJmespathLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewJmespathParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.SingleExpression()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}

func TestParser2(t *testing.T) {
	is := antlr.NewInputStream(testExpr)
	lexer := parser.NewJmespathLexer(is)
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n", lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}
