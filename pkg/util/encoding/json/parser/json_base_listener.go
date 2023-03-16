// Code generated from Json.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Json

import antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

// BaseJsonListener is a complete listener for a parse tree produced by JsonParser.
type BaseJsonListener struct{}

var _ JsonListener = &BaseJsonListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJsonListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJsonListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJsonListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJsonListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterJson is called when production json is entered.
func (s *BaseJsonListener) EnterJson(ctx *JsonContext) {}

// ExitJson is called when production json is exited.
func (s *BaseJsonListener) ExitJson(ctx *JsonContext) {}

// EnterObject is called when production object is entered.
func (s *BaseJsonListener) EnterObject(ctx *ObjectContext) {}

// ExitObject is called when production object is exited.
func (s *BaseJsonListener) ExitObject(ctx *ObjectContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseJsonListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseJsonListener) ExitPair(ctx *PairContext) {}

// EnterArray is called when production array is entered.
func (s *BaseJsonListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseJsonListener) ExitArray(ctx *ArrayContext) {}

// EnterValue is called when production value is entered.
func (s *BaseJsonListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseJsonListener) ExitValue(ctx *ValueContext) {}
