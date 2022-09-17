package c

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

type CPChar int  // C parser character. Unsigned ext. from char.
type CPToken int // C parser token.

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
	c   CPChar  // Current character.
	tok CPToken // Current token.
	val CPValue // Token value.
	str string  // Interned string of identifier/keyword.
	ct  *CType  // C type table entry.
	p   int     // Current position in input buffer.
	sb  []rune  // String buffer for tokens.
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
