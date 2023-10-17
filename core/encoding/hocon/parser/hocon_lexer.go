// Code generated from Hocon.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	antlr "github.com/wrmsr/bane/core/antlr/runtime"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type HoconLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var hoconlexerLexerStaticData struct {
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

func hoconlexerLexerInit() {
	staticData := &hoconlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "','", "'{'", "'}'", "'.'", "'['", "']'", "'-'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "COMMENT", "NUMBER", "STRING", "PATH_ELEMENT",
		"REFERENCE", "KV", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "ESC", "UNICODE",
		"ALPHANUM", "HEX", "INT", "EXP", "COMMENT", "NUMBER", "STRING", "PATH_ELEMENT",
		"REFERENCE", "KV", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 14, 173, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 7, 3, 7, 59, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 9, 3, 9, 68, 8, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 5, 11, 75, 8,
		11, 10, 11, 12, 11, 78, 9, 11, 3, 11, 80, 8, 11, 1, 12, 1, 12, 3, 12, 84,
		8, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 3, 13, 91, 8, 13, 1, 13, 5, 13,
		94, 8, 13, 10, 13, 12, 13, 97, 9, 13, 1, 13, 1, 13, 1, 14, 3, 14, 102,
		8, 14, 1, 14, 1, 14, 1, 14, 4, 14, 107, 8, 14, 11, 14, 12, 14, 108, 1,
		14, 3, 14, 112, 8, 14, 1, 14, 3, 14, 115, 8, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 3, 14, 121, 8, 14, 1, 14, 3, 14, 124, 8, 14, 1, 15, 1, 15, 1, 15, 5,
		15, 129, 8, 15, 10, 15, 12, 15, 132, 9, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		5, 15, 138, 8, 15, 10, 15, 12, 15, 141, 9, 15, 1, 15, 3, 15, 144, 8, 15,
		1, 16, 1, 16, 4, 16, 148, 8, 16, 11, 16, 12, 16, 149, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 5, 17, 158, 8, 17, 10, 17, 12, 17, 161, 9, 17,
		1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 4, 19, 168, 8, 19, 11, 19, 12, 19, 169,
		1, 19, 1, 19, 0, 0, 20, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		0, 17, 0, 19, 0, 21, 0, 23, 0, 25, 0, 27, 8, 29, 9, 31, 10, 33, 11, 35,
		12, 37, 13, 39, 14, 1, 0, 13, 8, 0, 34, 34, 47, 47, 92, 92, 98, 98, 102,
		102, 110, 110, 114, 114, 116, 116, 3, 0, 48, 57, 65, 90, 97, 122, 3, 0,
		48, 57, 65, 70, 97, 102, 1, 0, 49, 57, 1, 0, 48, 57, 2, 0, 69, 69, 101,
		101, 2, 0, 43, 43, 45, 45, 2, 0, 10, 10, 13, 13, 2, 0, 34, 34, 92, 92,
		2, 0, 39, 39, 92, 92, 2, 0, 45, 45, 95, 95, 2, 0, 58, 58, 61, 61, 3, 0,
		9, 10, 13, 13, 32, 32, 188, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1,
		0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0,
		33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0,
		1, 41, 1, 0, 0, 0, 3, 43, 1, 0, 0, 0, 5, 45, 1, 0, 0, 0, 7, 47, 1, 0, 0,
		0, 9, 49, 1, 0, 0, 0, 11, 51, 1, 0, 0, 0, 13, 53, 1, 0, 0, 0, 15, 55, 1,
		0, 0, 0, 17, 60, 1, 0, 0, 0, 19, 67, 1, 0, 0, 0, 21, 69, 1, 0, 0, 0, 23,
		79, 1, 0, 0, 0, 25, 81, 1, 0, 0, 0, 27, 90, 1, 0, 0, 0, 29, 123, 1, 0,
		0, 0, 31, 143, 1, 0, 0, 0, 33, 147, 1, 0, 0, 0, 35, 151, 1, 0, 0, 0, 37,
		164, 1, 0, 0, 0, 39, 167, 1, 0, 0, 0, 41, 42, 5, 44, 0, 0, 42, 2, 1, 0,
		0, 0, 43, 44, 5, 123, 0, 0, 44, 4, 1, 0, 0, 0, 45, 46, 5, 125, 0, 0, 46,
		6, 1, 0, 0, 0, 47, 48, 5, 46, 0, 0, 48, 8, 1, 0, 0, 0, 49, 50, 5, 91, 0,
		0, 50, 10, 1, 0, 0, 0, 51, 52, 5, 93, 0, 0, 52, 12, 1, 0, 0, 0, 53, 54,
		5, 45, 0, 0, 54, 14, 1, 0, 0, 0, 55, 58, 5, 92, 0, 0, 56, 59, 7, 0, 0,
		0, 57, 59, 3, 17, 8, 0, 58, 56, 1, 0, 0, 0, 58, 57, 1, 0, 0, 0, 59, 16,
		1, 0, 0, 0, 60, 61, 5, 117, 0, 0, 61, 62, 3, 21, 10, 0, 62, 63, 3, 21,
		10, 0, 63, 64, 3, 21, 10, 0, 64, 65, 3, 21, 10, 0, 65, 18, 1, 0, 0, 0,
		66, 68, 7, 1, 0, 0, 67, 66, 1, 0, 0, 0, 68, 20, 1, 0, 0, 0, 69, 70, 7,
		2, 0, 0, 70, 22, 1, 0, 0, 0, 71, 80, 5, 48, 0, 0, 72, 76, 7, 3, 0, 0, 73,
		75, 7, 4, 0, 0, 74, 73, 1, 0, 0, 0, 75, 78, 1, 0, 0, 0, 76, 74, 1, 0, 0,
		0, 76, 77, 1, 0, 0, 0, 77, 80, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 79, 71,
		1, 0, 0, 0, 79, 72, 1, 0, 0, 0, 80, 24, 1, 0, 0, 0, 81, 83, 7, 5, 0, 0,
		82, 84, 7, 6, 0, 0, 83, 82, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 85, 1,
		0, 0, 0, 85, 86, 3, 23, 11, 0, 86, 26, 1, 0, 0, 0, 87, 91, 5, 35, 0, 0,
		88, 89, 5, 47, 0, 0, 89, 91, 5, 47, 0, 0, 90, 87, 1, 0, 0, 0, 90, 88, 1,
		0, 0, 0, 91, 95, 1, 0, 0, 0, 92, 94, 8, 7, 0, 0, 93, 92, 1, 0, 0, 0, 94,
		97, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 98, 1, 0, 0,
		0, 97, 95, 1, 0, 0, 0, 98, 99, 6, 13, 0, 0, 99, 28, 1, 0, 0, 0, 100, 102,
		5, 45, 0, 0, 101, 100, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 103, 1, 0,
		0, 0, 103, 104, 3, 23, 11, 0, 104, 106, 5, 46, 0, 0, 105, 107, 7, 4, 0,
		0, 106, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 108,
		109, 1, 0, 0, 0, 109, 111, 1, 0, 0, 0, 110, 112, 3, 25, 12, 0, 111, 110,
		1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 124, 1, 0, 0, 0, 113, 115, 5, 45,
		0, 0, 114, 113, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0,
		116, 117, 3, 23, 11, 0, 117, 118, 3, 25, 12, 0, 118, 124, 1, 0, 0, 0, 119,
		121, 5, 45, 0, 0, 120, 119, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 122,
		1, 0, 0, 0, 122, 124, 3, 23, 11, 0, 123, 101, 1, 0, 0, 0, 123, 114, 1,
		0, 0, 0, 123, 120, 1, 0, 0, 0, 124, 30, 1, 0, 0, 0, 125, 130, 5, 34, 0,
		0, 126, 129, 3, 15, 7, 0, 127, 129, 8, 8, 0, 0, 128, 126, 1, 0, 0, 0, 128,
		127, 1, 0, 0, 0, 129, 132, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 130, 131,
		1, 0, 0, 0, 131, 133, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 133, 144, 5, 34,
		0, 0, 134, 139, 5, 39, 0, 0, 135, 138, 3, 15, 7, 0, 136, 138, 8, 9, 0,
		0, 137, 135, 1, 0, 0, 0, 137, 136, 1, 0, 0, 0, 138, 141, 1, 0, 0, 0, 139,
		137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 142, 1, 0, 0, 0, 141, 139,
		1, 0, 0, 0, 142, 144, 5, 39, 0, 0, 143, 125, 1, 0, 0, 0, 143, 134, 1, 0,
		0, 0, 144, 32, 1, 0, 0, 0, 145, 148, 3, 19, 9, 0, 146, 148, 7, 10, 0, 0,
		147, 145, 1, 0, 0, 0, 147, 146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149,
		147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 34, 1, 0, 0, 0, 151, 152, 5,
		36, 0, 0, 152, 153, 5, 123, 0, 0, 153, 154, 1, 0, 0, 0, 154, 159, 3, 33,
		16, 0, 155, 156, 5, 46, 0, 0, 156, 158, 3, 33, 16, 0, 157, 155, 1, 0, 0,
		0, 158, 161, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160,
		162, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 162, 163, 5, 125, 0, 0, 163, 36,
		1, 0, 0, 0, 164, 165, 7, 11, 0, 0, 165, 38, 1, 0, 0, 0, 166, 168, 7, 12,
		0, 0, 167, 166, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0,
		169, 170, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 172, 6, 19, 0, 0, 172,
		40, 1, 0, 0, 0, 23, 0, 58, 67, 76, 79, 83, 90, 95, 101, 108, 111, 114,
		120, 123, 128, 130, 137, 139, 143, 147, 149, 159, 169, 1, 6, 0, 0,
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

// HoconLexerInit initializes any static state used to implement HoconLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewHoconLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func HoconLexerInit() {
	staticData := &hoconlexerLexerStaticData
	staticData.once.Do(hoconlexerLexerInit)
}

// NewHoconLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewHoconLexer(input antlr.CharStream) *HoconLexer {
	HoconLexerInit()
	l := new(HoconLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &hoconlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Hocon.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// HoconLexer tokens.
const (
	HoconLexerT__0         = 1
	HoconLexerT__1         = 2
	HoconLexerT__2         = 3
	HoconLexerT__3         = 4
	HoconLexerT__4         = 5
	HoconLexerT__5         = 6
	HoconLexerT__6         = 7
	HoconLexerCOMMENT      = 8
	HoconLexerNUMBER       = 9
	HoconLexerSTRING       = 10
	HoconLexerPATH_ELEMENT = 11
	HoconLexerREFERENCE    = 12
	HoconLexerKV           = 13
	HoconLexerWS           = 14
)