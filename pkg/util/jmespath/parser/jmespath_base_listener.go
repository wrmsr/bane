// Code generated from Jmespath.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Jmespath

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJmespathListener is a complete listener for a parse tree produced by JmespathParser.
type BaseJmespathListener struct{}

var _ JmespathListener = &BaseJmespathListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJmespathListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJmespathListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJmespathListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJmespathListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSingleExpression is called when production singleExpression is entered.
func (s *BaseJmespathListener) EnterSingleExpression(ctx *SingleExpressionContext) {}

// ExitSingleExpression is called when production singleExpression is exited.
func (s *BaseJmespathListener) ExitSingleExpression(ctx *SingleExpressionContext) {}

// EnterPipeExpression is called when production pipeExpression is entered.
func (s *BaseJmespathListener) EnterPipeExpression(ctx *PipeExpressionContext) {}

// ExitPipeExpression is called when production pipeExpression is exited.
func (s *BaseJmespathListener) ExitPipeExpression(ctx *PipeExpressionContext) {}

// EnterParameterExpression is called when production parameterExpression is entered.
func (s *BaseJmespathListener) EnterParameterExpression(ctx *ParameterExpressionContext) {}

// ExitParameterExpression is called when production parameterExpression is exited.
func (s *BaseJmespathListener) ExitParameterExpression(ctx *ParameterExpressionContext) {}

// EnterIdentifierExpression is called when production identifierExpression is entered.
func (s *BaseJmespathListener) EnterIdentifierExpression(ctx *IdentifierExpressionContext) {}

// ExitIdentifierExpression is called when production identifierExpression is exited.
func (s *BaseJmespathListener) ExitIdentifierExpression(ctx *IdentifierExpressionContext) {}

// EnterNotExpression is called when production notExpression is entered.
func (s *BaseJmespathListener) EnterNotExpression(ctx *NotExpressionContext) {}

// ExitNotExpression is called when production notExpression is exited.
func (s *BaseJmespathListener) ExitNotExpression(ctx *NotExpressionContext) {}

// EnterRawStringExpression is called when production rawStringExpression is entered.
func (s *BaseJmespathListener) EnterRawStringExpression(ctx *RawStringExpressionContext) {}

// ExitRawStringExpression is called when production rawStringExpression is exited.
func (s *BaseJmespathListener) ExitRawStringExpression(ctx *RawStringExpressionContext) {}

// EnterComparisonExpression is called when production comparisonExpression is entered.
func (s *BaseJmespathListener) EnterComparisonExpression(ctx *ComparisonExpressionContext) {}

// ExitComparisonExpression is called when production comparisonExpression is exited.
func (s *BaseJmespathListener) ExitComparisonExpression(ctx *ComparisonExpressionContext) {}

// EnterParenExpression is called when production parenExpression is entered.
func (s *BaseJmespathListener) EnterParenExpression(ctx *ParenExpressionContext) {}

// ExitParenExpression is called when production parenExpression is exited.
func (s *BaseJmespathListener) ExitParenExpression(ctx *ParenExpressionContext) {}

// EnterBracketExpression is called when production bracketExpression is entered.
func (s *BaseJmespathListener) EnterBracketExpression(ctx *BracketExpressionContext) {}

// ExitBracketExpression is called when production bracketExpression is exited.
func (s *BaseJmespathListener) ExitBracketExpression(ctx *BracketExpressionContext) {}

// EnterOrExpression is called when production orExpression is entered.
func (s *BaseJmespathListener) EnterOrExpression(ctx *OrExpressionContext) {}

// ExitOrExpression is called when production orExpression is exited.
func (s *BaseJmespathListener) ExitOrExpression(ctx *OrExpressionContext) {}

// EnterCurrentNodeExpression is called when production currentNodeExpression is entered.
func (s *BaseJmespathListener) EnterCurrentNodeExpression(ctx *CurrentNodeExpressionContext) {}

// ExitCurrentNodeExpression is called when production currentNodeExpression is exited.
func (s *BaseJmespathListener) ExitCurrentNodeExpression(ctx *CurrentNodeExpressionContext) {}

// EnterChainExpression is called when production chainExpression is entered.
func (s *BaseJmespathListener) EnterChainExpression(ctx *ChainExpressionContext) {}

// ExitChainExpression is called when production chainExpression is exited.
func (s *BaseJmespathListener) ExitChainExpression(ctx *ChainExpressionContext) {}

// EnterAndExpression is called when production andExpression is entered.
func (s *BaseJmespathListener) EnterAndExpression(ctx *AndExpressionContext) {}

// ExitAndExpression is called when production andExpression is exited.
func (s *BaseJmespathListener) ExitAndExpression(ctx *AndExpressionContext) {}

// EnterMultiSelectHashExpression is called when production multiSelectHashExpression is entered.
func (s *BaseJmespathListener) EnterMultiSelectHashExpression(ctx *MultiSelectHashExpressionContext) {
}

// ExitMultiSelectHashExpression is called when production multiSelectHashExpression is exited.
func (s *BaseJmespathListener) ExitMultiSelectHashExpression(ctx *MultiSelectHashExpressionContext) {}

// EnterWildcardExpression is called when production wildcardExpression is entered.
func (s *BaseJmespathListener) EnterWildcardExpression(ctx *WildcardExpressionContext) {}

// ExitWildcardExpression is called when production wildcardExpression is exited.
func (s *BaseJmespathListener) ExitWildcardExpression(ctx *WildcardExpressionContext) {}

// EnterFunctionCallExpression is called when production functionCallExpression is entered.
func (s *BaseJmespathListener) EnterFunctionCallExpression(ctx *FunctionCallExpressionContext) {}

// ExitFunctionCallExpression is called when production functionCallExpression is exited.
func (s *BaseJmespathListener) ExitFunctionCallExpression(ctx *FunctionCallExpressionContext) {}

// EnterMultiSelectListExpression is called when production multiSelectListExpression is entered.
func (s *BaseJmespathListener) EnterMultiSelectListExpression(ctx *MultiSelectListExpressionContext) {
}

// ExitMultiSelectListExpression is called when production multiSelectListExpression is exited.
func (s *BaseJmespathListener) ExitMultiSelectListExpression(ctx *MultiSelectListExpressionContext) {}

// EnterBracketedExpression is called when production bracketedExpression is entered.
func (s *BaseJmespathListener) EnterBracketedExpression(ctx *BracketedExpressionContext) {}

// ExitBracketedExpression is called when production bracketedExpression is exited.
func (s *BaseJmespathListener) ExitBracketedExpression(ctx *BracketedExpressionContext) {}

// EnterLiteralExpression is called when production literalExpression is entered.
func (s *BaseJmespathListener) EnterLiteralExpression(ctx *LiteralExpressionContext) {}

// ExitLiteralExpression is called when production literalExpression is exited.
func (s *BaseJmespathListener) ExitLiteralExpression(ctx *LiteralExpressionContext) {}

// EnterChainedExpression is called when production chainedExpression is entered.
func (s *BaseJmespathListener) EnterChainedExpression(ctx *ChainedExpressionContext) {}

// ExitChainedExpression is called when production chainedExpression is exited.
func (s *BaseJmespathListener) ExitChainedExpression(ctx *ChainedExpressionContext) {}

// EnterWildcard is called when production wildcard is entered.
func (s *BaseJmespathListener) EnterWildcard(ctx *WildcardContext) {}

// ExitWildcard is called when production wildcard is exited.
func (s *BaseJmespathListener) ExitWildcard(ctx *WildcardContext) {}

// EnterBracketIndex is called when production bracketIndex is entered.
func (s *BaseJmespathListener) EnterBracketIndex(ctx *BracketIndexContext) {}

// ExitBracketIndex is called when production bracketIndex is exited.
func (s *BaseJmespathListener) ExitBracketIndex(ctx *BracketIndexContext) {}

// EnterBracketStar is called when production bracketStar is entered.
func (s *BaseJmespathListener) EnterBracketStar(ctx *BracketStarContext) {}

// ExitBracketStar is called when production bracketStar is exited.
func (s *BaseJmespathListener) ExitBracketStar(ctx *BracketStarContext) {}

// EnterBracketSlice is called when production bracketSlice is entered.
func (s *BaseJmespathListener) EnterBracketSlice(ctx *BracketSliceContext) {}

// ExitBracketSlice is called when production bracketSlice is exited.
func (s *BaseJmespathListener) ExitBracketSlice(ctx *BracketSliceContext) {}

// EnterBracketFlatten is called when production bracketFlatten is entered.
func (s *BaseJmespathListener) EnterBracketFlatten(ctx *BracketFlattenContext) {}

// ExitBracketFlatten is called when production bracketFlatten is exited.
func (s *BaseJmespathListener) ExitBracketFlatten(ctx *BracketFlattenContext) {}

// EnterSelect is called when production select is entered.
func (s *BaseJmespathListener) EnterSelect(ctx *SelectContext) {}

// ExitSelect is called when production select is exited.
func (s *BaseJmespathListener) ExitSelect(ctx *SelectContext) {}

// EnterMultiSelectList is called when production multiSelectList is entered.
func (s *BaseJmespathListener) EnterMultiSelectList(ctx *MultiSelectListContext) {}

// ExitMultiSelectList is called when production multiSelectList is exited.
func (s *BaseJmespathListener) ExitMultiSelectList(ctx *MultiSelectListContext) {}

// EnterMultiSelectHash is called when production multiSelectHash is entered.
func (s *BaseJmespathListener) EnterMultiSelectHash(ctx *MultiSelectHashContext) {}

// ExitMultiSelectHash is called when production multiSelectHash is exited.
func (s *BaseJmespathListener) ExitMultiSelectHash(ctx *MultiSelectHashContext) {}

// EnterKeyvalExpr is called when production keyvalExpr is entered.
func (s *BaseJmespathListener) EnterKeyvalExpr(ctx *KeyvalExprContext) {}

// ExitKeyvalExpr is called when production keyvalExpr is exited.
func (s *BaseJmespathListener) ExitKeyvalExpr(ctx *KeyvalExprContext) {}

// EnterSliceNode is called when production sliceNode is entered.
func (s *BaseJmespathListener) EnterSliceNode(ctx *SliceNodeContext) {}

// ExitSliceNode is called when production sliceNode is exited.
func (s *BaseJmespathListener) ExitSliceNode(ctx *SliceNodeContext) {}

// EnterNameParameter is called when production nameParameter is entered.
func (s *BaseJmespathListener) EnterNameParameter(ctx *NameParameterContext) {}

// ExitNameParameter is called when production nameParameter is exited.
func (s *BaseJmespathListener) ExitNameParameter(ctx *NameParameterContext) {}

// EnterNumberParameter is called when production numberParameter is entered.
func (s *BaseJmespathListener) EnterNumberParameter(ctx *NumberParameterContext) {}

// ExitNumberParameter is called when production numberParameter is exited.
func (s *BaseJmespathListener) ExitNumberParameter(ctx *NumberParameterContext) {}

// EnterFunctionExpression is called when production functionExpression is entered.
func (s *BaseJmespathListener) EnterFunctionExpression(ctx *FunctionExpressionContext) {}

// ExitFunctionExpression is called when production functionExpression is exited.
func (s *BaseJmespathListener) ExitFunctionExpression(ctx *FunctionExpressionContext) {}

// EnterFunctionArg is called when production functionArg is entered.
func (s *BaseJmespathListener) EnterFunctionArg(ctx *FunctionArgContext) {}

// ExitFunctionArg is called when production functionArg is exited.
func (s *BaseJmespathListener) ExitFunctionArg(ctx *FunctionArgContext) {}

// EnterCurrentNode is called when production currentNode is entered.
func (s *BaseJmespathListener) EnterCurrentNode(ctx *CurrentNodeContext) {}

// ExitCurrentNode is called when production currentNode is exited.
func (s *BaseJmespathListener) ExitCurrentNode(ctx *CurrentNodeContext) {}

// EnterExpressionType is called when production expressionType is entered.
func (s *BaseJmespathListener) EnterExpressionType(ctx *ExpressionTypeContext) {}

// ExitExpressionType is called when production expressionType is exited.
func (s *BaseJmespathListener) ExitExpressionType(ctx *ExpressionTypeContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseJmespathListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseJmespathListener) ExitLiteral(ctx *LiteralContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseJmespathListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseJmespathListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterJsonObject is called when production jsonObject is entered.
func (s *BaseJmespathListener) EnterJsonObject(ctx *JsonObjectContext) {}

// ExitJsonObject is called when production jsonObject is exited.
func (s *BaseJmespathListener) ExitJsonObject(ctx *JsonObjectContext) {}

// EnterJsonObjectPair is called when production jsonObjectPair is entered.
func (s *BaseJmespathListener) EnterJsonObjectPair(ctx *JsonObjectPairContext) {}

// ExitJsonObjectPair is called when production jsonObjectPair is exited.
func (s *BaseJmespathListener) ExitJsonObjectPair(ctx *JsonObjectPairContext) {}

// EnterJsonArray is called when production jsonArray is entered.
func (s *BaseJmespathListener) EnterJsonArray(ctx *JsonArrayContext) {}

// ExitJsonArray is called when production jsonArray is exited.
func (s *BaseJmespathListener) ExitJsonArray(ctx *JsonArrayContext) {}

// EnterJsonStringValue is called when production jsonStringValue is entered.
func (s *BaseJmespathListener) EnterJsonStringValue(ctx *JsonStringValueContext) {}

// ExitJsonStringValue is called when production jsonStringValue is exited.
func (s *BaseJmespathListener) ExitJsonStringValue(ctx *JsonStringValueContext) {}

// EnterJsonNumberValue is called when production jsonNumberValue is entered.
func (s *BaseJmespathListener) EnterJsonNumberValue(ctx *JsonNumberValueContext) {}

// ExitJsonNumberValue is called when production jsonNumberValue is exited.
func (s *BaseJmespathListener) ExitJsonNumberValue(ctx *JsonNumberValueContext) {}

// EnterJsonObjectValue is called when production jsonObjectValue is entered.
func (s *BaseJmespathListener) EnterJsonObjectValue(ctx *JsonObjectValueContext) {}

// ExitJsonObjectValue is called when production jsonObjectValue is exited.
func (s *BaseJmespathListener) ExitJsonObjectValue(ctx *JsonObjectValueContext) {}

// EnterJsonArrayValue is called when production jsonArrayValue is entered.
func (s *BaseJmespathListener) EnterJsonArrayValue(ctx *JsonArrayValueContext) {}

// ExitJsonArrayValue is called when production jsonArrayValue is exited.
func (s *BaseJmespathListener) ExitJsonArrayValue(ctx *JsonArrayValueContext) {}

// EnterJsonConstantValue is called when production jsonConstantValue is entered.
func (s *BaseJmespathListener) EnterJsonConstantValue(ctx *JsonConstantValueContext) {}

// ExitJsonConstantValue is called when production jsonConstantValue is exited.
func (s *BaseJmespathListener) ExitJsonConstantValue(ctx *JsonConstantValueContext) {}
