// Code generated from Jmespath.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Jmespath

import antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

type BaseJmespathVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseJmespathVisitor) VisitSingleExpression(ctx *SingleExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitPipeExpression(ctx *PipeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitParameterExpression(ctx *ParameterExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitNotExpression(ctx *NotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitRawStringExpression(ctx *RawStringExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitComparisonExpression(ctx *ComparisonExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitParenExpression(ctx *ParenExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitBracketExpression(ctx *BracketExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitOrExpression(ctx *OrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitCurrentNodeExpression(ctx *CurrentNodeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitChainExpression(ctx *ChainExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitAndExpression(ctx *AndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitMultiSelectHashExpression(ctx *MultiSelectHashExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitWildcardExpression(ctx *WildcardExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitMultiSelectListExpression(ctx *MultiSelectListExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitBracketedExpression(ctx *BracketedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitLiteralExpression(ctx *LiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitChainedExpression(ctx *ChainedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitWildcard(ctx *WildcardContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitBracketIndex(ctx *BracketIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitBracketStar(ctx *BracketStarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitBracketSlice(ctx *BracketSliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitBracketFlatten(ctx *BracketFlattenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitSelect(ctx *SelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitMultiSelectList(ctx *MultiSelectListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitMultiSelectHash(ctx *MultiSelectHashContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitKeyvalExpr(ctx *KeyvalExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitSliceNode(ctx *SliceNodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitNameParameter(ctx *NameParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitNumberParameter(ctx *NumberParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitFunctionExpression(ctx *FunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitFunctionArg(ctx *FunctionArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitCurrentNode(ctx *CurrentNodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitExpressionType(ctx *ExpressionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitJsonObject(ctx *JsonObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitJsonObjectPair(ctx *JsonObjectPairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitJsonArray(ctx *JsonArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitJsonStringValue(ctx *JsonStringValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitJsonNumberValue(ctx *JsonNumberValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitJsonObjectValue(ctx *JsonObjectValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitJsonArrayValue(ctx *JsonArrayValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitJsonConstantValue(ctx *JsonConstantValueContext) interface{} {
	return v.VisitChildren(ctx)
}
