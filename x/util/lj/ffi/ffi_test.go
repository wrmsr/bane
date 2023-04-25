package ffi

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestFfi(t *testing.T) {
	sleep := FindSleep()
	cs := CCallState{
		Fn: sleep,
	}
	cs.Gpr[0] = 3
	fmt.Println(Ffi_call(unsafe.Pointer(&cs)))
}

func TestSqrt(t *testing.T) {
	fmt.Println(Sqrt(9))
}
