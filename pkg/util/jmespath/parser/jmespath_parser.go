// Code generated from Jmespath.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Jmespath

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

type JmespathParser struct {
	*antlr.BaseParser
}

var jmespathParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jmespathParserInit() {
	staticData := &jmespathParserStaticData
	staticData.literalNames = []string{
		"", "'.'", "'&&'", "'||'", "'!'", "'('", "')'", "'|'", "'*'", "'['",
		"']'", "'[?'", "','", "'{'", "'}'", "':'", "'$'", "'@'", "'&'", "'`'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "COMPARATOR", "RAW_STRING", "JSON_CONSTANT", "NAME", "STRING",
		"REAL_OR_EXPONENT_NUMBER", "SIGNED_INT", "INT", "WS",
	}
	staticData.ruleNames = []string{
		"singleExpression", "expression", "chainedExpression", "wildcard", "bracketSpecifier",
		"multiSelectList", "multiSelectHash", "keyvalExpr", "sliceNode", "parameterNode",
		"functionExpression", "functionArg", "currentNode", "expressionType",
		"literal", "identifier", "jsonObject", "jsonObjectPair", "jsonArray",
		"jsonValue",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 28, 230, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 61, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 5, 1, 80, 8, 1, 10, 1, 12, 1, 83, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 3, 2, 90, 8, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 110, 8, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 116, 8, 5, 10, 5, 12, 5, 119, 9, 5, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 127, 8, 6, 10, 6, 12, 6, 130, 9, 6,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 3, 8, 139, 8, 8, 1, 8, 1, 8,
		3, 8, 143, 8, 8, 1, 8, 1, 8, 3, 8, 147, 8, 8, 3, 8, 149, 8, 8, 1, 9, 1,
		9, 1, 9, 1, 9, 3, 9, 155, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10,
		162, 8, 10, 10, 10, 12, 10, 165, 9, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 3, 10, 172, 8, 10, 1, 11, 1, 11, 3, 11, 176, 8, 11, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16,
		1, 16, 1, 16, 5, 16, 193, 8, 16, 10, 16, 12, 16, 196, 9, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 3, 16, 202, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18,
		1, 18, 1, 18, 1, 18, 5, 18, 212, 8, 18, 10, 18, 12, 18, 215, 9, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 3, 18, 221, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 3, 19, 228, 8, 19, 1, 19, 0, 1, 2, 20, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 0, 2, 1, 0, 22, 24, 1,
		0, 25, 26, 252, 0, 40, 1, 0, 0, 0, 2, 60, 1, 0, 0, 0, 4, 89, 1, 0, 0, 0,
		6, 91, 1, 0, 0, 0, 8, 109, 1, 0, 0, 0, 10, 111, 1, 0, 0, 0, 12, 122, 1,
		0, 0, 0, 14, 133, 1, 0, 0, 0, 16, 138, 1, 0, 0, 0, 18, 154, 1, 0, 0, 0,
		20, 171, 1, 0, 0, 0, 22, 175, 1, 0, 0, 0, 24, 177, 1, 0, 0, 0, 26, 179,
		1, 0, 0, 0, 28, 182, 1, 0, 0, 0, 30, 186, 1, 0, 0, 0, 32, 201, 1, 0, 0,
		0, 34, 203, 1, 0, 0, 0, 36, 220, 1, 0, 0, 0, 38, 227, 1, 0, 0, 0, 40, 41,
		3, 2, 1, 0, 41, 42, 5, 0, 0, 1, 42, 1, 1, 0, 0, 0, 43, 44, 6, 1, -1, 0,
		44, 61, 3, 8, 4, 0, 45, 61, 3, 30, 15, 0, 46, 47, 5, 4, 0, 0, 47, 61, 3,
		2, 1, 11, 48, 49, 5, 5, 0, 0, 49, 50, 3, 2, 1, 0, 50, 51, 5, 6, 0, 0, 51,
		61, 1, 0, 0, 0, 52, 61, 3, 6, 3, 0, 53, 61, 3, 10, 5, 0, 54, 61, 3, 12,
		6, 0, 55, 61, 3, 28, 14, 0, 56, 61, 3, 20, 10, 0, 57, 61, 5, 21, 0, 0,
		58, 61, 3, 24, 12, 0, 59, 61, 3, 18, 9, 0, 60, 43, 1, 0, 0, 0, 60, 45,
		1, 0, 0, 0, 60, 46, 1, 0, 0, 0, 60, 48, 1, 0, 0, 0, 60, 52, 1, 0, 0, 0,
		60, 53, 1, 0, 0, 0, 60, 54, 1, 0, 0, 0, 60, 55, 1, 0, 0, 0, 60, 56, 1,
		0, 0, 0, 60, 57, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 60, 59, 1, 0, 0, 0, 61,
		81, 1, 0, 0, 0, 62, 63, 10, 15, 0, 0, 63, 64, 5, 20, 0, 0, 64, 80, 3, 2,
		1, 16, 65, 66, 10, 14, 0, 0, 66, 67, 5, 2, 0, 0, 67, 80, 3, 2, 1, 15, 68,
		69, 10, 13, 0, 0, 69, 70, 5, 3, 0, 0, 70, 80, 3, 2, 1, 14, 71, 72, 10,
		4, 0, 0, 72, 73, 5, 7, 0, 0, 73, 80, 3, 2, 1, 5, 74, 75, 10, 18, 0, 0,
		75, 76, 5, 1, 0, 0, 76, 80, 3, 4, 2, 0, 77, 78, 10, 17, 0, 0, 78, 80, 3,
		8, 4, 0, 79, 62, 1, 0, 0, 0, 79, 65, 1, 0, 0, 0, 79, 68, 1, 0, 0, 0, 79,
		71, 1, 0, 0, 0, 79, 74, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 80, 83, 1, 0, 0,
		0, 81, 79, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 3, 1, 0, 0, 0, 83, 81, 1,
		0, 0, 0, 84, 90, 3, 30, 15, 0, 85, 90, 3, 10, 5, 0, 86, 90, 3, 12, 6, 0,
		87, 90, 3, 20, 10, 0, 88, 90, 3, 6, 3, 0, 89, 84, 1, 0, 0, 0, 89, 85, 1,
		0, 0, 0, 89, 86, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 88, 1, 0, 0, 0, 90,
		5, 1, 0, 0, 0, 91, 92, 5, 8, 0, 0, 92, 7, 1, 0, 0, 0, 93, 94, 5, 9, 0,
		0, 94, 95, 5, 26, 0, 0, 95, 110, 5, 10, 0, 0, 96, 97, 5, 9, 0, 0, 97, 98,
		5, 8, 0, 0, 98, 110, 5, 10, 0, 0, 99, 100, 5, 9, 0, 0, 100, 101, 3, 16,
		8, 0, 101, 102, 5, 10, 0, 0, 102, 110, 1, 0, 0, 0, 103, 104, 5, 9, 0, 0,
		104, 110, 5, 10, 0, 0, 105, 106, 5, 11, 0, 0, 106, 107, 3, 2, 1, 0, 107,
		108, 5, 10, 0, 0, 108, 110, 1, 0, 0, 0, 109, 93, 1, 0, 0, 0, 109, 96, 1,
		0, 0, 0, 109, 99, 1, 0, 0, 0, 109, 103, 1, 0, 0, 0, 109, 105, 1, 0, 0,
		0, 110, 9, 1, 0, 0, 0, 111, 112, 5, 9, 0, 0, 112, 117, 3, 2, 1, 0, 113,
		114, 5, 12, 0, 0, 114, 116, 3, 2, 1, 0, 115, 113, 1, 0, 0, 0, 116, 119,
		1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 120, 1, 0,
		0, 0, 119, 117, 1, 0, 0, 0, 120, 121, 5, 10, 0, 0, 121, 11, 1, 0, 0, 0,
		122, 123, 5, 13, 0, 0, 123, 128, 3, 14, 7, 0, 124, 125, 5, 12, 0, 0, 125,
		127, 3, 14, 7, 0, 126, 124, 1, 0, 0, 0, 127, 130, 1, 0, 0, 0, 128, 126,
		1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 131, 1, 0, 0, 0, 130, 128, 1, 0,
		0, 0, 131, 132, 5, 14, 0, 0, 132, 13, 1, 0, 0, 0, 133, 134, 3, 30, 15,
		0, 134, 135, 5, 15, 0, 0, 135, 136, 3, 2, 1, 0, 136, 15, 1, 0, 0, 0, 137,
		139, 5, 26, 0, 0, 138, 137, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 140,
		1, 0, 0, 0, 140, 142, 5, 15, 0, 0, 141, 143, 5, 26, 0, 0, 142, 141, 1,
		0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 148, 1, 0, 0, 0, 144, 146, 5, 15, 0,
		0, 145, 147, 5, 26, 0, 0, 146, 145, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147,
		149, 1, 0, 0, 0, 148, 144, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 17, 1,
		0, 0, 0, 150, 151, 5, 16, 0, 0, 151, 155, 5, 23, 0, 0, 152, 153, 5, 16,
		0, 0, 153, 155, 5, 27, 0, 0, 154, 150, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0,
		155, 19, 1, 0, 0, 0, 156, 157, 5, 23, 0, 0, 157, 158, 5, 5, 0, 0, 158,
		163, 3, 22, 11, 0, 159, 160, 5, 12, 0, 0, 160, 162, 3, 22, 11, 0, 161,
		159, 1, 0, 0, 0, 162, 165, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164,
		1, 0, 0, 0, 164, 166, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 166, 167, 5, 6,
		0, 0, 167, 172, 1, 0, 0, 0, 168, 169, 5, 23, 0, 0, 169, 170, 5, 5, 0, 0,
		170, 172, 5, 6, 0, 0, 171, 156, 1, 0, 0, 0, 171, 168, 1, 0, 0, 0, 172,
		21, 1, 0, 0, 0, 173, 176, 3, 2, 1, 0, 174, 176, 3, 26, 13, 0, 175, 173,
		1, 0, 0, 0, 175, 174, 1, 0, 0, 0, 176, 23, 1, 0, 0, 0, 177, 178, 5, 17,
		0, 0, 178, 25, 1, 0, 0, 0, 179, 180, 5, 18, 0, 0, 180, 181, 3, 2, 1, 0,
		181, 27, 1, 0, 0, 0, 182, 183, 5, 19, 0, 0, 183, 184, 3, 38, 19, 0, 184,
		185, 5, 19, 0, 0, 185, 29, 1, 0, 0, 0, 186, 187, 7, 0, 0, 0, 187, 31, 1,
		0, 0, 0, 188, 189, 5, 13, 0, 0, 189, 194, 3, 34, 17, 0, 190, 191, 5, 12,
		0, 0, 191, 193, 3, 34, 17, 0, 192, 190, 1, 0, 0, 0, 193, 196, 1, 0, 0,
		0, 194, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 197, 1, 0, 0, 0, 196,
		194, 1, 0, 0, 0, 197, 198, 5, 14, 0, 0, 198, 202, 1, 0, 0, 0, 199, 200,
		5, 13, 0, 0, 200, 202, 5, 14, 0, 0, 201, 188, 1, 0, 0, 0, 201, 199, 1,
		0, 0, 0, 202, 33, 1, 0, 0, 0, 203, 204, 5, 24, 0, 0, 204, 205, 5, 15, 0,
		0, 205, 206, 3, 38, 19, 0, 206, 35, 1, 0, 0, 0, 207, 208, 5, 9, 0, 0, 208,
		213, 3, 38, 19, 0, 209, 210, 5, 12, 0, 0, 210, 212, 3, 38, 19, 0, 211,
		209, 1, 0, 0, 0, 212, 215, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 213, 214,
		1, 0, 0, 0, 214, 216, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 216, 217, 5, 10,
		0, 0, 217, 221, 1, 0, 0, 0, 218, 219, 5, 9, 0, 0, 219, 221, 5, 10, 0, 0,
		220, 207, 1, 0, 0, 0, 220, 218, 1, 0, 0, 0, 221, 37, 1, 0, 0, 0, 222, 228,
		5, 24, 0, 0, 223, 228, 7, 1, 0, 0, 224, 228, 3, 32, 16, 0, 225, 228, 3,
		36, 18, 0, 226, 228, 5, 22, 0, 0, 227, 222, 1, 0, 0, 0, 227, 223, 1, 0,
		0, 0, 227, 224, 1, 0, 0, 0, 227, 225, 1, 0, 0, 0, 227, 226, 1, 0, 0, 0,
		228, 39, 1, 0, 0, 0, 20, 60, 79, 81, 89, 109, 117, 128, 138, 142, 146,
		148, 154, 163, 171, 175, 194, 201, 213, 220, 227,
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

// JmespathParserInit initializes any static state used to implement JmespathParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJmespathParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JmespathParserInit() {
	staticData := &jmespathParserStaticData
	staticData.once.Do(jmespathParserInit)
}

// NewJmespathParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJmespathParser(input antlr.TokenStream) *JmespathParser {
	JmespathParserInit()
	this := new(JmespathParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &jmespathParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Jmespath.g4"

	return this
}

// JmespathParser tokens.
const (
	JmespathParserEOF                     = antlr.TokenEOF
	JmespathParserT__0                    = 1
	JmespathParserT__1                    = 2
	JmespathParserT__2                    = 3
	JmespathParserT__3                    = 4
	JmespathParserT__4                    = 5
	JmespathParserT__5                    = 6
	JmespathParserT__6                    = 7
	JmespathParserT__7                    = 8
	JmespathParserT__8                    = 9
	JmespathParserT__9                    = 10
	JmespathParserT__10                   = 11
	JmespathParserT__11                   = 12
	JmespathParserT__12                   = 13
	JmespathParserT__13                   = 14
	JmespathParserT__14                   = 15
	JmespathParserT__15                   = 16
	JmespathParserT__16                   = 17
	JmespathParserT__17                   = 18
	JmespathParserT__18                   = 19
	JmespathParserCOMPARATOR              = 20
	JmespathParserRAW_STRING              = 21
	JmespathParserJSON_CONSTANT           = 22
	JmespathParserNAME                    = 23
	JmespathParserSTRING                  = 24
	JmespathParserREAL_OR_EXPONENT_NUMBER = 25
	JmespathParserSIGNED_INT              = 26
	JmespathParserINT                     = 27
	JmespathParserWS                      = 28
)

// JmespathParser rules.
const (
	JmespathParserRULE_singleExpression   = 0
	JmespathParserRULE_expression         = 1
	JmespathParserRULE_chainedExpression  = 2
	JmespathParserRULE_wildcard           = 3
	JmespathParserRULE_bracketSpecifier   = 4
	JmespathParserRULE_multiSelectList    = 5
	JmespathParserRULE_multiSelectHash    = 6
	JmespathParserRULE_keyvalExpr         = 7
	JmespathParserRULE_sliceNode          = 8
	JmespathParserRULE_parameterNode      = 9
	JmespathParserRULE_functionExpression = 10
	JmespathParserRULE_functionArg        = 11
	JmespathParserRULE_currentNode        = 12
	JmespathParserRULE_expressionType     = 13
	JmespathParserRULE_literal            = 14
	JmespathParserRULE_identifier         = 15
	JmespathParserRULE_jsonObject         = 16
	JmespathParserRULE_jsonObjectPair     = 17
	JmespathParserRULE_jsonArray          = 18
	JmespathParserRULE_jsonValue          = 19
)

// ISingleExpressionContext is an interface to support dynamic dispatch.
type ISingleExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	EOF() antlr.TerminalNode

	// IsSingleExpressionContext differentiates from other interfaces.
	IsSingleExpressionContext()
}

type SingleExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleExpressionContext() *SingleExpressionContext {
	var p = new(SingleExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JmespathParserRULE_singleExpression
	return p
}

func (*SingleExpressionContext) IsSingleExpressionContext() {}

func NewSingleExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleExpressionContext {
	var p = new(SingleExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JmespathParserRULE_singleExpression

	return p
}

func (s *SingleExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SingleExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(JmespathParserEOF, 0)
}

func (s *SingleExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterSingleExpression(s)
	}
}

func (s *SingleExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitSingleExpression(s)
	}
}

func (s *SingleExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitSingleExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JmespathParser) SingleExpression() (localctx ISingleExpressionContext) {
	this := p
	_ = this

	localctx = NewSingleExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JmespathParserRULE_singleExpression)

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
		p.SetState(40)
		p.expression(0)
	}
	{
		p.SetState(41)
		p.Match(JmespathParserEOF)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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
	p.RuleIndex = JmespathParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JmespathParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PipeExpressionContext struct {
	*ExpressionContext
}

func NewPipeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PipeExpressionContext {
	var p = new(PipeExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *PipeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeExpressionContext) AllExpression() []IExpressionContext {
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

func (s *PipeExpressionContext) Expression(i int) IExpressionContext {
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

func (s *PipeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterPipeExpression(s)
	}
}

func (s *PipeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitPipeExpression(s)
	}
}

func (s *PipeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitPipeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParameterExpressionContext struct {
	*ExpressionContext
}

func NewParameterExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParameterExpressionContext {
	var p = new(ParameterExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ParameterExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterExpressionContext) ParameterNode() IParameterNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterNodeContext)
}

func (s *ParameterExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterParameterExpression(s)
	}
}

func (s *ParameterExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitParameterExpression(s)
	}
}

func (s *ParameterExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitParameterExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierExpressionContext struct {
	*ExpressionContext
}

func NewIdentifierExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) Identifier() IIdentifierContext {
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

func (s *IdentifierExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitIdentifierExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExpressionContext struct {
	*ExpressionContext
}

func NewNotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExpressionContext {
	var p = new(NotExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NotExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterNotExpression(s)
	}
}

func (s *NotExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitNotExpression(s)
	}
}

func (s *NotExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitNotExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type RawStringExpressionContext struct {
	*ExpressionContext
}

func NewRawStringExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RawStringExpressionContext {
	var p = new(RawStringExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *RawStringExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RawStringExpressionContext) RAW_STRING() antlr.TerminalNode {
	return s.GetToken(JmespathParserRAW_STRING, 0)
}

func (s *RawStringExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterRawStringExpression(s)
	}
}

func (s *RawStringExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitRawStringExpression(s)
	}
}

func (s *RawStringExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitRawStringExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ComparisonExpressionContext struct {
	*ExpressionContext
}

func NewComparisonExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparisonExpressionContext {
	var p = new(ComparisonExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ComparisonExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonExpressionContext) AllExpression() []IExpressionContext {
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

func (s *ComparisonExpressionContext) Expression(i int) IExpressionContext {
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

func (s *ComparisonExpressionContext) COMPARATOR() antlr.TerminalNode {
	return s.GetToken(JmespathParserCOMPARATOR, 0)
}

func (s *ComparisonExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterComparisonExpression(s)
	}
}

func (s *ComparisonExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitComparisonExpression(s)
	}
}

func (s *ComparisonExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitComparisonExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExpressionContext struct {
	*ExpressionContext
}

func NewParenExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExpressionContext {
	var p = new(ParenExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ParenExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterParenExpression(s)
	}
}

func (s *ParenExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitParenExpression(s)
	}
}

func (s *ParenExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitParenExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type BracketExpressionContext struct {
	*ExpressionContext
}

func NewBracketExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BracketExpressionContext {
	var p = new(BracketExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BracketExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketExpressionContext) BracketSpecifier() IBracketSpecifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBracketSpecifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBracketSpecifierContext)
}

func (s *BracketExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterBracketExpression(s)
	}
}

func (s *BracketExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitBracketExpression(s)
	}
}

func (s *BracketExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitBracketExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrExpressionContext struct {
	*ExpressionContext
}

func NewOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExpressionContext {
	var p = new(OrExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *OrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExpressionContext) AllExpression() []IExpressionContext {
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

func (s *OrExpressionContext) Expression(i int) IExpressionContext {
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

func (s *OrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterOrExpression(s)
	}
}

func (s *OrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitOrExpression(s)
	}
}

func (s *OrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitOrExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type CurrentNodeExpressionContext struct {
	*ExpressionContext
}

func NewCurrentNodeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CurrentNodeExpressionContext {
	var p = new(CurrentNodeExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CurrentNodeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CurrentNodeExpressionContext) CurrentNode() ICurrentNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICurrentNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICurrentNodeContext)
}

func (s *CurrentNodeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterCurrentNodeExpression(s)
	}
}

func (s *CurrentNodeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitCurrentNodeExpression(s)
	}
}

func (s *CurrentNodeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitCurrentNodeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ChainExpressionContext struct {
	*ExpressionContext
}

func NewChainExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ChainExpressionContext {
	var p = new(ChainExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ChainExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChainExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ChainExpressionContext) ChainedExpression() IChainedExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChainedExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChainedExpressionContext)
}

func (s *ChainExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterChainExpression(s)
	}
}

func (s *ChainExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitChainExpression(s)
	}
}

func (s *ChainExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitChainExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndExpressionContext struct {
	*ExpressionContext
}

func NewAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExpressionContext {
	var p = new(AndExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExpressionContext) AllExpression() []IExpressionContext {
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

func (s *AndExpressionContext) Expression(i int) IExpressionContext {
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

func (s *AndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterAndExpression(s)
	}
}

func (s *AndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitAndExpression(s)
	}
}

func (s *AndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitAndExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultiSelectHashExpressionContext struct {
	*ExpressionContext
}

func NewMultiSelectHashExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiSelectHashExpressionContext {
	var p = new(MultiSelectHashExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MultiSelectHashExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiSelectHashExpressionContext) MultiSelectHash() IMultiSelectHashContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiSelectHashContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiSelectHashContext)
}

func (s *MultiSelectHashExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterMultiSelectHashExpression(s)
	}
}

func (s *MultiSelectHashExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitMultiSelectHashExpression(s)
	}
}

func (s *MultiSelectHashExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitMultiSelectHashExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type WildcardExpressionContext struct {
	*ExpressionContext
}

func NewWildcardExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WildcardExpressionContext {
	var p = new(WildcardExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *WildcardExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WildcardExpressionContext) Wildcard() IWildcardContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWildcardContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWildcardContext)
}

func (s *WildcardExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterWildcardExpression(s)
	}
}

func (s *WildcardExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitWildcardExpression(s)
	}
}

func (s *WildcardExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitWildcardExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallExpressionContext struct {
	*ExpressionContext
}

func NewFunctionCallExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallExpressionContext {
	var p = new(FunctionCallExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *FunctionCallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallExpressionContext) FunctionExpression() IFunctionExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionExpressionContext)
}

func (s *FunctionCallExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterFunctionCallExpression(s)
	}
}

func (s *FunctionCallExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitFunctionCallExpression(s)
	}
}

func (s *FunctionCallExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitFunctionCallExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultiSelectListExpressionContext struct {
	*ExpressionContext
}

func NewMultiSelectListExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiSelectListExpressionContext {
	var p = new(MultiSelectListExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MultiSelectListExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiSelectListExpressionContext) MultiSelectList() IMultiSelectListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiSelectListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiSelectListContext)
}

func (s *MultiSelectListExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterMultiSelectListExpression(s)
	}
}

func (s *MultiSelectListExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitMultiSelectListExpression(s)
	}
}

func (s *MultiSelectListExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitMultiSelectListExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type BracketedExpressionContext struct {
	*ExpressionContext
}

func NewBracketedExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BracketedExpressionContext {
	var p = new(BracketedExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BracketedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketedExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BracketedExpressionContext) BracketSpecifier() IBracketSpecifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBracketSpecifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBracketSpecifierContext)
}

func (s *BracketedExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterBracketedExpression(s)
	}
}

func (s *BracketedExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitBracketedExpression(s)
	}
}

func (s *BracketedExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitBracketedExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralExpressionContext struct {
	*ExpressionContext
}

func NewLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExpressionContext {
	var p = new(LiteralExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExpressionContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitLiteralExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JmespathParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *JmespathParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, JmespathParserRULE_expression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBracketExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(44)
			p.BracketSpecifier()
		}

	case 2:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(45)
			p.Identifier()
		}

	case 3:
		localctx = NewNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(46)
			p.Match(JmespathParserT__3)
		}
		{
			p.SetState(47)
			p.expression(11)
		}

	case 4:
		localctx = NewParenExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(48)
			p.Match(JmespathParserT__4)
		}
		{
			p.SetState(49)
			p.expression(0)
		}
		{
			p.SetState(50)
			p.Match(JmespathParserT__5)
		}

	case 5:
		localctx = NewWildcardExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(52)
			p.Wildcard()
		}

	case 6:
		localctx = NewMultiSelectListExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(53)
			p.MultiSelectList()
		}

	case 7:
		localctx = NewMultiSelectHashExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(54)
			p.MultiSelectHash()
		}

	case 8:
		localctx = NewLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(55)
			p.Literal()
		}

	case 9:
		localctx = NewFunctionCallExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(56)
			p.FunctionExpression()
		}

	case 10:
		localctx = NewRawStringExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(57)
			p.Match(JmespathParserRAW_STRING)
		}

	case 11:
		localctx = NewCurrentNodeExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(58)
			p.CurrentNode()
		}

	case 12:
		localctx = NewParameterExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(59)
			p.ParameterNode()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(79)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewComparisonExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JmespathParserRULE_expression)
				p.SetState(62)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(63)
					p.Match(JmespathParserCOMPARATOR)
				}
				{
					p.SetState(64)
					p.expression(16)
				}

			case 2:
				localctx = NewAndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JmespathParserRULE_expression)
				p.SetState(65)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(66)
					p.Match(JmespathParserT__1)
				}
				{
					p.SetState(67)
					p.expression(15)
				}

			case 3:
				localctx = NewOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JmespathParserRULE_expression)
				p.SetState(68)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(69)
					p.Match(JmespathParserT__2)
				}
				{
					p.SetState(70)
					p.expression(14)
				}

			case 4:
				localctx = NewPipeExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JmespathParserRULE_expression)
				p.SetState(71)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(72)
					p.Match(JmespathParserT__6)
				}
				{
					p.SetState(73)
					p.expression(5)
				}

			case 5:
				localctx = NewChainExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JmespathParserRULE_expression)
				p.SetState(74)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				{
					p.SetState(75)
					p.Match(JmespathParserT__0)
				}
				{
					p.SetState(76)
					p.ChainedExpression()
				}

			case 6:
				localctx = NewBracketedExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JmespathParserRULE_expression)
				p.SetState(77)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(78)
					p.BracketSpecifier()
				}

			}

		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IChainedExpressionContext is an interface to support dynamic dispatch.
type IChainedExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	MultiSelectList() IMultiSelectListContext
	MultiSelectHash() IMultiSelectHashContext
	FunctionExpression() IFunctionExpressionContext
	Wildcard() IWildcardContext

	// IsChainedExpressionContext differentiates from other interfaces.
	IsChainedExpressionContext()
}

type ChainedExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChainedExpressionContext() *ChainedExpressionContext {
	var p = new(ChainedExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JmespathParserRULE_chainedExpression
	return p
}

func (*ChainedExpressionContext) IsChainedExpressionContext() {}

func NewChainedExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChainedExpressionContext {
	var p = new(ChainedExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JmespathParserRULE_chainedExpression

	return p
}

func (s *ChainedExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ChainedExpressionContext) Identifier() IIdentifierContext {
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

func (s *ChainedExpressionContext) MultiSelectList() IMultiSelectListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiSelectListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiSelectListContext)
}

func (s *ChainedExpressionContext) MultiSelectHash() IMultiSelectHashContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiSelectHashContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiSelectHashContext)
}

func (s *ChainedExpressionContext) FunctionExpression() IFunctionExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionExpressionContext)
}

func (s *ChainedExpressionContext) Wildcard() IWildcardContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWildcardContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWildcardContext)
}

func (s *ChainedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChainedExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChainedExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterChainedExpression(s)
	}
}

func (s *ChainedExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitChainedExpression(s)
	}
}

func (s *ChainedExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitChainedExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JmespathParser) ChainedExpression() (localctx IChainedExpressionContext) {
	this := p
	_ = this

	localctx = NewChainedExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JmespathParserRULE_chainedExpression)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.Identifier()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.MultiSelectList()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(86)
			p.MultiSelectHash()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(87)
			p.FunctionExpression()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(88)
			p.Wildcard()
		}

	}

	return localctx
}

// IWildcardContext is an interface to support dynamic dispatch.
type IWildcardContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsWildcardContext differentiates from other interfaces.
	IsWildcardContext()
}

type WildcardContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWildcardContext() *WildcardContext {
	var p = new(WildcardContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JmespathParserRULE_wildcard
	return p
}

func (*WildcardContext) IsWildcardContext() {}

func NewWildcardContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WildcardContext {
	var p = new(WildcardContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JmespathParserRULE_wildcard

	return p
}

func (s *WildcardContext) GetParser() antlr.Parser { return s.parser }
func (s *WildcardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WildcardContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WildcardContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterWildcard(s)
	}
}

func (s *WildcardContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitWildcard(s)
	}
}

func (s *WildcardContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitWildcard(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JmespathParser) Wildcard() (localctx IWildcardContext) {
	this := p
	_ = this

	localctx = NewWildcardContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JmespathParserRULE_wildcard)

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
		p.Match(JmespathParserT__7)
	}

	return localctx
}

// IBracketSpecifierContext is an interface to support dynamic dispatch.
type IBracketSpecifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBracketSpecifierContext differentiates from other interfaces.
	IsBracketSpecifierContext()
}

type BracketSpecifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBracketSpecifierContext() *BracketSpecifierContext {
	var p = new(BracketSpecifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JmespathParserRULE_bracketSpecifier
	return p
}

func (*BracketSpecifierContext) IsBracketSpecifierContext() {}

func NewBracketSpecifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BracketSpecifierContext {
	var p = new(BracketSpecifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JmespathParserRULE_bracketSpecifier

	return p
}

func (s *BracketSpecifierContext) GetParser() antlr.Parser { return s.parser }

func (s *BracketSpecifierContext) CopyFrom(ctx *BracketSpecifierContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *BracketSpecifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketSpecifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SelectContext struct {
	*BracketSpecifierContext
}

func NewSelectContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectContext {
	var p = new(SelectContext)

	p.BracketSpecifierContext = NewEmptyBracketSpecifierContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BracketSpecifierContext))

	return p
}

func (s *SelectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SelectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterSelect(s)
	}
}

func (s *SelectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitSelect(s)
	}
}

func (s *SelectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitSelect(s)

	default:
		return t.VisitChildren(s)
	}
}

type BracketFlattenContext struct {
	*BracketSpecifierContext
}

func NewBracketFlattenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BracketFlattenContext {
	var p = new(BracketFlattenContext)

	p.BracketSpecifierContext = NewEmptyBracketSpecifierContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BracketSpecifierContext))

	return p
}

func (s *BracketFlattenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketFlattenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterBracketFlatten(s)
	}
}

func (s *BracketFlattenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitBracketFlatten(s)
	}
}

func (s *BracketFlattenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitBracketFlatten(s)

	default:
		return t.VisitChildren(s)
	}
}

type BracketSliceContext struct {
	*BracketSpecifierContext
}

func NewBracketSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BracketSliceContext {
	var p = new(BracketSliceContext)

	p.BracketSpecifierContext = NewEmptyBracketSpecifierContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BracketSpecifierContext))

	return p
}

func (s *BracketSliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketSliceContext) SliceNode() ISliceNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceNodeContext)
}

func (s *BracketSliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterBracketSlice(s)
	}
}

func (s *BracketSliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitBracketSlice(s)
	}
}

func (s *BracketSliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitBracketSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

type BracketIndexContext struct {
	*BracketSpecifierContext
}

func NewBracketIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BracketIndexContext {
	var p = new(BracketIndexContext)

	p.BracketSpecifierContext = NewEmptyBracketSpecifierContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BracketSpecifierContext))

	return p
}

func (s *BracketIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketIndexContext) SIGNED_INT() antlr.TerminalNode {
	return s.GetToken(JmespathParserSIGNED_INT, 0)
}

func (s *BracketIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterBracketIndex(s)
	}
}

func (s *BracketIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitBracketIndex(s)
	}
}

func (s *BracketIndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitBracketIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

type BracketStarContext struct {
	*BracketSpecifierContext
}

func NewBracketStarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BracketStarContext {
	var p = new(BracketStarContext)

	p.BracketSpecifierContext = NewEmptyBracketSpecifierContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BracketSpecifierContext))

	return p
}

func (s *BracketStarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketStarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterBracketStar(s)
	}
}

func (s *BracketStarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitBracketStar(s)
	}
}

func (s *BracketStarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitBracketStar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JmespathParser) BracketSpecifier() (localctx IBracketSpecifierContext) {
	this := p
	_ = this

	localctx = NewBracketSpecifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JmespathParserRULE_bracketSpecifier)

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

	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBracketIndexContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(93)
			p.Match(JmespathParserT__8)
		}
		{
			p.SetState(94)
			p.Match(JmespathParserSIGNED_INT)
		}
		{
			p.SetState(95)
			p.Match(JmespathParserT__9)
		}

	case 2:
		localctx = NewBracketStarContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(96)
			p.Match(JmespathParserT__8)
		}
		{
			p.SetState(97)
			p.Match(JmespathParserT__7)
		}
		{
			p.SetState(98)
			p.Match(JmespathParserT__9)
		}

	case 3:
		localctx = NewBracketSliceContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(99)
			p.Match(JmespathParserT__8)
		}
		{
			p.SetState(100)
			p.SliceNode()
		}
		{
			p.SetState(101)
			p.Match(JmespathParserT__9)
		}

	case 4:
		localctx = NewBracketFlattenContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(103)
			p.Match(JmespathParserT__8)
		}
		{
			p.SetState(104)
			p.Match(JmespathParserT__9)
		}

	case 5:
		localctx = NewSelectContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(105)
			p.Match(JmespathParserT__10)
		}
		{
			p.SetState(106)
			p.expression(0)
		}
		{
			p.SetState(107)
			p.Match(JmespathParserT__9)
		}

	}

	return localctx
}

// IMultiSelectListContext is an interface to support dynamic dispatch.
type IMultiSelectListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsMultiSelectListContext differentiates from other interfaces.
	IsMultiSelectListContext()
}

type MultiSelectListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiSelectListContext() *MultiSelectListContext {
	var p = new(MultiSelectListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JmespathParserRULE_multiSelectList
	return p
}

func (*MultiSelectListContext) IsMultiSelectListContext() {}

func NewMultiSelectListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiSelectListContext {
	var p = new(MultiSelectListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JmespathParserRULE_multiSelectList

	return p
}

func (s *MultiSelectListContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiSelectListContext) AllExpression() []IExpressionContext {
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

func (s *MultiSelectListContext) Expression(i int) IExpressionContext {
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

func (s *MultiSelectListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiSelectListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiSelectListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterMultiSelectList(s)
	}
}

func (s *MultiSelectListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitMultiSelectList(s)
	}
}

func (s *MultiSelectListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitMultiSelectList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JmespathParser) MultiSelectList() (localctx IMultiSelectListContext) {
	this := p
	_ = this

	localctx = NewMultiSelectListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JmespathParserRULE_multiSelectList)
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
		p.SetState(111)
		p.Match(JmespathParserT__8)
	}
	{
		p.SetState(112)
		p.expression(0)
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JmespathParserT__11 {
		{
			p.SetState(113)
			p.Match(JmespathParserT__11)
		}
		{
			p.SetState(114)
			p.expression(0)
		}

		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(120)
		p.Match(JmespathParserT__9)
	}

	return localctx
}

// IMultiSelectHashContext is an interface to support dynamic dispatch.
type IMultiSelectHashContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllKeyvalExpr() []IKeyvalExprContext
	KeyvalExpr(i int) IKeyvalExprContext

	// IsMultiSelectHashContext differentiates from other interfaces.
	IsMultiSelectHashContext()
}

type MultiSelectHashContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiSelectHashContext() *MultiSelectHashContext {
	var p = new(MultiSelectHashContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JmespathParserRULE_multiSelectHash
	return p
}

func (*MultiSelectHashContext) IsMultiSelectHashContext() {}

func NewMultiSelectHashContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiSelectHashContext {
	var p = new(MultiSelectHashContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JmespathParserRULE_multiSelectHash

	return p
}

func (s *MultiSelectHashContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiSelectHashContext) AllKeyvalExpr() []IKeyvalExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IKeyvalExprContext); ok {
			len++
		}
	}

	tst := make([]IKeyvalExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IKeyvalExprContext); ok {
			tst[i] = t.(IKeyvalExprContext)
			i++
		}
	}

	return tst
}

func (s *MultiSelectHashContext) KeyvalExpr(i int) IKeyvalExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyvalExprContext); ok {
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

	return t.(IKeyvalExprContext)
}

func (s *MultiSelectHashContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiSelectHashContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiSelectHashContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterMultiSelectHash(s)
	}
}

func (s *MultiSelectHashContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitMultiSelectHash(s)
	}
}

func (s *MultiSelectHashContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitMultiSelectHash(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JmespathParser) MultiSelectHash() (localctx IMultiSelectHashContext) {
	this := p
	_ = this

	localctx = NewMultiSelectHashContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JmespathParserRULE_multiSelectHash)
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
		p.SetState(122)
		p.Match(JmespathParserT__12)
	}
	{
		p.SetState(123)
		p.KeyvalExpr()
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JmespathParserT__11 {
		{
			p.SetState(124)
			p.Match(JmespathParserT__11)
		}
		{
			p.SetState(125)
			p.KeyvalExpr()
		}

		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(131)
		p.Match(JmespathParserT__13)
	}

	return localctx
}

// IKeyvalExprContext is an interface to support dynamic dispatch.
type IKeyvalExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	Expression() IExpressionContext

	// IsKeyvalExprContext differentiates from other interfaces.
	IsKeyvalExprContext()
}

type KeyvalExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyvalExprContext() *KeyvalExprContext {
	var p = new(KeyvalExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JmespathParserRULE_keyvalExpr
	return p
}

func (*KeyvalExprContext) IsKeyvalExprContext() {}

func NewKeyvalExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyvalExprContext {
	var p = new(KeyvalExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JmespathParserRULE_keyvalExpr

	return p
}

func (s *KeyvalExprContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyvalExprContext) Identifier() IIdentifierContext {
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

func (s *KeyvalExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *KeyvalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyvalExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyvalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterKeyvalExpr(s)
	}
}

func (s *KeyvalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitKeyvalExpr(s)
	}
}

func (s *KeyvalExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitKeyvalExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JmespathParser) KeyvalExpr() (localctx IKeyvalExprContext) {
	this := p
	_ = this

	localctx = NewKeyvalExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, JmespathParserRULE_keyvalExpr)

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
		p.SetState(133)
		p.Identifier()
	}
	{
		p.SetState(134)
		p.Match(JmespathParserT__14)
	}
	{
		p.SetState(135)
		p.expression(0)
	}

	return localctx
}

// ISliceNodeContext is an interface to support dynamic dispatch.
type ISliceNodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSliceStart returns the sliceStart token.
	GetSliceStart() antlr.Token

	// GetSliceStop returns the sliceStop token.
	GetSliceStop() antlr.Token

	// GetSliceStep returns the sliceStep token.
	GetSliceStep() antlr.Token

	// SetSliceStart sets the sliceStart token.
	SetSliceStart(antlr.Token)

	// SetSliceStop sets the sliceStop token.
	SetSliceStop(antlr.Token)

	// SetSliceStep sets the sliceStep token.
	SetSliceStep(antlr.Token)

	// Getter signatures
	AllSIGNED_INT() []antlr.TerminalNode
	SIGNED_INT(i int) antlr.TerminalNode

	// IsSliceNodeContext differentiates from other interfaces.
	IsSliceNodeContext()
}

type SliceNodeContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	sliceStart antlr.Token
	sliceStop  antlr.Token
	sliceStep  antlr.Token
}

func NewEmptySliceNodeContext() *SliceNodeContext {
	var p = new(SliceNodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JmespathParserRULE_sliceNode
	return p
}

func (*SliceNodeContext) IsSliceNodeContext() {}

func NewSliceNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceNodeContext {
	var p = new(SliceNodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JmespathParserRULE_sliceNode

	return p
}

func (s *SliceNodeContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceNodeContext) GetSliceStart() antlr.Token { return s.sliceStart }

func (s *SliceNodeContext) GetSliceStop() antlr.Token { return s.sliceStop }

func (s *SliceNodeContext) GetSliceStep() antlr.Token { return s.sliceStep }

func (s *SliceNodeContext) SetSliceStart(v antlr.Token) { s.sliceStart = v }

func (s *SliceNodeContext) SetSliceStop(v antlr.Token) { s.sliceStop = v }

func (s *SliceNodeContext) SetSliceStep(v antlr.Token) { s.sliceStep = v }

func (s *SliceNodeContext) AllSIGNED_INT() []antlr.TerminalNode {
	return s.GetTokens(JmespathParserSIGNED_INT)
}

func (s *SliceNodeContext) SIGNED_INT(i int) antlr.TerminalNode {
	return s.GetToken(JmespathParserSIGNED_INT, i)
}

func (s *SliceNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceNodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterSliceNode(s)
	}
}

func (s *SliceNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitSliceNode(s)
	}
}

func (s *SliceNodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitSliceNode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JmespathParser) SliceNode() (localctx ISliceNodeContext) {
	this := p
	_ = this

	localctx = NewSliceNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JmespathParserRULE_sliceNode)
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
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JmespathParserSIGNED_INT {
		{
			p.SetState(137)

			var _m = p.Match(JmespathParserSIGNED_INT)

			localctx.(*SliceNodeContext).sliceStart = _m
		}

	}
	{
		p.SetState(140)
		p.Match(JmespathParserT__14)
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JmespathParserSIGNED_INT {
		{
			p.SetState(141)

			var _m = p.Match(JmespathParserSIGNED_INT)

			localctx.(*SliceNodeContext).sliceStop = _m
		}

	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JmespathParserT__14 {
		{
			p.SetState(144)
			p.Match(JmespathParserT__14)
		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JmespathParserSIGNED_INT {
			{
				p.SetState(145)

				var _m = p.Match(JmespathParserSIGNED_INT)

				localctx.(*SliceNodeContext).sliceStep = _m
			}

		}

	}

	return localctx
}

// IParameterNodeContext is an interface to support dynamic dispatch.
type IParameterNodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsParameterNodeContext differentiates from other interfaces.
	IsParameterNodeContext()
}

type ParameterNodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterNodeContext() *ParameterNodeContext {
	var p = new(ParameterNodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JmespathParserRULE_parameterNode
	return p
}

func (*ParameterNodeContext) IsParameterNodeContext() {}

func NewParameterNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterNodeContext {
	var p = new(ParameterNodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JmespathParserRULE_parameterNode

	return p
}

func (s *ParameterNodeContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterNodeContext) CopyFrom(ctx *ParameterNodeContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ParameterNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterNodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NameParameterContext struct {
	*ParameterNodeContext
}

func NewNameParameterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NameParameterContext {
	var p = new(NameParameterContext)

	p.ParameterNodeContext = NewEmptyParameterNodeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParameterNodeContext))

	return p
}

func (s *NameParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameParameterContext) NAME() antlr.TerminalNode {
	return s.GetToken(JmespathParserNAME, 0)
}

func (s *NameParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterNameParameter(s)
	}
}

func (s *NameParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitNameParameter(s)
	}
}

func (s *NameParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitNameParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberParameterContext struct {
	*ParameterNodeContext
}

func NewNumberParameterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberParameterContext {
	var p = new(NumberParameterContext)

	p.ParameterNodeContext = NewEmptyParameterNodeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParameterNodeContext))

	return p
}

func (s *NumberParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberParameterContext) INT() antlr.TerminalNode {
	return s.GetToken(JmespathParserINT, 0)
}

func (s *NumberParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterNumberParameter(s)
	}
}

func (s *NumberParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitNumberParameter(s)
	}
}

func (s *NumberParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitNumberParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JmespathParser) ParameterNode() (localctx IParameterNodeContext) {
	this := p
	_ = this

	localctx = NewParameterNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, JmespathParserRULE_parameterNode)

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

	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNameParameterContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(150)
			p.Match(JmespathParserT__15)
		}
		{
			p.SetState(151)
			p.Match(JmespathParserNAME)
		}

	case 2:
		localctx = NewNumberParameterContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(152)
			p.Match(JmespathParserT__15)
		}
		{
			p.SetState(153)
			p.Match(JmespathParserINT)
		}

	}

	return localctx
}

// IFunctionExpressionContext is an interface to support dynamic dispatch.
type IFunctionExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode
	AllFunctionArg() []IFunctionArgContext
	FunctionArg(i int) IFunctionArgContext

	// IsFunctionExpressionContext differentiates from other interfaces.
	IsFunctionExpressionContext()
}

type FunctionExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionExpressionContext() *FunctionExpressionContext {
	var p = new(FunctionExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JmespathParserRULE_functionExpression
	return p
}

func (*FunctionExpressionContext) IsFunctionExpressionContext() {}

func NewFunctionExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionExpressionContext {
	var p = new(FunctionExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JmespathParserRULE_functionExpression

	return p
}

func (s *FunctionExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionExpressionContext) NAME() antlr.TerminalNode {
	return s.GetToken(JmespathParserNAME, 0)
}

func (s *FunctionExpressionContext) AllFunctionArg() []IFunctionArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionArgContext); ok {
			len++
		}
	}

	tst := make([]IFunctionArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionArgContext); ok {
			tst[i] = t.(IFunctionArgContext)
			i++
		}
	}

	return tst
}

func (s *FunctionExpressionContext) FunctionArg(i int) IFunctionArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionArgContext); ok {
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

	return t.(IFunctionArgContext)
}

func (s *FunctionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterFunctionExpression(s)
	}
}

func (s *FunctionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitFunctionExpression(s)
	}
}

func (s *FunctionExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitFunctionExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JmespathParser) FunctionExpression() (localctx IFunctionExpressionContext) {
	this := p
	_ = this

	localctx = NewFunctionExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, JmespathParserRULE_functionExpression)
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

	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(156)
			p.Match(JmespathParserNAME)
		}
		{
			p.SetState(157)
			p.Match(JmespathParserT__4)
		}
		{
			p.SetState(158)
			p.FunctionArg()
		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JmespathParserT__11 {
			{
				p.SetState(159)
				p.Match(JmespathParserT__11)
			}
			{
				p.SetState(160)
				p.FunctionArg()
			}

			p.SetState(165)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(166)
			p.Match(JmespathParserT__5)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(168)
			p.Match(JmespathParserNAME)
		}
		{
			p.SetState(169)
			p.Match(JmespathParserT__4)
		}
		{
			p.SetState(170)
			p.Match(JmespathParserT__5)
		}

	}

	return localctx
}

// IFunctionArgContext is an interface to support dynamic dispatch.
type IFunctionArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	ExpressionType() IExpressionTypeContext

	// IsFunctionArgContext differentiates from other interfaces.
	IsFunctionArgContext()
}

type FunctionArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionArgContext() *FunctionArgContext {
	var p = new(FunctionArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JmespathParserRULE_functionArg
	return p
}

func (*FunctionArgContext) IsFunctionArgContext() {}

func NewFunctionArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionArgContext {
	var p = new(FunctionArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JmespathParserRULE_functionArg

	return p
}

func (s *FunctionArgContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionArgContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FunctionArgContext) ExpressionType() IExpressionTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionTypeContext)
}

func (s *FunctionArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterFunctionArg(s)
	}
}

func (s *FunctionArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitFunctionArg(s)
	}
}

func (s *FunctionArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitFunctionArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JmespathParser) FunctionArg() (localctx IFunctionArgContext) {
	this := p
	_ = this

	localctx = NewFunctionArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, JmespathParserRULE_functionArg)

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

	p.SetState(175)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JmespathParserT__3, JmespathParserT__4, JmespathParserT__7, JmespathParserT__8, JmespathParserT__10, JmespathParserT__12, JmespathParserT__15, JmespathParserT__16, JmespathParserT__18, JmespathParserRAW_STRING, JmespathParserJSON_CONSTANT, JmespathParserNAME, JmespathParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(173)
			p.expression(0)
		}

	case JmespathParserT__17:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(174)
			p.ExpressionType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICurrentNodeContext is an interface to support dynamic dispatch.
type ICurrentNodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCurrentNodeContext differentiates from other interfaces.
	IsCurrentNodeContext()
}

type CurrentNodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCurrentNodeContext() *CurrentNodeContext {
	var p = new(CurrentNodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JmespathParserRULE_currentNode
	return p
}

func (*CurrentNodeContext) IsCurrentNodeContext() {}

func NewCurrentNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CurrentNodeContext {
	var p = new(CurrentNodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JmespathParserRULE_currentNode

	return p
}

func (s *CurrentNodeContext) GetParser() antlr.Parser { return s.parser }
func (s *CurrentNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CurrentNodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CurrentNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterCurrentNode(s)
	}
}

func (s *CurrentNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitCurrentNode(s)
	}
}

func (s *CurrentNodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitCurrentNode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JmespathParser) CurrentNode() (localctx ICurrentNodeContext) {
	this := p
	_ = this

	localctx = NewCurrentNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, JmespathParserRULE_currentNode)

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
		p.SetState(177)
		p.Match(JmespathParserT__16)
	}

	return localctx
}

// IExpressionTypeContext is an interface to support dynamic dispatch.
type IExpressionTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsExpressionTypeContext differentiates from other interfaces.
	IsExpressionTypeContext()
}

type ExpressionTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionTypeContext() *ExpressionTypeContext {
	var p = new(ExpressionTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JmespathParserRULE_expressionType
	return p
}

func (*ExpressionTypeContext) IsExpressionTypeContext() {}

func NewExpressionTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionTypeContext {
	var p = new(ExpressionTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JmespathParserRULE_expressionType

	return p
}

func (s *ExpressionTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionTypeContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterExpressionType(s)
	}
}

func (s *ExpressionTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitExpressionType(s)
	}
}

func (s *ExpressionTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitExpressionType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JmespathParser) ExpressionType() (localctx IExpressionTypeContext) {
	this := p
	_ = this

	localctx = NewExpressionTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, JmespathParserRULE_expressionType)

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
		p.SetState(179)
		p.Match(JmespathParserT__17)
	}
	{
		p.SetState(180)
		p.expression(0)
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	JsonValue() IJsonValueContext

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JmespathParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JmespathParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) JsonValue() IJsonValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJsonValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JmespathParser) Literal() (localctx ILiteralContext) {
	this := p
	_ = this

	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, JmespathParserRULE_literal)

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
		p.SetState(182)
		p.Match(JmespathParserT__18)
	}
	{
		p.SetState(183)
		p.JsonValue()
	}
	{
		p.SetState(184)
		p.Match(JmespathParserT__18)
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode
	STRING() antlr.TerminalNode
	JSON_CONSTANT() antlr.TerminalNode

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
	p.RuleIndex = JmespathParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JmespathParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) NAME() antlr.TerminalNode {
	return s.GetToken(JmespathParserNAME, 0)
}

func (s *IdentifierContext) STRING() antlr.TerminalNode {
	return s.GetToken(JmespathParserSTRING, 0)
}

func (s *IdentifierContext) JSON_CONSTANT() antlr.TerminalNode {
	return s.GetToken(JmespathParserJSON_CONSTANT, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JmespathParser) Identifier() (localctx IIdentifierContext) {
	this := p
	_ = this

	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, JmespathParserRULE_identifier)
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
		p.SetState(186)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&29360128) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IJsonObjectContext is an interface to support dynamic dispatch.
type IJsonObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllJsonObjectPair() []IJsonObjectPairContext
	JsonObjectPair(i int) IJsonObjectPairContext

	// IsJsonObjectContext differentiates from other interfaces.
	IsJsonObjectContext()
}

type JsonObjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonObjectContext() *JsonObjectContext {
	var p = new(JsonObjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JmespathParserRULE_jsonObject
	return p
}

func (*JsonObjectContext) IsJsonObjectContext() {}

func NewJsonObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonObjectContext {
	var p = new(JsonObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JmespathParserRULE_jsonObject

	return p
}

func (s *JsonObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *JsonObjectContext) AllJsonObjectPair() []IJsonObjectPairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJsonObjectPairContext); ok {
			len++
		}
	}

	tst := make([]IJsonObjectPairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJsonObjectPairContext); ok {
			tst[i] = t.(IJsonObjectPairContext)
			i++
		}
	}

	return tst
}

func (s *JsonObjectContext) JsonObjectPair(i int) IJsonObjectPairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJsonObjectPairContext); ok {
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

	return t.(IJsonObjectPairContext)
}

func (s *JsonObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JsonObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterJsonObject(s)
	}
}

func (s *JsonObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitJsonObject(s)
	}
}

func (s *JsonObjectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitJsonObject(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JmespathParser) JsonObject() (localctx IJsonObjectContext) {
	this := p
	_ = this

	localctx = NewJsonObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, JmespathParserRULE_jsonObject)
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

	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(188)
			p.Match(JmespathParserT__12)
		}
		{
			p.SetState(189)
			p.JsonObjectPair()
		}
		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JmespathParserT__11 {
			{
				p.SetState(190)
				p.Match(JmespathParserT__11)
			}
			{
				p.SetState(191)
				p.JsonObjectPair()
			}

			p.SetState(196)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(197)
			p.Match(JmespathParserT__13)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(199)
			p.Match(JmespathParserT__12)
		}
		{
			p.SetState(200)
			p.Match(JmespathParserT__13)
		}

	}

	return localctx
}

// IJsonObjectPairContext is an interface to support dynamic dispatch.
type IJsonObjectPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	JsonValue() IJsonValueContext

	// IsJsonObjectPairContext differentiates from other interfaces.
	IsJsonObjectPairContext()
}

type JsonObjectPairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonObjectPairContext() *JsonObjectPairContext {
	var p = new(JsonObjectPairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JmespathParserRULE_jsonObjectPair
	return p
}

func (*JsonObjectPairContext) IsJsonObjectPairContext() {}

func NewJsonObjectPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonObjectPairContext {
	var p = new(JsonObjectPairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JmespathParserRULE_jsonObjectPair

	return p
}

func (s *JsonObjectPairContext) GetParser() antlr.Parser { return s.parser }

func (s *JsonObjectPairContext) STRING() antlr.TerminalNode {
	return s.GetToken(JmespathParserSTRING, 0)
}

func (s *JsonObjectPairContext) JsonValue() IJsonValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJsonValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
}

func (s *JsonObjectPairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonObjectPairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JsonObjectPairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterJsonObjectPair(s)
	}
}

func (s *JsonObjectPairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitJsonObjectPair(s)
	}
}

func (s *JsonObjectPairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitJsonObjectPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JmespathParser) JsonObjectPair() (localctx IJsonObjectPairContext) {
	this := p
	_ = this

	localctx = NewJsonObjectPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, JmespathParserRULE_jsonObjectPair)

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
		p.SetState(203)
		p.Match(JmespathParserSTRING)
	}
	{
		p.SetState(204)
		p.Match(JmespathParserT__14)
	}
	{
		p.SetState(205)
		p.JsonValue()
	}

	return localctx
}

// IJsonArrayContext is an interface to support dynamic dispatch.
type IJsonArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllJsonValue() []IJsonValueContext
	JsonValue(i int) IJsonValueContext

	// IsJsonArrayContext differentiates from other interfaces.
	IsJsonArrayContext()
}

type JsonArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonArrayContext() *JsonArrayContext {
	var p = new(JsonArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JmespathParserRULE_jsonArray
	return p
}

func (*JsonArrayContext) IsJsonArrayContext() {}

func NewJsonArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonArrayContext {
	var p = new(JsonArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JmespathParserRULE_jsonArray

	return p
}

func (s *JsonArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *JsonArrayContext) AllJsonValue() []IJsonValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJsonValueContext); ok {
			len++
		}
	}

	tst := make([]IJsonValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJsonValueContext); ok {
			tst[i] = t.(IJsonValueContext)
			i++
		}
	}

	return tst
}

func (s *JsonArrayContext) JsonValue(i int) IJsonValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJsonValueContext); ok {
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

	return t.(IJsonValueContext)
}

func (s *JsonArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JsonArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterJsonArray(s)
	}
}

func (s *JsonArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitJsonArray(s)
	}
}

func (s *JsonArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitJsonArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JmespathParser) JsonArray() (localctx IJsonArrayContext) {
	this := p
	_ = this

	localctx = NewJsonArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, JmespathParserRULE_jsonArray)
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

	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(207)
			p.Match(JmespathParserT__8)
		}
		{
			p.SetState(208)
			p.JsonValue()
		}
		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JmespathParserT__11 {
			{
				p.SetState(209)
				p.Match(JmespathParserT__11)
			}
			{
				p.SetState(210)
				p.JsonValue()
			}

			p.SetState(215)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(216)
			p.Match(JmespathParserT__9)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(218)
			p.Match(JmespathParserT__8)
		}
		{
			p.SetState(219)
			p.Match(JmespathParserT__9)
		}

	}

	return localctx
}

// IJsonValueContext is an interface to support dynamic dispatch.
type IJsonValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsJsonValueContext differentiates from other interfaces.
	IsJsonValueContext()
}

type JsonValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonValueContext() *JsonValueContext {
	var p = new(JsonValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JmespathParserRULE_jsonValue
	return p
}

func (*JsonValueContext) IsJsonValueContext() {}

func NewJsonValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonValueContext {
	var p = new(JsonValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JmespathParserRULE_jsonValue

	return p
}

func (s *JsonValueContext) GetParser() antlr.Parser { return s.parser }

func (s *JsonValueContext) CopyFrom(ctx *JsonValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *JsonValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type JsonArrayValueContext struct {
	*JsonValueContext
}

func NewJsonArrayValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonArrayValueContext {
	var p = new(JsonArrayValueContext)

	p.JsonValueContext = NewEmptyJsonValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValueContext))

	return p
}

func (s *JsonArrayValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonArrayValueContext) JsonArray() IJsonArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJsonArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJsonArrayContext)
}

func (s *JsonArrayValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterJsonArrayValue(s)
	}
}

func (s *JsonArrayValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitJsonArrayValue(s)
	}
}

func (s *JsonArrayValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitJsonArrayValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type JsonStringValueContext struct {
	*JsonValueContext
}

func NewJsonStringValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonStringValueContext {
	var p = new(JsonStringValueContext)

	p.JsonValueContext = NewEmptyJsonValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValueContext))

	return p
}

func (s *JsonStringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonStringValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(JmespathParserSTRING, 0)
}

func (s *JsonStringValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterJsonStringValue(s)
	}
}

func (s *JsonStringValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitJsonStringValue(s)
	}
}

func (s *JsonStringValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitJsonStringValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type JsonObjectValueContext struct {
	*JsonValueContext
}

func NewJsonObjectValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonObjectValueContext {
	var p = new(JsonObjectValueContext)

	p.JsonValueContext = NewEmptyJsonValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValueContext))

	return p
}

func (s *JsonObjectValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonObjectValueContext) JsonObject() IJsonObjectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJsonObjectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJsonObjectContext)
}

func (s *JsonObjectValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterJsonObjectValue(s)
	}
}

func (s *JsonObjectValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitJsonObjectValue(s)
	}
}

func (s *JsonObjectValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitJsonObjectValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type JsonConstantValueContext struct {
	*JsonValueContext
}

func NewJsonConstantValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonConstantValueContext {
	var p = new(JsonConstantValueContext)

	p.JsonValueContext = NewEmptyJsonValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValueContext))

	return p
}

func (s *JsonConstantValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonConstantValueContext) JSON_CONSTANT() antlr.TerminalNode {
	return s.GetToken(JmespathParserJSON_CONSTANT, 0)
}

func (s *JsonConstantValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterJsonConstantValue(s)
	}
}

func (s *JsonConstantValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitJsonConstantValue(s)
	}
}

func (s *JsonConstantValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitJsonConstantValue(s)

	default:
		return t.VisitChildren(s)
	}
}

type JsonNumberValueContext struct {
	*JsonValueContext
}

func NewJsonNumberValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonNumberValueContext {
	var p = new(JsonNumberValueContext)

	p.JsonValueContext = NewEmptyJsonValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValueContext))

	return p
}

func (s *JsonNumberValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonNumberValueContext) REAL_OR_EXPONENT_NUMBER() antlr.TerminalNode {
	return s.GetToken(JmespathParserREAL_OR_EXPONENT_NUMBER, 0)
}

func (s *JsonNumberValueContext) SIGNED_INT() antlr.TerminalNode {
	return s.GetToken(JmespathParserSIGNED_INT, 0)
}

func (s *JsonNumberValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.EnterJsonNumberValue(s)
	}
}

func (s *JsonNumberValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JmespathListener); ok {
		listenerT.ExitJsonNumberValue(s)
	}
}

func (s *JsonNumberValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JmespathVisitor:
		return t.VisitJsonNumberValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JmespathParser) JsonValue() (localctx IJsonValueContext) {
	this := p
	_ = this

	localctx = NewJsonValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, JmespathParserRULE_jsonValue)
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

	p.SetState(227)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JmespathParserSTRING:
		localctx = NewJsonStringValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(222)
			p.Match(JmespathParserSTRING)
		}

	case JmespathParserREAL_OR_EXPONENT_NUMBER, JmespathParserSIGNED_INT:
		localctx = NewJsonNumberValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(223)
			_la = p.GetTokenStream().LA(1)

			if !(_la == JmespathParserREAL_OR_EXPONENT_NUMBER || _la == JmespathParserSIGNED_INT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case JmespathParserT__12:
		localctx = NewJsonObjectValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(224)
			p.JsonObject()
		}

	case JmespathParserT__8:
		localctx = NewJsonArrayValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(225)
			p.JsonArray()
		}

	case JmespathParserJSON_CONSTANT:
		localctx = NewJsonConstantValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(226)
			p.Match(JmespathParserJSON_CONSTANT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *JmespathParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *JmespathParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 17)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
