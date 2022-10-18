package ffi

import (
	"fmt"
	"testing"
	"unsafe"
)

const CCALL_NARG_GPR = 8
const CCALL_NRET_GPR = 2
const CCALL_NARG_FPR = 8
const CCALL_NRET_FPR = 4
const CCALL_SPS_FREE = 0

// #define CCALL_NUM_GPR (CCALL_NARG_GPR > CCALL_NRET_GPR ? CCALL_NARG_GPR : CCALL_NRET_GPR)
// #define CCALL_NUM_FPR (CCALL_NARG_FPR > CCALL_NRET_FPR ? CCALL_NARG_FPR : CCALL_NRET_FPR)

const CCALL_NUM_GPR = CCALL_NARG_GPR
const CCALL_NUM_FPR = CCALL_NARG_FPR

const CCALL_MAXSTACK = 32

type FPRArg uint64
type GPRArg uint64

type CCallState struct {
	fn uintptr // Pointer to called function.

	spadj uint32 // Stack pointer adjustment.
	nsp   uint8  // Number of stack slots.

	retref bool    // Return value by reference.
	retp   uintptr // Aggregate return pointer in x8.

	fpr   [CCALL_NUM_FPR]FPRArg  // Arguments/results in FPRs.
	gpr   [CCALL_NUM_GPR]GPRArg  // Arguments/results in GPRs.
	stack [CCALL_MAXSTACK]GPRArg // Stack slots.
}

func TestFfi(t *testing.T) {
	sleep := FindSleep()
	cs := CCallState{
		fn: sleep,
	}
	fmt.Println(Ffi_call(unsafe.Pointer(&cs)))
}

func TestSqrt(t *testing.T) {
	fmt.Println(Sqrt(9))
}
