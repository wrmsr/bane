// Code generated from Toml.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Toml

import antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

// A complete Visitor for a parse tree produced by TomlParser.
type TomlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TomlParser#document.
	VisitDocument(ctx *DocumentContext) interface{}

	// Visit a parse tree produced by TomlParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by TomlParser#comment.
	VisitComment(ctx *CommentContext) interface{}

	// Visit a parse tree produced by TomlParser#keyValue.
	VisitKeyValue(ctx *KeyValueContext) interface{}

	// Visit a parse tree produced by TomlParser#key.
	VisitKey(ctx *KeyContext) interface{}

	// Visit a parse tree produced by TomlParser#simpleKey.
	VisitSimpleKey(ctx *SimpleKeyContext) interface{}

	// Visit a parse tree produced by TomlParser#unquotedKey.
	VisitUnquotedKey(ctx *UnquotedKeyContext) interface{}

	// Visit a parse tree produced by TomlParser#quotedKey.
	VisitQuotedKey(ctx *QuotedKeyContext) interface{}

	// Visit a parse tree produced by TomlParser#dottedKey.
	VisitDottedKey(ctx *DottedKeyContext) interface{}

	// Visit a parse tree produced by TomlParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by TomlParser#stringValue.
	VisitStringValue(ctx *StringValueContext) interface{}

	// Visit a parse tree produced by TomlParser#integer.
	VisitInteger(ctx *IntegerContext) interface{}

	// Visit a parse tree produced by TomlParser#floatingPoint.
	VisitFloatingPoint(ctx *FloatingPointContext) interface{}

	// Visit a parse tree produced by TomlParser#boolean.
	VisitBoolean(ctx *BooleanContext) interface{}

	// Visit a parse tree produced by TomlParser#dateTime.
	VisitDateTime(ctx *DateTimeContext) interface{}

	// Visit a parse tree produced by TomlParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by TomlParser#arrayValues.
	VisitArrayValues(ctx *ArrayValuesContext) interface{}

	// Visit a parse tree produced by TomlParser#commentOrNl.
	VisitCommentOrNl(ctx *CommentOrNlContext) interface{}

	// Visit a parse tree produced by TomlParser#table.
	VisitTable(ctx *TableContext) interface{}

	// Visit a parse tree produced by TomlParser#standardTable.
	VisitStandardTable(ctx *StandardTableContext) interface{}

	// Visit a parse tree produced by TomlParser#inlineTable.
	VisitInlineTable(ctx *InlineTableContext) interface{}

	// Visit a parse tree produced by TomlParser#inlineTableKeyvals.
	VisitInlineTableKeyvals(ctx *InlineTableKeyvalsContext) interface{}

	// Visit a parse tree produced by TomlParser#inlineTableKeyvalsNonEmpty.
	VisitInlineTableKeyvalsNonEmpty(ctx *InlineTableKeyvalsNonEmptyContext) interface{}

	// Visit a parse tree produced by TomlParser#arrayTable.
	VisitArrayTable(ctx *ArrayTableContext) interface{}
}
