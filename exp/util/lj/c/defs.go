package c

// C lexer keywords.
type _CTOKDEF struct {
	dname string
	name  string
}

var _CTOKDEFS = []_CTOKDEF{
	{"IDENT", "<identifier>"},
	{"STRING", "<string>"},
	{"INTEGER", "<integer>"},
	{"EOF", "<eof>"},
	{"OROR", "||"},
	{"ANDAND", "&&"},
	{"EQ", "=="},
	{"NE", "!="},
	{"LE", "<="},
	{"GE", ">="},
	{"SHL", "<<"},
	{"SHR", ">>"},
	{"DEREF", "."},
}

// Simple declaration specifiers.
type _CDSDEF struct {
	dname string
}

var _CDSDEFS = []_CDSDEF{
	{"VOID"},
	{"BOOL"},
	{"CHAR"},
	{"INT"},
	{"FP"},
	{"LONG"},
	{"LONGLONG"},
	{"SHORT"},
	{"COMPLEX"},
	{"SIGNED"},
	{"UNSIGNED"},
	{"CONST"},
	{"VOLATILE"},
	{"RESTRICT"},
	{"INLINE"},
	{"TYPEDEF"},
	{"EXTERN"},
	{"STATIC"},
	{"AUTO"},
	{"REGISTER"},
}

// C keywords.
type _CKWDEF struct {
	dname string
}

var _CKWDEFS []_CKWDEF

func init() {
	for _, d := range _CDSDEFS {
		_CKWDEFS = append(_CKWDEFS, _CKWDEF{
			dname: d.dname,
		})
	}

	_CKWDEFS = append(_CKWDEFS, []_CKWDEF{
		{"EXTENSION"},
		{"ASM"},
		{"ATTRIBUTE"},
		{"DECLSPEC"},
		{"CCDECL"},
		{"PTRSZ"},
		{"STRUCT"},
		{"UNION"},
		{"ENUM"},
		{"SIZEOF"},
		{"ALIGNOF"},
	}...)
}

// Common types.
type _CTTYDEF struct {
	dname string
	sz    int
	ct    CTInfo
	info  CTInfo
}

var _CTTYDEFS = []_CTTYDEF{
	{"NONE", 0, CT_ATTRIB, CTATTRIB(CTA_BAD)},
	{"VOID", -1, CT_VOID, CTALIGNi(0)},
	{"CVOID", -1, CT_VOID, CTF_CONST | CTALIGNi(0)},
	{"BOOL", 1, CT_NUM, CTF_BOOL | CTF_UNSIGNED | CTALIGNi(0)},
	{"CCHAR", 1, CT_NUM, CTF_CONST | CTF_UCHAR | CTALIGNi(0)},
	{"INT8", 1, CT_NUM, CTALIGNi(0)},
	{"UINT8", 1, CT_NUM, CTF_UNSIGNED | CTALIGNi(0)},
	{"INT16", 2, CT_NUM, CTALIGNi(1)},
	{"UINT16", 2, CT_NUM, CTF_UNSIGNED | CTALIGNi(1)},
	{"INT32", 4, CT_NUM, CTALIGNi(2)},
	{"UINT32", 4, CT_NUM, CTF_UNSIGNED | CTALIGNi(2)},
	{"INT64", 8, CT_NUM, CTF_LONG | CTALIGNi(3)},
	{"UINT64", 8, CT_NUM, CTF_UNSIGNED | CTF_LONG | CTALIGNi(3)},
	{"FLOAT", 4, CT_NUM, CTF_FP | CTALIGNi(2)},
	{"DOUBLE", 8, CT_NUM, CTF_FP | CTALIGNi(3)},
	{"COMPLEX_FLOAT", 8, CT_ARRAY, CTF_COMPLEX | CTALIGNi(2) | CTID_FLOAT},
	{"COMPLEX_DOUBLE", 16, CT_ARRAY, CTF_COMPLEX | CTALIGNi(3) | CTID_DOUBLE},
	{"P_VOID", CTSIZE_PTR, CT_PTR, CTALIGN_PTR | CTID_VOID},
	{"P_CVOID", CTSIZE_PTR, CT_PTR, CTALIGN_PTR | CTID_CVOID},
	{"P_CCHAR", CTSIZE_PTR, CT_PTR, CTALIGN_PTR | CTID_CCHAR},
	{"P_UINT8", CTSIZE_PTR, CT_PTR, CTALIGN_PTR | CTID_UINT8},
	{"A_CCHAR", -1, CT_ARRAY, CTF_CONST | CTALIGNi(0) | CTID_CCHAR},
	{"CTYPEID", 4, CT_ENUM, CTALIGNi(2) | CTID_INT32},
}

// Predefined typedefs.
type _CTTDDEF struct {
	name   string
	idname string
	id     CTypeID
}

var _CTTDDEFS = []_CTTDDEF{
	// Vararg handling.
	{"va_list", "P_VOID", CTID_P_VOID},
	{"__builtin_va_list", "P_VOID", CTID_P_VOID},
	{"__gnuc_va_list", "P_VOID", CTID_P_VOID},
	// From stddef.h.
	{"ptrdiff_t", "INT_PSZ", CTID_INT_PSZ},
	{"size_t", "UINT_PSZ", CTID_UINT_PSZ},
	{"wchar_t", "WCHAR", CTID_WCHAR},
	// Subset of stdint.h.
	{"int8_t", "INT8", CTID_INT8},
	{"int16_t", "INT16", CTID_INT16},
	{"int32_t", "INT32", CTID_INT32},
	{"int64_t", "INT64", CTID_INT64},
	{"uint8_t", "UINT8", CTID_UINT8},
	{"uint16_t", "UINT16", CTID_UINT16},
	{"uint32_t", "UINT32", CTID_UINT32},
	{"uint64_t", "UINT64", CTID_UINT64},
	{"intptr_t", "INT_PSZ", CTID_INT_PSZ},
	{"uintptr_t", "UINT_PSZ", CTID_UINT_PSZ},
	// From POSIX.
	{"ssize_t", "INT_PSZ", CTID_INT_PSZ},
}

// Keywords (only the ones we actually care for).
type _CTKWDEF struct {
	name    string
	sz      int
	tokname string
	kw      int
}

var _CTKWDEFS = []_CTKWDEF{
	// Type specifiers.
	{"void", -1, "VOID", CTOK_VOID},
	{"_Bool", 0, "BOOL", CTOK_BOOL},
	{"bool", 1, "BOOL", CTOK_BOOL},
	{"char", 1, "CHAR", CTOK_CHAR},
	{"int", 4, "INT", CTOK_INT},
	{"__int8", 1, "INT", CTOK_INT},
	{"__int16", 2, "INT", CTOK_INT},
	{"__int32", 4, "INT", CTOK_INT},
	{"__int64", 8, "INT", CTOK_INT},
	{"float", 4, "FP", CTOK_FP},
	{"double", 8, "FP", CTOK_FP},
	{"long", 0, "LONG", CTOK_LONG},
	{"short", 0, "SHORT", CTOK_SHORT},
	{"_Complex", 0, "COMPLEX", CTOK_COMPLEX},
	{"complex", 0, "COMPLEX", CTOK_COMPLEX},
	{"__complex", 0, "COMPLEX", CTOK_COMPLEX},
	{"__complex__", 0, "COMPLEX", CTOK_COMPLEX},
	{"signed", 0, "SIGNED", CTOK_SIGNED},
	{"__signed", 0, "SIGNED", CTOK_SIGNED},
	{"__signed__", 0, "SIGNED", CTOK_SIGNED},
	{"unsigned", 0, "UNSIGNED", CTOK_UNSIGNED},
	// Type qualifiers.
	{"const", 0, "CONST", CTOK_CONST},
	{"__const", 0, "CONST", CTOK_CONST},
	{"__const__", 0, "CONST", CTOK_CONST},
	{"volatile", 0, "VOLATILE", CTOK_VOLATILE},
	{"__volatile", 0, "VOLATILE", CTOK_VOLATILE},
	{"__volatile__", 0, "VOLATILE", CTOK_VOLATILE},
	{"restrict", 0, "RESTRICT", CTOK_RESTRICT},
	{"__restrict", 0, "RESTRICT", CTOK_RESTRICT},
	{"__restrict__", 0, "RESTRICT", CTOK_RESTRICT},
	{"inline", 0, "INLINE", CTOK_INLINE},
	{"__inline", 0, "INLINE", CTOK_INLINE},
	{"__inline__", 0, "INLINE", CTOK_INLINE},
	// Storage class specifiers.
	{"typedef", 0, "TYPEDEF", CTOK_TYPEDEF},
	{"extern", 0, "EXTERN", CTOK_EXTERN},
	{"static", 0, "STATIC", CTOK_STATIC},
	{"auto", 0, "AUTO", CTOK_AUTO},
	{"register", 0, "REGISTER", CTOK_REGISTER},
	// GCC Attributes.
	{"__extension__", 0, "EXTENSION", CTOK_EXTENSION},
	{"__attribute", 0, "ATTRIBUTE", CTOK_ATTRIBUTE},
	{"__attribute__", 0, "ATTRIBUTE", CTOK_ATTRIBUTE},
	{"asm", 0, "ASM", CTOK_ASM},
	{"__asm", 0, "ASM", CTOK_ASM},
	{"__asm__", 0, "ASM", CTOK_ASM},
	// MSVC Attributes.
	{"__declspec", 0, "DECLSPEC", CTOK_DECLSPEC},
	{"__cdecl", CTCC_CDECL, "CCDECL", CTOK_CCDECL},
	{"__thiscall", CTCC_THISCALL, "CCDECL", CTOK_CCDECL},
	{"__fastcall", CTCC_FASTCALL, "CCDECL", CTOK_CCDECL},
	{"__stdcall", CTCC_STDCALL, "CCDECL", CTOK_CCDECL},
	{"__ptr32", 4, "PTRSZ", CTOK_PTRSZ},
	{"__ptr64", 8, "PTRSZ", CTOK_PTRSZ},
	// Other type specifiers.
	{"struct", 0, "STRUCT", CTOK_STRUCT},
	{"union", 0, "UNION", CTOK_UNION},
	{"enum", 0, "ENUM", CTOK_ENUM},
	// Operators.
	{"sizeof", 0, "SIZEOF", CTOK_SIZEOF},
	{"__alignof", 0, "ALIGNOF", CTOK_ALIGNOF},
	{"__alignof__", 0, "ALIGNOF", CTOK_ALIGNOF},
}
