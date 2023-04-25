package main

import (
	"fmt"
	"unsafe"

	"github.com/wrmsr/bane/x/lj/ffi"
)

func main() {
	sleep := ffi.FindSleep()
	cs := ffi.CCallState{
		Fn: sleep,
	}
	cs.Gpr[0] = 3
	fmt.Println(ffi.Ffi_call(unsafe.Pointer(&cs)))
}
