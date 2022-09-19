package c

import "strings"

// C parser limits.
const CPARSE_MAX_BUF = 32768     // Max. token buffer size.
const CPARSE_MAX_DECLSTACK = 100 // Max. declaration stack depth.
const CPARSE_MAX_DECLDEPTH = 20  // Max. recursive declaration depth.
const CPARSE_MAX_PACKSTACK = 7   // Max. pack pragma stack depth.

// Flags for C parser mode.
const CPARSE_MODE_MULTI = 1       // Process multiple declarations.
const CPARSE_MODE_ABSTRACT = 2    // Accept abstract declarators.
const CPARSE_MODE_DIRECT = 4      // Accept direct declarators.
const CPARSE_MODE_FIELD = 8       // Accept field width in bits, too.
const CPARSE_MODE_NOIMPLICIT = 16 // Reject implicit declarations.
const CPARSE_MODE_SKIP = 32       // Skip definitions, ignore errors.

type CPChar = rune // C parser character. Unsigned ext. from char.
type CPToken = int // C parser token.

// C parser internal value representation.
type CPValue struct {
	//union {
	//int32_t i32;    // Value for CTID_INT32.
	//uint32_t u32;    // Value for CTID_UINT32.
	//};
	i32 int32
	id  CTypeID // C Type ID of the value.
}

// C parser state.
type CPState struct {
	c   CPChar          // Current character.
	tok CPToken         // Current token.
	val CPValue         // Token value.
	str string          // Interned string of identifier/keyword.
	ct  *CType          // C type table entry.
	p   int             // Current position in input buffer.
	i   []rune          // Input buffer
	sb  strings.Builder // String buffer for tokens.
	//lua_State *L;  // Lua state.
	//CTState *cts;  // C type state.
	//TValue *param; // C type parameters.
	srcname    string                      // Current source name.
	linenumber int                         // Input line counter.
	depth      int                         // Recursive declaration depth.
	tmask      uint32                      // Type mask for next identifier.
	mode       uint32                      // C parser mode.
	packstack  [CPARSE_MAX_PACKSTACK]uint8 // Stack for pack pragmas.
	curpack    uint8                       // Current position in pack pragma stack.
}

const CPNS_DEFAULT = (1 << CT_KW) | (1 << CT_TYPEDEF) | (1 << CT_FUNC) | (1 << CT_EXTERN) | (1 << CT_CONSTVAL)

const CPNS_STRUCT = (1 << CT_KW) | (1 << CT_STRUCT) | (1 << CT_ENUM)

func cp_init(cp *CPState) {
	cp.linenumber = 1
	cp.depth = 0
	cp.curpack = 0
	cp.packstack[0] = 255
	//lj_buf_init(cp->L, &cp->sb);
	//lj_assertCP(cp->p != NULL, "uninitialized cp->p");
	cp_get(cp) // Read-ahead first char.
	cp.tok = 0
	cp.tmask = CPNS_DEFAULT
	cp_next(cp) // Read-ahead first token.
}

// Get next character.
func cp_get(cp *CPState) CPChar {
	cp.c = cp.i[cp.p)
	cp.p++
	if cp.c != '\\' {
		return cp.c
	}
	return cp_get_bs(cp)
}

// Transparently skip backslash-escaped line breaks.
func cp_get_bs(cp *CPState) CPChar {
	var c2 CPChar
	c := cp_rawpeek(cp)
	if !cp_iseol(c) {
		return cp.c
	}
	cp.p++
	c2 = cp_rawpeek(cp)
	if cp_iseol(c2) && c2 != c {
		cp.p++
	}
	cp.linenumber++
	return cp_get(cp)
}

// End-of-line?
func cp_iseol(c CPChar) bool {
	return c == '\n' || c == '\r'
}

// Peek next raw character.
func cp_rawpeek(cp *CPState) CPChar {
	return cp.i[cp.p]
}

func cp_next(cp *CPState) CPToken {
	cp.tok = cp_next_(cp)
	return cp.tok
}

// // C lexer keywords.
// #define CTOKDEF(_) \
// _(IDENT, "<identifier>")
// _(STRING, "<string>")
// _(INTEGER, "<integer>")
// _(EOF, "<eof>")
// _(OROR, "||")
// _(ANDAND, "&&")
// _(EQ, "==")
// _(NE, "!=")
// _(LE, "<=")
// _(GE, ">=")
// _(SHL, "<<")
// _(SHR, ">>")
// _(DEREF, "->")

// // Simple declaration specifiers.
// #define CDSDEF(_) \
// _(VOID)
// _(BOOL)
// _(CHAR)
// _(INT)
// _(FP) \
// _(LONG)
// _(LONGLONG)
// _(SHORT)
// _(COMPLEX)
// _(SIGNED)
// _(UNSIGNED) \
// _(CONST)
// _(VOLATILE)
// _(RESTRICT)
// _(INLINE) \
// _(TYPEDEF)
// _(EXTERN)
// _(STATIC)
// _(AUTO)
// _(REGISTER)

// // C keywords.
// #define CKWDEF(_) \
// CDSDEF(_)
// _(EXTENSION)
// _(ASM)
// _(ATTRIBUTE) \
// _(DECLSPEC)
// _(CCDECL)
// _(PTRSZ) \
// _(STRUCT)
// _(UNION)
// _(ENUM) \
// _(SIZEOF)
// _(ALIGNOF)

// C token numbers.
const (
	CTOK_OFS = iota + 255

	CTOK_IDENT
	CTOK_STRING
	CTOK_INTEGER
	CTOK_EOF
	CTOK_OROR
	CTOK_ANDAND
	CTOK_EQ
	CTOK_NE
	CTOK_LE
	CTOK_GE
	CTOK_SHL
	CTOK_SHR
	CTOK_DEREF

	CTOK_VOID
	CTOK_BOOL
	CTOK_CHAR
	CTOK_INT
	CTOK_FP
	CTOK_LONG
	CTOK_LONGLONG
	CTOK_SHORT
	CTOK_COMPLEX
	CTOK_SIGNED
	CTOK_UNSIGNED
	CTOK_CONST
	CTOK_VOLATILE
	CTOK_RESTRICT
	CTOK_INLINE
	CTOK_TYPEDEF
	CTOK_EXTERN
	CTOK_STATIC
	CTOK_AUTO
	CTOK_REGISTER

	CTOK_EXTENSION
	CTOK_ASM
	CTOK_ATTRIBUTE
	CTOK_DECLSPEC
	CTOK_CCDECL
	CTOK_PTRSZ
	CTOK_STRUCT
	CTOK_UNION
	CTOK_ENUM
	CTOK_SIZEOF
	CTOK_ALIGNOF

	CTOK_FIRSTDECL    = CTOK_VOID
	CTOK_FIRSTSCL     = CTOK_TYPEDEF
	CTOK_LASTDECLFLAG = CTOK_REGISTER
	CTOK_LASTDECL     = CTOK_ENUM
)

// Declaration specifier flags.
const (
	CDF_VOID = 1 << iota
	CDF_BOOL
	CDF_CHAR
	CDF_INT
	CDF_FP
	CDF_LONG
	CDF_LONGLONG
	CDF_SHORT
	CDF_COMPLEX
	CDF_SIGNED
	CDF_UNSIGNED
	CDF_CONST
	CDF_VOLATILE
	CDF_RESTRICT
	CDF_INLINE
	CDF_TYPEDEF
	CDF_EXTERN
	CDF_STATIC
	CDF_AUTO
	CDF_REGISTER
)

const CDF_SCL = CDF_TYPEDEF | CDF_EXTERN | CDF_STATIC | CDF_AUTO | CDF_REGISTER

func cp_next_(cp *CPState) CPToken {
	cp.sb.Reset()
	for {
		if lj_char_isident(cp.c) {
			if lj_char_isdigit(cp.c) {
				return cp_number(cp)
			} else {
				return cp_ident(cp)
			}
		}
		switch cp.c {
		case '\n',
			'\r':
			cp_newline(cp)
			fallthrough
		case ' ',
			'\t',
			'\v',
			'\f':
			cp_get(cp)
			break
		case '"',
			'\'':
			return cp_string(cp)
		case '/':
			if cp_get(cp) == '*' {
				cp_comment_c(cp)
			} else if cp.c == '/' {
				cp_comment_cpp(cp)
			} else {
				return '/'
			}
			break
		case '|':
			if cp_get(cp) != '|' {
				return '|'
			}
			cp_get(cp)
			return CTOK_OROR
		case '&':
			if cp_get(cp) != '&' {
				return '&'
			}
			cp_get(cp)
			return CTOK_ANDAND
		case '=':
			if cp_get(cp) != '=' {
				return '='
			}
			cp_get(cp)
			return CTOK_EQ
		case '!':
			if cp_get(cp) != '=' {
				return '!'
			}
			cp_get(cp)
			return CTOK_NE
		case '<':
			if cp_get(cp) == '=' {
				cp_get(cp)
				return CTOK_LE
			} else if cp.c == '<' {
				cp_get(cp)
				return CTOK_SHL
			}
			return '<'
		case '>':
			if cp_get(cp) == '=' {
				cp_get(cp)
				return CTOK_GE
			} else if cp.c == '>' {
				cp_get(cp)
				return CTOK_SHR
			}
			return '>'
		case '-':
			if cp_get(cp) != '>' {
				return '-'
			}
			cp_get(cp)
			return CTOK_DEREF
		case '$':
			return cp_param(cp)
		case 0:
			return CTOK_EOF
		default:
			{
				c := cp.c
				cp_get(cp)
				return c
			}
		}
	}
}
