// Code generated from Minml.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Minml

import (
	"fmt"
	"strconv"
	"sync"

	antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type MinmlParser struct {
	*antlr.BaseParser
}

var minmlParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func minmlParserInit() {
	staticData := &minmlParserStaticData
	staticData.literalNames = []string{
		"", "'{'", "','", "'}'", "':'", "'['", "']'", "'false'", "'null'", "'true'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "FALSE", "NULL", "TRUE", "DQ_STRING", "SQ_STRING",
		"TRI_DQ_STRING", "TRI_SQ_STRING", "IDENTIFIER", "NUMBER", "LINE_COMMENT",
		"BLOCK_COMMENT", "WS",
	}
	staticData.ruleNames = []string{
		"root", "value", "obj", "pair", "array", "identifier", "stringValue",
		"number", "true", "false", "null",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 18, 88, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 33,
		8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 39, 8, 2, 10, 2, 12, 2, 42, 9, 2, 1,
		2, 3, 2, 45, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 51, 8, 2, 1, 3, 1, 3,
		1, 3, 3, 3, 56, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 62, 8, 4, 10, 4, 12,
		4, 65, 9, 4, 1, 4, 3, 4, 68, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 74, 8,
		4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 10, 0, 0, 11, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 0, 1, 1, 0,
		10, 13, 90, 0, 22, 1, 0, 0, 0, 2, 32, 1, 0, 0, 0, 4, 50, 1, 0, 0, 0, 6,
		52, 1, 0, 0, 0, 8, 73, 1, 0, 0, 0, 10, 75, 1, 0, 0, 0, 12, 77, 1, 0, 0,
		0, 14, 79, 1, 0, 0, 0, 16, 81, 1, 0, 0, 0, 18, 83, 1, 0, 0, 0, 20, 85,
		1, 0, 0, 0, 22, 23, 3, 2, 1, 0, 23, 1, 1, 0, 0, 0, 24, 33, 3, 4, 2, 0,
		25, 33, 3, 8, 4, 0, 26, 33, 3, 12, 6, 0, 27, 33, 3, 14, 7, 0, 28, 33, 3,
		16, 8, 0, 29, 33, 3, 18, 9, 0, 30, 33, 3, 20, 10, 0, 31, 33, 3, 10, 5,
		0, 32, 24, 1, 0, 0, 0, 32, 25, 1, 0, 0, 0, 32, 26, 1, 0, 0, 0, 32, 27,
		1, 0, 0, 0, 32, 28, 1, 0, 0, 0, 32, 29, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0,
		32, 31, 1, 0, 0, 0, 33, 3, 1, 0, 0, 0, 34, 35, 5, 1, 0, 0, 35, 40, 3, 6,
		3, 0, 36, 37, 5, 2, 0, 0, 37, 39, 3, 6, 3, 0, 38, 36, 1, 0, 0, 0, 39, 42,
		1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 44, 1, 0, 0, 0,
		42, 40, 1, 0, 0, 0, 43, 45, 5, 2, 0, 0, 44, 43, 1, 0, 0, 0, 44, 45, 1,
		0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 47, 5, 3, 0, 0, 47, 51, 1, 0, 0, 0, 48,
		49, 5, 1, 0, 0, 49, 51, 5, 3, 0, 0, 50, 34, 1, 0, 0, 0, 50, 48, 1, 0, 0,
		0, 51, 5, 1, 0, 0, 0, 52, 55, 3, 2, 1, 0, 53, 54, 5, 4, 0, 0, 54, 56, 3,
		2, 1, 0, 55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 7, 1, 0, 0, 0, 57,
		58, 5, 5, 0, 0, 58, 63, 3, 2, 1, 0, 59, 60, 5, 2, 0, 0, 60, 62, 3, 2, 1,
		0, 61, 59, 1, 0, 0, 0, 62, 65, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 63, 64,
		1, 0, 0, 0, 64, 67, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 66, 68, 5, 2, 0, 0,
		67, 66, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 70, 5,
		6, 0, 0, 70, 74, 1, 0, 0, 0, 71, 72, 5, 5, 0, 0, 72, 74, 5, 6, 0, 0, 73,
		57, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 74, 9, 1, 0, 0, 0, 75, 76, 5, 14, 0,
		0, 76, 11, 1, 0, 0, 0, 77, 78, 7, 0, 0, 0, 78, 13, 1, 0, 0, 0, 79, 80,
		5, 15, 0, 0, 80, 15, 1, 0, 0, 0, 81, 82, 5, 9, 0, 0, 82, 17, 1, 0, 0, 0,
		83, 84, 5, 7, 0, 0, 84, 19, 1, 0, 0, 0, 85, 86, 5, 8, 0, 0, 86, 21, 1,
		0, 0, 0, 8, 32, 40, 44, 50, 55, 63, 67, 73,
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

// MinmlParserInit initializes any static state used to implement MinmlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMinmlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MinmlParserInit() {
	staticData := &minmlParserStaticData
	staticData.once.Do(minmlParserInit)
}

// NewMinmlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMinmlParser(input antlr.TokenStream) *MinmlParser {
	MinmlParserInit()
	this := new(MinmlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &minmlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Minml.g4"

	return this
}

// MinmlParser tokens.
const (
	MinmlParserEOF           = antlr.TokenEOF
	MinmlParserT__0          = 1
	MinmlParserT__1          = 2
	MinmlParserT__2          = 3
	MinmlParserT__3          = 4
	MinmlParserT__4          = 5
	MinmlParserT__5          = 6
	MinmlParserFALSE         = 7
	MinmlParserNULL          = 8
	MinmlParserTRUE          = 9
	MinmlParserDQ_STRING     = 10
	MinmlParserSQ_STRING     = 11
	MinmlParserTRI_DQ_STRING = 12
	MinmlParserTRI_SQ_STRING = 13
	MinmlParserIDENTIFIER    = 14
	MinmlParserNUMBER        = 15
	MinmlParserLINE_COMMENT  = 16
	MinmlParserBLOCK_COMMENT = 17
	MinmlParserWS            = 18
)

// MinmlParser rules.
const (
	MinmlParserRULE_root        = 0
	MinmlParserRULE_value       = 1
	MinmlParserRULE_obj         = 2
	MinmlParserRULE_pair        = 3
	MinmlParserRULE_array       = 4
	MinmlParserRULE_identifier  = 5
	MinmlParserRULE_stringValue = 6
	MinmlParserRULE_number      = 7
	MinmlParserRULE_true        = 8
	MinmlParserRULE_false       = 9
	MinmlParserRULE_null        = 10
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MinmlParserRULE_root
	return p
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinmlParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinmlListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinmlListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (s *RootContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinmlVisitor:
		return t.VisitRoot(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinmlParser) Root() (localctx IRootContext) {
	this := p
	_ = this

	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MinmlParserRULE_root)

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
		p.SetState(22)
		p.Value()
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MinmlParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinmlParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Obj() IObjContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjContext)
}

func (s *ValueContext) Array() IArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ValueContext) StringValue() IStringValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *ValueContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ValueContext) True() ITrueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITrueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITrueContext)
}

func (s *ValueContext) False() IFalseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFalseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFalseContext)
}

func (s *ValueContext) Null() INullContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INullContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INullContext)
}

func (s *ValueContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinmlListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinmlListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinmlVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinmlParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MinmlParserRULE_value)

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

	p.SetState(32)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MinmlParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(24)
			p.Obj()
		}

	case MinmlParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(25)
			p.Array()
		}

	case MinmlParserDQ_STRING, MinmlParserSQ_STRING, MinmlParserTRI_DQ_STRING, MinmlParserTRI_SQ_STRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(26)
			p.StringValue()
		}

	case MinmlParserNUMBER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(27)
			p.Number()
		}

	case MinmlParserTRUE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(28)
			p.True()
		}

	case MinmlParserFALSE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(29)
			p.False()
		}

	case MinmlParserNULL:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(30)
			p.Null()
		}

	case MinmlParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(31)
			p.Identifier()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IObjContext is an interface to support dynamic dispatch.
type IObjContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjContext differentiates from other interfaces.
	IsObjContext()
}

type ObjContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjContext() *ObjContext {
	var p = new(ObjContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MinmlParserRULE_obj
	return p
}

func (*ObjContext) IsObjContext() {}

func NewObjContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjContext {
	var p = new(ObjContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinmlParserRULE_obj

	return p
}

func (s *ObjContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjContext) AllPair() []IPairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPairContext); ok {
			len++
		}
	}

	tst := make([]IPairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPairContext); ok {
			tst[i] = t.(IPairContext)
			i++
		}
	}

	return tst
}

func (s *ObjContext) Pair(i int) IPairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPairContext); ok {
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

	return t.(IPairContext)
}

func (s *ObjContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinmlListener); ok {
		listenerT.EnterObj(s)
	}
}

func (s *ObjContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinmlListener); ok {
		listenerT.ExitObj(s)
	}
}

func (s *ObjContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinmlVisitor:
		return t.VisitObj(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinmlParser) Obj() (localctx IObjContext) {
	this := p
	_ = this

	localctx = NewObjContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MinmlParserRULE_obj)
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

	var _alt int

	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Match(MinmlParserT__0)
		}
		{
			p.SetState(35)
			p.Pair()
		}
		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(36)
					p.Match(MinmlParserT__1)
				}
				{
					p.SetState(37)
					p.Pair()
				}

			}
			p.SetState(42)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == MinmlParserT__1 {
			{
				p.SetState(43)
				p.Match(MinmlParserT__1)
			}

		}
		{
			p.SetState(46)
			p.Match(MinmlParserT__2)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(48)
			p.Match(MinmlParserT__0)
		}
		{
			p.SetState(49)
			p.Match(MinmlParserT__2)
		}

	}

	return localctx
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetK returns the k rule contexts.
	GetK() IValueContext

	// GetV returns the v rule contexts.
	GetV() IValueContext

	// SetK sets the k rule contexts.
	SetK(IValueContext)

	// SetV sets the v rule contexts.
	SetV(IValueContext)

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	k      IValueContext
	v      IValueContext
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MinmlParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinmlParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) GetK() IValueContext { return s.k }

func (s *PairContext) GetV() IValueContext { return s.v }

func (s *PairContext) SetK(v IValueContext) { s.k = v }

func (s *PairContext) SetV(v IValueContext) { s.v = v }

func (s *PairContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *PairContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
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

	return t.(IValueContext)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinmlListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinmlListener); ok {
		listenerT.ExitPair(s)
	}
}

func (s *PairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinmlVisitor:
		return t.VisitPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinmlParser) Pair() (localctx IPairContext) {
	this := p
	_ = this

	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MinmlParserRULE_pair)
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
		p.SetState(52)

		var _x = p.Value()

		localctx.(*PairContext).k = _x
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MinmlParserT__3 {
		{
			p.SetState(53)
			p.Match(MinmlParserT__3)
		}
		{
			p.SetState(54)

			var _x = p.Value()

			localctx.(*PairContext).v = _x
		}

	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MinmlParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinmlParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *ArrayContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
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

	return t.(IValueContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinmlListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinmlListener); ok {
		listenerT.ExitArray(s)
	}
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinmlVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinmlParser) Array() (localctx IArrayContext) {
	this := p
	_ = this

	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MinmlParserRULE_array)
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

	var _alt int

	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(57)
			p.Match(MinmlParserT__4)
		}
		{
			p.SetState(58)
			p.Value()
		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(59)
					p.Match(MinmlParserT__1)
				}
				{
					p.SetState(60)
					p.Value()
				}

			}
			p.SetState(65)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == MinmlParserT__1 {
			{
				p.SetState(66)
				p.Match(MinmlParserT__1)
			}

		}
		{
			p.SetState(69)
			p.Match(MinmlParserT__5)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.Match(MinmlParserT__4)
		}
		{
			p.SetState(72)
			p.Match(MinmlParserT__5)
		}

	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MinmlParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinmlParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MinmlParserIDENTIFIER, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinmlListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinmlListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinmlVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinmlParser) Identifier() (localctx IIdentifierContext) {
	this := p
	_ = this

	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MinmlParserRULE_identifier)

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
		p.SetState(75)
		p.Match(MinmlParserIDENTIFIER)
	}

	return localctx
}

// IStringValueContext is an interface to support dynamic dispatch.
type IStringValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringValueContext differentiates from other interfaces.
	IsStringValueContext()
}

type StringValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringValueContext() *StringValueContext {
	var p = new(StringValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MinmlParserRULE_stringValue
	return p
}

func (*StringValueContext) IsStringValueContext() {}

func NewStringValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringValueContext {
	var p = new(StringValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinmlParserRULE_stringValue

	return p
}

func (s *StringValueContext) GetParser() antlr.Parser { return s.parser }

func (s *StringValueContext) TRI_DQ_STRING() antlr.TerminalNode {
	return s.GetToken(MinmlParserTRI_DQ_STRING, 0)
}

func (s *StringValueContext) TRI_SQ_STRING() antlr.TerminalNode {
	return s.GetToken(MinmlParserTRI_SQ_STRING, 0)
}

func (s *StringValueContext) DQ_STRING() antlr.TerminalNode {
	return s.GetToken(MinmlParserDQ_STRING, 0)
}

func (s *StringValueContext) SQ_STRING() antlr.TerminalNode {
	return s.GetToken(MinmlParserSQ_STRING, 0)
}

func (s *StringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinmlListener); ok {
		listenerT.EnterStringValue(s)
	}
}

func (s *StringValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinmlListener); ok {
		listenerT.ExitStringValue(s)
	}
}

func (s *StringValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinmlVisitor:
		return t.VisitStringValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinmlParser) StringValue() (localctx IStringValueContext) {
	this := p
	_ = this

	localctx = NewStringValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MinmlParserRULE_stringValue)
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
		p.SetState(77)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<MinmlParserDQ_STRING)|(1<<MinmlParserSQ_STRING)|(1<<MinmlParserTRI_DQ_STRING)|(1<<MinmlParserTRI_SQ_STRING))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MinmlParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinmlParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MinmlParserNUMBER, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinmlListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinmlListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinmlVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinmlParser) Number() (localctx INumberContext) {
	this := p
	_ = this

	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MinmlParserRULE_number)

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
		p.SetState(79)
		p.Match(MinmlParserNUMBER)
	}

	return localctx
}

// ITrueContext is an interface to support dynamic dispatch.
type ITrueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTrueContext differentiates from other interfaces.
	IsTrueContext()
}

type TrueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTrueContext() *TrueContext {
	var p = new(TrueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MinmlParserRULE_true
	return p
}

func (*TrueContext) IsTrueContext() {}

func NewTrueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TrueContext {
	var p = new(TrueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinmlParserRULE_true

	return p
}

func (s *TrueContext) GetParser() antlr.Parser { return s.parser }

func (s *TrueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(MinmlParserTRUE, 0)
}

func (s *TrueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TrueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinmlListener); ok {
		listenerT.EnterTrue(s)
	}
}

func (s *TrueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinmlListener); ok {
		listenerT.ExitTrue(s)
	}
}

func (s *TrueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinmlVisitor:
		return t.VisitTrue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinmlParser) True() (localctx ITrueContext) {
	this := p
	_ = this

	localctx = NewTrueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MinmlParserRULE_true)

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
		p.SetState(81)
		p.Match(MinmlParserTRUE)
	}

	return localctx
}

// IFalseContext is an interface to support dynamic dispatch.
type IFalseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFalseContext differentiates from other interfaces.
	IsFalseContext()
}

type FalseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFalseContext() *FalseContext {
	var p = new(FalseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MinmlParserRULE_false
	return p
}

func (*FalseContext) IsFalseContext() {}

func NewFalseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FalseContext {
	var p = new(FalseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinmlParserRULE_false

	return p
}

func (s *FalseContext) GetParser() antlr.Parser { return s.parser }

func (s *FalseContext) FALSE() antlr.TerminalNode {
	return s.GetToken(MinmlParserFALSE, 0)
}

func (s *FalseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FalseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FalseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinmlListener); ok {
		listenerT.EnterFalse(s)
	}
}

func (s *FalseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinmlListener); ok {
		listenerT.ExitFalse(s)
	}
}

func (s *FalseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinmlVisitor:
		return t.VisitFalse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinmlParser) False() (localctx IFalseContext) {
	this := p
	_ = this

	localctx = NewFalseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MinmlParserRULE_false)

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
		p.SetState(83)
		p.Match(MinmlParserFALSE)
	}

	return localctx
}

// INullContext is an interface to support dynamic dispatch.
type INullContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNullContext differentiates from other interfaces.
	IsNullContext()
}

type NullContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNullContext() *NullContext {
	var p = new(NullContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MinmlParserRULE_null
	return p
}

func (*NullContext) IsNullContext() {}

func NewNullContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NullContext {
	var p = new(NullContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MinmlParserRULE_null

	return p
}

func (s *NullContext) GetParser() antlr.Parser { return s.parser }

func (s *NullContext) NULL() antlr.TerminalNode {
	return s.GetToken(MinmlParserNULL, 0)
}

func (s *NullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinmlListener); ok {
		listenerT.EnterNull(s)
	}
}

func (s *NullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MinmlListener); ok {
		listenerT.ExitNull(s)
	}
}

func (s *NullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MinmlVisitor:
		return t.VisitNull(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MinmlParser) Null() (localctx INullContext) {
	this := p
	_ = this

	localctx = NewNullContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MinmlParserRULE_null)

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
		p.SetState(85)
		p.Match(MinmlParserNULL)
	}

	return localctx
}
