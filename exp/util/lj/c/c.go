package c

// C type numbers. Highest 4 bits of C type info. ORDER CT.
const (
	// Externally visible types.
	CT_NUM     = iota // Integer or floating-point numbers.
	CT_STRUCT         // Struct or union.
	CT_PTR            // Pointer or reference.
	CT_ARRAY          // Array or complex type.
	CT_VOID           // Void type.
	CT_ENUM           // Enumeration.
	CT_FUNC           // Function.
	CT_TYPEDEF        // Typedef.
	CT_ATTRIB         // Miscellaneous attributes.
	// Internal element types.
	CT_FIELD    // Struct/union field or function parameter.
	CT_BITFIELD // Struct/union bitfield.
	CT_CONSTVAL // Constant value.
	CT_EXTERN   // External reference.
	CT_KW       // Keyword.

	CT_MAYCONVERT = CT_ARRAY
	CT_HASSIZE    = CT_ENUM // Last type where ct->size holds the actual size.
)

/*
 ---------- info ------------
|type      flags...  A   cid | size   |  sib  | next  | name  |
+----------------------------+--------+-------+-------+-------+--
|NUM       BFcvUL..  A       | size   |       | type  |       |
|STRUCT    ..cvU..V  A       | size   | field | name? | name? |
|PTR       ..cvR...  A   cid | size   |       | type  |       |
|ARRAY     VCcv...V  A   cid | size   |       | type  |       |
|VOID      ..cv....  A       | size   |       | type  |       |
|ENUM                A   cid | size   | const | name? | name? |
|FUNC      ....VS.. cc   cid | nargs  | field | name? | name? |
|TYPEDEF                 cid |        |       | name  | name  |
|ATTRIB        attrnum   cid | attr   | sib?  | type? |       |
|FIELD                   cid | offset | field |       | name? |
|BITFIELD  B.cvU csz bsz pos | offset | field |       | name? |
|CONSTVAL    c           cid | value  | const | name  | name  |
|EXTERN                  cid |        | sib?  | name  | name  |
|KW                      tok | size   |       | name  | name  |
+----------------------------+--------+-------+-------+-------+--
       ^^  ^^--- bits used for C type conversion dispatch
*/

// C type info flags.     TFFArrrr
const (
	CTF_BOOL       CTInfo = 0x08000000 // Boolean: NUM, BITFIELD.
	CTF_FP         CTInfo = 0x04000000 // Floating-point: NUM.
	CTF_CONST      CTInfo = 0x02000000 // Const qualifier.
	CTF_VOLATILE   CTInfo = 0x01000000 // Volatile qualifier.
	CTF_UNSIGNED   CTInfo = 0x00800000 // Unsigned: NUM, BITFIELD.
	CTF_LONG       CTInfo = 0x00400000 // Long: NUM.
	CTF_VLA        CTInfo = 0x00100000 // Variable-length: ARRAY, STRUCT.
	CTF_REF        CTInfo = 0x00800000 // Reference: PTR.
	CTF_VECTOR     CTInfo = 0x08000000 // Vector: ARRAY.
	CTF_COMPLEX    CTInfo = 0x04000000 // Complex: ARRAY.
	CTF_UNION      CTInfo = 0x00800000 // Union: STRUCT.
	CTF_VARARG     CTInfo = 0x00800000 // Vararg: FUNC.
	CTF_SSEREGPARM CTInfo = 0x00400000 // SSE register parameters: FUNC.

	CTF_QUAL  = CTF_CONST | CTF_VOLATILE
	CTF_ALIGN = CTMASK_ALIGN << CTSHIFT_ALIGN
)

const CTF_UCHAR = 0 // ((char)-1 > 0 ? CTF_UNSIGNED: 0)

// Flags used in parser.  .F.Ammvf   cp->attr
const CTFP_ALIGNED = 0x00000001 // cp->attr + ALIGN
const CTFP_PACKED = 0x00000002  // cp->attr
// ...C...f   cp->fattr
const CTFP_CCONV = 0x00000001 // cp->fattr + CCONV/[SSE]REGPARM

// C type info bitfields.
const CTMASK_CID CTInfo = 0x0000ffff // Max. 65536 type IDs.
const CTMASK_NUM = 0xf0000000        // Max. 16 type numbers.
const CTSHIFT_NUM = 28
const CTMASK_ALIGN = 15 // Max. alignment is 2^15.
const CTSHIFT_ALIGN = 16
const CTMASK_ATTRIB = 255 // Max. 256 attributes.
const CTSHIFT_ATTRIB = 16
const CTMASK_CCONV = 3 // Max. 4 calling conventions.
const CTSHIFT_CCONV = 16
const CTMASK_REGPARM = 3 // Max. 0-3 regparms.
const CTSHIFT_REGPARM = 18

// Bitfields only used in parser.
const CTMASK_VSIZEP = 15 // Max. vector size is 2^15.
const CTSHIFT_VSIZEP = 4
const CTMASK_MSIZEP = 255 // Max. type size (via mode) is 128.
const CTSHIFT_MSIZEP = 8

// Info bits for BITFIELD. Max. size of bitfield is 64 bits.
const CTBSZ_MAX = 32    // Max. size of bitfield is 32 bit.
const CTBSZ_FIELD = 127 // Temp. marker for regular field.
const CTMASK_BITPOS = 127
const CTMASK_BITBSZ = 127
const CTMASK_BITCSZ = 127
const CTSHIFT_BITPOS = 0
const CTSHIFT_BITBSZ = 8
const CTSHIFT_BITCSZ = 16

//#define CTF_INSERT(info, field, val) \
//	info = (info & ~(CTMASK_##field<<CTSHIFT_##field)) | \
//	(((CTSize)(val) & CTMASK_##field) << CTSHIFT_##field)

// Calling conventions. ORDER CC
const (
	CTCC_CDECL = iota
	CTCC_THISCALL
	CTCC_FASTCAL
	CTCC_STDCALL
)

// Attribute numbers.
const (
	CTA_NONE    = iota // Ignored attribute. Must be zero.
	CTA_QUAL           // Unmerged qualifiers.
	CTA_ALIGN          // Alignment override.
	CTA_SUBTYPE        // Transparent sub-type.
	CTA_REDIR          // Redirected symbol name.
	CTA_BAD            // To catch bad IDs.
	CTA__MAX
)

// Special sizes.
const CTSIZE_INVALID = 0xffffffff

type CTInfo uint32     // Type info.
type CTSize uint32     // Type size.
type CTypeID uint32    // Type ID.
type CTypeID1 = uint16 // Minimum-sized type ID.

// C type table element.
type CType struct {
	info CTInfo   // Type info.
	size CTSize   // Type size or other info.
	sib  CTypeID1 // Sibling element.
	next CTypeID1 // Next element in hash chain.
	name string   // Element name.
}

const CTHASH_SIZE = 128 // Number of hash anchors.
const CTHASH_MASK = CTHASH_SIZE - 1

// C type state.
type CTState struct {
	tab []*CType // C type table.
	top CTypeID  // Current top of C type table.
	//sizetab MSize   // Size of C type table.
	// lua_State *L;        // Lua state (needed for errors and allocations).
	// global_State *g;    // Global state.
	// GCtab *finalizer;    // Map of cdata to finalizer.
	// GCtab *miscmap;    // Map of -CTypeID to metatable and cb slot to func.
	// CCallback cb;        // Temporary callback state.
	//hash [CTHASH_SIZE]CTypeID1 // Hash anchors for C type table.
	hash map[string]CTypeID
}
