// Code generated from Jmespath.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Jmespath

import antlr "github.com/wrmsr/bane/core/antlr/runtime"

// JmespathListener is a complete listener for a parse tree produced by JmespathParser.
type JmespathListener interface {
	antlr.ParseTreeListener

	// EnterSingleExpression is called when entering the singleExpression production.
	EnterSingleExpression(c *SingleExpressionContext)

	// EnterPipeExpression is called when entering the pipeExpression production.
	EnterPipeExpression(c *PipeExpressionContext)

	// EnterParameterExpression is called when entering the parameterExpression production.
	EnterParameterExpression(c *ParameterExpressionContext)

	// EnterIdentifierExpression is called when entering the identifierExpression production.
	EnterIdentifierExpression(c *IdentifierExpressionContext)

	// EnterNotExpression is called when entering the notExpression production.
	EnterNotExpression(c *NotExpressionContext)

	// EnterRawStringExpression is called when entering the rawStringExpression production.
	EnterRawStringExpression(c *RawStringExpressionContext)

	// EnterComparisonExpression is called when entering the comparisonExpression production.
	EnterComparisonExpression(c *ComparisonExpressionContext)

	// EnterParenExpression is called when entering the parenExpression production.
	EnterParenExpression(c *ParenExpressionContext)

	// EnterBracketExpression is called when entering the bracketExpression production.
	EnterBracketExpression(c *BracketExpressionContext)

	// EnterOrExpression is called when entering the orExpression production.
	EnterOrExpression(c *OrExpressionContext)

	// EnterCurrentNodeExpression is called when entering the currentNodeExpression production.
	EnterCurrentNodeExpression(c *CurrentNodeExpressionContext)

	// EnterChainExpression is called when entering the chainExpression production.
	EnterChainExpression(c *ChainExpressionContext)

	// EnterAndExpression is called when entering the andExpression production.
	EnterAndExpression(c *AndExpressionContext)

	// EnterMultiSelectHashExpression is called when entering the multiSelectHashExpression production.
	EnterMultiSelectHashExpression(c *MultiSelectHashExpressionContext)

	// EnterWildcardExpression is called when entering the wildcardExpression production.
	EnterWildcardExpression(c *WildcardExpressionContext)

	// EnterFunctionCallExpression is called when entering the functionCallExpression production.
	EnterFunctionCallExpression(c *FunctionCallExpressionContext)

	// EnterMultiSelectListExpression is called when entering the multiSelectListExpression production.
	EnterMultiSelectListExpression(c *MultiSelectListExpressionContext)

	// EnterBracketedExpression is called when entering the bracketedExpression production.
	EnterBracketedExpression(c *BracketedExpressionContext)

	// EnterLiteralExpression is called when entering the literalExpression production.
	EnterLiteralExpression(c *LiteralExpressionContext)

	// EnterChainedExpression is called when entering the chainedExpression production.
	EnterChainedExpression(c *ChainedExpressionContext)

	// EnterWildcard is called when entering the wildcard production.
	EnterWildcard(c *WildcardContext)

	// EnterBracketIndex is called when entering the bracketIndex production.
	EnterBracketIndex(c *BracketIndexContext)

	// EnterBracketStar is called when entering the bracketStar production.
	EnterBracketStar(c *BracketStarContext)

	// EnterBracketSlice is called when entering the bracketSlice production.
	EnterBracketSlice(c *BracketSliceContext)

	// EnterBracketFlatten is called when entering the bracketFlatten production.
	EnterBracketFlatten(c *BracketFlattenContext)

	// EnterSelect is called when entering the select production.
	EnterSelect(c *SelectContext)

	// EnterMultiSelectList is called when entering the multiSelectList production.
	EnterMultiSelectList(c *MultiSelectListContext)

	// EnterMultiSelectHash is called when entering the multiSelectHash production.
	EnterMultiSelectHash(c *MultiSelectHashContext)

	// EnterKeyvalExpr is called when entering the keyvalExpr production.
	EnterKeyvalExpr(c *KeyvalExprContext)

	// EnterSliceNode is called when entering the sliceNode production.
	EnterSliceNode(c *SliceNodeContext)

	// EnterNameParameter is called when entering the nameParameter production.
	EnterNameParameter(c *NameParameterContext)

	// EnterNumberParameter is called when entering the numberParameter production.
	EnterNumberParameter(c *NumberParameterContext)

	// EnterFunctionExpression is called when entering the functionExpression production.
	EnterFunctionExpression(c *FunctionExpressionContext)

	// EnterFunctionArg is called when entering the functionArg production.
	EnterFunctionArg(c *FunctionArgContext)

	// EnterCurrentNode is called when entering the currentNode production.
	EnterCurrentNode(c *CurrentNodeContext)

	// EnterExpressionType is called when entering the expressionType production.
	EnterExpressionType(c *ExpressionTypeContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterJsonObject is called when entering the jsonObject production.
	EnterJsonObject(c *JsonObjectContext)

	// EnterJsonObjectPair is called when entering the jsonObjectPair production.
	EnterJsonObjectPair(c *JsonObjectPairContext)

	// EnterJsonArray is called when entering the jsonArray production.
	EnterJsonArray(c *JsonArrayContext)

	// EnterJsonStringValue is called when entering the jsonStringValue production.
	EnterJsonStringValue(c *JsonStringValueContext)

	// EnterJsonNumberValue is called when entering the jsonNumberValue production.
	EnterJsonNumberValue(c *JsonNumberValueContext)

	// EnterJsonObjectValue is called when entering the jsonObjectValue production.
	EnterJsonObjectValue(c *JsonObjectValueContext)

	// EnterJsonArrayValue is called when entering the jsonArrayValue production.
	EnterJsonArrayValue(c *JsonArrayValueContext)

	// EnterJsonConstantValue is called when entering the jsonConstantValue production.
	EnterJsonConstantValue(c *JsonConstantValueContext)

	// ExitSingleExpression is called when exiting the singleExpression production.
	ExitSingleExpression(c *SingleExpressionContext)

	// ExitPipeExpression is called when exiting the pipeExpression production.
	ExitPipeExpression(c *PipeExpressionContext)

	// ExitParameterExpression is called when exiting the parameterExpression production.
	ExitParameterExpression(c *ParameterExpressionContext)

	// ExitIdentifierExpression is called when exiting the identifierExpression production.
	ExitIdentifierExpression(c *IdentifierExpressionContext)

	// ExitNotExpression is called when exiting the notExpression production.
	ExitNotExpression(c *NotExpressionContext)

	// ExitRawStringExpression is called when exiting the rawStringExpression production.
	ExitRawStringExpression(c *RawStringExpressionContext)

	// ExitComparisonExpression is called when exiting the comparisonExpression production.
	ExitComparisonExpression(c *ComparisonExpressionContext)

	// ExitParenExpression is called when exiting the parenExpression production.
	ExitParenExpression(c *ParenExpressionContext)

	// ExitBracketExpression is called when exiting the bracketExpression production.
	ExitBracketExpression(c *BracketExpressionContext)

	// ExitOrExpression is called when exiting the orExpression production.
	ExitOrExpression(c *OrExpressionContext)

	// ExitCurrentNodeExpression is called when exiting the currentNodeExpression production.
	ExitCurrentNodeExpression(c *CurrentNodeExpressionContext)

	// ExitChainExpression is called when exiting the chainExpression production.
	ExitChainExpression(c *ChainExpressionContext)

	// ExitAndExpression is called when exiting the andExpression production.
	ExitAndExpression(c *AndExpressionContext)

	// ExitMultiSelectHashExpression is called when exiting the multiSelectHashExpression production.
	ExitMultiSelectHashExpression(c *MultiSelectHashExpressionContext)

	// ExitWildcardExpression is called when exiting the wildcardExpression production.
	ExitWildcardExpression(c *WildcardExpressionContext)

	// ExitFunctionCallExpression is called when exiting the functionCallExpression production.
	ExitFunctionCallExpression(c *FunctionCallExpressionContext)

	// ExitMultiSelectListExpression is called when exiting the multiSelectListExpression production.
	ExitMultiSelectListExpression(c *MultiSelectListExpressionContext)

	// ExitBracketedExpression is called when exiting the bracketedExpression production.
	ExitBracketedExpression(c *BracketedExpressionContext)

	// ExitLiteralExpression is called when exiting the literalExpression production.
	ExitLiteralExpression(c *LiteralExpressionContext)

	// ExitChainedExpression is called when exiting the chainedExpression production.
	ExitChainedExpression(c *ChainedExpressionContext)

	// ExitWildcard is called when exiting the wildcard production.
	ExitWildcard(c *WildcardContext)

	// ExitBracketIndex is called when exiting the bracketIndex production.
	ExitBracketIndex(c *BracketIndexContext)

	// ExitBracketStar is called when exiting the bracketStar production.
	ExitBracketStar(c *BracketStarContext)

	// ExitBracketSlice is called when exiting the bracketSlice production.
	ExitBracketSlice(c *BracketSliceContext)

	// ExitBracketFlatten is called when exiting the bracketFlatten production.
	ExitBracketFlatten(c *BracketFlattenContext)

	// ExitSelect is called when exiting the select production.
	ExitSelect(c *SelectContext)

	// ExitMultiSelectList is called when exiting the multiSelectList production.
	ExitMultiSelectList(c *MultiSelectListContext)

	// ExitMultiSelectHash is called when exiting the multiSelectHash production.
	ExitMultiSelectHash(c *MultiSelectHashContext)

	// ExitKeyvalExpr is called when exiting the keyvalExpr production.
	ExitKeyvalExpr(c *KeyvalExprContext)

	// ExitSliceNode is called when exiting the sliceNode production.
	ExitSliceNode(c *SliceNodeContext)

	// ExitNameParameter is called when exiting the nameParameter production.
	ExitNameParameter(c *NameParameterContext)

	// ExitNumberParameter is called when exiting the numberParameter production.
	ExitNumberParameter(c *NumberParameterContext)

	// ExitFunctionExpression is called when exiting the functionExpression production.
	ExitFunctionExpression(c *FunctionExpressionContext)

	// ExitFunctionArg is called when exiting the functionArg production.
	ExitFunctionArg(c *FunctionArgContext)

	// ExitCurrentNode is called when exiting the currentNode production.
	ExitCurrentNode(c *CurrentNodeContext)

	// ExitExpressionType is called when exiting the expressionType production.
	ExitExpressionType(c *ExpressionTypeContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitJsonObject is called when exiting the jsonObject production.
	ExitJsonObject(c *JsonObjectContext)

	// ExitJsonObjectPair is called when exiting the jsonObjectPair production.
	ExitJsonObjectPair(c *JsonObjectPairContext)

	// ExitJsonArray is called when exiting the jsonArray production.
	ExitJsonArray(c *JsonArrayContext)

	// ExitJsonStringValue is called when exiting the jsonStringValue production.
	ExitJsonStringValue(c *JsonStringValueContext)

	// ExitJsonNumberValue is called when exiting the jsonNumberValue production.
	ExitJsonNumberValue(c *JsonNumberValueContext)

	// ExitJsonObjectValue is called when exiting the jsonObjectValue production.
	ExitJsonObjectValue(c *JsonObjectValueContext)

	// ExitJsonArrayValue is called when exiting the jsonArrayValue production.
	ExitJsonArrayValue(c *JsonArrayValueContext)

	// ExitJsonConstantValue is called when exiting the jsonConstantValue production.
	ExitJsonConstantValue(c *JsonConstantValueContext)
}
