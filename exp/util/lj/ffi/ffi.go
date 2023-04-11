package ffi

import (
	"unsafe"
	_ "unsafe"

	"github.com/wrmsr/bane/exp/util/unsafe/dl"
	log "github.com/wrmsr/bane/pkg/util/slog"
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
	Fn uintptr // Pointer to called function.

	SpAdj uint32 // Stack pointer adjustment.
	Nsp   uint8  // Number of stack slots.

	RetRef bool    // Return value by reference.
	RetP   uintptr // Aggregate return pointer in x8.

	Fpr   [CCALL_NUM_FPR]FPRArg  // Arguments/results in FPRs.
	Gpr   [CCALL_NUM_GPR]GPRArg  // Arguments/results in GPRs.
	Stack [CCALL_MAXSTACK]GPRArg // Stack slots.
}

func Sqrt(x float32) float32

//go:nosplit
//go:noescape
func Ffi_call(cs unsafe.Pointer) uintptr

//go:linkname _runtime_procPin runtime.procPin
func _runtime_procPin() int

func runtime_procPin() int {
	return _runtime_procPin()
}

//go:linkname _runtime_procUnpin runtime.procUnpin
func _runtime_procUnpin()

func runtime_procUnpin() {
	_runtime_procUnpin()
}

func FindSleep() uintptr {
	var lib dl.Library
	var err error
	var ptr uintptr

	if lib, err = dl.Open(dl.Libc, dl.Lazy|dl.Local); err != nil {
		panic(err)
	}

	defer log.OrError(lib.Close)

	if ptr, err = lib.Symbol("sleep"); err != nil {
		panic(err)
	}

	if ptr == 0 {
		panic("not found")
	}

	return ptr
}
