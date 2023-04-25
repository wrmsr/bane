// Code generated from Toml.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Toml

import antlr "github.com/wrmsr/bane/core/antlr/runtime"

type BaseTomlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTomlVisitor) VisitDocument(ctx *DocumentContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitExpression(ctx *ExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitComment(ctx *CommentContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitKeyValue(ctx *KeyValueContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitKey(ctx *KeyContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitSimpleKey(ctx *SimpleKeyContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitUnquotedKey(ctx *UnquotedKeyContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitQuotedKey(ctx *QuotedKeyContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitDottedKey(ctx *DottedKeyContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitValue(ctx *ValueContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitStringValue(ctx *StringValueContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitInteger(ctx *IntegerContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitFloatingPoint(ctx *FloatingPointContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitBoolean(ctx *BooleanContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitDateTime(ctx *DateTimeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitArray(ctx *ArrayContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitArrayValues(ctx *ArrayValuesContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitCommentOrNl(ctx *CommentOrNlContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitTable(ctx *TableContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitStandardTable(ctx *StandardTableContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitInlineTable(ctx *InlineTableContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitInlineTableKeyvals(ctx *InlineTableKeyvalsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitInlineTableKeyvalsNonEmpty(ctx *InlineTableKeyvalsNonEmptyContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitArrayTable(ctx *ArrayTableContext) any {
	return v.VisitChildren(ctx)
}
