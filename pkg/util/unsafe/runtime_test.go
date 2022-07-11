package unsafe

import (
	"bytes"
	"hash/fnv"
	"testing"

	"github.com/cespare/xxhash"
	"github.com/segmentio/asm/base64"
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

func TestDecodeLines(t *testing.T) {
	src := []byte(`dGVzdCB0ZXN0IHRlc3QgdGVzdCB0ZXN0IHRlc3QgdGVzdCB0ZXN0IHRlc3QgdGVzdCB0ZXN0IHRl
c3QgdGVzdCB0ZXN0IHRlc3QgdGVzdCB0ZXN0IHRlc3QgdGVzdCB0ZXN0IHRlc3QgdGVzdCB0ZXN0
IHRlc3QgdGVzdCB0ZXN0IHRlc3QgdGVzdCB0ZXN0IHRlc3QgdGVzdCB0ZXN0IHRlc3QgdGVzdCB0
ZXN0IHRlc3QgdGVzdCB0ZXN0IHRlc3QgdGVzdCB0ZXN0IHRlc3QgdGVzdCB0ZXN0IHRlc3QgdGVz
dCB0ZXN0IHRlc3QgdGVzdA==`)

	expect := []byte(`test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test`)
	actual := make([]byte, base64.StdEncoding.DecodedLen(len(src)))
	n, err := base64.StdEncoding.Decode(actual, src)

	if err != nil {
		t.Fatalf("decode error: %v", err)
	}

	if !bytes.Equal(expect, actual[:n]) {
		t.Errorf("failed decode:\n\texpect = %v\n\tactual = %v", expect, actual)
	}
}
