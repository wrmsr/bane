package jmespath

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/wrmsr/bane/pkg/util/jmespath/parser"
)

//

type parseVisitor struct {
	chainedNode Node
}

func maybeNode(o any) Node {
	if o == nil {
		return nil
	}
	return o.(Node)
}

func (v *parseVisitor) createProjectionIfChained(node Node) Node {
	if v.chainedNode != nil {
		node = Sequence{items: []Node{node, Project{child: v.chainedNode}}}
		v.chainedNode = nil
	}
	return node
}

func (v *parseVisitor) createSequenceIfChained(node Node) Node {
	if v.chainedNode != nil {
		node = Sequence{items: []Node{node, v.chainedNode}}
		v.chainedNode = nil
	}
	return node
}

func (v *parseVisitor) nonChainingVisit(tree antlr.ParseTree) Node {
	stashedNextNode := v.chainedNode
	v.chainedNode = nil
	result := v.createSequenceIfChained(v.Visit(tree).(Node))
	v.chainedNode = stashedNextNode
	return result
}

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

//

var _ parser.JmespathVisitor = &parseVisitor{}

func (v *parseVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *parseVisitor) VisitChildren(node antlr.RuleNode) any {
	var result any = v.defaultResult()
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
	return v.defaultResult()
}

func (v *parseVisitor) VisitErrorNode(node antlr.ErrorNode) any {
	return v.defaultResult()
}

//

func (v *parseVisitor) VisitSingleExpression(ctx *parser.SingleExpressionContext) any {
	return v.Visit(ctx.Expression())
}

func (v *parseVisitor) VisitPipeExpression(ctx *parser.PipeExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitParameterExpression(ctx *parser.ParameterExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitIdentifierExpression(ctx *parser.IdentifierExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitNotExpression(ctx *parser.NotExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitRawStringExpression(ctx *parser.RawStringExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitComparisonExpression(ctx *parser.ComparisonExpressionContext) any {
	o := ParseCmpOp(ctx.COMPARATOR().GetText())
	right := v.nonChainingVisit(ctx.Expression(1))
	left := v.nonChainingVisit(ctx.Expression(0))
	return Cmp{o: o, left: left, right: right}

}

func (v *parseVisitor) VisitParenExpression(ctx *parser.ParenExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitBracketExpression(ctx *parser.BracketExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitOrExpression(ctx *parser.OrExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitCurrentNodeExpression(ctx *parser.CurrentNodeExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitChainExpression(ctx *parser.ChainExpressionContext) any {
	v.chainedNode = maybeNode(v.Visit(ctx.ChainedExpression()))
	return v.createSequenceIfChained(maybeNode(v.Visit(ctx.Expression())))
}

func (v *parseVisitor) VisitAndExpression(ctx *parser.AndExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitMultiSelectHashExpression(ctx *parser.MultiSelectHashExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitWildcardExpression(ctx *parser.WildcardExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitFunctionCallExpression(ctx *parser.FunctionCallExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitMultiSelectListExpression(ctx *parser.MultiSelectListExpressionContext) any {
	//var lst []Node
	//for _, c := range ctx.ExpressionContext.GetChildren() {
	//	lst = append(lst, v.nonChainingVisit(c.(antlr.ParseTree)).(Node))
	//}
	//return v.createSequenceIfChained(CreateArray{items: lst})
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitBracketedExpression(ctx *parser.BracketedExpressionContext) any {
	chainAfterExpression := maybeNode(v.Visit(ctx.BracketSpecifier()))
	expression := v.createSequenceIfChained(v.Visit(ctx.Expression()).(Node))
	v.chainedNode = chainAfterExpression
	return v.createSequenceIfChained(expression)
}

func (v *parseVisitor) VisitLiteralExpression(ctx *parser.LiteralExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitChainedExpression(ctx *parser.ChainedExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitWildcard(ctx *parser.WildcardContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitBracketIndex(ctx *parser.BracketIndexContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitBracketStar(ctx *parser.BracketStarContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitBracketSlice(ctx *parser.BracketSliceContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitBracketFlatten(ctx *parser.BracketFlattenContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitSelect(ctx *parser.SelectContext) any {
	v.chainedNode = v.createProjectionIfChained(Selection{child: v.nonChainingVisit(ctx.Expression())})
	return nil
}

func (v *parseVisitor) VisitMultiSelectList(ctx *parser.MultiSelectListContext) any {
	var lst []Node
	for _, c := range ctx.AllExpression() {
		lst = append(lst, v.nonChainingVisit(c).(Node))
	}
	return v.createSequenceIfChained(CreateArray{items: lst})
}

func (v *parseVisitor) VisitMultiSelectHash(ctx *parser.MultiSelectHashContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitKeyvalExpr(ctx *parser.KeyvalExprContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitSliceNode(ctx *parser.SliceNodeContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitNameParameter(ctx *parser.NameParameterContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitNumberParameter(ctx *parser.NumberParameterContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitFunctionExpression(ctx *parser.FunctionExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitFunctionArg(ctx *parser.FunctionArgContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitCurrentNode(ctx *parser.CurrentNodeContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitExpressionType(ctx *parser.ExpressionTypeContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitLiteral(ctx *parser.LiteralContext) any {
	// FIXME: unescape
	s := ctx.JsonValue().GetText()
	return JsonLiteral{text: s}
}

func (v *parseVisitor) VisitIdentifier(ctx *parser.IdentifierContext) any {
	// FIXME: unquote
	return v.createSequenceIfChained(Property{name: ctx.GetText()})
}

func (v *parseVisitor) VisitJsonObject(ctx *parser.JsonObjectContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitJsonObjectPair(ctx *parser.JsonObjectPairContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitJsonArray(ctx *parser.JsonArrayContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitJsonStringValue(ctx *parser.JsonStringValueContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitJsonNumberValue(ctx *parser.JsonNumberValueContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitJsonObjectValue(ctx *parser.JsonObjectValueContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitJsonArrayValue(ctx *parser.JsonArrayValueContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitJsonConstantValue(ctx *parser.JsonConstantValueContext) any {
	return v.VisitChildren(ctx)
}
