package unsafe

import (
	"hash/fnv"
	"testing"

	"github.com/cespare/xxhash"
)

func strhash0(s string) uint64 {
	h := fnv.New64()
	_, _ = h.Write([]byte(s))
	return h.Sum64()
}

func strhash1(s string) uint64 {
	return Strhash(s, 0)
}

func strhash2(s string) uint64 {
	return xxhash.Sum64String(s)
}

func benchmarkStrhash(b *testing.B, fn func(string) uint64) {
	for i := 0; i < b.N; i++ {
		fn("abcdabcabcabcabcabcabcabcddddddd")
	}
}

func BenchmarkStrhash0(b *testing.B) { benchmarkStrhash(b, strhash0) }
func BenchmarkStrhash1(b *testing.B) { benchmarkStrhash(b, strhash1) }
func BenchmarkStrhash2(b *testing.B) { benchmarkStrhash(b, strhash2) }
