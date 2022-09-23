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
	// cp.val.u32 = (uint32_t) o.i;
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
// #define ctype_isinteger(info) \
// 	(((info) & (CTMASK_NUM|CTF_BOOL|CTF_FP)) == CTINFO(CT_NUM, 0))

func ctype_isinteger_or_bool(info CTInfo) bool {
	return ((info) & (CTMASK_NUM | CTF_FP)) == CTINFO(CT_NUM, 0)
}

// #define ctype_isbool(info) \
// 	(((info) & (CTMASK_NUM|CTF_BOOL)) == CTINFO(CT_NUM, CTF_BOOL))
// #define ctype_isfp(info) \
// 	(((info) & (CTMASK_NUM|CTF_FP)) == CTINFO(CT_NUM, CTF_FP))

// #define ctype_ispointer(info) // Pointer or array.
// 	((ctype_type(info) >> 1) == (CT_PTR >> 1))  #define ctype_isref(info)  (((info) & (CTMASK_NUM|CTF_REF)) == CTINFO(CT_PTR, CTF_REF))

// #define ctype_isrefarray(info) \
// 	(((info) & (CTMASK_NUM|CTF_VECTOR|CTF_COMPLEX)) == CTINFO(CT_ARRAY, 0))
// #define ctype_isvector(info) \
// 	(((info) & (CTMASK_NUM|CTF_VECTOR)) == CTINFO(CT_ARRAY, CTF_VECTOR))
// #define ctype_iscomplex(info) \
// 	(((info) & (CTMASK_NUM|CTF_COMPLEX)) == CTINFO(CT_ARRAY, CTF_COMPLEX))

// #define ctype_isvltype(info) \ // VL array or VL struct.
// 	(((info) & ((CTMASK_NUM|CTF_VLA) - (2u<<CTSHIFT_NUM))) == CTINFO(CT_STRUCT, CTF_VLA))
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
	cts.hash[ct.name] = id
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
