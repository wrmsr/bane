package unsafe

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	tu "github.com/wrmsr/bane/core/dev/testing"
)

func TestPtrs(t *testing.T) {
	p := new(int)
	a0 := reflect.ValueOf(p).Pointer()
	a1 := uintptr(unsafe.Pointer(p))
	fmt.Printf("%x == %x\n", a0, a1)
	tu.AssertEqual(t, a0, a1)
}

//go:noinline
func ptr0(p *int) uintptr {
	return reflect.ValueOf(p).Pointer()
}

//go:noinline
func ptr1(p *int) uintptr {
	return uintptr(unsafe.Pointer(p))
}

func benchmarkPtr(b *testing.B, fn func(p *int) uintptr) {
	p := new(int)
	for n := 0; n < b.N; n++ {
		fn(p)
	}
}

func BenchmarkPtr0(b *testing.B) { benchmarkPtr(b, ptr0) }
func BenchmarkPtr1(b *testing.B) { benchmarkPtr(b, ptr1) }
