package ffi

import (
	"github.com/wrmsr/bane/exp/util/unsafe/dl"
	"github.com/wrmsr/bane/pkg/util/log"
)

func Sqrt(x float32) float32

//go:nosplit
//go:noescape
func Ffi_call(fn uintptr) uintptr

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
