// Code generated from SimpleSql.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SimpleSql

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSimpleSqlListener is a complete listener for a parse tree produced by SimpleSqlParser.
type BaseSimpleSqlListener struct{}

var _ SimpleSqlListener = &BaseSimpleSqlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSimpleSqlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSimpleSqlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSimpleSqlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSimpleSqlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSingleStatement is called when production singleStatement is entered.
func (s *BaseSimpleSqlListener) EnterSingleStatement(ctx *SingleStatementContext) {}

// ExitSingleStatement is called when production singleStatement is exited.
func (s *BaseSimpleSqlListener) ExitSingleStatement(ctx *SingleStatementContext) {}

// EnterSelect is called when production select is entered.
func (s *BaseSimpleSqlListener) EnterSelect(ctx *SelectContext) {}

// ExitSelect is called when production select is exited.
func (s *BaseSimpleSqlListener) ExitSelect(ctx *SelectContext) {}

// EnterCteSelect is called when production cteSelect is entered.
func (s *BaseSimpleSqlListener) EnterCteSelect(ctx *CteSelectContext) {}

// ExitCteSelect is called when production cteSelect is exited.
func (s *BaseSimpleSqlListener) ExitCteSelect(ctx *CteSelectContext) {}

// EnterCte is called when production cte is entered.
func (s *BaseSimpleSqlListener) EnterCte(ctx *CteContext) {}

// ExitCte is called when production cte is exited.
func (s *BaseSimpleSqlListener) ExitCte(ctx *CteContext) {}

// EnterUnionSelect is called when production unionSelect is entered.
func (s *BaseSimpleSqlListener) EnterUnionSelect(ctx *UnionSelectContext) {}

// ExitUnionSelect is called when production unionSelect is exited.
func (s *BaseSimpleSqlListener) ExitUnionSelect(ctx *UnionSelectContext) {}

// EnterUnionItem is called when production unionItem is entered.
func (s *BaseSimpleSqlListener) EnterUnionItem(ctx *UnionItemContext) {}

// ExitUnionItem is called when production unionItem is exited.
func (s *BaseSimpleSqlListener) ExitUnionItem(ctx *UnionItemContext) {}

// EnterPrimarySelect is called when production primarySelect is entered.
func (s *BaseSimpleSqlListener) EnterPrimarySelect(ctx *PrimarySelectContext) {}

// ExitPrimarySelect is called when production primarySelect is exited.
func (s *BaseSimpleSqlListener) ExitPrimarySelect(ctx *PrimarySelectContext) {}

// EnterAllSelectItem is called when production allSelectItem is entered.
func (s *BaseSimpleSqlListener) EnterAllSelectItem(ctx *AllSelectItemContext) {}

// ExitAllSelectItem is called when production allSelectItem is exited.
func (s *BaseSimpleSqlListener) ExitAllSelectItem(ctx *AllSelectItemContext) {}

// EnterExpressionSelectItem is called when production expressionSelectItem is entered.
func (s *BaseSimpleSqlListener) EnterExpressionSelectItem(ctx *ExpressionSelectItemContext) {}

// ExitExpressionSelectItem is called when production expressionSelectItem is exited.
func (s *BaseSimpleSqlListener) ExitExpressionSelectItem(ctx *ExpressionSelectItemContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSimpleSqlListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSimpleSqlListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBinaryBooleanExpression is called when production binaryBooleanExpression is entered.
func (s *BaseSimpleSqlListener) EnterBinaryBooleanExpression(ctx *BinaryBooleanExpressionContext) {}

// ExitBinaryBooleanExpression is called when production binaryBooleanExpression is exited.
func (s *BaseSimpleSqlListener) ExitBinaryBooleanExpression(ctx *BinaryBooleanExpressionContext) {}

// EnterPredicatedBooleanExpression is called when production predicatedBooleanExpression is entered.
func (s *BaseSimpleSqlListener) EnterPredicatedBooleanExpression(ctx *PredicatedBooleanExpressionContext) {
}

// ExitPredicatedBooleanExpression is called when production predicatedBooleanExpression is exited.
func (s *BaseSimpleSqlListener) ExitPredicatedBooleanExpression(ctx *PredicatedBooleanExpressionContext) {
}

// EnterUnaryBooleanExpression is called when production unaryBooleanExpression is entered.
func (s *BaseSimpleSqlListener) EnterUnaryBooleanExpression(ctx *UnaryBooleanExpressionContext) {}

// ExitUnaryBooleanExpression is called when production unaryBooleanExpression is exited.
func (s *BaseSimpleSqlListener) ExitUnaryBooleanExpression(ctx *UnaryBooleanExpressionContext) {}

// EnterCastBooleanExpression is called when production castBooleanExpression is entered.
func (s *BaseSimpleSqlListener) EnterCastBooleanExpression(ctx *CastBooleanExpressionContext) {}

// ExitCastBooleanExpression is called when production castBooleanExpression is exited.
func (s *BaseSimpleSqlListener) ExitCastBooleanExpression(ctx *CastBooleanExpressionContext) {}

// EnterCmpPredicate is called when production cmpPredicate is entered.
func (s *BaseSimpleSqlListener) EnterCmpPredicate(ctx *CmpPredicateContext) {}

// ExitCmpPredicate is called when production cmpPredicate is exited.
func (s *BaseSimpleSqlListener) ExitCmpPredicate(ctx *CmpPredicateContext) {}

// EnterIsNullPredicate is called when production isNullPredicate is entered.
func (s *BaseSimpleSqlListener) EnterIsNullPredicate(ctx *IsNullPredicateContext) {}

// ExitIsNullPredicate is called when production isNullPredicate is exited.
func (s *BaseSimpleSqlListener) ExitIsNullPredicate(ctx *IsNullPredicateContext) {}

// EnterInListPredicate is called when production inListPredicate is entered.
func (s *BaseSimpleSqlListener) EnterInListPredicate(ctx *InListPredicateContext) {}

// ExitInListPredicate is called when production inListPredicate is exited.
func (s *BaseSimpleSqlListener) ExitInListPredicate(ctx *InListPredicateContext) {}

// EnterInSelectPredicate is called when production inSelectPredicate is entered.
func (s *BaseSimpleSqlListener) EnterInSelectPredicate(ctx *InSelectPredicateContext) {}

// ExitInSelectPredicate is called when production inSelectPredicate is exited.
func (s *BaseSimpleSqlListener) ExitInSelectPredicate(ctx *InSelectPredicateContext) {}

// EnterLikePredicate is called when production likePredicate is entered.
func (s *BaseSimpleSqlListener) EnterLikePredicate(ctx *LikePredicateContext) {}

// ExitLikePredicate is called when production likePredicate is exited.
func (s *BaseSimpleSqlListener) ExitLikePredicate(ctx *LikePredicateContext) {}

// EnterPrimaryValueExpression is called when production primaryValueExpression is entered.
func (s *BaseSimpleSqlListener) EnterPrimaryValueExpression(ctx *PrimaryValueExpressionContext) {}

// ExitPrimaryValueExpression is called when production primaryValueExpression is exited.
func (s *BaseSimpleSqlListener) ExitPrimaryValueExpression(ctx *PrimaryValueExpressionContext) {}

// EnterUnaryValueExpression is called when production unaryValueExpression is entered.
func (s *BaseSimpleSqlListener) EnterUnaryValueExpression(ctx *UnaryValueExpressionContext) {}

// ExitUnaryValueExpression is called when production unaryValueExpression is exited.
func (s *BaseSimpleSqlListener) ExitUnaryValueExpression(ctx *UnaryValueExpressionContext) {}

// EnterArithValueExpression is called when production arithValueExpression is entered.
func (s *BaseSimpleSqlListener) EnterArithValueExpression(ctx *ArithValueExpressionContext) {}

// ExitArithValueExpression is called when production arithValueExpression is exited.
func (s *BaseSimpleSqlListener) ExitArithValueExpression(ctx *ArithValueExpressionContext) {}

// EnterFunctionCallExpression is called when production functionCallExpression is entered.
func (s *BaseSimpleSqlListener) EnterFunctionCallExpression(ctx *FunctionCallExpressionContext) {}

// ExitFunctionCallExpression is called when production functionCallExpression is exited.
func (s *BaseSimpleSqlListener) ExitFunctionCallExpression(ctx *FunctionCallExpressionContext) {}

// EnterStarFunctionCallExpression is called when production starFunctionCallExpression is entered.
func (s *BaseSimpleSqlListener) EnterStarFunctionCallExpression(ctx *StarFunctionCallExpressionContext) {
}

// ExitStarFunctionCallExpression is called when production starFunctionCallExpression is exited.
func (s *BaseSimpleSqlListener) ExitStarFunctionCallExpression(ctx *StarFunctionCallExpressionContext) {
}

// EnterCaseExpression is called when production caseExpression is entered.
func (s *BaseSimpleSqlListener) EnterCaseExpression(ctx *CaseExpressionContext) {}

// ExitCaseExpression is called when production caseExpression is exited.
func (s *BaseSimpleSqlListener) ExitCaseExpression(ctx *CaseExpressionContext) {}

// EnterSelectExpression is called when production selectExpression is entered.
func (s *BaseSimpleSqlListener) EnterSelectExpression(ctx *SelectExpressionContext) {}

// ExitSelectExpression is called when production selectExpression is exited.
func (s *BaseSimpleSqlListener) ExitSelectExpression(ctx *SelectExpressionContext) {}

// EnterParenExpression is called when production parenExpression is entered.
func (s *BaseSimpleSqlListener) EnterParenExpression(ctx *ParenExpressionContext) {}

// ExitParenExpression is called when production parenExpression is exited.
func (s *BaseSimpleSqlListener) ExitParenExpression(ctx *ParenExpressionContext) {}

// EnterSimplePrimaryExpression is called when production simplePrimaryExpression is entered.
func (s *BaseSimpleSqlListener) EnterSimplePrimaryExpression(ctx *SimplePrimaryExpressionContext) {}

// ExitSimplePrimaryExpression is called when production simplePrimaryExpression is exited.
func (s *BaseSimpleSqlListener) ExitSimplePrimaryExpression(ctx *SimplePrimaryExpressionContext) {}

// EnterSimpleExpression is called when production simpleExpression is entered.
func (s *BaseSimpleSqlListener) EnterSimpleExpression(ctx *SimpleExpressionContext) {}

// ExitSimpleExpression is called when production simpleExpression is exited.
func (s *BaseSimpleSqlListener) ExitSimpleExpression(ctx *SimpleExpressionContext) {}

// EnterCaseItem is called when production caseItem is entered.
func (s *BaseSimpleSqlListener) EnterCaseItem(ctx *CaseItemContext) {}

// ExitCaseItem is called when production caseItem is exited.
func (s *BaseSimpleSqlListener) ExitCaseItem(ctx *CaseItemContext) {}

// EnterOver is called when production over is entered.
func (s *BaseSimpleSqlListener) EnterOver(ctx *OverContext) {}

// ExitOver is called when production over is exited.
func (s *BaseSimpleSqlListener) ExitOver(ctx *OverContext) {}

// EnterSortItem is called when production sortItem is entered.
func (s *BaseSimpleSqlListener) EnterSortItem(ctx *SortItemContext) {}

// ExitSortItem is called when production sortItem is exited.
func (s *BaseSimpleSqlListener) ExitSortItem(ctx *SortItemContext) {}

// EnterAliasedRelation is called when production aliasedRelation is entered.
func (s *BaseSimpleSqlListener) EnterAliasedRelation(ctx *AliasedRelationContext) {}

// ExitAliasedRelation is called when production aliasedRelation is exited.
func (s *BaseSimpleSqlListener) ExitAliasedRelation(ctx *AliasedRelationContext) {}

// EnterJoinRelation is called when production joinRelation is entered.
func (s *BaseSimpleSqlListener) EnterJoinRelation(ctx *JoinRelationContext) {}

// ExitJoinRelation is called when production joinRelation is exited.
func (s *BaseSimpleSqlListener) ExitJoinRelation(ctx *JoinRelationContext) {}

// EnterSelectRelation is called when production selectRelation is entered.
func (s *BaseSimpleSqlListener) EnterSelectRelation(ctx *SelectRelationContext) {}

// ExitSelectRelation is called when production selectRelation is exited.
func (s *BaseSimpleSqlListener) ExitSelectRelation(ctx *SelectRelationContext) {}

// EnterTableRelation is called when production tableRelation is entered.
func (s *BaseSimpleSqlListener) EnterTableRelation(ctx *TableRelationContext) {}

// ExitTableRelation is called when production tableRelation is exited.
func (s *BaseSimpleSqlListener) ExitTableRelation(ctx *TableRelationContext) {}

// EnterParenRelation is called when production parenRelation is entered.
func (s *BaseSimpleSqlListener) EnterParenRelation(ctx *ParenRelationContext) {}

// ExitParenRelation is called when production parenRelation is exited.
func (s *BaseSimpleSqlListener) ExitParenRelation(ctx *ParenRelationContext) {}

// EnterGroupBy is called when production groupBy is entered.
func (s *BaseSimpleSqlListener) EnterGroupBy(ctx *GroupByContext) {}

// ExitGroupBy is called when production groupBy is exited.
func (s *BaseSimpleSqlListener) ExitGroupBy(ctx *GroupByContext) {}

// EnterQualifiedName is called when production qualifiedName is entered.
func (s *BaseSimpleSqlListener) EnterQualifiedName(ctx *QualifiedNameContext) {}

// ExitQualifiedName is called when production qualifiedName is exited.
func (s *BaseSimpleSqlListener) ExitQualifiedName(ctx *QualifiedNameContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseSimpleSqlListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseSimpleSqlListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseSimpleSqlListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseSimpleSqlListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterQuotedIdentifier is called when production quotedIdentifier is entered.
func (s *BaseSimpleSqlListener) EnterQuotedIdentifier(ctx *QuotedIdentifierContext) {}

// ExitQuotedIdentifier is called when production quotedIdentifier is exited.
func (s *BaseSimpleSqlListener) ExitQuotedIdentifier(ctx *QuotedIdentifierContext) {}

// EnterIntegerNumber is called when production integerNumber is entered.
func (s *BaseSimpleSqlListener) EnterIntegerNumber(ctx *IntegerNumberContext) {}

// ExitIntegerNumber is called when production integerNumber is exited.
func (s *BaseSimpleSqlListener) ExitIntegerNumber(ctx *IntegerNumberContext) {}

// EnterDecimalNumber is called when production decimalNumber is entered.
func (s *BaseSimpleSqlListener) EnterDecimalNumber(ctx *DecimalNumberContext) {}

// ExitDecimalNumber is called when production decimalNumber is exited.
func (s *BaseSimpleSqlListener) ExitDecimalNumber(ctx *DecimalNumberContext) {}

// EnterFloatNumber is called when production floatNumber is entered.
func (s *BaseSimpleSqlListener) EnterFloatNumber(ctx *FloatNumberContext) {}

// ExitFloatNumber is called when production floatNumber is exited.
func (s *BaseSimpleSqlListener) ExitFloatNumber(ctx *FloatNumberContext) {}

// EnterStr is called when production str is entered.
func (s *BaseSimpleSqlListener) EnterStr(ctx *StrContext) {}

// ExitStr is called when production str is exited.
func (s *BaseSimpleSqlListener) ExitStr(ctx *StrContext) {}

// EnterNull is called when production null is entered.
func (s *BaseSimpleSqlListener) EnterNull(ctx *NullContext) {}

// ExitNull is called when production null is exited.
func (s *BaseSimpleSqlListener) ExitNull(ctx *NullContext) {}

// EnterTrue is called when production true is entered.
func (s *BaseSimpleSqlListener) EnterTrue(ctx *TrueContext) {}

// ExitTrue is called when production true is exited.
func (s *BaseSimpleSqlListener) ExitTrue(ctx *TrueContext) {}

// EnterFalse is called when production false is entered.
func (s *BaseSimpleSqlListener) EnterFalse(ctx *FalseContext) {}

// ExitFalse is called when production false is exited.
func (s *BaseSimpleSqlListener) ExitFalse(ctx *FalseContext) {}

// EnterSetQuantifier is called when production setQuantifier is entered.
func (s *BaseSimpleSqlListener) EnterSetQuantifier(ctx *SetQuantifierContext) {}

// ExitSetQuantifier is called when production setQuantifier is exited.
func (s *BaseSimpleSqlListener) ExitSetQuantifier(ctx *SetQuantifierContext) {}

// EnterJoinType is called when production joinType is entered.
func (s *BaseSimpleSqlListener) EnterJoinType(ctx *JoinTypeContext) {}

// ExitJoinType is called when production joinType is exited.
func (s *BaseSimpleSqlListener) ExitJoinType(ctx *JoinTypeContext) {}

// EnterCmpOp is called when production cmpOp is entered.
func (s *BaseSimpleSqlListener) EnterCmpOp(ctx *CmpOpContext) {}

// ExitCmpOp is called when production cmpOp is exited.
func (s *BaseSimpleSqlListener) ExitCmpOp(ctx *CmpOpContext) {}

// EnterArithOp is called when production arithOp is entered.
func (s *BaseSimpleSqlListener) EnterArithOp(ctx *ArithOpContext) {}

// ExitArithOp is called when production arithOp is exited.
func (s *BaseSimpleSqlListener) ExitArithOp(ctx *ArithOpContext) {}

// EnterUnaryOp is called when production unaryOp is entered.
func (s *BaseSimpleSqlListener) EnterUnaryOp(ctx *UnaryOpContext) {}

// ExitUnaryOp is called when production unaryOp is exited.
func (s *BaseSimpleSqlListener) ExitUnaryOp(ctx *UnaryOpContext) {}

// EnterUnquotedIdentifier is called when production unquotedIdentifier is entered.
func (s *BaseSimpleSqlListener) EnterUnquotedIdentifier(ctx *UnquotedIdentifierContext) {}

// ExitUnquotedIdentifier is called when production unquotedIdentifier is exited.
func (s *BaseSimpleSqlListener) ExitUnquotedIdentifier(ctx *UnquotedIdentifierContext) {}
