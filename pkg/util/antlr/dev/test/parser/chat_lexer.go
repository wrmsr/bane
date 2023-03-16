//go:build !nodev

// Code generated from Chat.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ChatLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var chatlexerLexerStaticData struct {
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

func chatlexerLexerInit() {
	staticData := &chatlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "':'", "'-'", "')'", "'('", "'/'", "'@'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "SAYS", "SHOUTS", "WORD", "WHITESPACE",
		"NEWLINE", "TEXT", "BLOCK_COMMENT", "COMMENT",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "A", "S", "Y", "H",
		"O", "U", "T", "LOWERCASE", "UPPERCASE", "SAYS", "SHOUTS", "WORD", "WHITESPACE",
		"NEWLINE", "TEXT", "BLOCK_COMMENT", "COMMENT",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 14, 143, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 4, 17, 93, 8, 17, 11, 17, 12, 17,
		94, 1, 18, 4, 18, 98, 8, 18, 11, 18, 12, 18, 99, 1, 19, 3, 19, 103, 8,
		19, 1, 19, 1, 19, 4, 19, 107, 8, 19, 11, 19, 12, 19, 108, 1, 20, 1, 20,
		4, 20, 113, 8, 20, 11, 20, 12, 20, 114, 1, 20, 1, 20, 1, 21, 1, 21, 1,
		21, 1, 21, 5, 21, 123, 8, 21, 10, 21, 12, 21, 126, 9, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 137, 8, 22, 10,
		22, 12, 22, 140, 9, 22, 1, 22, 1, 22, 1, 124, 0, 23, 1, 1, 3, 2, 5, 3,
		7, 4, 9, 5, 11, 6, 13, 0, 15, 0, 17, 0, 19, 0, 21, 0, 23, 0, 25, 0, 27,
		0, 29, 0, 31, 7, 33, 8, 35, 9, 37, 10, 39, 11, 41, 12, 43, 13, 45, 14,
		1, 0, 13, 2, 0, 65, 65, 97, 97, 2, 0, 83, 83, 115, 115, 2, 0, 89, 89, 121,
		121, 2, 0, 72, 72, 104, 104, 2, 0, 79, 79, 111, 111, 2, 0, 85, 85, 117,
		117, 2, 0, 84, 84, 116, 116, 1, 0, 97, 122, 1, 0, 65, 90, 2, 0, 9, 9, 32,
		32, 2, 0, 40, 40, 91, 91, 2, 0, 41, 41, 93, 93, 2, 0, 10, 10, 13, 13, 143,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0,
		0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0,
		0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 1, 47, 1, 0, 0, 0, 3, 49, 1,
		0, 0, 0, 5, 51, 1, 0, 0, 0, 7, 53, 1, 0, 0, 0, 9, 55, 1, 0, 0, 0, 11, 57,
		1, 0, 0, 0, 13, 59, 1, 0, 0, 0, 15, 61, 1, 0, 0, 0, 17, 63, 1, 0, 0, 0,
		19, 65, 1, 0, 0, 0, 21, 67, 1, 0, 0, 0, 23, 69, 1, 0, 0, 0, 25, 71, 1,
		0, 0, 0, 27, 73, 1, 0, 0, 0, 29, 75, 1, 0, 0, 0, 31, 77, 1, 0, 0, 0, 33,
		82, 1, 0, 0, 0, 35, 92, 1, 0, 0, 0, 37, 97, 1, 0, 0, 0, 39, 106, 1, 0,
		0, 0, 41, 110, 1, 0, 0, 0, 43, 118, 1, 0, 0, 0, 45, 132, 1, 0, 0, 0, 47,
		48, 5, 58, 0, 0, 48, 2, 1, 0, 0, 0, 49, 50, 5, 45, 0, 0, 50, 4, 1, 0, 0,
		0, 51, 52, 5, 41, 0, 0, 52, 6, 1, 0, 0, 0, 53, 54, 5, 40, 0, 0, 54, 8,
		1, 0, 0, 0, 55, 56, 5, 47, 0, 0, 56, 10, 1, 0, 0, 0, 57, 58, 5, 64, 0,
		0, 58, 12, 1, 0, 0, 0, 59, 60, 7, 0, 0, 0, 60, 14, 1, 0, 0, 0, 61, 62,
		7, 1, 0, 0, 62, 16, 1, 0, 0, 0, 63, 64, 7, 2, 0, 0, 64, 18, 1, 0, 0, 0,
		65, 66, 7, 3, 0, 0, 66, 20, 1, 0, 0, 0, 67, 68, 7, 4, 0, 0, 68, 22, 1,
		0, 0, 0, 69, 70, 7, 5, 0, 0, 70, 24, 1, 0, 0, 0, 71, 72, 7, 6, 0, 0, 72,
		26, 1, 0, 0, 0, 73, 74, 7, 7, 0, 0, 74, 28, 1, 0, 0, 0, 75, 76, 7, 8, 0,
		0, 76, 30, 1, 0, 0, 0, 77, 78, 3, 15, 7, 0, 78, 79, 3, 13, 6, 0, 79, 80,
		3, 17, 8, 0, 80, 81, 3, 15, 7, 0, 81, 32, 1, 0, 0, 0, 82, 83, 3, 15, 7,
		0, 83, 84, 3, 19, 9, 0, 84, 85, 3, 21, 10, 0, 85, 86, 3, 23, 11, 0, 86,
		87, 3, 25, 12, 0, 87, 88, 3, 15, 7, 0, 88, 34, 1, 0, 0, 0, 89, 93, 3, 27,
		13, 0, 90, 93, 3, 29, 14, 0, 91, 93, 5, 95, 0, 0, 92, 89, 1, 0, 0, 0, 92,
		90, 1, 0, 0, 0, 92, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 92, 1, 0, 0,
		0, 94, 95, 1, 0, 0, 0, 95, 36, 1, 0, 0, 0, 96, 98, 7, 9, 0, 0, 97, 96,
		1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0,
		100, 38, 1, 0, 0, 0, 101, 103, 5, 13, 0, 0, 102, 101, 1, 0, 0, 0, 102,
		103, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 107, 5, 10, 0, 0, 105, 107,
		5, 13, 0, 0, 106, 102, 1, 0, 0, 0, 106, 105, 1, 0, 0, 0, 107, 108, 1, 0,
		0, 0, 108, 106, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 40, 1, 0, 0, 0,
		110, 112, 7, 10, 0, 0, 111, 113, 8, 11, 0, 0, 112, 111, 1, 0, 0, 0, 113,
		114, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 116,
		1, 0, 0, 0, 116, 117, 7, 11, 0, 0, 117, 42, 1, 0, 0, 0, 118, 119, 5, 47,
		0, 0, 119, 120, 5, 42, 0, 0, 120, 124, 1, 0, 0, 0, 121, 123, 9, 0, 0, 0,
		122, 121, 1, 0, 0, 0, 123, 126, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 124,
		122, 1, 0, 0, 0, 125, 127, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 127, 128,
		5, 42, 0, 0, 128, 129, 5, 47, 0, 0, 129, 130, 1, 0, 0, 0, 130, 131, 6,
		21, 0, 0, 131, 44, 1, 0, 0, 0, 132, 133, 5, 47, 0, 0, 133, 134, 5, 47,
		0, 0, 134, 138, 1, 0, 0, 0, 135, 137, 8, 12, 0, 0, 136, 135, 1, 0, 0, 0,
		137, 140, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139,
		141, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141, 142, 6, 22, 0, 0, 142, 46,
		1, 0, 0, 0, 10, 0, 92, 94, 99, 102, 106, 108, 114, 124, 138, 1, 0, 2, 0,
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

// ChatLexerInit initializes any static state used to implement ChatLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewChatLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ChatLexerInit() {
	staticData := &chatlexerLexerStaticData
	staticData.once.Do(chatlexerLexerInit)
}

// NewChatLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewChatLexer(input antlr.CharStream) *ChatLexer {
	ChatLexerInit()
	l := new(ChatLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &chatlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Chat.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ChatLexer tokens.
const (
	ChatLexerT__0          = 1
	ChatLexerT__1          = 2
	ChatLexerT__2          = 3
	ChatLexerT__3          = 4
	ChatLexerT__4          = 5
	ChatLexerT__5          = 6
	ChatLexerSAYS          = 7
	ChatLexerSHOUTS        = 8
	ChatLexerWORD          = 9
	ChatLexerWHITESPACE    = 10
	ChatLexerNEWLINE       = 11
	ChatLexerTEXT          = 12
	ChatLexerBLOCK_COMMENT = 13
	ChatLexerCOMMENT       = 14
)
