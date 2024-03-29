// Code generated from Expr.g4 by ANTLR 4.12.0. DO NOT EDIT.

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

type ExprLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var exprlexerLexerStaticData struct {
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

func exprlexerLexerInit() {
	staticData := &exprlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
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
		"SKIP_", "UNKNOWN_CHAR",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "STRING", "NUMBER",
		"INTEGER", "AND", "FALSE", "IN", "IS", "NONE", "NOT", "OR", "TRUE",
		"NAME", "STRING_LITERAL", "BYTES_LITERAL", "DECIMAL_INTEGER", "OCT_INTEGER",
		"HEX_INTEGER", "BIN_INTEGER", "FLOAT_NUMBER", "IMAG_NUMBER", "SKIP_",
		"UNKNOWN_CHAR", "SHORT_STRING", "LONG_STRING", "LONG_STRING_ITEM", "LONG_STRING_CHAR",
		"STRING_ESCAPE_SEQ", "NON_ZERO_DIGIT", "DIGIT", "OCT_DIGIT", "HEX_DIGIT",
		"BIN_DIGIT", "POINT_FLOAT", "EXPONENT_FLOAT", "INT_PART", "FRACTION",
		"EXPONENT", "SHORT_BYTES", "LONG_BYTES", "LONG_BYTES_ITEM", "SHORT_BYTES_CHAR_NO_SINGLE_QUOTE",
		"SHORT_BYTES_CHAR_NO_DOUBLE_QUOTE", "LONG_BYTES_CHAR", "BYTES_ESCAPE_SEQ",
		"SPACES", "WS", "COMMENT", "LINE_JOINING", "ID_START", "ID_CONTINUE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 53, 597, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
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
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2,
		73, 7, 73, 2, 74, 7, 74, 2, 75, 7, 75, 2, 76, 7, 76, 2, 77, 7, 77, 2, 78,
		7, 78, 2, 79, 7, 79, 2, 80, 7, 80, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16,
		1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1,
		30, 1, 30, 1, 31, 1, 31, 3, 31, 239, 8, 31, 1, 32, 1, 32, 1, 32, 3, 32,
		244, 8, 32, 1, 33, 1, 33, 1, 33, 1, 33, 3, 33, 250, 8, 33, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1,
		36, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39,
		1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1,
		42, 1, 42, 5, 42, 287, 8, 42, 10, 42, 12, 42, 290, 9, 42, 1, 43, 1, 43,
		1, 43, 1, 43, 1, 43, 3, 43, 297, 8, 43, 1, 43, 1, 43, 3, 43, 301, 8, 43,
		1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 3, 44, 308, 8, 44, 1, 44, 1, 44, 3,
		44, 312, 8, 44, 1, 45, 1, 45, 3, 45, 316, 8, 45, 1, 45, 5, 45, 319, 8,
		45, 10, 45, 12, 45, 322, 9, 45, 1, 45, 4, 45, 325, 8, 45, 11, 45, 12, 45,
		326, 3, 45, 329, 8, 45, 1, 46, 1, 46, 1, 46, 4, 46, 334, 8, 46, 11, 46,
		12, 46, 335, 1, 47, 1, 47, 1, 47, 4, 47, 341, 8, 47, 11, 47, 12, 47, 342,
		1, 48, 1, 48, 1, 48, 4, 48, 348, 8, 48, 11, 48, 12, 48, 349, 1, 49, 1,
		49, 3, 49, 354, 8, 49, 1, 50, 1, 50, 3, 50, 358, 8, 50, 1, 50, 1, 50, 1,
		51, 1, 51, 1, 51, 3, 51, 365, 8, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 53,
		1, 53, 1, 53, 5, 53, 374, 8, 53, 10, 53, 12, 53, 377, 9, 53, 1, 53, 1,
		53, 1, 53, 1, 53, 5, 53, 383, 8, 53, 10, 53, 12, 53, 386, 9, 53, 1, 53,
		3, 53, 389, 8, 53, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 5, 54, 396, 8, 54,
		10, 54, 12, 54, 399, 9, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1,
		54, 1, 54, 5, 54, 409, 8, 54, 10, 54, 12, 54, 412, 9, 54, 1, 54, 1, 54,
		1, 54, 3, 54, 417, 8, 54, 1, 55, 1, 55, 3, 55, 421, 8, 55, 1, 56, 1, 56,
		1, 57, 1, 57, 1, 57, 1, 57, 3, 57, 429, 8, 57, 1, 58, 1, 58, 1, 59, 1,
		59, 1, 60, 1, 60, 1, 61, 1, 61, 1, 62, 1, 62, 1, 63, 3, 63, 442, 8, 63,
		1, 63, 1, 63, 1, 63, 1, 63, 3, 63, 448, 8, 63, 1, 64, 1, 64, 3, 64, 452,
		8, 64, 1, 64, 1, 64, 1, 65, 1, 65, 3, 65, 458, 8, 65, 1, 65, 5, 65, 461,
		8, 65, 10, 65, 12, 65, 464, 9, 65, 1, 66, 1, 66, 1, 66, 3, 66, 469, 8,
		66, 1, 66, 5, 66, 472, 8, 66, 10, 66, 12, 66, 475, 9, 66, 1, 67, 1, 67,
		3, 67, 479, 8, 67, 1, 67, 1, 67, 3, 67, 483, 8, 67, 1, 67, 5, 67, 486,
		8, 67, 10, 67, 12, 67, 489, 9, 67, 1, 68, 1, 68, 1, 68, 5, 68, 494, 8,
		68, 10, 68, 12, 68, 497, 9, 68, 1, 68, 1, 68, 1, 68, 1, 68, 5, 68, 503,
		8, 68, 10, 68, 12, 68, 506, 9, 68, 1, 68, 3, 68, 509, 8, 68, 1, 69, 1,
		69, 1, 69, 1, 69, 1, 69, 5, 69, 516, 8, 69, 10, 69, 12, 69, 519, 9, 69,
		1, 69, 1, 69, 1, 69, 1, 69, 1, 69, 1, 69, 1, 69, 1, 69, 5, 69, 529, 8,
		69, 10, 69, 12, 69, 532, 9, 69, 1, 69, 1, 69, 1, 69, 3, 69, 537, 8, 69,
		1, 70, 1, 70, 3, 70, 541, 8, 70, 1, 71, 3, 71, 544, 8, 71, 1, 72, 3, 72,
		547, 8, 72, 1, 73, 3, 73, 550, 8, 73, 1, 74, 1, 74, 1, 74, 1, 75, 4, 75,
		556, 8, 75, 11, 75, 12, 75, 557, 1, 76, 1, 76, 3, 76, 562, 8, 76, 1, 76,
		1, 76, 3, 76, 566, 8, 76, 1, 76, 3, 76, 569, 8, 76, 3, 76, 571, 8, 76,
		1, 77, 1, 77, 5, 77, 575, 8, 77, 10, 77, 12, 77, 578, 9, 77, 1, 78, 1,
		78, 3, 78, 582, 8, 78, 1, 78, 3, 78, 585, 8, 78, 1, 78, 1, 78, 3, 78, 589,
		8, 78, 1, 79, 3, 79, 592, 8, 79, 1, 80, 1, 80, 3, 80, 596, 8, 80, 4, 397,
		410, 517, 530, 0, 81, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8,
		17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35,
		71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44,
		89, 45, 91, 46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105,
		53, 107, 0, 109, 0, 111, 0, 113, 0, 115, 0, 117, 0, 119, 0, 121, 0, 123,
		0, 125, 0, 127, 0, 129, 0, 131, 0, 133, 0, 135, 0, 137, 0, 139, 0, 141,
		0, 143, 0, 145, 0, 147, 0, 149, 0, 151, 0, 153, 0, 155, 0, 157, 0, 159,
		0, 161, 0, 1, 0, 25, 6, 0, 70, 70, 82, 82, 85, 85, 102, 102, 114, 114,
		117, 117, 2, 0, 70, 70, 102, 102, 2, 0, 82, 82, 114, 114, 2, 0, 66, 66,
		98, 98, 2, 0, 79, 79, 111, 111, 2, 0, 88, 88, 120, 120, 2, 0, 74, 74, 106,
		106, 4, 0, 10, 10, 12, 13, 39, 39, 92, 92, 4, 0, 10, 10, 12, 13, 34, 34,
		92, 92, 1, 0, 92, 92, 1, 0, 49, 57, 1, 0, 48, 57, 1, 0, 48, 55, 3, 0, 48,
		57, 65, 70, 97, 102, 1, 0, 48, 49, 2, 0, 69, 69, 101, 101, 2, 0, 43, 43,
		45, 45, 5, 0, 0, 9, 11, 12, 14, 38, 40, 91, 93, 127, 5, 0, 0, 9, 11, 12,
		14, 33, 35, 91, 93, 127, 2, 0, 0, 91, 93, 127, 1, 0, 0, 127, 2, 0, 9, 9,
		32, 32, 2, 0, 10, 10, 12, 13, 295, 0, 65, 90, 95, 95, 97, 122, 170, 170,
		181, 181, 186, 186, 192, 214, 216, 246, 248, 577, 592, 705, 710, 721, 736,
		740, 750, 750, 890, 890, 902, 902, 904, 906, 908, 908, 910, 929, 931, 974,
		976, 1013, 1015, 1153, 1162, 1230, 1232, 1273, 1280, 1295, 1329, 1366,
		1369, 1369, 1377, 1415, 1488, 1514, 1520, 1522, 1569, 1594, 1600, 1610,
		1646, 1647, 1649, 1747, 1749, 1749, 1765, 1766, 1774, 1775, 1786, 1788,
		1791, 1791, 1808, 1808, 1810, 1839, 1869, 1901, 1920, 1957, 1969, 1969,
		2308, 2361, 2365, 2365, 2384, 2384, 2392, 2401, 2429, 2429, 2437, 2444,
		2447, 2448, 2451, 2472, 2474, 2480, 2482, 2482, 2486, 2489, 2493, 2493,
		2510, 2510, 2524, 2525, 2527, 2529, 2544, 2545, 2565, 2570, 2575, 2576,
		2579, 2600, 2602, 2608, 2610, 2611, 2613, 2614, 2616, 2617, 2649, 2652,
		2654, 2654, 2674, 2676, 2693, 2701, 2703, 2705, 2707, 2728, 2730, 2736,
		2738, 2739, 2741, 2745, 2749, 2749, 2768, 2768, 2784, 2785, 2821, 2828,
		2831, 2832, 2835, 2856, 2858, 2864, 2866, 2867, 2869, 2873, 2877, 2877,
		2908, 2909, 2911, 2913, 2929, 2929, 2947, 2947, 2949, 2954, 2958, 2960,
		2962, 2965, 2969, 2970, 2972, 2972, 2974, 2975, 2979, 2980, 2984, 2986,
		2990, 3001, 3077, 3084, 3086, 3088, 3090, 3112, 3114, 3123, 3125, 3129,
		3168, 3169, 3205, 3212, 3214, 3216, 3218, 3240, 3242, 3251, 3253, 3257,
		3261, 3261, 3294, 3294, 3296, 3297, 3333, 3340, 3342, 3344, 3346, 3368,
		3370, 3385, 3424, 3425, 3461, 3478, 3482, 3505, 3507, 3515, 3517, 3517,
		3520, 3526, 3585, 3632, 3634, 3635, 3648, 3654, 3713, 3714, 3716, 3716,
		3719, 3720, 3722, 3722, 3725, 3725, 3732, 3735, 3737, 3743, 3745, 3747,
		3749, 3749, 3751, 3751, 3754, 3755, 3757, 3760, 3762, 3763, 3773, 3773,
		3776, 3780, 3782, 3782, 3804, 3805, 3840, 3840, 3904, 3911, 3913, 3946,
		3976, 3979, 4096, 4129, 4131, 4135, 4137, 4138, 4176, 4181, 4256, 4293,
		4304, 4346, 4348, 4348, 4352, 4441, 4447, 4514, 4520, 4601, 4608, 4680,
		4682, 4685, 4688, 4694, 4696, 4696, 4698, 4701, 4704, 4744, 4746, 4749,
		4752, 4784, 4786, 4789, 4792, 4798, 4800, 4800, 4802, 4805, 4808, 4822,
		4824, 4880, 4882, 4885, 4888, 4954, 4992, 5007, 5024, 5108, 5121, 5740,
		5743, 5750, 5761, 5786, 5792, 5866, 5870, 5872, 5888, 5900, 5902, 5905,
		5920, 5937, 5952, 5969, 5984, 5996, 5998, 6000, 6016, 6067, 6103, 6103,
		6108, 6108, 6176, 6263, 6272, 6312, 6400, 6428, 6480, 6509, 6512, 6516,
		6528, 6569, 6593, 6599, 6656, 6678, 7424, 7615, 7680, 7835, 7840, 7929,
		7936, 7957, 7960, 7965, 7968, 8005, 8008, 8013, 8016, 8023, 8025, 8025,
		8027, 8027, 8029, 8029, 8031, 8061, 8064, 8116, 8118, 8124, 8126, 8126,
		8130, 8132, 8134, 8140, 8144, 8147, 8150, 8155, 8160, 8172, 8178, 8180,
		8182, 8188, 8305, 8305, 8319, 8319, 8336, 8340, 8450, 8450, 8455, 8455,
		8458, 8467, 8469, 8469, 8472, 8477, 8484, 8484, 8486, 8486, 8488, 8488,
		8490, 8497, 8499, 8505, 8508, 8511, 8517, 8521, 8544, 8579, 11264, 11310,
		11312, 11358, 11392, 11492, 11520, 11557, 11568, 11621, 11631, 11631, 11648,
		11670, 11680, 11686, 11688, 11694, 11696, 11702, 11704, 11710, 11712, 11718,
		11720, 11726, 11728, 11734, 11736, 11742, 12293, 12295, 12321, 12329, 12337,
		12341, 12344, 12348, 12353, 12438, 12443, 12447, 12449, 12538, 12540, 12543,
		12549, 12588, 12593, 12686, 12704, 12727, 12784, 12799, 13312, 19893, 19968,
		40891, 40960, 42124, 43008, 43009, 43011, 43013, 43015, 43018, 43020, 43042,
		44032, 55203, 63744, 64045, 64048, 64106, 64112, 64217, 64256, 64262, 64275,
		64279, 64285, 64285, 64287, 64296, 64298, 64310, 64312, 64316, 64318, 64318,
		64320, 64321, 64323, 64324, 64326, 64433, 64467, 64829, 64848, 64911, 64914,
		64967, 65008, 65019, 65136, 65140, 65142, 65276, 65313, 65338, 65345, 65370,
		65382, 65470, 65474, 65479, 65482, 65487, 65490, 65495, 65498, 65500, 148,
		0, 48, 57, 768, 879, 1155, 1158, 1425, 1465, 1467, 1469, 1471, 1471, 1473,
		1474, 1476, 1477, 1479, 1479, 1552, 1557, 1611, 1630, 1632, 1641, 1648,
		1648, 1750, 1756, 1759, 1764, 1767, 1768, 1770, 1773, 1776, 1785, 1809,
		1809, 1840, 1866, 1958, 1968, 2305, 2307, 2364, 2364, 2366, 2381, 2385,
		2388, 2402, 2403, 2406, 2415, 2433, 2435, 2492, 2492, 2494, 2500, 2503,
		2504, 2507, 2509, 2519, 2519, 2530, 2531, 2534, 2543, 2561, 2563, 2620,
		2620, 2622, 2626, 2631, 2632, 2635, 2637, 2662, 2673, 2689, 2691, 2748,
		2748, 2750, 2757, 2759, 2761, 2763, 2765, 2786, 2787, 2790, 2799, 2817,
		2819, 2876, 2876, 2878, 2883, 2887, 2888, 2891, 2893, 2902, 2903, 2918,
		2927, 2946, 2946, 3006, 3010, 3014, 3016, 3018, 3021, 3031, 3031, 3046,
		3055, 3073, 3075, 3134, 3140, 3142, 3144, 3146, 3149, 3157, 3158, 3174,
		3183, 3202, 3203, 3260, 3260, 3262, 3268, 3270, 3272, 3274, 3277, 3285,
		3286, 3302, 3311, 3330, 3331, 3390, 3395, 3398, 3400, 3402, 3405, 3415,
		3415, 3430, 3439, 3458, 3459, 3530, 3530, 3535, 3540, 3542, 3542, 3544,
		3551, 3570, 3571, 3633, 3633, 3636, 3642, 3655, 3662, 3664, 3673, 3761,
		3761, 3764, 3769, 3771, 3772, 3784, 3789, 3792, 3801, 3864, 3865, 3872,
		3881, 3893, 3893, 3895, 3895, 3897, 3897, 3902, 3903, 3953, 3972, 3974,
		3975, 3984, 3991, 3993, 4028, 4038, 4038, 4140, 4146, 4150, 4153, 4160,
		4169, 4182, 4185, 4959, 4959, 4969, 4977, 5906, 5908, 5938, 5940, 5970,
		5971, 6002, 6003, 6070, 6099, 6109, 6109, 6112, 6121, 6155, 6157, 6160,
		6169, 6313, 6313, 6432, 6443, 6448, 6459, 6470, 6479, 6576, 6592, 6600,
		6601, 6608, 6617, 6679, 6683, 7616, 7619, 8255, 8256, 8276, 8276, 8400,
		8412, 8417, 8417, 8421, 8427, 12330, 12335, 12441, 12442, 43010, 43010,
		43014, 43014, 43019, 43019, 43043, 43047, 64286, 64286, 65024, 65039, 65056,
		65059, 65075, 65076, 65101, 65103, 65296, 65305, 65343, 65343, 632, 0,
		1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0,
		9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0,
		0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0,
		0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0,
		0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1,
		0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47,
		1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0,
		55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0,
		0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0,
		0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0,
		0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1,
		0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93,
		1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0,
		101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 1, 163, 1, 0,
		0, 0, 3, 165, 1, 0, 0, 0, 5, 167, 1, 0, 0, 0, 7, 170, 1, 0, 0, 0, 9, 173,
		1, 0, 0, 0, 11, 176, 1, 0, 0, 0, 13, 179, 1, 0, 0, 0, 15, 182, 1, 0, 0,
		0, 17, 184, 1, 0, 0, 0, 19, 186, 1, 0, 0, 0, 21, 188, 1, 0, 0, 0, 23, 191,
		1, 0, 0, 0, 25, 194, 1, 0, 0, 0, 27, 196, 1, 0, 0, 0, 29, 198, 1, 0, 0,
		0, 31, 200, 1, 0, 0, 0, 33, 202, 1, 0, 0, 0, 35, 204, 1, 0, 0, 0, 37, 206,
		1, 0, 0, 0, 39, 209, 1, 0, 0, 0, 41, 211, 1, 0, 0, 0, 43, 214, 1, 0, 0,
		0, 45, 216, 1, 0, 0, 0, 47, 218, 1, 0, 0, 0, 49, 220, 1, 0, 0, 0, 51, 222,
		1, 0, 0, 0, 53, 224, 1, 0, 0, 0, 55, 226, 1, 0, 0, 0, 57, 230, 1, 0, 0,
		0, 59, 232, 1, 0, 0, 0, 61, 234, 1, 0, 0, 0, 63, 238, 1, 0, 0, 0, 65, 243,
		1, 0, 0, 0, 67, 249, 1, 0, 0, 0, 69, 251, 1, 0, 0, 0, 71, 255, 1, 0, 0,
		0, 73, 261, 1, 0, 0, 0, 75, 264, 1, 0, 0, 0, 77, 267, 1, 0, 0, 0, 79, 272,
		1, 0, 0, 0, 81, 276, 1, 0, 0, 0, 83, 279, 1, 0, 0, 0, 85, 284, 1, 0, 0,
		0, 87, 296, 1, 0, 0, 0, 89, 307, 1, 0, 0, 0, 91, 328, 1, 0, 0, 0, 93, 330,
		1, 0, 0, 0, 95, 337, 1, 0, 0, 0, 97, 344, 1, 0, 0, 0, 99, 353, 1, 0, 0,
		0, 101, 357, 1, 0, 0, 0, 103, 364, 1, 0, 0, 0, 105, 368, 1, 0, 0, 0, 107,
		388, 1, 0, 0, 0, 109, 416, 1, 0, 0, 0, 111, 420, 1, 0, 0, 0, 113, 422,
		1, 0, 0, 0, 115, 428, 1, 0, 0, 0, 117, 430, 1, 0, 0, 0, 119, 432, 1, 0,
		0, 0, 121, 434, 1, 0, 0, 0, 123, 436, 1, 0, 0, 0, 125, 438, 1, 0, 0, 0,
		127, 447, 1, 0, 0, 0, 129, 451, 1, 0, 0, 0, 131, 455, 1, 0, 0, 0, 133,
		465, 1, 0, 0, 0, 135, 476, 1, 0, 0, 0, 137, 508, 1, 0, 0, 0, 139, 536,
		1, 0, 0, 0, 141, 540, 1, 0, 0, 0, 143, 543, 1, 0, 0, 0, 145, 546, 1, 0,
		0, 0, 147, 549, 1, 0, 0, 0, 149, 551, 1, 0, 0, 0, 151, 555, 1, 0, 0, 0,
		153, 570, 1, 0, 0, 0, 155, 572, 1, 0, 0, 0, 157, 579, 1, 0, 0, 0, 159,
		591, 1, 0, 0, 0, 161, 595, 1, 0, 0, 0, 163, 164, 5, 60, 0, 0, 164, 2, 1,
		0, 0, 0, 165, 166, 5, 62, 0, 0, 166, 4, 1, 0, 0, 0, 167, 168, 5, 61, 0,
		0, 168, 169, 5, 61, 0, 0, 169, 6, 1, 0, 0, 0, 170, 171, 5, 62, 0, 0, 171,
		172, 5, 61, 0, 0, 172, 8, 1, 0, 0, 0, 173, 174, 5, 60, 0, 0, 174, 175,
		5, 61, 0, 0, 175, 10, 1, 0, 0, 0, 176, 177, 5, 60, 0, 0, 177, 178, 5, 62,
		0, 0, 178, 12, 1, 0, 0, 0, 179, 180, 5, 33, 0, 0, 180, 181, 5, 61, 0, 0,
		181, 14, 1, 0, 0, 0, 182, 183, 5, 124, 0, 0, 183, 16, 1, 0, 0, 0, 184,
		185, 5, 94, 0, 0, 185, 18, 1, 0, 0, 0, 186, 187, 5, 38, 0, 0, 187, 20,
		1, 0, 0, 0, 188, 189, 5, 60, 0, 0, 189, 190, 5, 60, 0, 0, 190, 22, 1, 0,
		0, 0, 191, 192, 5, 62, 0, 0, 192, 193, 5, 62, 0, 0, 193, 24, 1, 0, 0, 0,
		194, 195, 5, 43, 0, 0, 195, 26, 1, 0, 0, 0, 196, 197, 5, 45, 0, 0, 197,
		28, 1, 0, 0, 0, 198, 199, 5, 42, 0, 0, 199, 30, 1, 0, 0, 0, 200, 201, 5,
		64, 0, 0, 201, 32, 1, 0, 0, 0, 202, 203, 5, 47, 0, 0, 203, 34, 1, 0, 0,
		0, 204, 205, 5, 37, 0, 0, 205, 36, 1, 0, 0, 0, 206, 207, 5, 47, 0, 0, 207,
		208, 5, 47, 0, 0, 208, 38, 1, 0, 0, 0, 209, 210, 5, 126, 0, 0, 210, 40,
		1, 0, 0, 0, 211, 212, 5, 42, 0, 0, 212, 213, 5, 42, 0, 0, 213, 42, 1, 0,
		0, 0, 214, 215, 5, 40, 0, 0, 215, 44, 1, 0, 0, 0, 216, 217, 5, 41, 0, 0,
		217, 46, 1, 0, 0, 0, 218, 219, 5, 91, 0, 0, 219, 48, 1, 0, 0, 0, 220, 221,
		5, 93, 0, 0, 221, 50, 1, 0, 0, 0, 222, 223, 5, 123, 0, 0, 223, 52, 1, 0,
		0, 0, 224, 225, 5, 125, 0, 0, 225, 54, 1, 0, 0, 0, 226, 227, 5, 46, 0,
		0, 227, 228, 5, 46, 0, 0, 228, 229, 5, 46, 0, 0, 229, 56, 1, 0, 0, 0, 230,
		231, 5, 44, 0, 0, 231, 58, 1, 0, 0, 0, 232, 233, 5, 46, 0, 0, 233, 60,
		1, 0, 0, 0, 234, 235, 5, 58, 0, 0, 235, 62, 1, 0, 0, 0, 236, 239, 3, 87,
		43, 0, 237, 239, 3, 89, 44, 0, 238, 236, 1, 0, 0, 0, 238, 237, 1, 0, 0,
		0, 239, 64, 1, 0, 0, 0, 240, 244, 3, 67, 33, 0, 241, 244, 3, 99, 49, 0,
		242, 244, 3, 101, 50, 0, 243, 240, 1, 0, 0, 0, 243, 241, 1, 0, 0, 0, 243,
		242, 1, 0, 0, 0, 244, 66, 1, 0, 0, 0, 245, 250, 3, 91, 45, 0, 246, 250,
		3, 93, 46, 0, 247, 250, 3, 95, 47, 0, 248, 250, 3, 97, 48, 0, 249, 245,
		1, 0, 0, 0, 249, 246, 1, 0, 0, 0, 249, 247, 1, 0, 0, 0, 249, 248, 1, 0,
		0, 0, 250, 68, 1, 0, 0, 0, 251, 252, 5, 97, 0, 0, 252, 253, 5, 110, 0,
		0, 253, 254, 5, 100, 0, 0, 254, 70, 1, 0, 0, 0, 255, 256, 5, 70, 0, 0,
		256, 257, 5, 97, 0, 0, 257, 258, 5, 108, 0, 0, 258, 259, 5, 115, 0, 0,
		259, 260, 5, 101, 0, 0, 260, 72, 1, 0, 0, 0, 261, 262, 5, 105, 0, 0, 262,
		263, 5, 110, 0, 0, 263, 74, 1, 0, 0, 0, 264, 265, 5, 105, 0, 0, 265, 266,
		5, 115, 0, 0, 266, 76, 1, 0, 0, 0, 267, 268, 5, 78, 0, 0, 268, 269, 5,
		111, 0, 0, 269, 270, 5, 110, 0, 0, 270, 271, 5, 101, 0, 0, 271, 78, 1,
		0, 0, 0, 272, 273, 5, 110, 0, 0, 273, 274, 5, 111, 0, 0, 274, 275, 5, 116,
		0, 0, 275, 80, 1, 0, 0, 0, 276, 277, 5, 111, 0, 0, 277, 278, 5, 114, 0,
		0, 278, 82, 1, 0, 0, 0, 279, 280, 5, 84, 0, 0, 280, 281, 5, 114, 0, 0,
		281, 282, 5, 117, 0, 0, 282, 283, 5, 101, 0, 0, 283, 84, 1, 0, 0, 0, 284,
		288, 3, 159, 79, 0, 285, 287, 3, 161, 80, 0, 286, 285, 1, 0, 0, 0, 287,
		290, 1, 0, 0, 0, 288, 286, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289, 86, 1,
		0, 0, 0, 290, 288, 1, 0, 0, 0, 291, 297, 7, 0, 0, 0, 292, 293, 7, 1, 0,
		0, 293, 297, 7, 2, 0, 0, 294, 295, 7, 2, 0, 0, 295, 297, 7, 1, 0, 0, 296,
		291, 1, 0, 0, 0, 296, 292, 1, 0, 0, 0, 296, 294, 1, 0, 0, 0, 296, 297,
		1, 0, 0, 0, 297, 300, 1, 0, 0, 0, 298, 301, 3, 107, 53, 0, 299, 301, 3,
		109, 54, 0, 300, 298, 1, 0, 0, 0, 300, 299, 1, 0, 0, 0, 301, 88, 1, 0,
		0, 0, 302, 308, 7, 3, 0, 0, 303, 304, 7, 3, 0, 0, 304, 308, 7, 2, 0, 0,
		305, 306, 7, 2, 0, 0, 306, 308, 7, 3, 0, 0, 307, 302, 1, 0, 0, 0, 307,
		303, 1, 0, 0, 0, 307, 305, 1, 0, 0, 0, 308, 311, 1, 0, 0, 0, 309, 312,
		3, 137, 68, 0, 310, 312, 3, 139, 69, 0, 311, 309, 1, 0, 0, 0, 311, 310,
		1, 0, 0, 0, 312, 90, 1, 0, 0, 0, 313, 320, 3, 117, 58, 0, 314, 316, 5,
		95, 0, 0, 315, 314, 1, 0, 0, 0, 315, 316, 1, 0, 0, 0, 316, 317, 1, 0, 0,
		0, 317, 319, 3, 119, 59, 0, 318, 315, 1, 0, 0, 0, 319, 322, 1, 0, 0, 0,
		320, 318, 1, 0, 0, 0, 320, 321, 1, 0, 0, 0, 321, 329, 1, 0, 0, 0, 322,
		320, 1, 0, 0, 0, 323, 325, 5, 48, 0, 0, 324, 323, 1, 0, 0, 0, 325, 326,
		1, 0, 0, 0, 326, 324, 1, 0, 0, 0, 326, 327, 1, 0, 0, 0, 327, 329, 1, 0,
		0, 0, 328, 313, 1, 0, 0, 0, 328, 324, 1, 0, 0, 0, 329, 92, 1, 0, 0, 0,
		330, 331, 5, 48, 0, 0, 331, 333, 7, 4, 0, 0, 332, 334, 3, 121, 60, 0, 333,
		332, 1, 0, 0, 0, 334, 335, 1, 0, 0, 0, 335, 333, 1, 0, 0, 0, 335, 336,
		1, 0, 0, 0, 336, 94, 1, 0, 0, 0, 337, 338, 5, 48, 0, 0, 338, 340, 7, 5,
		0, 0, 339, 341, 3, 123, 61, 0, 340, 339, 1, 0, 0, 0, 341, 342, 1, 0, 0,
		0, 342, 340, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 96, 1, 0, 0, 0, 344,
		345, 5, 48, 0, 0, 345, 347, 7, 3, 0, 0, 346, 348, 3, 125, 62, 0, 347, 346,
		1, 0, 0, 0, 348, 349, 1, 0, 0, 0, 349, 347, 1, 0, 0, 0, 349, 350, 1, 0,
		0, 0, 350, 98, 1, 0, 0, 0, 351, 354, 3, 127, 63, 0, 352, 354, 3, 129, 64,
		0, 353, 351, 1, 0, 0, 0, 353, 352, 1, 0, 0, 0, 354, 100, 1, 0, 0, 0, 355,
		358, 3, 99, 49, 0, 356, 358, 3, 131, 65, 0, 357, 355, 1, 0, 0, 0, 357,
		356, 1, 0, 0, 0, 358, 359, 1, 0, 0, 0, 359, 360, 7, 6, 0, 0, 360, 102,
		1, 0, 0, 0, 361, 365, 3, 153, 76, 0, 362, 365, 3, 155, 77, 0, 363, 365,
		3, 157, 78, 0, 364, 361, 1, 0, 0, 0, 364, 362, 1, 0, 0, 0, 364, 363, 1,
		0, 0, 0, 365, 366, 1, 0, 0, 0, 366, 367, 6, 51, 0, 0, 367, 104, 1, 0, 0,
		0, 368, 369, 9, 0, 0, 0, 369, 106, 1, 0, 0, 0, 370, 375, 5, 39, 0, 0, 371,
		374, 3, 115, 57, 0, 372, 374, 8, 7, 0, 0, 373, 371, 1, 0, 0, 0, 373, 372,
		1, 0, 0, 0, 374, 377, 1, 0, 0, 0, 375, 373, 1, 0, 0, 0, 375, 376, 1, 0,
		0, 0, 376, 378, 1, 0, 0, 0, 377, 375, 1, 0, 0, 0, 378, 389, 5, 39, 0, 0,
		379, 384, 5, 34, 0, 0, 380, 383, 3, 115, 57, 0, 381, 383, 8, 8, 0, 0, 382,
		380, 1, 0, 0, 0, 382, 381, 1, 0, 0, 0, 383, 386, 1, 0, 0, 0, 384, 382,
		1, 0, 0, 0, 384, 385, 1, 0, 0, 0, 385, 387, 1, 0, 0, 0, 386, 384, 1, 0,
		0, 0, 387, 389, 5, 34, 0, 0, 388, 370, 1, 0, 0, 0, 388, 379, 1, 0, 0, 0,
		389, 108, 1, 0, 0, 0, 390, 391, 5, 39, 0, 0, 391, 392, 5, 39, 0, 0, 392,
		393, 5, 39, 0, 0, 393, 397, 1, 0, 0, 0, 394, 396, 3, 111, 55, 0, 395, 394,
		1, 0, 0, 0, 396, 399, 1, 0, 0, 0, 397, 398, 1, 0, 0, 0, 397, 395, 1, 0,
		0, 0, 398, 400, 1, 0, 0, 0, 399, 397, 1, 0, 0, 0, 400, 401, 5, 39, 0, 0,
		401, 402, 5, 39, 0, 0, 402, 417, 5, 39, 0, 0, 403, 404, 5, 34, 0, 0, 404,
		405, 5, 34, 0, 0, 405, 406, 5, 34, 0, 0, 406, 410, 1, 0, 0, 0, 407, 409,
		3, 111, 55, 0, 408, 407, 1, 0, 0, 0, 409, 412, 1, 0, 0, 0, 410, 411, 1,
		0, 0, 0, 410, 408, 1, 0, 0, 0, 411, 413, 1, 0, 0, 0, 412, 410, 1, 0, 0,
		0, 413, 414, 5, 34, 0, 0, 414, 415, 5, 34, 0, 0, 415, 417, 5, 34, 0, 0,
		416, 390, 1, 0, 0, 0, 416, 403, 1, 0, 0, 0, 417, 110, 1, 0, 0, 0, 418,
		421, 3, 113, 56, 0, 419, 421, 3, 115, 57, 0, 420, 418, 1, 0, 0, 0, 420,
		419, 1, 0, 0, 0, 421, 112, 1, 0, 0, 0, 422, 423, 8, 9, 0, 0, 423, 114,
		1, 0, 0, 0, 424, 425, 5, 92, 0, 0, 425, 429, 9, 0, 0, 0, 426, 427, 5, 92,
		0, 0, 427, 429, 3, 153, 76, 0, 428, 424, 1, 0, 0, 0, 428, 426, 1, 0, 0,
		0, 429, 116, 1, 0, 0, 0, 430, 431, 7, 10, 0, 0, 431, 118, 1, 0, 0, 0, 432,
		433, 7, 11, 0, 0, 433, 120, 1, 0, 0, 0, 434, 435, 7, 12, 0, 0, 435, 122,
		1, 0, 0, 0, 436, 437, 7, 13, 0, 0, 437, 124, 1, 0, 0, 0, 438, 439, 7, 14,
		0, 0, 439, 126, 1, 0, 0, 0, 440, 442, 3, 131, 65, 0, 441, 440, 1, 0, 0,
		0, 441, 442, 1, 0, 0, 0, 442, 443, 1, 0, 0, 0, 443, 448, 3, 133, 66, 0,
		444, 445, 3, 131, 65, 0, 445, 446, 5, 46, 0, 0, 446, 448, 1, 0, 0, 0, 447,
		441, 1, 0, 0, 0, 447, 444, 1, 0, 0, 0, 448, 128, 1, 0, 0, 0, 449, 452,
		3, 131, 65, 0, 450, 452, 3, 127, 63, 0, 451, 449, 1, 0, 0, 0, 451, 450,
		1, 0, 0, 0, 452, 453, 1, 0, 0, 0, 453, 454, 3, 135, 67, 0, 454, 130, 1,
		0, 0, 0, 455, 462, 3, 119, 59, 0, 456, 458, 5, 95, 0, 0, 457, 456, 1, 0,
		0, 0, 457, 458, 1, 0, 0, 0, 458, 459, 1, 0, 0, 0, 459, 461, 3, 119, 59,
		0, 460, 457, 1, 0, 0, 0, 461, 464, 1, 0, 0, 0, 462, 460, 1, 0, 0, 0, 462,
		463, 1, 0, 0, 0, 463, 132, 1, 0, 0, 0, 464, 462, 1, 0, 0, 0, 465, 466,
		5, 46, 0, 0, 466, 473, 3, 119, 59, 0, 467, 469, 5, 95, 0, 0, 468, 467,
		1, 0, 0, 0, 468, 469, 1, 0, 0, 0, 469, 470, 1, 0, 0, 0, 470, 472, 3, 119,
		59, 0, 471, 468, 1, 0, 0, 0, 472, 475, 1, 0, 0, 0, 473, 471, 1, 0, 0, 0,
		473, 474, 1, 0, 0, 0, 474, 134, 1, 0, 0, 0, 475, 473, 1, 0, 0, 0, 476,
		478, 7, 15, 0, 0, 477, 479, 7, 16, 0, 0, 478, 477, 1, 0, 0, 0, 478, 479,
		1, 0, 0, 0, 479, 480, 1, 0, 0, 0, 480, 487, 3, 119, 59, 0, 481, 483, 5,
		95, 0, 0, 482, 481, 1, 0, 0, 0, 482, 483, 1, 0, 0, 0, 483, 484, 1, 0, 0,
		0, 484, 486, 3, 119, 59, 0, 485, 482, 1, 0, 0, 0, 486, 489, 1, 0, 0, 0,
		487, 485, 1, 0, 0, 0, 487, 488, 1, 0, 0, 0, 488, 136, 1, 0, 0, 0, 489,
		487, 1, 0, 0, 0, 490, 495, 5, 39, 0, 0, 491, 494, 3, 143, 71, 0, 492, 494,
		3, 149, 74, 0, 493, 491, 1, 0, 0, 0, 493, 492, 1, 0, 0, 0, 494, 497, 1,
		0, 0, 0, 495, 493, 1, 0, 0, 0, 495, 496, 1, 0, 0, 0, 496, 498, 1, 0, 0,
		0, 497, 495, 1, 0, 0, 0, 498, 509, 5, 39, 0, 0, 499, 504, 5, 34, 0, 0,
		500, 503, 3, 145, 72, 0, 501, 503, 3, 149, 74, 0, 502, 500, 1, 0, 0, 0,
		502, 501, 1, 0, 0, 0, 503, 506, 1, 0, 0, 0, 504, 502, 1, 0, 0, 0, 504,
		505, 1, 0, 0, 0, 505, 507, 1, 0, 0, 0, 506, 504, 1, 0, 0, 0, 507, 509,
		5, 34, 0, 0, 508, 490, 1, 0, 0, 0, 508, 499, 1, 0, 0, 0, 509, 138, 1, 0,
		0, 0, 510, 511, 5, 39, 0, 0, 511, 512, 5, 39, 0, 0, 512, 513, 5, 39, 0,
		0, 513, 517, 1, 0, 0, 0, 514, 516, 3, 141, 70, 0, 515, 514, 1, 0, 0, 0,
		516, 519, 1, 0, 0, 0, 517, 518, 1, 0, 0, 0, 517, 515, 1, 0, 0, 0, 518,
		520, 1, 0, 0, 0, 519, 517, 1, 0, 0, 0, 520, 521, 5, 39, 0, 0, 521, 522,
		5, 39, 0, 0, 522, 537, 5, 39, 0, 0, 523, 524, 5, 34, 0, 0, 524, 525, 5,
		34, 0, 0, 525, 526, 5, 34, 0, 0, 526, 530, 1, 0, 0, 0, 527, 529, 3, 141,
		70, 0, 528, 527, 1, 0, 0, 0, 529, 532, 1, 0, 0, 0, 530, 531, 1, 0, 0, 0,
		530, 528, 1, 0, 0, 0, 531, 533, 1, 0, 0, 0, 532, 530, 1, 0, 0, 0, 533,
		534, 5, 34, 0, 0, 534, 535, 5, 34, 0, 0, 535, 537, 5, 34, 0, 0, 536, 510,
		1, 0, 0, 0, 536, 523, 1, 0, 0, 0, 537, 140, 1, 0, 0, 0, 538, 541, 3, 147,
		73, 0, 539, 541, 3, 149, 74, 0, 540, 538, 1, 0, 0, 0, 540, 539, 1, 0, 0,
		0, 541, 142, 1, 0, 0, 0, 542, 544, 7, 17, 0, 0, 543, 542, 1, 0, 0, 0, 544,
		144, 1, 0, 0, 0, 545, 547, 7, 18, 0, 0, 546, 545, 1, 0, 0, 0, 547, 146,
		1, 0, 0, 0, 548, 550, 7, 19, 0, 0, 549, 548, 1, 0, 0, 0, 550, 148, 1, 0,
		0, 0, 551, 552, 5, 92, 0, 0, 552, 553, 7, 20, 0, 0, 553, 150, 1, 0, 0,
		0, 554, 556, 7, 21, 0, 0, 555, 554, 1, 0, 0, 0, 556, 557, 1, 0, 0, 0, 557,
		555, 1, 0, 0, 0, 557, 558, 1, 0, 0, 0, 558, 152, 1, 0, 0, 0, 559, 571,
		3, 151, 75, 0, 560, 562, 5, 13, 0, 0, 561, 560, 1, 0, 0, 0, 561, 562, 1,
		0, 0, 0, 562, 563, 1, 0, 0, 0, 563, 566, 5, 10, 0, 0, 564, 566, 2, 12,
		13, 0, 565, 561, 1, 0, 0, 0, 565, 564, 1, 0, 0, 0, 566, 568, 1, 0, 0, 0,
		567, 569, 3, 151, 75, 0, 568, 567, 1, 0, 0, 0, 568, 569, 1, 0, 0, 0, 569,
		571, 1, 0, 0, 0, 570, 559, 1, 0, 0, 0, 570, 565, 1, 0, 0, 0, 571, 154,
		1, 0, 0, 0, 572, 576, 5, 35, 0, 0, 573, 575, 8, 22, 0, 0, 574, 573, 1,
		0, 0, 0, 575, 578, 1, 0, 0, 0, 576, 574, 1, 0, 0, 0, 576, 577, 1, 0, 0,
		0, 577, 156, 1, 0, 0, 0, 578, 576, 1, 0, 0, 0, 579, 581, 5, 92, 0, 0, 580,
		582, 3, 151, 75, 0, 581, 580, 1, 0, 0, 0, 581, 582, 1, 0, 0, 0, 582, 588,
		1, 0, 0, 0, 583, 585, 5, 13, 0, 0, 584, 583, 1, 0, 0, 0, 584, 585, 1, 0,
		0, 0, 585, 586, 1, 0, 0, 0, 586, 589, 5, 10, 0, 0, 587, 589, 2, 12, 13,
		0, 588, 584, 1, 0, 0, 0, 588, 587, 1, 0, 0, 0, 589, 158, 1, 0, 0, 0, 590,
		592, 7, 23, 0, 0, 591, 590, 1, 0, 0, 0, 592, 160, 1, 0, 0, 0, 593, 596,
		3, 159, 79, 0, 594, 596, 7, 24, 0, 0, 595, 593, 1, 0, 0, 0, 595, 594, 1,
		0, 0, 0, 596, 162, 1, 0, 0, 0, 62, 0, 238, 243, 249, 288, 296, 300, 307,
		311, 315, 320, 326, 328, 335, 342, 349, 353, 357, 364, 373, 375, 382, 384,
		388, 397, 410, 416, 420, 428, 441, 447, 451, 457, 462, 468, 473, 478, 482,
		487, 493, 495, 502, 504, 508, 517, 530, 536, 540, 543, 546, 549, 557, 561,
		565, 568, 570, 576, 581, 584, 588, 591, 595, 1, 0, 1, 0,
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

// ExprLexerInit initializes any static state used to implement ExprLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewExprLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ExprLexerInit() {
	staticData := &exprlexerLexerStaticData
	staticData.once.Do(exprlexerLexerInit)
}

// NewExprLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewExprLexer(input antlr.CharStream) *ExprLexer {
	ExprLexerInit()
	l := new(ExprLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &exprlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Expr.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ExprLexer tokens.
const (
	ExprLexerT__0            = 1
	ExprLexerT__1            = 2
	ExprLexerT__2            = 3
	ExprLexerT__3            = 4
	ExprLexerT__4            = 5
	ExprLexerT__5            = 6
	ExprLexerT__6            = 7
	ExprLexerT__7            = 8
	ExprLexerT__8            = 9
	ExprLexerT__9            = 10
	ExprLexerT__10           = 11
	ExprLexerT__11           = 12
	ExprLexerT__12           = 13
	ExprLexerT__13           = 14
	ExprLexerT__14           = 15
	ExprLexerT__15           = 16
	ExprLexerT__16           = 17
	ExprLexerT__17           = 18
	ExprLexerT__18           = 19
	ExprLexerT__19           = 20
	ExprLexerT__20           = 21
	ExprLexerT__21           = 22
	ExprLexerT__22           = 23
	ExprLexerT__23           = 24
	ExprLexerT__24           = 25
	ExprLexerT__25           = 26
	ExprLexerT__26           = 27
	ExprLexerT__27           = 28
	ExprLexerT__28           = 29
	ExprLexerT__29           = 30
	ExprLexerT__30           = 31
	ExprLexerSTRING          = 32
	ExprLexerNUMBER          = 33
	ExprLexerINTEGER         = 34
	ExprLexerAND             = 35
	ExprLexerFALSE           = 36
	ExprLexerIN              = 37
	ExprLexerIS              = 38
	ExprLexerNONE            = 39
	ExprLexerNOT             = 40
	ExprLexerOR              = 41
	ExprLexerTRUE            = 42
	ExprLexerNAME            = 43
	ExprLexerSTRING_LITERAL  = 44
	ExprLexerBYTES_LITERAL   = 45
	ExprLexerDECIMAL_INTEGER = 46
	ExprLexerOCT_INTEGER     = 47
	ExprLexerHEX_INTEGER     = 48
	ExprLexerBIN_INTEGER     = 49
	ExprLexerFLOAT_NUMBER    = 50
	ExprLexerIMAG_NUMBER     = 51
	ExprLexerSKIP_           = 52
	ExprLexerUNKNOWN_CHAR    = 53
)
