// Code generated from Toml.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Toml

import antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

// BaseTomlListener is a complete listener for a parse tree produced by TomlParser.
type BaseTomlListener struct{}

var _ TomlListener = &BaseTomlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTomlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTomlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTomlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTomlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocument is called when production document is entered.
func (s *BaseTomlListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BaseTomlListener) ExitDocument(ctx *DocumentContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTomlListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTomlListener) ExitExpression(ctx *ExpressionContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseTomlListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseTomlListener) ExitComment(ctx *CommentContext) {}

// EnterKeyValue is called when production keyValue is entered.
func (s *BaseTomlListener) EnterKeyValue(ctx *KeyValueContext) {}

// ExitKeyValue is called when production keyValue is exited.
func (s *BaseTomlListener) ExitKeyValue(ctx *KeyValueContext) {}

// EnterKey is called when production key is entered.
func (s *BaseTomlListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BaseTomlListener) ExitKey(ctx *KeyContext) {}

// EnterSimpleKey is called when production simpleKey is entered.
func (s *BaseTomlListener) EnterSimpleKey(ctx *SimpleKeyContext) {}

// ExitSimpleKey is called when production simpleKey is exited.
func (s *BaseTomlListener) ExitSimpleKey(ctx *SimpleKeyContext) {}

// EnterUnquotedKey is called when production unquotedKey is entered.
func (s *BaseTomlListener) EnterUnquotedKey(ctx *UnquotedKeyContext) {}

// ExitUnquotedKey is called when production unquotedKey is exited.
func (s *BaseTomlListener) ExitUnquotedKey(ctx *UnquotedKeyContext) {}

// EnterQuotedKey is called when production quotedKey is entered.
func (s *BaseTomlListener) EnterQuotedKey(ctx *QuotedKeyContext) {}

// ExitQuotedKey is called when production quotedKey is exited.
func (s *BaseTomlListener) ExitQuotedKey(ctx *QuotedKeyContext) {}

// EnterDottedKey is called when production dottedKey is entered.
func (s *BaseTomlListener) EnterDottedKey(ctx *DottedKeyContext) {}

// ExitDottedKey is called when production dottedKey is exited.
func (s *BaseTomlListener) ExitDottedKey(ctx *DottedKeyContext) {}

// EnterValue is called when production value is entered.
func (s *BaseTomlListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseTomlListener) ExitValue(ctx *ValueContext) {}

// EnterStringValue is called when production stringValue is entered.
func (s *BaseTomlListener) EnterStringValue(ctx *StringValueContext) {}

// ExitStringValue is called when production stringValue is exited.
func (s *BaseTomlListener) ExitStringValue(ctx *StringValueContext) {}

// EnterInteger is called when production integer is entered.
func (s *BaseTomlListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BaseTomlListener) ExitInteger(ctx *IntegerContext) {}

// EnterFloatingPoint is called when production floatingPoint is entered.
func (s *BaseTomlListener) EnterFloatingPoint(ctx *FloatingPointContext) {}

// ExitFloatingPoint is called when production floatingPoint is exited.
func (s *BaseTomlListener) ExitFloatingPoint(ctx *FloatingPointContext) {}

// EnterBoolean is called when production boolean is entered.
func (s *BaseTomlListener) EnterBoolean(ctx *BooleanContext) {}

// ExitBoolean is called when production boolean is exited.
func (s *BaseTomlListener) ExitBoolean(ctx *BooleanContext) {}

// EnterDateTime is called when production dateTime is entered.
func (s *BaseTomlListener) EnterDateTime(ctx *DateTimeContext) {}

// ExitDateTime is called when production dateTime is exited.
func (s *BaseTomlListener) ExitDateTime(ctx *DateTimeContext) {}

// EnterArray is called when production array is entered.
func (s *BaseTomlListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseTomlListener) ExitArray(ctx *ArrayContext) {}

// EnterArrayValues is called when production arrayValues is entered.
func (s *BaseTomlListener) EnterArrayValues(ctx *ArrayValuesContext) {}

// ExitArrayValues is called when production arrayValues is exited.
func (s *BaseTomlListener) ExitArrayValues(ctx *ArrayValuesContext) {}

// EnterCommentOrNl is called when production commentOrNl is entered.
func (s *BaseTomlListener) EnterCommentOrNl(ctx *CommentOrNlContext) {}

// ExitCommentOrNl is called when production commentOrNl is exited.
func (s *BaseTomlListener) ExitCommentOrNl(ctx *CommentOrNlContext) {}

// EnterTable is called when production table is entered.
func (s *BaseTomlListener) EnterTable(ctx *TableContext) {}

// ExitTable is called when production table is exited.
func (s *BaseTomlListener) ExitTable(ctx *TableContext) {}

// EnterStandardTable is called when production standardTable is entered.
func (s *BaseTomlListener) EnterStandardTable(ctx *StandardTableContext) {}

// ExitStandardTable is called when production standardTable is exited.
func (s *BaseTomlListener) ExitStandardTable(ctx *StandardTableContext) {}

// EnterInlineTable is called when production inlineTable is entered.
func (s *BaseTomlListener) EnterInlineTable(ctx *InlineTableContext) {}

// ExitInlineTable is called when production inlineTable is exited.
func (s *BaseTomlListener) ExitInlineTable(ctx *InlineTableContext) {}

// EnterInlineTableKeyvals is called when production inlineTableKeyvals is entered.
func (s *BaseTomlListener) EnterInlineTableKeyvals(ctx *InlineTableKeyvalsContext) {}

// ExitInlineTableKeyvals is called when production inlineTableKeyvals is exited.
func (s *BaseTomlListener) ExitInlineTableKeyvals(ctx *InlineTableKeyvalsContext) {}

// EnterInlineTableKeyvalsNonEmpty is called when production inlineTableKeyvalsNonEmpty is entered.
func (s *BaseTomlListener) EnterInlineTableKeyvalsNonEmpty(ctx *InlineTableKeyvalsNonEmptyContext) {}

// ExitInlineTableKeyvalsNonEmpty is called when production inlineTableKeyvalsNonEmpty is exited.
func (s *BaseTomlListener) ExitInlineTableKeyvalsNonEmpty(ctx *InlineTableKeyvalsNonEmptyContext) {}

// EnterArrayTable is called when production arrayTable is entered.
func (s *BaseTomlListener) EnterArrayTable(ctx *ArrayTableContext) {}

// ExitArrayTable is called when production arrayTable is exited.
func (s *BaseTomlListener) ExitArrayTable(ctx *ArrayTableContext) {}
