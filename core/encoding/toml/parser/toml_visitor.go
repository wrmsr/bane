// Code generated from Toml.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Toml

import antlr "github.com/wrmsr/bane/core/antlr/runtime"

// A complete Visitor for a parse tree produced by TomlParser.
type TomlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TomlParser#document.
	VisitDocument(ctx *DocumentContext) any

	// Visit a parse tree produced by TomlParser#expression.
	VisitExpression(ctx *ExpressionContext) any

	// Visit a parse tree produced by TomlParser#comment.
	VisitComment(ctx *CommentContext) any

	// Visit a parse tree produced by TomlParser#keyValue.
	VisitKeyValue(ctx *KeyValueContext) any

	// Visit a parse tree produced by TomlParser#key.
	VisitKey(ctx *KeyContext) any

	// Visit a parse tree produced by TomlParser#simpleKey.
	VisitSimpleKey(ctx *SimpleKeyContext) any

	// Visit a parse tree produced by TomlParser#unquotedKey.
	VisitUnquotedKey(ctx *UnquotedKeyContext) any

	// Visit a parse tree produced by TomlParser#quotedKey.
	VisitQuotedKey(ctx *QuotedKeyContext) any

	// Visit a parse tree produced by TomlParser#dottedKey.
	VisitDottedKey(ctx *DottedKeyContext) any

	// Visit a parse tree produced by TomlParser#value.
	VisitValue(ctx *ValueContext) any

	// Visit a parse tree produced by TomlParser#stringValue.
	VisitStringValue(ctx *StringValueContext) any

	// Visit a parse tree produced by TomlParser#integer.
	VisitInteger(ctx *IntegerContext) any

	// Visit a parse tree produced by TomlParser#floatingPoint.
	VisitFloatingPoint(ctx *FloatingPointContext) any

	// Visit a parse tree produced by TomlParser#boolean.
	VisitBoolean(ctx *BooleanContext) any

	// Visit a parse tree produced by TomlParser#dateTime.
	VisitDateTime(ctx *DateTimeContext) any

	// Visit a parse tree produced by TomlParser#array.
	VisitArray(ctx *ArrayContext) any

	// Visit a parse tree produced by TomlParser#arrayValues.
	VisitArrayValues(ctx *ArrayValuesContext) any

	// Visit a parse tree produced by TomlParser#commentOrNl.
	VisitCommentOrNl(ctx *CommentOrNlContext) any

	// Visit a parse tree produced by TomlParser#table.
	VisitTable(ctx *TableContext) any

	// Visit a parse tree produced by TomlParser#standardTable.
	VisitStandardTable(ctx *StandardTableContext) any

	// Visit a parse tree produced by TomlParser#inlineTable.
	VisitInlineTable(ctx *InlineTableContext) any

	// Visit a parse tree produced by TomlParser#inlineTableKeyvals.
	VisitInlineTableKeyvals(ctx *InlineTableKeyvalsContext) any

	// Visit a parse tree produced by TomlParser#inlineTableKeyvalsNonEmpty.
	VisitInlineTableKeyvalsNonEmpty(ctx *InlineTableKeyvalsNonEmptyContext) any

	// Visit a parse tree produced by TomlParser#arrayTable.
	VisitArrayTable(ctx *ArrayTableContext) any
}
