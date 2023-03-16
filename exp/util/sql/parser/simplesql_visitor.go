// Code generated from SimpleSql.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // SimpleSql

import antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

// A complete Visitor for a parse tree produced by SimpleSqlParser.
type SimpleSqlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SimpleSqlParser#singleStatement.
	VisitSingleStatement(ctx *SingleStatementContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#select.
	VisitSelect(ctx *SelectContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#cteSelect.
	VisitCteSelect(ctx *CteSelectContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#cte.
	VisitCte(ctx *CteContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#unionSelect.
	VisitUnionSelect(ctx *UnionSelectContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#unionItem.
	VisitUnionItem(ctx *UnionItemContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#primarySelect.
	VisitPrimarySelect(ctx *PrimarySelectContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#allSelectItem.
	VisitAllSelectItem(ctx *AllSelectItemContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#expressionSelectItem.
	VisitExpressionSelectItem(ctx *ExpressionSelectItemContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#binaryBooleanExpression.
	VisitBinaryBooleanExpression(ctx *BinaryBooleanExpressionContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#predicatedBooleanExpression.
	VisitPredicatedBooleanExpression(ctx *PredicatedBooleanExpressionContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#unaryBooleanExpression.
	VisitUnaryBooleanExpression(ctx *UnaryBooleanExpressionContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#castBooleanExpression.
	VisitCastBooleanExpression(ctx *CastBooleanExpressionContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#cmpPredicate.
	VisitCmpPredicate(ctx *CmpPredicateContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#isNullPredicate.
	VisitIsNullPredicate(ctx *IsNullPredicateContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#inListPredicate.
	VisitInListPredicate(ctx *InListPredicateContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#inSelectPredicate.
	VisitInSelectPredicate(ctx *InSelectPredicateContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#likePredicate.
	VisitLikePredicate(ctx *LikePredicateContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#primaryValueExpression.
	VisitPrimaryValueExpression(ctx *PrimaryValueExpressionContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#unaryValueExpression.
	VisitUnaryValueExpression(ctx *UnaryValueExpressionContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#arithValueExpression.
	VisitArithValueExpression(ctx *ArithValueExpressionContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#functionCallExpression.
	VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#starFunctionCallExpression.
	VisitStarFunctionCallExpression(ctx *StarFunctionCallExpressionContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#caseExpression.
	VisitCaseExpression(ctx *CaseExpressionContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#selectExpression.
	VisitSelectExpression(ctx *SelectExpressionContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#parenExpression.
	VisitParenExpression(ctx *ParenExpressionContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#simplePrimaryExpression.
	VisitSimplePrimaryExpression(ctx *SimplePrimaryExpressionContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#simpleExpression.
	VisitSimpleExpression(ctx *SimpleExpressionContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#caseItem.
	VisitCaseItem(ctx *CaseItemContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#over.
	VisitOver(ctx *OverContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#sortItem.
	VisitSortItem(ctx *SortItemContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#aliasedRelation.
	VisitAliasedRelation(ctx *AliasedRelationContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#joinRelation.
	VisitJoinRelation(ctx *JoinRelationContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#selectRelation.
	VisitSelectRelation(ctx *SelectRelationContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#tableRelation.
	VisitTableRelation(ctx *TableRelationContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#parenRelation.
	VisitParenRelation(ctx *ParenRelationContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#groupBy.
	VisitGroupBy(ctx *GroupByContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#qualifiedName.
	VisitQualifiedName(ctx *QualifiedNameContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#quotedIdentifier.
	VisitQuotedIdentifier(ctx *QuotedIdentifierContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#integerNumber.
	VisitIntegerNumber(ctx *IntegerNumberContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#decimalNumber.
	VisitDecimalNumber(ctx *DecimalNumberContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#floatNumber.
	VisitFloatNumber(ctx *FloatNumberContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#str.
	VisitStr(ctx *StrContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#null.
	VisitNull(ctx *NullContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#true.
	VisitTrue(ctx *TrueContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#false.
	VisitFalse(ctx *FalseContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#setQuantifier.
	VisitSetQuantifier(ctx *SetQuantifierContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#joinType.
	VisitJoinType(ctx *JoinTypeContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#cmpOp.
	VisitCmpOp(ctx *CmpOpContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#arithOp.
	VisitArithOp(ctx *ArithOpContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#unaryOp.
	VisitUnaryOp(ctx *UnaryOpContext) interface{}

	// Visit a parse tree produced by SimpleSqlParser#unquotedIdentifier.
	VisitUnquotedIdentifier(ctx *UnquotedIdentifierContext) interface{}
}
