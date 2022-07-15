package jmespath

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/jmespath/parser"
	opt "github.com/wrmsr/bane/pkg/util/optional"
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
		node = Sequence{Items: []Node{node, Project{Child: v.chainedNode}}}
		v.chainedNode = nil
	}
	return node
}

func (v *parseVisitor) createSequenceIfChained(node Node) Node {
	if v.chainedNode != nil {
		node = Sequence{Items: []Node{node, v.chainedNode}}
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
	right := v.Visit(ctx.Expression(1)).(Node)
	left := v.Visit(ctx.Expression(0)).(Node)
	return Sequence{Items: []Node{left, right}}
}

func (v *parseVisitor) VisitParameterExpression(ctx *parser.ParameterExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitIdentifierExpression(ctx *parser.IdentifierExpressionContext) any {
	return v.Visit(ctx.Identifier())
}

func (v *parseVisitor) VisitNotExpression(ctx *parser.NotExpressionContext) any {
	return Negate{Item: v.Visit(ctx.Expression()).(Node)}
}

func (v *parseVisitor) VisitRawStringExpression(ctx *parser.RawStringExpressionContext) any {
	// FIXME: shared escaping with tree (core.util.StringEscaping) - really, in tok.util for java/sql?
	return String{Value: ctx.RAW_STRING().GetText()}
}

func (v *parseVisitor) VisitComparisonExpression(ctx *parser.ComparisonExpressionContext) any {
	op := ParseCmpOp(ctx.COMPARATOR().GetText())
	right := v.nonChainingVisit(ctx.Expression(1))
	left := v.nonChainingVisit(ctx.Expression(0))
	return Cmp{Op: op, Left: left, Right: right}

}

func (v *parseVisitor) VisitParenExpression(ctx *parser.ParenExpressionContext) any {
	return v.createSequenceIfChained(v.nonChainingVisit(ctx.Expression()))
}

func (v *parseVisitor) VisitBracketExpression(ctx *parser.BracketExpressionContext) any {
	result := v.Visit(ctx.BracketSpecifier())
	if result == nil {
		result = v.chainedNode
		v.chainedNode = nil
	}
	return result
}

func (v *parseVisitor) VisitOrExpression(ctx *parser.OrExpressionContext) any {
	left := v.nonChainingVisit(ctx.Expression(0))
	right := v.nonChainingVisit(ctx.Expression(1))
	return v.createSequenceIfChained(Or{Left: left, Right: right})
}

func (v *parseVisitor) VisitCurrentNodeExpression(ctx *parser.CurrentNodeExpressionContext) any {
	if v.chainedNode == nil {
		return Current{}
	}
	result := v.chainedNode
	v.chainedNode = nil
	return result
}

func (v *parseVisitor) VisitChainExpression(ctx *parser.ChainExpressionContext) any {
	v.chainedNode = maybeNode(v.Visit(ctx.ChainedExpression()))
	return v.createSequenceIfChained(maybeNode(v.Visit(ctx.Expression())))
}

func (v *parseVisitor) VisitAndExpression(ctx *parser.AndExpressionContext) any {
	left := v.nonChainingVisit(ctx.Expression(0))
	right := v.nonChainingVisit(ctx.Expression(1))
	return v.createSequenceIfChained(And{Left: left, Right: right})
}

func (v *parseVisitor) VisitMultiSelectHashExpression(ctx *parser.MultiSelectHashExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitWildcardExpression(ctx *parser.WildcardExpressionContext) any {
	return v.Visit(ctx.Wildcard())
}

func (v *parseVisitor) VisitFunctionCallExpression(ctx *parser.FunctionCallExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitMultiSelectListExpression(ctx *parser.MultiSelectListExpressionContext) any {
	var lst []Node
	for _, c := range ctx.ExpressionContext.GetChildren() {
		lst = append(lst, v.nonChainingVisit(c.(antlr.ParseTree)).(Node))
	}
	return v.createSequenceIfChained(CreateArray{Items: lst})
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
	return v.createProjectionIfChained(FlattenObject{})
}

func (v *parseVisitor) VisitBracketIndex(ctx *parser.BracketIndexContext) any {
	index := check.Must1(strconv.Atoi(ctx.SIGNED_INT().GetText()))
	v.chainedNode = v.createSequenceIfChained(Index{Value: index})
	return nil
}

func (v *parseVisitor) VisitBracketStar(ctx *parser.BracketStarContext) any {
	projection := v.chainedNode
	if projection == nil {
		projection = Current{}
	}
	v.chainedNode = Project{Child: projection}
	return nil
}

func (v *parseVisitor) VisitBracketSlice(ctx *parser.BracketSliceContext) any {
	var start, stop, step opt.Optional[int]
	sliceCtx := ctx.SliceNode().(*parser.SliceNodeContext)
	if sliceCtx.GetSliceStart() != nil {
		start = opt.Just(check.Must1(strconv.Atoi(sliceCtx.GetSliceStart().GetText())))
	}
	if sliceCtx.GetSliceStop() != nil {
		stop = opt.Just(check.Must1(strconv.Atoi(sliceCtx.GetSliceStop().GetText())))
	}
	if sliceCtx.GetSliceStep() != nil {
		step = opt.Just(check.Must1(strconv.Atoi(sliceCtx.GetSliceStep().GetText())))
		if step.Value() == 0 {
			panic("zero step")
		}
	}
	v.chainedNode = v.createProjectionIfChained(Slice{Start: start, Stop: stop, Step: step})
	return nil
}

func (v *parseVisitor) VisitBracketFlatten(ctx *parser.BracketFlattenContext) any {
	return v.createProjectionIfChained(FlattenArray{})
}

func (v *parseVisitor) VisitSelect(ctx *parser.SelectContext) any {
	v.chainedNode = v.createProjectionIfChained(Selection{Child: v.nonChainingVisit(ctx.Expression())})
	return nil
}

func (v *parseVisitor) VisitMultiSelectList(ctx *parser.MultiSelectListContext) any {
	var lst []Node
	for _, c := range ctx.AllExpression() {
		lst = append(lst, v.nonChainingVisit(c).(Node))
	}
	return v.createSequenceIfChained(CreateArray{Items: lst})
}

func (v *parseVisitor) VisitMultiSelectHash(ctx *parser.MultiSelectHashContext) any {
	dct := make(map[string]Node)
	kvs := ctx.AllKeyvalExpr()
	for _, kv := range kvs {
		// FIXME: unquote?
		kvCtx := kv.(*parser.KeyvalExprContext)
		key := kvCtx.Identifier().GetText()
		value := v.nonChainingVisit(kvCtx.Expression())
		dct[key] = value
	}
	return v.createSequenceIfChained(CreateObject{Fields: dct})
}

func (v *parseVisitor) VisitKeyvalExpr(ctx *parser.KeyvalExprContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitSliceNode(ctx *parser.SliceNodeContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitNameParameter(ctx *parser.NameParameterContext) any {
	return Parameter{Target: NameTarget(ctx.NAME().GetText())}
}

func (v *parseVisitor) VisitNumberParameter(ctx *parser.NumberParameterContext) any {
	return Parameter{Target: NumberTarget(check.Must1(strconv.Atoi(ctx.INT().GetText())))}
}

func (v *parseVisitor) VisitFunctionExpression(ctx *parser.FunctionExpressionContext) any {
	name := ctx.NAME().GetText()
	fas := ctx.AllFunctionArg()
	args := make([]Node, len(fas))
	for i, a := range fas {
		args[i] = v.nonChainingVisit(a).(Node)
	}
	return v.createSequenceIfChained(Call{Name: name, Args: args})
}

func (v *parseVisitor) VisitFunctionArg(ctx *parser.FunctionArgContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitCurrentNode(ctx *parser.CurrentNodeContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitExpressionType(ctx *parser.ExpressionTypeContext) any {
	expression := v.createSequenceIfChained(v.Visit(ctx.Expression()).(Node))
	return ExprRef{Expr: expression}
}

func (v *parseVisitor) VisitLiteral(ctx *parser.LiteralContext) any {
	// FIXME: unescape
	s := ctx.JsonValue().GetText()
	return JsonLiteral{Text: s}
}

func (v *parseVisitor) VisitIdentifier(ctx *parser.IdentifierContext) any {
	s := ctx.GetText()
	if s[0] == '"' {
		s = s[1 : len(s)-1]
	}
	return v.createSequenceIfChained(Property{Name: s})
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
