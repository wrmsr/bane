package jmespath

import (
	"fmt"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/wrmsr/bane/pkg/util/jmespath/parser"
)

//

type TreeShapeListener struct {
	*parser.BaseJmespathListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (l *TreeShapeListener) EnterIdentifier(ctx *parser.IdentifierContext) {
	fmt.Println("identifier!")
}

func (l *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

//

type TreeVisitor struct {
	parser.BaseJmespathVisitor
}

func (v *TreeVisitor) Visit(tree antlr.ParseTree) interface{}            { return tree.Accept(v) }
func (v *TreeVisitor) VisitChildren(node antlr.RuleNode) interface{}     { return nil }
func (v *TreeVisitor) VisitTerminal(node antlr.TerminalNode) interface{} { return nil }
func (v *TreeVisitor) VisitErrorNode(node antlr.ErrorNode) interface{}   { return nil }

var _ parser.JmespathVisitor = &TreeVisitor{}

func (v *TreeVisitor) VisitSingleExpression(ctx *parser.SingleExpressionContext) interface{} {
	return v.Visit(ctx.Expression())
}

func (v *TreeVisitor) VisitIdentifier(ctx *parser.IdentifierContext) interface{} {
	fmt.Println(ctx)
	return v.BaseJmespathVisitor.VisitIdentifier(ctx)
}

//func (v *TreeVisitor) VisitChildren(node antlr.RuleNode) interface{} {
//	for i := 0; i < node.GetChildCount(); i++ {
//		child := node.GetChild(i)
//		if pt, ok := child.(antlr.ParseTree); ok {
//			pt.Accept(v)
//		}
//	}
//	return nil
//}

//

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

	v := &TreeVisitor{}
	//v.BaseParseTreeVisitor = v
	tree.Accept(v)
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
