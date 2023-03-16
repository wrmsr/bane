//go:build !nodev

// Code generated from Chat.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Chat

import antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

// BaseChatListener is a complete listener for a parse tree produced by ChatParser.
type BaseChatListener struct{}

var _ ChatListener = &BaseChatListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseChatListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseChatListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseChatListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseChatListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterChat is called when production chat is entered.
func (s *BaseChatListener) EnterChat(ctx *ChatContext) {}

// ExitChat is called when production chat is exited.
func (s *BaseChatListener) ExitChat(ctx *ChatContext) {}

// EnterLine is called when production line is entered.
func (s *BaseChatListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaseChatListener) ExitLine(ctx *LineContext) {}

// EnterName is called when production name is entered.
func (s *BaseChatListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseChatListener) ExitName(ctx *NameContext) {}

// EnterCommand is called when production command is entered.
func (s *BaseChatListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BaseChatListener) ExitCommand(ctx *CommandContext) {}

// EnterMessage is called when production message is entered.
func (s *BaseChatListener) EnterMessage(ctx *MessageContext) {}

// ExitMessage is called when production message is exited.
func (s *BaseChatListener) ExitMessage(ctx *MessageContext) {}

// EnterEmoticon is called when production emoticon is entered.
func (s *BaseChatListener) EnterEmoticon(ctx *EmoticonContext) {}

// ExitEmoticon is called when production emoticon is exited.
func (s *BaseChatListener) ExitEmoticon(ctx *EmoticonContext) {}

// EnterLink is called when production link is entered.
func (s *BaseChatListener) EnterLink(ctx *LinkContext) {}

// ExitLink is called when production link is exited.
func (s *BaseChatListener) ExitLink(ctx *LinkContext) {}

// EnterColor is called when production color is entered.
func (s *BaseChatListener) EnterColor(ctx *ColorContext) {}

// ExitColor is called when production color is exited.
func (s *BaseChatListener) ExitColor(ctx *ColorContext) {}

// EnterMention is called when production mention is entered.
func (s *BaseChatListener) EnterMention(ctx *MentionContext) {}

// ExitMention is called when production mention is exited.
func (s *BaseChatListener) ExitMention(ctx *MentionContext) {}
