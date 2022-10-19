package malloc

import "unsafe"

const ALIGNMENT = 16 // must be 2^N
const ALIGNMENT_SHIFT = 4

func INDEX2SIZE(I uint) uint {
	return (uint(I) + 1) << ALIGNMENT_SHIFT
}

const SMALL_REQUEST_THRESHOLD = 512
const NB_SMALL_SIZE_CLASSES = SMALL_REQUEST_THRESHOLD / ALIGNMENT

type block = uint8

type pool_header struct {
	freeblock     *block       // pool's free list head
	nextpool      *pool_header // next pool of this size class
	prevpool      *pool_header // previous pool
	count         uint
	arenaindex    uint // index into arenas of base adr
	szidx         uint // block size class index
	nextoffset    uint // bytes to virgin block
	maxnextoffset uint // largest valid nextoffset
}

type arena_object struct {
	// The address of the arena, as returned by malloc. Note that 0 will never be returned by a successful malloc, and
	// is used here to mark an arena_object that doesn't correspond to an allocated arena.
	address uintptr

	// Pool-aligned pointer to the next pool to be carved off.
	pool_address *block

	// The number of available pools in the arena:  free pools + never- allocated pools.
	nfreepools uint

	// The total number of pools in the arena, whether or not available.
	ntotalpools uint

	// Singly-linked list of available pools.
	freepools *pool_header

	// Whenever this arena_object is not associated with an allocated arena, the nextarena member is used to link all
	// unassociated arena_objects in the singly-linked `unused_arena_objects` list. The prevarena member is unused in
	// this case.
	//
	// When this arena_object is associated with an allocated arena with at least one available pool, both members are
	// used in the doubly-linked `usable_arenas` list, which is maintained in increasing order of `nfreepools` values.
	//
	// Else this arena_object is associated with an allocated arena all of whose pools are in use.  `nextarena` and
	//  `prevarena` are both meaningless in this case.
	nextarena *arena_object
	prevarena *arena_object
}

func _Py_SIZE_ROUND_DOWN(n, a uintptr) uintptr {
	return n & ^(a - 1)
}

func _Py_SIZE_ROUND_UP(n, a uintptr) uintptr {
	return (n + (a - 1)) & ^(a - 1)
}

func _Py_ALIGN_DOWN(p, a uintptr) uintptr {
	return p & ^(a - 1)
}

func _Py_ALIGN_UP(p, a uintptr) uintptr {
	return (p + (a - 1)) & ^(a - 1)
}

func _Py_IS_ALIGNED(p, a uintptr) bool {
	return (p & (a - 1)) == 0
}

var POOL_OVERHEAD = (func() uint {
	return uint(_Py_SIZE_ROUND_UP(unsafe.Sizeof(pool_header{}), ALIGNMENT))
})()

const DUMMY_SIZE_IDX = 0xffff

const ARENA_BITS = 20
const ARENA_SIZE = 1 << ARENA_BITS
const ARENA_SIZE_MASK = ARENA_SIZE - 1

const POOL_BITS = 14
const POOL_SIZE = 1 << POOL_BITS
const POOL_SIZE_MASK = POOL_SIZE - 1

func POOL_ADDR(P uintptr) uintptr {
	return _Py_ALIGN_DOWN(P, POOL_SIZE)
}

func NUMBLOCKS(I uint) uint {
	return (POOL_SIZE - POOL_OVERHEAD) / INDEX2SIZE(I)
}
