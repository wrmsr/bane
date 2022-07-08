//go:build !nodev

// Code generated from Chat.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Chat

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ChatParser.
type ChatVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ChatParser#chat.
	VisitChat(ctx *ChatContext) interface{}

	// Visit a parse tree produced by ChatParser#line.
	VisitLine(ctx *LineContext) interface{}

	// Visit a parse tree produced by ChatParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by ChatParser#command.
	VisitCommand(ctx *CommandContext) interface{}

	// Visit a parse tree produced by ChatParser#message.
	VisitMessage(ctx *MessageContext) interface{}

	// Visit a parse tree produced by ChatParser#emoticon.
	VisitEmoticon(ctx *EmoticonContext) interface{}

	// Visit a parse tree produced by ChatParser#link.
	VisitLink(ctx *LinkContext) interface{}

	// Visit a parse tree produced by ChatParser#color.
	VisitColor(ctx *ColorContext) interface{}

	// Visit a parse tree produced by ChatParser#mention.
	VisitMention(ctx *MentionContext) interface{}
}
