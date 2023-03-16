//go:build !nodev

// Code generated from Chat.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Chat

import antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

type BaseChatVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseChatVisitor) VisitChat(ctx *ChatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseChatVisitor) VisitLine(ctx *LineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseChatVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseChatVisitor) VisitCommand(ctx *CommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseChatVisitor) VisitMessage(ctx *MessageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseChatVisitor) VisitEmoticon(ctx *EmoticonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseChatVisitor) VisitLink(ctx *LinkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseChatVisitor) VisitColor(ctx *ColorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseChatVisitor) VisitMention(ctx *MentionContext) interface{} {
	return v.VisitChildren(ctx)
}
