// Code generated from SimpleSql.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // SimpleSql

import antlr "github.com/wrmsr/bane/core/antlr/runtime"

// A complete Visitor for a parse tree produced by SimpleSqlParser.
type SimpleSqlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SimpleSqlParser#singleStatement.
	VisitSingleStatement(ctx *SingleStatementContext) any

	// Visit a parse tree produced by SimpleSqlParser#select.
	VisitSelect(ctx *SelectContext) any

	// Visit a parse tree produced by SimpleSqlParser#cteSelect.
	VisitCteSelect(ctx *CteSelectContext) any

	// Visit a parse tree produced by SimpleSqlParser#cte.
	VisitCte(ctx *CteContext) any

	// Visit a parse tree produced by SimpleSqlParser#unionSelect.
	VisitUnionSelect(ctx *UnionSelectContext) any

	// Visit a parse tree produced by SimpleSqlParser#unionItem.
	VisitUnionItem(ctx *UnionItemContext) any

	// Visit a parse tree produced by SimpleSqlParser#primarySelect.
	VisitPrimarySelect(ctx *PrimarySelectContext) any

	// Visit a parse tree produced by SimpleSqlParser#allSelectItem.
	VisitAllSelectItem(ctx *AllSelectItemContext) any

	// Visit a parse tree produced by SimpleSqlParser#expressionSelectItem.
	VisitExpressionSelectItem(ctx *ExpressionSelectItemContext) any

	// Visit a parse tree produced by SimpleSqlParser#expression.
	VisitExpression(ctx *ExpressionContext) any

	// Visit a parse tree produced by SimpleSqlParser#binaryBooleanExpression.
	VisitBinaryBooleanExpression(ctx *BinaryBooleanExpressionContext) any

	// Visit a parse tree produced by SimpleSqlParser#predicatedBooleanExpression.
	VisitPredicatedBooleanExpression(ctx *PredicatedBooleanExpressionContext) any

	// Visit a parse tree produced by SimpleSqlParser#unaryBooleanExpression.
	VisitUnaryBooleanExpression(ctx *UnaryBooleanExpressionContext) any

	// Visit a parse tree produced by SimpleSqlParser#castBooleanExpression.
	VisitCastBooleanExpression(ctx *CastBooleanExpressionContext) any

	// Visit a parse tree produced by SimpleSqlParser#cmpPredicate.
	VisitCmpPredicate(ctx *CmpPredicateContext) any

	// Visit a parse tree produced by SimpleSqlParser#isNullPredicate.
	VisitIsNullPredicate(ctx *IsNullPredicateContext) any

	// Visit a parse tree produced by SimpleSqlParser#inListPredicate.
	VisitInListPredicate(ctx *InListPredicateContext) any

	// Visit a parse tree produced by SimpleSqlParser#inSelectPredicate.
	VisitInSelectPredicate(ctx *InSelectPredicateContext) any

	// Visit a parse tree produced by SimpleSqlParser#likePredicate.
	VisitLikePredicate(ctx *LikePredicateContext) any

	// Visit a parse tree produced by SimpleSqlParser#primaryValueExpression.
	VisitPrimaryValueExpression(ctx *PrimaryValueExpressionContext) any

	// Visit a parse tree produced by SimpleSqlParser#unaryValueExpression.
	VisitUnaryValueExpression(ctx *UnaryValueExpressionContext) any

	// Visit a parse tree produced by SimpleSqlParser#arithValueExpression.
	VisitArithValueExpression(ctx *ArithValueExpressionContext) any

	// Visit a parse tree produced by SimpleSqlParser#functionCallExpression.
	VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) any

	// Visit a parse tree produced by SimpleSqlParser#starFunctionCallExpression.
	VisitStarFunctionCallExpression(ctx *StarFunctionCallExpressionContext) any

	// Visit a parse tree produced by SimpleSqlParser#caseExpression.
	VisitCaseExpression(ctx *CaseExpressionContext) any

	// Visit a parse tree produced by SimpleSqlParser#selectExpression.
	VisitSelectExpression(ctx *SelectExpressionContext) any

	// Visit a parse tree produced by SimpleSqlParser#parenExpression.
	VisitParenExpression(ctx *ParenExpressionContext) any

	// Visit a parse tree produced by SimpleSqlParser#simplePrimaryExpression.
	VisitSimplePrimaryExpression(ctx *SimplePrimaryExpressionContext) any

	// Visit a parse tree produced by SimpleSqlParser#simpleExpression.
	VisitSimpleExpression(ctx *SimpleExpressionContext) any

	// Visit a parse tree produced by SimpleSqlParser#caseItem.
	VisitCaseItem(ctx *CaseItemContext) any

	// Visit a parse tree produced by SimpleSqlParser#over.
	VisitOver(ctx *OverContext) any

	// Visit a parse tree produced by SimpleSqlParser#sortItem.
	VisitSortItem(ctx *SortItemContext) any

	// Visit a parse tree produced by SimpleSqlParser#aliasedRelation.
	VisitAliasedRelation(ctx *AliasedRelationContext) any

	// Visit a parse tree produced by SimpleSqlParser#joinRelation.
	VisitJoinRelation(ctx *JoinRelationContext) any

	// Visit a parse tree produced by SimpleSqlParser#selectRelation.
	VisitSelectRelation(ctx *SelectRelationContext) any

	// Visit a parse tree produced by SimpleSqlParser#tableRelation.
	VisitTableRelation(ctx *TableRelationContext) any

	// Visit a parse tree produced by SimpleSqlParser#parenRelation.
	VisitParenRelation(ctx *ParenRelationContext) any

	// Visit a parse tree produced by SimpleSqlParser#groupBy.
	VisitGroupBy(ctx *GroupByContext) any

	// Visit a parse tree produced by SimpleSqlParser#qualifiedName.
	VisitQualifiedName(ctx *QualifiedNameContext) any

	// Visit a parse tree produced by SimpleSqlParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) any

	// Visit a parse tree produced by SimpleSqlParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) any

	// Visit a parse tree produced by SimpleSqlParser#quotedIdentifier.
	VisitQuotedIdentifier(ctx *QuotedIdentifierContext) any

	// Visit a parse tree produced by SimpleSqlParser#integerNumber.
	VisitIntegerNumber(ctx *IntegerNumberContext) any

	// Visit a parse tree produced by SimpleSqlParser#decimalNumber.
	VisitDecimalNumber(ctx *DecimalNumberContext) any

	// Visit a parse tree produced by SimpleSqlParser#floatNumber.
	VisitFloatNumber(ctx *FloatNumberContext) any

	// Visit a parse tree produced by SimpleSqlParser#str.
	VisitStr(ctx *StrContext) any

	// Visit a parse tree produced by SimpleSqlParser#null.
	VisitNull(ctx *NullContext) any

	// Visit a parse tree produced by SimpleSqlParser#true.
	VisitTrue(ctx *TrueContext) any

	// Visit a parse tree produced by SimpleSqlParser#false.
	VisitFalse(ctx *FalseContext) any

	// Visit a parse tree produced by SimpleSqlParser#setQuantifier.
	VisitSetQuantifier(ctx *SetQuantifierContext) any

	// Visit a parse tree produced by SimpleSqlParser#joinType.
	VisitJoinType(ctx *JoinTypeContext) any

	// Visit a parse tree produced by SimpleSqlParser#cmpOp.
	VisitCmpOp(ctx *CmpOpContext) any

	// Visit a parse tree produced by SimpleSqlParser#arithOp.
	VisitArithOp(ctx *ArithOpContext) any

	// Visit a parse tree produced by SimpleSqlParser#unaryOp.
	VisitUnaryOp(ctx *UnaryOpContext) any

	// Visit a parse tree produced by SimpleSqlParser#unquotedIdentifier.
	VisitUnquotedIdentifier(ctx *UnquotedIdentifierContext) any
}
