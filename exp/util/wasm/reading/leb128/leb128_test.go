package leb128

import (
	"fmt"
	"testing"
)

func TestLeb128(t *testing.T) {
	fmt.Println(EncodeU32(uint32(624485)))
	fmt.Println(EncodeI32(int32(-123456)))
}
