// Code generated from Toml.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Toml

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

type TomlParser struct {
	*antlr.BaseParser
}

var tomlParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tomlParserInit() {
	staticData := &tomlParserStaticData
	staticData.literalNames = []string{
		"", "'='", "'.'", "'['", "']'", "','", "'{'", "'}'", "'[['", "']]'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "WS", "NL", "COMMENT", "BOOLEAN",
		"BASIC_STRING", "ML_BASIC_STRING", "LITERAL_STRING", "ML_LITERAL_STRING",
		"FLOAT", "INF", "NAN", "DEC_INT", "HEX_INT", "OCT_INT", "BIN_INT", "OFFSET_DATE_TIME",
		"LOCAL_DATE_TIME", "LOCAL_DATE", "LOCAL_TIME", "UNQUOTED_KEY",
	}
	staticData.ruleNames = []string{
		"document", "expression", "comment", "keyValue", "key", "simpleKey",
		"unquotedKey", "quotedKey", "dottedKey", "value", "stringValue", "integer",
		"floatingPoint", "boolean", "dateTime", "array", "arrayValues", "commentOrNl",
		"table", "standardTable", "inlineTable", "inlineTableKeyvals", "inlineTableKeyvalsNonEmpty",
		"arrayTable",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 29, 166, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 1, 0, 1, 0, 5, 0, 52, 8, 0,
		10, 0, 12, 0, 55, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1,
		64, 8, 1, 1, 2, 3, 2, 67, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 3,
		4, 75, 8, 4, 1, 5, 1, 5, 3, 5, 79, 8, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 4, 8, 88, 8, 8, 11, 8, 12, 8, 89, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 3, 9, 99, 8, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 3, 15, 113, 8, 15, 1, 15, 1,
		15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		3, 16, 127, 8, 16, 3, 16, 129, 8, 16, 1, 17, 3, 17, 132, 8, 17, 1, 17,
		5, 17, 135, 8, 17, 10, 17, 12, 17, 138, 9, 17, 1, 18, 1, 18, 3, 18, 142,
		8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 3,
		21, 153, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 160, 8, 22, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 0, 0, 24, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 0, 5, 2, 0,
		14, 14, 16, 16, 1, 0, 14, 17, 1, 0, 21, 24, 1, 0, 18, 20, 1, 0, 25, 28,
		162, 0, 48, 1, 0, 0, 0, 2, 63, 1, 0, 0, 0, 4, 66, 1, 0, 0, 0, 6, 68, 1,
		0, 0, 0, 8, 74, 1, 0, 0, 0, 10, 78, 1, 0, 0, 0, 12, 80, 1, 0, 0, 0, 14,
		82, 1, 0, 0, 0, 16, 84, 1, 0, 0, 0, 18, 98, 1, 0, 0, 0, 20, 100, 1, 0,
		0, 0, 22, 102, 1, 0, 0, 0, 24, 104, 1, 0, 0, 0, 26, 106, 1, 0, 0, 0, 28,
		108, 1, 0, 0, 0, 30, 110, 1, 0, 0, 0, 32, 128, 1, 0, 0, 0, 34, 136, 1,
		0, 0, 0, 36, 141, 1, 0, 0, 0, 38, 143, 1, 0, 0, 0, 40, 147, 1, 0, 0, 0,
		42, 152, 1, 0, 0, 0, 44, 154, 1, 0, 0, 0, 46, 161, 1, 0, 0, 0, 48, 53,
		3, 2, 1, 0, 49, 50, 5, 11, 0, 0, 50, 52, 3, 2, 1, 0, 51, 49, 1, 0, 0, 0,
		52, 55, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 1, 1, 0,
		0, 0, 55, 53, 1, 0, 0, 0, 56, 57, 3, 6, 3, 0, 57, 58, 3, 4, 2, 0, 58, 64,
		1, 0, 0, 0, 59, 60, 3, 36, 18, 0, 60, 61, 3, 4, 2, 0, 61, 64, 1, 0, 0,
		0, 62, 64, 3, 4, 2, 0, 63, 56, 1, 0, 0, 0, 63, 59, 1, 0, 0, 0, 63, 62,
		1, 0, 0, 0, 64, 3, 1, 0, 0, 0, 65, 67, 5, 12, 0, 0, 66, 65, 1, 0, 0, 0,
		66, 67, 1, 0, 0, 0, 67, 5, 1, 0, 0, 0, 68, 69, 3, 8, 4, 0, 69, 70, 5, 1,
		0, 0, 70, 71, 3, 18, 9, 0, 71, 7, 1, 0, 0, 0, 72, 75, 3, 10, 5, 0, 73,
		75, 3, 16, 8, 0, 74, 72, 1, 0, 0, 0, 74, 73, 1, 0, 0, 0, 75, 9, 1, 0, 0,
		0, 76, 79, 3, 14, 7, 0, 77, 79, 3, 12, 6, 0, 78, 76, 1, 0, 0, 0, 78, 77,
		1, 0, 0, 0, 79, 11, 1, 0, 0, 0, 80, 81, 5, 29, 0, 0, 81, 13, 1, 0, 0, 0,
		82, 83, 7, 0, 0, 0, 83, 15, 1, 0, 0, 0, 84, 87, 3, 10, 5, 0, 85, 86, 5,
		2, 0, 0, 86, 88, 3, 10, 5, 0, 87, 85, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89,
		87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 17, 1, 0, 0, 0, 91, 99, 3, 20,
		10, 0, 92, 99, 3, 22, 11, 0, 93, 99, 3, 24, 12, 0, 94, 99, 3, 26, 13, 0,
		95, 99, 3, 28, 14, 0, 96, 99, 3, 30, 15, 0, 97, 99, 3, 40, 20, 0, 98, 91,
		1, 0, 0, 0, 98, 92, 1, 0, 0, 0, 98, 93, 1, 0, 0, 0, 98, 94, 1, 0, 0, 0,
		98, 95, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 97, 1, 0, 0, 0, 99, 19, 1,
		0, 0, 0, 100, 101, 7, 1, 0, 0, 101, 21, 1, 0, 0, 0, 102, 103, 7, 2, 0,
		0, 103, 23, 1, 0, 0, 0, 104, 105, 7, 3, 0, 0, 105, 25, 1, 0, 0, 0, 106,
		107, 5, 13, 0, 0, 107, 27, 1, 0, 0, 0, 108, 109, 7, 4, 0, 0, 109, 29, 1,
		0, 0, 0, 110, 112, 5, 3, 0, 0, 111, 113, 3, 32, 16, 0, 112, 111, 1, 0,
		0, 0, 112, 113, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 115, 3, 34, 17,
		0, 115, 116, 5, 4, 0, 0, 116, 31, 1, 0, 0, 0, 117, 118, 3, 34, 17, 0, 118,
		119, 3, 18, 9, 0, 119, 120, 5, 5, 0, 0, 120, 121, 3, 32, 16, 0, 121, 122,
		3, 34, 17, 0, 122, 129, 1, 0, 0, 0, 123, 124, 3, 34, 17, 0, 124, 126, 3,
		18, 9, 0, 125, 127, 5, 5, 0, 0, 126, 125, 1, 0, 0, 0, 126, 127, 1, 0, 0,
		0, 127, 129, 1, 0, 0, 0, 128, 117, 1, 0, 0, 0, 128, 123, 1, 0, 0, 0, 129,
		33, 1, 0, 0, 0, 130, 132, 5, 12, 0, 0, 131, 130, 1, 0, 0, 0, 131, 132,
		1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 135, 5, 11, 0, 0, 134, 131, 1, 0,
		0, 0, 135, 138, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0,
		137, 35, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 139, 142, 3, 38, 19, 0, 140,
		142, 3, 46, 23, 0, 141, 139, 1, 0, 0, 0, 141, 140, 1, 0, 0, 0, 142, 37,
		1, 0, 0, 0, 143, 144, 5, 3, 0, 0, 144, 145, 3, 8, 4, 0, 145, 146, 5, 4,
		0, 0, 146, 39, 1, 0, 0, 0, 147, 148, 5, 6, 0, 0, 148, 149, 3, 42, 21, 0,
		149, 150, 5, 7, 0, 0, 150, 41, 1, 0, 0, 0, 151, 153, 3, 44, 22, 0, 152,
		151, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 43, 1, 0, 0, 0, 154, 155, 3,
		8, 4, 0, 155, 156, 5, 1, 0, 0, 156, 159, 3, 18, 9, 0, 157, 158, 5, 5, 0,
		0, 158, 160, 3, 44, 22, 0, 159, 157, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0,
		160, 45, 1, 0, 0, 0, 161, 162, 5, 8, 0, 0, 162, 163, 3, 8, 4, 0, 163, 164,
		5, 9, 0, 0, 164, 47, 1, 0, 0, 0, 15, 53, 63, 66, 74, 78, 89, 98, 112, 126,
		128, 131, 136, 141, 152, 159,
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

// TomlParserInit initializes any static state used to implement TomlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTomlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TomlParserInit() {
	staticData := &tomlParserStaticData
	staticData.once.Do(tomlParserInit)
}

// NewTomlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTomlParser(input antlr.TokenStream) *TomlParser {
	TomlParserInit()
	this := new(TomlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &tomlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Toml.g4"

	return this
}

// TomlParser tokens.
const (
	TomlParserEOF               = antlr.TokenEOF
	TomlParserT__0              = 1
	TomlParserT__1              = 2
	TomlParserT__2              = 3
	TomlParserT__3              = 4
	TomlParserT__4              = 5
	TomlParserT__5              = 6
	TomlParserT__6              = 7
	TomlParserT__7              = 8
	TomlParserT__8              = 9
	TomlParserWS                = 10
	TomlParserNL                = 11
	TomlParserCOMMENT           = 12
	TomlParserBOOLEAN           = 13
	TomlParserBASIC_STRING      = 14
	TomlParserML_BASIC_STRING   = 15
	TomlParserLITERAL_STRING    = 16
	TomlParserML_LITERAL_STRING = 17
	TomlParserFLOAT             = 18
	TomlParserINF               = 19
	TomlParserNAN               = 20
	TomlParserDEC_INT           = 21
	TomlParserHEX_INT           = 22
	TomlParserOCT_INT           = 23
	TomlParserBIN_INT           = 24
	TomlParserOFFSET_DATE_TIME  = 25
	TomlParserLOCAL_DATE_TIME   = 26
	TomlParserLOCAL_DATE        = 27
	TomlParserLOCAL_TIME        = 28
	TomlParserUNQUOTED_KEY      = 29
)

// TomlParser rules.
const (
	TomlParserRULE_document                   = 0
	TomlParserRULE_expression                 = 1
	TomlParserRULE_comment                    = 2
	TomlParserRULE_keyValue                   = 3
	TomlParserRULE_key                        = 4
	TomlParserRULE_simpleKey                  = 5
	TomlParserRULE_unquotedKey                = 6
	TomlParserRULE_quotedKey                  = 7
	TomlParserRULE_dottedKey                  = 8
	TomlParserRULE_value                      = 9
	TomlParserRULE_stringValue                = 10
	TomlParserRULE_integer                    = 11
	TomlParserRULE_floatingPoint              = 12
	TomlParserRULE_boolean                    = 13
	TomlParserRULE_dateTime                   = 14
	TomlParserRULE_array                      = 15
	TomlParserRULE_arrayValues                = 16
	TomlParserRULE_commentOrNl                = 17
	TomlParserRULE_table                      = 18
	TomlParserRULE_standardTable              = 19
	TomlParserRULE_inlineTable                = 20
	TomlParserRULE_inlineTableKeyvals         = 21
	TomlParserRULE_inlineTableKeyvalsNonEmpty = 22
	TomlParserRULE_arrayTable                 = 23
)

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TomlParserRULE_document
	return p
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *DocumentContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *DocumentContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(TomlParserNL)
}

func (s *DocumentContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(TomlParserNL, i)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (s *DocumentContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case TomlVisitor:
		return t.VisitDocument(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TomlParser) Document() (localctx IDocumentContext) {
	this := p
	_ = this

	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TomlParserRULE_document)
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
		p.SetState(48)
		p.Expression()
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TomlParserNL {
		{
			p.SetState(49)
			p.Match(TomlParserNL)
		}
		{
			p.SetState(50)
			p.Expression()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KeyValue() IKeyValueContext
	Comment() ICommentContext
	Table() ITableContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TomlParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) KeyValue() IKeyValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyValueContext)
}

func (s *ExpressionContext) Comment() ICommentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *ExpressionContext) Table() ITableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case TomlVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TomlParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TomlParserRULE_expression)

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

	p.SetState(63)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TomlParserBASIC_STRING, TomlParserLITERAL_STRING, TomlParserUNQUOTED_KEY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(56)
			p.KeyValue()
		}
		{
			p.SetState(57)
			p.Comment()
		}

	case TomlParserT__2, TomlParserT__7:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(59)
			p.Table()
		}
		{
			p.SetState(60)
			p.Comment()
		}

	case TomlParserEOF, TomlParserNL, TomlParserCOMMENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(62)
			p.Comment()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMENT() antlr.TerminalNode

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TomlParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(TomlParserCOMMENT, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.ExitComment(s)
	}
}

func (s *CommentContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case TomlVisitor:
		return t.VisitComment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TomlParser) Comment() (localctx ICommentContext) {
	this := p
	_ = this

	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TomlParserRULE_comment)
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

	if _la == TomlParserCOMMENT {
		{
			p.SetState(65)
			p.Match(TomlParserCOMMENT)
		}

	}

	return localctx
}

// IKeyValueContext is an interface to support dynamic dispatch.
type IKeyValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Key() IKeyContext
	Value() IValueContext

	// IsKeyValueContext differentiates from other interfaces.
	IsKeyValueContext()
}

type KeyValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyValueContext() *KeyValueContext {
	var p = new(KeyValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TomlParserRULE_keyValue
	return p
}

func (*KeyValueContext) IsKeyValueContext() {}

func NewKeyValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyValueContext {
	var p = new(KeyValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_keyValue

	return p
}

func (s *KeyValueContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyValueContext) Key() IKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *KeyValueContext) Value() IValueContext {
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

func (s *KeyValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.EnterKeyValue(s)
	}
}

func (s *KeyValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.ExitKeyValue(s)
	}
}

func (s *KeyValueContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case TomlVisitor:
		return t.VisitKeyValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TomlParser) KeyValue() (localctx IKeyValueContext) {
	this := p
	_ = this

	localctx = NewKeyValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TomlParserRULE_keyValue)

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
		p.SetState(68)
		p.Key()
	}
	{
		p.SetState(69)
		p.Match(TomlParserT__0)
	}
	{
		p.SetState(70)
		p.Value()
	}

	return localctx
}

// IKeyContext is an interface to support dynamic dispatch.
type IKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SimpleKey() ISimpleKeyContext
	DottedKey() IDottedKeyContext

	// IsKeyContext differentiates from other interfaces.
	IsKeyContext()
}

type KeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyContext() *KeyContext {
	var p = new(KeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TomlParserRULE_key
	return p
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyContext) SimpleKey() ISimpleKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleKeyContext)
}

func (s *KeyContext) DottedKey() IDottedKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDottedKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDottedKeyContext)
}

func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.EnterKey(s)
	}
}

func (s *KeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.ExitKey(s)
	}
}

func (s *KeyContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case TomlVisitor:
		return t.VisitKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TomlParser) Key() (localctx IKeyContext) {
	this := p
	_ = this

	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TomlParserRULE_key)

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

	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)
			p.SimpleKey()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.DottedKey()
		}

	}

	return localctx
}

// ISimpleKeyContext is an interface to support dynamic dispatch.
type ISimpleKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QuotedKey() IQuotedKeyContext
	UnquotedKey() IUnquotedKeyContext

	// IsSimpleKeyContext differentiates from other interfaces.
	IsSimpleKeyContext()
}

type SimpleKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleKeyContext() *SimpleKeyContext {
	var p = new(SimpleKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TomlParserRULE_simpleKey
	return p
}

func (*SimpleKeyContext) IsSimpleKeyContext() {}

func NewSimpleKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleKeyContext {
	var p = new(SimpleKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_simpleKey

	return p
}

func (s *SimpleKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleKeyContext) QuotedKey() IQuotedKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQuotedKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQuotedKeyContext)
}

func (s *SimpleKeyContext) UnquotedKey() IUnquotedKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnquotedKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnquotedKeyContext)
}

func (s *SimpleKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.EnterSimpleKey(s)
	}
}

func (s *SimpleKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.ExitSimpleKey(s)
	}
}

func (s *SimpleKeyContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case TomlVisitor:
		return t.VisitSimpleKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TomlParser) SimpleKey() (localctx ISimpleKeyContext) {
	this := p
	_ = this

	localctx = NewSimpleKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TomlParserRULE_simpleKey)

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

	p.SetState(78)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TomlParserBASIC_STRING, TomlParserLITERAL_STRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)
			p.QuotedKey()
		}

	case TomlParserUNQUOTED_KEY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)
			p.UnquotedKey()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IUnquotedKeyContext is an interface to support dynamic dispatch.
type IUnquotedKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UNQUOTED_KEY() antlr.TerminalNode

	// IsUnquotedKeyContext differentiates from other interfaces.
	IsUnquotedKeyContext()
}

type UnquotedKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnquotedKeyContext() *UnquotedKeyContext {
	var p = new(UnquotedKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TomlParserRULE_unquotedKey
	return p
}

func (*UnquotedKeyContext) IsUnquotedKeyContext() {}

func NewUnquotedKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnquotedKeyContext {
	var p = new(UnquotedKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_unquotedKey

	return p
}

func (s *UnquotedKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *UnquotedKeyContext) UNQUOTED_KEY() antlr.TerminalNode {
	return s.GetToken(TomlParserUNQUOTED_KEY, 0)
}

func (s *UnquotedKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnquotedKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnquotedKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.EnterUnquotedKey(s)
	}
}

func (s *UnquotedKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.ExitUnquotedKey(s)
	}
}

func (s *UnquotedKeyContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case TomlVisitor:
		return t.VisitUnquotedKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TomlParser) UnquotedKey() (localctx IUnquotedKeyContext) {
	this := p
	_ = this

	localctx = NewUnquotedKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TomlParserRULE_unquotedKey)

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
		p.SetState(80)
		p.Match(TomlParserUNQUOTED_KEY)
	}

	return localctx
}

// IQuotedKeyContext is an interface to support dynamic dispatch.
type IQuotedKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BASIC_STRING() antlr.TerminalNode
	LITERAL_STRING() antlr.TerminalNode

	// IsQuotedKeyContext differentiates from other interfaces.
	IsQuotedKeyContext()
}

type QuotedKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuotedKeyContext() *QuotedKeyContext {
	var p = new(QuotedKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TomlParserRULE_quotedKey
	return p
}

func (*QuotedKeyContext) IsQuotedKeyContext() {}

func NewQuotedKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuotedKeyContext {
	var p = new(QuotedKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_quotedKey

	return p
}

func (s *QuotedKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *QuotedKeyContext) BASIC_STRING() antlr.TerminalNode {
	return s.GetToken(TomlParserBASIC_STRING, 0)
}

func (s *QuotedKeyContext) LITERAL_STRING() antlr.TerminalNode {
	return s.GetToken(TomlParserLITERAL_STRING, 0)
}

func (s *QuotedKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuotedKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.EnterQuotedKey(s)
	}
}

func (s *QuotedKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.ExitQuotedKey(s)
	}
}

func (s *QuotedKeyContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case TomlVisitor:
		return t.VisitQuotedKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TomlParser) QuotedKey() (localctx IQuotedKeyContext) {
	this := p
	_ = this

	localctx = NewQuotedKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TomlParserRULE_quotedKey)
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
		p.SetState(82)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TomlParserBASIC_STRING || _la == TomlParserLITERAL_STRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDottedKeyContext is an interface to support dynamic dispatch.
type IDottedKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSimpleKey() []ISimpleKeyContext
	SimpleKey(i int) ISimpleKeyContext

	// IsDottedKeyContext differentiates from other interfaces.
	IsDottedKeyContext()
}

type DottedKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDottedKeyContext() *DottedKeyContext {
	var p = new(DottedKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TomlParserRULE_dottedKey
	return p
}

func (*DottedKeyContext) IsDottedKeyContext() {}

func NewDottedKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DottedKeyContext {
	var p = new(DottedKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_dottedKey

	return p
}

func (s *DottedKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *DottedKeyContext) AllSimpleKey() []ISimpleKeyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimpleKeyContext); ok {
			len++
		}
	}

	tst := make([]ISimpleKeyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimpleKeyContext); ok {
			tst[i] = t.(ISimpleKeyContext)
			i++
		}
	}

	return tst
}

func (s *DottedKeyContext) SimpleKey(i int) ISimpleKeyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleKeyContext); ok {
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

	return t.(ISimpleKeyContext)
}

func (s *DottedKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DottedKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DottedKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.EnterDottedKey(s)
	}
}

func (s *DottedKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.ExitDottedKey(s)
	}
}

func (s *DottedKeyContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case TomlVisitor:
		return t.VisitDottedKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TomlParser) DottedKey() (localctx IDottedKeyContext) {
	this := p
	_ = this

	localctx = NewDottedKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TomlParserRULE_dottedKey)
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
		p.SetState(84)
		p.SimpleKey()
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == TomlParserT__1 {
		{
			p.SetState(85)
			p.Match(TomlParserT__1)
		}
		{
			p.SetState(86)
			p.SimpleKey()
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StringValue() IStringValueContext
	Integer() IIntegerContext
	FloatingPoint() IFloatingPointContext
	Boolean() IBooleanContext
	DateTime() IDateTimeContext
	Array() IArrayContext
	InlineTable() IInlineTableContext

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
	p.RuleIndex = TomlParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

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

func (s *ValueContext) Integer() IIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *ValueContext) FloatingPoint() IFloatingPointContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatingPointContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloatingPointContext)
}

func (s *ValueContext) Boolean() IBooleanContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanContext)
}

func (s *ValueContext) DateTime() IDateTimeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateTimeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateTimeContext)
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

func (s *ValueContext) InlineTable() IInlineTableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlineTableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInlineTableContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case TomlVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TomlParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TomlParserRULE_value)

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

	p.SetState(98)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TomlParserBASIC_STRING, TomlParserML_BASIC_STRING, TomlParserLITERAL_STRING, TomlParserML_LITERAL_STRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(91)
			p.StringValue()
		}

	case TomlParserDEC_INT, TomlParserHEX_INT, TomlParserOCT_INT, TomlParserBIN_INT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(92)
			p.Integer()
		}

	case TomlParserFLOAT, TomlParserINF, TomlParserNAN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(93)
			p.FloatingPoint()
		}

	case TomlParserBOOLEAN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(94)
			p.Boolean()
		}

	case TomlParserOFFSET_DATE_TIME, TomlParserLOCAL_DATE_TIME, TomlParserLOCAL_DATE, TomlParserLOCAL_TIME:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(95)
			p.DateTime()
		}

	case TomlParserT__2:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(96)
			p.Array()
		}

	case TomlParserT__5:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(97)
			p.InlineTable()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStringValueContext is an interface to support dynamic dispatch.
type IStringValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BASIC_STRING() antlr.TerminalNode
	ML_BASIC_STRING() antlr.TerminalNode
	LITERAL_STRING() antlr.TerminalNode
	ML_LITERAL_STRING() antlr.TerminalNode

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
	p.RuleIndex = TomlParserRULE_stringValue
	return p
}

func (*StringValueContext) IsStringValueContext() {}

func NewStringValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringValueContext {
	var p = new(StringValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_stringValue

	return p
}

func (s *StringValueContext) GetParser() antlr.Parser { return s.parser }

func (s *StringValueContext) BASIC_STRING() antlr.TerminalNode {
	return s.GetToken(TomlParserBASIC_STRING, 0)
}

func (s *StringValueContext) ML_BASIC_STRING() antlr.TerminalNode {
	return s.GetToken(TomlParserML_BASIC_STRING, 0)
}

func (s *StringValueContext) LITERAL_STRING() antlr.TerminalNode {
	return s.GetToken(TomlParserLITERAL_STRING, 0)
}

func (s *StringValueContext) ML_LITERAL_STRING() antlr.TerminalNode {
	return s.GetToken(TomlParserML_LITERAL_STRING, 0)
}

func (s *StringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.EnterStringValue(s)
	}
}

func (s *StringValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.ExitStringValue(s)
	}
}

func (s *StringValueContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case TomlVisitor:
		return t.VisitStringValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TomlParser) StringValue() (localctx IStringValueContext) {
	this := p
	_ = this

	localctx = NewStringValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TomlParserRULE_stringValue)
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
		p.SetState(100)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&245760) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEC_INT() antlr.TerminalNode
	HEX_INT() antlr.TerminalNode
	OCT_INT() antlr.TerminalNode
	BIN_INT() antlr.TerminalNode

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TomlParserRULE_integer
	return p
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) DEC_INT() antlr.TerminalNode {
	return s.GetToken(TomlParserDEC_INT, 0)
}

func (s *IntegerContext) HEX_INT() antlr.TerminalNode {
	return s.GetToken(TomlParserHEX_INT, 0)
}

func (s *IntegerContext) OCT_INT() antlr.TerminalNode {
	return s.GetToken(TomlParserOCT_INT, 0)
}

func (s *IntegerContext) BIN_INT() antlr.TerminalNode {
	return s.GetToken(TomlParserBIN_INT, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (s *IntegerContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case TomlVisitor:
		return t.VisitInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TomlParser) Integer() (localctx IIntegerContext) {
	this := p
	_ = this

	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TomlParserRULE_integer)
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
		p.SetState(102)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&31457280) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFloatingPointContext is an interface to support dynamic dispatch.
type IFloatingPointContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FLOAT() antlr.TerminalNode
	INF() antlr.TerminalNode
	NAN() antlr.TerminalNode

	// IsFloatingPointContext differentiates from other interfaces.
	IsFloatingPointContext()
}

type FloatingPointContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatingPointContext() *FloatingPointContext {
	var p = new(FloatingPointContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TomlParserRULE_floatingPoint
	return p
}

func (*FloatingPointContext) IsFloatingPointContext() {}

func NewFloatingPointContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatingPointContext {
	var p = new(FloatingPointContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_floatingPoint

	return p
}

func (s *FloatingPointContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatingPointContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(TomlParserFLOAT, 0)
}

func (s *FloatingPointContext) INF() antlr.TerminalNode {
	return s.GetToken(TomlParserINF, 0)
}

func (s *FloatingPointContext) NAN() antlr.TerminalNode {
	return s.GetToken(TomlParserNAN, 0)
}

func (s *FloatingPointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatingPointContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatingPointContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.EnterFloatingPoint(s)
	}
}

func (s *FloatingPointContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.ExitFloatingPoint(s)
	}
}

func (s *FloatingPointContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case TomlVisitor:
		return t.VisitFloatingPoint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TomlParser) FloatingPoint() (localctx IFloatingPointContext) {
	this := p
	_ = this

	localctx = NewFloatingPointContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TomlParserRULE_floatingPoint)
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
		p.SetState(104)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1835008) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBooleanContext is an interface to support dynamic dispatch.
type IBooleanContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BOOLEAN() antlr.TerminalNode

	// IsBooleanContext differentiates from other interfaces.
	IsBooleanContext()
}

type BooleanContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanContext() *BooleanContext {
	var p = new(BooleanContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TomlParserRULE_boolean
	return p
}

func (*BooleanContext) IsBooleanContext() {}

func NewBooleanContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanContext {
	var p = new(BooleanContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_boolean

	return p
}

func (s *BooleanContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(TomlParserBOOLEAN, 0)
}

func (s *BooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.EnterBoolean(s)
	}
}

func (s *BooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.ExitBoolean(s)
	}
}

func (s *BooleanContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case TomlVisitor:
		return t.VisitBoolean(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TomlParser) Boolean() (localctx IBooleanContext) {
	this := p
	_ = this

	localctx = NewBooleanContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, TomlParserRULE_boolean)

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
		p.SetState(106)
		p.Match(TomlParserBOOLEAN)
	}

	return localctx
}

// IDateTimeContext is an interface to support dynamic dispatch.
type IDateTimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OFFSET_DATE_TIME() antlr.TerminalNode
	LOCAL_DATE_TIME() antlr.TerminalNode
	LOCAL_DATE() antlr.TerminalNode
	LOCAL_TIME() antlr.TerminalNode

	// IsDateTimeContext differentiates from other interfaces.
	IsDateTimeContext()
}

type DateTimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateTimeContext() *DateTimeContext {
	var p = new(DateTimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TomlParserRULE_dateTime
	return p
}

func (*DateTimeContext) IsDateTimeContext() {}

func NewDateTimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateTimeContext {
	var p = new(DateTimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_dateTime

	return p
}

func (s *DateTimeContext) GetParser() antlr.Parser { return s.parser }

func (s *DateTimeContext) OFFSET_DATE_TIME() antlr.TerminalNode {
	return s.GetToken(TomlParserOFFSET_DATE_TIME, 0)
}

func (s *DateTimeContext) LOCAL_DATE_TIME() antlr.TerminalNode {
	return s.GetToken(TomlParserLOCAL_DATE_TIME, 0)
}

func (s *DateTimeContext) LOCAL_DATE() antlr.TerminalNode {
	return s.GetToken(TomlParserLOCAL_DATE, 0)
}

func (s *DateTimeContext) LOCAL_TIME() antlr.TerminalNode {
	return s.GetToken(TomlParserLOCAL_TIME, 0)
}

func (s *DateTimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateTimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateTimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.EnterDateTime(s)
	}
}

func (s *DateTimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.ExitDateTime(s)
	}
}

func (s *DateTimeContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case TomlVisitor:
		return t.VisitDateTime(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TomlParser) DateTime() (localctx IDateTimeContext) {
	this := p
	_ = this

	localctx = NewDateTimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, TomlParserRULE_dateTime)
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
		p.SetState(108)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&503316480) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CommentOrNl() ICommentOrNlContext
	ArrayValues() IArrayValuesContext

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
	p.RuleIndex = TomlParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) CommentOrNl() ICommentOrNlContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentOrNlContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommentOrNlContext)
}

func (s *ArrayContext) ArrayValues() IArrayValuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayValuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayValuesContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.ExitArray(s)
	}
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case TomlVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TomlParser) Array() (localctx IArrayContext) {
	this := p
	_ = this

	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, TomlParserRULE_array)

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
		p.SetState(110)
		p.Match(TomlParserT__2)
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(111)
			p.ArrayValues()
		}

	}
	{
		p.SetState(114)
		p.CommentOrNl()
	}
	{
		p.SetState(115)
		p.Match(TomlParserT__3)
	}

	return localctx
}

// IArrayValuesContext is an interface to support dynamic dispatch.
type IArrayValuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCommentOrNl() []ICommentOrNlContext
	CommentOrNl(i int) ICommentOrNlContext
	Value() IValueContext
	ArrayValues() IArrayValuesContext

	// IsArrayValuesContext differentiates from other interfaces.
	IsArrayValuesContext()
}

type ArrayValuesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayValuesContext() *ArrayValuesContext {
	var p = new(ArrayValuesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TomlParserRULE_arrayValues
	return p
}

func (*ArrayValuesContext) IsArrayValuesContext() {}

func NewArrayValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayValuesContext {
	var p = new(ArrayValuesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_arrayValues

	return p
}

func (s *ArrayValuesContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayValuesContext) AllCommentOrNl() []ICommentOrNlContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICommentOrNlContext); ok {
			len++
		}
	}

	tst := make([]ICommentOrNlContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICommentOrNlContext); ok {
			tst[i] = t.(ICommentOrNlContext)
			i++
		}
	}

	return tst
}

func (s *ArrayValuesContext) CommentOrNl(i int) ICommentOrNlContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentOrNlContext); ok {
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

	return t.(ICommentOrNlContext)
}

func (s *ArrayValuesContext) Value() IValueContext {
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

func (s *ArrayValuesContext) ArrayValues() IArrayValuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayValuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayValuesContext)
}

func (s *ArrayValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.EnterArrayValues(s)
	}
}

func (s *ArrayValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.ExitArrayValues(s)
	}
}

func (s *ArrayValuesContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case TomlVisitor:
		return t.VisitArrayValues(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TomlParser) ArrayValues() (localctx IArrayValuesContext) {
	this := p
	_ = this

	localctx = NewArrayValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, TomlParserRULE_arrayValues)
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

	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(117)
			p.CommentOrNl()
		}
		{
			p.SetState(118)
			p.Value()
		}
		{
			p.SetState(119)
			p.Match(TomlParserT__4)
		}
		{
			p.SetState(120)
			p.ArrayValues()
		}
		{
			p.SetState(121)
			p.CommentOrNl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(123)
			p.CommentOrNl()
		}
		{
			p.SetState(124)
			p.Value()
		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TomlParserT__4 {
			{
				p.SetState(125)
				p.Match(TomlParserT__4)
			}

		}

	}

	return localctx
}

// ICommentOrNlContext is an interface to support dynamic dispatch.
type ICommentOrNlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllCOMMENT() []antlr.TerminalNode
	COMMENT(i int) antlr.TerminalNode

	// IsCommentOrNlContext differentiates from other interfaces.
	IsCommentOrNlContext()
}

type CommentOrNlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentOrNlContext() *CommentOrNlContext {
	var p = new(CommentOrNlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TomlParserRULE_commentOrNl
	return p
}

func (*CommentOrNlContext) IsCommentOrNlContext() {}

func NewCommentOrNlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentOrNlContext {
	var p = new(CommentOrNlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_commentOrNl

	return p
}

func (s *CommentOrNlContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentOrNlContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(TomlParserNL)
}

func (s *CommentOrNlContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(TomlParserNL, i)
}

func (s *CommentOrNlContext) AllCOMMENT() []antlr.TerminalNode {
	return s.GetTokens(TomlParserCOMMENT)
}

func (s *CommentOrNlContext) COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(TomlParserCOMMENT, i)
}

func (s *CommentOrNlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentOrNlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentOrNlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.EnterCommentOrNl(s)
	}
}

func (s *CommentOrNlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.ExitCommentOrNl(s)
	}
}

func (s *CommentOrNlContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case TomlVisitor:
		return t.VisitCommentOrNl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TomlParser) CommentOrNl() (localctx ICommentOrNlContext) {
	this := p
	_ = this

	localctx = NewCommentOrNlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, TomlParserRULE_commentOrNl)
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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(131)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == TomlParserCOMMENT {
				{
					p.SetState(130)
					p.Match(TomlParserCOMMENT)
				}

			}
			{
				p.SetState(133)
				p.Match(TomlParserNL)
			}

		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// ITableContext is an interface to support dynamic dispatch.
type ITableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StandardTable() IStandardTableContext
	ArrayTable() IArrayTableContext

	// IsTableContext differentiates from other interfaces.
	IsTableContext()
}

type TableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableContext() *TableContext {
	var p = new(TableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TomlParserRULE_table
	return p
}

func (*TableContext) IsTableContext() {}

func NewTableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableContext {
	var p = new(TableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_table

	return p
}

func (s *TableContext) GetParser() antlr.Parser { return s.parser }

func (s *TableContext) StandardTable() IStandardTableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStandardTableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStandardTableContext)
}

func (s *TableContext) ArrayTable() IArrayTableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayTableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayTableContext)
}

func (s *TableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.EnterTable(s)
	}
}

func (s *TableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.ExitTable(s)
	}
}

func (s *TableContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case TomlVisitor:
		return t.VisitTable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TomlParser) Table() (localctx ITableContext) {
	this := p
	_ = this

	localctx = NewTableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, TomlParserRULE_table)

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

	p.SetState(141)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TomlParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(139)
			p.StandardTable()
		}

	case TomlParserT__7:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(140)
			p.ArrayTable()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStandardTableContext is an interface to support dynamic dispatch.
type IStandardTableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Key() IKeyContext

	// IsStandardTableContext differentiates from other interfaces.
	IsStandardTableContext()
}

type StandardTableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStandardTableContext() *StandardTableContext {
	var p = new(StandardTableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TomlParserRULE_standardTable
	return p
}

func (*StandardTableContext) IsStandardTableContext() {}

func NewStandardTableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StandardTableContext {
	var p = new(StandardTableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_standardTable

	return p
}

func (s *StandardTableContext) GetParser() antlr.Parser { return s.parser }

func (s *StandardTableContext) Key() IKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *StandardTableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StandardTableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StandardTableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.EnterStandardTable(s)
	}
}

func (s *StandardTableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.ExitStandardTable(s)
	}
}

func (s *StandardTableContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case TomlVisitor:
		return t.VisitStandardTable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TomlParser) StandardTable() (localctx IStandardTableContext) {
	this := p
	_ = this

	localctx = NewStandardTableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, TomlParserRULE_standardTable)

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
		p.SetState(143)
		p.Match(TomlParserT__2)
	}
	{
		p.SetState(144)
		p.Key()
	}
	{
		p.SetState(145)
		p.Match(TomlParserT__3)
	}

	return localctx
}

// IInlineTableContext is an interface to support dynamic dispatch.
type IInlineTableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	InlineTableKeyvals() IInlineTableKeyvalsContext

	// IsInlineTableContext differentiates from other interfaces.
	IsInlineTableContext()
}

type InlineTableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineTableContext() *InlineTableContext {
	var p = new(InlineTableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TomlParserRULE_inlineTable
	return p
}

func (*InlineTableContext) IsInlineTableContext() {}

func NewInlineTableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineTableContext {
	var p = new(InlineTableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_inlineTable

	return p
}

func (s *InlineTableContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineTableContext) InlineTableKeyvals() IInlineTableKeyvalsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlineTableKeyvalsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInlineTableKeyvalsContext)
}

func (s *InlineTableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineTableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InlineTableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.EnterInlineTable(s)
	}
}

func (s *InlineTableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.ExitInlineTable(s)
	}
}

func (s *InlineTableContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case TomlVisitor:
		return t.VisitInlineTable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TomlParser) InlineTable() (localctx IInlineTableContext) {
	this := p
	_ = this

	localctx = NewInlineTableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, TomlParserRULE_inlineTable)

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
		p.SetState(147)
		p.Match(TomlParserT__5)
	}
	{
		p.SetState(148)
		p.InlineTableKeyvals()
	}
	{
		p.SetState(149)
		p.Match(TomlParserT__6)
	}

	return localctx
}

// IInlineTableKeyvalsContext is an interface to support dynamic dispatch.
type IInlineTableKeyvalsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	InlineTableKeyvalsNonEmpty() IInlineTableKeyvalsNonEmptyContext

	// IsInlineTableKeyvalsContext differentiates from other interfaces.
	IsInlineTableKeyvalsContext()
}

type InlineTableKeyvalsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineTableKeyvalsContext() *InlineTableKeyvalsContext {
	var p = new(InlineTableKeyvalsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TomlParserRULE_inlineTableKeyvals
	return p
}

func (*InlineTableKeyvalsContext) IsInlineTableKeyvalsContext() {}

func NewInlineTableKeyvalsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineTableKeyvalsContext {
	var p = new(InlineTableKeyvalsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_inlineTableKeyvals

	return p
}

func (s *InlineTableKeyvalsContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineTableKeyvalsContext) InlineTableKeyvalsNonEmpty() IInlineTableKeyvalsNonEmptyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlineTableKeyvalsNonEmptyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInlineTableKeyvalsNonEmptyContext)
}

func (s *InlineTableKeyvalsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineTableKeyvalsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InlineTableKeyvalsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.EnterInlineTableKeyvals(s)
	}
}

func (s *InlineTableKeyvalsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.ExitInlineTableKeyvals(s)
	}
}

func (s *InlineTableKeyvalsContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case TomlVisitor:
		return t.VisitInlineTableKeyvals(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TomlParser) InlineTableKeyvals() (localctx IInlineTableKeyvalsContext) {
	this := p
	_ = this

	localctx = NewInlineTableKeyvalsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, TomlParserRULE_inlineTableKeyvals)
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
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&536952832) != 0 {
		{
			p.SetState(151)
			p.InlineTableKeyvalsNonEmpty()
		}

	}

	return localctx
}

// IInlineTableKeyvalsNonEmptyContext is an interface to support dynamic dispatch.
type IInlineTableKeyvalsNonEmptyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Key() IKeyContext
	Value() IValueContext
	InlineTableKeyvalsNonEmpty() IInlineTableKeyvalsNonEmptyContext

	// IsInlineTableKeyvalsNonEmptyContext differentiates from other interfaces.
	IsInlineTableKeyvalsNonEmptyContext()
}

type InlineTableKeyvalsNonEmptyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineTableKeyvalsNonEmptyContext() *InlineTableKeyvalsNonEmptyContext {
	var p = new(InlineTableKeyvalsNonEmptyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TomlParserRULE_inlineTableKeyvalsNonEmpty
	return p
}

func (*InlineTableKeyvalsNonEmptyContext) IsInlineTableKeyvalsNonEmptyContext() {}

func NewInlineTableKeyvalsNonEmptyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineTableKeyvalsNonEmptyContext {
	var p = new(InlineTableKeyvalsNonEmptyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_inlineTableKeyvalsNonEmpty

	return p
}

func (s *InlineTableKeyvalsNonEmptyContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineTableKeyvalsNonEmptyContext) Key() IKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *InlineTableKeyvalsNonEmptyContext) Value() IValueContext {
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

func (s *InlineTableKeyvalsNonEmptyContext) InlineTableKeyvalsNonEmpty() IInlineTableKeyvalsNonEmptyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlineTableKeyvalsNonEmptyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInlineTableKeyvalsNonEmptyContext)
}

func (s *InlineTableKeyvalsNonEmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineTableKeyvalsNonEmptyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InlineTableKeyvalsNonEmptyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.EnterInlineTableKeyvalsNonEmpty(s)
	}
}

func (s *InlineTableKeyvalsNonEmptyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.ExitInlineTableKeyvalsNonEmpty(s)
	}
}

func (s *InlineTableKeyvalsNonEmptyContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case TomlVisitor:
		return t.VisitInlineTableKeyvalsNonEmpty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TomlParser) InlineTableKeyvalsNonEmpty() (localctx IInlineTableKeyvalsNonEmptyContext) {
	this := p
	_ = this

	localctx = NewInlineTableKeyvalsNonEmptyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, TomlParserRULE_inlineTableKeyvalsNonEmpty)
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
		p.SetState(154)
		p.Key()
	}
	{
		p.SetState(155)
		p.Match(TomlParserT__0)
	}
	{
		p.SetState(156)
		p.Value()
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TomlParserT__4 {
		{
			p.SetState(157)
			p.Match(TomlParserT__4)
		}
		{
			p.SetState(158)
			p.InlineTableKeyvalsNonEmpty()
		}

	}

	return localctx
}

// IArrayTableContext is an interface to support dynamic dispatch.
type IArrayTableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Key() IKeyContext

	// IsArrayTableContext differentiates from other interfaces.
	IsArrayTableContext()
}

type ArrayTableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayTableContext() *ArrayTableContext {
	var p = new(ArrayTableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TomlParserRULE_arrayTable
	return p
}

func (*ArrayTableContext) IsArrayTableContext() {}

func NewArrayTableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayTableContext {
	var p = new(ArrayTableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_arrayTable

	return p
}

func (s *ArrayTableContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayTableContext) Key() IKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *ArrayTableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayTableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayTableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.EnterArrayTable(s)
	}
}

func (s *ArrayTableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlListener); ok {
		listenerT.ExitArrayTable(s)
	}
}

func (s *ArrayTableContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case TomlVisitor:
		return t.VisitArrayTable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TomlParser) ArrayTable() (localctx IArrayTableContext) {
	this := p
	_ = this

	localctx = NewArrayTableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, TomlParserRULE_arrayTable)

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
		p.SetState(161)
		p.Match(TomlParserT__7)
	}
	{
		p.SetState(162)
		p.Key()
	}
	{
		p.SetState(163)
		p.Match(TomlParserT__8)
	}

	return localctx
}
