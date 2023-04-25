// Code generated from Jmespath.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Jmespath

import antlr "github.com/wrmsr/bane/core/antlr/runtime"

type BaseJmespathVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseJmespathVisitor) VisitSingleExpression(ctx *SingleExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitPipeExpression(ctx *PipeExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitParameterExpression(ctx *ParameterExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitIdentifierExpression(ctx *IdentifierExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitNotExpression(ctx *NotExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitRawStringExpression(ctx *RawStringExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitComparisonExpression(ctx *ComparisonExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitParenExpression(ctx *ParenExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitBracketExpression(ctx *BracketExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitOrExpression(ctx *OrExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitCurrentNodeExpression(ctx *CurrentNodeExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitChainExpression(ctx *ChainExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitAndExpression(ctx *AndExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitMultiSelectHashExpression(ctx *MultiSelectHashExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitWildcardExpression(ctx *WildcardExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitMultiSelectListExpression(ctx *MultiSelectListExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitBracketedExpression(ctx *BracketedExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitLiteralExpression(ctx *LiteralExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitChainedExpression(ctx *ChainedExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitWildcard(ctx *WildcardContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitBracketIndex(ctx *BracketIndexContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitBracketStar(ctx *BracketStarContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitBracketSlice(ctx *BracketSliceContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitBracketFlatten(ctx *BracketFlattenContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitSelect(ctx *SelectContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitMultiSelectList(ctx *MultiSelectListContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitMultiSelectHash(ctx *MultiSelectHashContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitKeyvalExpr(ctx *KeyvalExprContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitSliceNode(ctx *SliceNodeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitNameParameter(ctx *NameParameterContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitNumberParameter(ctx *NumberParameterContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitFunctionExpression(ctx *FunctionExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitFunctionArg(ctx *FunctionArgContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitCurrentNode(ctx *CurrentNodeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitExpressionType(ctx *ExpressionTypeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitLiteral(ctx *LiteralContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitIdentifier(ctx *IdentifierContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitJsonObject(ctx *JsonObjectContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitJsonObjectPair(ctx *JsonObjectPairContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitJsonArray(ctx *JsonArrayContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitJsonStringValue(ctx *JsonStringValueContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitJsonNumberValue(ctx *JsonNumberValueContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitJsonObjectValue(ctx *JsonObjectValueContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitJsonArrayValue(ctx *JsonArrayValueContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseJmespathVisitor) VisitJsonConstantValue(ctx *JsonConstantValueContext) any {
	return v.VisitChildren(ctx)
}
