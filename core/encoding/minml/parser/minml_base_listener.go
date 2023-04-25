// Code generated from Minml.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Minml

import antlr "github.com/wrmsr/bane/core/antlr/runtime"

// BaseMinmlListener is a complete listener for a parse tree produced by MinmlParser.
type BaseMinmlListener struct{}

var _ MinmlListener = &BaseMinmlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMinmlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMinmlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMinmlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMinmlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseMinmlListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseMinmlListener) ExitRoot(ctx *RootContext) {}

// EnterValue is called when production value is entered.
func (s *BaseMinmlListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseMinmlListener) ExitValue(ctx *ValueContext) {}

// EnterObj is called when production obj is entered.
func (s *BaseMinmlListener) EnterObj(ctx *ObjContext) {}

// ExitObj is called when production obj is exited.
func (s *BaseMinmlListener) ExitObj(ctx *ObjContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseMinmlListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseMinmlListener) ExitPair(ctx *PairContext) {}

// EnterArray is called when production array is entered.
func (s *BaseMinmlListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseMinmlListener) ExitArray(ctx *ArrayContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseMinmlListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseMinmlListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterStringValue is called when production stringValue is entered.
func (s *BaseMinmlListener) EnterStringValue(ctx *StringValueContext) {}

// ExitStringValue is called when production stringValue is exited.
func (s *BaseMinmlListener) ExitStringValue(ctx *StringValueContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseMinmlListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseMinmlListener) ExitNumber(ctx *NumberContext) {}

// EnterTrue is called when production true is entered.
func (s *BaseMinmlListener) EnterTrue(ctx *TrueContext) {}

// ExitTrue is called when production true is exited.
func (s *BaseMinmlListener) ExitTrue(ctx *TrueContext) {}

// EnterFalse is called when production false is entered.
func (s *BaseMinmlListener) EnterFalse(ctx *FalseContext) {}

// ExitFalse is called when production false is exited.
func (s *BaseMinmlListener) ExitFalse(ctx *FalseContext) {}

// EnterNull is called when production null is entered.
func (s *BaseMinmlListener) EnterNull(ctx *NullContext) {}

// ExitNull is called when production null is exited.
func (s *BaseMinmlListener) ExitNull(ctx *NullContext) {}
