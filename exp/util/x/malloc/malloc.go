package malloc

const ALIGNMENT = 16 // must be 2^N
const ALIGNMENT_SHIFT = 4

type block = uint8

type pool_header struct {
	//union {
	//	block *_padding;
	//	uint count;
	//} ref;                   // number of allocated blocks
	freeblock     *block       // pool's free list head
	nextpool      *pool_header // next pool of this size class
	prevpool      *pool_header // previous pool       ""
	arenaindex    uint         // index into arenas of base adr
	szidx         uint         // block size class index
	nextoffset    uint         // bytes to virgin block
	maxnextoffset uint         // largest valid nextoffset
}

type Allocator struct {
}
