// Code generated from Chat.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Chat

import antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

// A complete Visitor for a parse tree produced by ChatParser.
type ChatVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ChatParser#chat.
	VisitChat(ctx *ChatContext) any

	// Visit a parse tree produced by ChatParser#line.
	VisitLine(ctx *LineContext) any

	// Visit a parse tree produced by ChatParser#name.
	VisitName(ctx *NameContext) any

	// Visit a parse tree produced by ChatParser#command.
	VisitCommand(ctx *CommandContext) any

	// Visit a parse tree produced by ChatParser#message.
	VisitMessage(ctx *MessageContext) any

	// Visit a parse tree produced by ChatParser#emoticon.
	VisitEmoticon(ctx *EmoticonContext) any

	// Visit a parse tree produced by ChatParser#link.
	VisitLink(ctx *LinkContext) any

	// Visit a parse tree produced by ChatParser#color.
	VisitColor(ctx *ColorContext) any

	// Visit a parse tree produced by ChatParser#mention.
	VisitMention(ctx *MentionContext) any
}
