// Code generated from Toml.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Toml

import antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

// TomlListener is a complete listener for a parse tree produced by TomlParser.
type TomlListener interface {
	antlr.ParseTreeListener

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterKeyValue is called when entering the keyValue production.
	EnterKeyValue(c *KeyValueContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterSimpleKey is called when entering the simpleKey production.
	EnterSimpleKey(c *SimpleKeyContext)

	// EnterUnquotedKey is called when entering the unquotedKey production.
	EnterUnquotedKey(c *UnquotedKeyContext)

	// EnterQuotedKey is called when entering the quotedKey production.
	EnterQuotedKey(c *QuotedKeyContext)

	// EnterDottedKey is called when entering the dottedKey production.
	EnterDottedKey(c *DottedKeyContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterStringValue is called when entering the stringValue production.
	EnterStringValue(c *StringValueContext)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// EnterFloatingPoint is called when entering the floatingPoint production.
	EnterFloatingPoint(c *FloatingPointContext)

	// EnterBoolean is called when entering the boolean production.
	EnterBoolean(c *BooleanContext)

	// EnterDateTime is called when entering the dateTime production.
	EnterDateTime(c *DateTimeContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterArrayValues is called when entering the arrayValues production.
	EnterArrayValues(c *ArrayValuesContext)

	// EnterCommentOrNl is called when entering the commentOrNl production.
	EnterCommentOrNl(c *CommentOrNlContext)

	// EnterTable is called when entering the table production.
	EnterTable(c *TableContext)

	// EnterStandardTable is called when entering the standardTable production.
	EnterStandardTable(c *StandardTableContext)

	// EnterInlineTable is called when entering the inlineTable production.
	EnterInlineTable(c *InlineTableContext)

	// EnterInlineTableKeyvals is called when entering the inlineTableKeyvals production.
	EnterInlineTableKeyvals(c *InlineTableKeyvalsContext)

	// EnterInlineTableKeyvalsNonEmpty is called when entering the inlineTableKeyvalsNonEmpty production.
	EnterInlineTableKeyvalsNonEmpty(c *InlineTableKeyvalsNonEmptyContext)

	// EnterArrayTable is called when entering the arrayTable production.
	EnterArrayTable(c *ArrayTableContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitKeyValue is called when exiting the keyValue production.
	ExitKeyValue(c *KeyValueContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitSimpleKey is called when exiting the simpleKey production.
	ExitSimpleKey(c *SimpleKeyContext)

	// ExitUnquotedKey is called when exiting the unquotedKey production.
	ExitUnquotedKey(c *UnquotedKeyContext)

	// ExitQuotedKey is called when exiting the quotedKey production.
	ExitQuotedKey(c *QuotedKeyContext)

	// ExitDottedKey is called when exiting the dottedKey production.
	ExitDottedKey(c *DottedKeyContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitStringValue is called when exiting the stringValue production.
	ExitStringValue(c *StringValueContext)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)

	// ExitFloatingPoint is called when exiting the floatingPoint production.
	ExitFloatingPoint(c *FloatingPointContext)

	// ExitBoolean is called when exiting the boolean production.
	ExitBoolean(c *BooleanContext)

	// ExitDateTime is called when exiting the dateTime production.
	ExitDateTime(c *DateTimeContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitArrayValues is called when exiting the arrayValues production.
	ExitArrayValues(c *ArrayValuesContext)

	// ExitCommentOrNl is called when exiting the commentOrNl production.
	ExitCommentOrNl(c *CommentOrNlContext)

	// ExitTable is called when exiting the table production.
	ExitTable(c *TableContext)

	// ExitStandardTable is called when exiting the standardTable production.
	ExitStandardTable(c *StandardTableContext)

	// ExitInlineTable is called when exiting the inlineTable production.
	ExitInlineTable(c *InlineTableContext)

	// ExitInlineTableKeyvals is called when exiting the inlineTableKeyvals production.
	ExitInlineTableKeyvals(c *InlineTableKeyvalsContext)

	// ExitInlineTableKeyvalsNonEmpty is called when exiting the inlineTableKeyvalsNonEmpty production.
	ExitInlineTableKeyvalsNonEmpty(c *InlineTableKeyvalsNonEmptyContext)

	// ExitArrayTable is called when exiting the arrayTable production.
	ExitArrayTable(c *ArrayTableContext)
}
