// Code generated from SimpleSql.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SimpleSql

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SimpleSqlListener is a complete listener for a parse tree produced by SimpleSqlParser.
type SimpleSqlListener interface {
	antlr.ParseTreeListener

	// EnterSingleStatement is called when entering the singleStatement production.
	EnterSingleStatement(c *SingleStatementContext)

	// EnterSelect is called when entering the select production.
	EnterSelect(c *SelectContext)

	// EnterCteSelect is called when entering the cteSelect production.
	EnterCteSelect(c *CteSelectContext)

	// EnterCte is called when entering the cte production.
	EnterCte(c *CteContext)

	// EnterUnionSelect is called when entering the unionSelect production.
	EnterUnionSelect(c *UnionSelectContext)

	// EnterUnionItem is called when entering the unionItem production.
	EnterUnionItem(c *UnionItemContext)

	// EnterPrimarySelect is called when entering the primarySelect production.
	EnterPrimarySelect(c *PrimarySelectContext)

	// EnterAllSelectItem is called when entering the allSelectItem production.
	EnterAllSelectItem(c *AllSelectItemContext)

	// EnterExpressionSelectItem is called when entering the expressionSelectItem production.
	EnterExpressionSelectItem(c *ExpressionSelectItemContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBinaryBooleanExpression is called when entering the binaryBooleanExpression production.
	EnterBinaryBooleanExpression(c *BinaryBooleanExpressionContext)

	// EnterPredicatedBooleanExpression is called when entering the predicatedBooleanExpression production.
	EnterPredicatedBooleanExpression(c *PredicatedBooleanExpressionContext)

	// EnterUnaryBooleanExpression is called when entering the unaryBooleanExpression production.
	EnterUnaryBooleanExpression(c *UnaryBooleanExpressionContext)

	// EnterCastBooleanExpression is called when entering the castBooleanExpression production.
	EnterCastBooleanExpression(c *CastBooleanExpressionContext)

	// EnterCmpPredicate is called when entering the cmpPredicate production.
	EnterCmpPredicate(c *CmpPredicateContext)

	// EnterIsNullPredicate is called when entering the isNullPredicate production.
	EnterIsNullPredicate(c *IsNullPredicateContext)

	// EnterInListPredicate is called when entering the inListPredicate production.
	EnterInListPredicate(c *InListPredicateContext)

	// EnterInSelectPredicate is called when entering the inSelectPredicate production.
	EnterInSelectPredicate(c *InSelectPredicateContext)

	// EnterLikePredicate is called when entering the likePredicate production.
	EnterLikePredicate(c *LikePredicateContext)

	// EnterPrimaryValueExpression is called when entering the primaryValueExpression production.
	EnterPrimaryValueExpression(c *PrimaryValueExpressionContext)

	// EnterUnaryValueExpression is called when entering the unaryValueExpression production.
	EnterUnaryValueExpression(c *UnaryValueExpressionContext)

	// EnterArithValueExpression is called when entering the arithValueExpression production.
	EnterArithValueExpression(c *ArithValueExpressionContext)

	// EnterFunctionCallExpression is called when entering the functionCallExpression production.
	EnterFunctionCallExpression(c *FunctionCallExpressionContext)

	// EnterStarFunctionCallExpression is called when entering the starFunctionCallExpression production.
	EnterStarFunctionCallExpression(c *StarFunctionCallExpressionContext)

	// EnterCaseExpression is called when entering the caseExpression production.
	EnterCaseExpression(c *CaseExpressionContext)

	// EnterSelectExpression is called when entering the selectExpression production.
	EnterSelectExpression(c *SelectExpressionContext)

	// EnterParenExpression is called when entering the parenExpression production.
	EnterParenExpression(c *ParenExpressionContext)

	// EnterSimplePrimaryExpression is called when entering the simplePrimaryExpression production.
	EnterSimplePrimaryExpression(c *SimplePrimaryExpressionContext)

	// EnterSimpleExpression is called when entering the simpleExpression production.
	EnterSimpleExpression(c *SimpleExpressionContext)

	// EnterCaseItem is called when entering the caseItem production.
	EnterCaseItem(c *CaseItemContext)

	// EnterOver is called when entering the over production.
	EnterOver(c *OverContext)

	// EnterSortItem is called when entering the sortItem production.
	EnterSortItem(c *SortItemContext)

	// EnterAliasedRelation is called when entering the aliasedRelation production.
	EnterAliasedRelation(c *AliasedRelationContext)

	// EnterJoinRelation is called when entering the joinRelation production.
	EnterJoinRelation(c *JoinRelationContext)

	// EnterSelectRelation is called when entering the selectRelation production.
	EnterSelectRelation(c *SelectRelationContext)

	// EnterTableRelation is called when entering the tableRelation production.
	EnterTableRelation(c *TableRelationContext)

	// EnterParenRelation is called when entering the parenRelation production.
	EnterParenRelation(c *ParenRelationContext)

	// EnterGroupBy is called when entering the groupBy production.
	EnterGroupBy(c *GroupByContext)

	// EnterQualifiedName is called when entering the qualifiedName production.
	EnterQualifiedName(c *QualifiedNameContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterQuotedIdentifier is called when entering the quotedIdentifier production.
	EnterQuotedIdentifier(c *QuotedIdentifierContext)

	// EnterIntegerNumber is called when entering the integerNumber production.
	EnterIntegerNumber(c *IntegerNumberContext)

	// EnterDecimalNumber is called when entering the decimalNumber production.
	EnterDecimalNumber(c *DecimalNumberContext)

	// EnterFloatNumber is called when entering the floatNumber production.
	EnterFloatNumber(c *FloatNumberContext)

	// EnterStr is called when entering the str production.
	EnterStr(c *StrContext)

	// EnterNull is called when entering the null production.
	EnterNull(c *NullContext)

	// EnterTrue is called when entering the true production.
	EnterTrue(c *TrueContext)

	// EnterFalse is called when entering the false production.
	EnterFalse(c *FalseContext)

	// EnterSetQuantifier is called when entering the setQuantifier production.
	EnterSetQuantifier(c *SetQuantifierContext)

	// EnterJoinType is called when entering the joinType production.
	EnterJoinType(c *JoinTypeContext)

	// EnterCmpOp is called when entering the cmpOp production.
	EnterCmpOp(c *CmpOpContext)

	// EnterArithOp is called when entering the arithOp production.
	EnterArithOp(c *ArithOpContext)

	// EnterUnaryOp is called when entering the unaryOp production.
	EnterUnaryOp(c *UnaryOpContext)

	// EnterUnquotedIdentifier is called when entering the unquotedIdentifier production.
	EnterUnquotedIdentifier(c *UnquotedIdentifierContext)

	// ExitSingleStatement is called when exiting the singleStatement production.
	ExitSingleStatement(c *SingleStatementContext)

	// ExitSelect is called when exiting the select production.
	ExitSelect(c *SelectContext)

	// ExitCteSelect is called when exiting the cteSelect production.
	ExitCteSelect(c *CteSelectContext)

	// ExitCte is called when exiting the cte production.
	ExitCte(c *CteContext)

	// ExitUnionSelect is called when exiting the unionSelect production.
	ExitUnionSelect(c *UnionSelectContext)

	// ExitUnionItem is called when exiting the unionItem production.
	ExitUnionItem(c *UnionItemContext)

	// ExitPrimarySelect is called when exiting the primarySelect production.
	ExitPrimarySelect(c *PrimarySelectContext)

	// ExitAllSelectItem is called when exiting the allSelectItem production.
	ExitAllSelectItem(c *AllSelectItemContext)

	// ExitExpressionSelectItem is called when exiting the expressionSelectItem production.
	ExitExpressionSelectItem(c *ExpressionSelectItemContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBinaryBooleanExpression is called when exiting the binaryBooleanExpression production.
	ExitBinaryBooleanExpression(c *BinaryBooleanExpressionContext)

	// ExitPredicatedBooleanExpression is called when exiting the predicatedBooleanExpression production.
	ExitPredicatedBooleanExpression(c *PredicatedBooleanExpressionContext)

	// ExitUnaryBooleanExpression is called when exiting the unaryBooleanExpression production.
	ExitUnaryBooleanExpression(c *UnaryBooleanExpressionContext)

	// ExitCastBooleanExpression is called when exiting the castBooleanExpression production.
	ExitCastBooleanExpression(c *CastBooleanExpressionContext)

	// ExitCmpPredicate is called when exiting the cmpPredicate production.
	ExitCmpPredicate(c *CmpPredicateContext)

	// ExitIsNullPredicate is called when exiting the isNullPredicate production.
	ExitIsNullPredicate(c *IsNullPredicateContext)

	// ExitInListPredicate is called when exiting the inListPredicate production.
	ExitInListPredicate(c *InListPredicateContext)

	// ExitInSelectPredicate is called when exiting the inSelectPredicate production.
	ExitInSelectPredicate(c *InSelectPredicateContext)

	// ExitLikePredicate is called when exiting the likePredicate production.
	ExitLikePredicate(c *LikePredicateContext)

	// ExitPrimaryValueExpression is called when exiting the primaryValueExpression production.
	ExitPrimaryValueExpression(c *PrimaryValueExpressionContext)

	// ExitUnaryValueExpression is called when exiting the unaryValueExpression production.
	ExitUnaryValueExpression(c *UnaryValueExpressionContext)

	// ExitArithValueExpression is called when exiting the arithValueExpression production.
	ExitArithValueExpression(c *ArithValueExpressionContext)

	// ExitFunctionCallExpression is called when exiting the functionCallExpression production.
	ExitFunctionCallExpression(c *FunctionCallExpressionContext)

	// ExitStarFunctionCallExpression is called when exiting the starFunctionCallExpression production.
	ExitStarFunctionCallExpression(c *StarFunctionCallExpressionContext)

	// ExitCaseExpression is called when exiting the caseExpression production.
	ExitCaseExpression(c *CaseExpressionContext)

	// ExitSelectExpression is called when exiting the selectExpression production.
	ExitSelectExpression(c *SelectExpressionContext)

	// ExitParenExpression is called when exiting the parenExpression production.
	ExitParenExpression(c *ParenExpressionContext)

	// ExitSimplePrimaryExpression is called when exiting the simplePrimaryExpression production.
	ExitSimplePrimaryExpression(c *SimplePrimaryExpressionContext)

	// ExitSimpleExpression is called when exiting the simpleExpression production.
	ExitSimpleExpression(c *SimpleExpressionContext)

	// ExitCaseItem is called when exiting the caseItem production.
	ExitCaseItem(c *CaseItemContext)

	// ExitOver is called when exiting the over production.
	ExitOver(c *OverContext)

	// ExitSortItem is called when exiting the sortItem production.
	ExitSortItem(c *SortItemContext)

	// ExitAliasedRelation is called when exiting the aliasedRelation production.
	ExitAliasedRelation(c *AliasedRelationContext)

	// ExitJoinRelation is called when exiting the joinRelation production.
	ExitJoinRelation(c *JoinRelationContext)

	// ExitSelectRelation is called when exiting the selectRelation production.
	ExitSelectRelation(c *SelectRelationContext)

	// ExitTableRelation is called when exiting the tableRelation production.
	ExitTableRelation(c *TableRelationContext)

	// ExitParenRelation is called when exiting the parenRelation production.
	ExitParenRelation(c *ParenRelationContext)

	// ExitGroupBy is called when exiting the groupBy production.
	ExitGroupBy(c *GroupByContext)

	// ExitQualifiedName is called when exiting the qualifiedName production.
	ExitQualifiedName(c *QualifiedNameContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitQuotedIdentifier is called when exiting the quotedIdentifier production.
	ExitQuotedIdentifier(c *QuotedIdentifierContext)

	// ExitIntegerNumber is called when exiting the integerNumber production.
	ExitIntegerNumber(c *IntegerNumberContext)

	// ExitDecimalNumber is called when exiting the decimalNumber production.
	ExitDecimalNumber(c *DecimalNumberContext)

	// ExitFloatNumber is called when exiting the floatNumber production.
	ExitFloatNumber(c *FloatNumberContext)

	// ExitStr is called when exiting the str production.
	ExitStr(c *StrContext)

	// ExitNull is called when exiting the null production.
	ExitNull(c *NullContext)

	// ExitTrue is called when exiting the true production.
	ExitTrue(c *TrueContext)

	// ExitFalse is called when exiting the false production.
	ExitFalse(c *FalseContext)

	// ExitSetQuantifier is called when exiting the setQuantifier production.
	ExitSetQuantifier(c *SetQuantifierContext)

	// ExitJoinType is called when exiting the joinType production.
	ExitJoinType(c *JoinTypeContext)

	// ExitCmpOp is called when exiting the cmpOp production.
	ExitCmpOp(c *CmpOpContext)

	// ExitArithOp is called when exiting the arithOp production.
	ExitArithOp(c *ArithOpContext)

	// ExitUnaryOp is called when exiting the unaryOp production.
	ExitUnaryOp(c *UnaryOpContext)

	// ExitUnquotedIdentifier is called when exiting the unquotedIdentifier production.
	ExitUnquotedIdentifier(c *UnquotedIdentifierContext)
}
