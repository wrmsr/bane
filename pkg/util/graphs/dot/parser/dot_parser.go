// Code generated from Dot.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Dot

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

type DotParser struct {
	*antlr.BaseParser
}

var dotParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func dotParserInit() {
	staticData := &dotParserStaticData
	staticData.literalNames = []string{
		"", "'{'", "'}'", "';'", "'='", "'['", "']'", "','", "'->'", "'--'",
		"':'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "STRICT", "GRAPH", "DIGRAPH",
		"NODE", "EDGE", "SUBGRAPH", "NUMBER", "STRING", "IDENT", "HTML_STRING",
		"COMMENT", "LINE_COMMENT", "PREPROC", "WS",
	}
	staticData.ruleNames = []string{
		"graph", "stmtList", "stmt", "attrStmt", "attrList", "aList", "edgeStmt",
		"edgeRHS", "edgeop", "nodeStmt", "nodeId", "port", "subgraph", "ident",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 24, 128, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 3, 0, 30, 8, 0, 1,
		0, 1, 0, 3, 0, 34, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 42,
		8, 1, 5, 1, 44, 8, 1, 10, 1, 12, 1, 47, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 3, 2, 57, 8, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4,
		64, 8, 4, 1, 4, 4, 4, 67, 8, 4, 11, 4, 12, 4, 68, 1, 5, 1, 5, 1, 5, 3,
		5, 74, 8, 5, 1, 5, 3, 5, 77, 8, 5, 4, 5, 79, 8, 5, 11, 5, 12, 5, 80, 1,
		6, 1, 6, 3, 6, 85, 8, 6, 1, 6, 1, 6, 3, 6, 89, 8, 6, 1, 7, 1, 7, 1, 7,
		3, 7, 94, 8, 7, 4, 7, 96, 8, 7, 11, 7, 12, 7, 97, 1, 8, 1, 8, 1, 9, 1,
		9, 3, 9, 104, 8, 9, 1, 10, 1, 10, 3, 10, 108, 8, 10, 1, 11, 1, 11, 1, 11,
		1, 11, 3, 11, 114, 8, 11, 1, 12, 1, 12, 3, 12, 118, 8, 12, 3, 12, 120,
		8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 0, 0, 14, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 0, 4, 1, 0, 12, 13, 2, 0,
		12, 12, 14, 15, 1, 0, 8, 9, 1, 0, 17, 20, 135, 0, 29, 1, 0, 0, 0, 2, 45,
		1, 0, 0, 0, 4, 56, 1, 0, 0, 0, 6, 58, 1, 0, 0, 0, 8, 66, 1, 0, 0, 0, 10,
		78, 1, 0, 0, 0, 12, 84, 1, 0, 0, 0, 14, 95, 1, 0, 0, 0, 16, 99, 1, 0, 0,
		0, 18, 101, 1, 0, 0, 0, 20, 105, 1, 0, 0, 0, 22, 109, 1, 0, 0, 0, 24, 119,
		1, 0, 0, 0, 26, 125, 1, 0, 0, 0, 28, 30, 5, 11, 0, 0, 29, 28, 1, 0, 0,
		0, 29, 30, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 33, 7, 0, 0, 0, 32, 34,
		3, 26, 13, 0, 33, 32, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 35, 1, 0, 0,
		0, 35, 36, 5, 1, 0, 0, 36, 37, 3, 2, 1, 0, 37, 38, 5, 2, 0, 0, 38, 1, 1,
		0, 0, 0, 39, 41, 3, 4, 2, 0, 40, 42, 5, 3, 0, 0, 41, 40, 1, 0, 0, 0, 41,
		42, 1, 0, 0, 0, 42, 44, 1, 0, 0, 0, 43, 39, 1, 0, 0, 0, 44, 47, 1, 0, 0,
		0, 45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 3, 1, 0, 0, 0, 47, 45, 1,
		0, 0, 0, 48, 57, 3, 18, 9, 0, 49, 57, 3, 12, 6, 0, 50, 57, 3, 6, 3, 0,
		51, 52, 3, 26, 13, 0, 52, 53, 5, 4, 0, 0, 53, 54, 3, 26, 13, 0, 54, 57,
		1, 0, 0, 0, 55, 57, 3, 24, 12, 0, 56, 48, 1, 0, 0, 0, 56, 49, 1, 0, 0,
		0, 56, 50, 1, 0, 0, 0, 56, 51, 1, 0, 0, 0, 56, 55, 1, 0, 0, 0, 57, 5, 1,
		0, 0, 0, 58, 59, 7, 1, 0, 0, 59, 60, 3, 8, 4, 0, 60, 7, 1, 0, 0, 0, 61,
		63, 5, 5, 0, 0, 62, 64, 3, 10, 5, 0, 63, 62, 1, 0, 0, 0, 63, 64, 1, 0,
		0, 0, 64, 65, 1, 0, 0, 0, 65, 67, 5, 6, 0, 0, 66, 61, 1, 0, 0, 0, 67, 68,
		1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 9, 1, 0, 0, 0,
		70, 73, 3, 26, 13, 0, 71, 72, 5, 4, 0, 0, 72, 74, 3, 26, 13, 0, 73, 71,
		1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 76, 1, 0, 0, 0, 75, 77, 5, 7, 0, 0,
		76, 75, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 79, 1, 0, 0, 0, 78, 70, 1,
		0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81,
		11, 1, 0, 0, 0, 82, 85, 3, 20, 10, 0, 83, 85, 3, 24, 12, 0, 84, 82, 1,
		0, 0, 0, 84, 83, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 88, 3, 14, 7, 0, 87,
		89, 3, 8, 4, 0, 88, 87, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 13, 1, 0, 0,
		0, 90, 93, 3, 16, 8, 0, 91, 94, 3, 20, 10, 0, 92, 94, 3, 24, 12, 0, 93,
		91, 1, 0, 0, 0, 93, 92, 1, 0, 0, 0, 94, 96, 1, 0, 0, 0, 95, 90, 1, 0, 0,
		0, 96, 97, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 15,
		1, 0, 0, 0, 99, 100, 7, 2, 0, 0, 100, 17, 1, 0, 0, 0, 101, 103, 3, 20,
		10, 0, 102, 104, 3, 8, 4, 0, 103, 102, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0,
		104, 19, 1, 0, 0, 0, 105, 107, 3, 26, 13, 0, 106, 108, 3, 22, 11, 0, 107,
		106, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 21, 1, 0, 0, 0, 109, 110, 5,
		10, 0, 0, 110, 113, 3, 26, 13, 0, 111, 112, 5, 10, 0, 0, 112, 114, 3, 26,
		13, 0, 113, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 23, 1, 0, 0, 0,
		115, 117, 5, 16, 0, 0, 116, 118, 3, 26, 13, 0, 117, 116, 1, 0, 0, 0, 117,
		118, 1, 0, 0, 0, 118, 120, 1, 0, 0, 0, 119, 115, 1, 0, 0, 0, 119, 120,
		1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 122, 5, 1, 0, 0, 122, 123, 3, 2,
		1, 0, 123, 124, 5, 2, 0, 0, 124, 25, 1, 0, 0, 0, 125, 126, 7, 3, 0, 0,
		126, 27, 1, 0, 0, 0, 19, 29, 33, 41, 45, 56, 63, 68, 73, 76, 80, 84, 88,
		93, 97, 103, 107, 113, 117, 119,
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

// DotParserInit initializes any static state used to implement DotParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDotParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DotParserInit() {
	staticData := &dotParserStaticData
	staticData.once.Do(dotParserInit)
}

// NewDotParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDotParser(input antlr.TokenStream) *DotParser {
	DotParserInit()
	this := new(DotParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &dotParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Dot.g4"

	return this
}

// DotParser tokens.
const (
	DotParserEOF          = antlr.TokenEOF
	DotParserT__0         = 1
	DotParserT__1         = 2
	DotParserT__2         = 3
	DotParserT__3         = 4
	DotParserT__4         = 5
	DotParserT__5         = 6
	DotParserT__6         = 7
	DotParserT__7         = 8
	DotParserT__8         = 9
	DotParserT__9         = 10
	DotParserSTRICT       = 11
	DotParserGRAPH        = 12
	DotParserDIGRAPH      = 13
	DotParserNODE         = 14
	DotParserEDGE         = 15
	DotParserSUBGRAPH     = 16
	DotParserNUMBER       = 17
	DotParserSTRING       = 18
	DotParserIDENT        = 19
	DotParserHTML_STRING  = 20
	DotParserCOMMENT      = 21
	DotParserLINE_COMMENT = 22
	DotParserPREPROC      = 23
	DotParserWS           = 24
)

// DotParser rules.
const (
	DotParserRULE_graph    = 0
	DotParserRULE_stmtList = 1
	DotParserRULE_stmt     = 2
	DotParserRULE_attrStmt = 3
	DotParserRULE_attrList = 4
	DotParserRULE_aList    = 5
	DotParserRULE_edgeStmt = 6
	DotParserRULE_edgeRHS  = 7
	DotParserRULE_edgeop   = 8
	DotParserRULE_nodeStmt = 9
	DotParserRULE_nodeId   = 10
	DotParserRULE_port     = 11
	DotParserRULE_subgraph = 12
	DotParserRULE_ident    = 13
)

// IGraphContext is an interface to support dynamic dispatch.
type IGraphContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGraphContext differentiates from other interfaces.
	IsGraphContext()
}

type GraphContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGraphContext() *GraphContext {
	var p = new(GraphContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DotParserRULE_graph
	return p
}

func (*GraphContext) IsGraphContext() {}

func NewGraphContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GraphContext {
	var p = new(GraphContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DotParserRULE_graph

	return p
}

func (s *GraphContext) GetParser() antlr.Parser { return s.parser }

func (s *GraphContext) StmtList() IStmtListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtListContext)
}

func (s *GraphContext) GRAPH() antlr.TerminalNode {
	return s.GetToken(DotParserGRAPH, 0)
}

func (s *GraphContext) DIGRAPH() antlr.TerminalNode {
	return s.GetToken(DotParserDIGRAPH, 0)
}

func (s *GraphContext) STRICT() antlr.TerminalNode {
	return s.GetToken(DotParserSTRICT, 0)
}

func (s *GraphContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *GraphContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GraphContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GraphContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.EnterGraph(s)
	}
}

func (s *GraphContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.ExitGraph(s)
	}
}

func (s *GraphContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitGraph(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DotParser) Graph() (localctx IGraphContext) {
	this := p
	_ = this

	localctx = NewGraphContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DotParserRULE_graph)
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

	if _la == DotParserSTRICT {
		{
			p.SetState(28)
			p.Match(DotParserSTRICT)
		}

	}
	{
		p.SetState(31)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DotParserGRAPH || _la == DotParserDIGRAPH) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DotParserNUMBER)|(1<<DotParserSTRING)|(1<<DotParserIDENT)|(1<<DotParserHTML_STRING))) != 0 {
		{
			p.SetState(32)
			p.Ident()
		}

	}
	{
		p.SetState(35)
		p.Match(DotParserT__0)
	}
	{
		p.SetState(36)
		p.StmtList()
	}
	{
		p.SetState(37)
		p.Match(DotParserT__1)
	}

	return localctx
}

// IStmtListContext is an interface to support dynamic dispatch.
type IStmtListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtListContext differentiates from other interfaces.
	IsStmtListContext()
}

type StmtListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtListContext() *StmtListContext {
	var p = new(StmtListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DotParserRULE_stmtList
	return p
}

func (*StmtListContext) IsStmtListContext() {}

func NewStmtListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtListContext {
	var p = new(StmtListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DotParserRULE_stmtList

	return p
}

func (s *StmtListContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtListContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *StmtListContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
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

	return t.(IStmtContext)
}

func (s *StmtListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.EnterStmtList(s)
	}
}

func (s *StmtListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.ExitStmtList(s)
	}
}

func (s *StmtListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitStmtList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DotParser) StmtList() (localctx IStmtListContext) {
	this := p
	_ = this

	localctx = NewStmtListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DotParserRULE_stmtList)
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
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DotParserT__0)|(1<<DotParserGRAPH)|(1<<DotParserNODE)|(1<<DotParserEDGE)|(1<<DotParserSUBGRAPH)|(1<<DotParserNUMBER)|(1<<DotParserSTRING)|(1<<DotParserIDENT)|(1<<DotParserHTML_STRING))) != 0 {
		{
			p.SetState(39)
			p.Stmt()
		}
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DotParserT__2 {
			{
				p.SetState(40)
				p.Match(DotParserT__2)
			}

		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DotParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DotParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) NodeStmt() INodeStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INodeStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INodeStmtContext)
}

func (s *StmtContext) EdgeStmt() IEdgeStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEdgeStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEdgeStmtContext)
}

func (s *StmtContext) AttrStmt() IAttrStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrStmtContext)
}

func (s *StmtContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *StmtContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
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

	return t.(IIdentContext)
}

func (s *StmtContext) Subgraph() ISubgraphContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubgraphContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubgraphContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DotParser) Stmt() (localctx IStmtContext) {
	this := p
	_ = this

	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DotParserRULE_stmt)

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

	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(48)
			p.NodeStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)
			p.EdgeStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(50)
			p.AttrStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(51)
			p.Ident()
		}
		{
			p.SetState(52)
			p.Match(DotParserT__3)
		}
		{
			p.SetState(53)
			p.Ident()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(55)
			p.Subgraph()
		}

	}

	return localctx
}

// IAttrStmtContext is an interface to support dynamic dispatch.
type IAttrStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttrStmtContext differentiates from other interfaces.
	IsAttrStmtContext()
}

type AttrStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrStmtContext() *AttrStmtContext {
	var p = new(AttrStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DotParserRULE_attrStmt
	return p
}

func (*AttrStmtContext) IsAttrStmtContext() {}

func NewAttrStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrStmtContext {
	var p = new(AttrStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DotParserRULE_attrStmt

	return p
}

func (s *AttrStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrStmtContext) AttrList() IAttrListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrListContext)
}

func (s *AttrStmtContext) GRAPH() antlr.TerminalNode {
	return s.GetToken(DotParserGRAPH, 0)
}

func (s *AttrStmtContext) NODE() antlr.TerminalNode {
	return s.GetToken(DotParserNODE, 0)
}

func (s *AttrStmtContext) EDGE() antlr.TerminalNode {
	return s.GetToken(DotParserEDGE, 0)
}

func (s *AttrStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.EnterAttrStmt(s)
	}
}

func (s *AttrStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.ExitAttrStmt(s)
	}
}

func (s *AttrStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitAttrStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DotParser) AttrStmt() (localctx IAttrStmtContext) {
	this := p
	_ = this

	localctx = NewAttrStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DotParserRULE_attrStmt)
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
		p.SetState(58)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DotParserGRAPH)|(1<<DotParserNODE)|(1<<DotParserEDGE))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(59)
		p.AttrList()
	}

	return localctx
}

// IAttrListContext is an interface to support dynamic dispatch.
type IAttrListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttrListContext differentiates from other interfaces.
	IsAttrListContext()
}

type AttrListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrListContext() *AttrListContext {
	var p = new(AttrListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DotParserRULE_attrList
	return p
}

func (*AttrListContext) IsAttrListContext() {}

func NewAttrListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrListContext {
	var p = new(AttrListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DotParserRULE_attrList

	return p
}

func (s *AttrListContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrListContext) AllAList() []IAListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAListContext); ok {
			len++
		}
	}

	tst := make([]IAListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAListContext); ok {
			tst[i] = t.(IAListContext)
			i++
		}
	}

	return tst
}

func (s *AttrListContext) AList(i int) IAListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAListContext); ok {
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

	return t.(IAListContext)
}

func (s *AttrListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.EnterAttrList(s)
	}
}

func (s *AttrListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.ExitAttrList(s)
	}
}

func (s *AttrListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitAttrList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DotParser) AttrList() (localctx IAttrListContext) {
	this := p
	_ = this

	localctx = NewAttrListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DotParserRULE_attrList)
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
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DotParserT__4 {
		{
			p.SetState(61)
			p.Match(DotParserT__4)
		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DotParserNUMBER)|(1<<DotParserSTRING)|(1<<DotParserIDENT)|(1<<DotParserHTML_STRING))) != 0 {
			{
				p.SetState(62)
				p.AList()
			}

		}
		{
			p.SetState(65)
			p.Match(DotParserT__5)
		}

		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAListContext is an interface to support dynamic dispatch.
type IAListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAListContext differentiates from other interfaces.
	IsAListContext()
}

type AListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAListContext() *AListContext {
	var p = new(AListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DotParserRULE_aList
	return p
}

func (*AListContext) IsAListContext() {}

func NewAListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AListContext {
	var p = new(AListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DotParserRULE_aList

	return p
}

func (s *AListContext) GetParser() antlr.Parser { return s.parser }

func (s *AListContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *AListContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
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

	return t.(IIdentContext)
}

func (s *AListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.EnterAList(s)
	}
}

func (s *AListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.ExitAList(s)
	}
}

func (s *AListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitAList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DotParser) AList() (localctx IAListContext) {
	this := p
	_ = this

	localctx = NewAListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DotParserRULE_aList)
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
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DotParserNUMBER)|(1<<DotParserSTRING)|(1<<DotParserIDENT)|(1<<DotParserHTML_STRING))) != 0) {
		{
			p.SetState(70)
			p.Ident()
		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DotParserT__3 {
			{
				p.SetState(71)
				p.Match(DotParserT__3)
			}
			{
				p.SetState(72)
				p.Ident()
			}

		}
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DotParserT__6 {
			{
				p.SetState(75)
				p.Match(DotParserT__6)
			}

		}

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IEdgeStmtContext is an interface to support dynamic dispatch.
type IEdgeStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEdgeStmtContext differentiates from other interfaces.
	IsEdgeStmtContext()
}

type EdgeStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEdgeStmtContext() *EdgeStmtContext {
	var p = new(EdgeStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DotParserRULE_edgeStmt
	return p
}

func (*EdgeStmtContext) IsEdgeStmtContext() {}

func NewEdgeStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EdgeStmtContext {
	var p = new(EdgeStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DotParserRULE_edgeStmt

	return p
}

func (s *EdgeStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *EdgeStmtContext) EdgeRHS() IEdgeRHSContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEdgeRHSContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEdgeRHSContext)
}

func (s *EdgeStmtContext) NodeId() INodeIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INodeIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INodeIdContext)
}

func (s *EdgeStmtContext) Subgraph() ISubgraphContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubgraphContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubgraphContext)
}

func (s *EdgeStmtContext) AttrList() IAttrListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrListContext)
}

func (s *EdgeStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EdgeStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EdgeStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.EnterEdgeStmt(s)
	}
}

func (s *EdgeStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.ExitEdgeStmt(s)
	}
}

func (s *EdgeStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitEdgeStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DotParser) EdgeStmt() (localctx IEdgeStmtContext) {
	this := p
	_ = this

	localctx = NewEdgeStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DotParserRULE_edgeStmt)
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
	p.SetState(84)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DotParserNUMBER, DotParserSTRING, DotParserIDENT, DotParserHTML_STRING:
		{
			p.SetState(82)
			p.NodeId()
		}

	case DotParserT__0, DotParserSUBGRAPH:
		{
			p.SetState(83)
			p.Subgraph()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(86)
		p.EdgeRHS()
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DotParserT__4 {
		{
			p.SetState(87)
			p.AttrList()
		}

	}

	return localctx
}

// IEdgeRHSContext is an interface to support dynamic dispatch.
type IEdgeRHSContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEdgeRHSContext differentiates from other interfaces.
	IsEdgeRHSContext()
}

type EdgeRHSContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEdgeRHSContext() *EdgeRHSContext {
	var p = new(EdgeRHSContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DotParserRULE_edgeRHS
	return p
}

func (*EdgeRHSContext) IsEdgeRHSContext() {}

func NewEdgeRHSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EdgeRHSContext {
	var p = new(EdgeRHSContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DotParserRULE_edgeRHS

	return p
}

func (s *EdgeRHSContext) GetParser() antlr.Parser { return s.parser }

func (s *EdgeRHSContext) AllEdgeop() []IEdgeopContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEdgeopContext); ok {
			len++
		}
	}

	tst := make([]IEdgeopContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEdgeopContext); ok {
			tst[i] = t.(IEdgeopContext)
			i++
		}
	}

	return tst
}

func (s *EdgeRHSContext) Edgeop(i int) IEdgeopContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEdgeopContext); ok {
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

	return t.(IEdgeopContext)
}

func (s *EdgeRHSContext) AllNodeId() []INodeIdContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INodeIdContext); ok {
			len++
		}
	}

	tst := make([]INodeIdContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INodeIdContext); ok {
			tst[i] = t.(INodeIdContext)
			i++
		}
	}

	return tst
}

func (s *EdgeRHSContext) NodeId(i int) INodeIdContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INodeIdContext); ok {
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

	return t.(INodeIdContext)
}

func (s *EdgeRHSContext) AllSubgraph() []ISubgraphContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISubgraphContext); ok {
			len++
		}
	}

	tst := make([]ISubgraphContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISubgraphContext); ok {
			tst[i] = t.(ISubgraphContext)
			i++
		}
	}

	return tst
}

func (s *EdgeRHSContext) Subgraph(i int) ISubgraphContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubgraphContext); ok {
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

	return t.(ISubgraphContext)
}

func (s *EdgeRHSContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EdgeRHSContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EdgeRHSContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.EnterEdgeRHS(s)
	}
}

func (s *EdgeRHSContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.ExitEdgeRHS(s)
	}
}

func (s *EdgeRHSContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitEdgeRHS(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DotParser) EdgeRHS() (localctx IEdgeRHSContext) {
	this := p
	_ = this

	localctx = NewEdgeRHSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DotParserRULE_edgeRHS)
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
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DotParserT__7 || _la == DotParserT__8 {
		{
			p.SetState(90)
			p.Edgeop()
		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case DotParserNUMBER, DotParserSTRING, DotParserIDENT, DotParserHTML_STRING:
			{
				p.SetState(91)
				p.NodeId()
			}

		case DotParserT__0, DotParserSUBGRAPH:
			{
				p.SetState(92)
				p.Subgraph()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IEdgeopContext is an interface to support dynamic dispatch.
type IEdgeopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEdgeopContext differentiates from other interfaces.
	IsEdgeopContext()
}

type EdgeopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEdgeopContext() *EdgeopContext {
	var p = new(EdgeopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DotParserRULE_edgeop
	return p
}

func (*EdgeopContext) IsEdgeopContext() {}

func NewEdgeopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EdgeopContext {
	var p = new(EdgeopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DotParserRULE_edgeop

	return p
}

func (s *EdgeopContext) GetParser() antlr.Parser { return s.parser }
func (s *EdgeopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EdgeopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EdgeopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.EnterEdgeop(s)
	}
}

func (s *EdgeopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.ExitEdgeop(s)
	}
}

func (s *EdgeopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitEdgeop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DotParser) Edgeop() (localctx IEdgeopContext) {
	this := p
	_ = this

	localctx = NewEdgeopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DotParserRULE_edgeop)
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
		p.SetState(99)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DotParserT__7 || _la == DotParserT__8) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INodeStmtContext is an interface to support dynamic dispatch.
type INodeStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNodeStmtContext differentiates from other interfaces.
	IsNodeStmtContext()
}

type NodeStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNodeStmtContext() *NodeStmtContext {
	var p = new(NodeStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DotParserRULE_nodeStmt
	return p
}

func (*NodeStmtContext) IsNodeStmtContext() {}

func NewNodeStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NodeStmtContext {
	var p = new(NodeStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DotParserRULE_nodeStmt

	return p
}

func (s *NodeStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *NodeStmtContext) NodeId() INodeIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INodeIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INodeIdContext)
}

func (s *NodeStmtContext) AttrList() IAttrListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrListContext)
}

func (s *NodeStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NodeStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NodeStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.EnterNodeStmt(s)
	}
}

func (s *NodeStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.ExitNodeStmt(s)
	}
}

func (s *NodeStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitNodeStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DotParser) NodeStmt() (localctx INodeStmtContext) {
	this := p
	_ = this

	localctx = NewNodeStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DotParserRULE_nodeStmt)
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
		p.SetState(101)
		p.NodeId()
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DotParserT__4 {
		{
			p.SetState(102)
			p.AttrList()
		}

	}

	return localctx
}

// INodeIdContext is an interface to support dynamic dispatch.
type INodeIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNodeIdContext differentiates from other interfaces.
	IsNodeIdContext()
}

type NodeIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNodeIdContext() *NodeIdContext {
	var p = new(NodeIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DotParserRULE_nodeId
	return p
}

func (*NodeIdContext) IsNodeIdContext() {}

func NewNodeIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NodeIdContext {
	var p = new(NodeIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DotParserRULE_nodeId

	return p
}

func (s *NodeIdContext) GetParser() antlr.Parser { return s.parser }

func (s *NodeIdContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *NodeIdContext) Port() IPortContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPortContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPortContext)
}

func (s *NodeIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NodeIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NodeIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.EnterNodeId(s)
	}
}

func (s *NodeIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.ExitNodeId(s)
	}
}

func (s *NodeIdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitNodeId(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DotParser) NodeId() (localctx INodeIdContext) {
	this := p
	_ = this

	localctx = NewNodeIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DotParserRULE_nodeId)
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
		p.SetState(105)
		p.Ident()
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DotParserT__9 {
		{
			p.SetState(106)
			p.Port()
		}

	}

	return localctx
}

// IPortContext is an interface to support dynamic dispatch.
type IPortContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPortContext differentiates from other interfaces.
	IsPortContext()
}

type PortContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPortContext() *PortContext {
	var p = new(PortContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DotParserRULE_port
	return p
}

func (*PortContext) IsPortContext() {}

func NewPortContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PortContext {
	var p = new(PortContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DotParserRULE_port

	return p
}

func (s *PortContext) GetParser() antlr.Parser { return s.parser }

func (s *PortContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *PortContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
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

	return t.(IIdentContext)
}

func (s *PortContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PortContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PortContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.EnterPort(s)
	}
}

func (s *PortContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.ExitPort(s)
	}
}

func (s *PortContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitPort(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DotParser) Port() (localctx IPortContext) {
	this := p
	_ = this

	localctx = NewPortContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, DotParserRULE_port)
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
		p.SetState(109)
		p.Match(DotParserT__9)
	}
	{
		p.SetState(110)
		p.Ident()
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DotParserT__9 {
		{
			p.SetState(111)
			p.Match(DotParserT__9)
		}
		{
			p.SetState(112)
			p.Ident()
		}

	}

	return localctx
}

// ISubgraphContext is an interface to support dynamic dispatch.
type ISubgraphContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubgraphContext differentiates from other interfaces.
	IsSubgraphContext()
}

type SubgraphContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubgraphContext() *SubgraphContext {
	var p = new(SubgraphContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DotParserRULE_subgraph
	return p
}

func (*SubgraphContext) IsSubgraphContext() {}

func NewSubgraphContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubgraphContext {
	var p = new(SubgraphContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DotParserRULE_subgraph

	return p
}

func (s *SubgraphContext) GetParser() antlr.Parser { return s.parser }

func (s *SubgraphContext) StmtList() IStmtListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtListContext)
}

func (s *SubgraphContext) SUBGRAPH() antlr.TerminalNode {
	return s.GetToken(DotParserSUBGRAPH, 0)
}

func (s *SubgraphContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *SubgraphContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubgraphContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubgraphContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.EnterSubgraph(s)
	}
}

func (s *SubgraphContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.ExitSubgraph(s)
	}
}

func (s *SubgraphContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitSubgraph(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DotParser) Subgraph() (localctx ISubgraphContext) {
	this := p
	_ = this

	localctx = NewSubgraphContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, DotParserRULE_subgraph)
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
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DotParserSUBGRAPH {
		{
			p.SetState(115)
			p.Match(DotParserSUBGRAPH)
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DotParserNUMBER)|(1<<DotParserSTRING)|(1<<DotParserIDENT)|(1<<DotParserHTML_STRING))) != 0 {
			{
				p.SetState(116)
				p.Ident()
			}

		}

	}
	{
		p.SetState(121)
		p.Match(DotParserT__0)
	}
	{
		p.SetState(122)
		p.StmtList()
	}
	{
		p.SetState(123)
		p.Match(DotParserT__1)
	}

	return localctx
}

// IIdentContext is an interface to support dynamic dispatch.
type IIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentContext differentiates from other interfaces.
	IsIdentContext()
}

type IdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentContext() *IdentContext {
	var p = new(IdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DotParserRULE_ident
	return p
}

func (*IdentContext) IsIdentContext() {}

func NewIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentContext {
	var p = new(IdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DotParserRULE_ident

	return p
}

func (s *IdentContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentContext) IDENT() antlr.TerminalNode {
	return s.GetToken(DotParserIDENT, 0)
}

func (s *IdentContext) STRING() antlr.TerminalNode {
	return s.GetToken(DotParserSTRING, 0)
}

func (s *IdentContext) HTML_STRING() antlr.TerminalNode {
	return s.GetToken(DotParserHTML_STRING, 0)
}

func (s *IdentContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(DotParserNUMBER, 0)
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.EnterIdent(s)
	}
}

func (s *IdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DotListener); ok {
		listenerT.ExitIdent(s)
	}
}

func (s *IdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DotVisitor:
		return t.VisitIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DotParser) Ident() (localctx IIdentContext) {
	this := p
	_ = this

	localctx = NewIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, DotParserRULE_ident)
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
		p.SetState(125)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DotParserNUMBER)|(1<<DotParserSTRING)|(1<<DotParserIDENT)|(1<<DotParserHTML_STRING))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
