package arm64

type Action int8

const (
	DASM_STOP Action = iota
	DASM_SECTION
	DASM_ESC
	DASM_REL_EXT
	// The following actions need a buffer position.
	DASM_ALIGN
	DASM_REL_LG
	DASM_LABEL_LG
	// The following actions also have an argument.
	DASM_REL_PC
	DASM_LABEL_PC
	DASM_REL_A
	DASM_IMM
	DASM_IMM6
	DASM_IMM12
	DASM_IMM13W
	DASM_IMM13X
	DASM_IMML
	DASM_IMMV
	DASM_VREG
	DASM__MAX
)

// Maximum number of section buffer positions for a single dasm_put() call.
const DASM_MAXSECPOS = 25

// DynASM encoder status codes. Action list offset or number are or'ed in.
const DASM_S_OK = 0x00000000
const DASM_S_NOMEM = 0x01000000
const DASM_S_PHASE = 0x02000000
const DASM_S_MATCH_SEC = 0x03000000
const DASM_S_RANGE_I = 0x11000000
const DASM_S_RANGE_SEC = 0x12000000
const DASM_S_RANGE_LG = 0x13000000
const DASM_S_RANGE_PC = 0x14000000
const DASM_S_RANGE_REL = 0x15000000
const DASM_S_RANGE_VREG = 0x16000000
const DASM_S_UNDEF_LG = 0x21000000
const DASM_S_UNDEF_PC = 0x22000000

// Macros to convert positions (8 bit section + 24 bit index).
func DASM_POS2IDX(pos int) int  { return pos & 0x00ffffff }
func DASM_POS2BIAS(pos int) int { return pos & 0xff000000 }
func DASM_SEC2POS(sec int) int  { return sec << 24 }
func DASM_POS2SEC(pos int) int  { return pos >> 24 }

//func DASM_POS2PTR(D, pos int) int { return (D- > sections[DASM_POS2SEC(pos)].rbuf+(pos)) }

type dasm_ActList = []int

// Per-section structure.
type dasm_Section struct {
	rbuf  int // Biased buffer pointer (negative section bias).
	buf   int // True buffer pointer.
	bsize int // Buffer size in bytes.
	pos   int // Biased buffer position.
	epos  int // End of biased buffer position - max single put.
	ofs   int // Byte offset into section.
}

// Core structure holding the DynASM encoding state.
type dasm_State struct {
	psize      int          // Allocated size of this structure.
	actionlist dasm_ActList // Current actionlist pointer.
	lglabels   int          // Local/global chain/pos ptrs.
	lgsize     int
	pclabels   int // PC label chains/pos ptrs.
	pcsize     int
	globals    int            // Array of globals (bias -10).
	section    dasm_Section   // Pointer to active section.
	codesize   int            // Total size of all code sections.
	maxsection int            // 0 <= sectionidx < maxsection.
	status     int            // Status code.
	sections   []dasm_Section // All sections. Alloc-extended.
}
