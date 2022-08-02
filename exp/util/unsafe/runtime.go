package unsafe

import (
	"reflect"
	"unsafe"
)

//go:noescape
//go:linkname memhash runtime.memhash
func memhash(p unsafe.Pointer, h, s uintptr) uintptr

//go:noescape
//go:linkname strhash runtime.strhash
func strhash(p unsafe.Pointer, h uintptr) uintptr

func Memhash(b []byte, seed uint64) uint64 {
	pstring := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	return uint64(memhash(unsafe.Pointer(pstring.Data), uintptr(seed), uintptr(pstring.Len)))
}

func Strhash(str string, seed uint64) uint64 {
	return uint64(strhash(unsafe.Pointer(&str), uintptr(seed)))
}
