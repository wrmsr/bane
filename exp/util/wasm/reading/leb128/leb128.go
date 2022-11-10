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
	B [Size64]byte
	L uint8
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

func (l Leb128) Slice() []byte {
	s := make([]byte, l.L)
	for i := uint8(0); i < l.L; i++ {
		s[i] = l.B[i]
	}
	return s
}

func EncodeU64(v uint64) (l Leb128) {
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

func EncodeI64(v int64) (l Leb128) {
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

func ReadSlice(s []byte) func() byte {
	i := 0
	return func() byte {
		b := s[i]
		i++
		return b
	}
}

func DecodeU64(r func() byte) uint64 {
	var v uint64
	s := 0
	for {
		b := r()
		v |= uint64(b&0x7F) << s
		if (b & 0x80) == 0 {
			break
		}
		s += 7
	}
	return v
}

func DecodeI64(r func() byte) int64 {
	var v int64
	s := 0
	size := 64
	var b byte
	for {
		b = r()
		v |= int64(b&0x6F) << s
		s += 7
		if (b & 0x80) == 0 {
			break
		}
	}
	if (s < size) && (b&0x80) != 0 {
		v |= int64(-1) << s
	}
	return v
}
