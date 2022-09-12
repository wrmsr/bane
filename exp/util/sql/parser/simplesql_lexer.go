// Code generated from SimpleSql.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type SimpleSqlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var simplesqllexerLexerStaticData struct {
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

func simplesqllexerLexerInit() {
	staticData := &simplesqllexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "';'", "','", "'('", "')'", "'*'", "'::'", "'.'", "'='", "'!='",
		"'<>'", "'<'", "'<='", "'>'", "'>='", "'+'", "'-'", "'/'", "'%'", "'||'",
		"'all'", "'and'", "'as'", "'asc'", "'by'", "'case'", "'cross'", "'desc'",
		"'distinct'", "'else'", "'end'", "'false'", "'from'", "'full'", "'group'",
		"'having'", "'in'", "'inner'", "'is'", "'join'", "'left'", "'like'",
		"'natural'", "'not'", "'null'", "'on'", "'or'", "'order'", "'outer'",
		"'over'", "'right'", "'select'", "'then'", "'true'", "'union'", "'when'",
		"'where'", "'with'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "ALL", "AND", "AS", "ASC", "BY", "CASE", "CROSS", "DESC",
		"DISTINCT", "ELSE", "END", "FALSE", "FROM", "FULL", "GROUP", "HAVING",
		"IN", "INNER", "IS", "JOIN", "LEFT", "LIKE", "NATURAL", "NOT", "NULL",
		"ON", "OR", "ORDER", "OUTER", "OVER", "RIGHT", "SELECT", "THEN", "TRUE",
		"UNION", "WHEN", "WHERE", "WITH", "STRING", "INTEGER_VALUE", "DECIMAL_VALUE",
		"FLOAT_VALUE", "IDENTIFIER", "QUOTED_IDENTIFIER", "COMMENT", "BLOCK_COMMENT",
		"WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "ALL", "AND", "AS", "ASC", "BY", "CASE", "CROSS",
		"DESC", "DISTINCT", "ELSE", "END", "FALSE", "FROM", "FULL", "GROUP",
		"HAVING", "IN", "INNER", "IS", "JOIN", "LEFT", "LIKE", "NATURAL", "NOT",
		"NULL", "ON", "OR", "ORDER", "OUTER", "OVER", "RIGHT", "SELECT", "THEN",
		"TRUE", "UNION", "WHEN", "WHERE", "WITH", "STRING", "INTEGER_VALUE",
		"DECIMAL_VALUE", "FLOAT_VALUE", "IDENTIFIER", "QUOTED_IDENTIFIER", "EXPONENT",
		"DIGIT", "LETTER", "COMMENT", "BLOCK_COMMENT", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 66, 512, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1,
		18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21,
		1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29,
		1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1,
		35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37,
		1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1,
		40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41,
		1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1,
		43, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46,
		1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1,
		48, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50,
		1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1,
		52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53,
		1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1,
		55, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57, 1, 57, 5, 57,
		381, 8, 57, 10, 57, 12, 57, 384, 9, 57, 1, 57, 1, 57, 1, 58, 4, 58, 389,
		8, 58, 11, 58, 12, 58, 390, 1, 59, 4, 59, 394, 8, 59, 11, 59, 12, 59, 395,
		1, 59, 1, 59, 5, 59, 400, 8, 59, 10, 59, 12, 59, 403, 9, 59, 1, 59, 1,
		59, 4, 59, 407, 8, 59, 11, 59, 12, 59, 408, 3, 59, 411, 8, 59, 1, 60, 4,
		60, 414, 8, 60, 11, 60, 12, 60, 415, 1, 60, 1, 60, 5, 60, 420, 8, 60, 10,
		60, 12, 60, 423, 9, 60, 3, 60, 425, 8, 60, 1, 60, 1, 60, 1, 60, 1, 60,
		4, 60, 431, 8, 60, 11, 60, 12, 60, 432, 1, 60, 1, 60, 3, 60, 437, 8, 60,
		1, 61, 1, 61, 3, 61, 441, 8, 61, 1, 61, 1, 61, 1, 61, 5, 61, 446, 8, 61,
		10, 61, 12, 61, 449, 9, 61, 1, 62, 1, 62, 1, 62, 1, 62, 5, 62, 455, 8,
		62, 10, 62, 12, 62, 458, 9, 62, 1, 62, 1, 62, 1, 63, 1, 63, 3, 63, 464,
		8, 63, 1, 63, 4, 63, 467, 8, 63, 11, 63, 12, 63, 468, 1, 64, 1, 64, 1,
		65, 1, 65, 1, 66, 1, 66, 1, 66, 1, 66, 5, 66, 479, 8, 66, 10, 66, 12, 66,
		482, 9, 66, 1, 66, 3, 66, 485, 8, 66, 1, 66, 3, 66, 488, 8, 66, 1, 66,
		1, 66, 1, 67, 1, 67, 1, 67, 1, 67, 5, 67, 496, 8, 67, 10, 67, 12, 67, 499,
		9, 67, 1, 67, 1, 67, 1, 67, 1, 67, 1, 67, 1, 68, 4, 68, 507, 8, 68, 11,
		68, 12, 68, 508, 1, 68, 1, 68, 1, 497, 0, 69, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32,
		65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41,
		83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48, 97, 49, 99, 50,
		101, 51, 103, 52, 105, 53, 107, 54, 109, 55, 111, 56, 113, 57, 115, 58,
		117, 59, 119, 60, 121, 61, 123, 62, 125, 63, 127, 0, 129, 0, 131, 0, 133,
		64, 135, 65, 137, 66, 1, 0, 9, 1, 0, 39, 39, 4, 0, 36, 36, 58, 58, 64,
		64, 95, 95, 1, 0, 34, 34, 2, 0, 69, 69, 101, 101, 2, 0, 43, 43, 45, 45,
		1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10,
		13, 13, 32, 32, 533, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0,
		0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0,
		0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1,
		0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29,
		1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0,
		37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0,
		0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0,
		0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0,
		0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1,
		0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75,
		1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0,
		83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0,
		0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0,
		0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1,
		0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0,
		113, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1, 0,
		0, 0, 0, 121, 1, 0, 0, 0, 0, 123, 1, 0, 0, 0, 0, 125, 1, 0, 0, 0, 0, 133,
		1, 0, 0, 0, 0, 135, 1, 0, 0, 0, 0, 137, 1, 0, 0, 0, 1, 139, 1, 0, 0, 0,
		3, 141, 1, 0, 0, 0, 5, 143, 1, 0, 0, 0, 7, 145, 1, 0, 0, 0, 9, 147, 1,
		0, 0, 0, 11, 149, 1, 0, 0, 0, 13, 152, 1, 0, 0, 0, 15, 154, 1, 0, 0, 0,
		17, 156, 1, 0, 0, 0, 19, 159, 1, 0, 0, 0, 21, 162, 1, 0, 0, 0, 23, 164,
		1, 0, 0, 0, 25, 167, 1, 0, 0, 0, 27, 169, 1, 0, 0, 0, 29, 172, 1, 0, 0,
		0, 31, 174, 1, 0, 0, 0, 33, 176, 1, 0, 0, 0, 35, 178, 1, 0, 0, 0, 37, 180,
		1, 0, 0, 0, 39, 183, 1, 0, 0, 0, 41, 187, 1, 0, 0, 0, 43, 191, 1, 0, 0,
		0, 45, 194, 1, 0, 0, 0, 47, 198, 1, 0, 0, 0, 49, 201, 1, 0, 0, 0, 51, 206,
		1, 0, 0, 0, 53, 212, 1, 0, 0, 0, 55, 217, 1, 0, 0, 0, 57, 226, 1, 0, 0,
		0, 59, 231, 1, 0, 0, 0, 61, 235, 1, 0, 0, 0, 63, 241, 1, 0, 0, 0, 65, 246,
		1, 0, 0, 0, 67, 251, 1, 0, 0, 0, 69, 257, 1, 0, 0, 0, 71, 264, 1, 0, 0,
		0, 73, 267, 1, 0, 0, 0, 75, 273, 1, 0, 0, 0, 77, 276, 1, 0, 0, 0, 79, 281,
		1, 0, 0, 0, 81, 286, 1, 0, 0, 0, 83, 291, 1, 0, 0, 0, 85, 299, 1, 0, 0,
		0, 87, 303, 1, 0, 0, 0, 89, 308, 1, 0, 0, 0, 91, 311, 1, 0, 0, 0, 93, 314,
		1, 0, 0, 0, 95, 320, 1, 0, 0, 0, 97, 326, 1, 0, 0, 0, 99, 331, 1, 0, 0,
		0, 101, 337, 1, 0, 0, 0, 103, 344, 1, 0, 0, 0, 105, 349, 1, 0, 0, 0, 107,
		354, 1, 0, 0, 0, 109, 360, 1, 0, 0, 0, 111, 365, 1, 0, 0, 0, 113, 371,
		1, 0, 0, 0, 115, 376, 1, 0, 0, 0, 117, 388, 1, 0, 0, 0, 119, 410, 1, 0,
		0, 0, 121, 436, 1, 0, 0, 0, 123, 440, 1, 0, 0, 0, 125, 450, 1, 0, 0, 0,
		127, 461, 1, 0, 0, 0, 129, 470, 1, 0, 0, 0, 131, 472, 1, 0, 0, 0, 133,
		474, 1, 0, 0, 0, 135, 491, 1, 0, 0, 0, 137, 506, 1, 0, 0, 0, 139, 140,
		5, 59, 0, 0, 140, 2, 1, 0, 0, 0, 141, 142, 5, 44, 0, 0, 142, 4, 1, 0, 0,
		0, 143, 144, 5, 40, 0, 0, 144, 6, 1, 0, 0, 0, 145, 146, 5, 41, 0, 0, 146,
		8, 1, 0, 0, 0, 147, 148, 5, 42, 0, 0, 148, 10, 1, 0, 0, 0, 149, 150, 5,
		58, 0, 0, 150, 151, 5, 58, 0, 0, 151, 12, 1, 0, 0, 0, 152, 153, 5, 46,
		0, 0, 153, 14, 1, 0, 0, 0, 154, 155, 5, 61, 0, 0, 155, 16, 1, 0, 0, 0,
		156, 157, 5, 33, 0, 0, 157, 158, 5, 61, 0, 0, 158, 18, 1, 0, 0, 0, 159,
		160, 5, 60, 0, 0, 160, 161, 5, 62, 0, 0, 161, 20, 1, 0, 0, 0, 162, 163,
		5, 60, 0, 0, 163, 22, 1, 0, 0, 0, 164, 165, 5, 60, 0, 0, 165, 166, 5, 61,
		0, 0, 166, 24, 1, 0, 0, 0, 167, 168, 5, 62, 0, 0, 168, 26, 1, 0, 0, 0,
		169, 170, 5, 62, 0, 0, 170, 171, 5, 61, 0, 0, 171, 28, 1, 0, 0, 0, 172,
		173, 5, 43, 0, 0, 173, 30, 1, 0, 0, 0, 174, 175, 5, 45, 0, 0, 175, 32,
		1, 0, 0, 0, 176, 177, 5, 47, 0, 0, 177, 34, 1, 0, 0, 0, 178, 179, 5, 37,
		0, 0, 179, 36, 1, 0, 0, 0, 180, 181, 5, 124, 0, 0, 181, 182, 5, 124, 0,
		0, 182, 38, 1, 0, 0, 0, 183, 184, 5, 97, 0, 0, 184, 185, 5, 108, 0, 0,
		185, 186, 5, 108, 0, 0, 186, 40, 1, 0, 0, 0, 187, 188, 5, 97, 0, 0, 188,
		189, 5, 110, 0, 0, 189, 190, 5, 100, 0, 0, 190, 42, 1, 0, 0, 0, 191, 192,
		5, 97, 0, 0, 192, 193, 5, 115, 0, 0, 193, 44, 1, 0, 0, 0, 194, 195, 5,
		97, 0, 0, 195, 196, 5, 115, 0, 0, 196, 197, 5, 99, 0, 0, 197, 46, 1, 0,
		0, 0, 198, 199, 5, 98, 0, 0, 199, 200, 5, 121, 0, 0, 200, 48, 1, 0, 0,
		0, 201, 202, 5, 99, 0, 0, 202, 203, 5, 97, 0, 0, 203, 204, 5, 115, 0, 0,
		204, 205, 5, 101, 0, 0, 205, 50, 1, 0, 0, 0, 206, 207, 5, 99, 0, 0, 207,
		208, 5, 114, 0, 0, 208, 209, 5, 111, 0, 0, 209, 210, 5, 115, 0, 0, 210,
		211, 5, 115, 0, 0, 211, 52, 1, 0, 0, 0, 212, 213, 5, 100, 0, 0, 213, 214,
		5, 101, 0, 0, 214, 215, 5, 115, 0, 0, 215, 216, 5, 99, 0, 0, 216, 54, 1,
		0, 0, 0, 217, 218, 5, 100, 0, 0, 218, 219, 5, 105, 0, 0, 219, 220, 5, 115,
		0, 0, 220, 221, 5, 116, 0, 0, 221, 222, 5, 105, 0, 0, 222, 223, 5, 110,
		0, 0, 223, 224, 5, 99, 0, 0, 224, 225, 5, 116, 0, 0, 225, 56, 1, 0, 0,
		0, 226, 227, 5, 101, 0, 0, 227, 228, 5, 108, 0, 0, 228, 229, 5, 115, 0,
		0, 229, 230, 5, 101, 0, 0, 230, 58, 1, 0, 0, 0, 231, 232, 5, 101, 0, 0,
		232, 233, 5, 110, 0, 0, 233, 234, 5, 100, 0, 0, 234, 60, 1, 0, 0, 0, 235,
		236, 5, 102, 0, 0, 236, 237, 5, 97, 0, 0, 237, 238, 5, 108, 0, 0, 238,
		239, 5, 115, 0, 0, 239, 240, 5, 101, 0, 0, 240, 62, 1, 0, 0, 0, 241, 242,
		5, 102, 0, 0, 242, 243, 5, 114, 0, 0, 243, 244, 5, 111, 0, 0, 244, 245,
		5, 109, 0, 0, 245, 64, 1, 0, 0, 0, 246, 247, 5, 102, 0, 0, 247, 248, 5,
		117, 0, 0, 248, 249, 5, 108, 0, 0, 249, 250, 5, 108, 0, 0, 250, 66, 1,
		0, 0, 0, 251, 252, 5, 103, 0, 0, 252, 253, 5, 114, 0, 0, 253, 254, 5, 111,
		0, 0, 254, 255, 5, 117, 0, 0, 255, 256, 5, 112, 0, 0, 256, 68, 1, 0, 0,
		0, 257, 258, 5, 104, 0, 0, 258, 259, 5, 97, 0, 0, 259, 260, 5, 118, 0,
		0, 260, 261, 5, 105, 0, 0, 261, 262, 5, 110, 0, 0, 262, 263, 5, 103, 0,
		0, 263, 70, 1, 0, 0, 0, 264, 265, 5, 105, 0, 0, 265, 266, 5, 110, 0, 0,
		266, 72, 1, 0, 0, 0, 267, 268, 5, 105, 0, 0, 268, 269, 5, 110, 0, 0, 269,
		270, 5, 110, 0, 0, 270, 271, 5, 101, 0, 0, 271, 272, 5, 114, 0, 0, 272,
		74, 1, 0, 0, 0, 273, 274, 5, 105, 0, 0, 274, 275, 5, 115, 0, 0, 275, 76,
		1, 0, 0, 0, 276, 277, 5, 106, 0, 0, 277, 278, 5, 111, 0, 0, 278, 279, 5,
		105, 0, 0, 279, 280, 5, 110, 0, 0, 280, 78, 1, 0, 0, 0, 281, 282, 5, 108,
		0, 0, 282, 283, 5, 101, 0, 0, 283, 284, 5, 102, 0, 0, 284, 285, 5, 116,
		0, 0, 285, 80, 1, 0, 0, 0, 286, 287, 5, 108, 0, 0, 287, 288, 5, 105, 0,
		0, 288, 289, 5, 107, 0, 0, 289, 290, 5, 101, 0, 0, 290, 82, 1, 0, 0, 0,
		291, 292, 5, 110, 0, 0, 292, 293, 5, 97, 0, 0, 293, 294, 5, 116, 0, 0,
		294, 295, 5, 117, 0, 0, 295, 296, 5, 114, 0, 0, 296, 297, 5, 97, 0, 0,
		297, 298, 5, 108, 0, 0, 298, 84, 1, 0, 0, 0, 299, 300, 5, 110, 0, 0, 300,
		301, 5, 111, 0, 0, 301, 302, 5, 116, 0, 0, 302, 86, 1, 0, 0, 0, 303, 304,
		5, 110, 0, 0, 304, 305, 5, 117, 0, 0, 305, 306, 5, 108, 0, 0, 306, 307,
		5, 108, 0, 0, 307, 88, 1, 0, 0, 0, 308, 309, 5, 111, 0, 0, 309, 310, 5,
		110, 0, 0, 310, 90, 1, 0, 0, 0, 311, 312, 5, 111, 0, 0, 312, 313, 5, 114,
		0, 0, 313, 92, 1, 0, 0, 0, 314, 315, 5, 111, 0, 0, 315, 316, 5, 114, 0,
		0, 316, 317, 5, 100, 0, 0, 317, 318, 5, 101, 0, 0, 318, 319, 5, 114, 0,
		0, 319, 94, 1, 0, 0, 0, 320, 321, 5, 111, 0, 0, 321, 322, 5, 117, 0, 0,
		322, 323, 5, 116, 0, 0, 323, 324, 5, 101, 0, 0, 324, 325, 5, 114, 0, 0,
		325, 96, 1, 0, 0, 0, 326, 327, 5, 111, 0, 0, 327, 328, 5, 118, 0, 0, 328,
		329, 5, 101, 0, 0, 329, 330, 5, 114, 0, 0, 330, 98, 1, 0, 0, 0, 331, 332,
		5, 114, 0, 0, 332, 333, 5, 105, 0, 0, 333, 334, 5, 103, 0, 0, 334, 335,
		5, 104, 0, 0, 335, 336, 5, 116, 0, 0, 336, 100, 1, 0, 0, 0, 337, 338, 5,
		115, 0, 0, 338, 339, 5, 101, 0, 0, 339, 340, 5, 108, 0, 0, 340, 341, 5,
		101, 0, 0, 341, 342, 5, 99, 0, 0, 342, 343, 5, 116, 0, 0, 343, 102, 1,
		0, 0, 0, 344, 345, 5, 116, 0, 0, 345, 346, 5, 104, 0, 0, 346, 347, 5, 101,
		0, 0, 347, 348, 5, 110, 0, 0, 348, 104, 1, 0, 0, 0, 349, 350, 5, 116, 0,
		0, 350, 351, 5, 114, 0, 0, 351, 352, 5, 117, 0, 0, 352, 353, 5, 101, 0,
		0, 353, 106, 1, 0, 0, 0, 354, 355, 5, 117, 0, 0, 355, 356, 5, 110, 0, 0,
		356, 357, 5, 105, 0, 0, 357, 358, 5, 111, 0, 0, 358, 359, 5, 110, 0, 0,
		359, 108, 1, 0, 0, 0, 360, 361, 5, 119, 0, 0, 361, 362, 5, 104, 0, 0, 362,
		363, 5, 101, 0, 0, 363, 364, 5, 110, 0, 0, 364, 110, 1, 0, 0, 0, 365, 366,
		5, 119, 0, 0, 366, 367, 5, 104, 0, 0, 367, 368, 5, 101, 0, 0, 368, 369,
		5, 114, 0, 0, 369, 370, 5, 101, 0, 0, 370, 112, 1, 0, 0, 0, 371, 372, 5,
		119, 0, 0, 372, 373, 5, 105, 0, 0, 373, 374, 5, 116, 0, 0, 374, 375, 5,
		104, 0, 0, 375, 114, 1, 0, 0, 0, 376, 382, 5, 39, 0, 0, 377, 381, 8, 0,
		0, 0, 378, 379, 5, 39, 0, 0, 379, 381, 5, 39, 0, 0, 380, 377, 1, 0, 0,
		0, 380, 378, 1, 0, 0, 0, 381, 384, 1, 0, 0, 0, 382, 380, 1, 0, 0, 0, 382,
		383, 1, 0, 0, 0, 383, 385, 1, 0, 0, 0, 384, 382, 1, 0, 0, 0, 385, 386,
		5, 39, 0, 0, 386, 116, 1, 0, 0, 0, 387, 389, 3, 129, 64, 0, 388, 387, 1,
		0, 0, 0, 389, 390, 1, 0, 0, 0, 390, 388, 1, 0, 0, 0, 390, 391, 1, 0, 0,
		0, 391, 118, 1, 0, 0, 0, 392, 394, 3, 129, 64, 0, 393, 392, 1, 0, 0, 0,
		394, 395, 1, 0, 0, 0, 395, 393, 1, 0, 0, 0, 395, 396, 1, 0, 0, 0, 396,
		397, 1, 0, 0, 0, 397, 401, 5, 46, 0, 0, 398, 400, 3, 129, 64, 0, 399, 398,
		1, 0, 0, 0, 400, 403, 1, 0, 0, 0, 401, 399, 1, 0, 0, 0, 401, 402, 1, 0,
		0, 0, 402, 411, 1, 0, 0, 0, 403, 401, 1, 0, 0, 0, 404, 406, 5, 46, 0, 0,
		405, 407, 3, 129, 64, 0, 406, 405, 1, 0, 0, 0, 407, 408, 1, 0, 0, 0, 408,
		406, 1, 0, 0, 0, 408, 409, 1, 0, 0, 0, 409, 411, 1, 0, 0, 0, 410, 393,
		1, 0, 0, 0, 410, 404, 1, 0, 0, 0, 411, 120, 1, 0, 0, 0, 412, 414, 3, 129,
		64, 0, 413, 412, 1, 0, 0, 0, 414, 415, 1, 0, 0, 0, 415, 413, 1, 0, 0, 0,
		415, 416, 1, 0, 0, 0, 416, 424, 1, 0, 0, 0, 417, 421, 5, 46, 0, 0, 418,
		420, 3, 129, 64, 0, 419, 418, 1, 0, 0, 0, 420, 423, 1, 0, 0, 0, 421, 419,
		1, 0, 0, 0, 421, 422, 1, 0, 0, 0, 422, 425, 1, 0, 0, 0, 423, 421, 1, 0,
		0, 0, 424, 417, 1, 0, 0, 0, 424, 425, 1, 0, 0, 0, 425, 426, 1, 0, 0, 0,
		426, 427, 3, 127, 63, 0, 427, 437, 1, 0, 0, 0, 428, 430, 5, 46, 0, 0, 429,
		431, 3, 129, 64, 0, 430, 429, 1, 0, 0, 0, 431, 432, 1, 0, 0, 0, 432, 430,
		1, 0, 0, 0, 432, 433, 1, 0, 0, 0, 433, 434, 1, 0, 0, 0, 434, 435, 3, 127,
		63, 0, 435, 437, 1, 0, 0, 0, 436, 413, 1, 0, 0, 0, 436, 428, 1, 0, 0, 0,
		437, 122, 1, 0, 0, 0, 438, 441, 3, 131, 65, 0, 439, 441, 5, 95, 0, 0, 440,
		438, 1, 0, 0, 0, 440, 439, 1, 0, 0, 0, 441, 447, 1, 0, 0, 0, 442, 446,
		3, 131, 65, 0, 443, 446, 3, 129, 64, 0, 444, 446, 7, 1, 0, 0, 445, 442,
		1, 0, 0, 0, 445, 443, 1, 0, 0, 0, 445, 444, 1, 0, 0, 0, 446, 449, 1, 0,
		0, 0, 447, 445, 1, 0, 0, 0, 447, 448, 1, 0, 0, 0, 448, 124, 1, 0, 0, 0,
		449, 447, 1, 0, 0, 0, 450, 456, 5, 34, 0, 0, 451, 455, 8, 2, 0, 0, 452,
		453, 5, 34, 0, 0, 453, 455, 5, 34, 0, 0, 454, 451, 1, 0, 0, 0, 454, 452,
		1, 0, 0, 0, 455, 458, 1, 0, 0, 0, 456, 454, 1, 0, 0, 0, 456, 457, 1, 0,
		0, 0, 457, 459, 1, 0, 0, 0, 458, 456, 1, 0, 0, 0, 459, 460, 5, 34, 0, 0,
		460, 126, 1, 0, 0, 0, 461, 463, 7, 3, 0, 0, 462, 464, 7, 4, 0, 0, 463,
		462, 1, 0, 0, 0, 463, 464, 1, 0, 0, 0, 464, 466, 1, 0, 0, 0, 465, 467,
		3, 129, 64, 0, 466, 465, 1, 0, 0, 0, 467, 468, 1, 0, 0, 0, 468, 466, 1,
		0, 0, 0, 468, 469, 1, 0, 0, 0, 469, 128, 1, 0, 0, 0, 470, 471, 7, 5, 0,
		0, 471, 130, 1, 0, 0, 0, 472, 473, 7, 6, 0, 0, 473, 132, 1, 0, 0, 0, 474,
		475, 5, 45, 0, 0, 475, 476, 5, 45, 0, 0, 476, 480, 1, 0, 0, 0, 477, 479,
		8, 7, 0, 0, 478, 477, 1, 0, 0, 0, 479, 482, 1, 0, 0, 0, 480, 478, 1, 0,
		0, 0, 480, 481, 1, 0, 0, 0, 481, 484, 1, 0, 0, 0, 482, 480, 1, 0, 0, 0,
		483, 485, 5, 13, 0, 0, 484, 483, 1, 0, 0, 0, 484, 485, 1, 0, 0, 0, 485,
		487, 1, 0, 0, 0, 486, 488, 5, 10, 0, 0, 487, 486, 1, 0, 0, 0, 487, 488,
		1, 0, 0, 0, 488, 489, 1, 0, 0, 0, 489, 490, 6, 66, 0, 0, 490, 134, 1, 0,
		0, 0, 491, 492, 5, 47, 0, 0, 492, 493, 5, 42, 0, 0, 493, 497, 1, 0, 0,
		0, 494, 496, 9, 0, 0, 0, 495, 494, 1, 0, 0, 0, 496, 499, 1, 0, 0, 0, 497,
		498, 1, 0, 0, 0, 497, 495, 1, 0, 0, 0, 498, 500, 1, 0, 0, 0, 499, 497,
		1, 0, 0, 0, 500, 501, 5, 42, 0, 0, 501, 502, 5, 47, 0, 0, 502, 503, 1,
		0, 0, 0, 503, 504, 6, 67, 0, 0, 504, 136, 1, 0, 0, 0, 505, 507, 7, 8, 0,
		0, 506, 505, 1, 0, 0, 0, 507, 508, 1, 0, 0, 0, 508, 506, 1, 0, 0, 0, 508,
		509, 1, 0, 0, 0, 509, 510, 1, 0, 0, 0, 510, 511, 6, 68, 1, 0, 511, 138,
		1, 0, 0, 0, 25, 0, 380, 382, 390, 395, 401, 408, 410, 415, 421, 424, 432,
		436, 440, 445, 447, 454, 456, 463, 468, 480, 484, 487, 497, 508, 2, 0,
		1, 0, 6, 0, 0,
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

// SimpleSqlLexerInit initializes any static state used to implement SimpleSqlLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSimpleSqlLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SimpleSqlLexerInit() {
	staticData := &simplesqllexerLexerStaticData
	staticData.once.Do(simplesqllexerLexerInit)
}

// NewSimpleSqlLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSimpleSqlLexer(input antlr.CharStream) *SimpleSqlLexer {
	SimpleSqlLexerInit()
	l := new(SimpleSqlLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &simplesqllexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "SimpleSql.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SimpleSqlLexer tokens.
const (
	SimpleSqlLexerT__0              = 1
	SimpleSqlLexerT__1              = 2
	SimpleSqlLexerT__2              = 3
	SimpleSqlLexerT__3              = 4
	SimpleSqlLexerT__4              = 5
	SimpleSqlLexerT__5              = 6
	SimpleSqlLexerT__6              = 7
	SimpleSqlLexerT__7              = 8
	SimpleSqlLexerT__8              = 9
	SimpleSqlLexerT__9              = 10
	SimpleSqlLexerT__10             = 11
	SimpleSqlLexerT__11             = 12
	SimpleSqlLexerT__12             = 13
	SimpleSqlLexerT__13             = 14
	SimpleSqlLexerT__14             = 15
	SimpleSqlLexerT__15             = 16
	SimpleSqlLexerT__16             = 17
	SimpleSqlLexerT__17             = 18
	SimpleSqlLexerT__18             = 19
	SimpleSqlLexerALL               = 20
	SimpleSqlLexerAND               = 21
	SimpleSqlLexerAS                = 22
	SimpleSqlLexerASC               = 23
	SimpleSqlLexerBY                = 24
	SimpleSqlLexerCASE              = 25
	SimpleSqlLexerCROSS             = 26
	SimpleSqlLexerDESC              = 27
	SimpleSqlLexerDISTINCT          = 28
	SimpleSqlLexerELSE              = 29
	SimpleSqlLexerEND               = 30
	SimpleSqlLexerFALSE             = 31
	SimpleSqlLexerFROM              = 32
	SimpleSqlLexerFULL              = 33
	SimpleSqlLexerGROUP             = 34
	SimpleSqlLexerHAVING            = 35
	SimpleSqlLexerIN                = 36
	SimpleSqlLexerINNER             = 37
	SimpleSqlLexerIS                = 38
	SimpleSqlLexerJOIN              = 39
	SimpleSqlLexerLEFT              = 40
	SimpleSqlLexerLIKE              = 41
	SimpleSqlLexerNATURAL           = 42
	SimpleSqlLexerNOT               = 43
	SimpleSqlLexerNULL              = 44
	SimpleSqlLexerON                = 45
	SimpleSqlLexerOR                = 46
	SimpleSqlLexerORDER             = 47
	SimpleSqlLexerOUTER             = 48
	SimpleSqlLexerOVER              = 49
	SimpleSqlLexerRIGHT             = 50
	SimpleSqlLexerSELECT            = 51
	SimpleSqlLexerTHEN              = 52
	SimpleSqlLexerTRUE              = 53
	SimpleSqlLexerUNION             = 54
	SimpleSqlLexerWHEN              = 55
	SimpleSqlLexerWHERE             = 56
	SimpleSqlLexerWITH              = 57
	SimpleSqlLexerSTRING            = 58
	SimpleSqlLexerINTEGER_VALUE     = 59
	SimpleSqlLexerDECIMAL_VALUE     = 60
	SimpleSqlLexerFLOAT_VALUE       = 61
	SimpleSqlLexerIDENTIFIER        = 62
	SimpleSqlLexerQUOTED_IDENTIFIER = 63
	SimpleSqlLexerCOMMENT           = 64
	SimpleSqlLexerBLOCK_COMMENT     = 65
	SimpleSqlLexerWS                = 66
)
