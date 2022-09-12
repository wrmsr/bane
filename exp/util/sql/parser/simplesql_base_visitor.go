// Code generated from SimpleSql.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SimpleSql

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSimpleSqlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSimpleSqlVisitor) VisitSingleStatement(ctx *SingleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitSelect(ctx *SelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitCteSelect(ctx *CteSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitCte(ctx *CteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitUnionSelect(ctx *UnionSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitUnionItem(ctx *UnionItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitPrimarySelect(ctx *PrimarySelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitAllSelectItem(ctx *AllSelectItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitExpressionSelectItem(ctx *ExpressionSelectItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitBinaryBooleanExpression(ctx *BinaryBooleanExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitPredicatedBooleanExpression(ctx *PredicatedBooleanExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitUnaryBooleanExpression(ctx *UnaryBooleanExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitCastBooleanExpression(ctx *CastBooleanExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitCmpPredicate(ctx *CmpPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitIsNullPredicate(ctx *IsNullPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitInListPredicate(ctx *InListPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitInSelectPredicate(ctx *InSelectPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitLikePredicate(ctx *LikePredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitPrimaryValueExpression(ctx *PrimaryValueExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitUnaryValueExpression(ctx *UnaryValueExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitArithValueExpression(ctx *ArithValueExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitStarFunctionCallExpression(ctx *StarFunctionCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitCaseExpression(ctx *CaseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitSelectExpression(ctx *SelectExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitParenExpression(ctx *ParenExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitSimplePrimaryExpression(ctx *SimplePrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitSimpleExpression(ctx *SimpleExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitCaseItem(ctx *CaseItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitOver(ctx *OverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitSortItem(ctx *SortItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitAliasedRelation(ctx *AliasedRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitJoinRelation(ctx *JoinRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitSelectRelation(ctx *SelectRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitTableRelation(ctx *TableRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitParenRelation(ctx *ParenRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitGroupBy(ctx *GroupByContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitQualifiedName(ctx *QualifiedNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitQuotedIdentifier(ctx *QuotedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitIntegerNumber(ctx *IntegerNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitDecimalNumber(ctx *DecimalNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitFloatNumber(ctx *FloatNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitStr(ctx *StrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitNull(ctx *NullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitTrue(ctx *TrueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitFalse(ctx *FalseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitSetQuantifier(ctx *SetQuantifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitJoinType(ctx *JoinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitCmpOp(ctx *CmpOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitArithOp(ctx *ArithOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitUnaryOp(ctx *UnaryOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitUnquotedIdentifier(ctx *UnquotedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}
