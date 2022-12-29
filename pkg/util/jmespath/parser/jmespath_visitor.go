// Code generated from Jmespath.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Jmespath

import antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

// A complete Visitor for a parse tree produced by JmespathParser.
type JmespathVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by JmespathParser#singleExpression.
	VisitSingleExpression(ctx *SingleExpressionContext) interface{}

	// Visit a parse tree produced by JmespathParser#pipeExpression.
	VisitPipeExpression(ctx *PipeExpressionContext) interface{}

	// Visit a parse tree produced by JmespathParser#parameterExpression.
	VisitParameterExpression(ctx *ParameterExpressionContext) interface{}

	// Visit a parse tree produced by JmespathParser#identifierExpression.
	VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{}

	// Visit a parse tree produced by JmespathParser#notExpression.
	VisitNotExpression(ctx *NotExpressionContext) interface{}

	// Visit a parse tree produced by JmespathParser#rawStringExpression.
	VisitRawStringExpression(ctx *RawStringExpressionContext) interface{}

	// Visit a parse tree produced by JmespathParser#comparisonExpression.
	VisitComparisonExpression(ctx *ComparisonExpressionContext) interface{}

	// Visit a parse tree produced by JmespathParser#parenExpression.
	VisitParenExpression(ctx *ParenExpressionContext) interface{}

	// Visit a parse tree produced by JmespathParser#bracketExpression.
	VisitBracketExpression(ctx *BracketExpressionContext) interface{}

	// Visit a parse tree produced by JmespathParser#orExpression.
	VisitOrExpression(ctx *OrExpressionContext) interface{}

	// Visit a parse tree produced by JmespathParser#currentNodeExpression.
	VisitCurrentNodeExpression(ctx *CurrentNodeExpressionContext) interface{}

	// Visit a parse tree produced by JmespathParser#chainExpression.
	VisitChainExpression(ctx *ChainExpressionContext) interface{}

	// Visit a parse tree produced by JmespathParser#andExpression.
	VisitAndExpression(ctx *AndExpressionContext) interface{}

	// Visit a parse tree produced by JmespathParser#multiSelectHashExpression.
	VisitMultiSelectHashExpression(ctx *MultiSelectHashExpressionContext) interface{}

	// Visit a parse tree produced by JmespathParser#wildcardExpression.
	VisitWildcardExpression(ctx *WildcardExpressionContext) interface{}

	// Visit a parse tree produced by JmespathParser#functionCallExpression.
	VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{}

	// Visit a parse tree produced by JmespathParser#multiSelectListExpression.
	VisitMultiSelectListExpression(ctx *MultiSelectListExpressionContext) interface{}

	// Visit a parse tree produced by JmespathParser#bracketedExpression.
	VisitBracketedExpression(ctx *BracketedExpressionContext) interface{}

	// Visit a parse tree produced by JmespathParser#literalExpression.
	VisitLiteralExpression(ctx *LiteralExpressionContext) interface{}

	// Visit a parse tree produced by JmespathParser#chainedExpression.
	VisitChainedExpression(ctx *ChainedExpressionContext) interface{}

	// Visit a parse tree produced by JmespathParser#wildcard.
	VisitWildcard(ctx *WildcardContext) interface{}

	// Visit a parse tree produced by JmespathParser#bracketIndex.
	VisitBracketIndex(ctx *BracketIndexContext) interface{}

	// Visit a parse tree produced by JmespathParser#bracketStar.
	VisitBracketStar(ctx *BracketStarContext) interface{}

	// Visit a parse tree produced by JmespathParser#bracketSlice.
	VisitBracketSlice(ctx *BracketSliceContext) interface{}

	// Visit a parse tree produced by JmespathParser#bracketFlatten.
	VisitBracketFlatten(ctx *BracketFlattenContext) interface{}

	// Visit a parse tree produced by JmespathParser#select.
	VisitSelect(ctx *SelectContext) interface{}

	// Visit a parse tree produced by JmespathParser#multiSelectList.
	VisitMultiSelectList(ctx *MultiSelectListContext) interface{}

	// Visit a parse tree produced by JmespathParser#multiSelectHash.
	VisitMultiSelectHash(ctx *MultiSelectHashContext) interface{}

	// Visit a parse tree produced by JmespathParser#keyvalExpr.
	VisitKeyvalExpr(ctx *KeyvalExprContext) interface{}

	// Visit a parse tree produced by JmespathParser#sliceNode.
	VisitSliceNode(ctx *SliceNodeContext) interface{}

	// Visit a parse tree produced by JmespathParser#nameParameter.
	VisitNameParameter(ctx *NameParameterContext) interface{}

	// Visit a parse tree produced by JmespathParser#numberParameter.
	VisitNumberParameter(ctx *NumberParameterContext) interface{}

	// Visit a parse tree produced by JmespathParser#functionExpression.
	VisitFunctionExpression(ctx *FunctionExpressionContext) interface{}

	// Visit a parse tree produced by JmespathParser#functionArg.
	VisitFunctionArg(ctx *FunctionArgContext) interface{}

	// Visit a parse tree produced by JmespathParser#currentNode.
	VisitCurrentNode(ctx *CurrentNodeContext) interface{}

	// Visit a parse tree produced by JmespathParser#expressionType.
	VisitExpressionType(ctx *ExpressionTypeContext) interface{}

	// Visit a parse tree produced by JmespathParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by JmespathParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by JmespathParser#jsonObject.
	VisitJsonObject(ctx *JsonObjectContext) interface{}

	// Visit a parse tree produced by JmespathParser#jsonObjectPair.
	VisitJsonObjectPair(ctx *JsonObjectPairContext) interface{}

	// Visit a parse tree produced by JmespathParser#jsonArray.
	VisitJsonArray(ctx *JsonArrayContext) interface{}

	// Visit a parse tree produced by JmespathParser#jsonStringValue.
	VisitJsonStringValue(ctx *JsonStringValueContext) interface{}

	// Visit a parse tree produced by JmespathParser#jsonNumberValue.
	VisitJsonNumberValue(ctx *JsonNumberValueContext) interface{}

	// Visit a parse tree produced by JmespathParser#jsonObjectValue.
	VisitJsonObjectValue(ctx *JsonObjectValueContext) interface{}

	// Visit a parse tree produced by JmespathParser#jsonArrayValue.
	VisitJsonArrayValue(ctx *JsonArrayValueContext) interface{}

	// Visit a parse tree produced by JmespathParser#jsonConstantValue.
	VisitJsonConstantValue(ctx *JsonConstantValueContext) interface{}
}
