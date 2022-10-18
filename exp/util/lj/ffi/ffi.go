package ffi

import (
	"unsafe"
	_ "unsafe"

	"github.com/wrmsr/bane/exp/util/unsafe/dl"
	"github.com/wrmsr/bane/pkg/util/log"
)

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
