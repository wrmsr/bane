package main

import "github.com/wrmsr/bane/exp/util/lj/ffi"

func main() {
	ffi.Ffi_call(ffi.FindSleep())
}
