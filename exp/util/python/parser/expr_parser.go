// Code generated from Expr.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Expr

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

type ExprParser struct {
	*antlr.BaseParser
}

var exprParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func exprParserInit() {
	staticData := &exprParserStaticData
	staticData.literalNames = []string{
		"", "'<'", "'>'", "'=='", "'>='", "'<='", "'<>'", "'!='", "'|'", "'^'",
		"'&'", "'<<'", "'>>'", "'+'", "'-'", "'*'", "'@'", "'/'", "'%'", "'//'",
		"'~'", "'**'", "'('", "')'", "'['", "']'", "'{'", "'}'", "'...'", "','",
		"'.'", "':'", "", "", "", "'and'", "'False'", "'in'", "'is'", "'None'",
		"'not'", "'or'", "'True'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "STRING",
		"NUMBER", "INTEGER", "AND", "FALSE", "IN", "IS", "NONE", "NOT", "OR",
		"TRUE", "NAME", "STRING_LITERAL", "BYTES_LITERAL", "DECIMAL_INTEGER",
		"OCT_INTEGER", "HEX_INTEGER", "BIN_INTEGER", "FLOAT_NUMBER", "IMAG_NUMBER",
		"SKIP_", "UNKNOWN_CHAR", "INDENT", "DEDENT",
	}
	staticData.ruleNames = []string{
		"singleExpr", "expr", "orTest", "andTest", "notTest", "comparison",
		"compOp", "exprMain", "exprCont", "xorExpr", "xorExprCont", "andExpr",
		"andExprCont", "shiftExpr", "shiftExprCont", "arithExpr", "arithExprCont",
		"term", "termCont", "factor", "power", "atomExpr", "atom", "constVal",
		"testListComp", "exprOrStarExpr", "starExpr", "trailer", "subscriptList",
		"subscript", "sliceOp", "dictOrSetMaker", "dictItem",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 55, 311, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 5, 2,
		75, 8, 2, 10, 2, 12, 2, 78, 9, 2, 1, 3, 1, 3, 1, 3, 5, 3, 83, 8, 3, 10,
		3, 12, 3, 86, 9, 3, 1, 4, 1, 4, 1, 4, 3, 4, 91, 8, 4, 1, 5, 1, 5, 1, 5,
		1, 5, 5, 5, 97, 8, 5, 10, 5, 12, 5, 100, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 115, 8, 6,
		1, 7, 1, 7, 5, 7, 119, 8, 7, 10, 7, 12, 7, 122, 9, 7, 1, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 5, 9, 129, 8, 9, 10, 9, 12, 9, 132, 9, 9, 1, 10, 1, 10, 1,
		10, 1, 11, 1, 11, 5, 11, 139, 8, 11, 10, 11, 12, 11, 142, 9, 11, 1, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 5, 13, 149, 8, 13, 10, 13, 12, 13, 152, 9,
		13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 5, 15, 159, 8, 15, 10, 15, 12, 15,
		162, 9, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 5, 17, 169, 8, 17, 10, 17,
		12, 17, 172, 9, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 3, 19, 180,
		8, 19, 1, 20, 1, 20, 1, 20, 3, 20, 185, 8, 20, 1, 21, 1, 21, 5, 21, 189,
		8, 21, 10, 21, 12, 21, 192, 9, 21, 1, 22, 1, 22, 3, 22, 196, 8, 22, 1,
		22, 1, 22, 1, 22, 3, 22, 201, 8, 22, 1, 22, 1, 22, 1, 22, 3, 22, 206, 8,
		22, 1, 22, 1, 22, 3, 22, 210, 8, 22, 1, 23, 1, 23, 1, 23, 4, 23, 215, 8,
		23, 11, 23, 12, 23, 216, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 223, 8, 23,
		1, 24, 1, 24, 1, 24, 5, 24, 228, 8, 24, 10, 24, 12, 24, 231, 9, 24, 1,
		24, 3, 24, 234, 8, 24, 1, 25, 1, 25, 3, 25, 238, 8, 25, 1, 26, 1, 26, 1,
		26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 249, 8, 27, 1, 28,
		1, 28, 1, 28, 5, 28, 254, 8, 28, 10, 28, 12, 28, 257, 9, 28, 1, 28, 3,
		28, 260, 8, 28, 1, 29, 1, 29, 3, 29, 264, 8, 29, 1, 29, 1, 29, 3, 29, 268,
		8, 29, 1, 29, 3, 29, 271, 8, 29, 3, 29, 273, 8, 29, 1, 30, 1, 30, 3, 30,
		277, 8, 30, 1, 31, 1, 31, 1, 31, 5, 31, 282, 8, 31, 10, 31, 12, 31, 285,
		9, 31, 1, 31, 3, 31, 288, 8, 31, 1, 31, 1, 31, 1, 31, 5, 31, 293, 8, 31,
		10, 31, 12, 31, 296, 9, 31, 1, 31, 3, 31, 299, 8, 31, 3, 31, 301, 8, 31,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 309, 8, 32, 1, 32, 0,
		0, 33, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 0, 4, 1, 0,
		11, 12, 1, 0, 13, 14, 1, 0, 15, 19, 2, 0, 13, 14, 20, 20, 330, 0, 66, 1,
		0, 0, 0, 2, 69, 1, 0, 0, 0, 4, 71, 1, 0, 0, 0, 6, 79, 1, 0, 0, 0, 8, 90,
		1, 0, 0, 0, 10, 92, 1, 0, 0, 0, 12, 114, 1, 0, 0, 0, 14, 116, 1, 0, 0,
		0, 16, 123, 1, 0, 0, 0, 18, 126, 1, 0, 0, 0, 20, 133, 1, 0, 0, 0, 22, 136,
		1, 0, 0, 0, 24, 143, 1, 0, 0, 0, 26, 146, 1, 0, 0, 0, 28, 153, 1, 0, 0,
		0, 30, 156, 1, 0, 0, 0, 32, 163, 1, 0, 0, 0, 34, 166, 1, 0, 0, 0, 36, 173,
		1, 0, 0, 0, 38, 179, 1, 0, 0, 0, 40, 181, 1, 0, 0, 0, 42, 186, 1, 0, 0,
		0, 44, 209, 1, 0, 0, 0, 46, 222, 1, 0, 0, 0, 48, 224, 1, 0, 0, 0, 50, 237,
		1, 0, 0, 0, 52, 239, 1, 0, 0, 0, 54, 248, 1, 0, 0, 0, 56, 250, 1, 0, 0,
		0, 58, 272, 1, 0, 0, 0, 60, 274, 1, 0, 0, 0, 62, 300, 1, 0, 0, 0, 64, 308,
		1, 0, 0, 0, 66, 67, 3, 2, 1, 0, 67, 68, 5, 0, 0, 1, 68, 1, 1, 0, 0, 0,
		69, 70, 3, 4, 2, 0, 70, 3, 1, 0, 0, 0, 71, 76, 3, 6, 3, 0, 72, 73, 5, 41,
		0, 0, 73, 75, 3, 6, 3, 0, 74, 72, 1, 0, 0, 0, 75, 78, 1, 0, 0, 0, 76, 74,
		1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 5, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0,
		79, 84, 3, 8, 4, 0, 80, 81, 5, 35, 0, 0, 81, 83, 3, 8, 4, 0, 82, 80, 1,
		0, 0, 0, 83, 86, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85,
		7, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 87, 88, 5, 40, 0, 0, 88, 91, 3, 8, 4,
		0, 89, 91, 3, 10, 5, 0, 90, 87, 1, 0, 0, 0, 90, 89, 1, 0, 0, 0, 91, 9,
		1, 0, 0, 0, 92, 98, 3, 14, 7, 0, 93, 94, 3, 12, 6, 0, 94, 95, 3, 14, 7,
		0, 95, 97, 1, 0, 0, 0, 96, 93, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 96,
		1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 11, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0,
		101, 115, 5, 1, 0, 0, 102, 115, 5, 2, 0, 0, 103, 115, 5, 3, 0, 0, 104,
		115, 5, 4, 0, 0, 105, 115, 5, 5, 0, 0, 106, 115, 5, 6, 0, 0, 107, 115,
		5, 7, 0, 0, 108, 115, 5, 37, 0, 0, 109, 110, 5, 40, 0, 0, 110, 115, 5,
		37, 0, 0, 111, 115, 5, 38, 0, 0, 112, 113, 5, 38, 0, 0, 113, 115, 5, 40,
		0, 0, 114, 101, 1, 0, 0, 0, 114, 102, 1, 0, 0, 0, 114, 103, 1, 0, 0, 0,
		114, 104, 1, 0, 0, 0, 114, 105, 1, 0, 0, 0, 114, 106, 1, 0, 0, 0, 114,
		107, 1, 0, 0, 0, 114, 108, 1, 0, 0, 0, 114, 109, 1, 0, 0, 0, 114, 111,
		1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 115, 13, 1, 0, 0, 0, 116, 120, 3, 18,
		9, 0, 117, 119, 3, 16, 8, 0, 118, 117, 1, 0, 0, 0, 119, 122, 1, 0, 0, 0,
		120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 15, 1, 0, 0, 0, 122, 120,
		1, 0, 0, 0, 123, 124, 5, 8, 0, 0, 124, 125, 3, 18, 9, 0, 125, 17, 1, 0,
		0, 0, 126, 130, 3, 22, 11, 0, 127, 129, 3, 20, 10, 0, 128, 127, 1, 0, 0,
		0, 129, 132, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131,
		19, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 133, 134, 5, 9, 0, 0, 134, 135, 3,
		22, 11, 0, 135, 21, 1, 0, 0, 0, 136, 140, 3, 26, 13, 0, 137, 139, 3, 24,
		12, 0, 138, 137, 1, 0, 0, 0, 139, 142, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0,
		140, 141, 1, 0, 0, 0, 141, 23, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 143, 144,
		5, 10, 0, 0, 144, 145, 3, 26, 13, 0, 145, 25, 1, 0, 0, 0, 146, 150, 3,
		30, 15, 0, 147, 149, 3, 28, 14, 0, 148, 147, 1, 0, 0, 0, 149, 152, 1, 0,
		0, 0, 150, 148, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 27, 1, 0, 0, 0,
		152, 150, 1, 0, 0, 0, 153, 154, 7, 0, 0, 0, 154, 155, 3, 30, 15, 0, 155,
		29, 1, 0, 0, 0, 156, 160, 3, 34, 17, 0, 157, 159, 3, 32, 16, 0, 158, 157,
		1, 0, 0, 0, 159, 162, 1, 0, 0, 0, 160, 158, 1, 0, 0, 0, 160, 161, 1, 0,
		0, 0, 161, 31, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 163, 164, 7, 1, 0, 0,
		164, 165, 3, 34, 17, 0, 165, 33, 1, 0, 0, 0, 166, 170, 3, 38, 19, 0, 167,
		169, 3, 36, 18, 0, 168, 167, 1, 0, 0, 0, 169, 172, 1, 0, 0, 0, 170, 168,
		1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 35, 1, 0, 0, 0, 172, 170, 1, 0,
		0, 0, 173, 174, 7, 2, 0, 0, 174, 175, 3, 38, 19, 0, 175, 37, 1, 0, 0, 0,
		176, 177, 7, 3, 0, 0, 177, 180, 3, 38, 19, 0, 178, 180, 3, 40, 20, 0, 179,
		176, 1, 0, 0, 0, 179, 178, 1, 0, 0, 0, 180, 39, 1, 0, 0, 0, 181, 184, 3,
		42, 21, 0, 182, 183, 5, 21, 0, 0, 183, 185, 3, 38, 19, 0, 184, 182, 1,
		0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 41, 1, 0, 0, 0, 186, 190, 3, 44, 22,
		0, 187, 189, 3, 54, 27, 0, 188, 187, 1, 0, 0, 0, 189, 192, 1, 0, 0, 0,
		190, 188, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 43, 1, 0, 0, 0, 192, 190,
		1, 0, 0, 0, 193, 195, 5, 22, 0, 0, 194, 196, 3, 48, 24, 0, 195, 194, 1,
		0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 210, 5, 23, 0,
		0, 198, 200, 5, 24, 0, 0, 199, 201, 3, 48, 24, 0, 200, 199, 1, 0, 0, 0,
		200, 201, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 210, 5, 25, 0, 0, 203,
		205, 5, 26, 0, 0, 204, 206, 3, 62, 31, 0, 205, 204, 1, 0, 0, 0, 205, 206,
		1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 210, 5, 27, 0, 0, 208, 210, 3, 46,
		23, 0, 209, 193, 1, 0, 0, 0, 209, 198, 1, 0, 0, 0, 209, 203, 1, 0, 0, 0,
		209, 208, 1, 0, 0, 0, 210, 45, 1, 0, 0, 0, 211, 223, 5, 43, 0, 0, 212,
		223, 5, 33, 0, 0, 213, 215, 5, 32, 0, 0, 214, 213, 1, 0, 0, 0, 215, 216,
		1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 223, 1, 0,
		0, 0, 218, 223, 5, 28, 0, 0, 219, 223, 5, 39, 0, 0, 220, 223, 5, 42, 0,
		0, 221, 223, 5, 36, 0, 0, 222, 211, 1, 0, 0, 0, 222, 212, 1, 0, 0, 0, 222,
		214, 1, 0, 0, 0, 222, 218, 1, 0, 0, 0, 222, 219, 1, 0, 0, 0, 222, 220,
		1, 0, 0, 0, 222, 221, 1, 0, 0, 0, 223, 47, 1, 0, 0, 0, 224, 229, 3, 50,
		25, 0, 225, 226, 5, 29, 0, 0, 226, 228, 3, 50, 25, 0, 227, 225, 1, 0, 0,
		0, 228, 231, 1, 0, 0, 0, 229, 227, 1, 0, 0, 0, 229, 230, 1, 0, 0, 0, 230,
		233, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 232, 234, 5, 29, 0, 0, 233, 232,
		1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 49, 1, 0, 0, 0, 235, 238, 3, 2,
		1, 0, 236, 238, 3, 52, 26, 0, 237, 235, 1, 0, 0, 0, 237, 236, 1, 0, 0,
		0, 238, 51, 1, 0, 0, 0, 239, 240, 5, 15, 0, 0, 240, 241, 3, 2, 1, 0, 241,
		53, 1, 0, 0, 0, 242, 243, 5, 24, 0, 0, 243, 244, 3, 56, 28, 0, 244, 245,
		5, 25, 0, 0, 245, 249, 1, 0, 0, 0, 246, 247, 5, 30, 0, 0, 247, 249, 5,
		43, 0, 0, 248, 242, 1, 0, 0, 0, 248, 246, 1, 0, 0, 0, 249, 55, 1, 0, 0,
		0, 250, 255, 3, 58, 29, 0, 251, 252, 5, 29, 0, 0, 252, 254, 3, 58, 29,
		0, 253, 251, 1, 0, 0, 0, 254, 257, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 255,
		256, 1, 0, 0, 0, 256, 259, 1, 0, 0, 0, 257, 255, 1, 0, 0, 0, 258, 260,
		5, 29, 0, 0, 259, 258, 1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260, 57, 1, 0,
		0, 0, 261, 273, 3, 2, 1, 0, 262, 264, 3, 2, 1, 0, 263, 262, 1, 0, 0, 0,
		263, 264, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0, 265, 267, 5, 31, 0, 0, 266,
		268, 3, 2, 1, 0, 267, 266, 1, 0, 0, 0, 267, 268, 1, 0, 0, 0, 268, 270,
		1, 0, 0, 0, 269, 271, 3, 60, 30, 0, 270, 269, 1, 0, 0, 0, 270, 271, 1,
		0, 0, 0, 271, 273, 1, 0, 0, 0, 272, 261, 1, 0, 0, 0, 272, 263, 1, 0, 0,
		0, 273, 59, 1, 0, 0, 0, 274, 276, 5, 31, 0, 0, 275, 277, 3, 2, 1, 0, 276,
		275, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 61, 1, 0, 0, 0, 278, 283, 3,
		64, 32, 0, 279, 280, 5, 29, 0, 0, 280, 282, 3, 64, 32, 0, 281, 279, 1,
		0, 0, 0, 282, 285, 1, 0, 0, 0, 283, 281, 1, 0, 0, 0, 283, 284, 1, 0, 0,
		0, 284, 287, 1, 0, 0, 0, 285, 283, 1, 0, 0, 0, 286, 288, 5, 29, 0, 0, 287,
		286, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288, 301, 1, 0, 0, 0, 289, 294,
		3, 2, 1, 0, 290, 291, 5, 29, 0, 0, 291, 293, 3, 2, 1, 0, 292, 290, 1, 0,
		0, 0, 293, 296, 1, 0, 0, 0, 294, 292, 1, 0, 0, 0, 294, 295, 1, 0, 0, 0,
		295, 298, 1, 0, 0, 0, 296, 294, 1, 0, 0, 0, 297, 299, 5, 29, 0, 0, 298,
		297, 1, 0, 0, 0, 298, 299, 1, 0, 0, 0, 299, 301, 1, 0, 0, 0, 300, 278,
		1, 0, 0, 0, 300, 289, 1, 0, 0, 0, 301, 63, 1, 0, 0, 0, 302, 303, 3, 2,
		1, 0, 303, 304, 5, 31, 0, 0, 304, 305, 3, 2, 1, 0, 305, 309, 1, 0, 0, 0,
		306, 307, 5, 21, 0, 0, 307, 309, 3, 14, 7, 0, 308, 302, 1, 0, 0, 0, 308,
		306, 1, 0, 0, 0, 309, 65, 1, 0, 0, 0, 37, 76, 84, 90, 98, 114, 120, 130,
		140, 150, 160, 170, 179, 184, 190, 195, 200, 205, 209, 216, 222, 229, 233,
		237, 248, 255, 259, 263, 267, 270, 272, 276, 283, 287, 294, 298, 300, 308,
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

// ExprParserInit initializes any static state used to implement ExprParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewExprParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ExprParserInit() {
	staticData := &exprParserStaticData
	staticData.once.Do(exprParserInit)
}

// NewExprParser produces a new parser instance for the optional input antlr.TokenStream.
func NewExprParser(input antlr.TokenStream) *ExprParser {
	ExprParserInit()
	this := new(ExprParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &exprParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Expr.g4"

	return this
}

// ExprParser tokens.
const (
	ExprParserEOF             = antlr.TokenEOF
	ExprParserT__0            = 1
	ExprParserT__1            = 2
	ExprParserT__2            = 3
	ExprParserT__3            = 4
	ExprParserT__4            = 5
	ExprParserT__5            = 6
	ExprParserT__6            = 7
	ExprParserT__7            = 8
	ExprParserT__8            = 9
	ExprParserT__9            = 10
	ExprParserT__10           = 11
	ExprParserT__11           = 12
	ExprParserT__12           = 13
	ExprParserT__13           = 14
	ExprParserT__14           = 15
	ExprParserT__15           = 16
	ExprParserT__16           = 17
	ExprParserT__17           = 18
	ExprParserT__18           = 19
	ExprParserT__19           = 20
	ExprParserT__20           = 21
	ExprParserT__21           = 22
	ExprParserT__22           = 23
	ExprParserT__23           = 24
	ExprParserT__24           = 25
	ExprParserT__25           = 26
	ExprParserT__26           = 27
	ExprParserT__27           = 28
	ExprParserT__28           = 29
	ExprParserT__29           = 30
	ExprParserT__30           = 31
	ExprParserSTRING          = 32
	ExprParserNUMBER          = 33
	ExprParserINTEGER         = 34
	ExprParserAND             = 35
	ExprParserFALSE           = 36
	ExprParserIN              = 37
	ExprParserIS              = 38
	ExprParserNONE            = 39
	ExprParserNOT             = 40
	ExprParserOR              = 41
	ExprParserTRUE            = 42
	ExprParserNAME            = 43
	ExprParserSTRING_LITERAL  = 44
	ExprParserBYTES_LITERAL   = 45
	ExprParserDECIMAL_INTEGER = 46
	ExprParserOCT_INTEGER     = 47
	ExprParserHEX_INTEGER     = 48
	ExprParserBIN_INTEGER     = 49
	ExprParserFLOAT_NUMBER    = 50
	ExprParserIMAG_NUMBER     = 51
	ExprParserSKIP_           = 52
	ExprParserUNKNOWN_CHAR    = 53
	ExprParserINDENT          = 54
	ExprParserDEDENT          = 55
)

// ExprParser rules.
const (
	ExprParserRULE_singleExpr     = 0
	ExprParserRULE_expr           = 1
	ExprParserRULE_orTest         = 2
	ExprParserRULE_andTest        = 3
	ExprParserRULE_notTest        = 4
	ExprParserRULE_comparison     = 5
	ExprParserRULE_compOp         = 6
	ExprParserRULE_exprMain       = 7
	ExprParserRULE_exprCont       = 8
	ExprParserRULE_xorExpr        = 9
	ExprParserRULE_xorExprCont    = 10
	ExprParserRULE_andExpr        = 11
	ExprParserRULE_andExprCont    = 12
	ExprParserRULE_shiftExpr      = 13
	ExprParserRULE_shiftExprCont  = 14
	ExprParserRULE_arithExpr      = 15
	ExprParserRULE_arithExprCont  = 16
	ExprParserRULE_term           = 17
	ExprParserRULE_termCont       = 18
	ExprParserRULE_factor         = 19
	ExprParserRULE_power          = 20
	ExprParserRULE_atomExpr       = 21
	ExprParserRULE_atom           = 22
	ExprParserRULE_constVal       = 23
	ExprParserRULE_testListComp   = 24
	ExprParserRULE_exprOrStarExpr = 25
	ExprParserRULE_starExpr       = 26
	ExprParserRULE_trailer        = 27
	ExprParserRULE_subscriptList  = 28
	ExprParserRULE_subscript      = 29
	ExprParserRULE_sliceOp        = 30
	ExprParserRULE_dictOrSetMaker = 31
	ExprParserRULE_dictItem       = 32
)

// ISingleExprContext is an interface to support dynamic dispatch.
type ISingleExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	EOF() antlr.TerminalNode

	// IsSingleExprContext differentiates from other interfaces.
	IsSingleExprContext()
}

type SingleExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleExprContext() *SingleExprContext {
	var p = new(SingleExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_singleExpr
	return p
}

func (*SingleExprContext) IsSingleExprContext() {}

func NewSingleExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleExprContext {
	var p = new(SingleExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_singleExpr

	return p
}

func (s *SingleExprContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SingleExprContext) EOF() antlr.TerminalNode {
	return s.GetToken(ExprParserEOF, 0)
}

func (s *SingleExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterSingleExpr(s)
	}
}

func (s *SingleExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitSingleExpr(s)
	}
}

func (s *SingleExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitSingleExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) SingleExpr() (localctx ISingleExprContext) {
	this := p
	_ = this

	localctx = NewSingleExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ExprParserRULE_singleExpr)

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
		p.SetState(66)
		p.Expr()
	}
	{
		p.SetState(67)
		p.Match(ExprParserEOF)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OrTest() IOrTestContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) OrTest() IOrTestContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrTestContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrTestContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) Expr() (localctx IExprContext) {
	this := p
	_ = this

	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ExprParserRULE_expr)

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
		p.SetState(69)
		p.OrTest()
	}

	return localctx
}

// IOrTestContext is an interface to support dynamic dispatch.
type IOrTestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAndTest() []IAndTestContext
	AndTest(i int) IAndTestContext
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode

	// IsOrTestContext differentiates from other interfaces.
	IsOrTestContext()
}

type OrTestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrTestContext() *OrTestContext {
	var p = new(OrTestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_orTest
	return p
}

func (*OrTestContext) IsOrTestContext() {}

func NewOrTestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrTestContext {
	var p = new(OrTestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_orTest

	return p
}

func (s *OrTestContext) GetParser() antlr.Parser { return s.parser }

func (s *OrTestContext) AllAndTest() []IAndTestContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAndTestContext); ok {
			len++
		}
	}

	tst := make([]IAndTestContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAndTestContext); ok {
			tst[i] = t.(IAndTestContext)
			i++
		}
	}

	return tst
}

func (s *OrTestContext) AndTest(i int) IAndTestContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndTestContext); ok {
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

	return t.(IAndTestContext)
}

func (s *OrTestContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(ExprParserOR)
}

func (s *OrTestContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserOR, i)
}

func (s *OrTestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrTestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrTestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterOrTest(s)
	}
}

func (s *OrTestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitOrTest(s)
	}
}

func (s *OrTestContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitOrTest(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) OrTest() (localctx IOrTestContext) {
	this := p
	_ = this

	localctx = NewOrTestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ExprParserRULE_orTest)
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
		p.SetState(71)
		p.AndTest()
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ExprParserOR {
		{
			p.SetState(72)
			p.Match(ExprParserOR)
		}
		{
			p.SetState(73)
			p.AndTest()
		}

		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAndTestContext is an interface to support dynamic dispatch.
type IAndTestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNotTest() []INotTestContext
	NotTest(i int) INotTestContext
	AllAND() []antlr.TerminalNode
	AND(i int) antlr.TerminalNode

	// IsAndTestContext differentiates from other interfaces.
	IsAndTestContext()
}

type AndTestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndTestContext() *AndTestContext {
	var p = new(AndTestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_andTest
	return p
}

func (*AndTestContext) IsAndTestContext() {}

func NewAndTestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndTestContext {
	var p = new(AndTestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_andTest

	return p
}

func (s *AndTestContext) GetParser() antlr.Parser { return s.parser }

func (s *AndTestContext) AllNotTest() []INotTestContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INotTestContext); ok {
			len++
		}
	}

	tst := make([]INotTestContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INotTestContext); ok {
			tst[i] = t.(INotTestContext)
			i++
		}
	}

	return tst
}

func (s *AndTestContext) NotTest(i int) INotTestContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INotTestContext); ok {
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

	return t.(INotTestContext)
}

func (s *AndTestContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(ExprParserAND)
}

func (s *AndTestContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserAND, i)
}

func (s *AndTestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndTestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndTestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterAndTest(s)
	}
}

func (s *AndTestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitAndTest(s)
	}
}

func (s *AndTestContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitAndTest(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) AndTest() (localctx IAndTestContext) {
	this := p
	_ = this

	localctx = NewAndTestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ExprParserRULE_andTest)
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
		p.SetState(79)
		p.NotTest()
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ExprParserAND {
		{
			p.SetState(80)
			p.Match(ExprParserAND)
		}
		{
			p.SetState(81)
			p.NotTest()
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INotTestContext is an interface to support dynamic dispatch.
type INotTestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NOT() antlr.TerminalNode
	NotTest() INotTestContext
	Comparison() IComparisonContext

	// IsNotTestContext differentiates from other interfaces.
	IsNotTestContext()
}

type NotTestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotTestContext() *NotTestContext {
	var p = new(NotTestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_notTest
	return p
}

func (*NotTestContext) IsNotTestContext() {}

func NewNotTestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotTestContext {
	var p = new(NotTestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_notTest

	return p
}

func (s *NotTestContext) GetParser() antlr.Parser { return s.parser }

func (s *NotTestContext) NOT() antlr.TerminalNode {
	return s.GetToken(ExprParserNOT, 0)
}

func (s *NotTestContext) NotTest() INotTestContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INotTestContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INotTestContext)
}

func (s *NotTestContext) Comparison() IComparisonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonContext)
}

func (s *NotTestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotTestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotTestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterNotTest(s)
	}
}

func (s *NotTestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitNotTest(s)
	}
}

func (s *NotTestContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitNotTest(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) NotTest() (localctx INotTestContext) {
	this := p
	_ = this

	localctx = NewNotTestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ExprParserRULE_notTest)

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

	p.SetState(90)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ExprParserNOT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)
			p.Match(ExprParserNOT)
		}
		{
			p.SetState(88)
			p.NotTest()
		}

	case ExprParserT__12, ExprParserT__13, ExprParserT__19, ExprParserT__21, ExprParserT__23, ExprParserT__25, ExprParserT__27, ExprParserSTRING, ExprParserNUMBER, ExprParserFALSE, ExprParserNONE, ExprParserTRUE, ExprParserNAME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.Comparison()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IComparisonContext is an interface to support dynamic dispatch.
type IComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExprMain() []IExprMainContext
	ExprMain(i int) IExprMainContext
	AllCompOp() []ICompOpContext
	CompOp(i int) ICompOpContext

	// IsComparisonContext differentiates from other interfaces.
	IsComparisonContext()
}

type ComparisonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonContext() *ComparisonContext {
	var p = new(ComparisonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_comparison
	return p
}

func (*ComparisonContext) IsComparisonContext() {}

func NewComparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonContext {
	var p = new(ComparisonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_comparison

	return p
}

func (s *ComparisonContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonContext) AllExprMain() []IExprMainContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprMainContext); ok {
			len++
		}
	}

	tst := make([]IExprMainContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprMainContext); ok {
			tst[i] = t.(IExprMainContext)
			i++
		}
	}

	return tst
}

func (s *ComparisonContext) ExprMain(i int) IExprMainContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprMainContext); ok {
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

	return t.(IExprMainContext)
}

func (s *ComparisonContext) AllCompOp() []ICompOpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICompOpContext); ok {
			len++
		}
	}

	tst := make([]ICompOpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICompOpContext); ok {
			tst[i] = t.(ICompOpContext)
			i++
		}
	}

	return tst
}

func (s *ComparisonContext) CompOp(i int) ICompOpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompOpContext); ok {
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

	return t.(ICompOpContext)
}

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterComparison(s)
	}
}

func (s *ComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitComparison(s)
	}
}

func (s *ComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) Comparison() (localctx IComparisonContext) {
	this := p
	_ = this

	localctx = NewComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ExprParserRULE_comparison)
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
		p.SetState(92)
		p.ExprMain()
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1511828488446) != 0 {
		{
			p.SetState(93)
			p.CompOp()
		}
		{
			p.SetState(94)
			p.ExprMain()
		}

		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICompOpContext is an interface to support dynamic dispatch.
type ICompOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IN() antlr.TerminalNode
	NOT() antlr.TerminalNode
	IS() antlr.TerminalNode

	// IsCompOpContext differentiates from other interfaces.
	IsCompOpContext()
}

type CompOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompOpContext() *CompOpContext {
	var p = new(CompOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_compOp
	return p
}

func (*CompOpContext) IsCompOpContext() {}

func NewCompOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompOpContext {
	var p = new(CompOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_compOp

	return p
}

func (s *CompOpContext) GetParser() antlr.Parser { return s.parser }

func (s *CompOpContext) IN() antlr.TerminalNode {
	return s.GetToken(ExprParserIN, 0)
}

func (s *CompOpContext) NOT() antlr.TerminalNode {
	return s.GetToken(ExprParserNOT, 0)
}

func (s *CompOpContext) IS() antlr.TerminalNode {
	return s.GetToken(ExprParserIS, 0)
}

func (s *CompOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterCompOp(s)
	}
}

func (s *CompOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitCompOp(s)
	}
}

func (s *CompOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitCompOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) CompOp() (localctx ICompOpContext) {
	this := p
	_ = this

	localctx = NewCompOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ExprParserRULE_compOp)

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

	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(101)
			p.Match(ExprParserT__0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(102)
			p.Match(ExprParserT__1)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(103)
			p.Match(ExprParserT__2)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(104)
			p.Match(ExprParserT__3)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(105)
			p.Match(ExprParserT__4)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(106)
			p.Match(ExprParserT__5)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(107)
			p.Match(ExprParserT__6)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(108)
			p.Match(ExprParserIN)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(109)
			p.Match(ExprParserNOT)
		}
		{
			p.SetState(110)
			p.Match(ExprParserIN)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(111)
			p.Match(ExprParserIS)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(112)
			p.Match(ExprParserIS)
		}
		{
			p.SetState(113)
			p.Match(ExprParserNOT)
		}

	}

	return localctx
}

// IExprMainContext is an interface to support dynamic dispatch.
type IExprMainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	XorExpr() IXorExprContext
	AllExprCont() []IExprContContext
	ExprCont(i int) IExprContContext

	// IsExprMainContext differentiates from other interfaces.
	IsExprMainContext()
}

type ExprMainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprMainContext() *ExprMainContext {
	var p = new(ExprMainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_exprMain
	return p
}

func (*ExprMainContext) IsExprMainContext() {}

func NewExprMainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprMainContext {
	var p = new(ExprMainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_exprMain

	return p
}

func (s *ExprMainContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprMainContext) XorExpr() IXorExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IXorExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IXorExprContext)
}

func (s *ExprMainContext) AllExprCont() []IExprContContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContContext); ok {
			len++
		}
	}

	tst := make([]IExprContContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContContext); ok {
			tst[i] = t.(IExprContContext)
			i++
		}
	}

	return tst
}

func (s *ExprMainContext) ExprCont(i int) IExprContContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContContext); ok {
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

	return t.(IExprContContext)
}

func (s *ExprMainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprMainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprMainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterExprMain(s)
	}
}

func (s *ExprMainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitExprMain(s)
	}
}

func (s *ExprMainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitExprMain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) ExprMain() (localctx IExprMainContext) {
	this := p
	_ = this

	localctx = NewExprMainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ExprParserRULE_exprMain)
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
		p.SetState(116)
		p.XorExpr()
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ExprParserT__7 {
		{
			p.SetState(117)
			p.ExprCont()
		}

		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExprContContext is an interface to support dynamic dispatch.
type IExprContContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	XorExpr() IXorExprContext

	// IsExprContContext differentiates from other interfaces.
	IsExprContContext()
}

type ExprContContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyExprContContext() *ExprContContext {
	var p = new(ExprContContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_exprCont
	return p
}

func (*ExprContContext) IsExprContContext() {}

func NewExprContContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContContext {
	var p = new(ExprContContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_exprCont

	return p
}

func (s *ExprContContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContContext) GetOp() antlr.Token { return s.op }

func (s *ExprContContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContContext) XorExpr() IXorExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IXorExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IXorExprContext)
}

func (s *ExprContContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterExprCont(s)
	}
}

func (s *ExprContContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitExprCont(s)
	}
}

func (s *ExprContContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitExprCont(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) ExprCont() (localctx IExprContContext) {
	this := p
	_ = this

	localctx = NewExprContContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ExprParserRULE_exprCont)

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
		p.SetState(123)

		var _m = p.Match(ExprParserT__7)

		localctx.(*ExprContContext).op = _m
	}
	{
		p.SetState(124)
		p.XorExpr()
	}

	return localctx
}

// IXorExprContext is an interface to support dynamic dispatch.
type IXorExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AndExpr() IAndExprContext
	AllXorExprCont() []IXorExprContContext
	XorExprCont(i int) IXorExprContContext

	// IsXorExprContext differentiates from other interfaces.
	IsXorExprContext()
}

type XorExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyXorExprContext() *XorExprContext {
	var p = new(XorExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_xorExpr
	return p
}

func (*XorExprContext) IsXorExprContext() {}

func NewXorExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *XorExprContext {
	var p = new(XorExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_xorExpr

	return p
}

func (s *XorExprContext) GetParser() antlr.Parser { return s.parser }

func (s *XorExprContext) AndExpr() IAndExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndExprContext)
}

func (s *XorExprContext) AllXorExprCont() []IXorExprContContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IXorExprContContext); ok {
			len++
		}
	}

	tst := make([]IXorExprContContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IXorExprContContext); ok {
			tst[i] = t.(IXorExprContContext)
			i++
		}
	}

	return tst
}

func (s *XorExprContext) XorExprCont(i int) IXorExprContContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IXorExprContContext); ok {
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

	return t.(IXorExprContContext)
}

func (s *XorExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XorExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *XorExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterXorExpr(s)
	}
}

func (s *XorExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitXorExpr(s)
	}
}

func (s *XorExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitXorExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) XorExpr() (localctx IXorExprContext) {
	this := p
	_ = this

	localctx = NewXorExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ExprParserRULE_xorExpr)
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
		p.SetState(126)
		p.AndExpr()
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ExprParserT__8 {
		{
			p.SetState(127)
			p.XorExprCont()
		}

		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IXorExprContContext is an interface to support dynamic dispatch.
type IXorExprContContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	AndExpr() IAndExprContext

	// IsXorExprContContext differentiates from other interfaces.
	IsXorExprContContext()
}

type XorExprContContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyXorExprContContext() *XorExprContContext {
	var p = new(XorExprContContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_xorExprCont
	return p
}

func (*XorExprContContext) IsXorExprContContext() {}

func NewXorExprContContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *XorExprContContext {
	var p = new(XorExprContContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_xorExprCont

	return p
}

func (s *XorExprContContext) GetParser() antlr.Parser { return s.parser }

func (s *XorExprContContext) GetOp() antlr.Token { return s.op }

func (s *XorExprContContext) SetOp(v antlr.Token) { s.op = v }

func (s *XorExprContContext) AndExpr() IAndExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndExprContext)
}

func (s *XorExprContContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XorExprContContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *XorExprContContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterXorExprCont(s)
	}
}

func (s *XorExprContContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitXorExprCont(s)
	}
}

func (s *XorExprContContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitXorExprCont(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) XorExprCont() (localctx IXorExprContContext) {
	this := p
	_ = this

	localctx = NewXorExprContContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ExprParserRULE_xorExprCont)

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

		var _m = p.Match(ExprParserT__8)

		localctx.(*XorExprContContext).op = _m
	}
	{
		p.SetState(134)
		p.AndExpr()
	}

	return localctx
}

// IAndExprContext is an interface to support dynamic dispatch.
type IAndExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ShiftExpr() IShiftExprContext
	AllAndExprCont() []IAndExprContContext
	AndExprCont(i int) IAndExprContContext

	// IsAndExprContext differentiates from other interfaces.
	IsAndExprContext()
}

type AndExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndExprContext() *AndExprContext {
	var p = new(AndExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_andExpr
	return p
}

func (*AndExprContext) IsAndExprContext() {}

func NewAndExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndExprContext {
	var p = new(AndExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_andExpr

	return p
}

func (s *AndExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AndExprContext) ShiftExpr() IShiftExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShiftExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShiftExprContext)
}

func (s *AndExprContext) AllAndExprCont() []IAndExprContContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAndExprContContext); ok {
			len++
		}
	}

	tst := make([]IAndExprContContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAndExprContContext); ok {
			tst[i] = t.(IAndExprContContext)
			i++
		}
	}

	return tst
}

func (s *AndExprContext) AndExprCont(i int) IAndExprContContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndExprContContext); ok {
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

	return t.(IAndExprContContext)
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) AndExpr() (localctx IAndExprContext) {
	this := p
	_ = this

	localctx = NewAndExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ExprParserRULE_andExpr)
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
		p.SetState(136)
		p.ShiftExpr()
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ExprParserT__9 {
		{
			p.SetState(137)
			p.AndExprCont()
		}

		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAndExprContContext is an interface to support dynamic dispatch.
type IAndExprContContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	ShiftExpr() IShiftExprContext

	// IsAndExprContContext differentiates from other interfaces.
	IsAndExprContContext()
}

type AndExprContContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyAndExprContContext() *AndExprContContext {
	var p = new(AndExprContContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_andExprCont
	return p
}

func (*AndExprContContext) IsAndExprContContext() {}

func NewAndExprContContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndExprContContext {
	var p = new(AndExprContContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_andExprCont

	return p
}

func (s *AndExprContContext) GetParser() antlr.Parser { return s.parser }

func (s *AndExprContContext) GetOp() antlr.Token { return s.op }

func (s *AndExprContContext) SetOp(v antlr.Token) { s.op = v }

func (s *AndExprContContext) ShiftExpr() IShiftExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShiftExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShiftExprContext)
}

func (s *AndExprContContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndExprContContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterAndExprCont(s)
	}
}

func (s *AndExprContContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitAndExprCont(s)
	}
}

func (s *AndExprContContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitAndExprCont(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) AndExprCont() (localctx IAndExprContContext) {
	this := p
	_ = this

	localctx = NewAndExprContContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ExprParserRULE_andExprCont)

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

		var _m = p.Match(ExprParserT__9)

		localctx.(*AndExprContContext).op = _m
	}
	{
		p.SetState(144)
		p.ShiftExpr()
	}

	return localctx
}

// IShiftExprContext is an interface to support dynamic dispatch.
type IShiftExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ArithExpr() IArithExprContext
	AllShiftExprCont() []IShiftExprContContext
	ShiftExprCont(i int) IShiftExprContContext

	// IsShiftExprContext differentiates from other interfaces.
	IsShiftExprContext()
}

type ShiftExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShiftExprContext() *ShiftExprContext {
	var p = new(ShiftExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_shiftExpr
	return p
}

func (*ShiftExprContext) IsShiftExprContext() {}

func NewShiftExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShiftExprContext {
	var p = new(ShiftExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_shiftExpr

	return p
}

func (s *ShiftExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ShiftExprContext) ArithExpr() IArithExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithExprContext)
}

func (s *ShiftExprContext) AllShiftExprCont() []IShiftExprContContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IShiftExprContContext); ok {
			len++
		}
	}

	tst := make([]IShiftExprContContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IShiftExprContContext); ok {
			tst[i] = t.(IShiftExprContContext)
			i++
		}
	}

	return tst
}

func (s *ShiftExprContext) ShiftExprCont(i int) IShiftExprContContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShiftExprContContext); ok {
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

	return t.(IShiftExprContContext)
}

func (s *ShiftExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShiftExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShiftExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterShiftExpr(s)
	}
}

func (s *ShiftExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitShiftExpr(s)
	}
}

func (s *ShiftExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitShiftExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) ShiftExpr() (localctx IShiftExprContext) {
	this := p
	_ = this

	localctx = NewShiftExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ExprParserRULE_shiftExpr)
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
		p.SetState(146)
		p.ArithExpr()
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ExprParserT__10 || _la == ExprParserT__11 {
		{
			p.SetState(147)
			p.ShiftExprCont()
		}

		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IShiftExprContContext is an interface to support dynamic dispatch.
type IShiftExprContContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	ArithExpr() IArithExprContext

	// IsShiftExprContContext differentiates from other interfaces.
	IsShiftExprContContext()
}

type ShiftExprContContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyShiftExprContContext() *ShiftExprContContext {
	var p = new(ShiftExprContContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_shiftExprCont
	return p
}

func (*ShiftExprContContext) IsShiftExprContContext() {}

func NewShiftExprContContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShiftExprContContext {
	var p = new(ShiftExprContContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_shiftExprCont

	return p
}

func (s *ShiftExprContContext) GetParser() antlr.Parser { return s.parser }

func (s *ShiftExprContContext) GetOp() antlr.Token { return s.op }

func (s *ShiftExprContContext) SetOp(v antlr.Token) { s.op = v }

func (s *ShiftExprContContext) ArithExpr() IArithExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithExprContext)
}

func (s *ShiftExprContContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShiftExprContContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShiftExprContContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterShiftExprCont(s)
	}
}

func (s *ShiftExprContContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitShiftExprCont(s)
	}
}

func (s *ShiftExprContContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitShiftExprCont(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) ShiftExprCont() (localctx IShiftExprContContext) {
	this := p
	_ = this

	localctx = NewShiftExprContContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ExprParserRULE_shiftExprCont)
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
		p.SetState(153)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*ShiftExprContContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == ExprParserT__10 || _la == ExprParserT__11) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*ShiftExprContContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(154)
		p.ArithExpr()
	}

	return localctx
}

// IArithExprContext is an interface to support dynamic dispatch.
type IArithExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Term() ITermContext
	AllArithExprCont() []IArithExprContContext
	ArithExprCont(i int) IArithExprContContext

	// IsArithExprContext differentiates from other interfaces.
	IsArithExprContext()
}

type ArithExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArithExprContext() *ArithExprContext {
	var p = new(ArithExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_arithExpr
	return p
}

func (*ArithExprContext) IsArithExprContext() {}

func NewArithExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArithExprContext {
	var p = new(ArithExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_arithExpr

	return p
}

func (s *ArithExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ArithExprContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ArithExprContext) AllArithExprCont() []IArithExprContContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArithExprContContext); ok {
			len++
		}
	}

	tst := make([]IArithExprContContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArithExprContContext); ok {
			tst[i] = t.(IArithExprContContext)
			i++
		}
	}

	return tst
}

func (s *ArithExprContext) ArithExprCont(i int) IArithExprContContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithExprContContext); ok {
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

	return t.(IArithExprContContext)
}

func (s *ArithExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArithExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterArithExpr(s)
	}
}

func (s *ArithExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitArithExpr(s)
	}
}

func (s *ArithExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitArithExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) ArithExpr() (localctx IArithExprContext) {
	this := p
	_ = this

	localctx = NewArithExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ExprParserRULE_arithExpr)
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
		p.SetState(156)
		p.Term()
	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ExprParserT__12 || _la == ExprParserT__13 {
		{
			p.SetState(157)
			p.ArithExprCont()
		}

		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IArithExprContContext is an interface to support dynamic dispatch.
type IArithExprContContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	Term() ITermContext

	// IsArithExprContContext differentiates from other interfaces.
	IsArithExprContContext()
}

type ArithExprContContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyArithExprContContext() *ArithExprContContext {
	var p = new(ArithExprContContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_arithExprCont
	return p
}

func (*ArithExprContContext) IsArithExprContContext() {}

func NewArithExprContContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArithExprContContext {
	var p = new(ArithExprContContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_arithExprCont

	return p
}

func (s *ArithExprContContext) GetParser() antlr.Parser { return s.parser }

func (s *ArithExprContContext) GetOp() antlr.Token { return s.op }

func (s *ArithExprContContext) SetOp(v antlr.Token) { s.op = v }

func (s *ArithExprContContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ArithExprContContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithExprContContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArithExprContContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterArithExprCont(s)
	}
}

func (s *ArithExprContContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitArithExprCont(s)
	}
}

func (s *ArithExprContContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitArithExprCont(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) ArithExprCont() (localctx IArithExprContContext) {
	this := p
	_ = this

	localctx = NewArithExprContContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ExprParserRULE_arithExprCont)
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
		p.SetState(163)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*ArithExprContContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == ExprParserT__12 || _la == ExprParserT__13) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*ArithExprContContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(164)
		p.Term()
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Factor() IFactorContext
	AllTermCont() []ITermContContext
	TermCont(i int) ITermContContext

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Factor() IFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *TermContext) AllTermCont() []ITermContContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContContext); ok {
			len++
		}
	}

	tst := make([]ITermContContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContContext); ok {
			tst[i] = t.(ITermContContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) TermCont(i int) ITermContContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContContext); ok {
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

	return t.(ITermContContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (s *TermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) Term() (localctx ITermContext) {
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ExprParserRULE_term)
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
		p.SetState(166)
		p.Factor()
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1015808) != 0 {
		{
			p.SetState(167)
			p.TermCont()
		}

		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITermContContext is an interface to support dynamic dispatch.
type ITermContContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	Factor() IFactorContext

	// IsTermContContext differentiates from other interfaces.
	IsTermContContext()
}

type TermContContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyTermContContext() *TermContContext {
	var p = new(TermContContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_termCont
	return p
}

func (*TermContContext) IsTermContContext() {}

func NewTermContContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContContext {
	var p = new(TermContContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_termCont

	return p
}

func (s *TermContContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContContext) GetOp() antlr.Token { return s.op }

func (s *TermContContext) SetOp(v antlr.Token) { s.op = v }

func (s *TermContContext) Factor() IFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *TermContContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterTermCont(s)
	}
}

func (s *TermContContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitTermCont(s)
	}
}

func (s *TermContContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitTermCont(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) TermCont() (localctx ITermContContext) {
	this := p
	_ = this

	localctx = NewTermContContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ExprParserRULE_termCont)
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
		p.SetState(173)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*TermContContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1015808) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*TermContContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(174)
		p.Factor()
	}

	return localctx
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	Factor() IFactorContext
	Power() IPowerContext

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) GetOp() antlr.Token { return s.op }

func (s *FactorContext) SetOp(v antlr.Token) { s.op = v }

func (s *FactorContext) Factor() IFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *FactorContext) Power() IPowerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPowerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPowerContext)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (s *FactorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitFactor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) Factor() (localctx IFactorContext) {
	this := p
	_ = this

	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ExprParserRULE_factor)
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

	p.SetState(179)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ExprParserT__12, ExprParserT__13, ExprParserT__19:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(176)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*FactorContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1073152) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*FactorContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(177)
			p.Factor()
		}

	case ExprParserT__21, ExprParserT__23, ExprParserT__25, ExprParserT__27, ExprParserSTRING, ExprParserNUMBER, ExprParserFALSE, ExprParserNONE, ExprParserTRUE, ExprParserNAME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(178)
			p.Power()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPowerContext is an interface to support dynamic dispatch.
type IPowerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	AtomExpr() IAtomExprContext
	Factor() IFactorContext

	// IsPowerContext differentiates from other interfaces.
	IsPowerContext()
}

type PowerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyPowerContext() *PowerContext {
	var p = new(PowerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_power
	return p
}

func (*PowerContext) IsPowerContext() {}

func NewPowerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PowerContext {
	var p = new(PowerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_power

	return p
}

func (s *PowerContext) GetParser() antlr.Parser { return s.parser }

func (s *PowerContext) GetOp() antlr.Token { return s.op }

func (s *PowerContext) SetOp(v antlr.Token) { s.op = v }

func (s *PowerContext) AtomExpr() IAtomExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomExprContext)
}

func (s *PowerContext) Factor() IFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *PowerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PowerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterPower(s)
	}
}

func (s *PowerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitPower(s)
	}
}

func (s *PowerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitPower(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) Power() (localctx IPowerContext) {
	this := p
	_ = this

	localctx = NewPowerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ExprParserRULE_power)
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
		p.SetState(181)
		p.AtomExpr()
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ExprParserT__20 {
		{
			p.SetState(182)

			var _m = p.Match(ExprParserT__20)

			localctx.(*PowerContext).op = _m
		}
		{
			p.SetState(183)
			p.Factor()
		}

	}

	return localctx
}

// IAtomExprContext is an interface to support dynamic dispatch.
type IAtomExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Atom() IAtomContext
	AllTrailer() []ITrailerContext
	Trailer(i int) ITrailerContext

	// IsAtomExprContext differentiates from other interfaces.
	IsAtomExprContext()
}

type AtomExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomExprContext() *AtomExprContext {
	var p = new(AtomExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_atomExpr
	return p
}

func (*AtomExprContext) IsAtomExprContext() {}

func NewAtomExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomExprContext {
	var p = new(AtomExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_atomExpr

	return p
}

func (s *AtomExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomExprContext) Atom() IAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *AtomExprContext) AllTrailer() []ITrailerContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITrailerContext); ok {
			len++
		}
	}

	tst := make([]ITrailerContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITrailerContext); ok {
			tst[i] = t.(ITrailerContext)
			i++
		}
	}

	return tst
}

func (s *AtomExprContext) Trailer(i int) ITrailerContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITrailerContext); ok {
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

	return t.(ITrailerContext)
}

func (s *AtomExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterAtomExpr(s)
	}
}

func (s *AtomExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitAtomExpr(s)
	}
}

func (s *AtomExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitAtomExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) AtomExpr() (localctx IAtomExprContext) {
	this := p
	_ = this

	localctx = NewAtomExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ExprParserRULE_atomExpr)
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
		p.Atom()
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ExprParserT__23 || _la == ExprParserT__29 {
		{
			p.SetState(187)
			p.Trailer()
		}

		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) CopyFrom(ctx *AtomContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BracketAtomContext struct {
	*AtomContext
}

func NewBracketAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BracketAtomContext {
	var p = new(BracketAtomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *BracketAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketAtomContext) TestListComp() ITestListCompContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITestListCompContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITestListCompContext)
}

func (s *BracketAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterBracketAtom(s)
	}
}

func (s *BracketAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitBracketAtom(s)
	}
}

func (s *BracketAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitBracketAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type DictOrSetAtomContext struct {
	*AtomContext
}

func NewDictOrSetAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DictOrSetAtomContext {
	var p = new(DictOrSetAtomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *DictOrSetAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictOrSetAtomContext) DictOrSetMaker() IDictOrSetMakerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDictOrSetMakerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDictOrSetMakerContext)
}

func (s *DictOrSetAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterDictOrSetAtom(s)
	}
}

func (s *DictOrSetAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitDictOrSetAtom(s)
	}
}

func (s *DictOrSetAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitDictOrSetAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConstAtomContext struct {
	*AtomContext
}

func NewConstAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstAtomContext {
	var p = new(ConstAtomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *ConstAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstAtomContext) ConstVal() IConstValContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstValContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstValContext)
}

func (s *ConstAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterConstAtom(s)
	}
}

func (s *ConstAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitConstAtom(s)
	}
}

func (s *ConstAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitConstAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenAtomContext struct {
	*AtomContext
}

func NewParenAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenAtomContext {
	var p = new(ParenAtomContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *ParenAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenAtomContext) TestListComp() ITestListCompContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITestListCompContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITestListCompContext)
}

func (s *ParenAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterParenAtom(s)
	}
}

func (s *ParenAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitParenAtom(s)
	}
}

func (s *ParenAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitParenAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) Atom() (localctx IAtomContext) {
	this := p
	_ = this

	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ExprParserRULE_atom)
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

	p.SetState(209)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ExprParserT__21:
		localctx = NewParenAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(193)
			p.Match(ExprParserT__21)
		}
		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14925368975360) != 0 {
			{
				p.SetState(194)
				p.TestListComp()
			}

		}
		{
			p.SetState(197)
			p.Match(ExprParserT__22)
		}

	case ExprParserT__23:
		localctx = NewBracketAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(198)
			p.Match(ExprParserT__23)
		}
		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14925368975360) != 0 {
			{
				p.SetState(199)
				p.TestListComp()
			}

		}
		{
			p.SetState(202)
			p.Match(ExprParserT__24)
		}

	case ExprParserT__25:
		localctx = NewDictOrSetAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(203)
			p.Match(ExprParserT__25)
		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14925371039744) != 0 {
			{
				p.SetState(204)
				p.DictOrSetMaker()
			}

		}
		{
			p.SetState(207)
			p.Match(ExprParserT__26)
		}

	case ExprParserT__27, ExprParserSTRING, ExprParserNUMBER, ExprParserFALSE, ExprParserNONE, ExprParserTRUE, ExprParserNAME:
		localctx = NewConstAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(208)
			p.ConstVal()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConstValContext is an interface to support dynamic dispatch.
type IConstValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	NONE() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode

	// IsConstValContext differentiates from other interfaces.
	IsConstValContext()
}

type ConstValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstValContext() *ConstValContext {
	var p = new(ConstValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_constVal
	return p
}

func (*ConstValContext) IsConstValContext() {}

func NewConstValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstValContext {
	var p = new(ConstValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_constVal

	return p
}

func (s *ConstValContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstValContext) NAME() antlr.TerminalNode {
	return s.GetToken(ExprParserNAME, 0)
}

func (s *ConstValContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ExprParserNUMBER, 0)
}

func (s *ConstValContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(ExprParserSTRING)
}

func (s *ConstValContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(ExprParserSTRING, i)
}

func (s *ConstValContext) NONE() antlr.TerminalNode {
	return s.GetToken(ExprParserNONE, 0)
}

func (s *ConstValContext) TRUE() antlr.TerminalNode {
	return s.GetToken(ExprParserTRUE, 0)
}

func (s *ConstValContext) FALSE() antlr.TerminalNode {
	return s.GetToken(ExprParserFALSE, 0)
}

func (s *ConstValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterConstVal(s)
	}
}

func (s *ConstValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitConstVal(s)
	}
}

func (s *ConstValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitConstVal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) ConstVal() (localctx IConstValContext) {
	this := p
	_ = this

	localctx = NewConstValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ExprParserRULE_constVal)
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

	p.SetState(222)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ExprParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(211)
			p.Match(ExprParserNAME)
		}

	case ExprParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(212)
			p.Match(ExprParserNUMBER)
		}

	case ExprParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ExprParserSTRING {
			{
				p.SetState(213)
				p.Match(ExprParserSTRING)
			}

			p.SetState(216)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case ExprParserT__27:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(218)
			p.Match(ExprParserT__27)
		}

	case ExprParserNONE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(219)
			p.Match(ExprParserNONE)
		}

	case ExprParserTRUE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(220)
			p.Match(ExprParserTRUE)
		}

	case ExprParserFALSE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(221)
			p.Match(ExprParserFALSE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITestListCompContext is an interface to support dynamic dispatch.
type ITestListCompContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExprOrStarExpr() []IExprOrStarExprContext
	ExprOrStarExpr(i int) IExprOrStarExprContext

	// IsTestListCompContext differentiates from other interfaces.
	IsTestListCompContext()
}

type TestListCompContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTestListCompContext() *TestListCompContext {
	var p = new(TestListCompContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_testListComp
	return p
}

func (*TestListCompContext) IsTestListCompContext() {}

func NewTestListCompContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TestListCompContext {
	var p = new(TestListCompContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_testListComp

	return p
}

func (s *TestListCompContext) GetParser() antlr.Parser { return s.parser }

func (s *TestListCompContext) AllExprOrStarExpr() []IExprOrStarExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprOrStarExprContext); ok {
			len++
		}
	}

	tst := make([]IExprOrStarExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprOrStarExprContext); ok {
			tst[i] = t.(IExprOrStarExprContext)
			i++
		}
	}

	return tst
}

func (s *TestListCompContext) ExprOrStarExpr(i int) IExprOrStarExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprOrStarExprContext); ok {
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

	return t.(IExprOrStarExprContext)
}

func (s *TestListCompContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TestListCompContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TestListCompContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterTestListComp(s)
	}
}

func (s *TestListCompContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitTestListComp(s)
	}
}

func (s *TestListCompContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitTestListComp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) TestListComp() (localctx ITestListCompContext) {
	this := p
	_ = this

	localctx = NewTestListCompContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ExprParserRULE_testListComp)
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
	{
		p.SetState(224)
		p.ExprOrStarExpr()
	}

	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(225)
				p.Match(ExprParserT__28)
			}
			{
				p.SetState(226)
				p.ExprOrStarExpr()
			}

		}
		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ExprParserT__28 {
		{
			p.SetState(232)
			p.Match(ExprParserT__28)
		}

	}

	return localctx
}

// IExprOrStarExprContext is an interface to support dynamic dispatch.
type IExprOrStarExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	StarExpr() IStarExprContext

	// IsExprOrStarExprContext differentiates from other interfaces.
	IsExprOrStarExprContext()
}

type ExprOrStarExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprOrStarExprContext() *ExprOrStarExprContext {
	var p = new(ExprOrStarExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_exprOrStarExpr
	return p
}

func (*ExprOrStarExprContext) IsExprOrStarExprContext() {}

func NewExprOrStarExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprOrStarExprContext {
	var p = new(ExprOrStarExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_exprOrStarExpr

	return p
}

func (s *ExprOrStarExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprOrStarExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprOrStarExprContext) StarExpr() IStarExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStarExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStarExprContext)
}

func (s *ExprOrStarExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprOrStarExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprOrStarExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterExprOrStarExpr(s)
	}
}

func (s *ExprOrStarExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitExprOrStarExpr(s)
	}
}

func (s *ExprOrStarExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitExprOrStarExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) ExprOrStarExpr() (localctx IExprOrStarExprContext) {
	this := p
	_ = this

	localctx = NewExprOrStarExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ExprParserRULE_exprOrStarExpr)

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

	p.SetState(237)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ExprParserT__12, ExprParserT__13, ExprParserT__19, ExprParserT__21, ExprParserT__23, ExprParserT__25, ExprParserT__27, ExprParserSTRING, ExprParserNUMBER, ExprParserFALSE, ExprParserNONE, ExprParserNOT, ExprParserTRUE, ExprParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(235)
			p.Expr()
		}

	case ExprParserT__14:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(236)
			p.StarExpr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStarExprContext is an interface to support dynamic dispatch.
type IStarExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsStarExprContext differentiates from other interfaces.
	IsStarExprContext()
}

type StarExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStarExprContext() *StarExprContext {
	var p = new(StarExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_starExpr
	return p
}

func (*StarExprContext) IsStarExprContext() {}

func NewStarExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StarExprContext {
	var p = new(StarExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_starExpr

	return p
}

func (s *StarExprContext) GetParser() antlr.Parser { return s.parser }

func (s *StarExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StarExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StarExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StarExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterStarExpr(s)
	}
}

func (s *StarExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitStarExpr(s)
	}
}

func (s *StarExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitStarExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) StarExpr() (localctx IStarExprContext) {
	this := p
	_ = this

	localctx = NewStarExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ExprParserRULE_starExpr)

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
		p.SetState(239)
		p.Match(ExprParserT__14)
	}
	{
		p.SetState(240)
		p.Expr()
	}

	return localctx
}

// ITrailerContext is an interface to support dynamic dispatch.
type ITrailerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SubscriptList() ISubscriptListContext
	NAME() antlr.TerminalNode

	// IsTrailerContext differentiates from other interfaces.
	IsTrailerContext()
}

type TrailerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTrailerContext() *TrailerContext {
	var p = new(TrailerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_trailer
	return p
}

func (*TrailerContext) IsTrailerContext() {}

func NewTrailerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TrailerContext {
	var p = new(TrailerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_trailer

	return p
}

func (s *TrailerContext) GetParser() antlr.Parser { return s.parser }

func (s *TrailerContext) SubscriptList() ISubscriptListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubscriptListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubscriptListContext)
}

func (s *TrailerContext) NAME() antlr.TerminalNode {
	return s.GetToken(ExprParserNAME, 0)
}

func (s *TrailerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrailerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TrailerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterTrailer(s)
	}
}

func (s *TrailerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitTrailer(s)
	}
}

func (s *TrailerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitTrailer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) Trailer() (localctx ITrailerContext) {
	this := p
	_ = this

	localctx = NewTrailerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ExprParserRULE_trailer)

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

	p.SetState(248)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ExprParserT__23:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(242)
			p.Match(ExprParserT__23)
		}
		{
			p.SetState(243)
			p.SubscriptList()
		}
		{
			p.SetState(244)
			p.Match(ExprParserT__24)
		}

	case ExprParserT__29:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(246)
			p.Match(ExprParserT__29)
		}
		{
			p.SetState(247)
			p.Match(ExprParserNAME)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISubscriptListContext is an interface to support dynamic dispatch.
type ISubscriptListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSubscript() []ISubscriptContext
	Subscript(i int) ISubscriptContext

	// IsSubscriptListContext differentiates from other interfaces.
	IsSubscriptListContext()
}

type SubscriptListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubscriptListContext() *SubscriptListContext {
	var p = new(SubscriptListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_subscriptList
	return p
}

func (*SubscriptListContext) IsSubscriptListContext() {}

func NewSubscriptListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubscriptListContext {
	var p = new(SubscriptListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_subscriptList

	return p
}

func (s *SubscriptListContext) GetParser() antlr.Parser { return s.parser }

func (s *SubscriptListContext) AllSubscript() []ISubscriptContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISubscriptContext); ok {
			len++
		}
	}

	tst := make([]ISubscriptContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISubscriptContext); ok {
			tst[i] = t.(ISubscriptContext)
			i++
		}
	}

	return tst
}

func (s *SubscriptListContext) Subscript(i int) ISubscriptContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubscriptContext); ok {
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

	return t.(ISubscriptContext)
}

func (s *SubscriptListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubscriptListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubscriptListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterSubscriptList(s)
	}
}

func (s *SubscriptListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitSubscriptList(s)
	}
}

func (s *SubscriptListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitSubscriptList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) SubscriptList() (localctx ISubscriptListContext) {
	this := p
	_ = this

	localctx = NewSubscriptListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ExprParserRULE_subscriptList)
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
	{
		p.SetState(250)
		p.Subscript()
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(251)
				p.Match(ExprParserT__28)
			}
			{
				p.SetState(252)
				p.Subscript()
			}

		}
		p.SetState(257)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ExprParserT__28 {
		{
			p.SetState(258)
			p.Match(ExprParserT__28)
		}

	}

	return localctx
}

// ISubscriptContext is an interface to support dynamic dispatch.
type ISubscriptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSubscriptContext differentiates from other interfaces.
	IsSubscriptContext()
}

type SubscriptContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubscriptContext() *SubscriptContext {
	var p = new(SubscriptContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_subscript
	return p
}

func (*SubscriptContext) IsSubscriptContext() {}

func NewSubscriptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubscriptContext {
	var p = new(SubscriptContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_subscript

	return p
}

func (s *SubscriptContext) GetParser() antlr.Parser { return s.parser }

func (s *SubscriptContext) CopyFrom(ctx *SubscriptContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SubscriptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubscriptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SliceSubscriptContext struct {
	*SubscriptContext
	sta IExprContext
	sto IExprContext
	ste ISliceOpContext
}

func NewSliceSubscriptContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceSubscriptContext {
	var p = new(SliceSubscriptContext)

	p.SubscriptContext = NewEmptySubscriptContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SubscriptContext))

	return p
}

func (s *SliceSubscriptContext) GetSta() IExprContext { return s.sta }

func (s *SliceSubscriptContext) GetSto() IExprContext { return s.sto }

func (s *SliceSubscriptContext) GetSte() ISliceOpContext { return s.ste }

func (s *SliceSubscriptContext) SetSta(v IExprContext) { s.sta = v }

func (s *SliceSubscriptContext) SetSto(v IExprContext) { s.sto = v }

func (s *SliceSubscriptContext) SetSte(v ISliceOpContext) { s.ste = v }

func (s *SliceSubscriptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceSubscriptContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *SliceSubscriptContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *SliceSubscriptContext) SliceOp() ISliceOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceOpContext)
}

func (s *SliceSubscriptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterSliceSubscript(s)
	}
}

func (s *SliceSubscriptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitSliceSubscript(s)
	}
}

func (s *SliceSubscriptContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitSliceSubscript(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprSubscriptContext struct {
	*SubscriptContext
}

func NewExprSubscriptContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprSubscriptContext {
	var p = new(ExprSubscriptContext)

	p.SubscriptContext = NewEmptySubscriptContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SubscriptContext))

	return p
}

func (s *ExprSubscriptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprSubscriptContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprSubscriptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterExprSubscript(s)
	}
}

func (s *ExprSubscriptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitExprSubscript(s)
	}
}

func (s *ExprSubscriptContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitExprSubscript(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) Subscript() (localctx ISubscriptContext) {
	this := p
	_ = this

	localctx = NewSubscriptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ExprParserRULE_subscript)
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

	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprSubscriptContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(261)
			p.Expr()
		}

	case 2:
		localctx = NewSliceSubscriptContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14925368942592) != 0 {
			{
				p.SetState(262)

				var _x = p.Expr()

				localctx.(*SliceSubscriptContext).sta = _x
			}

		}
		{
			p.SetState(265)
			p.Match(ExprParserT__30)
		}
		p.SetState(267)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14925368942592) != 0 {
			{
				p.SetState(266)

				var _x = p.Expr()

				localctx.(*SliceSubscriptContext).sto = _x
			}

		}
		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ExprParserT__30 {
			{
				p.SetState(269)

				var _x = p.SliceOp()

				localctx.(*SliceSubscriptContext).ste = _x
			}

		}

	}

	return localctx
}

// ISliceOpContext is an interface to support dynamic dispatch.
type ISliceOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsSliceOpContext differentiates from other interfaces.
	IsSliceOpContext()
}

type SliceOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceOpContext() *SliceOpContext {
	var p = new(SliceOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_sliceOp
	return p
}

func (*SliceOpContext) IsSliceOpContext() {}

func NewSliceOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceOpContext {
	var p = new(SliceOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_sliceOp

	return p
}

func (s *SliceOpContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceOpContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SliceOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterSliceOp(s)
	}
}

func (s *SliceOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitSliceOp(s)
	}
}

func (s *SliceOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitSliceOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) SliceOp() (localctx ISliceOpContext) {
	this := p
	_ = this

	localctx = NewSliceOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ExprParserRULE_sliceOp)
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
		p.SetState(274)
		p.Match(ExprParserT__30)
	}
	p.SetState(276)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14925368942592) != 0 {
		{
			p.SetState(275)
			p.Expr()
		}

	}

	return localctx
}

// IDictOrSetMakerContext is an interface to support dynamic dispatch.
type IDictOrSetMakerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDictOrSetMakerContext differentiates from other interfaces.
	IsDictOrSetMakerContext()
}

type DictOrSetMakerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictOrSetMakerContext() *DictOrSetMakerContext {
	var p = new(DictOrSetMakerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_dictOrSetMaker
	return p
}

func (*DictOrSetMakerContext) IsDictOrSetMakerContext() {}

func NewDictOrSetMakerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictOrSetMakerContext {
	var p = new(DictOrSetMakerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_dictOrSetMaker

	return p
}

func (s *DictOrSetMakerContext) GetParser() antlr.Parser { return s.parser }

func (s *DictOrSetMakerContext) CopyFrom(ctx *DictOrSetMakerContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *DictOrSetMakerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictOrSetMakerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DictMakerContext struct {
	*DictOrSetMakerContext
}

func NewDictMakerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DictMakerContext {
	var p = new(DictMakerContext)

	p.DictOrSetMakerContext = NewEmptyDictOrSetMakerContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DictOrSetMakerContext))

	return p
}

func (s *DictMakerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictMakerContext) AllDictItem() []IDictItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDictItemContext); ok {
			len++
		}
	}

	tst := make([]IDictItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDictItemContext); ok {
			tst[i] = t.(IDictItemContext)
			i++
		}
	}

	return tst
}

func (s *DictMakerContext) DictItem(i int) IDictItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDictItemContext); ok {
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

	return t.(IDictItemContext)
}

func (s *DictMakerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterDictMaker(s)
	}
}

func (s *DictMakerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitDictMaker(s)
	}
}

func (s *DictMakerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitDictMaker(s)

	default:
		return t.VisitChildren(s)
	}
}

type SetMakerContext struct {
	*DictOrSetMakerContext
}

func NewSetMakerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SetMakerContext {
	var p = new(SetMakerContext)

	p.DictOrSetMakerContext = NewEmptyDictOrSetMakerContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DictOrSetMakerContext))

	return p
}

func (s *SetMakerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetMakerContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *SetMakerContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *SetMakerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterSetMaker(s)
	}
}

func (s *SetMakerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitSetMaker(s)
	}
}

func (s *SetMakerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitSetMaker(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) DictOrSetMaker() (localctx IDictOrSetMakerContext) {
	this := p
	_ = this

	localctx = NewDictOrSetMakerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ExprParserRULE_dictOrSetMaker)
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

	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDictMakerContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(278)
			p.DictItem()
		}
		p.SetState(283)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(279)
					p.Match(ExprParserT__28)
				}
				{
					p.SetState(280)
					p.DictItem()
				}

			}
			p.SetState(285)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())
		}
		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ExprParserT__28 {
			{
				p.SetState(286)
				p.Match(ExprParserT__28)
			}

		}

	case 2:
		localctx = NewSetMakerContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(289)
			p.Expr()
		}
		p.SetState(294)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(290)
					p.Match(ExprParserT__28)
				}
				{
					p.SetState(291)
					p.Expr()
				}

			}
			p.SetState(296)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())
		}
		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ExprParserT__28 {
			{
				p.SetState(297)
				p.Match(ExprParserT__28)
			}

		}

	}

	return localctx
}

// IDictItemContext is an interface to support dynamic dispatch.
type IDictItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDictItemContext differentiates from other interfaces.
	IsDictItemContext()
}

type DictItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictItemContext() *DictItemContext {
	var p = new(DictItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_dictItem
	return p
}

func (*DictItemContext) IsDictItemContext() {}

func NewDictItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictItemContext {
	var p = new(DictItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_dictItem

	return p
}

func (s *DictItemContext) GetParser() antlr.Parser { return s.parser }

func (s *DictItemContext) CopyFrom(ctx *DictItemContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *DictItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StarsDictItemContext struct {
	*DictItemContext
}

func NewStarsDictItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StarsDictItemContext {
	var p = new(StarsDictItemContext)

	p.DictItemContext = NewEmptyDictItemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DictItemContext))

	return p
}

func (s *StarsDictItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StarsDictItemContext) ExprMain() IExprMainContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprMainContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprMainContext)
}

func (s *StarsDictItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterStarsDictItem(s)
	}
}

func (s *StarsDictItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitStarsDictItem(s)
	}
}

func (s *StarsDictItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitStarsDictItem(s)

	default:
		return t.VisitChildren(s)
	}
}

type KvDictItemContext struct {
	*DictItemContext
	k IExprContext
	v IExprContext
}

func NewKvDictItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *KvDictItemContext {
	var p = new(KvDictItemContext)

	p.DictItemContext = NewEmptyDictItemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DictItemContext))

	return p
}

func (s *KvDictItemContext) GetK() IExprContext { return s.k }

func (s *KvDictItemContext) GetV() IExprContext { return s.v }

func (s *KvDictItemContext) SetK(v IExprContext) { s.k = v }

func (s *KvDictItemContext) SetV(v IExprContext) { s.v = v }

func (s *KvDictItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KvDictItemContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *KvDictItemContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *KvDictItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterKvDictItem(s)
	}
}

func (s *KvDictItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitKvDictItem(s)
	}
}

func (s *KvDictItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExprVisitor:
		return t.VisitKvDictItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExprParser) DictItem() (localctx IDictItemContext) {
	this := p
	_ = this

	localctx = NewDictItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ExprParserRULE_dictItem)

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

	p.SetState(308)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ExprParserT__12, ExprParserT__13, ExprParserT__19, ExprParserT__21, ExprParserT__23, ExprParserT__25, ExprParserT__27, ExprParserSTRING, ExprParserNUMBER, ExprParserFALSE, ExprParserNONE, ExprParserNOT, ExprParserTRUE, ExprParserNAME:
		localctx = NewKvDictItemContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(302)

			var _x = p.Expr()

			localctx.(*KvDictItemContext).k = _x
		}
		{
			p.SetState(303)
			p.Match(ExprParserT__30)
		}
		{
			p.SetState(304)

			var _x = p.Expr()

			localctx.(*KvDictItemContext).v = _x
		}

	case ExprParserT__20:
		localctx = NewStarsDictItemContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(306)
			p.Match(ExprParserT__20)
		}
		{
			p.SetState(307)
			p.ExprMain()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
