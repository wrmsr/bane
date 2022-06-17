// Code generated from Json.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type JsonLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var jsonlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jsonlexerLexerInit() {
	staticData := &jsonlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'{'", "','", "'}'", "':'", "'['", "']'", "'true'", "'false'", "'null'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "STRING", "NUMBER", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"STRING", "ESC", "UNICODE", "HEX", "SAFECODEPOINT", "NUMBER", "INT",
		"EXP", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 12, 128, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 9, 5, 9, 69, 8, 9, 10, 9, 12, 9, 72, 9, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 3, 10, 79, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 14, 3, 14, 92, 8, 14, 1, 14, 1, 14, 1, 14, 4, 14,
		97, 8, 14, 11, 14, 12, 14, 98, 3, 14, 101, 8, 14, 1, 14, 3, 14, 104, 8,
		14, 1, 15, 1, 15, 1, 15, 5, 15, 109, 8, 15, 10, 15, 12, 15, 112, 9, 15,
		3, 15, 114, 8, 15, 1, 16, 1, 16, 3, 16, 118, 8, 16, 1, 16, 1, 16, 1, 17,
		4, 17, 123, 8, 17, 11, 17, 12, 17, 124, 1, 17, 1, 17, 0, 0, 18, 1, 1, 3,
		2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 0, 23, 0,
		25, 0, 27, 0, 29, 11, 31, 0, 33, 0, 35, 12, 1, 0, 8, 8, 0, 34, 34, 47,
		47, 92, 92, 98, 98, 102, 102, 110, 110, 114, 114, 116, 116, 3, 0, 48, 57,
		65, 70, 97, 102, 3, 0, 0, 31, 34, 34, 92, 92, 1, 0, 48, 57, 1, 0, 49, 57,
		2, 0, 69, 69, 101, 101, 2, 0, 43, 43, 45, 45, 3, 0, 9, 10, 13, 13, 32,
		32, 132, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0,
		35, 1, 0, 0, 0, 1, 37, 1, 0, 0, 0, 3, 39, 1, 0, 0, 0, 5, 41, 1, 0, 0, 0,
		7, 43, 1, 0, 0, 0, 9, 45, 1, 0, 0, 0, 11, 47, 1, 0, 0, 0, 13, 49, 1, 0,
		0, 0, 15, 54, 1, 0, 0, 0, 17, 60, 1, 0, 0, 0, 19, 65, 1, 0, 0, 0, 21, 75,
		1, 0, 0, 0, 23, 80, 1, 0, 0, 0, 25, 86, 1, 0, 0, 0, 27, 88, 1, 0, 0, 0,
		29, 91, 1, 0, 0, 0, 31, 113, 1, 0, 0, 0, 33, 115, 1, 0, 0, 0, 35, 122,
		1, 0, 0, 0, 37, 38, 5, 123, 0, 0, 38, 2, 1, 0, 0, 0, 39, 40, 5, 44, 0,
		0, 40, 4, 1, 0, 0, 0, 41, 42, 5, 125, 0, 0, 42, 6, 1, 0, 0, 0, 43, 44,
		5, 58, 0, 0, 44, 8, 1, 0, 0, 0, 45, 46, 5, 91, 0, 0, 46, 10, 1, 0, 0, 0,
		47, 48, 5, 93, 0, 0, 48, 12, 1, 0, 0, 0, 49, 50, 5, 116, 0, 0, 50, 51,
		5, 114, 0, 0, 51, 52, 5, 117, 0, 0, 52, 53, 5, 101, 0, 0, 53, 14, 1, 0,
		0, 0, 54, 55, 5, 102, 0, 0, 55, 56, 5, 97, 0, 0, 56, 57, 5, 108, 0, 0,
		57, 58, 5, 115, 0, 0, 58, 59, 5, 101, 0, 0, 59, 16, 1, 0, 0, 0, 60, 61,
		5, 110, 0, 0, 61, 62, 5, 117, 0, 0, 62, 63, 5, 108, 0, 0, 63, 64, 5, 108,
		0, 0, 64, 18, 1, 0, 0, 0, 65, 70, 5, 34, 0, 0, 66, 69, 3, 21, 10, 0, 67,
		69, 3, 27, 13, 0, 68, 66, 1, 0, 0, 0, 68, 67, 1, 0, 0, 0, 69, 72, 1, 0,
		0, 0, 70, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 73, 1, 0, 0, 0, 72, 70,
		1, 0, 0, 0, 73, 74, 5, 34, 0, 0, 74, 20, 1, 0, 0, 0, 75, 78, 5, 92, 0,
		0, 76, 79, 7, 0, 0, 0, 77, 79, 3, 23, 11, 0, 78, 76, 1, 0, 0, 0, 78, 77,
		1, 0, 0, 0, 79, 22, 1, 0, 0, 0, 80, 81, 5, 117, 0, 0, 81, 82, 3, 25, 12,
		0, 82, 83, 3, 25, 12, 0, 83, 84, 3, 25, 12, 0, 84, 85, 3, 25, 12, 0, 85,
		24, 1, 0, 0, 0, 86, 87, 7, 1, 0, 0, 87, 26, 1, 0, 0, 0, 88, 89, 8, 2, 0,
		0, 89, 28, 1, 0, 0, 0, 90, 92, 5, 45, 0, 0, 91, 90, 1, 0, 0, 0, 91, 92,
		1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 100, 3, 31, 15, 0, 94, 96, 5, 46, 0,
		0, 95, 97, 7, 3, 0, 0, 96, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 96,
		1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 101, 1, 0, 0, 0, 100, 94, 1, 0, 0,
		0, 100, 101, 1, 0, 0, 0, 101, 103, 1, 0, 0, 0, 102, 104, 3, 33, 16, 0,
		103, 102, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 30, 1, 0, 0, 0, 105, 114,
		5, 48, 0, 0, 106, 110, 7, 4, 0, 0, 107, 109, 7, 3, 0, 0, 108, 107, 1, 0,
		0, 0, 109, 112, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0,
		111, 114, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 113, 105, 1, 0, 0, 0, 113,
		106, 1, 0, 0, 0, 114, 32, 1, 0, 0, 0, 115, 117, 7, 5, 0, 0, 116, 118, 7,
		6, 0, 0, 117, 116, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 119, 1, 0, 0,
		0, 119, 120, 3, 31, 15, 0, 120, 34, 1, 0, 0, 0, 121, 123, 7, 7, 0, 0, 122,
		121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 125,
		1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 127, 6, 17, 0, 0, 127, 36, 1, 0,
		0, 0, 12, 0, 68, 70, 78, 91, 98, 100, 103, 110, 113, 117, 124, 1, 6, 0,
		0,
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

// JsonLexerInit initializes any static state used to implement JsonLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewJsonLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func JsonLexerInit() {
	staticData := &jsonlexerLexerStaticData
	staticData.once.Do(jsonlexerLexerInit)
}

// NewJsonLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewJsonLexer(input antlr.CharStream) *JsonLexer {
	JsonLexerInit()
	l := new(JsonLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &jsonlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Json.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// JsonLexer tokens.
const (
	JsonLexerT__0   = 1
	JsonLexerT__1   = 2
	JsonLexerT__2   = 3
	JsonLexerT__3   = 4
	JsonLexerT__4   = 5
	JsonLexerT__5   = 6
	JsonLexerT__6   = 7
	JsonLexerT__7   = 8
	JsonLexerT__8   = 9
	JsonLexerSTRING = 10
	JsonLexerNUMBER = 11
	JsonLexerWS     = 12
)
