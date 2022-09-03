package python

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/wrmsr/bane/exp/util/python/parser"
	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/slices"
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

func (v *parseVisitor) NodeVisit(tree antlr.ParseTree) Node {
	return v.Visit(tree).(Node)
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
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitOrTest(ctx *parser.OrTestContext) any {
	cs := check.NotEmptySlice(ctx.AllAndTest())
	if len(cs) > 1 {
		return Or{
			Children: slices.Map(func(c parser.IAndTestContext) Node {
				return v.NodeVisit(c)
			}, cs),
		}
	}
	return v.Visit(check.Single(cs))
}

func (v *parseVisitor) VisitAndTest(ctx *parser.AndTestContext) any {
	cs := check.NotEmptySlice(ctx.AllNotTest())
	if len(cs) > 1 {
		return And{
			Children: slices.Map(func(c parser.INotTestContext) Node {
				return v.NodeVisit(c)
			}, cs),
		}
	}
	return v.Visit(check.Single(cs))
}

func (v *parseVisitor) VisitNotTest(ctx *parser.NotTestContext) any {
	if nt := ctx.NotTest(); nt != nil {
		return Not{
			Child: v.NodeVisit(nt),
		}
	}
	return v.Visit(ctx.Comparison())
}

func (v *parseVisitor) VisitComparison(ctx *parser.ComparisonContext) any {
	ems := ctx.AllExprMain()
	cos := ctx.AllCompOp()
	check.Equal(len(ems), len(cos)+1)
	left := v.NodeVisit(ems[0])
	if len(cos) < 1 {
		return left
	}
	ops := make([]CmpOp, len(cos))
	rights := make([]Node, len(cos))
	for i, co := range cos {
		ops[i] = ParseCmpOp(co.GetText())
		rights[i] = v.NodeVisit(ems[i+1])
	}
	return Comparison{
		Left:   left,
		Ops:    ops,
		Rights: rights,
	}
}

func (v *parseVisitor) VisitCompOp(ctx *parser.CompOpContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitExprMain(ctx *parser.ExprMainContext) any {
	cs := ctx.AllExprCont()
	check.EmptySlice(cs)

	return v.Visit(ctx.XorExpr())
}

func (v *parseVisitor) VisitExprCont(ctx *parser.ExprContContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitXorExpr(ctx *parser.XorExprContext) any {
	cs := ctx.AllXorExprCont()
	check.EmptySlice(cs)

	return v.Visit(ctx.AndExpr())
}

func (v *parseVisitor) VisitXorExprCont(ctx *parser.XorExprContContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitAndExpr(ctx *parser.AndExprContext) any {
	cs := ctx.AllAndExprCont()
	check.EmptySlice(cs)

	return v.Visit(ctx.ShiftExpr())
}

func (v *parseVisitor) VisitAndExprCont(ctx *parser.AndExprContContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitShiftExpr(ctx *parser.ShiftExprContext) any {
	cs := ctx.AllShiftExprCont()
	check.EmptySlice(cs)

	return v.Visit(ctx.ArithExpr())
}

func (v *parseVisitor) VisitShiftExprCont(ctx *parser.ShiftExprContContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitArithExpr(ctx *parser.ArithExprContext) any {
	cs := ctx.AllArithExprCont()
	check.EmptySlice(cs)

	return v.Visit(ctx.Term())
}

func (v *parseVisitor) VisitArithExprCont(ctx *parser.ArithExprContContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitTerm(ctx *parser.TermContext) any {
	cs := ctx.AllTermCont()
	check.EmptySlice(cs)

	return v.Visit(ctx.Factor())
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
