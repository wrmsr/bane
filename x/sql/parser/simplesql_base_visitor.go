// Code generated from SimpleSql.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // SimpleSql

import antlr "github.com/wrmsr/bane/core/antlr/runtime"

type BaseSimpleSqlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSimpleSqlVisitor) VisitSingleStatement(ctx *SingleStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitSelect(ctx *SelectContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitCteSelect(ctx *CteSelectContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitCte(ctx *CteContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitUnionSelect(ctx *UnionSelectContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitUnionItem(ctx *UnionItemContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitPrimarySelect(ctx *PrimarySelectContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitAllSelectItem(ctx *AllSelectItemContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitExpressionSelectItem(ctx *ExpressionSelectItemContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitExpression(ctx *ExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitBinaryBooleanExpression(ctx *BinaryBooleanExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitPredicatedBooleanExpression(ctx *PredicatedBooleanExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitUnaryBooleanExpression(ctx *UnaryBooleanExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitCastBooleanExpression(ctx *CastBooleanExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitCmpPredicate(ctx *CmpPredicateContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitIsNullPredicate(ctx *IsNullPredicateContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitInListPredicate(ctx *InListPredicateContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitInSelectPredicate(ctx *InSelectPredicateContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitLikePredicate(ctx *LikePredicateContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitPrimaryValueExpression(ctx *PrimaryValueExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitUnaryValueExpression(ctx *UnaryValueExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitArithValueExpression(ctx *ArithValueExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitStarFunctionCallExpression(ctx *StarFunctionCallExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitCaseExpression(ctx *CaseExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitSelectExpression(ctx *SelectExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitParenExpression(ctx *ParenExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitSimplePrimaryExpression(ctx *SimplePrimaryExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitSimpleExpression(ctx *SimpleExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitCaseItem(ctx *CaseItemContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitOver(ctx *OverContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitSortItem(ctx *SortItemContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitAliasedRelation(ctx *AliasedRelationContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitJoinRelation(ctx *JoinRelationContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitSelectRelation(ctx *SelectRelationContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitTableRelation(ctx *TableRelationContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitParenRelation(ctx *ParenRelationContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitGroupBy(ctx *GroupByContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitQualifiedName(ctx *QualifiedNameContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitIdentifierList(ctx *IdentifierListContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitIdentifier(ctx *IdentifierContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitQuotedIdentifier(ctx *QuotedIdentifierContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitIntegerNumber(ctx *IntegerNumberContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitDecimalNumber(ctx *DecimalNumberContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitFloatNumber(ctx *FloatNumberContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitStr(ctx *StrContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitNull(ctx *NullContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitTrue(ctx *TrueContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitFalse(ctx *FalseContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitSetQuantifier(ctx *SetQuantifierContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitJoinType(ctx *JoinTypeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitCmpOp(ctx *CmpOpContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitArithOp(ctx *ArithOpContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitUnaryOp(ctx *UnaryOpContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleSqlVisitor) VisitUnquotedIdentifier(ctx *UnquotedIdentifierContext) any {
	return v.VisitChildren(ctx)
}
