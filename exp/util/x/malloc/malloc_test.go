package malloc

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestMalloc(t *testing.T) {
	b := new([80]byte)
	for i := range b {
		b[i] = byte(i)
	}
	p := uintptr(unsafe.Pointer(&b[0]))
	fmt.Println(pool_header__arenaindex(p))
	fmt.Println(b)
}
