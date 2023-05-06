//go:build !nodev

// Code generated from Chat.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Chat

import antlr "github.com/wrmsr/bane/core/antlr/runtime"

type BaseChatVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseChatVisitor) VisitChat(ctx *ChatContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseChatVisitor) VisitLine(ctx *LineContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseChatVisitor) VisitName(ctx *NameContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseChatVisitor) VisitCommand(ctx *CommandContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseChatVisitor) VisitMessage(ctx *MessageContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseChatVisitor) VisitEmoticon(ctx *EmoticonContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseChatVisitor) VisitLink(ctx *LinkContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseChatVisitor) VisitColor(ctx *ColorContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseChatVisitor) VisitMention(ctx *MentionContext) any {
	return v.VisitChildren(ctx)
}
