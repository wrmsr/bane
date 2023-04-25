// Code generated from Jmespath.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Jmespath

import antlr "github.com/wrmsr/bane/core/antlr/runtime"

// A complete Visitor for a parse tree produced by JmespathParser.
type JmespathVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by JmespathParser#singleExpression.
	VisitSingleExpression(ctx *SingleExpressionContext) any

	// Visit a parse tree produced by JmespathParser#pipeExpression.
	VisitPipeExpression(ctx *PipeExpressionContext) any

	// Visit a parse tree produced by JmespathParser#parameterExpression.
	VisitParameterExpression(ctx *ParameterExpressionContext) any

	// Visit a parse tree produced by JmespathParser#identifierExpression.
	VisitIdentifierExpression(ctx *IdentifierExpressionContext) any

	// Visit a parse tree produced by JmespathParser#notExpression.
	VisitNotExpression(ctx *NotExpressionContext) any

	// Visit a parse tree produced by JmespathParser#rawStringExpression.
	VisitRawStringExpression(ctx *RawStringExpressionContext) any

	// Visit a parse tree produced by JmespathParser#comparisonExpression.
	VisitComparisonExpression(ctx *ComparisonExpressionContext) any

	// Visit a parse tree produced by JmespathParser#parenExpression.
	VisitParenExpression(ctx *ParenExpressionContext) any

	// Visit a parse tree produced by JmespathParser#bracketExpression.
	VisitBracketExpression(ctx *BracketExpressionContext) any

	// Visit a parse tree produced by JmespathParser#orExpression.
	VisitOrExpression(ctx *OrExpressionContext) any

	// Visit a parse tree produced by JmespathParser#currentNodeExpression.
	VisitCurrentNodeExpression(ctx *CurrentNodeExpressionContext) any

	// Visit a parse tree produced by JmespathParser#chainExpression.
	VisitChainExpression(ctx *ChainExpressionContext) any

	// Visit a parse tree produced by JmespathParser#andExpression.
	VisitAndExpression(ctx *AndExpressionContext) any

	// Visit a parse tree produced by JmespathParser#multiSelectHashExpression.
	VisitMultiSelectHashExpression(ctx *MultiSelectHashExpressionContext) any

	// Visit a parse tree produced by JmespathParser#wildcardExpression.
	VisitWildcardExpression(ctx *WildcardExpressionContext) any

	// Visit a parse tree produced by JmespathParser#functionCallExpression.
	VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) any

	// Visit a parse tree produced by JmespathParser#multiSelectListExpression.
	VisitMultiSelectListExpression(ctx *MultiSelectListExpressionContext) any

	// Visit a parse tree produced by JmespathParser#bracketedExpression.
	VisitBracketedExpression(ctx *BracketedExpressionContext) any

	// Visit a parse tree produced by JmespathParser#literalExpression.
	VisitLiteralExpression(ctx *LiteralExpressionContext) any

	// Visit a parse tree produced by JmespathParser#chainedExpression.
	VisitChainedExpression(ctx *ChainedExpressionContext) any

	// Visit a parse tree produced by JmespathParser#wildcard.
	VisitWildcard(ctx *WildcardContext) any

	// Visit a parse tree produced by JmespathParser#bracketIndex.
	VisitBracketIndex(ctx *BracketIndexContext) any

	// Visit a parse tree produced by JmespathParser#bracketStar.
	VisitBracketStar(ctx *BracketStarContext) any

	// Visit a parse tree produced by JmespathParser#bracketSlice.
	VisitBracketSlice(ctx *BracketSliceContext) any

	// Visit a parse tree produced by JmespathParser#bracketFlatten.
	VisitBracketFlatten(ctx *BracketFlattenContext) any

	// Visit a parse tree produced by JmespathParser#select.
	VisitSelect(ctx *SelectContext) any

	// Visit a parse tree produced by JmespathParser#multiSelectList.
	VisitMultiSelectList(ctx *MultiSelectListContext) any

	// Visit a parse tree produced by JmespathParser#multiSelectHash.
	VisitMultiSelectHash(ctx *MultiSelectHashContext) any

	// Visit a parse tree produced by JmespathParser#keyvalExpr.
	VisitKeyvalExpr(ctx *KeyvalExprContext) any

	// Visit a parse tree produced by JmespathParser#sliceNode.
	VisitSliceNode(ctx *SliceNodeContext) any

	// Visit a parse tree produced by JmespathParser#nameParameter.
	VisitNameParameter(ctx *NameParameterContext) any

	// Visit a parse tree produced by JmespathParser#numberParameter.
	VisitNumberParameter(ctx *NumberParameterContext) any

	// Visit a parse tree produced by JmespathParser#functionExpression.
	VisitFunctionExpression(ctx *FunctionExpressionContext) any

	// Visit a parse tree produced by JmespathParser#functionArg.
	VisitFunctionArg(ctx *FunctionArgContext) any

	// Visit a parse tree produced by JmespathParser#currentNode.
	VisitCurrentNode(ctx *CurrentNodeContext) any

	// Visit a parse tree produced by JmespathParser#expressionType.
	VisitExpressionType(ctx *ExpressionTypeContext) any

	// Visit a parse tree produced by JmespathParser#literal.
	VisitLiteral(ctx *LiteralContext) any

	// Visit a parse tree produced by JmespathParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) any

	// Visit a parse tree produced by JmespathParser#jsonObject.
	VisitJsonObject(ctx *JsonObjectContext) any

	// Visit a parse tree produced by JmespathParser#jsonObjectPair.
	VisitJsonObjectPair(ctx *JsonObjectPairContext) any

	// Visit a parse tree produced by JmespathParser#jsonArray.
	VisitJsonArray(ctx *JsonArrayContext) any

	// Visit a parse tree produced by JmespathParser#jsonStringValue.
	VisitJsonStringValue(ctx *JsonStringValueContext) any

	// Visit a parse tree produced by JmespathParser#jsonNumberValue.
	VisitJsonNumberValue(ctx *JsonNumberValueContext) any

	// Visit a parse tree produced by JmespathParser#jsonObjectValue.
	VisitJsonObjectValue(ctx *JsonObjectValueContext) any

	// Visit a parse tree produced by JmespathParser#jsonArrayValue.
	VisitJsonArrayValue(ctx *JsonArrayValueContext) any

	// Visit a parse tree produced by JmespathParser#jsonConstantValue.
	VisitJsonConstantValue(ctx *JsonConstantValueContext) any
}
