package python

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/wrmsr/bane/exp/util/python/parser"
	antlru "github.com/wrmsr/bane/pkg/util/antlr"
	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/slices"
)

//

type parseVisitor struct{}

//

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
	return v.error(fmt.Errorf("%s", node))
}

func (v *parseVisitor) VisitErrorNode(node antlr.ErrorNode) any {
	return v.error(fmt.Errorf("%s", node))
}

//

func (v *parseVisitor) NodeVisit(tree antlr.ParseTree) Node {
	return v.Visit(tree).(Node)
}

func visitArithWithCont[L, R antlr.ParserRuleContext](v *parseVisitor, ctx antlr.ParserRuleContext) Node {
	l := v.NodeVisit(check.Single(antlru.FindTreeChildren[L](ctx)))
	for _, rctx := range antlru.FindTreeChildren[R](ctx) {
		a := v.Visit(rctx).(Arith)
		a.Left = l
		l = a
	}
	return l
}

func visitArithCont(v *parseVisitor, op antlr.Token, rctx antlr.ParserRuleContext) Node {
	return Arith{
		Op:    ParseArithOp(op.GetText()),
		Right: v.NodeVisit(rctx),
	}
}

func (v *parseVisitor) nodeVisitTestListComp(ctx parser.ITestListCompContext) []Node {
	if ctx == nil {
		return nil
	}
	return slices.Map(func(c parser.IExprOrStarExprContext) Node {
		return v.NodeVisit(c)
	}, ctx.(*parser.TestListCompContext).AllExprOrStarExpr())
}

func (v *parseVisitor) error(e error) any {
	panic(e)
}

func (v *parseVisitor) unimplemented(ctx antlr.ParserRuleContext) any {
	return v.error(fmt.Errorf("unimplemented: %s", ctx))
}

//

func (v *parseVisitor) VisitSingleExpr(ctx *parser.SingleExprContext) any {
	return v.Visit(ctx.Expr())
}

func (v *parseVisitor) VisitExpr(ctx *parser.ExprContext) any {
	return v.Visit(ctx.OrTest())
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
	items := make([]CmpItem, len(cos))
	for i, co := range cos {
		items[i] = CmpItem{
			Op:    ParseCmpOp(co.GetText()),
			Right: v.NodeVisit(ems[i+1]),
		}
	}
	return Cmp{
		Left:  left,
		Items: items,
	}
}

func (v *parseVisitor) VisitCompOp(ctx *parser.CompOpContext) any {
	return v.unimplemented(ctx)
}

func (v *parseVisitor) VisitExprMain(ctx *parser.ExprMainContext) any {
	return visitArithWithCont[parser.IXorExprContext, parser.IExprContContext](v, ctx)
}

func (v *parseVisitor) VisitExprCont(ctx *parser.ExprContContext) any {
	return visitArithCont(v, ctx.GetOp(), ctx.XorExpr())
}

func (v *parseVisitor) VisitXorExpr(ctx *parser.XorExprContext) any {
	return visitArithWithCont[parser.IAndExprContext, parser.IXorExprContContext](v, ctx)
}

func (v *parseVisitor) VisitXorExprCont(ctx *parser.XorExprContContext) any {
	return visitArithCont(v, ctx.GetOp(), ctx.AndExpr())
}

func (v *parseVisitor) VisitAndExpr(ctx *parser.AndExprContext) any {
	return visitArithWithCont[parser.IShiftExprContext, parser.IAndExprContContext](v, ctx)
}

func (v *parseVisitor) VisitAndExprCont(ctx *parser.AndExprContContext) any {
	return visitArithCont(v, ctx.GetOp(), ctx.ShiftExpr())
}

func (v *parseVisitor) VisitShiftExpr(ctx *parser.ShiftExprContext) any {
	return visitArithWithCont[parser.IArithExprContext, parser.IShiftExprContContext](v, ctx)
}

func (v *parseVisitor) VisitShiftExprCont(ctx *parser.ShiftExprContContext) any {
	return visitArithCont(v, ctx.GetOp(), ctx.ArithExpr())
}

func (v *parseVisitor) VisitArithExpr(ctx *parser.ArithExprContext) any {
	return visitArithWithCont[parser.ITermContext, parser.IArithExprContContext](v, ctx)
}

func (v *parseVisitor) VisitArithExprCont(ctx *parser.ArithExprContContext) any {
	return visitArithCont(v, ctx.GetOp(), ctx.Term())
}

func (v *parseVisitor) VisitTerm(ctx *parser.TermContext) any {
	return visitArithWithCont[parser.IFactorContext, parser.ITermContContext](v, ctx)
}

func (v *parseVisitor) VisitTermCont(ctx *parser.TermContContext) any {
	return visitArithCont(v, ctx.GetOp(), ctx.Factor())
}

func (v *parseVisitor) VisitFactor(ctx *parser.FactorContext) any {
	if p := ctx.Power(); p != nil {
		return v.Visit(p)
	}
	return Unary{
		Op:    ParseUnaryOp(ctx.GetOp().GetText()),
		Child: v.NodeVisit(ctx.Factor()),
	}
}

func (v *parseVisitor) VisitPower(ctx *parser.PowerContext) any {
	e := v.NodeVisit(ctx.AtomExpr())
	if ctx.Factor() == nil {
		return e
	}
	return Arith{
		Op:    ArithPow,
		Left:  e,
		Right: v.NodeVisit(ctx.Factor()),
	}
}

func (v *parseVisitor) VisitAtomExpr(ctx *parser.AtomExprContext) any {
	ret := v.NodeVisit(ctx.Atom())
	ts := ctx.AllTrailer()
	for _, t := range ts {
		c := v.NodeVisit(t)
		switch c := c.(type) {
		case Subscript:
			c.Child = ret
			ret = c
		case Attr:
			c.Child = ret
			ret = c
		default:
			v.error(fmt.Errorf("unhandled trailer: %v", c))
		}
	}
	return ret
}

func (v *parseVisitor) VisitParenAtom(ctx *parser.ParenAtomContext) any {
	return Parens{
		Children: v.nodeVisitTestListComp(ctx.TestListComp()),
	}
}

func (v *parseVisitor) VisitBracketAtom(ctx *parser.BracketAtomContext) any {
	return Brackets{
		Children: v.nodeVisitTestListComp(ctx.TestListComp()),
	}
}

func (v *parseVisitor) VisitDictOrSetAtom(ctx *parser.DictOrSetAtomContext) any {
	return v.Visit(ctx.DictOrSetMaker())
}

func (v *parseVisitor) VisitDictMaker(ctx *parser.DictMakerContext) any {
	return Dict{
		Items: slices.Map(func(c parser.IDictItemContext) DictItem {
			return v.Visit(c).(DictItem)
		}, ctx.AllDictItem()),
	}
}

func (v *parseVisitor) VisitSetMaker(ctx *parser.SetMakerContext) any {
	return Set{
		Children: slices.Map(func(c parser.IExprContext) Node {
			return v.NodeVisit(c)
		}, ctx.AllExpr()),
	}
}

func (v *parseVisitor) VisitKvDictItem(ctx *parser.KvDictItemContext) any {
	return DictItem{
		K: v.NodeVisit(ctx.GetK()),
		V: v.NodeVisit(ctx.GetV()),
	}
}

func (v *parseVisitor) VisitStarsDictItem(ctx *parser.StarsDictItemContext) any {
	return DictItem{
		IsStars: true,

		K: v.NodeVisit(ctx.ExprMain()),
	}
}

func (v *parseVisitor) VisitConstAtom(ctx *parser.ConstAtomContext) any {
	return v.Visit(ctx.Const())
}

func (v *parseVisitor) VisitConst(ctx *parser.ConstContext) any {
	if ss := ctx.AllSTRING(); len(ss) > 0 {
		return String{S: slices.Map(antlr.TerminalNode.GetText, ss)}
	}
	if n := ctx.NUMBER(); n != nil {
		return Number{S: n.GetText()}
	}
	if na := ctx.NAME(); na != nil {
		return Name{S: na.GetText()}
	}
	txt := ctx.GetText()
	switch txt {
	case "True":
		return True{}
	case "False":
		return False{}
	case "None":
		return None{}
	case "...":
		return Ellipsis{}
	}
	return v.error(fmt.Errorf("invalid const value: %s", txt))
}

func (v *parseVisitor) VisitTestListComp(ctx *parser.TestListCompContext) any {
	return v.unimplemented(ctx)
}

func (v *parseVisitor) VisitExprOrStarExpr(ctx *parser.ExprOrStarExprContext) interface{} {
	if se := ctx.StarExpr(); se != nil {
		return v.Visit(se)
	}
	return v.Visit(ctx.Expr())
}

func (v *parseVisitor) VisitStarExpr(ctx *parser.StarExprContext) interface{} {
	return Star{
		Child: v.NodeVisit(ctx.Expr()),
	}
}

func (v *parseVisitor) VisitTrailer(ctx *parser.TrailerContext) any {
	if sl := ctx.SubscriptList(); sl != nil {
		return v.Visit(sl)
	}
	return Attr{
		Name: ctx.NAME().GetText(),
	}
}

func (v *parseVisitor) VisitSubscriptList(ctx *parser.SubscriptListContext) any {
	cs := ctx.AllSubscript()
	ret := Subscript{
		Items: make([]Node, len(cs)),
	}
	for i, c := range cs {
		ret.Items[i] = v.NodeVisit(c)
	}
	return ret
}

func (v *parseVisitor) VisitExprSubscript(ctx *parser.ExprSubscriptContext) any {
	return v.Visit(ctx.Expr())
}

func (v *parseVisitor) VisitSliceSubscript(ctx *parser.SliceSubscriptContext) any {
	r := Slice{}
	if ctx.GetSta() != nil {
		r.Start = v.NodeVisit(ctx.GetSta())
	}
	if ctx.GetSto() != nil {
		r.Stop = v.NodeVisit(ctx.GetSto())
	}
	if ctx.GetSte() != nil {
		ste := ctx.GetSte().(*parser.SliceOpContext)
		if stex := ste.Expr(); stex != nil {
			r.Step = v.NodeVisit(stex)
		}
	}
	return r
}

func (v *parseVisitor) VisitSliceOp(ctx *parser.SliceOpContext) any {
	return v.unimplemented(ctx)
}
