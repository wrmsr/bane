package types

import (
	"encoding/binary"
	"hash"
	"hash/fnv"
)

func DefaultHash() hash.Hash64 {
	return fnv.New64a()
}

func HashBytes(s []byte) uintptr {
	h := DefaultHash()
	_, _ = h.Write(s)
	return uintptr(h.Sum64())
}

func HashStr(s string) uintptr {
	h := DefaultHash()
	var ub [4]byte
	u := ub[:]
	for _, b := range s {
		binary.LittleEndian.PutUint32(u, uint32(b))
		_, _ = h.Write(u)
	}
	return uintptr(h.Sum64())
}
