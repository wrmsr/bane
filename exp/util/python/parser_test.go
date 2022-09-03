package python

import (
	"fmt"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/wrmsr/bane/exp/util/python/parser"
)

type parseVisitor struct{}

var _ parser.ExprVisitor = &parseVisitor{}

func (v *parseVisitor) defaultResult() any {
	return nil
}

func (v *parseVisitor) aggregateResult(aggregate, nextResult any) any {
	if aggregate != nil {
		panic("should be nil")
	}
	if nextResult == nil {
		panic("must not be nil")
	}
	return nextResult
}

func (v *parseVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *parseVisitor) VisitChildren(node antlr.RuleNode) any {
	var result = v.defaultResult()
	for i := 0; i < node.GetChildCount(); i++ {
		child := node.GetChild(i)
		if pt, ok := child.(antlr.ParseTree); ok {
			childResult := pt.Accept(v)
			result = v.aggregateResult(result, childResult)
		}
	}
	return result
}

func (v *parseVisitor) VisitTerminal(node antlr.TerminalNode) any {
	panic("implement me")
}

func (v *parseVisitor) VisitErrorNode(node antlr.ErrorNode) any {
	panic("implement me")
}

func (v *parseVisitor) VisitExpr(ctx *parser.ExprContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitOrTest(ctx *parser.OrTestContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitAndTest(ctx *parser.AndTestContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitNotTest(ctx *parser.NotTestContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitComparison(ctx *parser.ComparisonContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitCompOp(ctx *parser.CompOpContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitExprMain(ctx *parser.ExprMainContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitExprCont(ctx *parser.ExprContContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitXorExpr(ctx *parser.XorExprContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitXorExprCont(ctx *parser.XorExprContContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitAndExpr(ctx *parser.AndExprContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitAndExprCont(ctx *parser.AndExprContContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitShiftExpr(ctx *parser.ShiftExprContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitShiftExprCont(ctx *parser.ShiftExprContContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitArithExpr(ctx *parser.ArithExprContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitArithExprCont(ctx *parser.ArithExprContContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitTerm(ctx *parser.TermContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitTermCont(ctx *parser.TermContContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitFactor(ctx *parser.FactorContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitPower(ctx *parser.PowerContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitAtomExpr(ctx *parser.AtomExprContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitParenAtom(ctx *parser.ParenAtomContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitBraacketAtom(ctx *parser.BraacketAtomContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitDictOrSetAtom(ctx *parser.DictOrSetAtomContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitConstAtom(ctx *parser.ConstAtomContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitConst(ctx *parser.ConstContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitTestListComp(ctx *parser.TestListCompContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitTrailer(ctx *parser.TrailerContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitSubscriptList(ctx *parser.SubscriptListContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitSubscript(ctx *parser.SubscriptContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitSliceOp(ctx *parser.SliceOpContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitDictOrSetMaker(ctx *parser.DictOrSetMakerContext) any {
	panic("implement me")
}

func TestParser(t *testing.T) {
	testExpr := `420 * 69`

	is := antlr.NewInputStream(testExpr)
	lexer := parser.NewExprLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewExprParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Expr()

	fmt.Println(tree)

	//v := MyExprVisitor{}
	//v.Super = &v
	//tree.Accept(&v)
}
