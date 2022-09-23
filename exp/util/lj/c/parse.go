package c

import (
	"math/bits"
	"reflect"
	"strings"
	"unicode"

	bt "github.com/wrmsr/bane/pkg/util/types"
)

// C parser limits.
const CPARSE_MAX_BUF = 32768     // Max. token buffer size.
const CPARSE_MAX_DECLSTACK = 100 // Max. declaration stack depth.
const CPARSE_MAX_DECLDEPTH = 20  // Max. recursive declaration depth.
const CPARSE_MAX_PACKSTACK = 7   // Max. pack pragma stack depth.

// Flags for C parser mode.
const (
	CPARSE_MODE_MULTI      CPscl = 1  // Process multiple declarations.
	CPARSE_MODE_ABSTRACT   CPscl = 2  // Accept abstract declarators.
	CPARSE_MODE_DIRECT     CPscl = 4  // Accept direct declarators.
	CPARSE_MODE_FIELD      CPscl = 8  // Accept field width in bits, too.
	CPARSE_MODE_NOIMPLICIT CPscl = 16 // Reject implicit declarations.
	CPARSE_MODE_SKIP       CPscl = 32 // Skip definitions, ignore errors.
)

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
	cts *CTState // C type state.
	//TValue *param; // C type parameters.
	srcname    string                      // Current source name.
	linenumber int                         // Input line counter.
	depth      int                         // Recursive declaration depth.
	tmask      uint32                      // Type mask for next identifier.
	mode       CPscl                       // C parser mode.
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
	//lj_buf_init(cp.L, &cp.sb);
	//lj_assertCP(cp.p != NULL, "uninitialized cp.p");
	cp_get(cp) // Read-ahead first char.
	cp.tok = 0
	cp.tmask = CPNS_DEFAULT
	cp_next(cp) // Read-ahead first token.
}

// Get next character.
func cp_get(cp *CPState) CPChar {
	cp.c = cp.i[cp.p]
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

// Check and consume optional token.
func cp_opt(cp *CPState, tok CPToken) bool {
	if cp.tok == tok {
		cp_next(cp)
		return true
	}
	return false
}

// Check and consume token.
func cp_check(cp *CPState, tok CPToken) {
	if cp.tok != tok {
		//cp_err_token(cp, tok)
		panic(cp)
	}
	cp_next(cp)
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
// _(IDENT,   "<identifier>")
// _(STRING,  "<string>")
// _(INTEGER, "<integer>")
// _(EOF,     "<eof>")
// _(OROR,    "||")
// _(ANDAND,  "&&")
// _(EQ,      "==")
// _(NE,      "!=")
// _(LE,      "<=")
// _(GE,      ">=")
// _(SHL,     "<<")
// _(SHR,     ">>")
// _(DEREF,   ".")

// // Simple declaration specifiers.
// #define CDSDEF(_) \
// _(VOID)
// _(BOOL)
// _(CHAR)
// _(INT)
// _(FP)
// _(LONG)
// _(LONGLONG)
// _(SHORT)
// _(COMPLEX)
// _(SIGNED)
// _(UNSIGNED)
// _(CONST)
// _(VOLATILE)
// _(RESTRICT)
// _(INLINE)
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
// _(ATTRIBUTE)
// _(DECLSPEC)
// _(CCDECL)
// _(PTRSZ)
// _(STRUCT)
// _(UNION)
// _(ENUM)
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
	CDF_VOID = CPscl(1) << iota
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

// Skip line break. Handles "\n", "\r", "\r\n" or "\n\r".
func cp_newline(cp *CPState) {
	c := cp_rawpeek(cp)
	if cp_iseol(c) && c != cp.c {
		cp.p++
	}
	cp.linenumber++
}

/*
// Common types.
#define CTTYDEF(_) \
_(NONE,           0,          CT_ATTRIB, CTATTRIB(CTA_BAD))
_(VOID,           -1,         CT_VOID,   CTALIGN(0))
_(CVOID,          -1,         CT_VOID,   CTF_CONST|CTALIGN(0))
_(BOOL,           1,          CT_NUM,    CTF_BOOL|CTF_UNSIGNED|CTALIGN(0))
_(CCHAR,          1,          CT_NUM,    CTF_CONST|CTF_UCHAR|CTALIGN(0))
_(INT8,           1,          CT_NUM,    CTALIGN(0))
_(UINT8,          1,          CT_NUM,    CTF_UNSIGNED|CTALIGN(0))
_(INT16,          2,          CT_NUM,    CTALIGN(1))
_(UINT16,         2,          CT_NUM,    CTF_UNSIGNED|CTALIGN(1))
_(INT32,          4,          CT_NUM,    CTALIGN(2))
_(UINT32,         4,          CT_NUM,    CTF_UNSIGNED|CTALIGN(2))
_(INT64,          8,          CT_NUM,    CTF_LONG|CTALIGN(3))
_(UINT64,         8,          CT_NUM,    CTF_UNSIGNED|CTF_LONG|CTALIGN(3))
_(FLOAT,          4,          CT_NUM,    CTF_FP|CTALIGN(2))
_(DOUBLE,         8,          CT_NUM,    CTF_FP|CTALIGN(3))
_(COMPLEX_FLOAT,  8,          CT_ARRAY,  CTF_COMPLEX|CTALIGN(2)|CTID_FLOAT)
_(COMPLEX_DOUBLE, 16,         CT_ARRAY,  CTF_COMPLEX|CTALIGN(3)|CTID_DOUBLE)
_(P_VOID,         CTSIZE_PTR, CT_PTR,    CTALIGN_PTR|CTID_VOID)
_(P_CVOID,        CTSIZE_PTR, CT_PTR,    CTALIGN_PTR|CTID_CVOID)
_(P_CCHAR,        CTSIZE_PTR, CT_PTR,    CTALIGN_PTR|CTID_CCHAR)
_(P_UINT8,        CTSIZE_PTR, CT_PTR,    CTALIGN_PTR|CTID_UINT8)
_(A_CCHAR,        -1,         CT_ARRAY,  CTF_CONST|CTALIGN(0)|CTID_CCHAR)
_(CTYPEID,        4,          CT_ENUM,   CTALIGN(2)|CTID_INT32)
// End of type list.
*/

// Public predefined type IDs.
const (
	CTID_NONE = iota
	CTID_VOID
	CTID_CVOID
	CTID_BOOL
	CTID_CCHAR
	CTID_INT8
	CTID_UINT8
	CTID_INT16
	CTID_UINT16
	CTID_INT32
	CTID_UINT32
	CTID_INT64
	CTID_UINT64
	CTID_FLOAT
	CTID_DOUBLE
	CTID_COMPLEX_FLOAT
	CTID_COMPLEX_DOUBLE
	CTID_P_VOID
	CTID_P_CVOID
	CTID_P_CCHAR
	CTID_P_UINT8
	CTID_A_CCHAR
	CTID_CTYPEID
	CTID_MAX = 65536
)

// Save character in buffer.
func cp_save(cp *CPState, c CPChar) {
	cp.sb.WriteRune(c)
}

func isxdigit(r rune) bool {
	return unicode.IsDigit(r) ||
		(r >= 'a' && r <= 'f') ||
		(r >= 'A' && r <= 'F')
}

func isident(r rune) bool {
	return unicode.IsDigit(r) ||
		(r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'z') ||
		r == '_'
}

// Parse string or character constant.
func cp_string(cp *CPState) CPToken {
	delim := cp.c
	cp_get(cp)
	for cp.c != delim {
		c := cp.c
		if c == 0 {
			//cp_errmsg(cp, CTOK_EOF, LJ_ERR_XSTR)
			panic(cp)
		}
		if c == '\\' {
			c = cp_get(cp)
			switch c {
			case 0:
				//cp_errmsg(cp, CTOK_EOF, LJ_ERR_XSTR)
				panic(cp)
			case 'a':
				c = '\a'
			case 'b':
				c = '\b'
			case 'f':
				c = '\f'
			case 'n':
				c = '\n'
			case 'r':
				c = '\r'
			case 't':
				c = '\t'
			case 'v':
				c = '\v'
			case 'e':
				c = 27
			case 'x':
				c = 0
				for isxdigit(cp_get(cp)) {
					c = (c << 4) + bt.Choose(unicode.IsDigit(cp.c), cp.c-'0', (cp.c&15)+9)
				}
				cp_save(cp, c&0xff)
				continue
			default:
				if unicode.IsDigit(c) {
					c -= '0'
					if unicode.IsDigit(cp_get(cp)) {
						c = c*8 + (cp.c - '0')
						if unicode.IsDigit(cp_get(cp)) {
							c = c*8 + (cp.c - '0')
							cp_get(cp)
						}
					}
					cp_save(cp, c&0xff)
					continue
				}
				break
			}
		}
		cp_save(cp, c)
		cp_get(cp)
	}
	cp_get(cp)
	if delim == '"' {
		cp.str = cp.sb.String()
		return CTOK_STRING
	} else {
		if cp.sb.Len() != 1 {
			//cp_err_token(cp, '\'')
			panic(cp)
		}
		// FIXME:
		// cp.val.i32 = (int32_t)(char) * cp.sb.b
		// cp.val.id = CTID_INT32
		// return CTOK_INTEGER
		panic("nyi")
	}
}

func cp_number(cp *CPState) CPToken {
	// TODO: strscan :|
	// var  fmt StrScanFmt;
	// var  o TValue;
	// for {
	// 	cp_save(cp, cp.c);
	// 	if !isident(cp_get(cp)) {
	// 		break
	// 	}
	// }
	// cp_save(cp, 0);
	// fmt = lj_strscan_scan((const uint8_t *) (cp.sb.b), sbuflen(&cp.sb) - 1, &o, STRSCAN_OPT_C);
	// if fmt == STRSCAN_INT {
	// 	cp.val.id = CTID_INT32;
	// } else if fmt == STRSCAN_U32 {
	// 	cp.val.id = CTID_UINT32;
	// } else if (cp.mode & CPARSE_MODE_SKIP) == 0 {
	//     //cp_errmsg(cp, CTOK_INTEGER, LJ_ERR_XNUMBER);
	// 	panic(cp)
	// }
	// cp.val.i32 = (uint32_t) o.i;
	// return CTOK_INTEGER;
	panic("nyi")
}

// Skip C comment.
func cp_comment_c(cp *CPState) {
	for {
		if cp_get(cp) == '*' {
			for {
				if cp_get(cp) == '/' {
					cp_get(cp)
					return
				}
				if cp.c == '*' {
					break
				}
			}
		}
		if cp_iseol(cp.c) {
			cp_newline(cp)
		}
		if cp.c == 0 {
			break
		}
	}
}

// Skip C++ comment.
func cp_comment_cpp(cp *CPState) {
	for !cp_iseol(cp_get(cp)) && cp.c != 0 {
	}
}

// Get a C type by name, matching the type mask.
func ctype_getname(cts *CTState, ctp **CType, name string, tmask uint32) CTypeID {
	// CTypeID id = cts.hash[ct_hashname(name)];
	// while (id) {
	// 	CType *ct = ctype_get(cts, id);
	// 	if (gcref(ct.name) == obj2gco(name) && ((tmask >> ctype_type(ct.info)) & 1)) {
	// 		*ctp = ct;
	// 		return id;
	// 	}
	// 	id = ct.next;
	// }
	// *ctp = &cts.tab[0];  // Simplify caller logic. ctype_get() would assert.
	return 0
}

/*
// Predefined typedefs.
#define CTTDDEF(_) \
// Vararg handling.
_("va_list",           P_VOID)
_("__builtin_va_list", P_VOID)
_("__gnuc_va_list",    P_VOID)
// From stddef.h.
_("ptrdiff_t",         INT_PSZ)
_("size_t",            UINT_PSZ)
_("wchar_t",           WCHAR)
// Subset of stdint.h.
_("int8_t",            INT8)
_("int16_t",           INT16)
_("int32_t",           INT32)
_("int64_t",           INT64)
_("uint8_t",           UINT8)
_("uint16_t",          UINT16)
_("uint32_t",          UINT32)
_("uint64_t",          UINT64)
_("intptr_t",          INT_PSZ)
_("uintptr_t",         UINT_PSZ)
// From POSIX.
_("ssize_t",           INT_PSZ)
// End of typedef list.

// Keywords (only the ones we actually care for).
#define CTKWDEF(_) \
// Type specifiers.
_("void",           -1,                   CTOK_VOID)
_("_Bool",          0,                    CTOK_BOOL)
_("bool",           1,                    CTOK_BOOL)
_("char",           1,                    CTOK_CHAR)
_("int",            4,                    CTOK_INT)
_("__int8",         1,                    CTOK_INT)
_("__int16",        2,                    CTOK_INT)
_("__int32",        4,                    CTOK_INT)
_("__int64",        8,                    CTOK_INT)
_("float",          4,                    CTOK_FP)
_("double",         8,                    CTOK_FP)
_("long",           0,                    CTOK_LONG)
_("short",          0,                    CTOK_SHORT)
_("_Complex",       0,                    CTOK_COMPLEX)
_("complex",        0,                    CTOK_COMPLEX)
_("__complex",      0,                    CTOK_COMPLEX)
_("__complex__",    0,                    CTOK_COMPLEX)
_("signed",         0,                    CTOK_SIGNED)
_("__signed",       0,                    CTOK_SIGNED)
_("__signed__",     0,                    CTOK_SIGNED)
_("unsigned",       0,                    CTOK_UNSIGNED)
// Type qualifiers.
_("const",          0,                    CTOK_CONST)
_("__const",        0,                    CTOK_CONST)
_("__const__",      0,                    CTOK_CONST)
_("volatile",       0,                    CTOK_VOLATILE)
_("__volatile",     0,                    CTOK_VOLATILE)
_("__volatile__",   0,                    CTOK_VOLATILE)
_("restrict",       0,                    CTOK_RESTRICT)
_("__restrict",     0,                    CTOK_RESTRICT)
_("__restrict__",   0,                    CTOK_RESTRICT)
_("inline",         0,                    CTOK_INLINE)
_("__inline",       0,                    CTOK_INLINE)
_("__inline__",     0,                    CTOK_INLINE)
// Storage class specifiers.
_("typedef",        0,                    CTOK_TYPEDEF)
_("extern",         0,                    CTOK_EXTERN)
_("static",         0,                    CTOK_STATIC)
_("auto",           0,                    CTOK_AUTO)
_("register",       0,                    CTOK_REGISTER)
// GCC Attributes.
_("__extension__",  0,                    CTOK_EXTENSION)
_("__attribute",    0,                    CTOK_ATTRIBUTE)
_("__attribute__",  0,                    CTOK_ATTRIBUTE)
_("asm",            0,                    CTOK_ASM)
_("__asm",          0,                    CTOK_ASM)
_("__asm__",        0,                    CTOK_ASM)
// MSVC Attributes.
_("__declspec",     0,                    CTOK_DECLSPEC)
_("__cdecl",        CTCC_CDECL,           CTOK_CCDECL)
_("__thiscall",     CTCC_THISCALL,        CTOK_CCDECL)
_("__fastcall",     CTCC_FASTCALL,        CTOK_CCDECL)
_("__stdcall",      CTCC_STDCALL,         CTOK_CCDECL)
_("__ptr32",        4,                    CTOK_PTRSZ)
_("__ptr64",        8,                    CTOK_PTRSZ)
// Other type specifiers.
_("struct",         0,                    CTOK_STRUCT)
_("union",          0,                    CTOK_UNION)
_("enum",           0,                    CTOK_ENUM)
// Operators.
_("sizeof",         0,                    CTOK_SIZEOF)
_("__alignof",      0,                    CTOK_ALIGNOF)
_("__alignof__",    0,                    CTOK_ALIGNOF)
// End of keyword list.
*/

// Parse identifier or keyword.
func cp_ident(cp *CPState) CPToken {
	for {
		cp_save(cp, cp.c)
		if !isident(cp_get(cp)) {
			break
		}
	}
	cp.str = cp.sb.String()
	//cp.val.id = ctype_getname(cp.cts, cp.ct, cp.str, cp.tmask)
	cp.val.id = ctype_getname(nil, nil, cp.str, cp.tmask)
	//if ctype_type(cp.ct.info) == CT_KW {
	//	return ctype_cid(cp.ct.info)
	//}
	return CTOK_IDENT
}

// Parse parameter.
func cp_param(cp *CPState) CPToken {
	//c := cp_get(cp)
	//o := cp.param
	//if lj_char_isident(c) || c == '$' {  // Reserve $xyz for future extensions.
	//    // cp_errmsg(cp, c, LJ_ERR_XSYNTAX);
	//	panic(cp)
	//}
	//if !o || o >= cp.L.top {
	//    // cp_err(cp, LJ_ERR_FFI_NUMPARAM);
	//	panic(cp)
	//}
	//cp.param = o + 1;
	//if tvisstr(o) {
	//    cp.str = strV(o);
	//    cp.val.id = 0;
	//    cp.ct = &cp.cts.tab[0];
	//    return CTOK_IDENT;
	//} else if tvisnumber(o) {
	//    cp.val.i32 = numberVint(o);
	//    cp.val.id = CTID_INT32;
	//    return CTOK_INTEGER;
	//} else {
	//    var cd *GCcdata
	//    if !tviscdata(o) {
	//        lj_err_argtype(cp.L, (int) (o - cp.L.base) + 1, "type parameter");
	//	}
	//    cd = cdataV(o);
	//    if cd.ctypeid == CTID_CTYPEID {
	//        cp.val.id = *(CTypeID *) cdataptr(cd);
	//    } else {
	//        cp.val.id = cd.ctypeid;
	//	}
	//    return '$';
	//}
	panic("nyi")
}

func cp_next_(cp *CPState) CPToken {
	cp.sb.Reset()
	for {
		if isident(cp.c) {
			if unicode.IsDigit(cp.c) {
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
				return CPToken(c)
			}
		}
	}
}

type CPDeclIdx = CTypeID // Index into declaration stack.
type CPscl uint32        // Storage class flags.

// Type declaration context.
type CPDecl struct {
	top       CPDeclIdx                   // Top of declaration stack.
	pos       CPDeclIdx                   // Insertion position in declaration chain.
	specpos   CPDeclIdx                   // Saved position for declaration specifier.
	mode      CPscl                       // Declarator mode.
	cp        *CPState                    // C parser state.
	name      string                      // Name of declared identifier (if direct).
	redir     string                      // Redirected symbol name.
	nameid    CTypeID                     // Existing typedef for declared identifier.
	attr      CTInfo                      // Attributes.
	fattr     CTInfo                      // Function attributes.
	specattr  CTInfo                      // Saved attributes.
	specfattr CTInfo                      // Saved function attributes.
	bits      CTSize                      // Field size in bits (if any).
	stack     [CPARSE_MAX_DECLSTACK]CType // Type declaration stack.
}

// Get raw type of the child of a C type.
func ctype_rawchild(cts *CTState, ct *CType) *CType {
	for {
		ct = ctype_child(cts, ct)
		if !(ctype_isattrib(ct.info)) {
			break
		}
	}
	return ct
}

// Parse a single C type declaration.
func cp_decl_single(cp *CPState) {
	var decl CPDecl
	cp_decl_spec(cp, &decl, 0)
	cp_declarator(cp, &decl)
	cp.val.id = cp_decl_intern(cp, &decl)
	if cp.tok != CTOK_EOF {
		//cp_err_token(cp, CTOK_EOF)
		panic(cp)
	}
}

// Handle pragmas.
func cp_pragma(cp *CPState, pragmaline int) {
	cp_next(cp)
	if cp.tok == CTOK_IDENT && cp.str == "pack" {
		cp_next(cp)
		cp_check(cp, '(')
		if cp.tok == CTOK_IDENT {
			if cp.str == "push" {
				if cp.curpack < CPARSE_MAX_PACKSTACK {
					cp.packstack[cp.curpack+1] = cp.packstack[cp.curpack]
					cp.curpack++
				}
			} else if cp.str == "pop" {
				if cp.curpack > 0 {
					cp.curpack--
				}
			} else {
				//cp_errmsg(cp, cp.tok, LJ_ERR_XSYMBOL);
				panic(cp)
			}
			cp_next(cp)
			if !cp_opt(cp, ',') {
				goto end_pack
			}
		}
		if cp.tok == CTOK_INTEGER {
			cp.packstack[cp.curpack] = bt.Choose(cp.val.i32 != 0, uint8(lj_fls(uint32(cp.val.i32))), 0)
			cp_next(cp)
		} else {
			cp.packstack[cp.curpack] = 255
		}
	end_pack:
		cp_check(cp, ')')
	} else { // Ignore all other pragmas.
		for {
			cp_next(cp)
			if !(cp.tok != CTOK_EOF && cp.linenumber == pragmaline) {
				break
			}
		}
	}
}

// Parse multiple C declarations of types or extern identifiers.
func cp_decl_multi(cp *CPState) {
	var first = true
	for cp.tok != CTOK_EOF {
		var decl CPDecl
		var scl CPscl
		if cp_opt(cp, ';') { // Skip empty statements.
			first = false
			continue
		}
		if cp.tok == '#' { // Workaround, since we have no preprocessor, yet.
			var hashline int = cp.linenumber
			var tok CPToken = cp_next(cp)
			if tok == CTOK_INTEGER {
				cp_line(cp, hashline)
				continue
			} else if tok == CTOK_IDENT && cp.str == "line" {
				if cp_next(cp) != CTOK_INTEGER {
					//cp_err_token(cp, tok)
					panic(cp)
				}
				cp_line(cp, hashline)
				continue
			} else if tok == CTOK_IDENT && cp.str == "pragma" {
				cp_pragma(cp, hashline)
				continue
			} else {
				//cp_errmsg(cp, cp.tok, LJ_ERR_XSYMBOL);
				panic(cp)
			}
		}
		scl = cp_decl_spec(cp, &decl, CDF_TYPEDEF|CDF_EXTERN|CDF_STATIC)
		if (cp.tok == ';' || cp.tok == CTOK_EOF) &&
			ctype_istypedef(decl.stack[0].info) {
			var info CTInfo = ctype_rawchild(cp.cts, &decl.stack[0]).info
			if ctype_isstruct(info) || ctype_isenum(info) {
				//goto decl_end;  // Accept empty declaration of struct/union/enum.
				if cp.tok == CTOK_EOF && first {
					break
				} // May omit ';' for 1 decl.
				first = false
				cp_check(cp, ';')
			}
			continue
		}
		for {
			var ctypeid CTypeID
			cp_declarator(cp, &decl)
			ctypeid = cp_decl_intern(cp, &decl)
			if decl.name != "" && decl.nameid == 0 { // NYI: redeclarations are ignored.
				var ct *CType
				var id CTypeID
				if (scl & CDF_TYPEDEF) != 0 { // Create new typedef.
					id = lj_ctype_new(cp.cts, &ct)
					ct.info = CTINFO(CT_TYPEDEF, CTInfo(ctypeid))
					goto noredir
				} else if ctype_isfunc(ctype_get(cp.cts, ctypeid).info) {
					// Treat both static and extern function declarations as extern.
					ct = ctype_get(cp.cts, ctypeid)
					// We always get new anonymous functions (typedefs are copied).
					//lj_assertCP(gcref(ct.name) == NULL, "unexpected named function");
					id = ctypeid // Just name it.
				} else if (scl & CDF_STATIC) != 0 { // Accept static constants.
					id = cp_decl_constinit(cp, &ct, ctypeid)
					goto noredir
				} else { // External references have extern or no storage class.
					id = lj_ctype_new(cp.cts, &ct)
					ct.info = CTINFO(CT_EXTERN, CTInfo(ctypeid))
				}
				if decl.redir != "" { // Add attribute for redirected symbol name.
					var cta *CType
					var aid CTypeID = lj_ctype_new(cp.cts, &cta)
					ct = ctype_get(cp.cts, id) // Table may have been reallocated.
					cta.info = CTINFO(CT_ATTRIB, CTATTRIB(CTA_REDIR))
					cta.sib = ct.sib
					ct.sib = CTypeID1(aid)
					ctype_setname(cta, decl.redir)
				}
			noredir:
				ctype_setname(ct, decl.name)
				lj_ctype_addname(cp.cts, ct, id)
			}
			if !cp_opt(cp, ',') {
				break
			}
			cp_decl_reset(&decl)
		}
		if cp.tok == CTOK_EOF && first {
			break
		} // May omit ';' for 1 decl.
		first = false
		cp_check(cp, ';')
	}
}

// Handle line number.
func cp_line(cp *CPState, hashline int) {
	newline := cp.val.i32
	// TODO: Handle file name and include it in error messages.
	for cp.tok != CTOK_EOF && cp.linenumber == hashline {
		cp_next(cp)
	}
	cp.linenumber = int(newline)
}

func CTINFO(ct, flags CTInfo) CTInfo { return (ct << CTSHIFT_NUM) + flags }
func CTALIGN(al int) CTSize          { return CTSize(al) << CTSHIFT_ALIGN }
func CTATTRIB(at int) CTInfo         { return CTInfo(at) << CTSHIFT_ATTRIB }

func ctype_type(info CTInfo) CTInfo { return info >> CTSHIFT_NUM }

func ctype_cid(info CTInfo) CTInfo    { return info & CTMASK_CID }
func ctype_align(info CTInfo) CTInfo  { return (info >> CTSHIFT_ALIGN) & CTMASK_ALIGN }
func ctype_attrib(info CTInfo) CTInfo { return (info >> CTSHIFT_ATTRIB) & CTMASK_ATTRIB }
func ctype_bitpos(info CTInfo) CTInfo { return (info >> CTSHIFT_BITPOS) & CTMASK_BITPOS }
func ctype_bitbsz(info CTInfo) CTInfo { return (info >> CTSHIFT_BITBSZ) & CTMASK_BITBSZ }
func ctype_bitcsz(info CTInfo) CTInfo { return (info >> CTSHIFT_BITCSZ) & CTMASK_BITCSZ }
func ctype_vsizeP(info CTInfo) CTInfo { return (info >> CTSHIFT_VSIZEP) & CTMASK_VSIZEP }
func ctype_msizeP(info CTInfo) CTInfo { return (info >> CTSHIFT_MSIZEP) & CTMASK_MSIZEP }
func ctype_cconv(info CTInfo) CTInfo  { return (info >> CTSHIFT_CCONV) & CTMASK_CCONV }

// Simple type checks.
func ctype_isnum(info CTInfo) bool      { return ctype_type(info) == CT_NUM }
func ctype_isvoid(info CTInfo) bool     { return ctype_type(info) == CT_VOID }
func ctype_isptr(info CTInfo) bool      { return ctype_type(info) == CT_PTR }
func ctype_isarray(info CTInfo) bool    { return ctype_type(info) == CT_ARRAY }
func ctype_isstruct(info CTInfo) bool   { return ctype_type(info) == CT_STRUCT }
func ctype_isfunc(info CTInfo) bool     { return ctype_type(info) == CT_FUNC }
func ctype_isenum(info CTInfo) bool     { return ctype_type(info) == CT_ENUM }
func ctype_istypedef(info CTInfo) bool  { return ctype_type(info) == CT_TYPEDEF }
func ctype_isattrib(info CTInfo) bool   { return ctype_type(info) == CT_ATTRIB }
func ctype_isfield(info CTInfo) bool    { return ctype_type(info) == CT_FIELD }
func ctype_isbitfield(info CTInfo) bool { return ctype_type(info) == CT_BITFIELD }
func ctype_isconstval(info CTInfo) bool { return ctype_type(info) == CT_CONSTVAL }
func ctype_isextern(info CTInfo) bool   { return ctype_type(info) == CT_EXTERN }
func ctype_hassize(info CTInfo) bool    { return ctype_type(info) <= CT_HASSIZE }

// Combined type and flag checks.

func ctype_isinteger(info CTInfo) bool {
	return ((info) & (CTMASK_NUM | CTF_BOOL | CTF_FP)) == CTINFO(CT_NUM, 0)
}

func ctype_isinteger_or_bool(info CTInfo) bool {
	return (info & (CTMASK_NUM | CTF_FP)) == CTINFO(CT_NUM, 0)
}

func ctype_isbool(info CTInfo) bool {
	return ((info) & (CTMASK_NUM | CTF_BOOL)) == CTINFO(CT_NUM, CTF_BOOL)
}

func ctype_isfp(info CTInfo) bool {
	return ((info) & (CTMASK_NUM | CTF_FP)) == CTINFO(CT_NUM, CTF_FP)
}

func ctype_ispointer(info CTInfo) bool { // Pointer or array.
	return (ctype_type(info) >> 1) == (CT_PTR >> 1)
}

func ctype_isref(info CTInfo) bool {
	return ((info) & (CTMASK_NUM | CTF_REF)) == CTINFO(CT_PTR, CTF_REF)
}

func ctype_isrefarray(info CTInfo) bool {
	return (info & (CTMASK_NUM | CTF_VECTOR | CTF_COMPLEX)) == CTINFO(CT_ARRAY, 0)
}

// #define ctype_isvector(info) \
// 	(((info) & (CTMASK_NUM|CTF_VECTOR)) == CTINFO(CT_ARRAY, CTF_VECTOR))
// #define ctype_iscomplex(info) \
// 	(((info) & (CTMASK_NUM|CTF_COMPLEX)) == CTINFO(CT_ARRAY, CTF_COMPLEX))

func ctype_isvltype(info CTInfo) bool { // VL array or VL struct.
	return (info & ((CTMASK_NUM | CTF_VLA) - (2 << CTSHIFT_NUM))) == CTINFO(CT_STRUCT, CTF_VLA)
}

// #define ctype_isvlarray(info) \
// 	(((info) & (CTMASK_NUM|CTF_VLA)) == CTINFO(CT_ARRAY, CTF_VLA))

func ctype_isxattrib(info CTInfo, at int) bool {
	return (info & (CTMASK_NUM | CTATTRIB(CTMASK_ATTRIB))) == CTINFO(CT_ATTRIB, CTATTRIB(at))
}

// #define CTF_INSERT(info, field, val) \
// info = (info & ~(CTMASK_##field<<CTSHIFT_##field)) | (((CTSize)(val) & CTMASK_##field) << CTSHIFT_##field)
func CTF_INSERT(info, mask, shift, val CTInfo) CTInfo {
	return (info & ^(mask << shift)) | ((val & mask) << shift)
}

// Parse declaration attributes (and common qualifiers).
func cp_decl_attributes(cp *CPState, decl *CPDecl) {
	for {
		switch cp.tok {
		case CTOK_CONST:
			decl.attr |= CTF_CONST
			break
		case CTOK_VOLATILE:
			decl.attr |= CTF_VOLATILE
			break
		case CTOK_RESTRICT:
			break // Ignore.
		case CTOK_EXTENSION:
			break // Ignore.
		case CTOK_ATTRIBUTE:
			//cp_decl_gccattribute(cp, decl)
			//continue
			panic("fixme")
		case CTOK_ASM:
			//cp_decl_asm(cp, decl)
			//continue
			panic("fixme")
		case CTOK_DECLSPEC:
			//cp_decl_msvcattribute(cp, decl)
			//continue
			panic("fixme")
		case CTOK_CCDECL:
			// #if LJ_TARGET_X86
			// CTF_INSERT(decl.fattr, CCONV, cp.ct.size);
			// decl.fattr |= CTFP_CCONV;
			// #endif
			break
		case CTOK_PTRSZ:
			decl.attr = CTF_INSERT(decl.attr, CTMASK_MSIZEP, CTSHIFT_MSIZEP, CTInfo(cp.ct.size))
			break
		default:
			return
		}
		cp_next(cp)
	}
}

var sizeofInt = reflect.TypeOf(0).Size()

func lj_fls(x uint32) int {
	return bits.LeadingZeros32(x) ^ 31
}

// Check C type ID for validity when assertions are enabled.
func ctype_check(cts *CTState, id CTypeID) CTypeID {
	//lj_assertCTS(id > 0 && id < cts->top, "bad CTID %d", id);
	return id
}

// Get C type for C type ID.
func ctype_get(cts *CTState, id CTypeID) *CType {
	return cts.tab[ctype_check(cts, id)]
}

// Push unrolled type to declaration stack and merge qualifiers.
func cp_push_type(decl *CPDecl, id CTypeID) {
	var ct *CType = ctype_get(decl.cp.cts, id)
	var info = ct.info
	var size = ct.size
	switch ctype_type(info) {
	case CT_STRUCT, CT_ENUM:
		cp_push(decl, CTINFO(CT_TYPEDEF, CTInfo(id)), 0) // Don't copy unique types.
		if (decl.attr & CTF_QUAL) != 0 {                 // Push unmerged qualifiers.
			cp_push(decl, CTINFO(CT_ATTRIB, CTATTRIB(CTA_QUAL)), CTSize(decl.attr&CTF_QUAL))
			decl.attr &= ^CTF_QUAL
		}
		break
	case CT_ATTRIB:
		if ctype_isxattrib(info, CTA_QUAL) {
			decl.attr &= CTInfo(^size) // Remove redundant qualifiers.
		}
		cp_push_type(decl, CTypeID(ctype_cid(info))) // Unroll.
		cp_push(decl, info&^CTMASK_CID, size)        // Copy type.
		break
	case CT_ARRAY:
		if (ct.info & (CTF_VECTOR | CTF_COMPLEX)) != 0 {
			info |= decl.attr & CTF_QUAL
			decl.attr &= ^CTF_QUAL
		}
		cp_push_type(decl, CTypeID(ctype_cid(info))) // Unroll.
		cp_push(decl, info&^CTMASK_CID, size)        // Copy type.
		decl.stack[decl.pos].sib = 1                 // Mark as already checked and sized.
		// Note: this is not copied to the ct.sib in the C type table.
		break
	case CT_FUNC:
		// Copy type, link parameters (shared).
		decl.stack[cp_push(decl, info, size)].sib = ct.sib
		break
	default:
		// Copy type, merge common qualifiers.
		cp_push(decl, info|(decl.attr&CTF_QUAL), size)
		decl.attr &= ^CTF_QUAL
		break
	}
}

// Add declaration element behind the insertion position.
func cp_add(decl *CPDecl, info CTInfo, size CTSize) CPDeclIdx {
	var top = decl.top
	if top >= CPARSE_MAX_DECLSTACK {
		//cp_err(decl.cp, LJ_ERR_XLEVELS);
		panic(decl.cp)
	}
	decl.stack[top].info = info
	decl.stack[top].size = size
	decl.stack[top].sib = 0
	//setgcrefnull(decl.stack[top].name)
	decl.stack[top].next = decl.stack[decl.pos].next
	decl.stack[decl.pos].next = CTypeID1(top)
	decl.top = top + 1
	return top
}

// Push declaration element before the insertion position.
func cp_push(decl *CPDecl, info CTInfo, size CTSize) CPDeclIdx {
	decl.pos = cp_add(decl, info, size)
	return decl.pos
}

// Create new type element.
func lj_ctype_new(cts *CTState, ctp **CType) CTypeID {
	var id CTypeID = cts.top
	var ct *CType = &CType{}
	*ctp = ct
	cts.top = id + 1
	cts.tab = append(cts.tab, ct)
	ct.info = 0
	ct.size = 0
	ct.sib = 0
	ct.next = 0
	//setgcrefnull(ct.name);
	return id
}

// Set the name of a C type table element.
func ctype_setname(ct *CType, s string) {
	ct.name = s
}

// Add named element to hash table.
func lj_ctype_addname(cts *CTState, ct *CType, id CTypeID) {
	cts.nhash[ct.name] = id
}

// Parse struct/union/enum name.
func cp_struct_name(cp *CPState, sdecl *CPDecl, info CTInfo) CTypeID {
	var sid CTypeID
	var ct *CType
	cp.tmask = CPNS_STRUCT
	cp_next(cp)
	cp_decl_attributes(cp, sdecl)
	cp.tmask = CPNS_DEFAULT
	if cp.tok != '{' {
		if cp.tok != CTOK_IDENT {
			//cp_err_token(cp, CTOK_IDENT);
			panic(cp)
		}
		if cp.val.id != 0 { // Name of existing struct/union/enum.
			sid = cp.val.id
			ct = cp.ct
			if ((ct.info ^ info) & (CTMASK_NUM | CTF_UNION)) != 0 { // Wrong type.
				//cp_errmsg(cp, 0, LJ_ERR_FFI_REDEF, strdata(gco2str(gcref(ct.name))));
				panic(cp)
			}
		} else { // Create named, incomplete struct/union/enum.
			if (cp.mode & CPARSE_MODE_NOIMPLICIT) != 0 {
				//cp_errmsg(cp, 0, LJ_ERR_FFI_BADTAG, strdata(cp.str));
				panic(cp)
			}
			sid = lj_ctype_new(cp.cts, &ct)
			ct.info = info
			ct.size = CTSIZE_INVALID
			ctype_setname(ct, cp.str)
			lj_ctype_addname(cp.cts, ct, sid)
		}
		cp_next(cp)
	} else { // Create anonymous, incomplete struct/union/enum.
		sid = lj_ctype_new(cp.cts, &ct)
		ct.info = info
		ct.size = CTSIZE_INVALID
	}
	if cp.tok == '{' {
		if ct.size != CTSIZE_INVALID || (ct.sib != 0) {
			//cp_errmsg(cp, 0, LJ_ERR_FFI_REDEF, strdata(gco2str(gcref(ct.name))));
			panic(cp)
		}
		ct.sib = 1 // Indicate the type is currently being defined.
	}
	return sid
}

const CTSIZE_PTR = 8
const CTALIGN_PTR = 8

func CTINFO_REF(ref CTInfo) CTInfo {
	return CTINFO(CT_PTR, (CTF_CONST|CTF_REF|CTALIGN_PTR)+ref)
}

// Check if the next token may start a type declaration.
func cp_istypedecl(cp *CPState) bool {
	if cp.tok >= CTOK_FIRSTDECL && cp.tok <= CTOK_LASTDECL {
		return true
	}
	if cp.tok == CTOK_IDENT && ctype_istypedef(cp.ct.info) {
		return true
	}
	if cp.tok == '$' {
		return true
	}
	return false
}

// Intern a type element.
func lj_ctype_intern(cts *CTState, info CTInfo, size CTSize) CTypeID {
	t := CTInfoSize{info, size}
	if id, ok := cts.thash[t]; ok {
		return id
	}
	id := cts.top
	cts.top = id + 1
	cts.tab[id].info = info
	cts.tab[id].size = size
	cts.tab[id].sib = 0
	//cts.tab[id].next = cts.hash[h]
	cts.thash[t] = id
	return id
}

// Get child C type.
func ctype_child(cts *CTState, ct *CType) *CType {
	//lj_assertCTS(!(ctype_isvoid(ct->info) || ctype_isstruct(ct->info) || ctype_isbitfield(ct->info)), "ctype %08x has no children", ct->info);
	return ctype_get(cts, CTypeID(ctype_cid(ct.info)))
}

// Get raw type for a C type ID.
func ctype_raw(cts *CTState, id CTypeID) *CType {
	var ct *CType = ctype_get(cts, id)
	for ctype_isattrib(ct.info) {
		ct = ctype_child(cts, ct)
	}
	return ct
}

// Consume the declaration element chain and intern the C type.
func cp_decl_intern(cp *CPState, decl *CPDecl) CTypeID {
	var id CTypeID = 0
	var idx CPDeclIdx = 0
	var csize CTSize = CTSIZE_INVALID
	var cinfo CTInfo = 0
	for {
		var ct *CType = &decl.stack[idx]
		var info CTInfo = ct.info
		var size CTInfo = CTInfo(ct.size)
		// The cid is already part of info for copies of pointers/functions.
		idx = CPDeclIdx(ct.next)
		if ctype_istypedef(info) {
			//lj_assertCP(id == 0, "typedef not at toplevel")
			id = CTypeID(ctype_cid(info))
			// Always refetch info/size, since struct/enum may have been completed.
			cinfo = ctype_get(cp.cts, id).info
			csize = ctype_get(cp.cts, id).size
			//lj_assertCP(ctype_isstruct(cinfo) || ctype_isenum(cinfo), "typedef of bad type");
		} else if ctype_isfunc(info) { // Intern function.
			var fct *CType
			var fid CTypeID
			var sib CTypeID
			if id != 0 {
				var refct *CType = ctype_raw(cp.cts, id)
				// Reject function or refarray return types.
				if ctype_isfunc(refct.info) || ctype_isrefarray(refct.info) {
					//cp_err(cp, LJ_ERR_FFI_INVTYPE);
					panic(cp)
				}
			}
			// No intervening attributes allowed, skip forward.
			for idx != 0 {
				var ctn *CType = &decl.stack[idx]
				if !ctype_isattrib(ctn.info) {
					break
				}
				idx = CPDeclIdx(ctn.next) // Skip attribute.
			}
			sib = CTypeID(ct.sib) // Next line may reallocate the C type table.
			fid = lj_ctype_new(cp.cts, &fct)
			csize = CTSIZE_INVALID
			cinfo = info + CTInfo(id)
			fct.info = CTInfo(cinfo)
			fct.size = CTSize(size)
			fct.sib = CTypeID1(sib)
			id = fid
		} else if ctype_isattrib(info) {
			if ctype_isxattrib(info, CTA_QUAL) {
				cinfo |= size
			} else if ctype_isxattrib(info, CTA_ALIGN) {
				cinfo = CTF_INSERT(cinfo, CTMASK_ALIGN, CTSHIFT_ALIGN, size)
			}
			id = lj_ctype_intern(cp.cts, info+CTInfo(id), CTSize(size))
			// Inherit csize/cinfo from original type.
		} else {
			if ctype_isnum(info) { // Handle mode/vector-size attributes.
				//lj_assertCP(id == 0, "number not at toplevel")
				if (info & CTF_BOOL) == 0 {
					var msize CTSize = CTSize(ctype_msizeP(decl.attr))
					var vsize CTSize = CTSize(ctype_vsizeP(decl.attr))
					if msize != 0 && ((info&CTF_FP) == 0 || (msize == 4 || msize == 8)) {
						var malign CTSize = CTSize(lj_fls(uint32(msize)))
						if malign > 4 {
							malign = 4
						} // Limit alignment.
						info = CTF_INSERT(info, CTMASK_ALIGN, CTSHIFT_ALIGN, CTInfo(malign))
						size = CTInfo(msize) // Override size via mode.
					}
					if vsize != 0 { // Vector size set?
						var esize CTSize = CTSize(lj_fls(uint32(size)))
						if vsize >= esize {
							// Intern the element type first.
							id = lj_ctype_intern(cp.cts, info, CTSize(size))
							// Then create a vector (array) with vsize alignment.
							size = 1 << vsize
							if vsize > 4 {
								vsize = 4
							} // Limit alignment.
							if CTSize(ctype_align(info)) > vsize {
								vsize = CTSize(ctype_align(info))
							}
							info = CTINFO(CT_ARRAY, (info&CTF_QUAL)+CTF_VECTOR+CTInfo(CTALIGN(int(vsize))))
						}
					}
				}
			} else if ctype_isptr(info) {
				// Reject pointer/ref to ref.
				if id != 0 && ctype_isref(ctype_raw(cp.cts, id).info) {
					//cp_err(cp, LJ_ERR_FFI_INVTYPE);
					panic(cp)
				}
				if ctype_isref(info) {
					info &= ^CTF_VOLATILE // Refs are always const, never volatile.
					// No intervening attributes allowed, skip forward.
					for idx != 0 {
						var ctn *CType = &decl.stack[idx]
						if !ctype_isattrib(ctn.info) {
							break
						}
						idx = CPDeclIdx(ctn.next) // Skip attribute.
					}
				}
			} else if ctype_isarray(info) { // Check for valid array size etc.
				if ct.sib == 0 { // Only check/size arrays not copied by unroll.
					if ctype_isref(CTInfo(cinfo)) { // Reject arrays of refs.
						//cp_err(cp, LJ_ERR_FFI_INVTYPE)
						panic(cp)
					}
					// Reject VLS or unknown-sized types.
					if ctype_isvltype(cinfo) || csize == CTSIZE_INVALID {
						//cp_err(cp, LJ_ERR_FFI_INVSIZE)
						panic(cp)
					}
					// a[] and a[?] keep their invalid size.
					if size != CTSIZE_INVALID {
						var xsz uint64 = uint64(size * CTInfo(csize))
						if xsz >= 0x80000000 {
							//cp_err(cp, LJ_ERR_FFI_INVSIZE);
							panic(cp)
						}
						size = CTInfo(xsz)
					}
				}
				if (cinfo & CTF_ALIGN) > (info & CTF_ALIGN) { // Find max. align.
					info = (info & ^CTF_ALIGN) | (cinfo & CTF_ALIGN)
				}
				info |= cinfo & CTF_QUAL // Inherit qual.
			} else {
				//lj_assertCP(ctype_isvoid(info), "bad ctype %08x", info);
			}
			csize = CTSize(size)
			cinfo = info + CTInfo(id)
			id = lj_ctype_intern(cp.cts, CTInfo(cinfo), CTSize(size))
		}
		if idx == 0 {
			break
		}
	}
	return id
}

// Parse function declaration.
func cp_decl_func(cp *CPState, fdecl *CPDecl) {
	var nargs CTSize = 0
	var info CTInfo = CTINFO(CT_FUNC, 0)
	var lastid CTypeID = 0
	var anchor CTypeID = 0
	if cp.tok != ')' {
		for {
			var decl CPDecl
			var ctypeid CTypeID
			var fieldid CTypeID
			var ct *CType
			if cp_opt(cp, '.') { // Vararg function.
				cp_check(cp, '.') // Workaround for the minimalistic lexer.
				cp_check(cp, '.')
				info |= CTF_VARARG
				break
			}
			cp_decl_spec(cp, &decl, CDF_REGISTER)
			decl.mode = CPARSE_MODE_DIRECT | CPARSE_MODE_ABSTRACT
			cp_declarator(cp, &decl)
			ctypeid = cp_decl_intern(cp, &decl)
			ct = ctype_raw(cp.cts, ctypeid)
			if ctype_isvoid(ct.info) {
				break
			} else if ctype_isrefarray(ct.info) {
				ctypeid = lj_ctype_intern(cp.cts, CTINFO(CT_PTR, CTALIGN_PTR|ctype_cid(ct.info)), CTSIZE_PTR)
			} else if ctype_isfunc(ct.info) {
				ctypeid = lj_ctype_intern(cp.cts, CTINFO(CT_PTR, CTInfo(CTALIGN_PTR|ctypeid)), CTSIZE_PTR)
			}
			// Add new parameter.
			fieldid = lj_ctype_new(cp.cts, &ct)
			if anchor != 0 {
				ctype_get(cp.cts, lastid).sib = CTypeID1(fieldid)
			} else {
				anchor = fieldid
			}
			lastid = fieldid
			if decl.name != "" {
				ctype_setname(ct, decl.name)
			}
			ct.info = CTINFO(CT_FIELD, CTInfo(ctypeid))
			ct.size = nargs
			nargs++
			if !cp_opt(cp, ',') {
				break
			}
		}
	}
	cp_check(cp, ')')
	if cp_opt(cp, '{') { // Skip function definition.
		var level = 1
		cp.mode |= CPARSE_MODE_SKIP
		for {
			if cp.tok == '{' {
				level++
			} else {
				level--
				if cp.tok == '}' && level == 0 {
					break
				} else if cp.tok == CTOK_EOF {
					//cp_err_token(cp, '}');
					panic(cp)
				}
			}
			cp_next(cp)
		}
		cp.mode &= ^CPARSE_MODE_SKIP
		cp.tok = ';' // Ok for cp_decl_multi(), error in cp_decl_single().
	}
	info |= fdecl.fattr & ^CTMASK_CID
	fdecl.fattr = 0
	fdecl.stack[cp_add(fdecl, info, nargs)].sib = CTypeID1(anchor)
}

// Please note that type handling is very weak here. Most ops simply
// assume integer operands. Accessors are only needed to compute types and
// return synthetic values. The only purpose of the expression evaluator
// is to compute the values of constant expressions one would typically
// find in C header files. And again: this is NOT a validating C parser!

// Parse comma separated expression and return last result.
func cp_expr_comma(cp *CPState, k *CPValue) {
	for {
		cp_expr_sub(cp, k, 0)
		if !cp_opt(cp, ',') {
			break
		}
	}
}

// Get type, qualifiers, size and alignment for a C type ID.
func lj_ctype_info(cts *CTState, id CTypeID, szp *CTSize) CTInfo {
	var qual CTInfo = 0
	var ct *CType = ctype_get(cts, id)
	for {
		var info CTInfo = ct.info
		if ctype_isenum(info) {
			// Follow child. Need to look at its attributes, too.
		} else if ctype_isattrib(info) {
			if ctype_isxattrib(info, CTA_QUAL) {
				qual |= CTInfo(ct.size)
			} else if ctype_isxattrib(info, CTA_ALIGN) && (qual&CTFP_ALIGNED) == 0 {
				qual |= CTInfo(CTFP_ALIGNED + CTALIGN(int(ct.size)))
			}
		} else {
			if (qual & CTFP_ALIGNED) == 0 {
				qual |= info & CTF_ALIGN
			}
			qual |= info & ^(CTF_ALIGN | CTMASK_CID)
			//lj_assertCTS(ctype_hassize(info) || ctype_isfunc(info), "ctype without size");
			*szp = bt.Choose(ctype_isfunc(info), CTSIZE_INVALID, ct.size)
			break
		}
		ct = ctype_get(cts, CTypeID(ctype_cid(info)))
	}
	return qual
}

// Parse an abstract type declaration and return it's C type ID.
func cp_decl_abstract(cp *CPState) CTypeID {
	var decl CPDecl
	cp_decl_spec(cp, &decl, 0)
	decl.mode = CPARSE_MODE_ABSTRACT
	cp_declarator(cp, &decl)
	return cp_decl_intern(cp, &decl)
}

// Parse sizeof/alignof operator.
func cp_expr_sizeof(cp *CPState, k *CPValue, wantsz int) {
	var sz CTSize
	var info CTInfo
	if cp_opt(cp, '(') {
		if cp_istypedecl(cp) {
			k.id = cp_decl_abstract(cp)
		} else {
			cp_expr_comma(cp, k)
		}
		cp_check(cp, ')')
	} else {
		cp_expr_unary(cp, k)
	}
	info = lj_ctype_info(cp.cts, k.id, &sz)
	if wantsz != 0 {
		if sz != CTSIZE_INVALID {
			k.i32 = int32(sz)
		} else if k.id != CTID_A_CCHAR { // Special case for sizeof("string").
			//cp_err(cp, LJ_ERR_FFI_INVSIZE)
			panic(cp)
		}
	} else {
		k.i32 = 1 << ctype_align(info)
	}
	k.id = CTID_UINT32 // Really size_t.
}

// Parse prefix operators.
func cp_expr_prefix(cp *CPState, k *CPValue) {
	if cp.tok == CTOK_INTEGER {
		*k = cp.val
		cp_next(cp)
	} else if cp_opt(cp, '+') {
		cp_expr_unary(cp, k) // Nothing to do (well, integer promotion).
	} else if cp_opt(cp, '-') {
		cp_expr_unary(cp, k)
		k.i32 = -k.i32
	} else if cp_opt(cp, '~') {
		cp_expr_unary(cp, k)
		k.i32 = ^k.i32
	} else if cp_opt(cp, '!') {
		cp_expr_unary(cp, k)
		k.i32 = bt.Choose[int32](k.i32 != 0, 0, 1)
		k.id = CTID_INT32
	} else if cp_opt(cp, '(') {
		if cp_istypedecl(cp) { // Cast operator.
			var id CTypeID = cp_decl_abstract(cp)
			cp_check(cp, ')')
			cp_expr_unary(cp, k)
			k.id = id // No conversion performed.
		} else { // Sub-expression.
			cp_expr_comma(cp, k)
			cp_check(cp, ')')
		}
	} else if cp_opt(cp, '*') { // Indirection.
		var ct *CType
		cp_expr_unary(cp, k)
		ct = lj_ctype_rawref(cp.cts, k.id)
		if !ctype_ispointer(ct.info) {
			//cp_err_badidx(cp, ct)
			panic(cp)
		}
		k.i32 = 0
		k.id = CTypeID(ctype_cid(ct.info))
	} else if cp_opt(cp, '&') { // Address operator.
		cp_expr_unary(cp, k)
		k.id = lj_ctype_intern(cp.cts, CTINFO(CT_PTR, CTInfo(CTALIGN_PTR+k.id)), CTSIZE_PTR)
	} else if cp_opt(cp, CTOK_SIZEOF) {
		cp_expr_sizeof(cp, k, 1)
	} else if cp_opt(cp, CTOK_ALIGNOF) {
		cp_expr_sizeof(cp, k, 0)
	} else if cp.tok == CTOK_IDENT {
		if ctype_type(cp.ct.info) == CT_CONSTVAL {
			k.i32 = int32(cp.ct.size)
			k.id = CTypeID(ctype_cid(cp.ct.info))
		} else if ctype_type(cp.ct.info) == CT_EXTERN {
			k.i32 = int32(cp.val.id)
			k.id = CTypeID(ctype_cid(cp.ct.info))
		} else if ctype_type(cp.ct.info) == CT_FUNC {
			k.i32 = int32(cp.val.id)
			k.id = cp.val.id
		} else {
			//goto err_expr
			panic(cp)
		}
		cp_next(cp)
	} else if cp.tok == CTOK_STRING {
		var sz CTSize = CTSize(len(cp.str))
		for cp_next(cp) == CTOK_STRING {
			sz += CTSize(len(cp.str))
		}
		k.i32 = int32(sz) + 1
		k.id = CTID_A_CCHAR
	} else {
		//cp_errmsg(cp, cp.tok, LJ_ERR_XSYMBOL);
		panic(cp)
	}
}

// Follow references and get raw type for a C type ID.
func lj_ctype_rawref(cts *CTState, id CTypeID) *CType {
	var ct *CType = ctype_get(cts, id)
	for ctype_isattrib(ct.info) || ctype_isref(ct.info) {
		ct = ctype_child(cts, ct)
	}
	return ct
}

// Get a struct/union/enum/function field by name.
func lj_ctype_getfieldq(cts *CTState, ct *CType, name string, ofs *CTSize, qual *CTInfo) *CType {
	for ct.sib != 0 {
		ct = ctype_get(cts, CTypeID(ct.sib))
		if ct.name == name {
			*ofs = ct.size
			return ct
		}
		if ctype_isxattrib(ct.info, CTA_SUBTYPE) {
			var fct *CType
			cct := ctype_child(cts, ct)
			var q CTInfo = 0
			for ctype_isattrib(cct.info) {
				if ctype_attrib(cct.info) == CTA_QUAL {
					q |= CTInfo(cct.size)
				}
				cct = ctype_child(cts, cct)
			}
			fct = lj_ctype_getfieldq(cts, cct, name, ofs, qual)
			if fct != nil {
				if qual != nil {
					*qual |= q
				}
				*ofs += ct.size
				return fct
			}
		}
	}
	return nil // Not found.
}

func lj_ctype_getfield(cts *CTState, ct *CType, name string, ofs *CTSize) *CType {
	return lj_ctype_getfieldq(cts, ct, name, ofs, nil)
}

// Parse postfix operators.
func cp_expr_postfix(cp *CPState, k *CPValue) {
	for {
		var ct *CType
		if cp_opt(cp, '[') { // Array/pointer index.
			var k2 CPValue
			cp_expr_comma(cp, &k2)
			ct = lj_ctype_rawref(cp.cts, k.id)
			if !ctype_ispointer(ct.info) {
				ct = lj_ctype_rawref(cp.cts, k2.id)
				if !ctype_ispointer(ct.info) {
					//cp_err_badidx(cp, ct);
					panic(cp)
				}
			}
			cp_check(cp, ']')
			k.i32 = 0
		} else if cp.tok == '.' || cp.tok == CTOK_DEREF { // Struct deref.
			var ofs CTSize
			var fct *CType
			ct = lj_ctype_rawref(cp.cts, k.id)
			if cp.tok == CTOK_DEREF {
				if !ctype_ispointer(ct.info) {
					//cp_err_badidx(cp, ct);
					panic(cp)
				}
				ct = lj_ctype_rawref(cp.cts, CTypeID(ctype_cid(ct.info)))
			}
			cp_next(cp)
			if cp.tok != CTOK_IDENT {
				//cp_err_token(cp, CTOK_IDENT);
				panic(cp)
			}
			if !ctype_isstruct(ct.info) || ct.size == CTSIZE_INVALID {
				panic(cp)
			}
			fct = lj_ctype_getfield(cp.cts, ct, cp.str, &ofs)
			if fct == nil || ctype_isbitfield(fct.info) {
				//s := lj_ctype_repr(cp.cts.L, ctype_typeid(cp.cts, ct), NULL);
				//cp_errmsg(cp, 0, LJ_ERR_FFI_BADMEMBER, strdata(s), strdata(cp.str));
				panic(cp)
			}
			ct = fct
			k.i32 = bt.Choose(ctype_isconstval(ct.info), int32(ct.size), 0)
			cp_next(cp)
		} else {
			return
		}
		k.id = CTypeID(ctype_cid(ct.info))
	}
}

func i32b(i int32) bool {
	if i != 0 {
		return true
	}
	return false
}

func bi32(b bool) int32 {
	if b {
		return 1
	}
	return 0
}

// Parse infix operators.
func cp_expr_infix(cp *CPState, k *CPValue, pri int) {
	var k2 CPValue
	k2.i32 = 0
	k2.id = 0 // Silence the compiler.
	for {
		switch pri {
		case 0:
			if cp_opt(cp, '?') {
				var k3 CPValue
				cp_expr_comma(cp, &k2) // Right-associative.
				cp_check(cp, ':')
				cp_expr_sub(cp, &k3, 0)
				k.i32 = bt.Choose(k.i32 != 0, k2.i32, k3.i32)
				k.id = bt.Choose(k2.id > k3.id, k2.id, k3.id)
				continue
			}
			// fallthrough
		case 1:
			if cp_opt(cp, CTOK_OROR) {
				cp_expr_sub(cp, &k2, 2)
				k.i32 = bi32(i32b(k.i32) || i32b(k2.i32))
				k.id = CTID_INT32
				continue
			}
			// fallthrough
		case 2:
			if cp_opt(cp, CTOK_ANDAND) {
				cp_expr_sub(cp, &k2, 3)
				k.i32 = bi32(i32b(k.i32) && i32b(k2.i32))
				k.id = CTID_INT32
				continue
			}
			// fallthrough
		case 3:
			if cp_opt(cp, '|') {
				cp_expr_sub(cp, &k2, 4)
				k.i32 = k.i32 | k2.i32
				if k2.id > k.id {
					k.id = k2.id
				} // Trivial promotion to unsigned.
				continue
			}
			// fallthrough
		case 4:
			if cp_opt(cp, '^') {
				cp_expr_sub(cp, &k2, 5)
				k.i32 = k.i32 ^ k2.i32
				if k2.id > k.id {
					k.id = k2.id
				} // Trivial promotion to unsigned.
				continue
			}
			// fallthrough
		case 5:
			if cp_opt(cp, '&') {
				cp_expr_sub(cp, &k2, 6)
				k.i32 = k.i32 & k2.i32
				if k2.id > k.id {
					k.id = k2.id
				} // Trivial promotion to unsigned.
				continue
			}
			// fallthrough
		case 6:
			if cp_opt(cp, CTOK_EQ) {
				cp_expr_sub(cp, &k2, 7)
				k.i32 = bi32(k.i32 == k2.i32)
				k.id = CTID_INT32
				continue
			} else if cp_opt(cp, CTOK_NE) {
				cp_expr_sub(cp, &k2, 7)
				k.i32 = bi32(k.i32 != k2.i32)
				k.id = CTID_INT32
				continue
			}
			// fallthrough
		case 7:
			if cp_opt(cp, '<') {
				cp_expr_sub(cp, &k2, 8)
				if k.id == CTID_INT32 && k2.id == CTID_INT32 {
					k.i32 = bi32(k.i32 < k2.i32)
				} else {
					k.i32 = bi32(uint32(k.i32) < uint32(k2.i32))
				}
				k.id = CTID_INT32
				continue
			} else if cp_opt(cp, '>') {
				cp_expr_sub(cp, &k2, 8)
				if k.id == CTID_INT32 && k2.id == CTID_INT32 {
					k.i32 = bi32(k.i32 > k2.i32)
				} else {
					k.i32 = bi32(uint32(k.i32) > uint32(k2.i32))
				}
				k.id = CTID_INT32
				continue
			} else if cp_opt(cp, CTOK_LE) {
				cp_expr_sub(cp, &k2, 8)
				if k.id == CTID_INT32 && k2.id == CTID_INT32 {
					k.i32 = bi32(k.i32 <= k2.i32)
				} else {
					k.i32 = bi32(uint32(k.i32) <= uint32(k2.i32))
				}
				k.id = CTID_INT32
				continue
			} else if cp_opt(cp, CTOK_GE) {
				cp_expr_sub(cp, &k2, 8)
				if k.id == CTID_INT32 && k2.id == CTID_INT32 {
					k.i32 = bi32(k.i32 >= k2.i32)
				} else {
					k.i32 = bi32(uint32(k.i32) >= uint32(k2.i32))
				}
				k.id = CTID_INT32
				continue
			}
			// fallthrough
		case 8:
			if cp_opt(cp, CTOK_SHL) {
				cp_expr_sub(cp, &k2, 9)
				k.i32 = k.i32 << k2.i32
				continue
			} else if cp_opt(cp, CTOK_SHR) {
				cp_expr_sub(cp, &k2, 9)
				if k.id == CTID_INT32 {
					k.i32 = k.i32 >> k2.i32
				} else {
					k.i32 = int32(uint32(k.i32) >> uint32(k2.i32))
				}
				continue
			}
			// fallthrough
		case 9:
			if cp_opt(cp, '+') {
				cp_expr_sub(cp, &k2, 10)
				k.i32 = k.i32 + k2.i32
				if k2.id > k.id {
					k.id = k2.id
				} // Trivial promotion to unsigned.
				continue
			} else if cp_opt(cp, '-') {
				cp_expr_sub(cp, &k2, 10)
				k.i32 = k.i32 - k2.i32
				if k2.id > k.id {
					k.id = k2.id
				} // Trivial promotion to unsigned.
				continue
			}
			// fallthrough
		case 10:
			if cp_opt(cp, '*') {
				cp_expr_unary(cp, &k2)
				k.i32 = k.i32 * k2.i32
				if k2.id > k.id {
					k.id = k2.id
				} // Trivial promotion to unsigned.
				continue
			} else if cp_opt(cp, '/') {
				cp_expr_unary(cp, &k2)
				if k2.id > k.id {
					// Trivial promotion to unsigned.
					k.id = k2.id
				}
				if k2.i32 == 0 || (k.id == CTID_INT32 && uint32(k.i32) == 0x80000000 && k2.i32 == -1) {
					//cp_err(cp, LJ_ERR_BADVAL);
					panic(cp)
				}
				if k.id == CTID_INT32 {
					k.i32 = k.i32 / k2.i32
				} else {
					k.i32 = int32(uint32(k.i32) / uint32(k2.i32))
				}
				continue
			} else if cp_opt(cp, '%') {
				cp_expr_unary(cp, &k2)
				if k2.id > k.id {
					// Trivial promotion to unsigned.
					k.id = k2.id
				}
				if k2.i32 == 0 || (k.id == CTID_INT32 && uint32(k.i32) == 0x80000000 && k2.i32 == -1) {
					//cp_err(cp, LJ_ERR_BADVAL);
					panic(cp)
				}
				if k.id == CTID_INT32 {
					k.i32 = k.i32 % k2.i32
				} else {
					k.i32 = int32(uint32(k.i32) % uint32(k2.i32))
				}
				continue
			}
		default:
			return
		}
	}
}

// Parse and evaluate unary expression.
func cp_expr_unary(cp *CPState, k *CPValue) {
	cp.depth++
	if cp.depth > CPARSE_MAX_DECLDEPTH {
		//cp_err(cp, LJ_ERR_XLEVELS);
		panic(cp)
	}
	cp_expr_prefix(cp, k)
	cp_expr_postfix(cp, k)
	cp.depth--
}

// Parse and evaluate sub-expression.
func cp_expr_sub(cp *CPState, k *CPValue, pri int) {
	cp_expr_unary(cp, k)
	cp_expr_infix(cp, k, pri)
}

// Parse constant integer expression.
func cp_expr_kint(cp *CPState, k *CPValue) {
	var ct *CType
	cp_expr_sub(cp, k, 0)
	ct = ctype_raw(cp.cts, k.id)
	if !ctype_isinteger(ct.info) {
		//cp_err(cp, LJ_ERR_BADVAL)
		panic(cp)
	}
}

// Parse (non-negative) size expression.
func cp_expr_ksize(cp *CPState) CTSize {
	var k CPValue
	cp_expr_kint(cp, &k)
	if uint(k.i32) >= 0x80000000 {
		//cp_err(cp, LJ_ERR_FFI_INVSIZE)
		panic(cp)
	}
	return CTSize(k.i32)
}

// Parse array declaration.
func cp_decl_array(cp *CPState, decl *CPDecl) {
	var info CTInfo = CTINFO(CT_ARRAY, 0)
	var nelem CTSize = CTSIZE_INVALID // Default size for a[] or a[?].
	cp_decl_attributes(cp, decl)
	if cp_opt(cp, '?') {
		info |= CTF_VLA // Create variable-length array a[?].
	} else if cp.tok != ']' {
		nelem = cp_expr_ksize(cp)
	}
	cp_check(cp, ']')
	cp_add(decl, info, nelem)
}

// Parse declarator.
func cp_declarator(cp *CPState, decl *CPDecl) {
	cp.depth++
	if cp.depth > CPARSE_MAX_DECLDEPTH {
		//cp_err(cp, LJ_ERR_XLEVELS);
		panic(cp)
	}

	for { // Head of declarator.
		if cp_opt(cp, '*') { // Pointer.
			var sz CTSize
			var info CTInfo
			cp_decl_attributes(cp, decl)
			sz = CTSIZE_PTR
			info = CTINFO(CT_PTR, CTALIGN_PTR)
			if ctype_msizeP(decl.attr) == 4 {
				sz = 4
				info = CTINFO(CT_PTR, CTInfo(CTALIGN(2)))
			}
			info += decl.attr & (CTF_QUAL | CTF_REF)
			decl.attr &= ^(CTF_QUAL | (CTMASK_MSIZEP << CTSHIFT_MSIZEP))
			cp_push(decl, info, sz)
		} else if cp_opt(cp, '&') || cp_opt(cp, CTOK_ANDAND) { // Reference.
			decl.attr &= ^(CTF_QUAL | (CTMASK_MSIZEP << CTSHIFT_MSIZEP))
			cp_push(decl, CTINFO_REF(0), CTSIZE_PTR)
		} else {
			break
		}
	}

	if cp_opt(cp, '(') { // Inner declarator.
		var pos CPDeclIdx
		cp_decl_attributes(cp, decl)
		// Resolve ambiguity between inner declarator and 1st function parameter.
		if (decl.mode&CPARSE_MODE_ABSTRACT) != 0 && (cp.tok == ')' || cp_istypedecl(cp)) {
			//goto func_decl
			cp_decl_func(cp, decl)
		} else {
			pos = decl.pos
			cp_declarator(cp, decl)
			cp_check(cp, ')')
			decl.pos = pos
		}
	} else if cp.tok == CTOK_IDENT { // Direct declarator.
		if (decl.mode & CPARSE_MODE_DIRECT) == 0 {
			//cp_err_token(cp, CTOK_EOF)
			panic(cp)
		}
		decl.name = cp.str
		decl.nameid = cp.val.id
		cp_next(cp)
	} else { // Abstract declarator.
		if (decl.mode & CPARSE_MODE_ABSTRACT) == 0 {
			//cp_err_token(cp, CTOK_IDENT);
			panic(cp)
		}
	}

	for { // Tail of declarator.
		if cp_opt(cp, '[') { // Array.
			cp_decl_array(cp, decl)
		} else if cp_opt(cp, '(') { // Function.
			//func_decl:
			cp_decl_func(cp, decl)
		} else {
			break
		}
	}

	if (decl.mode&CPARSE_MODE_FIELD) != 0 && cp_opt(cp, ':') { // Field width.
		decl.bits = cp_expr_ksize(cp)
	}

	// Process postfix attributes.
	cp_decl_attributes(cp, decl)
	cp_push_attributes(decl)

	cp.depth--
}

// Push or merge attributes.
func cp_push_attributes(decl *CPDecl) {
	var ct *CType = &decl.stack[decl.pos]
	if ctype_isfunc(ct.info) { // Ok to modify in-place.
	} else {
		if (decl.attr&CTFP_ALIGNED) != 0 && (decl.mode&CPARSE_MODE_FIELD) == 0 {
			cp_push(decl, CTINFO(CT_ATTRIB, CTATTRIB(CTA_ALIGN)), CTSize(ctype_align(decl.attr)))
		}
	}
}

// Parse constant initializer.
// NYI: FP constants and strings as initializers.
func cp_decl_constinit(cp *CPState, ctp **CType, ctypeid CTypeID) CTypeID {
	var ctt *CType = ctype_get(cp.cts, ctypeid)
	var info CTInfo
	var size CTSize
	var k CPValue
	var constid CTypeID
	for ctype_isattrib(ctt.info) { // Skip attributes.
		ctypeid = CTypeID(ctype_cid(ctt.info)) // Update ID, too.
		ctt = ctype_get(cp.cts, ctypeid)
	}
	info = ctt.info
	size = ctt.size
	if !ctype_isinteger(info) || (info&CTF_CONST) == 0 || size > 4 {
		//cp_err(cp, LJ_ERR_FFI_INVTYPE);
		panic(cp)
	}
	cp_check(cp, '=')
	cp_expr_sub(cp, &k, 0)
	constid = lj_ctype_new(cp.cts, ctp)
	(*ctp).info = CTINFO(CT_CONSTVAL, CTF_CONST|CTInfo(ctypeid))
	k.i32 <<= 8 * (4 - size)
	if (info & CTF_UNSIGNED) != 0 {
		k.i32 >>= int32(uint32(k.i32) >> (uint32(8) * (4 - uint32(size))))
	} else {
		k.i32 >>= (8 * (4 - size))
	}
	(*ctp).size = CTSize(k.i32)
	return constid
}

// Parse struct/union declaration.
func cp_decl_struct(cp *CPState, sdecl *CPDecl, sinfo CTInfo) CTypeID {
	var sid CTypeID = cp_struct_name(cp, sdecl, sinfo)
	if cp_opt(cp, '{') { // Struct/union definition.
		var lastid CTypeID = sid
		var lastdecl = 0
		for cp.tok != '}' {
			var decl CPDecl
			var scl CPscl = cp_decl_spec(cp, &decl, CDF_STATIC)
			decl.mode = bt.Choose(scl != 0, CPARSE_MODE_DIRECT, CPARSE_MODE_DIRECT|CPARSE_MODE_ABSTRACT|CPARSE_MODE_FIELD)

			for {
				var ctypeid CTypeID
				if lastdecl != 0 {
					//cp_err_token(cp, '}')
					panic(cp)
				}

				// Parse field declarator.
				decl.bits = CTSIZE_INVALID
				cp_declarator(cp, &decl)
				ctypeid = cp_decl_intern(cp, &decl)

				if (scl & CDF_STATIC) != 0 { // Static constant in struct namespace.
					var ct *CType
					var fieldid CTypeID = cp_decl_constinit(cp, &ct, ctypeid)
					ctype_get(cp.cts, lastid).sib = CTypeID1(fieldid)
					lastid = fieldid
					ctype_setname(ct, decl.name)
				} else {
					var bsz CTSize = CTBSZ_FIELD // Temp. for layout phase.
					var ct *CType
					var fieldid CTypeID = lj_ctype_new(cp.cts, &ct) // Do this first.
					var tct *CType = ctype_raw(cp.cts, ctypeid)

					if decl.bits == CTSIZE_INVALID { // Regular field.
						if ctype_isarray(tct.info) && tct.size == CTSIZE_INVALID {
							lastdecl = 1 // a[] or a[?] must be the last declared field.
						}

						// Accept transparent struct/union/enum.
						if decl.name == "" {
							if !((ctype_isstruct(tct.info) && (tct.info&CTF_VLA) == 0) || ctype_isenum(tct.info)) {
								//cp_err_token(cp, CTOK_IDENT);
								panic(cp)
							}
							ct.info = CTINFO(CT_ATTRIB, CTATTRIB(CTA_SUBTYPE)+CTInfo(ctypeid))
							ct.size = bt.Choose(ctype_isstruct(tct.info), CTSize(decl.attr|0x80000000), 0) // For layout phase.
							goto add_field
						}
					} else { // Bitfield.
						bsz = decl.bits
						if !ctype_isinteger_or_bool(tct.info) ||
							(bsz == 0 && decl.name != "") ||
							8*tct.size > CTBSZ_MAX ||
							bsz > (bt.Choose(tct.info&CTF_BOOL != 0, 1, 8*tct.size)) {
							//cp_errmsg(cp, ':', LJ_ERR_BADVAL)
							panic(cp)
						}
					}

					// Create temporary field for layout phase.
					ct.info = CTINFO(CT_FIELD, CTInfo(ctypeid)+(CTInfo(bsz)<<CTSHIFT_BITCSZ))
					ct.size = CTSize(decl.attr)
					if decl.name != "" {
						ctype_setname(ct, decl.name)
					}

				add_field:
					ctype_get(cp.cts, lastid).sib = CTypeID1(fieldid)
					lastid = fieldid
				}
				if !cp_opt(cp, ',') {
					break
				}
				cp_decl_reset(&decl)
			}
			cp_check(cp, ';')
		}
		cp_check(cp, '}')
		ctype_get(cp.cts, lastid).sib = 0 // Drop sib = 1 for empty structs.
		cp_decl_attributes(cp, sdecl)     // Layout phase needs postfix attributes.
		cp_struct_layout(cp, sid, sdecl.attr)
	}
	return sid
}

// Determine field alignment.
func cp_field_align(cp *CPState, ct *CType, info CTInfo) CTSize {
	var align CTSize = CTSize(ctype_align(info))
	return align
}

// Layout struct/union fields.
func cp_struct_layout(cp *CPState, sid CTypeID, sattr CTInfo) {
	var bofs CTSize = 0
	var bmaxofs CTSize = 0 // Bit offset and max. bit offset.
	var maxalign CTSize = CTSize(ctype_align(sattr))
	var sct *CType = ctype_get(cp.cts, sid)
	var sinfo CTInfo = sct.info
	var fieldid CTypeID = CTypeID(sct.sib)
	for fieldid != 0 {
		var ct *CType = ctype_get(cp.cts, fieldid)
		var attr CTInfo = CTInfo(ct.size) // Field declaration attributes (temp.).

		if ctype_isfield(ct.info) || (ctype_isxattrib(ct.info, CTA_SUBTYPE) && (attr != 0)) {
			var align, amask CTSize // Alignment (pow2) and alignment mask (bits).
			var sz CTSize
			var info CTInfo = lj_ctype_info(cp.cts, CTypeID(ctype_cid(ct.info)), &sz)
			var bsz CTSize
			var csz CTSize = 8 * sz              // Field size and container size (in bits).
			sinfo |= info & (CTF_QUAL | CTF_VLA) // Merge pseudo-qualifiers.

			// Check for size overflow and determine alignment.
			if sz >= 0x20000000 || bofs+csz < bofs || (info&CTF_VLA) != 0 {
				if !(sz == CTSIZE_INVALID && ctype_isarray(info) && (sinfo&CTF_UNION) == 0) {
					//cp_err(cp, LJ_ERR_FFI_INVSIZE);
					panic(cp)
				}
				csz = 0
				sz = 0 // Treat a[] and a[?] as zero-sized.
			}
			align = cp_field_align(cp, ct, info)
			if ((attr|sattr)&CTFP_PACKED) != 0 || ((attr&CTFP_ALIGNED) != 0 && CTSize(ctype_align(attr)) > align) {
				align = CTSize(ctype_align(attr))
			}
			if cp.packstack[cp.curpack] < uint8(align) {
				align = CTSize(cp.packstack[cp.curpack])
			}
			if align > maxalign {
				maxalign = align
			}
			amask = (8 << align) - 1

			bsz = CTSize(ctype_bitcsz(ct.info)) // Bitfield size (temp.).
			if bsz == CTBSZ_FIELD || !ctype_isfield(ct.info) {
				bsz = csz                      // Regular fields or subtypes always fill the container.
				bofs = (bofs + amask) & ^amask // Start new aligned field.
				ct.size = (bofs >> 3)          // Store field offset.
			} else { // Bitfield.
				if bsz == 0 || (attr&CTFP_ALIGNED) != 0 || (((attr|sattr)&CTFP_PACKED) == 0 && (bofs&amask)+bsz > csz) {
					bofs = (bofs + amask) & ^amask // Start new aligned field.
				}

				// Prefer regular field over bitfield.
				if bsz == csz && (bofs&amask) == 0 {
					ct.info = CTINFO(CT_FIELD, ctype_cid(ct.info))
					ct.size = bofs >> 3 // Store field offset.
				} else {
					ct.info = CTINFO(CT_BITFIELD, (info&(CTF_QUAL|CTF_UNSIGNED|CTF_BOOL))+(CTInfo(csz)<<(CTInfo(CTSHIFT_BITCSZ)-3))+(CTInfo(bsz)<<CTInfo(CTSHIFT_BITBSZ)))
					ct.info += CTInfo((bofs & (csz - 1)) << CTSHIFT_BITPOS)
					ct.size = (bofs & ^(csz - 1)) >> 3 // Store container offset.
				}
			}

			// Determine next offset or max. offset.
			if (sinfo & CTF_UNION) != 0 {
				if bsz > bmaxofs {
					bmaxofs = bsz
				}
			} else {
				bofs += bsz
			}
		} // All other fields in the chain are already set up.

		fieldid = CTypeID(ct.sib)
	}

	// Complete struct/union.
	sct.info = sinfo + CTInfo(CTALIGN(int(maxalign)))
	bofs = bt.Choose((sinfo&CTF_UNION) != 0, bmaxofs, bofs)
	maxalign = (8 << maxalign) - 1
	sct.size = ((bofs + maxalign) & ^maxalign) >> 3
}

// Reset declaration state to declaration specifier.
func cp_decl_reset(decl *CPDecl) {
	decl.pos = decl.specpos
	decl.top = decl.specpos + 1
	decl.stack[decl.specpos].next = 0
	decl.attr = decl.specattr
	decl.fattr = decl.specfattr
	decl.name = ""
	decl.redir = ""
}

// Parse enum declaration.
func cp_decl_enum(cp *CPState, sdecl *CPDecl) CTypeID {
	var eid CTypeID = cp_struct_name(cp, sdecl, CTINFO(CT_ENUM, CTID_VOID))
	var einfo CTInfo = CTINFO(CT_ENUM, CTInfo(CTALIGN(2)+CTID_UINT32))
	var esize CTSize = 4 // Only 32 bit enums are supported.
	if cp_opt(cp, '{') { // Enum definition.
		var k CPValue
		var lastid CTypeID = eid
		k.i32 = 0
		k.id = CTID_INT32
		for {
			name := cp.str
			if cp.tok != CTOK_IDENT {
				//cp_err_token(cp, CTOK_IDENT);
				panic(cp)
			}
			if cp.val.id != 0 {
				//cp_errmsg(cp, 0, LJ_ERR_FFI_REDEF, strdata(name));
				panic(cp)
			}
			cp_next(cp)
			if cp_opt(cp, '=') {
				cp_expr_kint(cp, &k)
				if k.id == CTID_UINT32 {
					// C99 says that enum constants are always (signed) integers.
					// But since unsigned constants like 0x80000000 are quite common,
					// those are left as uint32_t.

					if k.i32 >= 0 {
						k.id = CTID_INT32
					}
				} else {
					// OTOH it's common practice and even mandated by some ABIs
					// that the enum type itself is unsigned, unless there are any
					// negative constants.

					k.id = CTID_INT32
					if k.i32 < 0 {
						einfo = CTINFO(CT_ENUM, CTInfo(CTALIGN(2)+CTID_INT32))
					}
				}
			}
			// Add named enum constant.
			{
				var ct *CType
				var constid CTypeID = lj_ctype_new(cp.cts, &ct)
				ctype_get(cp.cts, lastid).sib = CTypeID1(constid)
				lastid = constid
				ctype_setname(ct, name)
				ct.info = CTINFO(CT_CONSTVAL, CTF_CONST|CTInfo(k.id))
				k.i32++
				ct.size = CTSize(k.i32)
				if uint32(k.i32) == 0x80000000 {
					k.id = CTID_UINT32
				}
				lj_ctype_addname(cp.cts, ct, constid)
			}
			if !cp_opt(cp, ',') {
				break
			}
			if !(cp.tok != '}') {
				break
			} // Trailing ',' is ok.
		}
		cp_check(cp, '}')
		// Complete enum.
		ctype_get(cp.cts, eid).info = einfo
		ctype_get(cp.cts, eid).size = esize
	}
	return eid
}

// Parse declaration specifiers.
func cp_decl_spec(cp *CPState, decl *CPDecl, scl CPscl) CPscl {
	var cds = CPscl(0)
	var sz = CTSize(0)
	var tdef = CTypeID(0)

	decl.cp = cp
	decl.mode = cp.mode
	decl.name = ""
	decl.redir = ""
	decl.attr = 0
	decl.fattr = 0
	decl.pos = 0
	decl.top = 0
	decl.stack[0].next = 0

	for { // Parse basic types.
		cp_decl_attributes(cp, decl)
		if cp.tok >= CTOK_FIRSTDECL && cp.tok <= CTOK_LASTDECLFLAG {
			var cbit CPscl
			if cp.ct.size != 0 {
				if sz != 0 {
					goto end_decl
				}
				sz = cp.ct.size
			}
			cbit = 1 << (cp.tok - CTOK_FIRSTDECL)
			cds = cds | cbit | ((cbit & cds & CDF_LONG) << 1)
			if cp.tok >= CTOK_FIRSTSCL {
				if (scl & cbit) == 0 {
					// cp_errmsg(cp, cp.tok, LJ_ERR_FFI_BADSCL);
					panic(cp)
				}
			} else if tdef != 0 {
				goto end_decl
			}
			cp_next(cp)
			continue
		}
		if sz != 0 || tdef != 0 || (cds&(CDF_SHORT|CDF_LONG|CDF_SIGNED|CDF_UNSIGNED|CDF_COMPLEX)) != 0 {
			break
		}
		switch cp.tok {
		case CTOK_STRUCT:
			tdef = cp_decl_struct(cp, decl, CTINFO(CT_STRUCT, 0))
			continue
		case CTOK_UNION:
			tdef = cp_decl_struct(cp, decl, CTINFO(CT_STRUCT, CTF_UNION))
			continue
		case CTOK_ENUM:
			tdef = cp_decl_enum(cp, decl)
			continue
		case CTOK_IDENT:
			if ctype_istypedef(cp.ct.info) {
				tdef = CTypeID(ctype_cid(cp.ct.info)) // Get typedef.
				cp_next(cp)
				continue
			}
			break
		case '$':
			tdef = cp.val.id
			cp_next(cp)
			continue
		default:
			break
		}
		break
	}
end_decl:

	if (cds & CDF_COMPLEX) != 0 { // Use predefined complex types.
		tdef = CTypeID(bt.Choose(sz == 4, CTID_COMPLEX_FLOAT, CTID_COMPLEX_DOUBLE))
	}

	if tdef != 0 {
		cp_push_type(decl, tdef)
	} else if (cds & CDF_VOID) != 0 {
		cp_push(decl, CTINFO(CT_VOID, decl.attr&CTF_QUAL), CTSIZE_INVALID)
		decl.attr &= ^CTF_QUAL
	} else {
		// Determine type info and size.
		var info = CTINFO(CT_NUM, bt.Choose(cds&CDF_UNSIGNED != 0, CTF_UNSIGNED, 0))
		if (cds & CDF_BOOL) != 0 {
			if (cds & ^(CDF_SCL | CDF_BOOL | CDF_INT | CDF_SIGNED | CDF_UNSIGNED)) != 0 {
				//cp_errmsg(cp, 0, LJ_ERR_FFI_INVTYPE)
				panic(cp)
			}
			info |= CTF_BOOL
			if (cds & CDF_SIGNED) == 0 {
				info |= CTF_UNSIGNED
			}
			if sz == 0 {
				sz = 1
			}
		} else if (cds & CDF_FP) != 0 {
			info = CTINFO(CT_NUM, CTF_FP)
			if (cds & CDF_LONG) != 0 {
				sz = 8 // sizeof(long double)
			}
		} else if (cds & CDF_CHAR) != 0 {
			if (cds & (CDF_CHAR | CDF_SIGNED | CDF_UNSIGNED)) == CDF_CHAR {
				info |= CTF_UCHAR // Handle platforms where char is unsigned.
			}
		} else if (cds & CDF_SHORT) != 0 {
			sz = 2 // sizeof(short)
		} else if (cds & CDF_LONGLONG) != 0 {
			sz = 8
		} else if (cds & CDF_LONG) != 0 {
			info |= CTF_LONG
			sz = 4 // sizeof(long)
		} else if sz == 0 {
			if (cds & (CDF_SIGNED | CDF_UNSIGNED)) == 0 {
				//cp_errmsg(cp, cp.tok, LJ_ERR_FFI_DECLSPEC)
				panic(cp)
			}
			sz = CTSize(sizeofInt) // sizeof(int)
		}
		//lj_assertCP(sz != 0, "basic ctype with zero size")
		info += CTInfo(CTALIGN(lj_fls(uint32(sz)))) // Use natural alignment.
		info += decl.attr & CTF_QUAL                // Merge qualifiers.
		cp_push(decl, info, sz)
		decl.attr &= ^CTF_QUAL
	}
	decl.specpos = decl.pos
	decl.specattr = decl.attr
	decl.specfattr = decl.fattr
	return cds & CDF_SCL // Return storage class.
}
