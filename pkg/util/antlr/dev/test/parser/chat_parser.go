//go:build !nodev

// Code generated from Chat.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Chat

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ChatParser struct {
	*antlr.BaseParser
}

var chatParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func chatParserInit() {
	staticData := &chatParserStaticData
	staticData.literalNames = []string{
		"", "':'", "'-'", "')'", "'('", "'/'", "'@'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "SAYS", "SHOUTS", "WORD", "WHITESPACE",
		"NEWLINE", "TEXT", "BLOCK_COMMENT", "COMMENT",
	}
	staticData.ruleNames = []string{
		"chat", "line", "name", "command", "message", "emoticon", "link", "color",
		"mention",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 14, 75, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 4, 0, 20, 8, 0,
		11, 0, 12, 0, 21, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 30, 8, 1, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 4, 4, 47, 8, 4, 11, 4, 12, 4, 48, 1, 5, 1, 5, 3, 5, 53,
		8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 58, 8, 5, 1, 5, 3, 5, 61, 8, 5, 1, 6, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 0,
		0, 9, 0, 2, 4, 6, 8, 10, 12, 14, 16, 0, 1, 1, 0, 7, 8, 76, 0, 19, 1, 0,
		0, 0, 2, 29, 1, 0, 0, 0, 4, 33, 1, 0, 0, 0, 6, 36, 1, 0, 0, 0, 8, 46, 1,
		0, 0, 0, 10, 60, 1, 0, 0, 0, 12, 62, 1, 0, 0, 0, 14, 65, 1, 0, 0, 0, 16,
		71, 1, 0, 0, 0, 18, 20, 3, 2, 1, 0, 19, 18, 1, 0, 0, 0, 20, 21, 1, 0, 0,
		0, 21, 19, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 24,
		5, 0, 0, 1, 24, 1, 1, 0, 0, 0, 25, 26, 3, 4, 2, 0, 26, 27, 3, 6, 3, 0,
		27, 28, 3, 8, 4, 0, 28, 30, 1, 0, 0, 0, 29, 25, 1, 0, 0, 0, 29, 30, 1,
		0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 32, 5, 11, 0, 0, 32, 3, 1, 0, 0, 0, 33,
		34, 5, 9, 0, 0, 34, 35, 5, 10, 0, 0, 35, 5, 1, 0, 0, 0, 36, 37, 7, 0, 0,
		0, 37, 38, 5, 1, 0, 0, 38, 39, 5, 10, 0, 0, 39, 7, 1, 0, 0, 0, 40, 47,
		3, 10, 5, 0, 41, 47, 3, 12, 6, 0, 42, 47, 3, 14, 7, 0, 43, 47, 3, 16, 8,
		0, 44, 47, 5, 9, 0, 0, 45, 47, 5, 10, 0, 0, 46, 40, 1, 0, 0, 0, 46, 41,
		1, 0, 0, 0, 46, 42, 1, 0, 0, 0, 46, 43, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0,
		46, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 48, 49, 1,
		0, 0, 0, 49, 9, 1, 0, 0, 0, 50, 52, 5, 1, 0, 0, 51, 53, 5, 2, 0, 0, 52,
		51, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 61, 5, 3, 0,
		0, 55, 57, 5, 1, 0, 0, 56, 58, 5, 2, 0, 0, 57, 56, 1, 0, 0, 0, 57, 58,
		1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 61, 5, 4, 0, 0, 60, 50, 1, 0, 0, 0,
		60, 55, 1, 0, 0, 0, 61, 11, 1, 0, 0, 0, 62, 63, 5, 12, 0, 0, 63, 64, 5,
		12, 0, 0, 64, 13, 1, 0, 0, 0, 65, 66, 5, 5, 0, 0, 66, 67, 5, 9, 0, 0, 67,
		68, 5, 5, 0, 0, 68, 69, 3, 8, 4, 0, 69, 70, 5, 5, 0, 0, 70, 15, 1, 0, 0,
		0, 71, 72, 5, 6, 0, 0, 72, 73, 5, 9, 0, 0, 73, 17, 1, 0, 0, 0, 7, 21, 29,
		46, 48, 52, 57, 60,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ChatParserInit initializes any static state used to implement ChatParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewChatParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ChatParserInit() {
	staticData := &chatParserStaticData
	staticData.once.Do(chatParserInit)
}

// NewChatParser produces a new parser instance for the optional input antlr.TokenStream.
func NewChatParser(input antlr.TokenStream) *ChatParser {
	ChatParserInit()
	this := new(ChatParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &chatParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Chat.g4"

	return this
}

// ChatParser tokens.
const (
	ChatParserEOF           = antlr.TokenEOF
	ChatParserT__0          = 1
	ChatParserT__1          = 2
	ChatParserT__2          = 3
	ChatParserT__3          = 4
	ChatParserT__4          = 5
	ChatParserT__5          = 6
	ChatParserSAYS          = 7
	ChatParserSHOUTS        = 8
	ChatParserWORD          = 9
	ChatParserWHITESPACE    = 10
	ChatParserNEWLINE       = 11
	ChatParserTEXT          = 12
	ChatParserBLOCK_COMMENT = 13
	ChatParserCOMMENT       = 14
)

// ChatParser rules.
const (
	ChatParserRULE_chat     = 0
	ChatParserRULE_line     = 1
	ChatParserRULE_name     = 2
	ChatParserRULE_command  = 3
	ChatParserRULE_message  = 4
	ChatParserRULE_emoticon = 5
	ChatParserRULE_link     = 6
	ChatParserRULE_color    = 7
	ChatParserRULE_mention  = 8
)

// IChatContext is an interface to support dynamic dispatch.
type IChatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChatContext differentiates from other interfaces.
	IsChatContext()
}

type ChatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChatContext() *ChatContext {
	var p = new(ChatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChatParserRULE_chat
	return p
}

func (*ChatContext) IsChatContext() {}

func NewChatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChatContext {
	var p = new(ChatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChatParserRULE_chat

	return p
}

func (s *ChatContext) GetParser() antlr.Parser { return s.parser }

func (s *ChatContext) EOF() antlr.TerminalNode {
	return s.GetToken(ChatParserEOF, 0)
}

func (s *ChatContext) AllLine() []ILineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILineContext); ok {
			len++
		}
	}

	tst := make([]ILineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILineContext); ok {
			tst[i] = t.(ILineContext)
			i++
		}
	}

	return tst
}

func (s *ChatContext) Line(i int) ILineContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *ChatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.EnterChat(s)
	}
}

func (s *ChatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.ExitChat(s)
	}
}

func (s *ChatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ChatVisitor:
		return t.VisitChat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ChatParser) Chat() (localctx IChatContext) {
	this := p
	_ = this

	localctx = NewChatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ChatParserRULE_chat)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ChatParserWORD || _la == ChatParserNEWLINE {
		{
			p.SetState(18)
			p.Line()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(23)
		p.Match(ChatParserEOF)
	}

	return localctx
}

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChatParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChatParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(ChatParserNEWLINE, 0)
}

func (s *LineContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *LineContext) Command() ICommandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *LineContext) Message() IMessageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageContext)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.ExitLine(s)
	}
}

func (s *LineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ChatVisitor:
		return t.VisitLine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ChatParser) Line() (localctx ILineContext) {
	this := p
	_ = this

	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ChatParserRULE_line)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ChatParserWORD {
		{
			p.SetState(25)
			p.Name()
		}
		{
			p.SetState(26)
			p.Command()
		}
		{
			p.SetState(27)
			p.Message()
		}

	}
	{
		p.SetState(31)
		p.Match(ChatParserNEWLINE)
	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChatParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChatParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) WORD() antlr.TerminalNode {
	return s.GetToken(ChatParserWORD, 0)
}

func (s *NameContext) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(ChatParserWHITESPACE, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.ExitName(s)
	}
}

func (s *NameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ChatVisitor:
		return t.VisitName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ChatParser) Name() (localctx INameContext) {
	this := p
	_ = this

	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ChatParserRULE_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(33)
		p.Match(ChatParserWORD)
	}
	{
		p.SetState(34)
		p.Match(ChatParserWHITESPACE)
	}

	return localctx
}

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChatParserRULE_command
	return p
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChatParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(ChatParserWHITESPACE, 0)
}

func (s *CommandContext) SAYS() antlr.TerminalNode {
	return s.GetToken(ChatParserSAYS, 0)
}

func (s *CommandContext) SHOUTS() antlr.TerminalNode {
	return s.GetToken(ChatParserSHOUTS, 0)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.EnterCommand(s)
	}
}

func (s *CommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.ExitCommand(s)
	}
}

func (s *CommandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ChatVisitor:
		return t.VisitCommand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ChatParser) Command() (localctx ICommandContext) {
	this := p
	_ = this

	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ChatParserRULE_command)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ChatParserSAYS || _la == ChatParserSHOUTS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(37)
		p.Match(ChatParserT__0)
	}
	{
		p.SetState(38)
		p.Match(ChatParserWHITESPACE)
	}

	return localctx
}

// IMessageContext is an interface to support dynamic dispatch.
type IMessageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMessageContext differentiates from other interfaces.
	IsMessageContext()
}

type MessageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageContext() *MessageContext {
	var p = new(MessageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChatParserRULE_message
	return p
}

func (*MessageContext) IsMessageContext() {}

func NewMessageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageContext {
	var p = new(MessageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChatParserRULE_message

	return p
}

func (s *MessageContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageContext) AllEmoticon() []IEmoticonContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEmoticonContext); ok {
			len++
		}
	}

	tst := make([]IEmoticonContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEmoticonContext); ok {
			tst[i] = t.(IEmoticonContext)
			i++
		}
	}

	return tst
}

func (s *MessageContext) Emoticon(i int) IEmoticonContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmoticonContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmoticonContext)
}

func (s *MessageContext) AllLink() []ILinkContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILinkContext); ok {
			len++
		}
	}

	tst := make([]ILinkContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILinkContext); ok {
			tst[i] = t.(ILinkContext)
			i++
		}
	}

	return tst
}

func (s *MessageContext) Link(i int) ILinkContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILinkContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILinkContext)
}

func (s *MessageContext) AllColor() []IColorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColorContext); ok {
			len++
		}
	}

	tst := make([]IColorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColorContext); ok {
			tst[i] = t.(IColorContext)
			i++
		}
	}

	return tst
}

func (s *MessageContext) Color(i int) IColorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColorContext)
}

func (s *MessageContext) AllMention() []IMentionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMentionContext); ok {
			len++
		}
	}

	tst := make([]IMentionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMentionContext); ok {
			tst[i] = t.(IMentionContext)
			i++
		}
	}

	return tst
}

func (s *MessageContext) Mention(i int) IMentionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMentionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMentionContext)
}

func (s *MessageContext) AllWORD() []antlr.TerminalNode {
	return s.GetTokens(ChatParserWORD)
}

func (s *MessageContext) WORD(i int) antlr.TerminalNode {
	return s.GetToken(ChatParserWORD, i)
}

func (s *MessageContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(ChatParserWHITESPACE)
}

func (s *MessageContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(ChatParserWHITESPACE, i)
}

func (s *MessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.EnterMessage(s)
	}
}

func (s *MessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.ExitMessage(s)
	}
}

func (s *MessageContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ChatVisitor:
		return t.VisitMessage(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ChatParser) Message() (localctx IMessageContext) {
	this := p
	_ = this

	localctx = NewMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ChatParserRULE_message)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(46)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case ChatParserT__0:
				{
					p.SetState(40)
					p.Emoticon()
				}

			case ChatParserTEXT:
				{
					p.SetState(41)
					p.Link()
				}

			case ChatParserT__4:
				{
					p.SetState(42)
					p.Color()
				}

			case ChatParserT__5:
				{
					p.SetState(43)
					p.Mention()
				}

			case ChatParserWORD:
				{
					p.SetState(44)
					p.Match(ChatParserWORD)
				}

			case ChatParserWHITESPACE:
				{
					p.SetState(45)
					p.Match(ChatParserWHITESPACE)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IEmoticonContext is an interface to support dynamic dispatch.
type IEmoticonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEmoticonContext differentiates from other interfaces.
	IsEmoticonContext()
}

type EmoticonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmoticonContext() *EmoticonContext {
	var p = new(EmoticonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChatParserRULE_emoticon
	return p
}

func (*EmoticonContext) IsEmoticonContext() {}

func NewEmoticonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmoticonContext {
	var p = new(EmoticonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChatParserRULE_emoticon

	return p
}

func (s *EmoticonContext) GetParser() antlr.Parser { return s.parser }
func (s *EmoticonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmoticonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmoticonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.EnterEmoticon(s)
	}
}

func (s *EmoticonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.ExitEmoticon(s)
	}
}

func (s *EmoticonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ChatVisitor:
		return t.VisitEmoticon(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ChatParser) Emoticon() (localctx IEmoticonContext) {
	this := p
	_ = this

	localctx = NewEmoticonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ChatParserRULE_emoticon)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)
			p.Match(ChatParserT__0)
		}
		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ChatParserT__1 {
			{
				p.SetState(51)
				p.Match(ChatParserT__1)
			}

		}
		{
			p.SetState(54)
			p.Match(ChatParserT__2)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(55)
			p.Match(ChatParserT__0)
		}
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ChatParserT__1 {
			{
				p.SetState(56)
				p.Match(ChatParserT__1)
			}

		}
		{
			p.SetState(59)
			p.Match(ChatParserT__3)
		}

	}

	return localctx
}

// ILinkContext is an interface to support dynamic dispatch.
type ILinkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLinkContext differentiates from other interfaces.
	IsLinkContext()
}

type LinkContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLinkContext() *LinkContext {
	var p = new(LinkContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChatParserRULE_link
	return p
}

func (*LinkContext) IsLinkContext() {}

func NewLinkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LinkContext {
	var p = new(LinkContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChatParserRULE_link

	return p
}

func (s *LinkContext) GetParser() antlr.Parser { return s.parser }

func (s *LinkContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(ChatParserTEXT)
}

func (s *LinkContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(ChatParserTEXT, i)
}

func (s *LinkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LinkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LinkContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.EnterLink(s)
	}
}

func (s *LinkContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.ExitLink(s)
	}
}

func (s *LinkContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ChatVisitor:
		return t.VisitLink(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ChatParser) Link() (localctx ILinkContext) {
	this := p
	_ = this

	localctx = NewLinkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ChatParserRULE_link)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Match(ChatParserTEXT)
	}
	{
		p.SetState(63)
		p.Match(ChatParserTEXT)
	}

	return localctx
}

// IColorContext is an interface to support dynamic dispatch.
type IColorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColorContext differentiates from other interfaces.
	IsColorContext()
}

type ColorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColorContext() *ColorContext {
	var p = new(ColorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChatParserRULE_color
	return p
}

func (*ColorContext) IsColorContext() {}

func NewColorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColorContext {
	var p = new(ColorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChatParserRULE_color

	return p
}

func (s *ColorContext) GetParser() antlr.Parser { return s.parser }

func (s *ColorContext) WORD() antlr.TerminalNode {
	return s.GetToken(ChatParserWORD, 0)
}

func (s *ColorContext) Message() IMessageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageContext)
}

func (s *ColorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.EnterColor(s)
	}
}

func (s *ColorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.ExitColor(s)
	}
}

func (s *ColorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ChatVisitor:
		return t.VisitColor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ChatParser) Color() (localctx IColorContext) {
	this := p
	_ = this

	localctx = NewColorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ChatParserRULE_color)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Match(ChatParserT__4)
	}
	{
		p.SetState(66)
		p.Match(ChatParserWORD)
	}
	{
		p.SetState(67)
		p.Match(ChatParserT__4)
	}
	{
		p.SetState(68)
		p.Message()
	}
	{
		p.SetState(69)
		p.Match(ChatParserT__4)
	}

	return localctx
}

// IMentionContext is an interface to support dynamic dispatch.
type IMentionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMentionContext differentiates from other interfaces.
	IsMentionContext()
}

type MentionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMentionContext() *MentionContext {
	var p = new(MentionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChatParserRULE_mention
	return p
}

func (*MentionContext) IsMentionContext() {}

func NewMentionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MentionContext {
	var p = new(MentionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChatParserRULE_mention

	return p
}

func (s *MentionContext) GetParser() antlr.Parser { return s.parser }

func (s *MentionContext) WORD() antlr.TerminalNode {
	return s.GetToken(ChatParserWORD, 0)
}

func (s *MentionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MentionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MentionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.EnterMention(s)
	}
}

func (s *MentionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChatListener); ok {
		listenerT.ExitMention(s)
	}
}

func (s *MentionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ChatVisitor:
		return t.VisitMention(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ChatParser) Mention() (localctx IMentionContext) {
	this := p
	_ = this

	localctx = NewMentionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ChatParserRULE_mention)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Match(ChatParserT__5)
	}
	{
		p.SetState(72)
		p.Match(ChatParserWORD)
	}

	return localctx
}
