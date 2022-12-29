//go:build !nodev

// Code generated from Chat.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Chat

import antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

// ChatListener is a complete listener for a parse tree produced by ChatParser.
type ChatListener interface {
	antlr.ParseTreeListener

	// EnterChat is called when entering the chat production.
	EnterChat(c *ChatContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterMessage is called when entering the message production.
	EnterMessage(c *MessageContext)

	// EnterEmoticon is called when entering the emoticon production.
	EnterEmoticon(c *EmoticonContext)

	// EnterLink is called when entering the link production.
	EnterLink(c *LinkContext)

	// EnterColor is called when entering the color production.
	EnterColor(c *ColorContext)

	// EnterMention is called when entering the mention production.
	EnterMention(c *MentionContext)

	// ExitChat is called when exiting the chat production.
	ExitChat(c *ChatContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitMessage is called when exiting the message production.
	ExitMessage(c *MessageContext)

	// ExitEmoticon is called when exiting the emoticon production.
	ExitEmoticon(c *EmoticonContext)

	// ExitLink is called when exiting the link production.
	ExitLink(c *LinkContext)

	// ExitColor is called when exiting the color production.
	ExitColor(c *ColorContext)

	// ExitMention is called when exiting the mention production.
	ExitMention(c *MentionContext)
}
