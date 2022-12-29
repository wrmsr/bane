// Code generated from Hocon.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Hocon

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

type HoconParser struct {
	*antlr.BaseParser
}

var hoconParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func hoconParserInit() {
	staticData := &hoconParserStaticData
	staticData.literalNames = []string{
		"", "','", "'{'", "'}'", "'.'", "'['", "']'", "'-'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "COMMENT", "NUMBER", "STRING", "PATH_ELEMENT",
		"REFERENCE", "KV", "WS",
	}
	staticData.ruleNames = []string{
		"hocon", "prop", "obj", "objectBegin", "objectEnd", "objectData", "arrayData",
		"stringData", "referenceData", "numberData", "key", "path", "arrayBegin",
		"arrayEnd", "array", "arrayValue", "arrayString", "arrayReference",
		"arrayNumber", "arrayObj", "arrayArray", "stringValue", "rawstring",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 14, 197, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 1, 0, 5, 0, 48, 8, 0, 10, 0, 12, 0, 51, 9, 0,
		1, 0, 5, 0, 54, 8, 0, 10, 0, 12, 0, 57, 9, 0, 1, 0, 5, 0, 60, 8, 0, 10,
		0, 12, 0, 63, 9, 0, 3, 0, 65, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1,
		72, 8, 1, 1, 2, 1, 2, 1, 2, 3, 2, 77, 8, 2, 1, 2, 5, 2, 80, 8, 2, 10, 2,
		12, 2, 83, 9, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 90, 8, 2, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 5, 1, 5, 3, 5, 98, 8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 3, 10, 120, 8, 10, 1, 11, 1, 11, 1, 11, 5, 11, 125,
		8, 11, 10, 11, 12, 11, 128, 9, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 14, 3, 14, 137, 8, 14, 1, 14, 5, 14, 140, 8, 14, 10, 14, 12, 14,
		143, 9, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 150, 8, 14, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 157, 8, 15, 1, 16, 1, 16, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 5, 21, 171,
		8, 21, 10, 21, 12, 21, 174, 9, 21, 1, 21, 1, 21, 5, 21, 178, 8, 21, 10,
		21, 12, 21, 181, 9, 21, 1, 21, 1, 21, 5, 21, 185, 8, 21, 10, 21, 12, 21,
		188, 9, 21, 3, 21, 190, 8, 21, 1, 22, 4, 22, 193, 8, 22, 11, 22, 12, 22,
		194, 1, 22, 0, 0, 23, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
		28, 30, 32, 34, 36, 38, 40, 42, 44, 0, 1, 2, 0, 7, 7, 11, 11, 201, 0, 64,
		1, 0, 0, 0, 2, 71, 1, 0, 0, 0, 4, 89, 1, 0, 0, 0, 6, 91, 1, 0, 0, 0, 8,
		93, 1, 0, 0, 0, 10, 95, 1, 0, 0, 0, 12, 101, 1, 0, 0, 0, 14, 105, 1, 0,
		0, 0, 16, 109, 1, 0, 0, 0, 18, 113, 1, 0, 0, 0, 20, 119, 1, 0, 0, 0, 22,
		121, 1, 0, 0, 0, 24, 129, 1, 0, 0, 0, 26, 131, 1, 0, 0, 0, 28, 149, 1,
		0, 0, 0, 30, 156, 1, 0, 0, 0, 32, 158, 1, 0, 0, 0, 34, 160, 1, 0, 0, 0,
		36, 162, 1, 0, 0, 0, 38, 164, 1, 0, 0, 0, 40, 166, 1, 0, 0, 0, 42, 189,
		1, 0, 0, 0, 44, 192, 1, 0, 0, 0, 46, 48, 3, 4, 2, 0, 47, 46, 1, 0, 0, 0,
		48, 51, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 65, 1,
		0, 0, 0, 51, 49, 1, 0, 0, 0, 52, 54, 3, 28, 14, 0, 53, 52, 1, 0, 0, 0,
		54, 57, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 65, 1,
		0, 0, 0, 57, 55, 1, 0, 0, 0, 58, 60, 3, 2, 1, 0, 59, 58, 1, 0, 0, 0, 60,
		63, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 65, 1, 0, 0,
		0, 63, 61, 1, 0, 0, 0, 64, 49, 1, 0, 0, 0, 64, 55, 1, 0, 0, 0, 64, 61,
		1, 0, 0, 0, 65, 1, 1, 0, 0, 0, 66, 72, 3, 10, 5, 0, 67, 72, 3, 12, 6, 0,
		68, 72, 3, 14, 7, 0, 69, 72, 3, 16, 8, 0, 70, 72, 3, 18, 9, 0, 71, 66,
		1, 0, 0, 0, 71, 67, 1, 0, 0, 0, 71, 68, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0,
		71, 70, 1, 0, 0, 0, 72, 3, 1, 0, 0, 0, 73, 74, 3, 6, 3, 0, 74, 81, 3, 2,
		1, 0, 75, 77, 5, 1, 0, 0, 76, 75, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 78,
		1, 0, 0, 0, 78, 80, 3, 2, 1, 0, 79, 76, 1, 0, 0, 0, 80, 83, 1, 0, 0, 0,
		81, 79, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 84, 1, 0, 0, 0, 83, 81, 1,
		0, 0, 0, 84, 85, 3, 8, 4, 0, 85, 90, 1, 0, 0, 0, 86, 87, 3, 6, 3, 0, 87,
		88, 3, 8, 4, 0, 88, 90, 1, 0, 0, 0, 89, 73, 1, 0, 0, 0, 89, 86, 1, 0, 0,
		0, 90, 5, 1, 0, 0, 0, 91, 92, 5, 2, 0, 0, 92, 7, 1, 0, 0, 0, 93, 94, 5,
		3, 0, 0, 94, 9, 1, 0, 0, 0, 95, 97, 3, 20, 10, 0, 96, 98, 5, 13, 0, 0,
		97, 96, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 100, 3,
		4, 2, 0, 100, 11, 1, 0, 0, 0, 101, 102, 3, 20, 10, 0, 102, 103, 5, 13,
		0, 0, 103, 104, 3, 28, 14, 0, 104, 13, 1, 0, 0, 0, 105, 106, 3, 20, 10,
		0, 106, 107, 5, 13, 0, 0, 107, 108, 3, 42, 21, 0, 108, 15, 1, 0, 0, 0,
		109, 110, 3, 20, 10, 0, 110, 111, 5, 13, 0, 0, 111, 112, 5, 12, 0, 0, 112,
		17, 1, 0, 0, 0, 113, 114, 3, 20, 10, 0, 114, 115, 5, 13, 0, 0, 115, 116,
		5, 9, 0, 0, 116, 19, 1, 0, 0, 0, 117, 120, 3, 22, 11, 0, 118, 120, 5, 10,
		0, 0, 119, 117, 1, 0, 0, 0, 119, 118, 1, 0, 0, 0, 120, 21, 1, 0, 0, 0,
		121, 126, 5, 11, 0, 0, 122, 123, 5, 4, 0, 0, 123, 125, 5, 11, 0, 0, 124,
		122, 1, 0, 0, 0, 125, 128, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 126, 127,
		1, 0, 0, 0, 127, 23, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 129, 130, 5, 5,
		0, 0, 130, 25, 1, 0, 0, 0, 131, 132, 5, 6, 0, 0, 132, 27, 1, 0, 0, 0, 133,
		134, 3, 24, 12, 0, 134, 141, 3, 30, 15, 0, 135, 137, 5, 1, 0, 0, 136, 135,
		1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 140, 3, 30,
		15, 0, 139, 136, 1, 0, 0, 0, 140, 143, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0,
		141, 142, 1, 0, 0, 0, 142, 144, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 144,
		145, 3, 26, 13, 0, 145, 150, 1, 0, 0, 0, 146, 147, 3, 24, 12, 0, 147, 148,
		3, 26, 13, 0, 148, 150, 1, 0, 0, 0, 149, 133, 1, 0, 0, 0, 149, 146, 1,
		0, 0, 0, 150, 29, 1, 0, 0, 0, 151, 157, 3, 32, 16, 0, 152, 157, 3, 34,
		17, 0, 153, 157, 3, 36, 18, 0, 154, 157, 3, 38, 19, 0, 155, 157, 3, 40,
		20, 0, 156, 151, 1, 0, 0, 0, 156, 152, 1, 0, 0, 0, 156, 153, 1, 0, 0, 0,
		156, 154, 1, 0, 0, 0, 156, 155, 1, 0, 0, 0, 157, 31, 1, 0, 0, 0, 158, 159,
		3, 42, 21, 0, 159, 33, 1, 0, 0, 0, 160, 161, 5, 12, 0, 0, 161, 35, 1, 0,
		0, 0, 162, 163, 5, 9, 0, 0, 163, 37, 1, 0, 0, 0, 164, 165, 3, 4, 2, 0,
		165, 39, 1, 0, 0, 0, 166, 167, 3, 28, 14, 0, 167, 41, 1, 0, 0, 0, 168,
		172, 5, 10, 0, 0, 169, 171, 3, 42, 21, 0, 170, 169, 1, 0, 0, 0, 171, 174,
		1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 190, 1, 0,
		0, 0, 174, 172, 1, 0, 0, 0, 175, 179, 3, 44, 22, 0, 176, 178, 3, 42, 21,
		0, 177, 176, 1, 0, 0, 0, 178, 181, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 179,
		180, 1, 0, 0, 0, 180, 190, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 182, 186,
		5, 12, 0, 0, 183, 185, 3, 42, 21, 0, 184, 183, 1, 0, 0, 0, 185, 188, 1,
		0, 0, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 190, 1, 0, 0,
		0, 188, 186, 1, 0, 0, 0, 189, 168, 1, 0, 0, 0, 189, 175, 1, 0, 0, 0, 189,
		182, 1, 0, 0, 0, 190, 43, 1, 0, 0, 0, 191, 193, 7, 0, 0, 0, 192, 191, 1,
		0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0,
		0, 195, 45, 1, 0, 0, 0, 20, 49, 55, 61, 64, 71, 76, 81, 89, 97, 119, 126,
		136, 141, 149, 156, 172, 179, 186, 189, 194,
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

// HoconParserInit initializes any static state used to implement HoconParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewHoconParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func HoconParserInit() {
	staticData := &hoconParserStaticData
	staticData.once.Do(hoconParserInit)
}

// NewHoconParser produces a new parser instance for the optional input antlr.TokenStream.
func NewHoconParser(input antlr.TokenStream) *HoconParser {
	HoconParserInit()
	this := new(HoconParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &hoconParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Hocon.g4"

	return this
}

// HoconParser tokens.
const (
	HoconParserEOF          = antlr.TokenEOF
	HoconParserT__0         = 1
	HoconParserT__1         = 2
	HoconParserT__2         = 3
	HoconParserT__3         = 4
	HoconParserT__4         = 5
	HoconParserT__5         = 6
	HoconParserT__6         = 7
	HoconParserCOMMENT      = 8
	HoconParserNUMBER       = 9
	HoconParserSTRING       = 10
	HoconParserPATH_ELEMENT = 11
	HoconParserREFERENCE    = 12
	HoconParserKV           = 13
	HoconParserWS           = 14
)

// HoconParser rules.
const (
	HoconParserRULE_hocon          = 0
	HoconParserRULE_prop           = 1
	HoconParserRULE_obj            = 2
	HoconParserRULE_objectBegin    = 3
	HoconParserRULE_objectEnd      = 4
	HoconParserRULE_objectData     = 5
	HoconParserRULE_arrayData      = 6
	HoconParserRULE_stringData     = 7
	HoconParserRULE_referenceData  = 8
	HoconParserRULE_numberData     = 9
	HoconParserRULE_key            = 10
	HoconParserRULE_path           = 11
	HoconParserRULE_arrayBegin     = 12
	HoconParserRULE_arrayEnd       = 13
	HoconParserRULE_array          = 14
	HoconParserRULE_arrayValue     = 15
	HoconParserRULE_arrayString    = 16
	HoconParserRULE_arrayReference = 17
	HoconParserRULE_arrayNumber    = 18
	HoconParserRULE_arrayObj       = 19
	HoconParserRULE_arrayArray     = 20
	HoconParserRULE_stringValue    = 21
	HoconParserRULE_rawstring      = 22
)

// IHoconContext is an interface to support dynamic dispatch.
type IHoconContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHoconContext differentiates from other interfaces.
	IsHoconContext()
}

type HoconContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHoconContext() *HoconContext {
	var p = new(HoconContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HoconParserRULE_hocon
	return p
}

func (*HoconContext) IsHoconContext() {}

func NewHoconContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HoconContext {
	var p = new(HoconContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HoconParserRULE_hocon

	return p
}

func (s *HoconContext) GetParser() antlr.Parser { return s.parser }

func (s *HoconContext) AllObj() []IObjContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IObjContext); ok {
			len++
		}
	}

	tst := make([]IObjContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IObjContext); ok {
			tst[i] = t.(IObjContext)
			i++
		}
	}

	return tst
}

func (s *HoconContext) Obj(i int) IObjContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjContext); ok {
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

	return t.(IObjContext)
}

func (s *HoconContext) AllArray() []IArrayContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArrayContext); ok {
			len++
		}
	}

	tst := make([]IArrayContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArrayContext); ok {
			tst[i] = t.(IArrayContext)
			i++
		}
	}

	return tst
}

func (s *HoconContext) Array(i int) IArrayContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
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

	return t.(IArrayContext)
}

func (s *HoconContext) AllProp() []IPropContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPropContext); ok {
			len++
		}
	}

	tst := make([]IPropContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPropContext); ok {
			tst[i] = t.(IPropContext)
			i++
		}
	}

	return tst
}

func (s *HoconContext) Prop(i int) IPropContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropContext); ok {
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

	return t.(IPropContext)
}

func (s *HoconContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HoconContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HoconContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.EnterHocon(s)
	}
}

func (s *HoconContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.ExitHocon(s)
	}
}

func (s *HoconContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HoconVisitor:
		return t.VisitHocon(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HoconParser) Hocon() (localctx IHoconContext) {
	this := p
	_ = this

	localctx = NewHoconContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HoconParserRULE_hocon)
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

	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == HoconParserT__1 {
			{
				p.SetState(46)
				p.Obj()
			}

			p.SetState(51)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == HoconParserT__4 {
			{
				p.SetState(52)
				p.Array()
			}

			p.SetState(57)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == HoconParserSTRING || _la == HoconParserPATH_ELEMENT {
			{
				p.SetState(58)
				p.Prop()
			}

			p.SetState(63)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IPropContext is an interface to support dynamic dispatch.
type IPropContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropContext differentiates from other interfaces.
	IsPropContext()
}

type PropContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropContext() *PropContext {
	var p = new(PropContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HoconParserRULE_prop
	return p
}

func (*PropContext) IsPropContext() {}

func NewPropContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropContext {
	var p = new(PropContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HoconParserRULE_prop

	return p
}

func (s *PropContext) GetParser() antlr.Parser { return s.parser }

func (s *PropContext) ObjectData() IObjectDataContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectDataContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectDataContext)
}

func (s *PropContext) ArrayData() IArrayDataContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayDataContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayDataContext)
}

func (s *PropContext) StringData() IStringDataContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringDataContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringDataContext)
}

func (s *PropContext) ReferenceData() IReferenceDataContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceDataContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceDataContext)
}

func (s *PropContext) NumberData() INumberDataContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberDataContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberDataContext)
}

func (s *PropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.EnterProp(s)
	}
}

func (s *PropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.ExitProp(s)
	}
}

func (s *PropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HoconVisitor:
		return t.VisitProp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HoconParser) Prop() (localctx IPropContext) {
	this := p
	_ = this

	localctx = NewPropContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HoconParserRULE_prop)

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

	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)
			p.ObjectData()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)
			p.ArrayData()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(68)
			p.StringData()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(69)
			p.ReferenceData()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(70)
			p.NumberData()
		}

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
	p.RuleIndex = HoconParserRULE_obj
	return p
}

func (*ObjContext) IsObjContext() {}

func NewObjContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjContext {
	var p = new(ObjContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HoconParserRULE_obj

	return p
}

func (s *ObjContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjContext) ObjectBegin() IObjectBeginContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectBeginContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectBeginContext)
}

func (s *ObjContext) AllProp() []IPropContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPropContext); ok {
			len++
		}
	}

	tst := make([]IPropContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPropContext); ok {
			tst[i] = t.(IPropContext)
			i++
		}
	}

	return tst
}

func (s *ObjContext) Prop(i int) IPropContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropContext); ok {
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

	return t.(IPropContext)
}

func (s *ObjContext) ObjectEnd() IObjectEndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectEndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectEndContext)
}

func (s *ObjContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.EnterObj(s)
	}
}

func (s *ObjContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.ExitObj(s)
	}
}

func (s *ObjContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HoconVisitor:
		return t.VisitObj(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HoconParser) Obj() (localctx IObjContext) {
	this := p
	_ = this

	localctx = NewObjContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, HoconParserRULE_obj)
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

	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(73)
			p.ObjectBegin()
		}
		{
			p.SetState(74)
			p.Prop()
		}
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<HoconParserT__0)|(1<<HoconParserSTRING)|(1<<HoconParserPATH_ELEMENT))) != 0 {
			p.SetState(76)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == HoconParserT__0 {
				{
					p.SetState(75)
					p.Match(HoconParserT__0)
				}

			}
			{
				p.SetState(78)
				p.Prop()
			}

			p.SetState(83)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(84)
			p.ObjectEnd()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(86)
			p.ObjectBegin()
		}
		{
			p.SetState(87)
			p.ObjectEnd()
		}

	}

	return localctx
}

// IObjectBeginContext is an interface to support dynamic dispatch.
type IObjectBeginContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectBeginContext differentiates from other interfaces.
	IsObjectBeginContext()
}

type ObjectBeginContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectBeginContext() *ObjectBeginContext {
	var p = new(ObjectBeginContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HoconParserRULE_objectBegin
	return p
}

func (*ObjectBeginContext) IsObjectBeginContext() {}

func NewObjectBeginContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectBeginContext {
	var p = new(ObjectBeginContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HoconParserRULE_objectBegin

	return p
}

func (s *ObjectBeginContext) GetParser() antlr.Parser { return s.parser }
func (s *ObjectBeginContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectBeginContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectBeginContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.EnterObjectBegin(s)
	}
}

func (s *ObjectBeginContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.ExitObjectBegin(s)
	}
}

func (s *ObjectBeginContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HoconVisitor:
		return t.VisitObjectBegin(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HoconParser) ObjectBegin() (localctx IObjectBeginContext) {
	this := p
	_ = this

	localctx = NewObjectBeginContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, HoconParserRULE_objectBegin)

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
		p.SetState(91)
		p.Match(HoconParserT__1)
	}

	return localctx
}

// IObjectEndContext is an interface to support dynamic dispatch.
type IObjectEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectEndContext differentiates from other interfaces.
	IsObjectEndContext()
}

type ObjectEndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectEndContext() *ObjectEndContext {
	var p = new(ObjectEndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HoconParserRULE_objectEnd
	return p
}

func (*ObjectEndContext) IsObjectEndContext() {}

func NewObjectEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectEndContext {
	var p = new(ObjectEndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HoconParserRULE_objectEnd

	return p
}

func (s *ObjectEndContext) GetParser() antlr.Parser { return s.parser }
func (s *ObjectEndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectEndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectEndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.EnterObjectEnd(s)
	}
}

func (s *ObjectEndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.ExitObjectEnd(s)
	}
}

func (s *ObjectEndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HoconVisitor:
		return t.VisitObjectEnd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HoconParser) ObjectEnd() (localctx IObjectEndContext) {
	this := p
	_ = this

	localctx = NewObjectEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, HoconParserRULE_objectEnd)

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
		p.SetState(93)
		p.Match(HoconParserT__2)
	}

	return localctx
}

// IObjectDataContext is an interface to support dynamic dispatch.
type IObjectDataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectDataContext differentiates from other interfaces.
	IsObjectDataContext()
}

type ObjectDataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectDataContext() *ObjectDataContext {
	var p = new(ObjectDataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HoconParserRULE_objectData
	return p
}

func (*ObjectDataContext) IsObjectDataContext() {}

func NewObjectDataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectDataContext {
	var p = new(ObjectDataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HoconParserRULE_objectData

	return p
}

func (s *ObjectDataContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectDataContext) Key() IKeyContext {
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

func (s *ObjectDataContext) Obj() IObjContext {
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

func (s *ObjectDataContext) KV() antlr.TerminalNode {
	return s.GetToken(HoconParserKV, 0)
}

func (s *ObjectDataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectDataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectDataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.EnterObjectData(s)
	}
}

func (s *ObjectDataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.ExitObjectData(s)
	}
}

func (s *ObjectDataContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HoconVisitor:
		return t.VisitObjectData(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HoconParser) ObjectData() (localctx IObjectDataContext) {
	this := p
	_ = this

	localctx = NewObjectDataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, HoconParserRULE_objectData)
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
		p.SetState(95)
		p.Key()
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == HoconParserKV {
		{
			p.SetState(96)
			p.Match(HoconParserKV)
		}

	}
	{
		p.SetState(99)
		p.Obj()
	}

	return localctx
}

// IArrayDataContext is an interface to support dynamic dispatch.
type IArrayDataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayDataContext differentiates from other interfaces.
	IsArrayDataContext()
}

type ArrayDataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayDataContext() *ArrayDataContext {
	var p = new(ArrayDataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HoconParserRULE_arrayData
	return p
}

func (*ArrayDataContext) IsArrayDataContext() {}

func NewArrayDataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayDataContext {
	var p = new(ArrayDataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HoconParserRULE_arrayData

	return p
}

func (s *ArrayDataContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayDataContext) Key() IKeyContext {
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

func (s *ArrayDataContext) KV() antlr.TerminalNode {
	return s.GetToken(HoconParserKV, 0)
}

func (s *ArrayDataContext) Array() IArrayContext {
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

func (s *ArrayDataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayDataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayDataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.EnterArrayData(s)
	}
}

func (s *ArrayDataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.ExitArrayData(s)
	}
}

func (s *ArrayDataContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HoconVisitor:
		return t.VisitArrayData(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HoconParser) ArrayData() (localctx IArrayDataContext) {
	this := p
	_ = this

	localctx = NewArrayDataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, HoconParserRULE_arrayData)

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
		p.Key()
	}
	{
		p.SetState(102)
		p.Match(HoconParserKV)
	}
	{
		p.SetState(103)
		p.Array()
	}

	return localctx
}

// IStringDataContext is an interface to support dynamic dispatch.
type IStringDataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringDataContext differentiates from other interfaces.
	IsStringDataContext()
}

type StringDataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringDataContext() *StringDataContext {
	var p = new(StringDataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HoconParserRULE_stringData
	return p
}

func (*StringDataContext) IsStringDataContext() {}

func NewStringDataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringDataContext {
	var p = new(StringDataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HoconParserRULE_stringData

	return p
}

func (s *StringDataContext) GetParser() antlr.Parser { return s.parser }

func (s *StringDataContext) Key() IKeyContext {
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

func (s *StringDataContext) KV() antlr.TerminalNode {
	return s.GetToken(HoconParserKV, 0)
}

func (s *StringDataContext) StringValue() IStringValueContext {
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

func (s *StringDataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringDataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringDataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.EnterStringData(s)
	}
}

func (s *StringDataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.ExitStringData(s)
	}
}

func (s *StringDataContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HoconVisitor:
		return t.VisitStringData(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HoconParser) StringData() (localctx IStringDataContext) {
	this := p
	_ = this

	localctx = NewStringDataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, HoconParserRULE_stringData)

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
		p.Key()
	}
	{
		p.SetState(106)
		p.Match(HoconParserKV)
	}
	{
		p.SetState(107)
		p.StringValue()
	}

	return localctx
}

// IReferenceDataContext is an interface to support dynamic dispatch.
type IReferenceDataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReferenceDataContext differentiates from other interfaces.
	IsReferenceDataContext()
}

type ReferenceDataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferenceDataContext() *ReferenceDataContext {
	var p = new(ReferenceDataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HoconParserRULE_referenceData
	return p
}

func (*ReferenceDataContext) IsReferenceDataContext() {}

func NewReferenceDataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceDataContext {
	var p = new(ReferenceDataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HoconParserRULE_referenceData

	return p
}

func (s *ReferenceDataContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferenceDataContext) Key() IKeyContext {
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

func (s *ReferenceDataContext) KV() antlr.TerminalNode {
	return s.GetToken(HoconParserKV, 0)
}

func (s *ReferenceDataContext) REFERENCE() antlr.TerminalNode {
	return s.GetToken(HoconParserREFERENCE, 0)
}

func (s *ReferenceDataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceDataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferenceDataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.EnterReferenceData(s)
	}
}

func (s *ReferenceDataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.ExitReferenceData(s)
	}
}

func (s *ReferenceDataContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HoconVisitor:
		return t.VisitReferenceData(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HoconParser) ReferenceData() (localctx IReferenceDataContext) {
	this := p
	_ = this

	localctx = NewReferenceDataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, HoconParserRULE_referenceData)

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
		p.Key()
	}
	{
		p.SetState(110)
		p.Match(HoconParserKV)
	}
	{
		p.SetState(111)
		p.Match(HoconParserREFERENCE)
	}

	return localctx
}

// INumberDataContext is an interface to support dynamic dispatch.
type INumberDataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberDataContext differentiates from other interfaces.
	IsNumberDataContext()
}

type NumberDataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberDataContext() *NumberDataContext {
	var p = new(NumberDataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HoconParserRULE_numberData
	return p
}

func (*NumberDataContext) IsNumberDataContext() {}

func NewNumberDataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberDataContext {
	var p = new(NumberDataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HoconParserRULE_numberData

	return p
}

func (s *NumberDataContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberDataContext) Key() IKeyContext {
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

func (s *NumberDataContext) KV() antlr.TerminalNode {
	return s.GetToken(HoconParserKV, 0)
}

func (s *NumberDataContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(HoconParserNUMBER, 0)
}

func (s *NumberDataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberDataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberDataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.EnterNumberData(s)
	}
}

func (s *NumberDataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.ExitNumberData(s)
	}
}

func (s *NumberDataContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HoconVisitor:
		return t.VisitNumberData(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HoconParser) NumberData() (localctx INumberDataContext) {
	this := p
	_ = this

	localctx = NewNumberDataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, HoconParserRULE_numberData)

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
		p.SetState(113)
		p.Key()
	}
	{
		p.SetState(114)
		p.Match(HoconParserKV)
	}
	{
		p.SetState(115)
		p.Match(HoconParserNUMBER)
	}

	return localctx
}

// IKeyContext is an interface to support dynamic dispatch.
type IKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	p.RuleIndex = HoconParserRULE_key
	return p
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HoconParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *KeyContext) STRING() antlr.TerminalNode {
	return s.GetToken(HoconParserSTRING, 0)
}

func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.EnterKey(s)
	}
}

func (s *KeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.ExitKey(s)
	}
}

func (s *KeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HoconVisitor:
		return t.VisitKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HoconParser) Key() (localctx IKeyContext) {
	this := p
	_ = this

	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, HoconParserRULE_key)

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

	p.SetState(119)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case HoconParserPATH_ELEMENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(117)
			p.Path()
		}

	case HoconParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
			p.Match(HoconParserSTRING)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPathContext is an interface to support dynamic dispatch.
type IPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPathContext differentiates from other interfaces.
	IsPathContext()
}

type PathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathContext() *PathContext {
	var p = new(PathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HoconParserRULE_path
	return p
}

func (*PathContext) IsPathContext() {}

func NewPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathContext {
	var p = new(PathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HoconParserRULE_path

	return p
}

func (s *PathContext) GetParser() antlr.Parser { return s.parser }

func (s *PathContext) AllPATH_ELEMENT() []antlr.TerminalNode {
	return s.GetTokens(HoconParserPATH_ELEMENT)
}

func (s *PathContext) PATH_ELEMENT(i int) antlr.TerminalNode {
	return s.GetToken(HoconParserPATH_ELEMENT, i)
}

func (s *PathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.EnterPath(s)
	}
}

func (s *PathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.ExitPath(s)
	}
}

func (s *PathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HoconVisitor:
		return t.VisitPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HoconParser) Path() (localctx IPathContext) {
	this := p
	_ = this

	localctx = NewPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, HoconParserRULE_path)
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
		p.SetState(121)
		p.Match(HoconParserPATH_ELEMENT)
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == HoconParserT__3 {
		{
			p.SetState(122)
			p.Match(HoconParserT__3)
		}
		{
			p.SetState(123)
			p.Match(HoconParserPATH_ELEMENT)
		}

		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IArrayBeginContext is an interface to support dynamic dispatch.
type IArrayBeginContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayBeginContext differentiates from other interfaces.
	IsArrayBeginContext()
}

type ArrayBeginContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayBeginContext() *ArrayBeginContext {
	var p = new(ArrayBeginContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HoconParserRULE_arrayBegin
	return p
}

func (*ArrayBeginContext) IsArrayBeginContext() {}

func NewArrayBeginContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayBeginContext {
	var p = new(ArrayBeginContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HoconParserRULE_arrayBegin

	return p
}

func (s *ArrayBeginContext) GetParser() antlr.Parser { return s.parser }
func (s *ArrayBeginContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayBeginContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayBeginContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.EnterArrayBegin(s)
	}
}

func (s *ArrayBeginContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.ExitArrayBegin(s)
	}
}

func (s *ArrayBeginContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HoconVisitor:
		return t.VisitArrayBegin(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HoconParser) ArrayBegin() (localctx IArrayBeginContext) {
	this := p
	_ = this

	localctx = NewArrayBeginContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, HoconParserRULE_arrayBegin)

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
		p.SetState(129)
		p.Match(HoconParserT__4)
	}

	return localctx
}

// IArrayEndContext is an interface to support dynamic dispatch.
type IArrayEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayEndContext differentiates from other interfaces.
	IsArrayEndContext()
}

type ArrayEndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayEndContext() *ArrayEndContext {
	var p = new(ArrayEndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HoconParserRULE_arrayEnd
	return p
}

func (*ArrayEndContext) IsArrayEndContext() {}

func NewArrayEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayEndContext {
	var p = new(ArrayEndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HoconParserRULE_arrayEnd

	return p
}

func (s *ArrayEndContext) GetParser() antlr.Parser { return s.parser }
func (s *ArrayEndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayEndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayEndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.EnterArrayEnd(s)
	}
}

func (s *ArrayEndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.ExitArrayEnd(s)
	}
}

func (s *ArrayEndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HoconVisitor:
		return t.VisitArrayEnd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HoconParser) ArrayEnd() (localctx IArrayEndContext) {
	this := p
	_ = this

	localctx = NewArrayEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, HoconParserRULE_arrayEnd)

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
		p.SetState(131)
		p.Match(HoconParserT__5)
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
	p.RuleIndex = HoconParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HoconParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) ArrayBegin() IArrayBeginContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayBeginContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayBeginContext)
}

func (s *ArrayContext) AllArrayValue() []IArrayValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArrayValueContext); ok {
			len++
		}
	}

	tst := make([]IArrayValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArrayValueContext); ok {
			tst[i] = t.(IArrayValueContext)
			i++
		}
	}

	return tst
}

func (s *ArrayContext) ArrayValue(i int) IArrayValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayValueContext); ok {
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

	return t.(IArrayValueContext)
}

func (s *ArrayContext) ArrayEnd() IArrayEndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayEndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayEndContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.ExitArray(s)
	}
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HoconVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HoconParser) Array() (localctx IArrayContext) {
	this := p
	_ = this

	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, HoconParserRULE_array)
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

	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(133)
			p.ArrayBegin()
		}
		{
			p.SetState(134)
			p.ArrayValue()
		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<HoconParserT__0)|(1<<HoconParserT__1)|(1<<HoconParserT__4)|(1<<HoconParserT__6)|(1<<HoconParserNUMBER)|(1<<HoconParserSTRING)|(1<<HoconParserPATH_ELEMENT)|(1<<HoconParserREFERENCE))) != 0 {
			p.SetState(136)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == HoconParserT__0 {
				{
					p.SetState(135)
					p.Match(HoconParserT__0)
				}

			}
			{
				p.SetState(138)
				p.ArrayValue()
			}

			p.SetState(143)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(144)
			p.ArrayEnd()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(146)
			p.ArrayBegin()
		}
		{
			p.SetState(147)
			p.ArrayEnd()
		}

	}

	return localctx
}

// IArrayValueContext is an interface to support dynamic dispatch.
type IArrayValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayValueContext differentiates from other interfaces.
	IsArrayValueContext()
}

type ArrayValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayValueContext() *ArrayValueContext {
	var p = new(ArrayValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HoconParserRULE_arrayValue
	return p
}

func (*ArrayValueContext) IsArrayValueContext() {}

func NewArrayValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayValueContext {
	var p = new(ArrayValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HoconParserRULE_arrayValue

	return p
}

func (s *ArrayValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayValueContext) ArrayString() IArrayStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayStringContext)
}

func (s *ArrayValueContext) ArrayReference() IArrayReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayReferenceContext)
}

func (s *ArrayValueContext) ArrayNumber() IArrayNumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayNumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayNumberContext)
}

func (s *ArrayValueContext) ArrayObj() IArrayObjContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayObjContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayObjContext)
}

func (s *ArrayValueContext) ArrayArray() IArrayArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayArrayContext)
}

func (s *ArrayValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.EnterArrayValue(s)
	}
}

func (s *ArrayValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.ExitArrayValue(s)
	}
}

func (s *ArrayValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HoconVisitor:
		return t.VisitArrayValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HoconParser) ArrayValue() (localctx IArrayValueContext) {
	this := p
	_ = this

	localctx = NewArrayValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, HoconParserRULE_arrayValue)

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

	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(151)
			p.ArrayString()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(152)
			p.ArrayReference()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(153)
			p.ArrayNumber()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(154)
			p.ArrayObj()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(155)
			p.ArrayArray()
		}

	}

	return localctx
}

// IArrayStringContext is an interface to support dynamic dispatch.
type IArrayStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayStringContext differentiates from other interfaces.
	IsArrayStringContext()
}

type ArrayStringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayStringContext() *ArrayStringContext {
	var p = new(ArrayStringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HoconParserRULE_arrayString
	return p
}

func (*ArrayStringContext) IsArrayStringContext() {}

func NewArrayStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayStringContext {
	var p = new(ArrayStringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HoconParserRULE_arrayString

	return p
}

func (s *ArrayStringContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayStringContext) StringValue() IStringValueContext {
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

func (s *ArrayStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.EnterArrayString(s)
	}
}

func (s *ArrayStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.ExitArrayString(s)
	}
}

func (s *ArrayStringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HoconVisitor:
		return t.VisitArrayString(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HoconParser) ArrayString() (localctx IArrayStringContext) {
	this := p
	_ = this

	localctx = NewArrayStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, HoconParserRULE_arrayString)

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
		p.SetState(158)
		p.StringValue()
	}

	return localctx
}

// IArrayReferenceContext is an interface to support dynamic dispatch.
type IArrayReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayReferenceContext differentiates from other interfaces.
	IsArrayReferenceContext()
}

type ArrayReferenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayReferenceContext() *ArrayReferenceContext {
	var p = new(ArrayReferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HoconParserRULE_arrayReference
	return p
}

func (*ArrayReferenceContext) IsArrayReferenceContext() {}

func NewArrayReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayReferenceContext {
	var p = new(ArrayReferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HoconParserRULE_arrayReference

	return p
}

func (s *ArrayReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayReferenceContext) REFERENCE() antlr.TerminalNode {
	return s.GetToken(HoconParserREFERENCE, 0)
}

func (s *ArrayReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.EnterArrayReference(s)
	}
}

func (s *ArrayReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.ExitArrayReference(s)
	}
}

func (s *ArrayReferenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HoconVisitor:
		return t.VisitArrayReference(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HoconParser) ArrayReference() (localctx IArrayReferenceContext) {
	this := p
	_ = this

	localctx = NewArrayReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, HoconParserRULE_arrayReference)

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
		p.SetState(160)
		p.Match(HoconParserREFERENCE)
	}

	return localctx
}

// IArrayNumberContext is an interface to support dynamic dispatch.
type IArrayNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayNumberContext differentiates from other interfaces.
	IsArrayNumberContext()
}

type ArrayNumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayNumberContext() *ArrayNumberContext {
	var p = new(ArrayNumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HoconParserRULE_arrayNumber
	return p
}

func (*ArrayNumberContext) IsArrayNumberContext() {}

func NewArrayNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayNumberContext {
	var p = new(ArrayNumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HoconParserRULE_arrayNumber

	return p
}

func (s *ArrayNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayNumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(HoconParserNUMBER, 0)
}

func (s *ArrayNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.EnterArrayNumber(s)
	}
}

func (s *ArrayNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.ExitArrayNumber(s)
	}
}

func (s *ArrayNumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HoconVisitor:
		return t.VisitArrayNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HoconParser) ArrayNumber() (localctx IArrayNumberContext) {
	this := p
	_ = this

	localctx = NewArrayNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, HoconParserRULE_arrayNumber)

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
		p.SetState(162)
		p.Match(HoconParserNUMBER)
	}

	return localctx
}

// IArrayObjContext is an interface to support dynamic dispatch.
type IArrayObjContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayObjContext differentiates from other interfaces.
	IsArrayObjContext()
}

type ArrayObjContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayObjContext() *ArrayObjContext {
	var p = new(ArrayObjContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HoconParserRULE_arrayObj
	return p
}

func (*ArrayObjContext) IsArrayObjContext() {}

func NewArrayObjContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayObjContext {
	var p = new(ArrayObjContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HoconParserRULE_arrayObj

	return p
}

func (s *ArrayObjContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayObjContext) Obj() IObjContext {
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

func (s *ArrayObjContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayObjContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayObjContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.EnterArrayObj(s)
	}
}

func (s *ArrayObjContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.ExitArrayObj(s)
	}
}

func (s *ArrayObjContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HoconVisitor:
		return t.VisitArrayObj(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HoconParser) ArrayObj() (localctx IArrayObjContext) {
	this := p
	_ = this

	localctx = NewArrayObjContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, HoconParserRULE_arrayObj)

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
		p.SetState(164)
		p.Obj()
	}

	return localctx
}

// IArrayArrayContext is an interface to support dynamic dispatch.
type IArrayArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayArrayContext differentiates from other interfaces.
	IsArrayArrayContext()
}

type ArrayArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayArrayContext() *ArrayArrayContext {
	var p = new(ArrayArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HoconParserRULE_arrayArray
	return p
}

func (*ArrayArrayContext) IsArrayArrayContext() {}

func NewArrayArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayArrayContext {
	var p = new(ArrayArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HoconParserRULE_arrayArray

	return p
}

func (s *ArrayArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayArrayContext) Array() IArrayContext {
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

func (s *ArrayArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.EnterArrayArray(s)
	}
}

func (s *ArrayArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.ExitArrayArray(s)
	}
}

func (s *ArrayArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HoconVisitor:
		return t.VisitArrayArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HoconParser) ArrayArray() (localctx IArrayArrayContext) {
	this := p
	_ = this

	localctx = NewArrayArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, HoconParserRULE_arrayArray)

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
		p.SetState(166)
		p.Array()
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
	p.RuleIndex = HoconParserRULE_stringValue
	return p
}

func (*StringValueContext) IsStringValueContext() {}

func NewStringValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringValueContext {
	var p = new(StringValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HoconParserRULE_stringValue

	return p
}

func (s *StringValueContext) GetParser() antlr.Parser { return s.parser }

func (s *StringValueContext) CopyFrom(ctx *StringValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type V_rawstringContext struct {
	*StringValueContext
}

func NewV_rawstringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *V_rawstringContext {
	var p = new(V_rawstringContext)

	p.StringValueContext = NewEmptyStringValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StringValueContext))

	return p
}

func (s *V_rawstringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *V_rawstringContext) Rawstring() IRawstringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRawstringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRawstringContext)
}

func (s *V_rawstringContext) AllStringValue() []IStringValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStringValueContext); ok {
			len++
		}
	}

	tst := make([]IStringValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStringValueContext); ok {
			tst[i] = t.(IStringValueContext)
			i++
		}
	}

	return tst
}

func (s *V_rawstringContext) StringValue(i int) IStringValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringValueContext); ok {
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

	return t.(IStringValueContext)
}

func (s *V_rawstringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.EnterV_rawstring(s)
	}
}

func (s *V_rawstringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.ExitV_rawstring(s)
	}
}

func (s *V_rawstringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HoconVisitor:
		return t.VisitV_rawstring(s)

	default:
		return t.VisitChildren(s)
	}
}

type V_referenceContext struct {
	*StringValueContext
}

func NewV_referenceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *V_referenceContext {
	var p = new(V_referenceContext)

	p.StringValueContext = NewEmptyStringValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StringValueContext))

	return p
}

func (s *V_referenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *V_referenceContext) REFERENCE() antlr.TerminalNode {
	return s.GetToken(HoconParserREFERENCE, 0)
}

func (s *V_referenceContext) AllStringValue() []IStringValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStringValueContext); ok {
			len++
		}
	}

	tst := make([]IStringValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStringValueContext); ok {
			tst[i] = t.(IStringValueContext)
			i++
		}
	}

	return tst
}

func (s *V_referenceContext) StringValue(i int) IStringValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringValueContext); ok {
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

	return t.(IStringValueContext)
}

func (s *V_referenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.EnterV_reference(s)
	}
}

func (s *V_referenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.ExitV_reference(s)
	}
}

func (s *V_referenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HoconVisitor:
		return t.VisitV_reference(s)

	default:
		return t.VisitChildren(s)
	}
}

type V_stringContext struct {
	*StringValueContext
}

func NewV_stringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *V_stringContext {
	var p = new(V_stringContext)

	p.StringValueContext = NewEmptyStringValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StringValueContext))

	return p
}

func (s *V_stringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *V_stringContext) STRING() antlr.TerminalNode {
	return s.GetToken(HoconParserSTRING, 0)
}

func (s *V_stringContext) AllStringValue() []IStringValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStringValueContext); ok {
			len++
		}
	}

	tst := make([]IStringValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStringValueContext); ok {
			tst[i] = t.(IStringValueContext)
			i++
		}
	}

	return tst
}

func (s *V_stringContext) StringValue(i int) IStringValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringValueContext); ok {
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

	return t.(IStringValueContext)
}

func (s *V_stringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.EnterV_string(s)
	}
}

func (s *V_stringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.ExitV_string(s)
	}
}

func (s *V_stringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HoconVisitor:
		return t.VisitV_string(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HoconParser) StringValue() (localctx IStringValueContext) {
	this := p
	_ = this

	localctx = NewStringValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, HoconParserRULE_stringValue)

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

	p.SetState(189)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case HoconParserSTRING:
		localctx = NewV_stringContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(168)
			p.Match(HoconParserSTRING)
		}
		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(169)
					p.StringValue()
				}

			}
			p.SetState(174)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
		}

	case HoconParserT__6, HoconParserPATH_ELEMENT:
		localctx = NewV_rawstringContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(175)
			p.Rawstring()
		}
		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(176)
					p.StringValue()
				}

			}
			p.SetState(181)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
		}

	case HoconParserREFERENCE:
		localctx = NewV_referenceContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(182)
			p.Match(HoconParserREFERENCE)
		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(183)
					p.StringValue()
				}

			}
			p.SetState(188)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRawstringContext is an interface to support dynamic dispatch.
type IRawstringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRawstringContext differentiates from other interfaces.
	IsRawstringContext()
}

type RawstringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRawstringContext() *RawstringContext {
	var p = new(RawstringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HoconParserRULE_rawstring
	return p
}

func (*RawstringContext) IsRawstringContext() {}

func NewRawstringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RawstringContext {
	var p = new(RawstringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HoconParserRULE_rawstring

	return p
}

func (s *RawstringContext) GetParser() antlr.Parser { return s.parser }

func (s *RawstringContext) AllPATH_ELEMENT() []antlr.TerminalNode {
	return s.GetTokens(HoconParserPATH_ELEMENT)
}

func (s *RawstringContext) PATH_ELEMENT(i int) antlr.TerminalNode {
	return s.GetToken(HoconParserPATH_ELEMENT, i)
}

func (s *RawstringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RawstringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RawstringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.EnterRawstring(s)
	}
}

func (s *RawstringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoconListener); ok {
		listenerT.ExitRawstring(s)
	}
}

func (s *RawstringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case HoconVisitor:
		return t.VisitRawstring(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *HoconParser) Rawstring() (localctx IRawstringContext) {
	this := p
	_ = this

	localctx = NewRawstringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, HoconParserRULE_rawstring)
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
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(191)
				_la = p.GetTokenStream().LA(1)

				if !(_la == HoconParserT__6 || _la == HoconParserPATH_ELEMENT) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}

	return localctx
}
