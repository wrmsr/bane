/*
https://en.wikipedia.org/wiki/LEB128
*/
package leb128

import (
	"fmt"
	"strings"
)

const (
	Size32 = 5
	Size64 = 10
)

type Leb128 struct {
	L uint8
	B [Size64]byte
}

func (l Leb128) String() string {
	if l.L == 0 {
		return "00"
	}
	var sb strings.Builder
	for i := uint8(0); i < l.L; i++ {
		if i > 0 {
			sb.WriteRune('_')
		}
		_, _ = fmt.Fprintf(&sb, "%02x", l.B[i])
	}
	return sb.String()
}

type Int128 struct {
	Hi int64
	Lo int64
}

func EncodeU32(v uint32) (l Leb128) {
	for {
		c := byte(v & 0x7F)
		v >>= 7
		if v == 0 {
			l.B[l.L] = c
			l.L++
			return
		}
		l.B[l.L] = c | 0x80
		l.L++
	}
}

func EncodeI32(v int32) (l Leb128) {
	if v < 0 {
		for {
			c := byte(v & 0x7F)
			v >>= 7
			if v == -1 && (c&0x40) != 0 {
				l.B[l.L] = c
				l.L++
				return
			}
			l.B[l.L] = c | 0x80
			l.L++
		}
	} else {
		for {
			c := byte(v & 0x7F)
			v >>= 7
			if v == 0 && (c&0x40) == 0 {
				l.B[l.L] = c
				l.L++
				return
			}
			l.B[l.L] = c | 0x80
			l.L++
		}
	}
}
