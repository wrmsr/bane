package leb128

import (
	"fmt"
	"testing"
)

func TestLeb128(t *testing.T) {
	l0 := EncodeU64(624485)
	l1 := EncodeI64(-123456)
	fmt.Println(l0)
	fmt.Println(l1)
	fmt.Println(DecodeU64(ReadSlice(l0.Slice())))
	fmt.Println(DecodeI64(ReadSlice(l1.Slice())))
}
