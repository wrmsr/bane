// Code generated from Toml.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Toml

import antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

type BaseTomlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTomlVisitor) VisitDocument(ctx *DocumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitKeyValue(ctx *KeyValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitKey(ctx *KeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitSimpleKey(ctx *SimpleKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitUnquotedKey(ctx *UnquotedKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitQuotedKey(ctx *QuotedKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitDottedKey(ctx *DottedKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitStringValue(ctx *StringValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitInteger(ctx *IntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitFloatingPoint(ctx *FloatingPointContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitBoolean(ctx *BooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitDateTime(ctx *DateTimeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitArrayValues(ctx *ArrayValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitCommentOrNl(ctx *CommentOrNlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitTable(ctx *TableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitStandardTable(ctx *StandardTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitInlineTable(ctx *InlineTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitInlineTableKeyvals(ctx *InlineTableKeyvalsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitInlineTableKeyvalsNonEmpty(ctx *InlineTableKeyvalsNonEmptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTomlVisitor) VisitArrayTable(ctx *ArrayTableContext) interface{} {
	return v.VisitChildren(ctx)
}
