// Code generated from Protobuf3.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Protobuf3

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

type Protobuf3Parser struct {
	*antlr.BaseParser
}

var protobuf3ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func protobuf3ParserInit() {
	staticData := &protobuf3ParserStaticData
	staticData.literalNames = []string{
		"", "'='", "';'", "'('", "')'", "'.'", "'{'", "'}'", "':'", "'-'", "'['",
		"','", "']'", "'<'", "'>'", "'+'", "'bool'", "'bytes'", "'double'",
		"'enum'", "'extend'", "'fixed32'", "'fixed64'", "'float'", "'import'",
		"'int32'", "'int64'", "'map'", "'message'", "'oneof'", "'option'", "'package'",
		"'\"proto3\"'", "''proto3''", "'public'", "'repeated'", "'reserved'",
		"'returns'", "'rpc'", "'service'", "'sfixed32'", "'sfixed64'", "'sint32'",
		"'sint64'", "'stream'", "'string'", "'syntax'", "'to'", "'uint32'",
		"'uint64'", "'weak'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "BOOL",
		"BYTES", "DOUBLE", "ENUM", "EXTEND", "FIXED32", "FIXED64", "FLOAT",
		"IMPORT", "INT32", "INT64", "MAP", "MESSAGE", "ONEOF", "OPTION", "PACKAGE",
		"PROTO3_DOUBLE", "PROTO3_SINGLE", "PUBLIC", "REPEATED", "RESERVED",
		"RETURNS", "RPC", "SERVICE", "SFIXED32", "SFIXED64", "SINT32", "SINT64",
		"STREAM", "STRING", "SYNTAX", "TO", "UINT32", "UINT64", "WEAK", "IDENT",
		"INT_LIT", "FLOAT_LIT", "BOOL_LIT", "STR_LIT", "QUOTE", "WS", "COMMENT",
		"LINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"proto", "syntax", "syntaxExtra", "importStatement", "packageStatement",
		"option", "optionName", "optionBody", "optionBodyVariable", "topLevelDef",
		"message", "messageBody", "messageBodyContent", "enumDefinition", "enumBody",
		"enumField", "enumValueOption", "extend", "service", "rpc", "reserved",
		"ranges", "rangeRule", "fieldNames", "typeRule", "simpleType", "fieldNumber",
		"field", "fieldOptions", "fieldOption", "oneof", "oneofField", "mapField",
		"keyType", "reservedWord", "fullIdent", "messageName", "enumName", "messageOrEnumName",
		"fieldName", "oneofName", "mapName", "serviceName", "rpcName", "messageType",
		"messageOrEnumType", "emptyStatement", "constant",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 59, 464, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 1, 0, 1, 0, 5, 0, 99, 8, 0, 10, 0, 12, 0, 102, 9, 0, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 116,
		8, 2, 1, 3, 1, 3, 3, 3, 120, 8, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 134, 8, 5, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 143, 8, 6, 1, 6, 1, 6, 1, 6, 3, 6, 148, 8,
		6, 5, 6, 150, 8, 6, 10, 6, 12, 6, 153, 9, 6, 1, 7, 1, 7, 5, 7, 157, 8,
		7, 10, 7, 12, 7, 160, 9, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 9, 1, 9, 3, 9, 172, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		5, 11, 180, 8, 11, 10, 11, 12, 11, 183, 9, 11, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 196, 8, 12,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 206, 8,
		14, 10, 14, 12, 14, 209, 9, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 3, 15,
		216, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 223, 8, 15, 10, 15,
		12, 15, 226, 9, 15, 1, 15, 1, 15, 3, 15, 230, 8, 15, 1, 15, 1, 15, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 243, 8,
		17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 253,
		8, 18, 10, 18, 12, 18, 256, 9, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1,
		19, 3, 19, 264, 8, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 271, 8,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 278, 8, 19, 10, 19, 12, 19,
		281, 9, 19, 1, 19, 1, 19, 3, 19, 285, 8, 19, 1, 20, 1, 20, 1, 20, 3, 20,
		290, 8, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 5, 21, 297, 8, 21, 10, 21,
		12, 21, 300, 9, 21, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 306, 8, 22, 1, 23,
		1, 23, 1, 23, 5, 23, 311, 8, 23, 10, 23, 12, 23, 314, 9, 23, 1, 24, 1,
		24, 3, 24, 318, 8, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 3, 27, 325, 8,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 335,
		8, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 5, 28, 342, 8, 28, 10, 28, 12,
		28, 345, 9, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 5, 30, 356, 8, 30, 10, 30, 12, 30, 359, 9, 30, 1, 30, 1, 30, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 371, 8, 31,
		1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 388, 8, 32, 1, 32, 1, 32, 1, 33,
		1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 5, 35, 399, 8, 35, 10, 35, 12,
		35, 402, 9, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39,
		3, 39, 412, 8, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 1,
		43, 1, 44, 3, 44, 423, 8, 44, 1, 44, 1, 44, 5, 44, 427, 8, 44, 10, 44,
		12, 44, 430, 9, 44, 1, 44, 1, 44, 1, 45, 3, 45, 435, 8, 45, 1, 45, 1, 45,
		3, 45, 439, 8, 45, 1, 45, 5, 45, 442, 8, 45, 10, 45, 12, 45, 445, 9, 45,
		1, 45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47, 3, 47, 453, 8, 47, 1, 47, 1,
		47, 3, 47, 457, 8, 47, 1, 47, 1, 47, 1, 47, 3, 47, 462, 8, 47, 1, 47, 0,
		0, 48, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
		72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 0, 6, 1, 0, 32, 33, 2,
		0, 34, 34, 50, 50, 6, 0, 16, 18, 21, 23, 25, 26, 40, 43, 45, 45, 48, 49,
		6, 0, 16, 16, 21, 22, 25, 26, 40, 43, 45, 45, 48, 49, 6, 0, 20, 20, 28,
		28, 30, 31, 38, 39, 44, 46, 50, 50, 2, 0, 9, 9, 15, 15, 478, 0, 96, 1,
		0, 0, 0, 2, 105, 1, 0, 0, 0, 4, 115, 1, 0, 0, 0, 6, 117, 1, 0, 0, 0, 8,
		124, 1, 0, 0, 0, 10, 128, 1, 0, 0, 0, 12, 142, 1, 0, 0, 0, 14, 154, 1,
		0, 0, 0, 16, 163, 1, 0, 0, 0, 18, 171, 1, 0, 0, 0, 20, 173, 1, 0, 0, 0,
		22, 177, 1, 0, 0, 0, 24, 195, 1, 0, 0, 0, 26, 197, 1, 0, 0, 0, 28, 201,
		1, 0, 0, 0, 30, 212, 1, 0, 0, 0, 32, 233, 1, 0, 0, 0, 34, 237, 1, 0, 0,
		0, 36, 246, 1, 0, 0, 0, 38, 259, 1, 0, 0, 0, 40, 286, 1, 0, 0, 0, 42, 293,
		1, 0, 0, 0, 44, 305, 1, 0, 0, 0, 46, 307, 1, 0, 0, 0, 48, 317, 1, 0, 0,
		0, 50, 319, 1, 0, 0, 0, 52, 321, 1, 0, 0, 0, 54, 324, 1, 0, 0, 0, 56, 338,
		1, 0, 0, 0, 58, 346, 1, 0, 0, 0, 60, 350, 1, 0, 0, 0, 62, 362, 1, 0, 0,
		0, 64, 374, 1, 0, 0, 0, 66, 391, 1, 0, 0, 0, 68, 393, 1, 0, 0, 0, 70, 395,
		1, 0, 0, 0, 72, 403, 1, 0, 0, 0, 74, 405, 1, 0, 0, 0, 76, 407, 1, 0, 0,
		0, 78, 411, 1, 0, 0, 0, 80, 413, 1, 0, 0, 0, 82, 415, 1, 0, 0, 0, 84, 417,
		1, 0, 0, 0, 86, 419, 1, 0, 0, 0, 88, 422, 1, 0, 0, 0, 90, 434, 1, 0, 0,
		0, 92, 448, 1, 0, 0, 0, 94, 461, 1, 0, 0, 0, 96, 100, 3, 2, 1, 0, 97, 99,
		3, 4, 2, 0, 98, 97, 1, 0, 0, 0, 99, 102, 1, 0, 0, 0, 100, 98, 1, 0, 0,
		0, 100, 101, 1, 0, 0, 0, 101, 103, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 103,
		104, 5, 0, 0, 1, 104, 1, 1, 0, 0, 0, 105, 106, 5, 46, 0, 0, 106, 107, 5,
		1, 0, 0, 107, 108, 7, 0, 0, 0, 108, 109, 5, 2, 0, 0, 109, 3, 1, 0, 0, 0,
		110, 116, 3, 6, 3, 0, 111, 116, 3, 8, 4, 0, 112, 116, 3, 10, 5, 0, 113,
		116, 3, 18, 9, 0, 114, 116, 3, 92, 46, 0, 115, 110, 1, 0, 0, 0, 115, 111,
		1, 0, 0, 0, 115, 112, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 114, 1, 0,
		0, 0, 116, 5, 1, 0, 0, 0, 117, 119, 5, 24, 0, 0, 118, 120, 7, 1, 0, 0,
		119, 118, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121,
		122, 5, 55, 0, 0, 122, 123, 5, 2, 0, 0, 123, 7, 1, 0, 0, 0, 124, 125, 5,
		31, 0, 0, 125, 126, 3, 70, 35, 0, 126, 127, 5, 2, 0, 0, 127, 9, 1, 0, 0,
		0, 128, 129, 5, 30, 0, 0, 129, 130, 3, 12, 6, 0, 130, 133, 5, 1, 0, 0,
		131, 134, 3, 94, 47, 0, 132, 134, 3, 14, 7, 0, 133, 131, 1, 0, 0, 0, 133,
		132, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 136, 5, 2, 0, 0, 136, 11, 1,
		0, 0, 0, 137, 143, 5, 51, 0, 0, 138, 139, 5, 3, 0, 0, 139, 140, 3, 70,
		35, 0, 140, 141, 5, 4, 0, 0, 141, 143, 1, 0, 0, 0, 142, 137, 1, 0, 0, 0,
		142, 138, 1, 0, 0, 0, 143, 151, 1, 0, 0, 0, 144, 147, 5, 5, 0, 0, 145,
		148, 5, 51, 0, 0, 146, 148, 3, 68, 34, 0, 147, 145, 1, 0, 0, 0, 147, 146,
		1, 0, 0, 0, 148, 150, 1, 0, 0, 0, 149, 144, 1, 0, 0, 0, 150, 153, 1, 0,
		0, 0, 151, 149, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 13, 1, 0, 0, 0,
		153, 151, 1, 0, 0, 0, 154, 158, 5, 6, 0, 0, 155, 157, 3, 16, 8, 0, 156,
		155, 1, 0, 0, 0, 157, 160, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 158, 159,
		1, 0, 0, 0, 159, 161, 1, 0, 0, 0, 160, 158, 1, 0, 0, 0, 161, 162, 5, 7,
		0, 0, 162, 15, 1, 0, 0, 0, 163, 164, 3, 12, 6, 0, 164, 165, 5, 8, 0, 0,
		165, 166, 3, 94, 47, 0, 166, 17, 1, 0, 0, 0, 167, 172, 3, 20, 10, 0, 168,
		172, 3, 26, 13, 0, 169, 172, 3, 34, 17, 0, 170, 172, 3, 36, 18, 0, 171,
		167, 1, 0, 0, 0, 171, 168, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 171, 170,
		1, 0, 0, 0, 172, 19, 1, 0, 0, 0, 173, 174, 5, 28, 0, 0, 174, 175, 3, 72,
		36, 0, 175, 176, 3, 22, 11, 0, 176, 21, 1, 0, 0, 0, 177, 181, 5, 6, 0,
		0, 178, 180, 3, 24, 12, 0, 179, 178, 1, 0, 0, 0, 180, 183, 1, 0, 0, 0,
		181, 179, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 184, 1, 0, 0, 0, 183,
		181, 1, 0, 0, 0, 184, 185, 5, 7, 0, 0, 185, 23, 1, 0, 0, 0, 186, 196, 3,
		54, 27, 0, 187, 196, 3, 26, 13, 0, 188, 196, 3, 20, 10, 0, 189, 196, 3,
		34, 17, 0, 190, 196, 3, 10, 5, 0, 191, 196, 3, 60, 30, 0, 192, 196, 3,
		64, 32, 0, 193, 196, 3, 40, 20, 0, 194, 196, 3, 92, 46, 0, 195, 186, 1,
		0, 0, 0, 195, 187, 1, 0, 0, 0, 195, 188, 1, 0, 0, 0, 195, 189, 1, 0, 0,
		0, 195, 190, 1, 0, 0, 0, 195, 191, 1, 0, 0, 0, 195, 192, 1, 0, 0, 0, 195,
		193, 1, 0, 0, 0, 195, 194, 1, 0, 0, 0, 196, 25, 1, 0, 0, 0, 197, 198, 5,
		19, 0, 0, 198, 199, 3, 74, 37, 0, 199, 200, 3, 28, 14, 0, 200, 27, 1, 0,
		0, 0, 201, 207, 5, 6, 0, 0, 202, 206, 3, 10, 5, 0, 203, 206, 3, 30, 15,
		0, 204, 206, 3, 92, 46, 0, 205, 202, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0,
		205, 204, 1, 0, 0, 0, 206, 209, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 207,
		208, 1, 0, 0, 0, 208, 210, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 210, 211,
		5, 7, 0, 0, 211, 29, 1, 0, 0, 0, 212, 213, 5, 51, 0, 0, 213, 215, 5, 1,
		0, 0, 214, 216, 5, 9, 0, 0, 215, 214, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0,
		216, 217, 1, 0, 0, 0, 217, 229, 5, 52, 0, 0, 218, 219, 5, 10, 0, 0, 219,
		224, 3, 32, 16, 0, 220, 221, 5, 11, 0, 0, 221, 223, 3, 32, 16, 0, 222,
		220, 1, 0, 0, 0, 223, 226, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 224, 225,
		1, 0, 0, 0, 225, 227, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 227, 228, 5, 12,
		0, 0, 228, 230, 1, 0, 0, 0, 229, 218, 1, 0, 0, 0, 229, 230, 1, 0, 0, 0,
		230, 231, 1, 0, 0, 0, 231, 232, 5, 2, 0, 0, 232, 31, 1, 0, 0, 0, 233, 234,
		3, 12, 6, 0, 234, 235, 5, 1, 0, 0, 235, 236, 3, 94, 47, 0, 236, 33, 1,
		0, 0, 0, 237, 238, 5, 20, 0, 0, 238, 239, 3, 88, 44, 0, 239, 242, 5, 6,
		0, 0, 240, 243, 3, 54, 27, 0, 241, 243, 3, 92, 46, 0, 242, 240, 1, 0, 0,
		0, 242, 241, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244, 245, 5, 7, 0, 0, 245,
		35, 1, 0, 0, 0, 246, 247, 5, 39, 0, 0, 247, 248, 3, 84, 42, 0, 248, 254,
		5, 6, 0, 0, 249, 253, 3, 10, 5, 0, 250, 253, 3, 38, 19, 0, 251, 253, 3,
		92, 46, 0, 252, 249, 1, 0, 0, 0, 252, 250, 1, 0, 0, 0, 252, 251, 1, 0,
		0, 0, 253, 256, 1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0,
		255, 257, 1, 0, 0, 0, 256, 254, 1, 0, 0, 0, 257, 258, 5, 7, 0, 0, 258,
		37, 1, 0, 0, 0, 259, 260, 5, 38, 0, 0, 260, 261, 3, 86, 43, 0, 261, 263,
		5, 3, 0, 0, 262, 264, 5, 44, 0, 0, 263, 262, 1, 0, 0, 0, 263, 264, 1, 0,
		0, 0, 264, 265, 1, 0, 0, 0, 265, 266, 3, 88, 44, 0, 266, 267, 5, 4, 0,
		0, 267, 268, 5, 37, 0, 0, 268, 270, 5, 3, 0, 0, 269, 271, 5, 44, 0, 0,
		270, 269, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272,
		273, 3, 88, 44, 0, 273, 284, 5, 4, 0, 0, 274, 279, 5, 6, 0, 0, 275, 278,
		3, 10, 5, 0, 276, 278, 3, 92, 46, 0, 277, 275, 1, 0, 0, 0, 277, 276, 1,
		0, 0, 0, 278, 281, 1, 0, 0, 0, 279, 277, 1, 0, 0, 0, 279, 280, 1, 0, 0,
		0, 280, 282, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 282, 285, 5, 7, 0, 0, 283,
		285, 5, 2, 0, 0, 284, 274, 1, 0, 0, 0, 284, 283, 1, 0, 0, 0, 285, 39, 1,
		0, 0, 0, 286, 289, 5, 36, 0, 0, 287, 290, 3, 42, 21, 0, 288, 290, 3, 46,
		23, 0, 289, 287, 1, 0, 0, 0, 289, 288, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0,
		291, 292, 5, 2, 0, 0, 292, 41, 1, 0, 0, 0, 293, 298, 3, 44, 22, 0, 294,
		295, 5, 11, 0, 0, 295, 297, 3, 44, 22, 0, 296, 294, 1, 0, 0, 0, 297, 300,
		1, 0, 0, 0, 298, 296, 1, 0, 0, 0, 298, 299, 1, 0, 0, 0, 299, 43, 1, 0,
		0, 0, 300, 298, 1, 0, 0, 0, 301, 306, 5, 52, 0, 0, 302, 303, 5, 52, 0,
		0, 303, 304, 5, 47, 0, 0, 304, 306, 5, 52, 0, 0, 305, 301, 1, 0, 0, 0,
		305, 302, 1, 0, 0, 0, 306, 45, 1, 0, 0, 0, 307, 312, 5, 55, 0, 0, 308,
		309, 5, 11, 0, 0, 309, 311, 5, 55, 0, 0, 310, 308, 1, 0, 0, 0, 311, 314,
		1, 0, 0, 0, 312, 310, 1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 47, 1, 0,
		0, 0, 314, 312, 1, 0, 0, 0, 315, 318, 3, 50, 25, 0, 316, 318, 3, 90, 45,
		0, 317, 315, 1, 0, 0, 0, 317, 316, 1, 0, 0, 0, 318, 49, 1, 0, 0, 0, 319,
		320, 7, 2, 0, 0, 320, 51, 1, 0, 0, 0, 321, 322, 5, 52, 0, 0, 322, 53, 1,
		0, 0, 0, 323, 325, 5, 35, 0, 0, 324, 323, 1, 0, 0, 0, 324, 325, 1, 0, 0,
		0, 325, 326, 1, 0, 0, 0, 326, 327, 3, 48, 24, 0, 327, 328, 3, 78, 39, 0,
		328, 329, 5, 1, 0, 0, 329, 334, 3, 52, 26, 0, 330, 331, 5, 10, 0, 0, 331,
		332, 3, 56, 28, 0, 332, 333, 5, 12, 0, 0, 333, 335, 1, 0, 0, 0, 334, 330,
		1, 0, 0, 0, 334, 335, 1, 0, 0, 0, 335, 336, 1, 0, 0, 0, 336, 337, 5, 2,
		0, 0, 337, 55, 1, 0, 0, 0, 338, 343, 3, 58, 29, 0, 339, 340, 5, 11, 0,
		0, 340, 342, 3, 58, 29, 0, 341, 339, 1, 0, 0, 0, 342, 345, 1, 0, 0, 0,
		343, 341, 1, 0, 0, 0, 343, 344, 1, 0, 0, 0, 344, 57, 1, 0, 0, 0, 345, 343,
		1, 0, 0, 0, 346, 347, 3, 12, 6, 0, 347, 348, 5, 1, 0, 0, 348, 349, 3, 94,
		47, 0, 349, 59, 1, 0, 0, 0, 350, 351, 5, 29, 0, 0, 351, 352, 3, 80, 40,
		0, 352, 357, 5, 6, 0, 0, 353, 356, 3, 62, 31, 0, 354, 356, 3, 92, 46, 0,
		355, 353, 1, 0, 0, 0, 355, 354, 1, 0, 0, 0, 356, 359, 1, 0, 0, 0, 357,
		355, 1, 0, 0, 0, 357, 358, 1, 0, 0, 0, 358, 360, 1, 0, 0, 0, 359, 357,
		1, 0, 0, 0, 360, 361, 5, 7, 0, 0, 361, 61, 1, 0, 0, 0, 362, 363, 3, 48,
		24, 0, 363, 364, 3, 78, 39, 0, 364, 365, 5, 1, 0, 0, 365, 370, 3, 52, 26,
		0, 366, 367, 5, 10, 0, 0, 367, 368, 3, 56, 28, 0, 368, 369, 5, 12, 0, 0,
		369, 371, 1, 0, 0, 0, 370, 366, 1, 0, 0, 0, 370, 371, 1, 0, 0, 0, 371,
		372, 1, 0, 0, 0, 372, 373, 5, 2, 0, 0, 373, 63, 1, 0, 0, 0, 374, 375, 5,
		27, 0, 0, 375, 376, 5, 13, 0, 0, 376, 377, 3, 66, 33, 0, 377, 378, 5, 11,
		0, 0, 378, 379, 3, 48, 24, 0, 379, 380, 5, 14, 0, 0, 380, 381, 3, 82, 41,
		0, 381, 382, 5, 1, 0, 0, 382, 387, 3, 52, 26, 0, 383, 384, 5, 10, 0, 0,
		384, 385, 3, 56, 28, 0, 385, 386, 5, 12, 0, 0, 386, 388, 1, 0, 0, 0, 387,
		383, 1, 0, 0, 0, 387, 388, 1, 0, 0, 0, 388, 389, 1, 0, 0, 0, 389, 390,
		5, 2, 0, 0, 390, 65, 1, 0, 0, 0, 391, 392, 7, 3, 0, 0, 392, 67, 1, 0, 0,
		0, 393, 394, 7, 4, 0, 0, 394, 69, 1, 0, 0, 0, 395, 400, 5, 51, 0, 0, 396,
		397, 5, 5, 0, 0, 397, 399, 5, 51, 0, 0, 398, 396, 1, 0, 0, 0, 399, 402,
		1, 0, 0, 0, 400, 398, 1, 0, 0, 0, 400, 401, 1, 0, 0, 0, 401, 71, 1, 0,
		0, 0, 402, 400, 1, 0, 0, 0, 403, 404, 5, 51, 0, 0, 404, 73, 1, 0, 0, 0,
		405, 406, 5, 51, 0, 0, 406, 75, 1, 0, 0, 0, 407, 408, 5, 51, 0, 0, 408,
		77, 1, 0, 0, 0, 409, 412, 5, 51, 0, 0, 410, 412, 3, 68, 34, 0, 411, 409,
		1, 0, 0, 0, 411, 410, 1, 0, 0, 0, 412, 79, 1, 0, 0, 0, 413, 414, 5, 51,
		0, 0, 414, 81, 1, 0, 0, 0, 415, 416, 5, 51, 0, 0, 416, 83, 1, 0, 0, 0,
		417, 418, 5, 51, 0, 0, 418, 85, 1, 0, 0, 0, 419, 420, 5, 51, 0, 0, 420,
		87, 1, 0, 0, 0, 421, 423, 5, 5, 0, 0, 422, 421, 1, 0, 0, 0, 422, 423, 1,
		0, 0, 0, 423, 428, 1, 0, 0, 0, 424, 425, 5, 51, 0, 0, 425, 427, 5, 5, 0,
		0, 426, 424, 1, 0, 0, 0, 427, 430, 1, 0, 0, 0, 428, 426, 1, 0, 0, 0, 428,
		429, 1, 0, 0, 0, 429, 431, 1, 0, 0, 0, 430, 428, 1, 0, 0, 0, 431, 432,
		3, 72, 36, 0, 432, 89, 1, 0, 0, 0, 433, 435, 5, 5, 0, 0, 434, 433, 1, 0,
		0, 0, 434, 435, 1, 0, 0, 0, 435, 443, 1, 0, 0, 0, 436, 439, 5, 51, 0, 0,
		437, 439, 3, 68, 34, 0, 438, 436, 1, 0, 0, 0, 438, 437, 1, 0, 0, 0, 439,
		440, 1, 0, 0, 0, 440, 442, 5, 5, 0, 0, 441, 438, 1, 0, 0, 0, 442, 445,
		1, 0, 0, 0, 443, 441, 1, 0, 0, 0, 443, 444, 1, 0, 0, 0, 444, 446, 1, 0,
		0, 0, 445, 443, 1, 0, 0, 0, 446, 447, 3, 76, 38, 0, 447, 91, 1, 0, 0, 0,
		448, 449, 5, 2, 0, 0, 449, 93, 1, 0, 0, 0, 450, 462, 3, 70, 35, 0, 451,
		453, 7, 5, 0, 0, 452, 451, 1, 0, 0, 0, 452, 453, 1, 0, 0, 0, 453, 454,
		1, 0, 0, 0, 454, 462, 5, 52, 0, 0, 455, 457, 7, 5, 0, 0, 456, 455, 1, 0,
		0, 0, 456, 457, 1, 0, 0, 0, 457, 458, 1, 0, 0, 0, 458, 462, 5, 53, 0, 0,
		459, 462, 5, 55, 0, 0, 460, 462, 5, 54, 0, 0, 461, 450, 1, 0, 0, 0, 461,
		452, 1, 0, 0, 0, 461, 456, 1, 0, 0, 0, 461, 459, 1, 0, 0, 0, 461, 460,
		1, 0, 0, 0, 462, 95, 1, 0, 0, 0, 46, 100, 115, 119, 133, 142, 147, 151,
		158, 171, 181, 195, 205, 207, 215, 224, 229, 242, 252, 254, 263, 270, 277,
		279, 284, 289, 298, 305, 312, 317, 324, 334, 343, 355, 357, 370, 387, 400,
		411, 422, 428, 434, 438, 443, 452, 456, 461,
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

// Protobuf3ParserInit initializes any static state used to implement Protobuf3Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewProtobuf3Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Protobuf3ParserInit() {
	staticData := &protobuf3ParserStaticData
	staticData.once.Do(protobuf3ParserInit)
}

// NewProtobuf3Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewProtobuf3Parser(input antlr.TokenStream) *Protobuf3Parser {
	Protobuf3ParserInit()
	this := new(Protobuf3Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &protobuf3ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Protobuf3.g4"

	return this
}

// Protobuf3Parser tokens.
const (
	Protobuf3ParserEOF           = antlr.TokenEOF
	Protobuf3ParserT__0          = 1
	Protobuf3ParserT__1          = 2
	Protobuf3ParserT__2          = 3
	Protobuf3ParserT__3          = 4
	Protobuf3ParserT__4          = 5
	Protobuf3ParserT__5          = 6
	Protobuf3ParserT__6          = 7
	Protobuf3ParserT__7          = 8
	Protobuf3ParserT__8          = 9
	Protobuf3ParserT__9          = 10
	Protobuf3ParserT__10         = 11
	Protobuf3ParserT__11         = 12
	Protobuf3ParserT__12         = 13
	Protobuf3ParserT__13         = 14
	Protobuf3ParserT__14         = 15
	Protobuf3ParserBOOL          = 16
	Protobuf3ParserBYTES         = 17
	Protobuf3ParserDOUBLE        = 18
	Protobuf3ParserENUM          = 19
	Protobuf3ParserEXTEND        = 20
	Protobuf3ParserFIXED32       = 21
	Protobuf3ParserFIXED64       = 22
	Protobuf3ParserFLOAT         = 23
	Protobuf3ParserIMPORT        = 24
	Protobuf3ParserINT32         = 25
	Protobuf3ParserINT64         = 26
	Protobuf3ParserMAP           = 27
	Protobuf3ParserMESSAGE       = 28
	Protobuf3ParserONEOF         = 29
	Protobuf3ParserOPTION        = 30
	Protobuf3ParserPACKAGE       = 31
	Protobuf3ParserPROTO3_DOUBLE = 32
	Protobuf3ParserPROTO3_SINGLE = 33
	Protobuf3ParserPUBLIC        = 34
	Protobuf3ParserREPEATED      = 35
	Protobuf3ParserRESERVED      = 36
	Protobuf3ParserRETURNS       = 37
	Protobuf3ParserRPC           = 38
	Protobuf3ParserSERVICE       = 39
	Protobuf3ParserSFIXED32      = 40
	Protobuf3ParserSFIXED64      = 41
	Protobuf3ParserSINT32        = 42
	Protobuf3ParserSINT64        = 43
	Protobuf3ParserSTREAM        = 44
	Protobuf3ParserSTRING        = 45
	Protobuf3ParserSYNTAX        = 46
	Protobuf3ParserTO            = 47
	Protobuf3ParserUINT32        = 48
	Protobuf3ParserUINT64        = 49
	Protobuf3ParserWEAK          = 50
	Protobuf3ParserIDENT         = 51
	Protobuf3ParserINT_LIT       = 52
	Protobuf3ParserFLOAT_LIT     = 53
	Protobuf3ParserBOOL_LIT      = 54
	Protobuf3ParserSTR_LIT       = 55
	Protobuf3ParserQUOTE         = 56
	Protobuf3ParserWS            = 57
	Protobuf3ParserCOMMENT       = 58
	Protobuf3ParserLINE_COMMENT  = 59
)

// Protobuf3Parser rules.
const (
	Protobuf3ParserRULE_proto              = 0
	Protobuf3ParserRULE_syntax             = 1
	Protobuf3ParserRULE_syntaxExtra        = 2
	Protobuf3ParserRULE_importStatement    = 3
	Protobuf3ParserRULE_packageStatement   = 4
	Protobuf3ParserRULE_option             = 5
	Protobuf3ParserRULE_optionName         = 6
	Protobuf3ParserRULE_optionBody         = 7
	Protobuf3ParserRULE_optionBodyVariable = 8
	Protobuf3ParserRULE_topLevelDef        = 9
	Protobuf3ParserRULE_message            = 10
	Protobuf3ParserRULE_messageBody        = 11
	Protobuf3ParserRULE_messageBodyContent = 12
	Protobuf3ParserRULE_enumDefinition     = 13
	Protobuf3ParserRULE_enumBody           = 14
	Protobuf3ParserRULE_enumField          = 15
	Protobuf3ParserRULE_enumValueOption    = 16
	Protobuf3ParserRULE_extend             = 17
	Protobuf3ParserRULE_service            = 18
	Protobuf3ParserRULE_rpc                = 19
	Protobuf3ParserRULE_reserved           = 20
	Protobuf3ParserRULE_ranges             = 21
	Protobuf3ParserRULE_rangeRule          = 22
	Protobuf3ParserRULE_fieldNames         = 23
	Protobuf3ParserRULE_typeRule           = 24
	Protobuf3ParserRULE_simpleType         = 25
	Protobuf3ParserRULE_fieldNumber        = 26
	Protobuf3ParserRULE_field              = 27
	Protobuf3ParserRULE_fieldOptions       = 28
	Protobuf3ParserRULE_fieldOption        = 29
	Protobuf3ParserRULE_oneof              = 30
	Protobuf3ParserRULE_oneofField         = 31
	Protobuf3ParserRULE_mapField           = 32
	Protobuf3ParserRULE_keyType            = 33
	Protobuf3ParserRULE_reservedWord       = 34
	Protobuf3ParserRULE_fullIdent          = 35
	Protobuf3ParserRULE_messageName        = 36
	Protobuf3ParserRULE_enumName           = 37
	Protobuf3ParserRULE_messageOrEnumName  = 38
	Protobuf3ParserRULE_fieldName          = 39
	Protobuf3ParserRULE_oneofName          = 40
	Protobuf3ParserRULE_mapName            = 41
	Protobuf3ParserRULE_serviceName        = 42
	Protobuf3ParserRULE_rpcName            = 43
	Protobuf3ParserRULE_messageType        = 44
	Protobuf3ParserRULE_messageOrEnumType  = 45
	Protobuf3ParserRULE_emptyStatement     = 46
	Protobuf3ParserRULE_constant           = 47
)

// IProtoContext is an interface to support dynamic dispatch.
type IProtoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Syntax() ISyntaxContext
	EOF() antlr.TerminalNode
	AllSyntaxExtra() []ISyntaxExtraContext
	SyntaxExtra(i int) ISyntaxExtraContext

	// IsProtoContext differentiates from other interfaces.
	IsProtoContext()
}

type ProtoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProtoContext() *ProtoContext {
	var p = new(ProtoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_proto
	return p
}

func (*ProtoContext) IsProtoContext() {}

func NewProtoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProtoContext {
	var p = new(ProtoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_proto

	return p
}

func (s *ProtoContext) GetParser() antlr.Parser { return s.parser }

func (s *ProtoContext) Syntax() ISyntaxContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISyntaxContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISyntaxContext)
}

func (s *ProtoContext) EOF() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserEOF, 0)
}

func (s *ProtoContext) AllSyntaxExtra() []ISyntaxExtraContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISyntaxExtraContext); ok {
			len++
		}
	}

	tst := make([]ISyntaxExtraContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISyntaxExtraContext); ok {
			tst[i] = t.(ISyntaxExtraContext)
			i++
		}
	}

	return tst
}

func (s *ProtoContext) SyntaxExtra(i int) ISyntaxExtraContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISyntaxExtraContext); ok {
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

	return t.(ISyntaxExtraContext)
}

func (s *ProtoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProtoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProtoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterProto(s)
	}
}

func (s *ProtoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitProto(s)
	}
}

func (s *ProtoContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitProto(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Proto() (localctx IProtoContext) {
	this := p
	_ = this

	localctx = NewProtoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Protobuf3ParserRULE_proto)
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
		p.SetState(96)
		p.Syntax()
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&553263824900) != 0 {
		{
			p.SetState(97)
			p.SyntaxExtra()
		}

		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(103)
		p.Match(Protobuf3ParserEOF)
	}

	return localctx
}

// ISyntaxContext is an interface to support dynamic dispatch.
type ISyntaxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SYNTAX() antlr.TerminalNode
	PROTO3_DOUBLE() antlr.TerminalNode
	PROTO3_SINGLE() antlr.TerminalNode

	// IsSyntaxContext differentiates from other interfaces.
	IsSyntaxContext()
}

type SyntaxContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySyntaxContext() *SyntaxContext {
	var p = new(SyntaxContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_syntax
	return p
}

func (*SyntaxContext) IsSyntaxContext() {}

func NewSyntaxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SyntaxContext {
	var p = new(SyntaxContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_syntax

	return p
}

func (s *SyntaxContext) GetParser() antlr.Parser { return s.parser }

func (s *SyntaxContext) SYNTAX() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSYNTAX, 0)
}

func (s *SyntaxContext) PROTO3_DOUBLE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPROTO3_DOUBLE, 0)
}

func (s *SyntaxContext) PROTO3_SINGLE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPROTO3_SINGLE, 0)
}

func (s *SyntaxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SyntaxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SyntaxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterSyntax(s)
	}
}

func (s *SyntaxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitSyntax(s)
	}
}

func (s *SyntaxContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitSyntax(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Syntax() (localctx ISyntaxContext) {
	this := p
	_ = this

	localctx = NewSyntaxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Protobuf3ParserRULE_syntax)
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
		p.SetState(105)
		p.Match(Protobuf3ParserSYNTAX)
	}
	{
		p.SetState(106)
		p.Match(Protobuf3ParserT__0)
	}
	{
		p.SetState(107)
		_la = p.GetTokenStream().LA(1)

		if !(_la == Protobuf3ParserPROTO3_DOUBLE || _la == Protobuf3ParserPROTO3_SINGLE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(108)
		p.Match(Protobuf3ParserT__1)
	}

	return localctx
}

// ISyntaxExtraContext is an interface to support dynamic dispatch.
type ISyntaxExtraContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ImportStatement() IImportStatementContext
	PackageStatement() IPackageStatementContext
	Option() IOptionContext
	TopLevelDef() ITopLevelDefContext
	EmptyStatement() IEmptyStatementContext

	// IsSyntaxExtraContext differentiates from other interfaces.
	IsSyntaxExtraContext()
}

type SyntaxExtraContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySyntaxExtraContext() *SyntaxExtraContext {
	var p = new(SyntaxExtraContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_syntaxExtra
	return p
}

func (*SyntaxExtraContext) IsSyntaxExtraContext() {}

func NewSyntaxExtraContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SyntaxExtraContext {
	var p = new(SyntaxExtraContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_syntaxExtra

	return p
}

func (s *SyntaxExtraContext) GetParser() antlr.Parser { return s.parser }

func (s *SyntaxExtraContext) ImportStatement() IImportStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportStatementContext)
}

func (s *SyntaxExtraContext) PackageStatement() IPackageStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPackageStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPackageStatementContext)
}

func (s *SyntaxExtraContext) Option() IOptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *SyntaxExtraContext) TopLevelDef() ITopLevelDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopLevelDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITopLevelDefContext)
}

func (s *SyntaxExtraContext) EmptyStatement() IEmptyStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptyStatementContext)
}

func (s *SyntaxExtraContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SyntaxExtraContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SyntaxExtraContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterSyntaxExtra(s)
	}
}

func (s *SyntaxExtraContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitSyntaxExtra(s)
	}
}

func (s *SyntaxExtraContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitSyntaxExtra(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) SyntaxExtra() (localctx ISyntaxExtraContext) {
	this := p
	_ = this

	localctx = NewSyntaxExtraContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Protobuf3ParserRULE_syntaxExtra)

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

	p.SetState(115)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserIMPORT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.ImportStatement()
		}

	case Protobuf3ParserPACKAGE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(111)
			p.PackageStatement()
		}

	case Protobuf3ParserOPTION:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(112)
			p.Option()
		}

	case Protobuf3ParserENUM, Protobuf3ParserEXTEND, Protobuf3ParserMESSAGE, Protobuf3ParserSERVICE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(113)
			p.TopLevelDef()
		}

	case Protobuf3ParserT__1:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(114)
			p.EmptyStatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IImportStatementContext is an interface to support dynamic dispatch.
type IImportStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMPORT() antlr.TerminalNode
	STR_LIT() antlr.TerminalNode
	WEAK() antlr.TerminalNode
	PUBLIC() antlr.TerminalNode

	// IsImportStatementContext differentiates from other interfaces.
	IsImportStatementContext()
}

type ImportStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStatementContext() *ImportStatementContext {
	var p = new(ImportStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_importStatement
	return p
}

func (*ImportStatementContext) IsImportStatementContext() {}

func NewImportStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStatementContext {
	var p = new(ImportStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_importStatement

	return p
}

func (s *ImportStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStatementContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIMPORT, 0)
}

func (s *ImportStatementContext) STR_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTR_LIT, 0)
}

func (s *ImportStatementContext) WEAK() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserWEAK, 0)
}

func (s *ImportStatementContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPUBLIC, 0)
}

func (s *ImportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterImportStatement(s)
	}
}

func (s *ImportStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitImportStatement(s)
	}
}

func (s *ImportStatementContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitImportStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) ImportStatement() (localctx IImportStatementContext) {
	this := p
	_ = this

	localctx = NewImportStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Protobuf3ParserRULE_importStatement)
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
		p.SetState(117)
		p.Match(Protobuf3ParserIMPORT)
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserPUBLIC || _la == Protobuf3ParserWEAK {
		{
			p.SetState(118)
			_la = p.GetTokenStream().LA(1)

			if !(_la == Protobuf3ParserPUBLIC || _la == Protobuf3ParserWEAK) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(121)
		p.Match(Protobuf3ParserSTR_LIT)
	}
	{
		p.SetState(122)
		p.Match(Protobuf3ParserT__1)
	}

	return localctx
}

// IPackageStatementContext is an interface to support dynamic dispatch.
type IPackageStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PACKAGE() antlr.TerminalNode
	FullIdent() IFullIdentContext

	// IsPackageStatementContext differentiates from other interfaces.
	IsPackageStatementContext()
}

type PackageStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageStatementContext() *PackageStatementContext {
	var p = new(PackageStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_packageStatement
	return p
}

func (*PackageStatementContext) IsPackageStatementContext() {}

func NewPackageStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageStatementContext {
	var p = new(PackageStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_packageStatement

	return p
}

func (s *PackageStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageStatementContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPACKAGE, 0)
}

func (s *PackageStatementContext) FullIdent() IFullIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFullIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *PackageStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterPackageStatement(s)
	}
}

func (s *PackageStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitPackageStatement(s)
	}
}

func (s *PackageStatementContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitPackageStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) PackageStatement() (localctx IPackageStatementContext) {
	this := p
	_ = this

	localctx = NewPackageStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, Protobuf3ParserRULE_packageStatement)

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
		p.SetState(124)
		p.Match(Protobuf3ParserPACKAGE)
	}
	{
		p.SetState(125)
		p.FullIdent()
	}
	{
		p.SetState(126)
		p.Match(Protobuf3ParserT__1)
	}

	return localctx
}

// IOptionContext is an interface to support dynamic dispatch.
type IOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPTION() antlr.TerminalNode
	OptionName() IOptionNameContext
	Constant() IConstantContext
	OptionBody() IOptionBodyContext

	// IsOptionContext differentiates from other interfaces.
	IsOptionContext()
}

type OptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionContext() *OptionContext {
	var p = new(OptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_option
	return p
}

func (*OptionContext) IsOptionContext() {}

func NewOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionContext {
	var p = new(OptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_option

	return p
}

func (s *OptionContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionContext) OPTION() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserOPTION, 0)
}

func (s *OptionContext) OptionName() IOptionNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *OptionContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *OptionContext) OptionBody() IOptionBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionBodyContext)
}

func (s *OptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterOption(s)
	}
}

func (s *OptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitOption(s)
	}
}

func (s *OptionContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitOption(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Option() (localctx IOptionContext) {
	this := p
	_ = this

	localctx = NewOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, Protobuf3ParserRULE_option)

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
		p.SetState(128)
		p.Match(Protobuf3ParserOPTION)
	}
	{
		p.SetState(129)
		p.OptionName()
	}
	{
		p.SetState(130)
		p.Match(Protobuf3ParserT__0)
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserT__8, Protobuf3ParserT__14, Protobuf3ParserIDENT, Protobuf3ParserINT_LIT, Protobuf3ParserFLOAT_LIT, Protobuf3ParserBOOL_LIT, Protobuf3ParserSTR_LIT:
		{
			p.SetState(131)
			p.Constant()
		}

	case Protobuf3ParserT__5:
		{
			p.SetState(132)
			p.OptionBody()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(135)
		p.Match(Protobuf3ParserT__1)
	}

	return localctx
}

// IOptionNameContext is an interface to support dynamic dispatch.
type IOptionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	FullIdent() IFullIdentContext
	AllReservedWord() []IReservedWordContext
	ReservedWord(i int) IReservedWordContext

	// IsOptionNameContext differentiates from other interfaces.
	IsOptionNameContext()
}

type OptionNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionNameContext() *OptionNameContext {
	var p = new(OptionNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_optionName
	return p
}

func (*OptionNameContext) IsOptionNameContext() {}

func NewOptionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionNameContext {
	var p = new(OptionNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_optionName

	return p
}

func (s *OptionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionNameContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserIDENT)
}

func (s *OptionNameContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIDENT, i)
}

func (s *OptionNameContext) FullIdent() IFullIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFullIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *OptionNameContext) AllReservedWord() []IReservedWordContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IReservedWordContext); ok {
			len++
		}
	}

	tst := make([]IReservedWordContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IReservedWordContext); ok {
			tst[i] = t.(IReservedWordContext)
			i++
		}
	}

	return tst
}

func (s *OptionNameContext) ReservedWord(i int) IReservedWordContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReservedWordContext); ok {
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

	return t.(IReservedWordContext)
}

func (s *OptionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterOptionName(s)
	}
}

func (s *OptionNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitOptionName(s)
	}
}

func (s *OptionNameContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitOptionName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) OptionName() (localctx IOptionNameContext) {
	this := p
	_ = this

	localctx = NewOptionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, Protobuf3ParserRULE_optionName)
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
	p.SetState(142)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserIDENT:
		{
			p.SetState(137)
			p.Match(Protobuf3ParserIDENT)
		}

	case Protobuf3ParserT__2:
		{
			p.SetState(138)
			p.Match(Protobuf3ParserT__2)
		}
		{
			p.SetState(139)
			p.FullIdent()
		}
		{
			p.SetState(140)
			p.Match(Protobuf3ParserT__3)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserT__4 {
		{
			p.SetState(144)
			p.Match(Protobuf3ParserT__4)
		}
		p.SetState(147)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case Protobuf3ParserIDENT:
			{
				p.SetState(145)
				p.Match(Protobuf3ParserIDENT)
			}

		case Protobuf3ParserEXTEND, Protobuf3ParserMESSAGE, Protobuf3ParserOPTION, Protobuf3ParserPACKAGE, Protobuf3ParserRPC, Protobuf3ParserSERVICE, Protobuf3ParserSTREAM, Protobuf3ParserSTRING, Protobuf3ParserSYNTAX, Protobuf3ParserWEAK:
			{
				p.SetState(146)
				p.ReservedWord()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOptionBodyContext is an interface to support dynamic dispatch.
type IOptionBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllOptionBodyVariable() []IOptionBodyVariableContext
	OptionBodyVariable(i int) IOptionBodyVariableContext

	// IsOptionBodyContext differentiates from other interfaces.
	IsOptionBodyContext()
}

type OptionBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionBodyContext() *OptionBodyContext {
	var p = new(OptionBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_optionBody
	return p
}

func (*OptionBodyContext) IsOptionBodyContext() {}

func NewOptionBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionBodyContext {
	var p = new(OptionBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_optionBody

	return p
}

func (s *OptionBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionBodyContext) AllOptionBodyVariable() []IOptionBodyVariableContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOptionBodyVariableContext); ok {
			len++
		}
	}

	tst := make([]IOptionBodyVariableContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOptionBodyVariableContext); ok {
			tst[i] = t.(IOptionBodyVariableContext)
			i++
		}
	}

	return tst
}

func (s *OptionBodyContext) OptionBodyVariable(i int) IOptionBodyVariableContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionBodyVariableContext); ok {
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

	return t.(IOptionBodyVariableContext)
}

func (s *OptionBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterOptionBody(s)
	}
}

func (s *OptionBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitOptionBody(s)
	}
}

func (s *OptionBodyContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitOptionBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) OptionBody() (localctx IOptionBodyContext) {
	this := p
	_ = this

	localctx = NewOptionBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, Protobuf3ParserRULE_optionBody)
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
		p.Match(Protobuf3ParserT__5)
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserT__2 || _la == Protobuf3ParserIDENT {
		{
			p.SetState(155)
			p.OptionBodyVariable()
		}

		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(161)
		p.Match(Protobuf3ParserT__6)
	}

	return localctx
}

// IOptionBodyVariableContext is an interface to support dynamic dispatch.
type IOptionBodyVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OptionName() IOptionNameContext
	Constant() IConstantContext

	// IsOptionBodyVariableContext differentiates from other interfaces.
	IsOptionBodyVariableContext()
}

type OptionBodyVariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionBodyVariableContext() *OptionBodyVariableContext {
	var p = new(OptionBodyVariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_optionBodyVariable
	return p
}

func (*OptionBodyVariableContext) IsOptionBodyVariableContext() {}

func NewOptionBodyVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionBodyVariableContext {
	var p = new(OptionBodyVariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_optionBodyVariable

	return p
}

func (s *OptionBodyVariableContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionBodyVariableContext) OptionName() IOptionNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *OptionBodyVariableContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *OptionBodyVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionBodyVariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionBodyVariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterOptionBodyVariable(s)
	}
}

func (s *OptionBodyVariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitOptionBodyVariable(s)
	}
}

func (s *OptionBodyVariableContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitOptionBodyVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) OptionBodyVariable() (localctx IOptionBodyVariableContext) {
	this := p
	_ = this

	localctx = NewOptionBodyVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, Protobuf3ParserRULE_optionBodyVariable)

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
		p.OptionName()
	}
	{
		p.SetState(164)
		p.Match(Protobuf3ParserT__7)
	}
	{
		p.SetState(165)
		p.Constant()
	}

	return localctx
}

// ITopLevelDefContext is an interface to support dynamic dispatch.
type ITopLevelDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Message() IMessageContext
	EnumDefinition() IEnumDefinitionContext
	Extend() IExtendContext
	Service() IServiceContext

	// IsTopLevelDefContext differentiates from other interfaces.
	IsTopLevelDefContext()
}

type TopLevelDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopLevelDefContext() *TopLevelDefContext {
	var p = new(TopLevelDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_topLevelDef
	return p
}

func (*TopLevelDefContext) IsTopLevelDefContext() {}

func NewTopLevelDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelDefContext {
	var p = new(TopLevelDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_topLevelDef

	return p
}

func (s *TopLevelDefContext) GetParser() antlr.Parser { return s.parser }

func (s *TopLevelDefContext) Message() IMessageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageContext)
}

func (s *TopLevelDefContext) EnumDefinition() IEnumDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumDefinitionContext)
}

func (s *TopLevelDefContext) Extend() IExtendContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtendContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtendContext)
}

func (s *TopLevelDefContext) Service() IServiceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceContext)
}

func (s *TopLevelDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopLevelDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterTopLevelDef(s)
	}
}

func (s *TopLevelDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitTopLevelDef(s)
	}
}

func (s *TopLevelDefContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitTopLevelDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) TopLevelDef() (localctx ITopLevelDefContext) {
	this := p
	_ = this

	localctx = NewTopLevelDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, Protobuf3ParserRULE_topLevelDef)

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

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserMESSAGE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(167)
			p.Message()
		}

	case Protobuf3ParserENUM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(168)
			p.EnumDefinition()
		}

	case Protobuf3ParserEXTEND:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(169)
			p.Extend()
		}

	case Protobuf3ParserSERVICE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(170)
			p.Service()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMessageContext is an interface to support dynamic dispatch.
type IMessageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MESSAGE() antlr.TerminalNode
	MessageName() IMessageNameContext
	MessageBody() IMessageBodyContext

	// IsMessageContext differentiates from other interfaces.
	IsMessageContext()
}

type MessageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageContext() *MessageContext {
	var p = new(MessageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_message
	return p
}

func (*MessageContext) IsMessageContext() {}

func NewMessageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageContext {
	var p = new(MessageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_message

	return p
}

func (s *MessageContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageContext) MESSAGE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserMESSAGE, 0)
}

func (s *MessageContext) MessageName() IMessageNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageNameContext)
}

func (s *MessageContext) MessageBody() IMessageBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageBodyContext)
}

func (s *MessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMessage(s)
	}
}

func (s *MessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMessage(s)
	}
}

func (s *MessageContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitMessage(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Message() (localctx IMessageContext) {
	this := p
	_ = this

	localctx = NewMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, Protobuf3ParserRULE_message)

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
		p.Match(Protobuf3ParserMESSAGE)
	}
	{
		p.SetState(174)
		p.MessageName()
	}
	{
		p.SetState(175)
		p.MessageBody()
	}

	return localctx
}

// IMessageBodyContext is an interface to support dynamic dispatch.
type IMessageBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMessageBodyContent() []IMessageBodyContentContext
	MessageBodyContent(i int) IMessageBodyContentContext

	// IsMessageBodyContext differentiates from other interfaces.
	IsMessageBodyContext()
}

type MessageBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageBodyContext() *MessageBodyContext {
	var p = new(MessageBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageBody
	return p
}

func (*MessageBodyContext) IsMessageBodyContext() {}

func NewMessageBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageBodyContext {
	var p = new(MessageBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_messageBody

	return p
}

func (s *MessageBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageBodyContext) AllMessageBodyContent() []IMessageBodyContentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMessageBodyContentContext); ok {
			len++
		}
	}

	tst := make([]IMessageBodyContentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMessageBodyContentContext); ok {
			tst[i] = t.(IMessageBodyContentContext)
			i++
		}
	}

	return tst
}

func (s *MessageBodyContext) MessageBodyContent(i int) IMessageBodyContentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageBodyContentContext); ok {
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

	return t.(IMessageBodyContentContext)
}

func (s *MessageBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMessageBody(s)
	}
}

func (s *MessageBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMessageBody(s)
	}
}

func (s *MessageBodyContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitMessageBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) MessageBody() (localctx IMessageBodyContext) {
	this := p
	_ = this

	localctx = NewMessageBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, Protobuf3ParserRULE_messageBody)
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
		p.SetState(177)
		p.Match(Protobuf3ParserT__5)
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4362694618447908) != 0 {
		{
			p.SetState(178)
			p.MessageBodyContent()
		}

		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(184)
		p.Match(Protobuf3ParserT__6)
	}

	return localctx
}

// IMessageBodyContentContext is an interface to support dynamic dispatch.
type IMessageBodyContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Field() IFieldContext
	EnumDefinition() IEnumDefinitionContext
	Message() IMessageContext
	Extend() IExtendContext
	Option() IOptionContext
	Oneof() IOneofContext
	MapField() IMapFieldContext
	Reserved() IReservedContext
	EmptyStatement() IEmptyStatementContext

	// IsMessageBodyContentContext differentiates from other interfaces.
	IsMessageBodyContentContext()
}

type MessageBodyContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageBodyContentContext() *MessageBodyContentContext {
	var p = new(MessageBodyContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageBodyContent
	return p
}

func (*MessageBodyContentContext) IsMessageBodyContentContext() {}

func NewMessageBodyContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageBodyContentContext {
	var p = new(MessageBodyContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_messageBodyContent

	return p
}

func (s *MessageBodyContentContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageBodyContentContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *MessageBodyContentContext) EnumDefinition() IEnumDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumDefinitionContext)
}

func (s *MessageBodyContentContext) Message() IMessageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageContext)
}

func (s *MessageBodyContentContext) Extend() IExtendContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtendContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtendContext)
}

func (s *MessageBodyContentContext) Option() IOptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *MessageBodyContentContext) Oneof() IOneofContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOneofContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOneofContext)
}

func (s *MessageBodyContentContext) MapField() IMapFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapFieldContext)
}

func (s *MessageBodyContentContext) Reserved() IReservedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReservedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReservedContext)
}

func (s *MessageBodyContentContext) EmptyStatement() IEmptyStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptyStatementContext)
}

func (s *MessageBodyContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageBodyContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageBodyContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMessageBodyContent(s)
	}
}

func (s *MessageBodyContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMessageBodyContent(s)
	}
}

func (s *MessageBodyContentContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitMessageBodyContent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) MessageBodyContent() (localctx IMessageBodyContentContext) {
	this := p
	_ = this

	localctx = NewMessageBodyContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, Protobuf3ParserRULE_messageBodyContent)

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

	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(186)
			p.Field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(187)
			p.EnumDefinition()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(188)
			p.Message()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(189)
			p.Extend()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(190)
			p.Option()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(191)
			p.Oneof()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(192)
			p.MapField()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(193)
			p.Reserved()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(194)
			p.EmptyStatement()
		}

	}

	return localctx
}

// IEnumDefinitionContext is an interface to support dynamic dispatch.
type IEnumDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENUM() antlr.TerminalNode
	EnumName() IEnumNameContext
	EnumBody() IEnumBodyContext

	// IsEnumDefinitionContext differentiates from other interfaces.
	IsEnumDefinitionContext()
}

type EnumDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumDefinitionContext() *EnumDefinitionContext {
	var p = new(EnumDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumDefinition
	return p
}

func (*EnumDefinitionContext) IsEnumDefinitionContext() {}

func NewEnumDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDefinitionContext {
	var p = new(EnumDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumDefinition

	return p
}

func (s *EnumDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDefinitionContext) ENUM() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserENUM, 0)
}

func (s *EnumDefinitionContext) EnumName() IEnumNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumNameContext)
}

func (s *EnumDefinitionContext) EnumBody() IEnumBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *EnumDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEnumDefinition(s)
	}
}

func (s *EnumDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEnumDefinition(s)
	}
}

func (s *EnumDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitEnumDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) EnumDefinition() (localctx IEnumDefinitionContext) {
	this := p
	_ = this

	localctx = NewEnumDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, Protobuf3ParserRULE_enumDefinition)

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
		p.SetState(197)
		p.Match(Protobuf3ParserENUM)
	}
	{
		p.SetState(198)
		p.EnumName()
	}
	{
		p.SetState(199)
		p.EnumBody()
	}

	return localctx
}

// IEnumBodyContext is an interface to support dynamic dispatch.
type IEnumBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllOption() []IOptionContext
	Option(i int) IOptionContext
	AllEnumField() []IEnumFieldContext
	EnumField(i int) IEnumFieldContext
	AllEmptyStatement() []IEmptyStatementContext
	EmptyStatement(i int) IEmptyStatementContext

	// IsEnumBodyContext differentiates from other interfaces.
	IsEnumBodyContext()
}

type EnumBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumBodyContext() *EnumBodyContext {
	var p = new(EnumBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumBody
	return p
}

func (*EnumBodyContext) IsEnumBodyContext() {}

func NewEnumBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumBodyContext {
	var p = new(EnumBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumBody

	return p
}

func (s *EnumBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumBodyContext) AllOption() []IOptionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOptionContext); ok {
			len++
		}
	}

	tst := make([]IOptionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOptionContext); ok {
			tst[i] = t.(IOptionContext)
			i++
		}
	}

	return tst
}

func (s *EnumBodyContext) Option(i int) IOptionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionContext); ok {
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

	return t.(IOptionContext)
}

func (s *EnumBodyContext) AllEnumField() []IEnumFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnumFieldContext); ok {
			len++
		}
	}

	tst := make([]IEnumFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnumFieldContext); ok {
			tst[i] = t.(IEnumFieldContext)
			i++
		}
	}

	return tst
}

func (s *EnumBodyContext) EnumField(i int) IEnumFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumFieldContext); ok {
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

	return t.(IEnumFieldContext)
}

func (s *EnumBodyContext) AllEmptyStatement() []IEmptyStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEmptyStatementContext); ok {
			len++
		}
	}

	tst := make([]IEmptyStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEmptyStatementContext); ok {
			tst[i] = t.(IEmptyStatementContext)
			i++
		}
	}

	return tst
}

func (s *EnumBodyContext) EmptyStatement(i int) IEmptyStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatementContext); ok {
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

	return t.(IEmptyStatementContext)
}

func (s *EnumBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEnumBody(s)
	}
}

func (s *EnumBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEnumBody(s)
	}
}

func (s *EnumBodyContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitEnumBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) EnumBody() (localctx IEnumBodyContext) {
	this := p
	_ = this

	localctx = NewEnumBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, Protobuf3ParserRULE_enumBody)
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
		p.SetState(201)
		p.Match(Protobuf3ParserT__5)
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2251800887427076) != 0 {
		p.SetState(205)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case Protobuf3ParserOPTION:
			{
				p.SetState(202)
				p.Option()
			}

		case Protobuf3ParserIDENT:
			{
				p.SetState(203)
				p.EnumField()
			}

		case Protobuf3ParserT__1:
			{
				p.SetState(204)
				p.EmptyStatement()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(210)
		p.Match(Protobuf3ParserT__6)
	}

	return localctx
}

// IEnumFieldContext is an interface to support dynamic dispatch.
type IEnumFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	INT_LIT() antlr.TerminalNode
	AllEnumValueOption() []IEnumValueOptionContext
	EnumValueOption(i int) IEnumValueOptionContext

	// IsEnumFieldContext differentiates from other interfaces.
	IsEnumFieldContext()
}

type EnumFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumFieldContext() *EnumFieldContext {
	var p = new(EnumFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumField
	return p
}

func (*EnumFieldContext) IsEnumFieldContext() {}

func NewEnumFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumFieldContext {
	var p = new(EnumFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumField

	return p
}

func (s *EnumFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumFieldContext) IDENT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIDENT, 0)
}

func (s *EnumFieldContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT_LIT, 0)
}

func (s *EnumFieldContext) AllEnumValueOption() []IEnumValueOptionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnumValueOptionContext); ok {
			len++
		}
	}

	tst := make([]IEnumValueOptionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnumValueOptionContext); ok {
			tst[i] = t.(IEnumValueOptionContext)
			i++
		}
	}

	return tst
}

func (s *EnumFieldContext) EnumValueOption(i int) IEnumValueOptionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumValueOptionContext); ok {
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

	return t.(IEnumValueOptionContext)
}

func (s *EnumFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEnumField(s)
	}
}

func (s *EnumFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEnumField(s)
	}
}

func (s *EnumFieldContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitEnumField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) EnumField() (localctx IEnumFieldContext) {
	this := p
	_ = this

	localctx = NewEnumFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, Protobuf3ParserRULE_enumField)
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
		p.SetState(212)
		p.Match(Protobuf3ParserIDENT)
	}
	{
		p.SetState(213)
		p.Match(Protobuf3ParserT__0)
	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserT__8 {
		{
			p.SetState(214)
			p.Match(Protobuf3ParserT__8)
		}

	}
	{
		p.SetState(217)
		p.Match(Protobuf3ParserINT_LIT)
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserT__9 {
		{
			p.SetState(218)
			p.Match(Protobuf3ParserT__9)
		}
		{
			p.SetState(219)
			p.EnumValueOption()
		}
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == Protobuf3ParserT__10 {
			{
				p.SetState(220)
				p.Match(Protobuf3ParserT__10)
			}
			{
				p.SetState(221)
				p.EnumValueOption()
			}

			p.SetState(226)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(227)
			p.Match(Protobuf3ParserT__11)
		}

	}
	{
		p.SetState(231)
		p.Match(Protobuf3ParserT__1)
	}

	return localctx
}

// IEnumValueOptionContext is an interface to support dynamic dispatch.
type IEnumValueOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OptionName() IOptionNameContext
	Constant() IConstantContext

	// IsEnumValueOptionContext differentiates from other interfaces.
	IsEnumValueOptionContext()
}

type EnumValueOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueOptionContext() *EnumValueOptionContext {
	var p = new(EnumValueOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumValueOption
	return p
}

func (*EnumValueOptionContext) IsEnumValueOptionContext() {}

func NewEnumValueOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueOptionContext {
	var p = new(EnumValueOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumValueOption

	return p
}

func (s *EnumValueOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueOptionContext) OptionName() IOptionNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *EnumValueOptionContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *EnumValueOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEnumValueOption(s)
	}
}

func (s *EnumValueOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEnumValueOption(s)
	}
}

func (s *EnumValueOptionContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitEnumValueOption(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) EnumValueOption() (localctx IEnumValueOptionContext) {
	this := p
	_ = this

	localctx = NewEnumValueOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, Protobuf3ParserRULE_enumValueOption)

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
		p.SetState(233)
		p.OptionName()
	}
	{
		p.SetState(234)
		p.Match(Protobuf3ParserT__0)
	}
	{
		p.SetState(235)
		p.Constant()
	}

	return localctx
}

// IExtendContext is an interface to support dynamic dispatch.
type IExtendContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXTEND() antlr.TerminalNode
	MessageType() IMessageTypeContext
	Field() IFieldContext
	EmptyStatement() IEmptyStatementContext

	// IsExtendContext differentiates from other interfaces.
	IsExtendContext()
}

type ExtendContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtendContext() *ExtendContext {
	var p = new(ExtendContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_extend
	return p
}

func (*ExtendContext) IsExtendContext() {}

func NewExtendContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtendContext {
	var p = new(ExtendContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_extend

	return p
}

func (s *ExtendContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtendContext) EXTEND() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserEXTEND, 0)
}

func (s *ExtendContext) MessageType() IMessageTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageTypeContext)
}

func (s *ExtendContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *ExtendContext) EmptyStatement() IEmptyStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptyStatementContext)
}

func (s *ExtendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtendContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterExtend(s)
	}
}

func (s *ExtendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitExtend(s)
	}
}

func (s *ExtendContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitExtend(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Extend() (localctx IExtendContext) {
	this := p
	_ = this

	localctx = NewExtendContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, Protobuf3ParserRULE_extend)

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
		p.SetState(237)
		p.Match(Protobuf3ParserEXTEND)
	}
	{
		p.SetState(238)
		p.MessageType()
	}
	{
		p.SetState(239)
		p.Match(Protobuf3ParserT__5)
	}
	p.SetState(242)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserT__4, Protobuf3ParserBOOL, Protobuf3ParserBYTES, Protobuf3ParserDOUBLE, Protobuf3ParserEXTEND, Protobuf3ParserFIXED32, Protobuf3ParserFIXED64, Protobuf3ParserFLOAT, Protobuf3ParserINT32, Protobuf3ParserINT64, Protobuf3ParserMESSAGE, Protobuf3ParserOPTION, Protobuf3ParserPACKAGE, Protobuf3ParserREPEATED, Protobuf3ParserRPC, Protobuf3ParserSERVICE, Protobuf3ParserSFIXED32, Protobuf3ParserSFIXED64, Protobuf3ParserSINT32, Protobuf3ParserSINT64, Protobuf3ParserSTREAM, Protobuf3ParserSTRING, Protobuf3ParserSYNTAX, Protobuf3ParserUINT32, Protobuf3ParserUINT64, Protobuf3ParserWEAK, Protobuf3ParserIDENT:
		{
			p.SetState(240)
			p.Field()
		}

	case Protobuf3ParserT__1:
		{
			p.SetState(241)
			p.EmptyStatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(244)
		p.Match(Protobuf3ParserT__6)
	}

	return localctx
}

// IServiceContext is an interface to support dynamic dispatch.
type IServiceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SERVICE() antlr.TerminalNode
	ServiceName() IServiceNameContext
	AllOption() []IOptionContext
	Option(i int) IOptionContext
	AllRpc() []IRpcContext
	Rpc(i int) IRpcContext
	AllEmptyStatement() []IEmptyStatementContext
	EmptyStatement(i int) IEmptyStatementContext

	// IsServiceContext differentiates from other interfaces.
	IsServiceContext()
}

type ServiceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceContext() *ServiceContext {
	var p = new(ServiceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_service
	return p
}

func (*ServiceContext) IsServiceContext() {}

func NewServiceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceContext {
	var p = new(ServiceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_service

	return p
}

func (s *ServiceContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceContext) SERVICE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSERVICE, 0)
}

func (s *ServiceContext) ServiceName() IServiceNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceNameContext)
}

func (s *ServiceContext) AllOption() []IOptionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOptionContext); ok {
			len++
		}
	}

	tst := make([]IOptionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOptionContext); ok {
			tst[i] = t.(IOptionContext)
			i++
		}
	}

	return tst
}

func (s *ServiceContext) Option(i int) IOptionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionContext); ok {
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

	return t.(IOptionContext)
}

func (s *ServiceContext) AllRpc() []IRpcContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRpcContext); ok {
			len++
		}
	}

	tst := make([]IRpcContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRpcContext); ok {
			tst[i] = t.(IRpcContext)
			i++
		}
	}

	return tst
}

func (s *ServiceContext) Rpc(i int) IRpcContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRpcContext); ok {
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

	return t.(IRpcContext)
}

func (s *ServiceContext) AllEmptyStatement() []IEmptyStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEmptyStatementContext); ok {
			len++
		}
	}

	tst := make([]IEmptyStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEmptyStatementContext); ok {
			tst[i] = t.(IEmptyStatementContext)
			i++
		}
	}

	return tst
}

func (s *ServiceContext) EmptyStatement(i int) IEmptyStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatementContext); ok {
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

	return t.(IEmptyStatementContext)
}

func (s *ServiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterService(s)
	}
}

func (s *ServiceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitService(s)
	}
}

func (s *ServiceContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitService(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Service() (localctx IServiceContext) {
	this := p
	_ = this

	localctx = NewServiceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, Protobuf3ParserRULE_service)
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
		p.SetState(246)
		p.Match(Protobuf3ParserSERVICE)
	}
	{
		p.SetState(247)
		p.ServiceName()
	}
	{
		p.SetState(248)
		p.Match(Protobuf3ParserT__5)
	}
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275951648772) != 0 {
		p.SetState(252)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case Protobuf3ParserOPTION:
			{
				p.SetState(249)
				p.Option()
			}

		case Protobuf3ParserRPC:
			{
				p.SetState(250)
				p.Rpc()
			}

		case Protobuf3ParserT__1:
			{
				p.SetState(251)
				p.EmptyStatement()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(256)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(257)
		p.Match(Protobuf3ParserT__6)
	}

	return localctx
}

// IRpcContext is an interface to support dynamic dispatch.
type IRpcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RPC() antlr.TerminalNode
	RpcName() IRpcNameContext
	AllMessageType() []IMessageTypeContext
	MessageType(i int) IMessageTypeContext
	RETURNS() antlr.TerminalNode
	AllSTREAM() []antlr.TerminalNode
	STREAM(i int) antlr.TerminalNode
	AllOption() []IOptionContext
	Option(i int) IOptionContext
	AllEmptyStatement() []IEmptyStatementContext
	EmptyStatement(i int) IEmptyStatementContext

	// IsRpcContext differentiates from other interfaces.
	IsRpcContext()
}

type RpcContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcContext() *RpcContext {
	var p = new(RpcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_rpc
	return p
}

func (*RpcContext) IsRpcContext() {}

func NewRpcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcContext {
	var p = new(RpcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_rpc

	return p
}

func (s *RpcContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcContext) RPC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRPC, 0)
}

func (s *RpcContext) RpcName() IRpcNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRpcNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRpcNameContext)
}

func (s *RpcContext) AllMessageType() []IMessageTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMessageTypeContext); ok {
			len++
		}
	}

	tst := make([]IMessageTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMessageTypeContext); ok {
			tst[i] = t.(IMessageTypeContext)
			i++
		}
	}

	return tst
}

func (s *RpcContext) MessageType(i int) IMessageTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageTypeContext); ok {
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

	return t.(IMessageTypeContext)
}

func (s *RpcContext) RETURNS() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRETURNS, 0)
}

func (s *RpcContext) AllSTREAM() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserSTREAM)
}

func (s *RpcContext) STREAM(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTREAM, i)
}

func (s *RpcContext) AllOption() []IOptionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOptionContext); ok {
			len++
		}
	}

	tst := make([]IOptionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOptionContext); ok {
			tst[i] = t.(IOptionContext)
			i++
		}
	}

	return tst
}

func (s *RpcContext) Option(i int) IOptionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionContext); ok {
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

	return t.(IOptionContext)
}

func (s *RpcContext) AllEmptyStatement() []IEmptyStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEmptyStatementContext); ok {
			len++
		}
	}

	tst := make([]IEmptyStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEmptyStatementContext); ok {
			tst[i] = t.(IEmptyStatementContext)
			i++
		}
	}

	return tst
}

func (s *RpcContext) EmptyStatement(i int) IEmptyStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatementContext); ok {
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

	return t.(IEmptyStatementContext)
}

func (s *RpcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterRpc(s)
	}
}

func (s *RpcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitRpc(s)
	}
}

func (s *RpcContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitRpc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Rpc() (localctx IRpcContext) {
	this := p
	_ = this

	localctx = NewRpcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, Protobuf3ParserRULE_rpc)
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
		p.SetState(259)
		p.Match(Protobuf3ParserRPC)
	}
	{
		p.SetState(260)
		p.RpcName()
	}
	{
		p.SetState(261)
		p.Match(Protobuf3ParserT__2)
	}
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserSTREAM {
		{
			p.SetState(262)
			p.Match(Protobuf3ParserSTREAM)
		}

	}
	{
		p.SetState(265)
		p.MessageType()
	}
	{
		p.SetState(266)
		p.Match(Protobuf3ParserT__3)
	}
	{
		p.SetState(267)
		p.Match(Protobuf3ParserRETURNS)
	}
	{
		p.SetState(268)
		p.Match(Protobuf3ParserT__2)
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserSTREAM {
		{
			p.SetState(269)
			p.Match(Protobuf3ParserSTREAM)
		}

	}
	{
		p.SetState(272)
		p.MessageType()
	}
	{
		p.SetState(273)
		p.Match(Protobuf3ParserT__3)
	}
	p.SetState(284)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserT__5:
		{
			p.SetState(274)
			p.Match(Protobuf3ParserT__5)
		}
		p.SetState(279)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == Protobuf3ParserT__1 || _la == Protobuf3ParserOPTION {
			p.SetState(277)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case Protobuf3ParserOPTION:
				{
					p.SetState(275)
					p.Option()
				}

			case Protobuf3ParserT__1:
				{
					p.SetState(276)
					p.EmptyStatement()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(281)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(282)
			p.Match(Protobuf3ParserT__6)
		}

	case Protobuf3ParserT__1:
		{
			p.SetState(283)
			p.Match(Protobuf3ParserT__1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IReservedContext is an interface to support dynamic dispatch.
type IReservedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RESERVED() antlr.TerminalNode
	Ranges() IRangesContext
	FieldNames() IFieldNamesContext

	// IsReservedContext differentiates from other interfaces.
	IsReservedContext()
}

type ReservedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReservedContext() *ReservedContext {
	var p = new(ReservedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_reserved
	return p
}

func (*ReservedContext) IsReservedContext() {}

func NewReservedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReservedContext {
	var p = new(ReservedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_reserved

	return p
}

func (s *ReservedContext) GetParser() antlr.Parser { return s.parser }

func (s *ReservedContext) RESERVED() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRESERVED, 0)
}

func (s *ReservedContext) Ranges() IRangesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangesContext)
}

func (s *ReservedContext) FieldNames() IFieldNamesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNamesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNamesContext)
}

func (s *ReservedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReservedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReservedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterReserved(s)
	}
}

func (s *ReservedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitReserved(s)
	}
}

func (s *ReservedContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitReserved(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Reserved() (localctx IReservedContext) {
	this := p
	_ = this

	localctx = NewReservedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, Protobuf3ParserRULE_reserved)

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
		p.SetState(286)
		p.Match(Protobuf3ParserRESERVED)
	}
	p.SetState(289)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserINT_LIT:
		{
			p.SetState(287)
			p.Ranges()
		}

	case Protobuf3ParserSTR_LIT:
		{
			p.SetState(288)
			p.FieldNames()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(291)
		p.Match(Protobuf3ParserT__1)
	}

	return localctx
}

// IRangesContext is an interface to support dynamic dispatch.
type IRangesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRangeRule() []IRangeRuleContext
	RangeRule(i int) IRangeRuleContext

	// IsRangesContext differentiates from other interfaces.
	IsRangesContext()
}

type RangesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangesContext() *RangesContext {
	var p = new(RangesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_ranges
	return p
}

func (*RangesContext) IsRangesContext() {}

func NewRangesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangesContext {
	var p = new(RangesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_ranges

	return p
}

func (s *RangesContext) GetParser() antlr.Parser { return s.parser }

func (s *RangesContext) AllRangeRule() []IRangeRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRangeRuleContext); ok {
			len++
		}
	}

	tst := make([]IRangeRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRangeRuleContext); ok {
			tst[i] = t.(IRangeRuleContext)
			i++
		}
	}

	return tst
}

func (s *RangesContext) RangeRule(i int) IRangeRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangeRuleContext); ok {
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

	return t.(IRangeRuleContext)
}

func (s *RangesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterRanges(s)
	}
}

func (s *RangesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitRanges(s)
	}
}

func (s *RangesContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitRanges(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Ranges() (localctx IRangesContext) {
	this := p
	_ = this

	localctx = NewRangesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, Protobuf3ParserRULE_ranges)
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
		p.SetState(293)
		p.RangeRule()
	}
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserT__10 {
		{
			p.SetState(294)
			p.Match(Protobuf3ParserT__10)
		}
		{
			p.SetState(295)
			p.RangeRule()
		}

		p.SetState(300)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRangeRuleContext is an interface to support dynamic dispatch.
type IRangeRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllINT_LIT() []antlr.TerminalNode
	INT_LIT(i int) antlr.TerminalNode
	TO() antlr.TerminalNode

	// IsRangeRuleContext differentiates from other interfaces.
	IsRangeRuleContext()
}

type RangeRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeRuleContext() *RangeRuleContext {
	var p = new(RangeRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_rangeRule
	return p
}

func (*RangeRuleContext) IsRangeRuleContext() {}

func NewRangeRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeRuleContext {
	var p = new(RangeRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_rangeRule

	return p
}

func (s *RangeRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeRuleContext) AllINT_LIT() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserINT_LIT)
}

func (s *RangeRuleContext) INT_LIT(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT_LIT, i)
}

func (s *RangeRuleContext) TO() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserTO, 0)
}

func (s *RangeRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterRangeRule(s)
	}
}

func (s *RangeRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitRangeRule(s)
	}
}

func (s *RangeRuleContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitRangeRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) RangeRule() (localctx IRangeRuleContext) {
	this := p
	_ = this

	localctx = NewRangeRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, Protobuf3ParserRULE_rangeRule)

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

	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(301)
			p.Match(Protobuf3ParserINT_LIT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(302)
			p.Match(Protobuf3ParserINT_LIT)
		}
		{
			p.SetState(303)
			p.Match(Protobuf3ParserTO)
		}
		{
			p.SetState(304)
			p.Match(Protobuf3ParserINT_LIT)
		}

	}

	return localctx
}

// IFieldNamesContext is an interface to support dynamic dispatch.
type IFieldNamesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSTR_LIT() []antlr.TerminalNode
	STR_LIT(i int) antlr.TerminalNode

	// IsFieldNamesContext differentiates from other interfaces.
	IsFieldNamesContext()
}

type FieldNamesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNamesContext() *FieldNamesContext {
	var p = new(FieldNamesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldNames
	return p
}

func (*FieldNamesContext) IsFieldNamesContext() {}

func NewFieldNamesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNamesContext {
	var p = new(FieldNamesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fieldNames

	return p
}

func (s *FieldNamesContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNamesContext) AllSTR_LIT() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserSTR_LIT)
}

func (s *FieldNamesContext) STR_LIT(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTR_LIT, i)
}

func (s *FieldNamesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNamesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNamesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterFieldNames(s)
	}
}

func (s *FieldNamesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitFieldNames(s)
	}
}

func (s *FieldNamesContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitFieldNames(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) FieldNames() (localctx IFieldNamesContext) {
	this := p
	_ = this

	localctx = NewFieldNamesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, Protobuf3ParserRULE_fieldNames)
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
		p.SetState(307)
		p.Match(Protobuf3ParserSTR_LIT)
	}
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserT__10 {
		{
			p.SetState(308)
			p.Match(Protobuf3ParserT__10)
		}
		{
			p.SetState(309)
			p.Match(Protobuf3ParserSTR_LIT)
		}

		p.SetState(314)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITypeRuleContext is an interface to support dynamic dispatch.
type ITypeRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SimpleType() ISimpleTypeContext
	MessageOrEnumType() IMessageOrEnumTypeContext

	// IsTypeRuleContext differentiates from other interfaces.
	IsTypeRuleContext()
}

type TypeRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeRuleContext() *TypeRuleContext {
	var p = new(TypeRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_typeRule
	return p
}

func (*TypeRuleContext) IsTypeRuleContext() {}

func NewTypeRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeRuleContext {
	var p = new(TypeRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_typeRule

	return p
}

func (s *TypeRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeRuleContext) SimpleType() ISimpleTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleTypeContext)
}

func (s *TypeRuleContext) MessageOrEnumType() IMessageOrEnumTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageOrEnumTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageOrEnumTypeContext)
}

func (s *TypeRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterTypeRule(s)
	}
}

func (s *TypeRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitTypeRule(s)
	}
}

func (s *TypeRuleContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitTypeRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) TypeRule() (localctx ITypeRuleContext) {
	this := p
	_ = this

	localctx = NewTypeRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, Protobuf3ParserRULE_typeRule)

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

	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(315)
			p.SimpleType()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(316)
			p.MessageOrEnumType()
		}

	}

	return localctx
}

// ISimpleTypeContext is an interface to support dynamic dispatch.
type ISimpleTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOUBLE() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	INT32() antlr.TerminalNode
	INT64() antlr.TerminalNode
	UINT32() antlr.TerminalNode
	UINT64() antlr.TerminalNode
	SINT32() antlr.TerminalNode
	SINT64() antlr.TerminalNode
	FIXED32() antlr.TerminalNode
	FIXED64() antlr.TerminalNode
	SFIXED32() antlr.TerminalNode
	SFIXED64() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	STRING() antlr.TerminalNode
	BYTES() antlr.TerminalNode

	// IsSimpleTypeContext differentiates from other interfaces.
	IsSimpleTypeContext()
}

type SimpleTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleTypeContext() *SimpleTypeContext {
	var p = new(SimpleTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_simpleType
	return p
}

func (*SimpleTypeContext) IsSimpleTypeContext() {}

func NewSimpleTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleTypeContext {
	var p = new(SimpleTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_simpleType

	return p
}

func (s *SimpleTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleTypeContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserDOUBLE, 0)
}

func (s *SimpleTypeContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFLOAT, 0)
}

func (s *SimpleTypeContext) INT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT32, 0)
}

func (s *SimpleTypeContext) INT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT64, 0)
}

func (s *SimpleTypeContext) UINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserUINT32, 0)
}

func (s *SimpleTypeContext) UINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserUINT64, 0)
}

func (s *SimpleTypeContext) SINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSINT32, 0)
}

func (s *SimpleTypeContext) SINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSINT64, 0)
}

func (s *SimpleTypeContext) FIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFIXED32, 0)
}

func (s *SimpleTypeContext) FIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFIXED64, 0)
}

func (s *SimpleTypeContext) SFIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSFIXED32, 0)
}

func (s *SimpleTypeContext) SFIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSFIXED64, 0)
}

func (s *SimpleTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserBOOL, 0)
}

func (s *SimpleTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTRING, 0)
}

func (s *SimpleTypeContext) BYTES() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserBYTES, 0)
}

func (s *SimpleTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterSimpleType(s)
	}
}

func (s *SimpleTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitSimpleType(s)
	}
}

func (s *SimpleTypeContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitSimpleType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) SimpleType() (localctx ISimpleTypeContext) {
	this := p
	_ = this

	localctx = NewSimpleTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, Protobuf3ParserRULE_simpleType)
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
		p.SetState(319)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&896102092439552) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFieldNumberContext is an interface to support dynamic dispatch.
type IFieldNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT_LIT() antlr.TerminalNode

	// IsFieldNumberContext differentiates from other interfaces.
	IsFieldNumberContext()
}

type FieldNumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNumberContext() *FieldNumberContext {
	var p = new(FieldNumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldNumber
	return p
}

func (*FieldNumberContext) IsFieldNumberContext() {}

func NewFieldNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNumberContext {
	var p = new(FieldNumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fieldNumber

	return p
}

func (s *FieldNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNumberContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT_LIT, 0)
}

func (s *FieldNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterFieldNumber(s)
	}
}

func (s *FieldNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitFieldNumber(s)
	}
}

func (s *FieldNumberContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitFieldNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) FieldNumber() (localctx IFieldNumberContext) {
	this := p
	_ = this

	localctx = NewFieldNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, Protobuf3ParserRULE_fieldNumber)

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
		p.SetState(321)
		p.Match(Protobuf3ParserINT_LIT)
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TypeRule() ITypeRuleContext
	FieldName() IFieldNameContext
	FieldNumber() IFieldNumberContext
	REPEATED() antlr.TerminalNode
	FieldOptions() IFieldOptionsContext

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) TypeRule() ITypeRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeRuleContext)
}

func (s *FieldContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *FieldContext) FieldNumber() IFieldNumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNumberContext)
}

func (s *FieldContext) REPEATED() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserREPEATED, 0)
}

func (s *FieldContext) FieldOptions() IFieldOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Field() (localctx IFieldContext) {
	this := p
	_ = this

	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, Protobuf3ParserRULE_field)
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
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserREPEATED {
		{
			p.SetState(323)
			p.Match(Protobuf3ParserREPEATED)
		}

	}
	{
		p.SetState(326)
		p.TypeRule()
	}
	{
		p.SetState(327)
		p.FieldName()
	}
	{
		p.SetState(328)
		p.Match(Protobuf3ParserT__0)
	}
	{
		p.SetState(329)
		p.FieldNumber()
	}
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserT__9 {
		{
			p.SetState(330)
			p.Match(Protobuf3ParserT__9)
		}
		{
			p.SetState(331)
			p.FieldOptions()
		}
		{
			p.SetState(332)
			p.Match(Protobuf3ParserT__11)
		}

	}
	{
		p.SetState(336)
		p.Match(Protobuf3ParserT__1)
	}

	return localctx
}

// IFieldOptionsContext is an interface to support dynamic dispatch.
type IFieldOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFieldOption() []IFieldOptionContext
	FieldOption(i int) IFieldOptionContext

	// IsFieldOptionsContext differentiates from other interfaces.
	IsFieldOptionsContext()
}

type FieldOptionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldOptionsContext() *FieldOptionsContext {
	var p = new(FieldOptionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldOptions
	return p
}

func (*FieldOptionsContext) IsFieldOptionsContext() {}

func NewFieldOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOptionsContext {
	var p = new(FieldOptionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fieldOptions

	return p
}

func (s *FieldOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldOptionsContext) AllFieldOption() []IFieldOptionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldOptionContext); ok {
			len++
		}
	}

	tst := make([]IFieldOptionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldOptionContext); ok {
			tst[i] = t.(IFieldOptionContext)
			i++
		}
	}

	return tst
}

func (s *FieldOptionsContext) FieldOption(i int) IFieldOptionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldOptionContext); ok {
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

	return t.(IFieldOptionContext)
}

func (s *FieldOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterFieldOptions(s)
	}
}

func (s *FieldOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitFieldOptions(s)
	}
}

func (s *FieldOptionsContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitFieldOptions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) FieldOptions() (localctx IFieldOptionsContext) {
	this := p
	_ = this

	localctx = NewFieldOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, Protobuf3ParserRULE_fieldOptions)
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
		p.SetState(338)
		p.FieldOption()
	}
	p.SetState(343)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserT__10 {
		{
			p.SetState(339)
			p.Match(Protobuf3ParserT__10)
		}
		{
			p.SetState(340)
			p.FieldOption()
		}

		p.SetState(345)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFieldOptionContext is an interface to support dynamic dispatch.
type IFieldOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OptionName() IOptionNameContext
	Constant() IConstantContext

	// IsFieldOptionContext differentiates from other interfaces.
	IsFieldOptionContext()
}

type FieldOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldOptionContext() *FieldOptionContext {
	var p = new(FieldOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldOption
	return p
}

func (*FieldOptionContext) IsFieldOptionContext() {}

func NewFieldOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOptionContext {
	var p = new(FieldOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fieldOption

	return p
}

func (s *FieldOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldOptionContext) OptionName() IOptionNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *FieldOptionContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *FieldOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterFieldOption(s)
	}
}

func (s *FieldOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitFieldOption(s)
	}
}

func (s *FieldOptionContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitFieldOption(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) FieldOption() (localctx IFieldOptionContext) {
	this := p
	_ = this

	localctx = NewFieldOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, Protobuf3ParserRULE_fieldOption)

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
		p.SetState(346)
		p.OptionName()
	}
	{
		p.SetState(347)
		p.Match(Protobuf3ParserT__0)
	}
	{
		p.SetState(348)
		p.Constant()
	}

	return localctx
}

// IOneofContext is an interface to support dynamic dispatch.
type IOneofContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ONEOF() antlr.TerminalNode
	OneofName() IOneofNameContext
	AllOneofField() []IOneofFieldContext
	OneofField(i int) IOneofFieldContext
	AllEmptyStatement() []IEmptyStatementContext
	EmptyStatement(i int) IEmptyStatementContext

	// IsOneofContext differentiates from other interfaces.
	IsOneofContext()
}

type OneofContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneofContext() *OneofContext {
	var p = new(OneofContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_oneof
	return p
}

func (*OneofContext) IsOneofContext() {}

func NewOneofContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofContext {
	var p = new(OneofContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_oneof

	return p
}

func (s *OneofContext) GetParser() antlr.Parser { return s.parser }

func (s *OneofContext) ONEOF() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserONEOF, 0)
}

func (s *OneofContext) OneofName() IOneofNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOneofNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOneofNameContext)
}

func (s *OneofContext) AllOneofField() []IOneofFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOneofFieldContext); ok {
			len++
		}
	}

	tst := make([]IOneofFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOneofFieldContext); ok {
			tst[i] = t.(IOneofFieldContext)
			i++
		}
	}

	return tst
}

func (s *OneofContext) OneofField(i int) IOneofFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOneofFieldContext); ok {
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

	return t.(IOneofFieldContext)
}

func (s *OneofContext) AllEmptyStatement() []IEmptyStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEmptyStatementContext); ok {
			len++
		}
	}

	tst := make([]IEmptyStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEmptyStatementContext); ok {
			tst[i] = t.(IEmptyStatementContext)
			i++
		}
	}

	return tst
}

func (s *OneofContext) EmptyStatement(i int) IEmptyStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatementContext); ok {
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

	return t.(IEmptyStatementContext)
}

func (s *OneofContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneofContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterOneof(s)
	}
}

func (s *OneofContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitOneof(s)
	}
}

func (s *OneofContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitOneof(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Oneof() (localctx IOneofContext) {
	this := p
	_ = this

	localctx = NewOneofContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, Protobuf3ParserRULE_oneof)
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
		p.SetState(350)
		p.Match(Protobuf3ParserONEOF)
	}
	{
		p.SetState(351)
		p.OneofName()
	}
	{
		p.SetState(352)
		p.Match(Protobuf3ParserT__5)
	}
	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4362590867619876) != 0 {
		p.SetState(355)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case Protobuf3ParserT__4, Protobuf3ParserBOOL, Protobuf3ParserBYTES, Protobuf3ParserDOUBLE, Protobuf3ParserEXTEND, Protobuf3ParserFIXED32, Protobuf3ParserFIXED64, Protobuf3ParserFLOAT, Protobuf3ParserINT32, Protobuf3ParserINT64, Protobuf3ParserMESSAGE, Protobuf3ParserOPTION, Protobuf3ParserPACKAGE, Protobuf3ParserRPC, Protobuf3ParserSERVICE, Protobuf3ParserSFIXED32, Protobuf3ParserSFIXED64, Protobuf3ParserSINT32, Protobuf3ParserSINT64, Protobuf3ParserSTREAM, Protobuf3ParserSTRING, Protobuf3ParserSYNTAX, Protobuf3ParserUINT32, Protobuf3ParserUINT64, Protobuf3ParserWEAK, Protobuf3ParserIDENT:
			{
				p.SetState(353)
				p.OneofField()
			}

		case Protobuf3ParserT__1:
			{
				p.SetState(354)
				p.EmptyStatement()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(359)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(360)
		p.Match(Protobuf3ParserT__6)
	}

	return localctx
}

// IOneofFieldContext is an interface to support dynamic dispatch.
type IOneofFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TypeRule() ITypeRuleContext
	FieldName() IFieldNameContext
	FieldNumber() IFieldNumberContext
	FieldOptions() IFieldOptionsContext

	// IsOneofFieldContext differentiates from other interfaces.
	IsOneofFieldContext()
}

type OneofFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneofFieldContext() *OneofFieldContext {
	var p = new(OneofFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_oneofField
	return p
}

func (*OneofFieldContext) IsOneofFieldContext() {}

func NewOneofFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofFieldContext {
	var p = new(OneofFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_oneofField

	return p
}

func (s *OneofFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *OneofFieldContext) TypeRule() ITypeRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeRuleContext)
}

func (s *OneofFieldContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *OneofFieldContext) FieldNumber() IFieldNumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNumberContext)
}

func (s *OneofFieldContext) FieldOptions() IFieldOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *OneofFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneofFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterOneofField(s)
	}
}

func (s *OneofFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitOneofField(s)
	}
}

func (s *OneofFieldContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitOneofField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) OneofField() (localctx IOneofFieldContext) {
	this := p
	_ = this

	localctx = NewOneofFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, Protobuf3ParserRULE_oneofField)
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
		p.SetState(362)
		p.TypeRule()
	}
	{
		p.SetState(363)
		p.FieldName()
	}
	{
		p.SetState(364)
		p.Match(Protobuf3ParserT__0)
	}
	{
		p.SetState(365)
		p.FieldNumber()
	}
	p.SetState(370)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserT__9 {
		{
			p.SetState(366)
			p.Match(Protobuf3ParserT__9)
		}
		{
			p.SetState(367)
			p.FieldOptions()
		}
		{
			p.SetState(368)
			p.Match(Protobuf3ParserT__11)
		}

	}
	{
		p.SetState(372)
		p.Match(Protobuf3ParserT__1)
	}

	return localctx
}

// IMapFieldContext is an interface to support dynamic dispatch.
type IMapFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MAP() antlr.TerminalNode
	KeyType() IKeyTypeContext
	TypeRule() ITypeRuleContext
	MapName() IMapNameContext
	FieldNumber() IFieldNumberContext
	FieldOptions() IFieldOptionsContext

	// IsMapFieldContext differentiates from other interfaces.
	IsMapFieldContext()
}

type MapFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapFieldContext() *MapFieldContext {
	var p = new(MapFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_mapField
	return p
}

func (*MapFieldContext) IsMapFieldContext() {}

func NewMapFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapFieldContext {
	var p = new(MapFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_mapField

	return p
}

func (s *MapFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *MapFieldContext) MAP() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserMAP, 0)
}

func (s *MapFieldContext) KeyType() IKeyTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyTypeContext)
}

func (s *MapFieldContext) TypeRule() ITypeRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeRuleContext)
}

func (s *MapFieldContext) MapName() IMapNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapNameContext)
}

func (s *MapFieldContext) FieldNumber() IFieldNumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNumberContext)
}

func (s *MapFieldContext) FieldOptions() IFieldOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *MapFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMapField(s)
	}
}

func (s *MapFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMapField(s)
	}
}

func (s *MapFieldContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitMapField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) MapField() (localctx IMapFieldContext) {
	this := p
	_ = this

	localctx = NewMapFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, Protobuf3ParserRULE_mapField)
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
		p.SetState(374)
		p.Match(Protobuf3ParserMAP)
	}
	{
		p.SetState(375)
		p.Match(Protobuf3ParserT__12)
	}
	{
		p.SetState(376)
		p.KeyType()
	}
	{
		p.SetState(377)
		p.Match(Protobuf3ParserT__10)
	}
	{
		p.SetState(378)
		p.TypeRule()
	}
	{
		p.SetState(379)
		p.Match(Protobuf3ParserT__13)
	}
	{
		p.SetState(380)
		p.MapName()
	}
	{
		p.SetState(381)
		p.Match(Protobuf3ParserT__0)
	}
	{
		p.SetState(382)
		p.FieldNumber()
	}
	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserT__9 {
		{
			p.SetState(383)
			p.Match(Protobuf3ParserT__9)
		}
		{
			p.SetState(384)
			p.FieldOptions()
		}
		{
			p.SetState(385)
			p.Match(Protobuf3ParserT__11)
		}

	}
	{
		p.SetState(389)
		p.Match(Protobuf3ParserT__1)
	}

	return localctx
}

// IKeyTypeContext is an interface to support dynamic dispatch.
type IKeyTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT32() antlr.TerminalNode
	INT64() antlr.TerminalNode
	UINT32() antlr.TerminalNode
	UINT64() antlr.TerminalNode
	SINT32() antlr.TerminalNode
	SINT64() antlr.TerminalNode
	FIXED32() antlr.TerminalNode
	FIXED64() antlr.TerminalNode
	SFIXED32() antlr.TerminalNode
	SFIXED64() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsKeyTypeContext differentiates from other interfaces.
	IsKeyTypeContext()
}

type KeyTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyTypeContext() *KeyTypeContext {
	var p = new(KeyTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_keyType
	return p
}

func (*KeyTypeContext) IsKeyTypeContext() {}

func NewKeyTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyTypeContext {
	var p = new(KeyTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_keyType

	return p
}

func (s *KeyTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyTypeContext) INT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT32, 0)
}

func (s *KeyTypeContext) INT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT64, 0)
}

func (s *KeyTypeContext) UINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserUINT32, 0)
}

func (s *KeyTypeContext) UINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserUINT64, 0)
}

func (s *KeyTypeContext) SINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSINT32, 0)
}

func (s *KeyTypeContext) SINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSINT64, 0)
}

func (s *KeyTypeContext) FIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFIXED32, 0)
}

func (s *KeyTypeContext) FIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFIXED64, 0)
}

func (s *KeyTypeContext) SFIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSFIXED32, 0)
}

func (s *KeyTypeContext) SFIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSFIXED64, 0)
}

func (s *KeyTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserBOOL, 0)
}

func (s *KeyTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTRING, 0)
}

func (s *KeyTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterKeyType(s)
	}
}

func (s *KeyTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitKeyType(s)
	}
}

func (s *KeyTypeContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitKeyType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) KeyType() (localctx IKeyTypeContext) {
	this := p
	_ = this

	localctx = NewKeyTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, Protobuf3ParserRULE_keyType)
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
		p.SetState(391)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&896102083657728) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IReservedWordContext is an interface to support dynamic dispatch.
type IReservedWordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXTEND() antlr.TerminalNode
	MESSAGE() antlr.TerminalNode
	OPTION() antlr.TerminalNode
	PACKAGE() antlr.TerminalNode
	RPC() antlr.TerminalNode
	SERVICE() antlr.TerminalNode
	STREAM() antlr.TerminalNode
	STRING() antlr.TerminalNode
	SYNTAX() antlr.TerminalNode
	WEAK() antlr.TerminalNode

	// IsReservedWordContext differentiates from other interfaces.
	IsReservedWordContext()
}

type ReservedWordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReservedWordContext() *ReservedWordContext {
	var p = new(ReservedWordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_reservedWord
	return p
}

func (*ReservedWordContext) IsReservedWordContext() {}

func NewReservedWordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReservedWordContext {
	var p = new(ReservedWordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_reservedWord

	return p
}

func (s *ReservedWordContext) GetParser() antlr.Parser { return s.parser }

func (s *ReservedWordContext) EXTEND() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserEXTEND, 0)
}

func (s *ReservedWordContext) MESSAGE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserMESSAGE, 0)
}

func (s *ReservedWordContext) OPTION() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserOPTION, 0)
}

func (s *ReservedWordContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPACKAGE, 0)
}

func (s *ReservedWordContext) RPC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRPC, 0)
}

func (s *ReservedWordContext) SERVICE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSERVICE, 0)
}

func (s *ReservedWordContext) STREAM() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTREAM, 0)
}

func (s *ReservedWordContext) STRING() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTRING, 0)
}

func (s *ReservedWordContext) SYNTAX() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSYNTAX, 0)
}

func (s *ReservedWordContext) WEAK() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserWEAK, 0)
}

func (s *ReservedWordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReservedWordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReservedWordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterReservedWord(s)
	}
}

func (s *ReservedWordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitReservedWord(s)
	}
}

func (s *ReservedWordContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitReservedWord(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) ReservedWord() (localctx IReservedWordContext) {
	this := p
	_ = this

	localctx = NewReservedWordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, Protobuf3ParserRULE_reservedWord)
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
		p.SetState(393)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1249873333583872) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFullIdentContext is an interface to support dynamic dispatch.
type IFullIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode

	// IsFullIdentContext differentiates from other interfaces.
	IsFullIdentContext()
}

type FullIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFullIdentContext() *FullIdentContext {
	var p = new(FullIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fullIdent
	return p
}

func (*FullIdentContext) IsFullIdentContext() {}

func NewFullIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FullIdentContext {
	var p = new(FullIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fullIdent

	return p
}

func (s *FullIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *FullIdentContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserIDENT)
}

func (s *FullIdentContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIDENT, i)
}

func (s *FullIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FullIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterFullIdent(s)
	}
}

func (s *FullIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitFullIdent(s)
	}
}

func (s *FullIdentContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitFullIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) FullIdent() (localctx IFullIdentContext) {
	this := p
	_ = this

	localctx = NewFullIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, Protobuf3ParserRULE_fullIdent)
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
		p.SetState(395)
		p.Match(Protobuf3ParserIDENT)
	}
	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserT__4 {
		{
			p.SetState(396)
			p.Match(Protobuf3ParserT__4)
		}
		{
			p.SetState(397)
			p.Match(Protobuf3ParserIDENT)
		}

		p.SetState(402)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMessageNameContext is an interface to support dynamic dispatch.
type IMessageNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode

	// IsMessageNameContext differentiates from other interfaces.
	IsMessageNameContext()
}

type MessageNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageNameContext() *MessageNameContext {
	var p = new(MessageNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageName
	return p
}

func (*MessageNameContext) IsMessageNameContext() {}

func NewMessageNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageNameContext {
	var p = new(MessageNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_messageName

	return p
}

func (s *MessageNameContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageNameContext) IDENT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIDENT, 0)
}

func (s *MessageNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMessageName(s)
	}
}

func (s *MessageNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMessageName(s)
	}
}

func (s *MessageNameContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitMessageName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) MessageName() (localctx IMessageNameContext) {
	this := p
	_ = this

	localctx = NewMessageNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, Protobuf3ParserRULE_messageName)

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
		p.SetState(403)
		p.Match(Protobuf3ParserIDENT)
	}

	return localctx
}

// IEnumNameContext is an interface to support dynamic dispatch.
type IEnumNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode

	// IsEnumNameContext differentiates from other interfaces.
	IsEnumNameContext()
}

type EnumNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumNameContext() *EnumNameContext {
	var p = new(EnumNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumName
	return p
}

func (*EnumNameContext) IsEnumNameContext() {}

func NewEnumNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumNameContext {
	var p = new(EnumNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumName

	return p
}

func (s *EnumNameContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumNameContext) IDENT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIDENT, 0)
}

func (s *EnumNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEnumName(s)
	}
}

func (s *EnumNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEnumName(s)
	}
}

func (s *EnumNameContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitEnumName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) EnumName() (localctx IEnumNameContext) {
	this := p
	_ = this

	localctx = NewEnumNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, Protobuf3ParserRULE_enumName)

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
		p.SetState(405)
		p.Match(Protobuf3ParserIDENT)
	}

	return localctx
}

// IMessageOrEnumNameContext is an interface to support dynamic dispatch.
type IMessageOrEnumNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode

	// IsMessageOrEnumNameContext differentiates from other interfaces.
	IsMessageOrEnumNameContext()
}

type MessageOrEnumNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageOrEnumNameContext() *MessageOrEnumNameContext {
	var p = new(MessageOrEnumNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageOrEnumName
	return p
}

func (*MessageOrEnumNameContext) IsMessageOrEnumNameContext() {}

func NewMessageOrEnumNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageOrEnumNameContext {
	var p = new(MessageOrEnumNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_messageOrEnumName

	return p
}

func (s *MessageOrEnumNameContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageOrEnumNameContext) IDENT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIDENT, 0)
}

func (s *MessageOrEnumNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageOrEnumNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageOrEnumNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMessageOrEnumName(s)
	}
}

func (s *MessageOrEnumNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMessageOrEnumName(s)
	}
}

func (s *MessageOrEnumNameContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitMessageOrEnumName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) MessageOrEnumName() (localctx IMessageOrEnumNameContext) {
	this := p
	_ = this

	localctx = NewMessageOrEnumNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, Protobuf3ParserRULE_messageOrEnumName)

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
		p.SetState(407)
		p.Match(Protobuf3ParserIDENT)
	}

	return localctx
}

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	ReservedWord() IReservedWordContext

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldName
	return p
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) IDENT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIDENT, 0)
}

func (s *FieldNameContext) ReservedWord() IReservedWordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReservedWordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReservedWordContext)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterFieldName(s)
	}
}

func (s *FieldNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitFieldName(s)
	}
}

func (s *FieldNameContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitFieldName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) FieldName() (localctx IFieldNameContext) {
	this := p
	_ = this

	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, Protobuf3ParserRULE_fieldName)

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

	p.SetState(411)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(409)
			p.Match(Protobuf3ParserIDENT)
		}

	case Protobuf3ParserEXTEND, Protobuf3ParserMESSAGE, Protobuf3ParserOPTION, Protobuf3ParserPACKAGE, Protobuf3ParserRPC, Protobuf3ParserSERVICE, Protobuf3ParserSTREAM, Protobuf3ParserSTRING, Protobuf3ParserSYNTAX, Protobuf3ParserWEAK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(410)
			p.ReservedWord()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOneofNameContext is an interface to support dynamic dispatch.
type IOneofNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode

	// IsOneofNameContext differentiates from other interfaces.
	IsOneofNameContext()
}

type OneofNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneofNameContext() *OneofNameContext {
	var p = new(OneofNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_oneofName
	return p
}

func (*OneofNameContext) IsOneofNameContext() {}

func NewOneofNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofNameContext {
	var p = new(OneofNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_oneofName

	return p
}

func (s *OneofNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OneofNameContext) IDENT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIDENT, 0)
}

func (s *OneofNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneofNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterOneofName(s)
	}
}

func (s *OneofNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitOneofName(s)
	}
}

func (s *OneofNameContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitOneofName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) OneofName() (localctx IOneofNameContext) {
	this := p
	_ = this

	localctx = NewOneofNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, Protobuf3ParserRULE_oneofName)

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
		p.SetState(413)
		p.Match(Protobuf3ParserIDENT)
	}

	return localctx
}

// IMapNameContext is an interface to support dynamic dispatch.
type IMapNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode

	// IsMapNameContext differentiates from other interfaces.
	IsMapNameContext()
}

type MapNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapNameContext() *MapNameContext {
	var p = new(MapNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_mapName
	return p
}

func (*MapNameContext) IsMapNameContext() {}

func NewMapNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapNameContext {
	var p = new(MapNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_mapName

	return p
}

func (s *MapNameContext) GetParser() antlr.Parser { return s.parser }

func (s *MapNameContext) IDENT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIDENT, 0)
}

func (s *MapNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMapName(s)
	}
}

func (s *MapNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMapName(s)
	}
}

func (s *MapNameContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitMapName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) MapName() (localctx IMapNameContext) {
	this := p
	_ = this

	localctx = NewMapNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, Protobuf3ParserRULE_mapName)

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
		p.SetState(415)
		p.Match(Protobuf3ParserIDENT)
	}

	return localctx
}

// IServiceNameContext is an interface to support dynamic dispatch.
type IServiceNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode

	// IsServiceNameContext differentiates from other interfaces.
	IsServiceNameContext()
}

type ServiceNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceNameContext() *ServiceNameContext {
	var p = new(ServiceNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_serviceName
	return p
}

func (*ServiceNameContext) IsServiceNameContext() {}

func NewServiceNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceNameContext {
	var p = new(ServiceNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_serviceName

	return p
}

func (s *ServiceNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceNameContext) IDENT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIDENT, 0)
}

func (s *ServiceNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterServiceName(s)
	}
}

func (s *ServiceNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitServiceName(s)
	}
}

func (s *ServiceNameContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitServiceName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) ServiceName() (localctx IServiceNameContext) {
	this := p
	_ = this

	localctx = NewServiceNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, Protobuf3ParserRULE_serviceName)

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
		p.SetState(417)
		p.Match(Protobuf3ParserIDENT)
	}

	return localctx
}

// IRpcNameContext is an interface to support dynamic dispatch.
type IRpcNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode

	// IsRpcNameContext differentiates from other interfaces.
	IsRpcNameContext()
}

type RpcNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcNameContext() *RpcNameContext {
	var p = new(RpcNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_rpcName
	return p
}

func (*RpcNameContext) IsRpcNameContext() {}

func NewRpcNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcNameContext {
	var p = new(RpcNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_rpcName

	return p
}

func (s *RpcNameContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcNameContext) IDENT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIDENT, 0)
}

func (s *RpcNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterRpcName(s)
	}
}

func (s *RpcNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitRpcName(s)
	}
}

func (s *RpcNameContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitRpcName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) RpcName() (localctx IRpcNameContext) {
	this := p
	_ = this

	localctx = NewRpcNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, Protobuf3ParserRULE_rpcName)

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
		p.SetState(419)
		p.Match(Protobuf3ParserIDENT)
	}

	return localctx
}

// IMessageTypeContext is an interface to support dynamic dispatch.
type IMessageTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MessageName() IMessageNameContext
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode

	// IsMessageTypeContext differentiates from other interfaces.
	IsMessageTypeContext()
}

type MessageTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageTypeContext() *MessageTypeContext {
	var p = new(MessageTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageType
	return p
}

func (*MessageTypeContext) IsMessageTypeContext() {}

func NewMessageTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageTypeContext {
	var p = new(MessageTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_messageType

	return p
}

func (s *MessageTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageTypeContext) MessageName() IMessageNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageNameContext)
}

func (s *MessageTypeContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserIDENT)
}

func (s *MessageTypeContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIDENT, i)
}

func (s *MessageTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMessageType(s)
	}
}

func (s *MessageTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMessageType(s)
	}
}

func (s *MessageTypeContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitMessageType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) MessageType() (localctx IMessageTypeContext) {
	this := p
	_ = this

	localctx = NewMessageTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, Protobuf3ParserRULE_messageType)
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
	p.SetState(422)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserT__4 {
		{
			p.SetState(421)
			p.Match(Protobuf3ParserT__4)
		}

	}
	p.SetState(428)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(424)
				p.Match(Protobuf3ParserIDENT)
			}
			{
				p.SetState(425)
				p.Match(Protobuf3ParserT__4)
			}

		}
		p.SetState(430)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())
	}
	{
		p.SetState(431)
		p.MessageName()
	}

	return localctx
}

// IMessageOrEnumTypeContext is an interface to support dynamic dispatch.
type IMessageOrEnumTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MessageOrEnumName() IMessageOrEnumNameContext
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	AllReservedWord() []IReservedWordContext
	ReservedWord(i int) IReservedWordContext

	// IsMessageOrEnumTypeContext differentiates from other interfaces.
	IsMessageOrEnumTypeContext()
}

type MessageOrEnumTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageOrEnumTypeContext() *MessageOrEnumTypeContext {
	var p = new(MessageOrEnumTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageOrEnumType
	return p
}

func (*MessageOrEnumTypeContext) IsMessageOrEnumTypeContext() {}

func NewMessageOrEnumTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageOrEnumTypeContext {
	var p = new(MessageOrEnumTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_messageOrEnumType

	return p
}

func (s *MessageOrEnumTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageOrEnumTypeContext) MessageOrEnumName() IMessageOrEnumNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageOrEnumNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageOrEnumNameContext)
}

func (s *MessageOrEnumTypeContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserIDENT)
}

func (s *MessageOrEnumTypeContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIDENT, i)
}

func (s *MessageOrEnumTypeContext) AllReservedWord() []IReservedWordContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IReservedWordContext); ok {
			len++
		}
	}

	tst := make([]IReservedWordContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IReservedWordContext); ok {
			tst[i] = t.(IReservedWordContext)
			i++
		}
	}

	return tst
}

func (s *MessageOrEnumTypeContext) ReservedWord(i int) IReservedWordContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReservedWordContext); ok {
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

	return t.(IReservedWordContext)
}

func (s *MessageOrEnumTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageOrEnumTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageOrEnumTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterMessageOrEnumType(s)
	}
}

func (s *MessageOrEnumTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitMessageOrEnumType(s)
	}
}

func (s *MessageOrEnumTypeContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitMessageOrEnumType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) MessageOrEnumType() (localctx IMessageOrEnumTypeContext) {
	this := p
	_ = this

	localctx = NewMessageOrEnumTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, Protobuf3ParserRULE_messageOrEnumType)
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
	p.SetState(434)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserT__4 {
		{
			p.SetState(433)
			p.Match(Protobuf3ParserT__4)
		}

	}
	p.SetState(443)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(438)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case Protobuf3ParserIDENT:
				{
					p.SetState(436)
					p.Match(Protobuf3ParserIDENT)
				}

			case Protobuf3ParserEXTEND, Protobuf3ParserMESSAGE, Protobuf3ParserOPTION, Protobuf3ParserPACKAGE, Protobuf3ParserRPC, Protobuf3ParserSERVICE, Protobuf3ParserSTREAM, Protobuf3ParserSTRING, Protobuf3ParserSYNTAX, Protobuf3ParserWEAK:
				{
					p.SetState(437)
					p.ReservedWord()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}
			{
				p.SetState(440)
				p.Match(Protobuf3ParserT__4)
			}

		}
		p.SetState(445)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext())
	}
	{
		p.SetState(446)
		p.MessageOrEnumName()
	}

	return localctx
}

// IEmptyStatementContext is an interface to support dynamic dispatch.
type IEmptyStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsEmptyStatementContext differentiates from other interfaces.
	IsEmptyStatementContext()
}

type EmptyStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmptyStatementContext() *EmptyStatementContext {
	var p = new(EmptyStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_emptyStatement
	return p
}

func (*EmptyStatementContext) IsEmptyStatementContext() {}

func NewEmptyStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyStatementContext {
	var p = new(EmptyStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_emptyStatement

	return p
}

func (s *EmptyStatementContext) GetParser() antlr.Parser { return s.parser }
func (s *EmptyStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptyStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterEmptyStatement(s)
	}
}

func (s *EmptyStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitEmptyStatement(s)
	}
}

func (s *EmptyStatementContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitEmptyStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) EmptyStatement() (localctx IEmptyStatementContext) {
	this := p
	_ = this

	localctx = NewEmptyStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, Protobuf3ParserRULE_emptyStatement)

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
		p.SetState(448)
		p.Match(Protobuf3ParserT__1)
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FullIdent() IFullIdentContext
	INT_LIT() antlr.TerminalNode
	FLOAT_LIT() antlr.TerminalNode
	STR_LIT() antlr.TerminalNode
	BOOL_LIT() antlr.TerminalNode

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) FullIdent() IFullIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFullIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *ConstantContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT_LIT, 0)
}

func (s *ConstantContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFLOAT_LIT, 0)
}

func (s *ConstantContext) STR_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTR_LIT, 0)
}

func (s *ConstantContext) BOOL_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserBOOL_LIT, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Protobuf3Listener); ok {
		listenerT.ExitConstant(s)
	}
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) any {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Constant() (localctx IConstantContext) {
	this := p
	_ = this

	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, Protobuf3ParserRULE_constant)
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

	p.SetState(461)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(450)
			p.FullIdent()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(452)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == Protobuf3ParserT__8 || _la == Protobuf3ParserT__14 {
			{
				p.SetState(451)
				_la = p.GetTokenStream().LA(1)

				if !(_la == Protobuf3ParserT__8 || _la == Protobuf3ParserT__14) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(454)
			p.Match(Protobuf3ParserINT_LIT)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(456)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == Protobuf3ParserT__8 || _la == Protobuf3ParserT__14 {
			{
				p.SetState(455)
				_la = p.GetTokenStream().LA(1)

				if !(_la == Protobuf3ParserT__8 || _la == Protobuf3ParserT__14) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(458)
			p.Match(Protobuf3ParserFLOAT_LIT)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(459)
			p.Match(Protobuf3ParserSTR_LIT)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(460)
			p.Match(Protobuf3ParserBOOL_LIT)
		}

	}

	return localctx
}
