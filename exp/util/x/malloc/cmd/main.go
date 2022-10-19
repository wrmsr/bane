package main

import (
	"fmt"
	"unsafe"
)

func pool_header__arenaindex(p uintptr) uint {
	return *(*uint)(unsafe.Pointer(p + 24))
}

func main() {
	b := new([80]byte)
	for i := range b {
		b[i] = byte(i)
	}
	p := uintptr(unsafe.Pointer(&b[0]))
	fmt.Println(pool_header__arenaindex(p))
	fmt.Println(b)
}
