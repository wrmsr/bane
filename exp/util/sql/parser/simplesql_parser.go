// Code generated from SimpleSql.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SimpleSql

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

type SimpleSqlParser struct {
	*antlr.BaseParser
}

var simplesqlParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func simplesqlParserInit() {
	staticData := &simplesqlParserStaticData
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
		"singleStatement", "select", "cteSelect", "cte", "unionSelect", "unionItem",
		"primarySelect", "selectItem", "expression", "booleanExpression", "predicate",
		"valueExpression", "primaryExpression", "simpleExpression", "caseItem",
		"over", "sortItem", "relation", "groupBy", "qualifiedName", "identifierList",
		"identifier", "quotedIdentifier", "number", "str", "null", "true", "false",
		"setQuantifier", "joinType", "cmpOp", "arithOp", "unaryOp", "unquotedIdentifier",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 66, 419, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		2, 1, 2, 1, 2, 1, 2, 5, 2, 79, 8, 2, 10, 2, 12, 2, 82, 9, 2, 3, 2, 84,
		8, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 5, 4,
		96, 8, 4, 10, 4, 12, 4, 99, 9, 4, 1, 5, 1, 5, 3, 5, 103, 8, 5, 1, 5, 1,
		5, 1, 6, 1, 6, 3, 6, 109, 8, 6, 1, 6, 1, 6, 1, 6, 5, 6, 114, 8, 6, 10,
		6, 12, 6, 117, 9, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 123, 8, 6, 10, 6, 12,
		6, 126, 9, 6, 3, 6, 128, 8, 6, 1, 6, 1, 6, 3, 6, 132, 8, 6, 1, 6, 1, 6,
		1, 6, 3, 6, 137, 8, 6, 1, 6, 1, 6, 3, 6, 141, 8, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 5, 6, 148, 8, 6, 10, 6, 12, 6, 151, 9, 6, 3, 6, 153, 8, 6, 1,
		7, 1, 7, 1, 7, 3, 7, 158, 8, 7, 1, 7, 3, 7, 161, 8, 7, 3, 7, 163, 8, 7,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 3, 9, 170, 8, 9, 1, 9, 1, 9, 3, 9, 174, 8,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 182, 8, 9, 10, 9, 12, 9, 185,
		9, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 192, 8, 10, 1, 10, 1, 10,
		3, 10, 196, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 203, 8, 10,
		10, 10, 12, 10, 206, 9, 10, 1, 10, 1, 10, 1, 10, 3, 10, 211, 8, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 219, 8, 10, 1, 10, 1, 10,
		3, 10, 223, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 230, 8, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 236, 8, 11, 10, 11, 12, 11, 239, 9,
		11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 246, 8, 12, 10, 12, 12, 12,
		249, 9, 12, 3, 12, 251, 8, 12, 1, 12, 1, 12, 3, 12, 255, 8, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 262, 8, 12, 1, 12, 1, 12, 5, 12, 266,
		8, 12, 10, 12, 12, 12, 269, 9, 12, 1, 12, 1, 12, 3, 12, 273, 8, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12,
		285, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 293, 8, 13,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 5, 15, 307, 8, 15, 10, 15, 12, 15, 310, 9, 15, 3, 15, 312, 8,
		15, 1, 15, 1, 15, 1, 16, 1, 16, 3, 16, 318, 8, 16, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 330, 8, 17, 1,
		17, 1, 17, 3, 17, 334, 8, 17, 1, 17, 1, 17, 1, 17, 3, 17, 339, 8, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 3, 17, 345, 8, 17, 5, 17, 347, 8, 17, 10, 17,
		12, 17, 350, 9, 17, 1, 18, 1, 18, 1, 18, 5, 18, 355, 8, 18, 10, 18, 12,
		18, 358, 9, 18, 1, 19, 1, 19, 1, 19, 5, 19, 363, 8, 19, 10, 19, 12, 19,
		366, 9, 19, 1, 20, 1, 20, 1, 20, 5, 20, 371, 8, 20, 10, 20, 12, 20, 374,
		9, 20, 1, 21, 1, 21, 3, 21, 378, 8, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		23, 3, 23, 385, 8, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27,
		1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 409, 8, 29, 1, 30, 1, 30, 1, 31,
		1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 0, 3, 18, 22, 34, 34, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 0, 7, 2, 0, 21, 21, 46,
		46, 2, 0, 23, 23, 27, 27, 2, 0, 20, 20, 28, 28, 1, 0, 8, 14, 2, 0, 5, 5,
		15, 19, 1, 0, 15, 16, 3, 0, 40, 40, 50, 50, 62, 62, 455, 0, 68, 1, 0, 0,
		0, 2, 72, 1, 0, 0, 0, 4, 83, 1, 0, 0, 0, 6, 87, 1, 0, 0, 0, 8, 93, 1, 0,
		0, 0, 10, 100, 1, 0, 0, 0, 12, 106, 1, 0, 0, 0, 14, 162, 1, 0, 0, 0, 16,
		164, 1, 0, 0, 0, 18, 173, 1, 0, 0, 0, 20, 222, 1, 0, 0, 0, 22, 229, 1,
		0, 0, 0, 24, 284, 1, 0, 0, 0, 26, 292, 1, 0, 0, 0, 28, 294, 1, 0, 0, 0,
		30, 299, 1, 0, 0, 0, 32, 315, 1, 0, 0, 0, 34, 329, 1, 0, 0, 0, 36, 351,
		1, 0, 0, 0, 38, 359, 1, 0, 0, 0, 40, 367, 1, 0, 0, 0, 42, 377, 1, 0, 0,
		0, 44, 379, 1, 0, 0, 0, 46, 384, 1, 0, 0, 0, 48, 386, 1, 0, 0, 0, 50, 388,
		1, 0, 0, 0, 52, 390, 1, 0, 0, 0, 54, 392, 1, 0, 0, 0, 56, 394, 1, 0, 0,
		0, 58, 408, 1, 0, 0, 0, 60, 410, 1, 0, 0, 0, 62, 412, 1, 0, 0, 0, 64, 414,
		1, 0, 0, 0, 66, 416, 1, 0, 0, 0, 68, 69, 3, 2, 1, 0, 69, 70, 5, 1, 0, 0,
		70, 71, 5, 0, 0, 1, 71, 1, 1, 0, 0, 0, 72, 73, 3, 4, 2, 0, 73, 3, 1, 0,
		0, 0, 74, 75, 5, 57, 0, 0, 75, 80, 3, 6, 3, 0, 76, 77, 5, 2, 0, 0, 77,
		79, 3, 6, 3, 0, 78, 76, 1, 0, 0, 0, 79, 82, 1, 0, 0, 0, 80, 78, 1, 0, 0,
		0, 80, 81, 1, 0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 83, 74,
		1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 86, 3, 8, 4, 0,
		86, 5, 1, 0, 0, 0, 87, 88, 3, 42, 21, 0, 88, 89, 5, 22, 0, 0, 89, 90, 5,
		3, 0, 0, 90, 91, 3, 2, 1, 0, 91, 92, 5, 4, 0, 0, 92, 7, 1, 0, 0, 0, 93,
		97, 3, 12, 6, 0, 94, 96, 3, 10, 5, 0, 95, 94, 1, 0, 0, 0, 96, 99, 1, 0,
		0, 0, 97, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 9, 1, 0, 0, 0, 99, 97,
		1, 0, 0, 0, 100, 102, 5, 54, 0, 0, 101, 103, 3, 56, 28, 0, 102, 101, 1,
		0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 105, 3, 12, 6,
		0, 105, 11, 1, 0, 0, 0, 106, 108, 5, 51, 0, 0, 107, 109, 3, 56, 28, 0,
		108, 107, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110,
		115, 3, 14, 7, 0, 111, 112, 5, 2, 0, 0, 112, 114, 3, 14, 7, 0, 113, 111,
		1, 0, 0, 0, 114, 117, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0,
		0, 0, 116, 127, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 118, 119, 5, 32, 0, 0,
		119, 124, 3, 34, 17, 0, 120, 121, 5, 2, 0, 0, 121, 123, 3, 34, 17, 0, 122,
		120, 1, 0, 0, 0, 123, 126, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 125,
		1, 0, 0, 0, 125, 128, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 127, 118, 1, 0,
		0, 0, 127, 128, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0, 129, 130, 5, 56, 0, 0,
		130, 132, 3, 18, 9, 0, 131, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132,
		136, 1, 0, 0, 0, 133, 134, 5, 34, 0, 0, 134, 135, 5, 24, 0, 0, 135, 137,
		3, 36, 18, 0, 136, 133, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 140, 1,
		0, 0, 0, 138, 139, 5, 35, 0, 0, 139, 141, 3, 18, 9, 0, 140, 138, 1, 0,
		0, 0, 140, 141, 1, 0, 0, 0, 141, 152, 1, 0, 0, 0, 142, 143, 5, 47, 0, 0,
		143, 144, 5, 24, 0, 0, 144, 149, 3, 32, 16, 0, 145, 146, 5, 2, 0, 0, 146,
		148, 3, 32, 16, 0, 147, 145, 1, 0, 0, 0, 148, 151, 1, 0, 0, 0, 149, 147,
		1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 153, 1, 0, 0, 0, 151, 149, 1, 0,
		0, 0, 152, 142, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 13, 1, 0, 0, 0,
		154, 163, 5, 5, 0, 0, 155, 160, 3, 16, 8, 0, 156, 158, 5, 22, 0, 0, 157,
		156, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 161,
		3, 42, 21, 0, 160, 157, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 163, 1,
		0, 0, 0, 162, 154, 1, 0, 0, 0, 162, 155, 1, 0, 0, 0, 163, 15, 1, 0, 0,
		0, 164, 165, 3, 18, 9, 0, 165, 17, 1, 0, 0, 0, 166, 167, 6, 9, -1, 0, 167,
		169, 3, 22, 11, 0, 168, 170, 3, 20, 10, 0, 169, 168, 1, 0, 0, 0, 169, 170,
		1, 0, 0, 0, 170, 174, 1, 0, 0, 0, 171, 172, 5, 43, 0, 0, 172, 174, 3, 18,
		9, 3, 173, 166, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 174, 183, 1, 0, 0, 0,
		175, 176, 10, 2, 0, 0, 176, 177, 7, 0, 0, 0, 177, 182, 3, 18, 9, 3, 178,
		179, 10, 1, 0, 0, 179, 180, 5, 6, 0, 0, 180, 182, 3, 42, 21, 0, 181, 175,
		1, 0, 0, 0, 181, 178, 1, 0, 0, 0, 182, 185, 1, 0, 0, 0, 183, 181, 1, 0,
		0, 0, 183, 184, 1, 0, 0, 0, 184, 19, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0,
		186, 187, 3, 60, 30, 0, 187, 188, 3, 22, 11, 0, 188, 223, 1, 0, 0, 0, 189,
		191, 5, 38, 0, 0, 190, 192, 5, 43, 0, 0, 191, 190, 1, 0, 0, 0, 191, 192,
		1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 223, 5, 44, 0, 0, 194, 196, 5, 43,
		0, 0, 195, 194, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0,
		197, 198, 5, 36, 0, 0, 198, 199, 5, 3, 0, 0, 199, 204, 3, 16, 8, 0, 200,
		201, 5, 2, 0, 0, 201, 203, 3, 16, 8, 0, 202, 200, 1, 0, 0, 0, 203, 206,
		1, 0, 0, 0, 204, 202, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205, 207, 1, 0,
		0, 0, 206, 204, 1, 0, 0, 0, 207, 208, 5, 4, 0, 0, 208, 223, 1, 0, 0, 0,
		209, 211, 5, 43, 0, 0, 210, 209, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211,
		212, 1, 0, 0, 0, 212, 213, 5, 36, 0, 0, 213, 214, 5, 3, 0, 0, 214, 215,
		3, 2, 1, 0, 215, 216, 5, 4, 0, 0, 216, 223, 1, 0, 0, 0, 217, 219, 5, 43,
		0, 0, 218, 217, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0,
		220, 221, 5, 41, 0, 0, 221, 223, 3, 16, 8, 0, 222, 186, 1, 0, 0, 0, 222,
		189, 1, 0, 0, 0, 222, 195, 1, 0, 0, 0, 222, 210, 1, 0, 0, 0, 222, 218,
		1, 0, 0, 0, 223, 21, 1, 0, 0, 0, 224, 225, 6, 11, -1, 0, 225, 230, 3, 24,
		12, 0, 226, 227, 3, 64, 32, 0, 227, 228, 3, 22, 11, 2, 228, 230, 1, 0,
		0, 0, 229, 224, 1, 0, 0, 0, 229, 226, 1, 0, 0, 0, 230, 237, 1, 0, 0, 0,
		231, 232, 10, 1, 0, 0, 232, 233, 3, 62, 31, 0, 233, 234, 3, 22, 11, 2,
		234, 236, 1, 0, 0, 0, 235, 231, 1, 0, 0, 0, 236, 239, 1, 0, 0, 0, 237,
		235, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 23, 1, 0, 0, 0, 239, 237, 1,
		0, 0, 0, 240, 241, 3, 38, 19, 0, 241, 250, 5, 3, 0, 0, 242, 247, 3, 16,
		8, 0, 243, 244, 5, 2, 0, 0, 244, 246, 3, 16, 8, 0, 245, 243, 1, 0, 0, 0,
		246, 249, 1, 0, 0, 0, 247, 245, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248,
		251, 1, 0, 0, 0, 249, 247, 1, 0, 0, 0, 250, 242, 1, 0, 0, 0, 250, 251,
		1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252, 254, 5, 4, 0, 0, 253, 255, 3, 30,
		15, 0, 254, 253, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 285, 1, 0, 0, 0,
		256, 257, 3, 38, 19, 0, 257, 258, 5, 3, 0, 0, 258, 259, 5, 5, 0, 0, 259,
		261, 5, 4, 0, 0, 260, 262, 3, 30, 15, 0, 261, 260, 1, 0, 0, 0, 261, 262,
		1, 0, 0, 0, 262, 285, 1, 0, 0, 0, 263, 267, 5, 25, 0, 0, 264, 266, 3, 28,
		14, 0, 265, 264, 1, 0, 0, 0, 266, 269, 1, 0, 0, 0, 267, 265, 1, 0, 0, 0,
		267, 268, 1, 0, 0, 0, 268, 272, 1, 0, 0, 0, 269, 267, 1, 0, 0, 0, 270,
		271, 5, 29, 0, 0, 271, 273, 3, 16, 8, 0, 272, 270, 1, 0, 0, 0, 272, 273,
		1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 285, 5, 30, 0, 0, 275, 276, 5, 3,
		0, 0, 276, 277, 3, 2, 1, 0, 277, 278, 5, 4, 0, 0, 278, 285, 1, 0, 0, 0,
		279, 280, 5, 3, 0, 0, 280, 281, 3, 16, 8, 0, 281, 282, 5, 4, 0, 0, 282,
		285, 1, 0, 0, 0, 283, 285, 3, 26, 13, 0, 284, 240, 1, 0, 0, 0, 284, 256,
		1, 0, 0, 0, 284, 263, 1, 0, 0, 0, 284, 275, 1, 0, 0, 0, 284, 279, 1, 0,
		0, 0, 284, 283, 1, 0, 0, 0, 285, 25, 1, 0, 0, 0, 286, 293, 3, 38, 19, 0,
		287, 293, 3, 46, 23, 0, 288, 293, 3, 48, 24, 0, 289, 293, 3, 50, 25, 0,
		290, 293, 3, 52, 26, 0, 291, 293, 3, 54, 27, 0, 292, 286, 1, 0, 0, 0, 292,
		287, 1, 0, 0, 0, 292, 288, 1, 0, 0, 0, 292, 289, 1, 0, 0, 0, 292, 290,
		1, 0, 0, 0, 292, 291, 1, 0, 0, 0, 293, 27, 1, 0, 0, 0, 294, 295, 5, 55,
		0, 0, 295, 296, 3, 16, 8, 0, 296, 297, 5, 52, 0, 0, 297, 298, 3, 16, 8,
		0, 298, 29, 1, 0, 0, 0, 299, 300, 5, 49, 0, 0, 300, 311, 5, 3, 0, 0, 301,
		302, 5, 47, 0, 0, 302, 303, 5, 24, 0, 0, 303, 308, 3, 32, 16, 0, 304, 305,
		5, 2, 0, 0, 305, 307, 3, 32, 16, 0, 306, 304, 1, 0, 0, 0, 307, 310, 1,
		0, 0, 0, 308, 306, 1, 0, 0, 0, 308, 309, 1, 0, 0, 0, 309, 312, 1, 0, 0,
		0, 310, 308, 1, 0, 0, 0, 311, 301, 1, 0, 0, 0, 311, 312, 1, 0, 0, 0, 312,
		313, 1, 0, 0, 0, 313, 314, 5, 4, 0, 0, 314, 31, 1, 0, 0, 0, 315, 317, 3,
		16, 8, 0, 316, 318, 7, 1, 0, 0, 317, 316, 1, 0, 0, 0, 317, 318, 1, 0, 0,
		0, 318, 33, 1, 0, 0, 0, 319, 320, 6, 17, -1, 0, 320, 321, 5, 3, 0, 0, 321,
		322, 3, 2, 1, 0, 322, 323, 5, 4, 0, 0, 323, 330, 1, 0, 0, 0, 324, 325,
		5, 3, 0, 0, 325, 326, 3, 34, 17, 0, 326, 327, 5, 4, 0, 0, 327, 330, 1,
		0, 0, 0, 328, 330, 3, 38, 19, 0, 329, 319, 1, 0, 0, 0, 329, 324, 1, 0,
		0, 0, 329, 328, 1, 0, 0, 0, 330, 348, 1, 0, 0, 0, 331, 333, 10, 5, 0, 0,
		332, 334, 5, 22, 0, 0, 333, 332, 1, 0, 0, 0, 333, 334, 1, 0, 0, 0, 334,
		335, 1, 0, 0, 0, 335, 347, 3, 42, 21, 0, 336, 338, 10, 4, 0, 0, 337, 339,
		3, 58, 29, 0, 338, 337, 1, 0, 0, 0, 338, 339, 1, 0, 0, 0, 339, 340, 1,
		0, 0, 0, 340, 341, 5, 39, 0, 0, 341, 344, 3, 34, 17, 0, 342, 343, 5, 45,
		0, 0, 343, 345, 3, 18, 9, 0, 344, 342, 1, 0, 0, 0, 344, 345, 1, 0, 0, 0,
		345, 347, 1, 0, 0, 0, 346, 331, 1, 0, 0, 0, 346, 336, 1, 0, 0, 0, 347,
		350, 1, 0, 0, 0, 348, 346, 1, 0, 0, 0, 348, 349, 1, 0, 0, 0, 349, 35, 1,
		0, 0, 0, 350, 348, 1, 0, 0, 0, 351, 356, 3, 16, 8, 0, 352, 353, 5, 2, 0,
		0, 353, 355, 3, 16, 8, 0, 354, 352, 1, 0, 0, 0, 355, 358, 1, 0, 0, 0, 356,
		354, 1, 0, 0, 0, 356, 357, 1, 0, 0, 0, 357, 37, 1, 0, 0, 0, 358, 356, 1,
		0, 0, 0, 359, 364, 3, 42, 21, 0, 360, 361, 5, 7, 0, 0, 361, 363, 3, 42,
		21, 0, 362, 360, 1, 0, 0, 0, 363, 366, 1, 0, 0, 0, 364, 362, 1, 0, 0, 0,
		364, 365, 1, 0, 0, 0, 365, 39, 1, 0, 0, 0, 366, 364, 1, 0, 0, 0, 367, 372,
		3, 42, 21, 0, 368, 369, 5, 2, 0, 0, 369, 371, 3, 42, 21, 0, 370, 368, 1,
		0, 0, 0, 371, 374, 1, 0, 0, 0, 372, 370, 1, 0, 0, 0, 372, 373, 1, 0, 0,
		0, 373, 41, 1, 0, 0, 0, 374, 372, 1, 0, 0, 0, 375, 378, 3, 66, 33, 0, 376,
		378, 3, 44, 22, 0, 377, 375, 1, 0, 0, 0, 377, 376, 1, 0, 0, 0, 378, 43,
		1, 0, 0, 0, 379, 380, 5, 63, 0, 0, 380, 45, 1, 0, 0, 0, 381, 385, 5, 59,
		0, 0, 382, 385, 5, 60, 0, 0, 383, 385, 5, 61, 0, 0, 384, 381, 1, 0, 0,
		0, 384, 382, 1, 0, 0, 0, 384, 383, 1, 0, 0, 0, 385, 47, 1, 0, 0, 0, 386,
		387, 5, 58, 0, 0, 387, 49, 1, 0, 0, 0, 388, 389, 5, 44, 0, 0, 389, 51,
		1, 0, 0, 0, 390, 391, 5, 53, 0, 0, 391, 53, 1, 0, 0, 0, 392, 393, 5, 31,
		0, 0, 393, 55, 1, 0, 0, 0, 394, 395, 7, 2, 0, 0, 395, 57, 1, 0, 0, 0, 396,
		409, 5, 37, 0, 0, 397, 409, 5, 40, 0, 0, 398, 399, 5, 40, 0, 0, 399, 409,
		5, 48, 0, 0, 400, 409, 5, 50, 0, 0, 401, 402, 5, 50, 0, 0, 402, 409, 5,
		48, 0, 0, 403, 409, 5, 33, 0, 0, 404, 405, 5, 33, 0, 0, 405, 409, 5, 48,
		0, 0, 406, 409, 5, 26, 0, 0, 407, 409, 5, 42, 0, 0, 408, 396, 1, 0, 0,
		0, 408, 397, 1, 0, 0, 0, 408, 398, 1, 0, 0, 0, 408, 400, 1, 0, 0, 0, 408,
		401, 1, 0, 0, 0, 408, 403, 1, 0, 0, 0, 408, 404, 1, 0, 0, 0, 408, 406,
		1, 0, 0, 0, 408, 407, 1, 0, 0, 0, 409, 59, 1, 0, 0, 0, 410, 411, 7, 3,
		0, 0, 411, 61, 1, 0, 0, 0, 412, 413, 7, 4, 0, 0, 413, 63, 1, 0, 0, 0, 414,
		415, 7, 5, 0, 0, 415, 65, 1, 0, 0, 0, 416, 417, 7, 6, 0, 0, 417, 67, 1,
		0, 0, 0, 51, 80, 83, 97, 102, 108, 115, 124, 127, 131, 136, 140, 149, 152,
		157, 160, 162, 169, 173, 181, 183, 191, 195, 204, 210, 218, 222, 229, 237,
		247, 250, 254, 261, 267, 272, 284, 292, 308, 311, 317, 329, 333, 338, 344,
		346, 348, 356, 364, 372, 377, 384, 408,
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

// SimpleSqlParserInit initializes any static state used to implement SimpleSqlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSimpleSqlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SimpleSqlParserInit() {
	staticData := &simplesqlParserStaticData
	staticData.once.Do(simplesqlParserInit)
}

// NewSimpleSqlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSimpleSqlParser(input antlr.TokenStream) *SimpleSqlParser {
	SimpleSqlParserInit()
	this := new(SimpleSqlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &simplesqlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "SimpleSql.g4"

	return this
}

// SimpleSqlParser tokens.
const (
	SimpleSqlParserEOF               = antlr.TokenEOF
	SimpleSqlParserT__0              = 1
	SimpleSqlParserT__1              = 2
	SimpleSqlParserT__2              = 3
	SimpleSqlParserT__3              = 4
	SimpleSqlParserT__4              = 5
	SimpleSqlParserT__5              = 6
	SimpleSqlParserT__6              = 7
	SimpleSqlParserT__7              = 8
	SimpleSqlParserT__8              = 9
	SimpleSqlParserT__9              = 10
	SimpleSqlParserT__10             = 11
	SimpleSqlParserT__11             = 12
	SimpleSqlParserT__12             = 13
	SimpleSqlParserT__13             = 14
	SimpleSqlParserT__14             = 15
	SimpleSqlParserT__15             = 16
	SimpleSqlParserT__16             = 17
	SimpleSqlParserT__17             = 18
	SimpleSqlParserT__18             = 19
	SimpleSqlParserALL               = 20
	SimpleSqlParserAND               = 21
	SimpleSqlParserAS                = 22
	SimpleSqlParserASC               = 23
	SimpleSqlParserBY                = 24
	SimpleSqlParserCASE              = 25
	SimpleSqlParserCROSS             = 26
	SimpleSqlParserDESC              = 27
	SimpleSqlParserDISTINCT          = 28
	SimpleSqlParserELSE              = 29
	SimpleSqlParserEND               = 30
	SimpleSqlParserFALSE             = 31
	SimpleSqlParserFROM              = 32
	SimpleSqlParserFULL              = 33
	SimpleSqlParserGROUP             = 34
	SimpleSqlParserHAVING            = 35
	SimpleSqlParserIN                = 36
	SimpleSqlParserINNER             = 37
	SimpleSqlParserIS                = 38
	SimpleSqlParserJOIN              = 39
	SimpleSqlParserLEFT              = 40
	SimpleSqlParserLIKE              = 41
	SimpleSqlParserNATURAL           = 42
	SimpleSqlParserNOT               = 43
	SimpleSqlParserNULL              = 44
	SimpleSqlParserON                = 45
	SimpleSqlParserOR                = 46
	SimpleSqlParserORDER             = 47
	SimpleSqlParserOUTER             = 48
	SimpleSqlParserOVER              = 49
	SimpleSqlParserRIGHT             = 50
	SimpleSqlParserSELECT            = 51
	SimpleSqlParserTHEN              = 52
	SimpleSqlParserTRUE              = 53
	SimpleSqlParserUNION             = 54
	SimpleSqlParserWHEN              = 55
	SimpleSqlParserWHERE             = 56
	SimpleSqlParserWITH              = 57
	SimpleSqlParserSTRING            = 58
	SimpleSqlParserINTEGER_VALUE     = 59
	SimpleSqlParserDECIMAL_VALUE     = 60
	SimpleSqlParserFLOAT_VALUE       = 61
	SimpleSqlParserIDENTIFIER        = 62
	SimpleSqlParserQUOTED_IDENTIFIER = 63
	SimpleSqlParserCOMMENT           = 64
	SimpleSqlParserBLOCK_COMMENT     = 65
	SimpleSqlParserWS                = 66
)

// SimpleSqlParser rules.
const (
	SimpleSqlParserRULE_singleStatement    = 0
	SimpleSqlParserRULE_select             = 1
	SimpleSqlParserRULE_cteSelect          = 2
	SimpleSqlParserRULE_cte                = 3
	SimpleSqlParserRULE_unionSelect        = 4
	SimpleSqlParserRULE_unionItem          = 5
	SimpleSqlParserRULE_primarySelect      = 6
	SimpleSqlParserRULE_selectItem         = 7
	SimpleSqlParserRULE_expression         = 8
	SimpleSqlParserRULE_booleanExpression  = 9
	SimpleSqlParserRULE_predicate          = 10
	SimpleSqlParserRULE_valueExpression    = 11
	SimpleSqlParserRULE_primaryExpression  = 12
	SimpleSqlParserRULE_simpleExpression   = 13
	SimpleSqlParserRULE_caseItem           = 14
	SimpleSqlParserRULE_over               = 15
	SimpleSqlParserRULE_sortItem           = 16
	SimpleSqlParserRULE_relation           = 17
	SimpleSqlParserRULE_groupBy            = 18
	SimpleSqlParserRULE_qualifiedName      = 19
	SimpleSqlParserRULE_identifierList     = 20
	SimpleSqlParserRULE_identifier         = 21
	SimpleSqlParserRULE_quotedIdentifier   = 22
	SimpleSqlParserRULE_number             = 23
	SimpleSqlParserRULE_str                = 24
	SimpleSqlParserRULE_null               = 25
	SimpleSqlParserRULE_true               = 26
	SimpleSqlParserRULE_false              = 27
	SimpleSqlParserRULE_setQuantifier      = 28
	SimpleSqlParserRULE_joinType           = 29
	SimpleSqlParserRULE_cmpOp              = 30
	SimpleSqlParserRULE_arithOp            = 31
	SimpleSqlParserRULE_unaryOp            = 32
	SimpleSqlParserRULE_unquotedIdentifier = 33
)

// ISingleStatementContext is an interface to support dynamic dispatch.
type ISingleStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleStatementContext differentiates from other interfaces.
	IsSingleStatementContext()
}

type SingleStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleStatementContext() *SingleStatementContext {
	var p = new(SingleStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_singleStatement
	return p
}

func (*SingleStatementContext) IsSingleStatementContext() {}

func NewSingleStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleStatementContext {
	var p = new(SingleStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_singleStatement

	return p
}

func (s *SingleStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleStatementContext) Select() ISelectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectContext)
}

func (s *SingleStatementContext) EOF() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserEOF, 0)
}

func (s *SingleStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterSingleStatement(s)
	}
}

func (s *SingleStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitSingleStatement(s)
	}
}

func (s *SingleStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitSingleStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) SingleStatement() (localctx ISingleStatementContext) {
	this := p
	_ = this

	localctx = NewSingleStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SimpleSqlParserRULE_singleStatement)

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
		p.Select()
	}
	{
		p.SetState(69)
		p.Match(SimpleSqlParserT__0)
	}
	{
		p.SetState(70)
		p.Match(SimpleSqlParserEOF)
	}

	return localctx
}

// ISelectContext is an interface to support dynamic dispatch.
type ISelectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectContext differentiates from other interfaces.
	IsSelectContext()
}

type SelectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectContext() *SelectContext {
	var p = new(SelectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_select
	return p
}

func (*SelectContext) IsSelectContext() {}

func NewSelectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectContext {
	var p = new(SelectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_select

	return p
}

func (s *SelectContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectContext) CteSelect() ICteSelectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICteSelectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICteSelectContext)
}

func (s *SelectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterSelect(s)
	}
}

func (s *SelectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitSelect(s)
	}
}

func (s *SelectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitSelect(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Select() (localctx ISelectContext) {
	this := p
	_ = this

	localctx = NewSelectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SimpleSqlParserRULE_select)

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
		p.SetState(72)
		p.CteSelect()
	}

	return localctx
}

// ICteSelectContext is an interface to support dynamic dispatch.
type ICteSelectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCteSelectContext differentiates from other interfaces.
	IsCteSelectContext()
}

type CteSelectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCteSelectContext() *CteSelectContext {
	var p = new(CteSelectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_cteSelect
	return p
}

func (*CteSelectContext) IsCteSelectContext() {}

func NewCteSelectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CteSelectContext {
	var p = new(CteSelectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_cteSelect

	return p
}

func (s *CteSelectContext) GetParser() antlr.Parser { return s.parser }

func (s *CteSelectContext) UnionSelect() IUnionSelectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnionSelectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnionSelectContext)
}

func (s *CteSelectContext) WITH() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserWITH, 0)
}

func (s *CteSelectContext) AllCte() []ICteContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICteContext); ok {
			len++
		}
	}

	tst := make([]ICteContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICteContext); ok {
			tst[i] = t.(ICteContext)
			i++
		}
	}

	return tst
}

func (s *CteSelectContext) Cte(i int) ICteContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICteContext); ok {
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

	return t.(ICteContext)
}

func (s *CteSelectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CteSelectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CteSelectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterCteSelect(s)
	}
}

func (s *CteSelectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitCteSelect(s)
	}
}

func (s *CteSelectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitCteSelect(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) CteSelect() (localctx ICteSelectContext) {
	this := p
	_ = this

	localctx = NewCteSelectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SimpleSqlParserRULE_cteSelect)
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
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleSqlParserWITH {
		{
			p.SetState(74)
			p.Match(SimpleSqlParserWITH)
		}
		{
			p.SetState(75)
			p.Cte()
		}
		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SimpleSqlParserT__1 {
			{
				p.SetState(76)
				p.Match(SimpleSqlParserT__1)
			}
			{
				p.SetState(77)
				p.Cte()
			}

			p.SetState(82)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(85)
		p.UnionSelect()
	}

	return localctx
}

// ICteContext is an interface to support dynamic dispatch.
type ICteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCteContext differentiates from other interfaces.
	IsCteContext()
}

type CteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCteContext() *CteContext {
	var p = new(CteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_cte
	return p
}

func (*CteContext) IsCteContext() {}

func NewCteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CteContext {
	var p = new(CteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_cte

	return p
}

func (s *CteContext) GetParser() antlr.Parser { return s.parser }

func (s *CteContext) Identifier() IIdentifierContext {
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

func (s *CteContext) AS() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserAS, 0)
}

func (s *CteContext) Select() ISelectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectContext)
}

func (s *CteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterCte(s)
	}
}

func (s *CteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitCte(s)
	}
}

func (s *CteContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitCte(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Cte() (localctx ICteContext) {
	this := p
	_ = this

	localctx = NewCteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SimpleSqlParserRULE_cte)

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
		p.SetState(87)
		p.Identifier()
	}
	{
		p.SetState(88)
		p.Match(SimpleSqlParserAS)
	}
	{
		p.SetState(89)
		p.Match(SimpleSqlParserT__2)
	}
	{
		p.SetState(90)
		p.Select()
	}
	{
		p.SetState(91)
		p.Match(SimpleSqlParserT__3)
	}

	return localctx
}

// IUnionSelectContext is an interface to support dynamic dispatch.
type IUnionSelectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionSelectContext differentiates from other interfaces.
	IsUnionSelectContext()
}

type UnionSelectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionSelectContext() *UnionSelectContext {
	var p = new(UnionSelectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_unionSelect
	return p
}

func (*UnionSelectContext) IsUnionSelectContext() {}

func NewUnionSelectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionSelectContext {
	var p = new(UnionSelectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_unionSelect

	return p
}

func (s *UnionSelectContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionSelectContext) PrimarySelect() IPrimarySelectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimarySelectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimarySelectContext)
}

func (s *UnionSelectContext) AllUnionItem() []IUnionItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnionItemContext); ok {
			len++
		}
	}

	tst := make([]IUnionItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnionItemContext); ok {
			tst[i] = t.(IUnionItemContext)
			i++
		}
	}

	return tst
}

func (s *UnionSelectContext) UnionItem(i int) IUnionItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnionItemContext); ok {
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

	return t.(IUnionItemContext)
}

func (s *UnionSelectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionSelectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionSelectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterUnionSelect(s)
	}
}

func (s *UnionSelectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitUnionSelect(s)
	}
}

func (s *UnionSelectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitUnionSelect(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) UnionSelect() (localctx IUnionSelectContext) {
	this := p
	_ = this

	localctx = NewUnionSelectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SimpleSqlParserRULE_unionSelect)
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
		p.SetState(93)
		p.PrimarySelect()
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SimpleSqlParserUNION {
		{
			p.SetState(94)
			p.UnionItem()
		}

		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IUnionItemContext is an interface to support dynamic dispatch.
type IUnionItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionItemContext differentiates from other interfaces.
	IsUnionItemContext()
}

type UnionItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionItemContext() *UnionItemContext {
	var p = new(UnionItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_unionItem
	return p
}

func (*UnionItemContext) IsUnionItemContext() {}

func NewUnionItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionItemContext {
	var p = new(UnionItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_unionItem

	return p
}

func (s *UnionItemContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionItemContext) UNION() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserUNION, 0)
}

func (s *UnionItemContext) PrimarySelect() IPrimarySelectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimarySelectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimarySelectContext)
}

func (s *UnionItemContext) SetQuantifier() ISetQuantifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetQuantifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetQuantifierContext)
}

func (s *UnionItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterUnionItem(s)
	}
}

func (s *UnionItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitUnionItem(s)
	}
}

func (s *UnionItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitUnionItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) UnionItem() (localctx IUnionItemContext) {
	this := p
	_ = this

	localctx = NewUnionItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SimpleSqlParserRULE_unionItem)
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
		p.Match(SimpleSqlParserUNION)
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleSqlParserALL || _la == SimpleSqlParserDISTINCT {
		{
			p.SetState(101)
			p.SetQuantifier()
		}

	}
	{
		p.SetState(104)
		p.PrimarySelect()
	}

	return localctx
}

// IPrimarySelectContext is an interface to support dynamic dispatch.
type IPrimarySelectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetWhere returns the where rule contexts.
	GetWhere() IBooleanExpressionContext

	// GetHaving returns the having rule contexts.
	GetHaving() IBooleanExpressionContext

	// SetWhere sets the where rule contexts.
	SetWhere(IBooleanExpressionContext)

	// SetHaving sets the having rule contexts.
	SetHaving(IBooleanExpressionContext)

	// IsPrimarySelectContext differentiates from other interfaces.
	IsPrimarySelectContext()
}

type PrimarySelectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	where  IBooleanExpressionContext
	having IBooleanExpressionContext
}

func NewEmptyPrimarySelectContext() *PrimarySelectContext {
	var p = new(PrimarySelectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_primarySelect
	return p
}

func (*PrimarySelectContext) IsPrimarySelectContext() {}

func NewPrimarySelectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimarySelectContext {
	var p = new(PrimarySelectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_primarySelect

	return p
}

func (s *PrimarySelectContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimarySelectContext) GetWhere() IBooleanExpressionContext { return s.where }

func (s *PrimarySelectContext) GetHaving() IBooleanExpressionContext { return s.having }

func (s *PrimarySelectContext) SetWhere(v IBooleanExpressionContext) { s.where = v }

func (s *PrimarySelectContext) SetHaving(v IBooleanExpressionContext) { s.having = v }

func (s *PrimarySelectContext) SELECT() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserSELECT, 0)
}

func (s *PrimarySelectContext) AllSelectItem() []ISelectItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISelectItemContext); ok {
			len++
		}
	}

	tst := make([]ISelectItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISelectItemContext); ok {
			tst[i] = t.(ISelectItemContext)
			i++
		}
	}

	return tst
}

func (s *PrimarySelectContext) SelectItem(i int) ISelectItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectItemContext); ok {
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

	return t.(ISelectItemContext)
}

func (s *PrimarySelectContext) SetQuantifier() ISetQuantifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetQuantifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetQuantifierContext)
}

func (s *PrimarySelectContext) FROM() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserFROM, 0)
}

func (s *PrimarySelectContext) AllRelation() []IRelationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationContext); ok {
			len++
		}
	}

	tst := make([]IRelationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationContext); ok {
			tst[i] = t.(IRelationContext)
			i++
		}
	}

	return tst
}

func (s *PrimarySelectContext) Relation(i int) IRelationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationContext); ok {
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

	return t.(IRelationContext)
}

func (s *PrimarySelectContext) WHERE() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserWHERE, 0)
}

func (s *PrimarySelectContext) GROUP() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserGROUP, 0)
}

func (s *PrimarySelectContext) AllBY() []antlr.TerminalNode {
	return s.GetTokens(SimpleSqlParserBY)
}

func (s *PrimarySelectContext) BY(i int) antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserBY, i)
}

func (s *PrimarySelectContext) GroupBy() IGroupByContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupByContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupByContext)
}

func (s *PrimarySelectContext) HAVING() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserHAVING, 0)
}

func (s *PrimarySelectContext) ORDER() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserORDER, 0)
}

func (s *PrimarySelectContext) AllSortItem() []ISortItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISortItemContext); ok {
			len++
		}
	}

	tst := make([]ISortItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISortItemContext); ok {
			tst[i] = t.(ISortItemContext)
			i++
		}
	}

	return tst
}

func (s *PrimarySelectContext) SortItem(i int) ISortItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortItemContext); ok {
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

	return t.(ISortItemContext)
}

func (s *PrimarySelectContext) AllBooleanExpression() []IBooleanExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBooleanExpressionContext); ok {
			len++
		}
	}

	tst := make([]IBooleanExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBooleanExpressionContext); ok {
			tst[i] = t.(IBooleanExpressionContext)
			i++
		}
	}

	return tst
}

func (s *PrimarySelectContext) BooleanExpression(i int) IBooleanExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanExpressionContext); ok {
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

	return t.(IBooleanExpressionContext)
}

func (s *PrimarySelectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimarySelectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimarySelectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterPrimarySelect(s)
	}
}

func (s *PrimarySelectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitPrimarySelect(s)
	}
}

func (s *PrimarySelectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitPrimarySelect(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) PrimarySelect() (localctx IPrimarySelectContext) {
	this := p
	_ = this

	localctx = NewPrimarySelectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SimpleSqlParserRULE_primarySelect)
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
		p.SetState(106)
		p.Match(SimpleSqlParserSELECT)
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleSqlParserALL || _la == SimpleSqlParserDISTINCT {
		{
			p.SetState(107)
			p.SetQuantifier()
		}

	}
	{
		p.SetState(110)
		p.SelectItem()
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SimpleSqlParserT__1 {
		{
			p.SetState(111)
			p.Match(SimpleSqlParserT__1)
		}
		{
			p.SetState(112)
			p.SelectItem()
		}

		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleSqlParserFROM {
		{
			p.SetState(118)
			p.Match(SimpleSqlParserFROM)
		}
		{
			p.SetState(119)
			p.relation(0)
		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SimpleSqlParserT__1 {
			{
				p.SetState(120)
				p.Match(SimpleSqlParserT__1)
			}
			{
				p.SetState(121)
				p.relation(0)
			}

			p.SetState(126)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleSqlParserWHERE {
		{
			p.SetState(129)
			p.Match(SimpleSqlParserWHERE)
		}
		{
			p.SetState(130)

			var _x = p.booleanExpression(0)

			localctx.(*PrimarySelectContext).where = _x
		}

	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleSqlParserGROUP {
		{
			p.SetState(133)
			p.Match(SimpleSqlParserGROUP)
		}
		{
			p.SetState(134)
			p.Match(SimpleSqlParserBY)
		}
		{
			p.SetState(135)
			p.GroupBy()
		}

	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleSqlParserHAVING {
		{
			p.SetState(138)
			p.Match(SimpleSqlParserHAVING)
		}
		{
			p.SetState(139)

			var _x = p.booleanExpression(0)

			localctx.(*PrimarySelectContext).having = _x
		}

	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleSqlParserORDER {
		{
			p.SetState(142)
			p.Match(SimpleSqlParserORDER)
		}
		{
			p.SetState(143)
			p.Match(SimpleSqlParserBY)
		}
		{
			p.SetState(144)
			p.SortItem()
		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SimpleSqlParserT__1 {
			{
				p.SetState(145)
				p.Match(SimpleSqlParserT__1)
			}
			{
				p.SetState(146)
				p.SortItem()
			}

			p.SetState(151)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// ISelectItemContext is an interface to support dynamic dispatch.
type ISelectItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectItemContext differentiates from other interfaces.
	IsSelectItemContext()
}

type SelectItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectItemContext() *SelectItemContext {
	var p = new(SelectItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_selectItem
	return p
}

func (*SelectItemContext) IsSelectItemContext() {}

func NewSelectItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectItemContext {
	var p = new(SelectItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_selectItem

	return p
}

func (s *SelectItemContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectItemContext) CopyFrom(ctx *SelectItemContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SelectItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionSelectItemContext struct {
	*SelectItemContext
}

func NewExpressionSelectItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionSelectItemContext {
	var p = new(ExpressionSelectItemContext)

	p.SelectItemContext = NewEmptySelectItemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SelectItemContext))

	return p
}

func (s *ExpressionSelectItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionSelectItemContext) Expression() IExpressionContext {
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

func (s *ExpressionSelectItemContext) Identifier() IIdentifierContext {
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

func (s *ExpressionSelectItemContext) AS() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserAS, 0)
}

func (s *ExpressionSelectItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterExpressionSelectItem(s)
	}
}

func (s *ExpressionSelectItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitExpressionSelectItem(s)
	}
}

func (s *ExpressionSelectItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitExpressionSelectItem(s)

	default:
		return t.VisitChildren(s)
	}
}

type AllSelectItemContext struct {
	*SelectItemContext
}

func NewAllSelectItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllSelectItemContext {
	var p = new(AllSelectItemContext)

	p.SelectItemContext = NewEmptySelectItemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SelectItemContext))

	return p
}

func (s *AllSelectItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllSelectItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterAllSelectItem(s)
	}
}

func (s *AllSelectItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitAllSelectItem(s)
	}
}

func (s *AllSelectItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitAllSelectItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) SelectItem() (localctx ISelectItemContext) {
	this := p
	_ = this

	localctx = NewSelectItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SimpleSqlParserRULE_selectItem)
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

	p.SetState(162)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SimpleSqlParserT__4:
		localctx = NewAllSelectItemContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(154)
			p.Match(SimpleSqlParserT__4)
		}

	case SimpleSqlParserT__2, SimpleSqlParserT__14, SimpleSqlParserT__15, SimpleSqlParserCASE, SimpleSqlParserFALSE, SimpleSqlParserLEFT, SimpleSqlParserNOT, SimpleSqlParserNULL, SimpleSqlParserRIGHT, SimpleSqlParserTRUE, SimpleSqlParserSTRING, SimpleSqlParserINTEGER_VALUE, SimpleSqlParserDECIMAL_VALUE, SimpleSqlParserFLOAT_VALUE, SimpleSqlParserIDENTIFIER, SimpleSqlParserQUOTED_IDENTIFIER:
		localctx = NewExpressionSelectItemContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(155)
			p.Expression()
		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SimpleSqlParserAS || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(SimpleSqlParserLEFT-40))|(1<<(SimpleSqlParserRIGHT-40))|(1<<(SimpleSqlParserIDENTIFIER-40))|(1<<(SimpleSqlParserQUOTED_IDENTIFIER-40)))) != 0) {
			p.SetState(157)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == SimpleSqlParserAS {
				{
					p.SetState(156)
					p.Match(SimpleSqlParserAS)
				}

			}
			{
				p.SetState(159)
				p.Identifier()
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = SimpleSqlParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) BooleanExpression() IBooleanExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SimpleSqlParserRULE_expression)

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
		p.booleanExpression(0)
	}

	return localctx
}

// IBooleanExpressionContext is an interface to support dynamic dispatch.
type IBooleanExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanExpressionContext differentiates from other interfaces.
	IsBooleanExpressionContext()
}

type BooleanExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanExpressionContext() *BooleanExpressionContext {
	var p = new(BooleanExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_booleanExpression
	return p
}

func (*BooleanExpressionContext) IsBooleanExpressionContext() {}

func NewBooleanExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanExpressionContext {
	var p = new(BooleanExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_booleanExpression

	return p
}

func (s *BooleanExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanExpressionContext) CopyFrom(ctx *BooleanExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *BooleanExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BinaryBooleanExpressionContext struct {
	*BooleanExpressionContext
	op antlr.Token
}

func NewBinaryBooleanExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryBooleanExpressionContext {
	var p = new(BinaryBooleanExpressionContext)

	p.BooleanExpressionContext = NewEmptyBooleanExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BooleanExpressionContext))

	return p
}

func (s *BinaryBooleanExpressionContext) GetOp() antlr.Token { return s.op }

func (s *BinaryBooleanExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinaryBooleanExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryBooleanExpressionContext) AllBooleanExpression() []IBooleanExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBooleanExpressionContext); ok {
			len++
		}
	}

	tst := make([]IBooleanExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBooleanExpressionContext); ok {
			tst[i] = t.(IBooleanExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BinaryBooleanExpressionContext) BooleanExpression(i int) IBooleanExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanExpressionContext); ok {
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

	return t.(IBooleanExpressionContext)
}

func (s *BinaryBooleanExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserAND, 0)
}

func (s *BinaryBooleanExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserOR, 0)
}

func (s *BinaryBooleanExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterBinaryBooleanExpression(s)
	}
}

func (s *BinaryBooleanExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitBinaryBooleanExpression(s)
	}
}

func (s *BinaryBooleanExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitBinaryBooleanExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type PredicatedBooleanExpressionContext struct {
	*BooleanExpressionContext
	_valueExpression IValueExpressionContext
}

func NewPredicatedBooleanExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PredicatedBooleanExpressionContext {
	var p = new(PredicatedBooleanExpressionContext)

	p.BooleanExpressionContext = NewEmptyBooleanExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BooleanExpressionContext))

	return p
}

func (s *PredicatedBooleanExpressionContext) Get_valueExpression() IValueExpressionContext {
	return s._valueExpression
}

func (s *PredicatedBooleanExpressionContext) Set_valueExpression(v IValueExpressionContext) {
	s._valueExpression = v
}

func (s *PredicatedBooleanExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicatedBooleanExpressionContext) ValueExpression() IValueExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueExpressionContext)
}

func (s *PredicatedBooleanExpressionContext) Predicate() IPredicateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPredicateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *PredicatedBooleanExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterPredicatedBooleanExpression(s)
	}
}

func (s *PredicatedBooleanExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitPredicatedBooleanExpression(s)
	}
}

func (s *PredicatedBooleanExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitPredicatedBooleanExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryBooleanExpressionContext struct {
	*BooleanExpressionContext
	op antlr.Token
}

func NewUnaryBooleanExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryBooleanExpressionContext {
	var p = new(UnaryBooleanExpressionContext)

	p.BooleanExpressionContext = NewEmptyBooleanExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BooleanExpressionContext))

	return p
}

func (s *UnaryBooleanExpressionContext) GetOp() antlr.Token { return s.op }

func (s *UnaryBooleanExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnaryBooleanExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryBooleanExpressionContext) BooleanExpression() IBooleanExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanExpressionContext)
}

func (s *UnaryBooleanExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserNOT, 0)
}

func (s *UnaryBooleanExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterUnaryBooleanExpression(s)
	}
}

func (s *UnaryBooleanExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitUnaryBooleanExpression(s)
	}
}

func (s *UnaryBooleanExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitUnaryBooleanExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type CastBooleanExpressionContext struct {
	*BooleanExpressionContext
}

func NewCastBooleanExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CastBooleanExpressionContext {
	var p = new(CastBooleanExpressionContext)

	p.BooleanExpressionContext = NewEmptyBooleanExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BooleanExpressionContext))

	return p
}

func (s *CastBooleanExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CastBooleanExpressionContext) BooleanExpression() IBooleanExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanExpressionContext)
}

func (s *CastBooleanExpressionContext) Identifier() IIdentifierContext {
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

func (s *CastBooleanExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterCastBooleanExpression(s)
	}
}

func (s *CastBooleanExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitCastBooleanExpression(s)
	}
}

func (s *CastBooleanExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitCastBooleanExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) BooleanExpression() (localctx IBooleanExpressionContext) {
	return p.booleanExpression(0)
}

func (p *SimpleSqlParser) booleanExpression(_p int) (localctx IBooleanExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBooleanExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBooleanExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, SimpleSqlParserRULE_booleanExpression, _p)
	var _la int

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
	p.SetState(173)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SimpleSqlParserT__2, SimpleSqlParserT__14, SimpleSqlParserT__15, SimpleSqlParserCASE, SimpleSqlParserFALSE, SimpleSqlParserLEFT, SimpleSqlParserNULL, SimpleSqlParserRIGHT, SimpleSqlParserTRUE, SimpleSqlParserSTRING, SimpleSqlParserINTEGER_VALUE, SimpleSqlParserDECIMAL_VALUE, SimpleSqlParserFLOAT_VALUE, SimpleSqlParserIDENTIFIER, SimpleSqlParserQUOTED_IDENTIFIER:
		localctx = NewPredicatedBooleanExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(167)

			var _x = p.valueExpression(0)

			localctx.(*PredicatedBooleanExpressionContext)._valueExpression = _x
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(168)
				p.Predicate(localctx.(*PredicatedBooleanExpressionContext).Get_valueExpression())
			}

		}

	case SimpleSqlParserNOT:
		localctx = NewUnaryBooleanExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(171)

			var _m = p.Match(SimpleSqlParserNOT)

			localctx.(*UnaryBooleanExpressionContext).op = _m
		}
		{
			p.SetState(172)
			p.booleanExpression(3)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(181)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryBooleanExpressionContext(p, NewBooleanExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SimpleSqlParserRULE_booleanExpression)
				p.SetState(175)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(176)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryBooleanExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SimpleSqlParserAND || _la == SimpleSqlParserOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryBooleanExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(177)
					p.booleanExpression(3)
				}

			case 2:
				localctx = NewCastBooleanExpressionContext(p, NewBooleanExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SimpleSqlParserRULE_booleanExpression)
				p.SetState(178)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(179)
					p.Match(SimpleSqlParserT__5)
				}
				{
					p.SetState(180)
					p.Identifier()
				}

			}

		}
		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}

	return localctx
}

// IPredicateContext is an interface to support dynamic dispatch.
type IPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value attribute.
	GetValue() antlr.ParserRuleContext

	// SetValue sets the value attribute.
	SetValue(antlr.ParserRuleContext)

	// IsPredicateContext differentiates from other interfaces.
	IsPredicateContext()
}

type PredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.ParserRuleContext
}

func NewEmptyPredicateContext() *PredicateContext {
	var p = new(PredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_predicate
	return p
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int, value antlr.ParserRuleContext) *PredicateContext {
	var p = new(PredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_predicate

	p.value = value

	return p
}

func (s *PredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateContext) GetValue() antlr.ParserRuleContext { return s.value }

func (s *PredicateContext) SetValue(v antlr.ParserRuleContext) { s.value = v }

func (s *PredicateContext) CopyFrom(ctx *PredicateContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
	s.value = ctx.value
}

func (s *PredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IsNullPredicateContext struct {
	*PredicateContext
}

func NewIsNullPredicateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsNullPredicateContext {
	var p = new(IsNullPredicateContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}

func (s *IsNullPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsNullPredicateContext) IS() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserIS, 0)
}

func (s *IsNullPredicateContext) NULL() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserNULL, 0)
}

func (s *IsNullPredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserNOT, 0)
}

func (s *IsNullPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterIsNullPredicate(s)
	}
}

func (s *IsNullPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitIsNullPredicate(s)
	}
}

func (s *IsNullPredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitIsNullPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}

type CmpPredicateContext struct {
	*PredicateContext
	right IValueExpressionContext
}

func NewCmpPredicateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CmpPredicateContext {
	var p = new(CmpPredicateContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}

func (s *CmpPredicateContext) GetRight() IValueExpressionContext { return s.right }

func (s *CmpPredicateContext) SetRight(v IValueExpressionContext) { s.right = v }

func (s *CmpPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmpPredicateContext) CmpOp() ICmpOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICmpOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICmpOpContext)
}

func (s *CmpPredicateContext) ValueExpression() IValueExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueExpressionContext)
}

func (s *CmpPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterCmpPredicate(s)
	}
}

func (s *CmpPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitCmpPredicate(s)
	}
}

func (s *CmpPredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitCmpPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}

type LikePredicateContext struct {
	*PredicateContext
}

func NewLikePredicateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LikePredicateContext {
	var p = new(LikePredicateContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}

func (s *LikePredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LikePredicateContext) LIKE() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserLIKE, 0)
}

func (s *LikePredicateContext) Expression() IExpressionContext {
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

func (s *LikePredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserNOT, 0)
}

func (s *LikePredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterLikePredicate(s)
	}
}

func (s *LikePredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitLikePredicate(s)
	}
}

func (s *LikePredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitLikePredicate(s)

	default:
		return t.VisitChildren(s)
	}
}

type InSelectPredicateContext struct {
	*PredicateContext
}

func NewInSelectPredicateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InSelectPredicateContext {
	var p = new(InSelectPredicateContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}

func (s *InSelectPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InSelectPredicateContext) IN() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserIN, 0)
}

func (s *InSelectPredicateContext) Select() ISelectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectContext)
}

func (s *InSelectPredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserNOT, 0)
}

func (s *InSelectPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterInSelectPredicate(s)
	}
}

func (s *InSelectPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitInSelectPredicate(s)
	}
}

func (s *InSelectPredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitInSelectPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}

type InListPredicateContext struct {
	*PredicateContext
}

func NewInListPredicateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InListPredicateContext {
	var p = new(InListPredicateContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}

func (s *InListPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InListPredicateContext) IN() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserIN, 0)
}

func (s *InListPredicateContext) AllExpression() []IExpressionContext {
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

func (s *InListPredicateContext) Expression(i int) IExpressionContext {
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

func (s *InListPredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserNOT, 0)
}

func (s *InListPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterInListPredicate(s)
	}
}

func (s *InListPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitInListPredicate(s)
	}
}

func (s *InListPredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitInListPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Predicate(value antlr.ParserRuleContext) (localctx IPredicateContext) {
	this := p
	_ = this

	localctx = NewPredicateContext(p, p.GetParserRuleContext(), p.GetState(), value)
	p.EnterRule(localctx, 20, SimpleSqlParserRULE_predicate)
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		localctx = NewCmpPredicateContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(186)
			p.CmpOp()
		}
		{
			p.SetState(187)

			var _x = p.valueExpression(0)

			localctx.(*CmpPredicateContext).right = _x
		}

	case 2:
		localctx = NewIsNullPredicateContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(189)
			p.Match(SimpleSqlParserIS)
		}
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SimpleSqlParserNOT {
			{
				p.SetState(190)
				p.Match(SimpleSqlParserNOT)
			}

		}
		{
			p.SetState(193)
			p.Match(SimpleSqlParserNULL)
		}

	case 3:
		localctx = NewInListPredicateContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SimpleSqlParserNOT {
			{
				p.SetState(194)
				p.Match(SimpleSqlParserNOT)
			}

		}
		{
			p.SetState(197)
			p.Match(SimpleSqlParserIN)
		}
		{
			p.SetState(198)
			p.Match(SimpleSqlParserT__2)
		}
		{
			p.SetState(199)
			p.Expression()
		}
		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SimpleSqlParserT__1 {
			{
				p.SetState(200)
				p.Match(SimpleSqlParserT__1)
			}
			{
				p.SetState(201)
				p.Expression()
			}

			p.SetState(206)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(207)
			p.Match(SimpleSqlParserT__3)
		}

	case 4:
		localctx = NewInSelectPredicateContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SimpleSqlParserNOT {
			{
				p.SetState(209)
				p.Match(SimpleSqlParserNOT)
			}

		}
		{
			p.SetState(212)
			p.Match(SimpleSqlParserIN)
		}
		{
			p.SetState(213)
			p.Match(SimpleSqlParserT__2)
		}
		{
			p.SetState(214)
			p.Select()
		}
		{
			p.SetState(215)
			p.Match(SimpleSqlParserT__3)
		}

	case 5:
		localctx = NewLikePredicateContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SimpleSqlParserNOT {
			{
				p.SetState(217)
				p.Match(SimpleSqlParserNOT)
			}

		}
		{
			p.SetState(220)
			p.Match(SimpleSqlParserLIKE)
		}
		{
			p.SetState(221)
			p.Expression()
		}

	}

	return localctx
}

// IValueExpressionContext is an interface to support dynamic dispatch.
type IValueExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueExpressionContext differentiates from other interfaces.
	IsValueExpressionContext()
}

type ValueExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueExpressionContext() *ValueExpressionContext {
	var p = new(ValueExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_valueExpression
	return p
}

func (*ValueExpressionContext) IsValueExpressionContext() {}

func NewValueExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueExpressionContext {
	var p = new(ValueExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_valueExpression

	return p
}

func (s *ValueExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueExpressionContext) CopyFrom(ctx *ValueExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ValueExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrimaryValueExpressionContext struct {
	*ValueExpressionContext
}

func NewPrimaryValueExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryValueExpressionContext {
	var p = new(PrimaryValueExpressionContext)

	p.ValueExpressionContext = NewEmptyValueExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueExpressionContext))

	return p
}

func (s *PrimaryValueExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryValueExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *PrimaryValueExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterPrimaryValueExpression(s)
	}
}

func (s *PrimaryValueExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitPrimaryValueExpression(s)
	}
}

func (s *PrimaryValueExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitPrimaryValueExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryValueExpressionContext struct {
	*ValueExpressionContext
	op IUnaryOpContext
}

func NewUnaryValueExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryValueExpressionContext {
	var p = new(UnaryValueExpressionContext)

	p.ValueExpressionContext = NewEmptyValueExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueExpressionContext))

	return p
}

func (s *UnaryValueExpressionContext) GetOp() IUnaryOpContext { return s.op }

func (s *UnaryValueExpressionContext) SetOp(v IUnaryOpContext) { s.op = v }

func (s *UnaryValueExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryValueExpressionContext) ValueExpression() IValueExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueExpressionContext)
}

func (s *UnaryValueExpressionContext) UnaryOp() IUnaryOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryOpContext)
}

func (s *UnaryValueExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterUnaryValueExpression(s)
	}
}

func (s *UnaryValueExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitUnaryValueExpression(s)
	}
}

func (s *UnaryValueExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitUnaryValueExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArithValueExpressionContext struct {
	*ValueExpressionContext
	left  IValueExpressionContext
	op    IArithOpContext
	right IValueExpressionContext
}

func NewArithValueExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArithValueExpressionContext {
	var p = new(ArithValueExpressionContext)

	p.ValueExpressionContext = NewEmptyValueExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueExpressionContext))

	return p
}

func (s *ArithValueExpressionContext) GetLeft() IValueExpressionContext { return s.left }

func (s *ArithValueExpressionContext) GetOp() IArithOpContext { return s.op }

func (s *ArithValueExpressionContext) GetRight() IValueExpressionContext { return s.right }

func (s *ArithValueExpressionContext) SetLeft(v IValueExpressionContext) { s.left = v }

func (s *ArithValueExpressionContext) SetOp(v IArithOpContext) { s.op = v }

func (s *ArithValueExpressionContext) SetRight(v IValueExpressionContext) { s.right = v }

func (s *ArithValueExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithValueExpressionContext) AllValueExpression() []IValueExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueExpressionContext); ok {
			len++
		}
	}

	tst := make([]IValueExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueExpressionContext); ok {
			tst[i] = t.(IValueExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArithValueExpressionContext) ValueExpression(i int) IValueExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExpressionContext); ok {
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

	return t.(IValueExpressionContext)
}

func (s *ArithValueExpressionContext) ArithOp() IArithOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithOpContext)
}

func (s *ArithValueExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterArithValueExpression(s)
	}
}

func (s *ArithValueExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitArithValueExpression(s)
	}
}

func (s *ArithValueExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitArithValueExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) ValueExpression() (localctx IValueExpressionContext) {
	return p.valueExpression(0)
}

func (p *SimpleSqlParser) valueExpression(_p int) (localctx IValueExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewValueExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IValueExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, SimpleSqlParserRULE_valueExpression, _p)

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
	p.SetState(229)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SimpleSqlParserT__2, SimpleSqlParserCASE, SimpleSqlParserFALSE, SimpleSqlParserLEFT, SimpleSqlParserNULL, SimpleSqlParserRIGHT, SimpleSqlParserTRUE, SimpleSqlParserSTRING, SimpleSqlParserINTEGER_VALUE, SimpleSqlParserDECIMAL_VALUE, SimpleSqlParserFLOAT_VALUE, SimpleSqlParserIDENTIFIER, SimpleSqlParserQUOTED_IDENTIFIER:
		localctx = NewPrimaryValueExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(225)
			p.PrimaryExpression()
		}

	case SimpleSqlParserT__14, SimpleSqlParserT__15:
		localctx = NewUnaryValueExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(226)

			var _x = p.UnaryOp()

			localctx.(*UnaryValueExpressionContext).op = _x
		}
		{
			p.SetState(227)
			p.valueExpression(2)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewArithValueExpressionContext(p, NewValueExpressionContext(p, _parentctx, _parentState))
			localctx.(*ArithValueExpressionContext).left = _prevctx

			p.PushNewRecursionContext(localctx, _startState, SimpleSqlParserRULE_valueExpression)
			p.SetState(231)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(232)

				var _x = p.ArithOp()

				localctx.(*ArithValueExpressionContext).op = _x
			}
			{
				p.SetState(233)

				var _x = p.valueExpression(2)

				localctx.(*ArithValueExpressionContext).right = _x
			}

		}
		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_primaryExpression
	return p
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) CopyFrom(ctx *PrimaryExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StarFunctionCallExpressionContext struct {
	*PrimaryExpressionContext
}

func NewStarFunctionCallExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StarFunctionCallExpressionContext {
	var p = new(StarFunctionCallExpressionContext)

	p.PrimaryExpressionContext = NewEmptyPrimaryExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *StarFunctionCallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StarFunctionCallExpressionContext) QualifiedName() IQualifiedNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualifiedNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualifiedNameContext)
}

func (s *StarFunctionCallExpressionContext) Over() IOverContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOverContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOverContext)
}

func (s *StarFunctionCallExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterStarFunctionCallExpression(s)
	}
}

func (s *StarFunctionCallExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitStarFunctionCallExpression(s)
	}
}

func (s *StarFunctionCallExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitStarFunctionCallExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type SimplePrimaryExpressionContext struct {
	*PrimaryExpressionContext
}

func NewSimplePrimaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimplePrimaryExpressionContext {
	var p = new(SimplePrimaryExpressionContext)

	p.PrimaryExpressionContext = NewEmptyPrimaryExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *SimplePrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimplePrimaryExpressionContext) SimpleExpression() ISimpleExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleExpressionContext)
}

func (s *SimplePrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterSimplePrimaryExpression(s)
	}
}

func (s *SimplePrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitSimplePrimaryExpression(s)
	}
}

func (s *SimplePrimaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitSimplePrimaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type CaseExpressionContext struct {
	*PrimaryExpressionContext
}

func NewCaseExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CaseExpressionContext {
	var p = new(CaseExpressionContext)

	p.PrimaryExpressionContext = NewEmptyPrimaryExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *CaseExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseExpressionContext) CASE() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserCASE, 0)
}

func (s *CaseExpressionContext) END() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserEND, 0)
}

func (s *CaseExpressionContext) AllCaseItem() []ICaseItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICaseItemContext); ok {
			len++
		}
	}

	tst := make([]ICaseItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICaseItemContext); ok {
			tst[i] = t.(ICaseItemContext)
			i++
		}
	}

	return tst
}

func (s *CaseExpressionContext) CaseItem(i int) ICaseItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseItemContext); ok {
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

	return t.(ICaseItemContext)
}

func (s *CaseExpressionContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserELSE, 0)
}

func (s *CaseExpressionContext) Expression() IExpressionContext {
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

func (s *CaseExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterCaseExpression(s)
	}
}

func (s *CaseExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitCaseExpression(s)
	}
}

func (s *CaseExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitCaseExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallExpressionContext struct {
	*PrimaryExpressionContext
}

func NewFunctionCallExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallExpressionContext {
	var p = new(FunctionCallExpressionContext)

	p.PrimaryExpressionContext = NewEmptyPrimaryExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *FunctionCallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallExpressionContext) QualifiedName() IQualifiedNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualifiedNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualifiedNameContext)
}

func (s *FunctionCallExpressionContext) AllExpression() []IExpressionContext {
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

func (s *FunctionCallExpressionContext) Expression(i int) IExpressionContext {
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

func (s *FunctionCallExpressionContext) Over() IOverContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOverContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOverContext)
}

func (s *FunctionCallExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterFunctionCallExpression(s)
	}
}

func (s *FunctionCallExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitFunctionCallExpression(s)
	}
}

func (s *FunctionCallExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitFunctionCallExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExpressionContext struct {
	*PrimaryExpressionContext
}

func NewParenExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExpressionContext {
	var p = new(ParenExpressionContext)

	p.PrimaryExpressionContext = NewEmptyPrimaryExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpressionContext))

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
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterParenExpression(s)
	}
}

func (s *ParenExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitParenExpression(s)
	}
}

func (s *ParenExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitParenExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type SelectExpressionContext struct {
	*PrimaryExpressionContext
}

func NewSelectExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectExpressionContext {
	var p = new(SelectExpressionContext)

	p.PrimaryExpressionContext = NewEmptyPrimaryExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *SelectExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectExpressionContext) Select() ISelectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectContext)
}

func (s *SelectExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterSelectExpression(s)
	}
}

func (s *SelectExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitSelectExpression(s)
	}
}

func (s *SelectExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitSelectExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	this := p
	_ = this

	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SimpleSqlParserRULE_primaryExpression)
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

	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFunctionCallExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(240)
			p.QualifiedName()
		}
		{
			p.SetState(241)
			p.Match(SimpleSqlParserT__2)
		}
		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SimpleSqlParserT__2)|(1<<SimpleSqlParserT__14)|(1<<SimpleSqlParserT__15)|(1<<SimpleSqlParserCASE)|(1<<SimpleSqlParserFALSE))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(SimpleSqlParserLEFT-40))|(1<<(SimpleSqlParserNOT-40))|(1<<(SimpleSqlParserNULL-40))|(1<<(SimpleSqlParserRIGHT-40))|(1<<(SimpleSqlParserTRUE-40))|(1<<(SimpleSqlParserSTRING-40))|(1<<(SimpleSqlParserINTEGER_VALUE-40))|(1<<(SimpleSqlParserDECIMAL_VALUE-40))|(1<<(SimpleSqlParserFLOAT_VALUE-40))|(1<<(SimpleSqlParserIDENTIFIER-40))|(1<<(SimpleSqlParserQUOTED_IDENTIFIER-40)))) != 0) {
			{
				p.SetState(242)
				p.Expression()
			}
			p.SetState(247)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == SimpleSqlParserT__1 {
				{
					p.SetState(243)
					p.Match(SimpleSqlParserT__1)
				}
				{
					p.SetState(244)
					p.Expression()
				}

				p.SetState(249)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(252)
			p.Match(SimpleSqlParserT__3)
		}
		p.SetState(254)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(253)
				p.Over()
			}

		}

	case 2:
		localctx = NewStarFunctionCallExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(256)
			p.QualifiedName()
		}
		{
			p.SetState(257)
			p.Match(SimpleSqlParserT__2)
		}
		{
			p.SetState(258)
			p.Match(SimpleSqlParserT__4)
		}
		{
			p.SetState(259)
			p.Match(SimpleSqlParserT__3)
		}
		p.SetState(261)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(260)
				p.Over()
			}

		}

	case 3:
		localctx = NewCaseExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(263)
			p.Match(SimpleSqlParserCASE)
		}
		p.SetState(267)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SimpleSqlParserWHEN {
			{
				p.SetState(264)
				p.CaseItem()
			}

			p.SetState(269)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SimpleSqlParserELSE {
			{
				p.SetState(270)
				p.Match(SimpleSqlParserELSE)
			}
			{
				p.SetState(271)
				p.Expression()
			}

		}
		{
			p.SetState(274)
			p.Match(SimpleSqlParserEND)
		}

	case 4:
		localctx = NewSelectExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(275)
			p.Match(SimpleSqlParserT__2)
		}
		{
			p.SetState(276)
			p.Select()
		}
		{
			p.SetState(277)
			p.Match(SimpleSqlParserT__3)
		}

	case 5:
		localctx = NewParenExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(279)
			p.Match(SimpleSqlParserT__2)
		}
		{
			p.SetState(280)
			p.Expression()
		}
		{
			p.SetState(281)
			p.Match(SimpleSqlParserT__3)
		}

	case 6:
		localctx = NewSimplePrimaryExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(283)
			p.SimpleExpression()
		}

	}

	return localctx
}

// ISimpleExpressionContext is an interface to support dynamic dispatch.
type ISimpleExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimpleExpressionContext differentiates from other interfaces.
	IsSimpleExpressionContext()
}

type SimpleExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleExpressionContext() *SimpleExpressionContext {
	var p = new(SimpleExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_simpleExpression
	return p
}

func (*SimpleExpressionContext) IsSimpleExpressionContext() {}

func NewSimpleExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleExpressionContext {
	var p = new(SimpleExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_simpleExpression

	return p
}

func (s *SimpleExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleExpressionContext) QualifiedName() IQualifiedNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualifiedNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualifiedNameContext)
}

func (s *SimpleExpressionContext) Number() INumberContext {
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

func (s *SimpleExpressionContext) Str() IStrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrContext)
}

func (s *SimpleExpressionContext) Null() INullContext {
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

func (s *SimpleExpressionContext) True() ITrueContext {
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

func (s *SimpleExpressionContext) False() IFalseContext {
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

func (s *SimpleExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterSimpleExpression(s)
	}
}

func (s *SimpleExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitSimpleExpression(s)
	}
}

func (s *SimpleExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitSimpleExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) SimpleExpression() (localctx ISimpleExpressionContext) {
	this := p
	_ = this

	localctx = NewSimpleExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SimpleSqlParserRULE_simpleExpression)

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

	p.SetState(292)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SimpleSqlParserLEFT, SimpleSqlParserRIGHT, SimpleSqlParserIDENTIFIER, SimpleSqlParserQUOTED_IDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(286)
			p.QualifiedName()
		}

	case SimpleSqlParserINTEGER_VALUE, SimpleSqlParserDECIMAL_VALUE, SimpleSqlParserFLOAT_VALUE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(287)
			p.Number()
		}

	case SimpleSqlParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(288)
			p.Str()
		}

	case SimpleSqlParserNULL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(289)
			p.Null()
		}

	case SimpleSqlParserTRUE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(290)
			p.True()
		}

	case SimpleSqlParserFALSE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(291)
			p.False()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICaseItemContext is an interface to support dynamic dispatch.
type ICaseItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCaseItemContext differentiates from other interfaces.
	IsCaseItemContext()
}

type CaseItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseItemContext() *CaseItemContext {
	var p = new(CaseItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_caseItem
	return p
}

func (*CaseItemContext) IsCaseItemContext() {}

func NewCaseItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseItemContext {
	var p = new(CaseItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_caseItem

	return p
}

func (s *CaseItemContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseItemContext) WHEN() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserWHEN, 0)
}

func (s *CaseItemContext) AllExpression() []IExpressionContext {
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

func (s *CaseItemContext) Expression(i int) IExpressionContext {
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

func (s *CaseItemContext) THEN() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserTHEN, 0)
}

func (s *CaseItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterCaseItem(s)
	}
}

func (s *CaseItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitCaseItem(s)
	}
}

func (s *CaseItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitCaseItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) CaseItem() (localctx ICaseItemContext) {
	this := p
	_ = this

	localctx = NewCaseItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SimpleSqlParserRULE_caseItem)

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
		p.SetState(294)
		p.Match(SimpleSqlParserWHEN)
	}
	{
		p.SetState(295)
		p.Expression()
	}
	{
		p.SetState(296)
		p.Match(SimpleSqlParserTHEN)
	}
	{
		p.SetState(297)
		p.Expression()
	}

	return localctx
}

// IOverContext is an interface to support dynamic dispatch.
type IOverContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOverContext differentiates from other interfaces.
	IsOverContext()
}

type OverContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOverContext() *OverContext {
	var p = new(OverContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_over
	return p
}

func (*OverContext) IsOverContext() {}

func NewOverContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OverContext {
	var p = new(OverContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_over

	return p
}

func (s *OverContext) GetParser() antlr.Parser { return s.parser }

func (s *OverContext) OVER() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserOVER, 0)
}

func (s *OverContext) ORDER() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserORDER, 0)
}

func (s *OverContext) BY() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserBY, 0)
}

func (s *OverContext) AllSortItem() []ISortItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISortItemContext); ok {
			len++
		}
	}

	tst := make([]ISortItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISortItemContext); ok {
			tst[i] = t.(ISortItemContext)
			i++
		}
	}

	return tst
}

func (s *OverContext) SortItem(i int) ISortItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortItemContext); ok {
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

	return t.(ISortItemContext)
}

func (s *OverContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OverContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OverContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterOver(s)
	}
}

func (s *OverContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitOver(s)
	}
}

func (s *OverContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitOver(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Over() (localctx IOverContext) {
	this := p
	_ = this

	localctx = NewOverContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SimpleSqlParserRULE_over)
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
		p.SetState(299)
		p.Match(SimpleSqlParserOVER)
	}
	{
		p.SetState(300)
		p.Match(SimpleSqlParserT__2)
	}
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleSqlParserORDER {
		{
			p.SetState(301)
			p.Match(SimpleSqlParserORDER)
		}
		{
			p.SetState(302)
			p.Match(SimpleSqlParserBY)
		}
		{
			p.SetState(303)
			p.SortItem()
		}
		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SimpleSqlParserT__1 {
			{
				p.SetState(304)
				p.Match(SimpleSqlParserT__1)
			}
			{
				p.SetState(305)
				p.SortItem()
			}

			p.SetState(310)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(313)
		p.Match(SimpleSqlParserT__3)
	}

	return localctx
}

// ISortItemContext is an interface to support dynamic dispatch.
type ISortItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDirection returns the direction token.
	GetDirection() antlr.Token

	// SetDirection sets the direction token.
	SetDirection(antlr.Token)

	// IsSortItemContext differentiates from other interfaces.
	IsSortItemContext()
}

type SortItemContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	direction antlr.Token
}

func NewEmptySortItemContext() *SortItemContext {
	var p = new(SortItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_sortItem
	return p
}

func (*SortItemContext) IsSortItemContext() {}

func NewSortItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortItemContext {
	var p = new(SortItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_sortItem

	return p
}

func (s *SortItemContext) GetParser() antlr.Parser { return s.parser }

func (s *SortItemContext) GetDirection() antlr.Token { return s.direction }

func (s *SortItemContext) SetDirection(v antlr.Token) { s.direction = v }

func (s *SortItemContext) Expression() IExpressionContext {
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

func (s *SortItemContext) ASC() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserASC, 0)
}

func (s *SortItemContext) DESC() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserDESC, 0)
}

func (s *SortItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SortItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterSortItem(s)
	}
}

func (s *SortItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitSortItem(s)
	}
}

func (s *SortItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitSortItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) SortItem() (localctx ISortItemContext) {
	this := p
	_ = this

	localctx = NewSortItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SimpleSqlParserRULE_sortItem)
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
		p.SetState(315)
		p.Expression()
	}
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleSqlParserASC || _la == SimpleSqlParserDESC {
		{
			p.SetState(316)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*SortItemContext).direction = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SimpleSqlParserASC || _la == SimpleSqlParserDESC) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*SortItemContext).direction = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}

// IRelationContext is an interface to support dynamic dispatch.
type IRelationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationContext differentiates from other interfaces.
	IsRelationContext()
}

type RelationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationContext() *RelationContext {
	var p = new(RelationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_relation
	return p
}

func (*RelationContext) IsRelationContext() {}

func NewRelationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationContext {
	var p = new(RelationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_relation

	return p
}

func (s *RelationContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationContext) CopyFrom(ctx *RelationContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *RelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AliasedRelationContext struct {
	*RelationContext
}

func NewAliasedRelationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AliasedRelationContext {
	var p = new(AliasedRelationContext)

	p.RelationContext = NewEmptyRelationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RelationContext))

	return p
}

func (s *AliasedRelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasedRelationContext) Relation() IRelationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationContext)
}

func (s *AliasedRelationContext) Identifier() IIdentifierContext {
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

func (s *AliasedRelationContext) AS() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserAS, 0)
}

func (s *AliasedRelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterAliasedRelation(s)
	}
}

func (s *AliasedRelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitAliasedRelation(s)
	}
}

func (s *AliasedRelationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitAliasedRelation(s)

	default:
		return t.VisitChildren(s)
	}
}

type JoinRelationContext struct {
	*RelationContext
	left  IRelationContext
	ty    IJoinTypeContext
	right IRelationContext
	cond  IBooleanExpressionContext
}

func NewJoinRelationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JoinRelationContext {
	var p = new(JoinRelationContext)

	p.RelationContext = NewEmptyRelationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RelationContext))

	return p
}

func (s *JoinRelationContext) GetLeft() IRelationContext { return s.left }

func (s *JoinRelationContext) GetTy() IJoinTypeContext { return s.ty }

func (s *JoinRelationContext) GetRight() IRelationContext { return s.right }

func (s *JoinRelationContext) GetCond() IBooleanExpressionContext { return s.cond }

func (s *JoinRelationContext) SetLeft(v IRelationContext) { s.left = v }

func (s *JoinRelationContext) SetTy(v IJoinTypeContext) { s.ty = v }

func (s *JoinRelationContext) SetRight(v IRelationContext) { s.right = v }

func (s *JoinRelationContext) SetCond(v IBooleanExpressionContext) { s.cond = v }

func (s *JoinRelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoinRelationContext) JOIN() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserJOIN, 0)
}

func (s *JoinRelationContext) AllRelation() []IRelationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationContext); ok {
			len++
		}
	}

	tst := make([]IRelationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationContext); ok {
			tst[i] = t.(IRelationContext)
			i++
		}
	}

	return tst
}

func (s *JoinRelationContext) Relation(i int) IRelationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationContext); ok {
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

	return t.(IRelationContext)
}

func (s *JoinRelationContext) ON() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserON, 0)
}

func (s *JoinRelationContext) JoinType() IJoinTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJoinTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJoinTypeContext)
}

func (s *JoinRelationContext) BooleanExpression() IBooleanExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanExpressionContext)
}

func (s *JoinRelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterJoinRelation(s)
	}
}

func (s *JoinRelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitJoinRelation(s)
	}
}

func (s *JoinRelationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitJoinRelation(s)

	default:
		return t.VisitChildren(s)
	}
}

type SelectRelationContext struct {
	*RelationContext
}

func NewSelectRelationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectRelationContext {
	var p = new(SelectRelationContext)

	p.RelationContext = NewEmptyRelationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RelationContext))

	return p
}

func (s *SelectRelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectRelationContext) Select() ISelectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectContext)
}

func (s *SelectRelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterSelectRelation(s)
	}
}

func (s *SelectRelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitSelectRelation(s)
	}
}

func (s *SelectRelationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitSelectRelation(s)

	default:
		return t.VisitChildren(s)
	}
}

type TableRelationContext struct {
	*RelationContext
}

func NewTableRelationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TableRelationContext {
	var p = new(TableRelationContext)

	p.RelationContext = NewEmptyRelationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RelationContext))

	return p
}

func (s *TableRelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableRelationContext) QualifiedName() IQualifiedNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualifiedNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualifiedNameContext)
}

func (s *TableRelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterTableRelation(s)
	}
}

func (s *TableRelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitTableRelation(s)
	}
}

func (s *TableRelationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitTableRelation(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenRelationContext struct {
	*RelationContext
}

func NewParenRelationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenRelationContext {
	var p = new(ParenRelationContext)

	p.RelationContext = NewEmptyRelationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RelationContext))

	return p
}

func (s *ParenRelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenRelationContext) Relation() IRelationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationContext)
}

func (s *ParenRelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterParenRelation(s)
	}
}

func (s *ParenRelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitParenRelation(s)
	}
}

func (s *ParenRelationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitParenRelation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Relation() (localctx IRelationContext) {
	return p.relation(0)
}

func (p *SimpleSqlParser) relation(_p int) (localctx IRelationContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewRelationContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IRelationContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, SimpleSqlParserRULE_relation, _p)
	var _la int

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
	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSelectRelationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(320)
			p.Match(SimpleSqlParserT__2)
		}
		{
			p.SetState(321)
			p.Select()
		}
		{
			p.SetState(322)
			p.Match(SimpleSqlParserT__3)
		}

	case 2:
		localctx = NewParenRelationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(324)
			p.Match(SimpleSqlParserT__2)
		}
		{
			p.SetState(325)
			p.relation(0)
		}
		{
			p.SetState(326)
			p.Match(SimpleSqlParserT__3)
		}

	case 3:
		localctx = NewTableRelationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(328)
			p.QualifiedName()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(346)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAliasedRelationContext(p, NewRelationContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, SimpleSqlParserRULE_relation)
				p.SetState(331)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				p.SetState(333)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == SimpleSqlParserAS {
					{
						p.SetState(332)
						p.Match(SimpleSqlParserAS)
					}

				}
				{
					p.SetState(335)
					p.Identifier()
				}

			case 2:
				localctx = NewJoinRelationContext(p, NewRelationContext(p, _parentctx, _parentState))
				localctx.(*JoinRelationContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, SimpleSqlParserRULE_relation)
				p.SetState(336)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				p.SetState(338)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if ((_la-26)&-(0x1f+1)) == 0 && ((1<<uint((_la-26)))&((1<<(SimpleSqlParserCROSS-26))|(1<<(SimpleSqlParserFULL-26))|(1<<(SimpleSqlParserINNER-26))|(1<<(SimpleSqlParserLEFT-26))|(1<<(SimpleSqlParserNATURAL-26))|(1<<(SimpleSqlParserRIGHT-26)))) != 0 {
					{
						p.SetState(337)

						var _x = p.JoinType()

						localctx.(*JoinRelationContext).ty = _x
					}

				}
				{
					p.SetState(340)
					p.Match(SimpleSqlParserJOIN)
				}
				{
					p.SetState(341)

					var _x = p.relation(0)

					localctx.(*JoinRelationContext).right = _x
				}
				p.SetState(344)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(342)
						p.Match(SimpleSqlParserON)
					}
					{
						p.SetState(343)

						var _x = p.booleanExpression(0)

						localctx.(*JoinRelationContext).cond = _x
					}

				}

			}

		}
		p.SetState(350)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())
	}

	return localctx
}

// IGroupByContext is an interface to support dynamic dispatch.
type IGroupByContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupByContext differentiates from other interfaces.
	IsGroupByContext()
}

type GroupByContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupByContext() *GroupByContext {
	var p = new(GroupByContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_groupBy
	return p
}

func (*GroupByContext) IsGroupByContext() {}

func NewGroupByContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupByContext {
	var p = new(GroupByContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_groupBy

	return p
}

func (s *GroupByContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupByContext) AllExpression() []IExpressionContext {
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

func (s *GroupByContext) Expression(i int) IExpressionContext {
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

func (s *GroupByContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupByContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupByContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterGroupBy(s)
	}
}

func (s *GroupByContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitGroupBy(s)
	}
}

func (s *GroupByContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitGroupBy(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) GroupBy() (localctx IGroupByContext) {
	this := p
	_ = this

	localctx = NewGroupByContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SimpleSqlParserRULE_groupBy)
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
		p.SetState(351)
		p.Expression()
	}
	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SimpleSqlParserT__1 {
		{
			p.SetState(352)
			p.Match(SimpleSqlParserT__1)
		}
		{
			p.SetState(353)
			p.Expression()
		}

		p.SetState(358)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IQualifiedNameContext is an interface to support dynamic dispatch.
type IQualifiedNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualifiedNameContext differentiates from other interfaces.
	IsQualifiedNameContext()
}

type QualifiedNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifiedNameContext() *QualifiedNameContext {
	var p = new(QualifiedNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_qualifiedName
	return p
}

func (*QualifiedNameContext) IsQualifiedNameContext() {}

func NewQualifiedNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedNameContext {
	var p = new(QualifiedNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_qualifiedName

	return p
}

func (s *QualifiedNameContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedNameContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *QualifiedNameContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
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

	return t.(IIdentifierContext)
}

func (s *QualifiedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterQualifiedName(s)
	}
}

func (s *QualifiedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitQualifiedName(s)
	}
}

func (s *QualifiedNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitQualifiedName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) QualifiedName() (localctx IQualifiedNameContext) {
	this := p
	_ = this

	localctx = NewQualifiedNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SimpleSqlParserRULE_qualifiedName)

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
		p.SetState(359)
		p.Identifier()
	}
	p.SetState(364)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(360)
				p.Match(SimpleSqlParserT__6)
			}
			{
				p.SetState(361)
				p.Identifier()
			}

		}
		p.SetState(366)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())
	}

	return localctx
}

// IIdentifierListContext is an interface to support dynamic dispatch.
type IIdentifierListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierListContext differentiates from other interfaces.
	IsIdentifierListContext()
}

type IdentifierListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierListContext() *IdentifierListContext {
	var p = new(IdentifierListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_identifierList
	return p
}

func (*IdentifierListContext) IsIdentifierListContext() {}

func NewIdentifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierListContext {
	var p = new(IdentifierListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_identifierList

	return p
}

func (s *IdentifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierListContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *IdentifierListContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
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

	return t.(IIdentifierContext)
}

func (s *IdentifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterIdentifierList(s)
	}
}

func (s *IdentifierListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitIdentifierList(s)
	}
}

func (s *IdentifierListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitIdentifierList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) IdentifierList() (localctx IIdentifierListContext) {
	this := p
	_ = this

	localctx = NewIdentifierListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SimpleSqlParserRULE_identifierList)
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
		p.SetState(367)
		p.Identifier()
	}
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SimpleSqlParserT__1 {
		{
			p.SetState(368)
			p.Match(SimpleSqlParserT__1)
		}
		{
			p.SetState(369)
			p.Identifier()
		}

		p.SetState(374)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = SimpleSqlParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) UnquotedIdentifier() IUnquotedIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnquotedIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnquotedIdentifierContext)
}

func (s *IdentifierContext) QuotedIdentifier() IQuotedIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQuotedIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQuotedIdentifierContext)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Identifier() (localctx IIdentifierContext) {
	this := p
	_ = this

	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SimpleSqlParserRULE_identifier)

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

	p.SetState(377)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SimpleSqlParserLEFT, SimpleSqlParserRIGHT, SimpleSqlParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(375)
			p.UnquotedIdentifier()
		}

	case SimpleSqlParserQUOTED_IDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(376)
			p.QuotedIdentifier()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IQuotedIdentifierContext is an interface to support dynamic dispatch.
type IQuotedIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuotedIdentifierContext differentiates from other interfaces.
	IsQuotedIdentifierContext()
}

type QuotedIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuotedIdentifierContext() *QuotedIdentifierContext {
	var p = new(QuotedIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_quotedIdentifier
	return p
}

func (*QuotedIdentifierContext) IsQuotedIdentifierContext() {}

func NewQuotedIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuotedIdentifierContext {
	var p = new(QuotedIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_quotedIdentifier

	return p
}

func (s *QuotedIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *QuotedIdentifierContext) QUOTED_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserQUOTED_IDENTIFIER, 0)
}

func (s *QuotedIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuotedIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterQuotedIdentifier(s)
	}
}

func (s *QuotedIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitQuotedIdentifier(s)
	}
}

func (s *QuotedIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitQuotedIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) QuotedIdentifier() (localctx IQuotedIdentifierContext) {
	this := p
	_ = this

	localctx = NewQuotedIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SimpleSqlParserRULE_quotedIdentifier)

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
		p.SetState(379)
		p.Match(SimpleSqlParserQUOTED_IDENTIFIER)
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
	p.RuleIndex = SimpleSqlParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) CopyFrom(ctx *NumberContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DecimalNumberContext struct {
	*NumberContext
}

func NewDecimalNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecimalNumberContext {
	var p = new(DecimalNumberContext)

	p.NumberContext = NewEmptyNumberContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumberContext))

	return p
}

func (s *DecimalNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalNumberContext) DECIMAL_VALUE() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserDECIMAL_VALUE, 0)
}

func (s *DecimalNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterDecimalNumber(s)
	}
}

func (s *DecimalNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitDecimalNumber(s)
	}
}

func (s *DecimalNumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitDecimalNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntegerNumberContext struct {
	*NumberContext
}

func NewIntegerNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerNumberContext {
	var p = new(IntegerNumberContext)

	p.NumberContext = NewEmptyNumberContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumberContext))

	return p
}

func (s *IntegerNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerNumberContext) INTEGER_VALUE() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserINTEGER_VALUE, 0)
}

func (s *IntegerNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterIntegerNumber(s)
	}
}

func (s *IntegerNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitIntegerNumber(s)
	}
}

func (s *IntegerNumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitIntegerNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatNumberContext struct {
	*NumberContext
}

func NewFloatNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatNumberContext {
	var p = new(FloatNumberContext)

	p.NumberContext = NewEmptyNumberContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumberContext))

	return p
}

func (s *FloatNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatNumberContext) FLOAT_VALUE() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserFLOAT_VALUE, 0)
}

func (s *FloatNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterFloatNumber(s)
	}
}

func (s *FloatNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitFloatNumber(s)
	}
}

func (s *FloatNumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitFloatNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Number() (localctx INumberContext) {
	this := p
	_ = this

	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SimpleSqlParserRULE_number)

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

	p.SetState(384)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SimpleSqlParserINTEGER_VALUE:
		localctx = NewIntegerNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(381)
			p.Match(SimpleSqlParserINTEGER_VALUE)
		}

	case SimpleSqlParserDECIMAL_VALUE:
		localctx = NewDecimalNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(382)
			p.Match(SimpleSqlParserDECIMAL_VALUE)
		}

	case SimpleSqlParserFLOAT_VALUE:
		localctx = NewFloatNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(383)
			p.Match(SimpleSqlParserFLOAT_VALUE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStrContext is an interface to support dynamic dispatch.
type IStrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStrContext differentiates from other interfaces.
	IsStrContext()
}

type StrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrContext() *StrContext {
	var p = new(StrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_str
	return p
}

func (*StrContext) IsStrContext() {}

func NewStrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrContext {
	var p = new(StrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_str

	return p
}

func (s *StrContext) GetParser() antlr.Parser { return s.parser }

func (s *StrContext) STRING() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserSTRING, 0)
}

func (s *StrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterStr(s)
	}
}

func (s *StrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitStr(s)
	}
}

func (s *StrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitStr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Str() (localctx IStrContext) {
	this := p
	_ = this

	localctx = NewStrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SimpleSqlParserRULE_str)

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
		p.SetState(386)
		p.Match(SimpleSqlParserSTRING)
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
	p.RuleIndex = SimpleSqlParserRULE_null
	return p
}

func (*NullContext) IsNullContext() {}

func NewNullContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NullContext {
	var p = new(NullContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_null

	return p
}

func (s *NullContext) GetParser() antlr.Parser { return s.parser }

func (s *NullContext) NULL() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserNULL, 0)
}

func (s *NullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterNull(s)
	}
}

func (s *NullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitNull(s)
	}
}

func (s *NullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitNull(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) Null() (localctx INullContext) {
	this := p
	_ = this

	localctx = NewNullContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SimpleSqlParserRULE_null)

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
		p.SetState(388)
		p.Match(SimpleSqlParserNULL)
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
	p.RuleIndex = SimpleSqlParserRULE_true
	return p
}

func (*TrueContext) IsTrueContext() {}

func NewTrueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TrueContext {
	var p = new(TrueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_true

	return p
}

func (s *TrueContext) GetParser() antlr.Parser { return s.parser }

func (s *TrueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserTRUE, 0)
}

func (s *TrueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TrueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterTrue(s)
	}
}

func (s *TrueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitTrue(s)
	}
}

func (s *TrueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitTrue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) True() (localctx ITrueContext) {
	this := p
	_ = this

	localctx = NewTrueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SimpleSqlParserRULE_true)

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
		p.SetState(390)
		p.Match(SimpleSqlParserTRUE)
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
	p.RuleIndex = SimpleSqlParserRULE_false
	return p
}

func (*FalseContext) IsFalseContext() {}

func NewFalseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FalseContext {
	var p = new(FalseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_false

	return p
}

func (s *FalseContext) GetParser() antlr.Parser { return s.parser }

func (s *FalseContext) FALSE() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserFALSE, 0)
}

func (s *FalseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FalseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FalseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterFalse(s)
	}
}

func (s *FalseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitFalse(s)
	}
}

func (s *FalseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitFalse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) False() (localctx IFalseContext) {
	this := p
	_ = this

	localctx = NewFalseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SimpleSqlParserRULE_false)

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
		p.SetState(392)
		p.Match(SimpleSqlParserFALSE)
	}

	return localctx
}

// ISetQuantifierContext is an interface to support dynamic dispatch.
type ISetQuantifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetQuantifierContext differentiates from other interfaces.
	IsSetQuantifierContext()
}

type SetQuantifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetQuantifierContext() *SetQuantifierContext {
	var p = new(SetQuantifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_setQuantifier
	return p
}

func (*SetQuantifierContext) IsSetQuantifierContext() {}

func NewSetQuantifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetQuantifierContext {
	var p = new(SetQuantifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_setQuantifier

	return p
}

func (s *SetQuantifierContext) GetParser() antlr.Parser { return s.parser }

func (s *SetQuantifierContext) DISTINCT() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserDISTINCT, 0)
}

func (s *SetQuantifierContext) ALL() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserALL, 0)
}

func (s *SetQuantifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetQuantifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetQuantifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterSetQuantifier(s)
	}
}

func (s *SetQuantifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitSetQuantifier(s)
	}
}

func (s *SetQuantifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitSetQuantifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) SetQuantifier() (localctx ISetQuantifierContext) {
	this := p
	_ = this

	localctx = NewSetQuantifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SimpleSqlParserRULE_setQuantifier)
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
		p.SetState(394)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SimpleSqlParserALL || _la == SimpleSqlParserDISTINCT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IJoinTypeContext is an interface to support dynamic dispatch.
type IJoinTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJoinTypeContext differentiates from other interfaces.
	IsJoinTypeContext()
}

type JoinTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoinTypeContext() *JoinTypeContext {
	var p = new(JoinTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_joinType
	return p
}

func (*JoinTypeContext) IsJoinTypeContext() {}

func NewJoinTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JoinTypeContext {
	var p = new(JoinTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_joinType

	return p
}

func (s *JoinTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *JoinTypeContext) INNER() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserINNER, 0)
}

func (s *JoinTypeContext) LEFT() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserLEFT, 0)
}

func (s *JoinTypeContext) OUTER() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserOUTER, 0)
}

func (s *JoinTypeContext) RIGHT() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserRIGHT, 0)
}

func (s *JoinTypeContext) FULL() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserFULL, 0)
}

func (s *JoinTypeContext) CROSS() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserCROSS, 0)
}

func (s *JoinTypeContext) NATURAL() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserNATURAL, 0)
}

func (s *JoinTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoinTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JoinTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterJoinType(s)
	}
}

func (s *JoinTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitJoinType(s)
	}
}

func (s *JoinTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitJoinType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) JoinType() (localctx IJoinTypeContext) {
	this := p
	_ = this

	localctx = NewJoinTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SimpleSqlParserRULE_joinType)

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

	p.SetState(408)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(396)
			p.Match(SimpleSqlParserINNER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(397)
			p.Match(SimpleSqlParserLEFT)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(398)
			p.Match(SimpleSqlParserLEFT)
		}
		{
			p.SetState(399)
			p.Match(SimpleSqlParserOUTER)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(400)
			p.Match(SimpleSqlParserRIGHT)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(401)
			p.Match(SimpleSqlParserRIGHT)
		}
		{
			p.SetState(402)
			p.Match(SimpleSqlParserOUTER)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(403)
			p.Match(SimpleSqlParserFULL)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(404)
			p.Match(SimpleSqlParserFULL)
		}
		{
			p.SetState(405)
			p.Match(SimpleSqlParserOUTER)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(406)
			p.Match(SimpleSqlParserCROSS)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(407)
			p.Match(SimpleSqlParserNATURAL)
		}

	}

	return localctx
}

// ICmpOpContext is an interface to support dynamic dispatch.
type ICmpOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCmpOpContext differentiates from other interfaces.
	IsCmpOpContext()
}

type CmpOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmpOpContext() *CmpOpContext {
	var p = new(CmpOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_cmpOp
	return p
}

func (*CmpOpContext) IsCmpOpContext() {}

func NewCmpOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmpOpContext {
	var p = new(CmpOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_cmpOp

	return p
}

func (s *CmpOpContext) GetParser() antlr.Parser { return s.parser }
func (s *CmpOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmpOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmpOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterCmpOp(s)
	}
}

func (s *CmpOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitCmpOp(s)
	}
}

func (s *CmpOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitCmpOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) CmpOp() (localctx ICmpOpContext) {
	this := p
	_ = this

	localctx = NewCmpOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SimpleSqlParserRULE_cmpOp)
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
		p.SetState(410)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SimpleSqlParserT__7)|(1<<SimpleSqlParserT__8)|(1<<SimpleSqlParserT__9)|(1<<SimpleSqlParserT__10)|(1<<SimpleSqlParserT__11)|(1<<SimpleSqlParserT__12)|(1<<SimpleSqlParserT__13))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IArithOpContext is an interface to support dynamic dispatch.
type IArithOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArithOpContext differentiates from other interfaces.
	IsArithOpContext()
}

type ArithOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArithOpContext() *ArithOpContext {
	var p = new(ArithOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_arithOp
	return p
}

func (*ArithOpContext) IsArithOpContext() {}

func NewArithOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArithOpContext {
	var p = new(ArithOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_arithOp

	return p
}

func (s *ArithOpContext) GetParser() antlr.Parser { return s.parser }
func (s *ArithOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArithOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterArithOp(s)
	}
}

func (s *ArithOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitArithOp(s)
	}
}

func (s *ArithOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitArithOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) ArithOp() (localctx IArithOpContext) {
	this := p
	_ = this

	localctx = NewArithOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SimpleSqlParserRULE_arithOp)
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
		p.SetState(412)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SimpleSqlParserT__4)|(1<<SimpleSqlParserT__14)|(1<<SimpleSqlParserT__15)|(1<<SimpleSqlParserT__16)|(1<<SimpleSqlParserT__17)|(1<<SimpleSqlParserT__18))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IUnaryOpContext is an interface to support dynamic dispatch.
type IUnaryOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryOpContext differentiates from other interfaces.
	IsUnaryOpContext()
}

type UnaryOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOpContext() *UnaryOpContext {
	var p = new(UnaryOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_unaryOp
	return p
}

func (*UnaryOpContext) IsUnaryOpContext() {}

func NewUnaryOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOpContext {
	var p = new(UnaryOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_unaryOp

	return p
}

func (s *UnaryOpContext) GetParser() antlr.Parser { return s.parser }
func (s *UnaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterUnaryOp(s)
	}
}

func (s *UnaryOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitUnaryOp(s)
	}
}

func (s *UnaryOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitUnaryOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) UnaryOp() (localctx IUnaryOpContext) {
	this := p
	_ = this

	localctx = NewUnaryOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SimpleSqlParserRULE_unaryOp)
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
		p.SetState(414)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SimpleSqlParserT__14 || _la == SimpleSqlParserT__15) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IUnquotedIdentifierContext is an interface to support dynamic dispatch.
type IUnquotedIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnquotedIdentifierContext differentiates from other interfaces.
	IsUnquotedIdentifierContext()
}

type UnquotedIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnquotedIdentifierContext() *UnquotedIdentifierContext {
	var p = new(UnquotedIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleSqlParserRULE_unquotedIdentifier
	return p
}

func (*UnquotedIdentifierContext) IsUnquotedIdentifierContext() {}

func NewUnquotedIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnquotedIdentifierContext {
	var p = new(UnquotedIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleSqlParserRULE_unquotedIdentifier

	return p
}

func (s *UnquotedIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *UnquotedIdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserIDENTIFIER, 0)
}

func (s *UnquotedIdentifierContext) LEFT() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserLEFT, 0)
}

func (s *UnquotedIdentifierContext) RIGHT() antlr.TerminalNode {
	return s.GetToken(SimpleSqlParserRIGHT, 0)
}

func (s *UnquotedIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnquotedIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnquotedIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.EnterUnquotedIdentifier(s)
	}
}

func (s *UnquotedIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleSqlListener); ok {
		listenerT.ExitUnquotedIdentifier(s)
	}
}

func (s *UnquotedIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SimpleSqlVisitor:
		return t.VisitUnquotedIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SimpleSqlParser) UnquotedIdentifier() (localctx IUnquotedIdentifierContext) {
	this := p
	_ = this

	localctx = NewUnquotedIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SimpleSqlParserRULE_unquotedIdentifier)
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
		p.SetState(416)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(SimpleSqlParserLEFT-40))|(1<<(SimpleSqlParserRIGHT-40))|(1<<(SimpleSqlParserIDENTIFIER-40)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *SimpleSqlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
		var t *BooleanExpressionContext = nil
		if localctx != nil {
			t = localctx.(*BooleanExpressionContext)
		}
		return p.BooleanExpression_Sempred(t, predIndex)

	case 11:
		var t *ValueExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ValueExpressionContext)
		}
		return p.ValueExpression_Sempred(t, predIndex)

	case 17:
		var t *RelationContext = nil
		if localctx != nil {
			t = localctx.(*RelationContext)
		}
		return p.Relation_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SimpleSqlParser) BooleanExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SimpleSqlParser) ValueExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SimpleSqlParser) Relation_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
