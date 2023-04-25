package types

import (
	"testing"

	tu "github.com/wrmsr/bane/core/dev/testing"
)

func TestBytes(t *testing.T) {
	bs := []Bytes{
		BytesOf([]byte{0}),
		BytesOf([]byte{0, 1}),
		BytesOf([]byte{1}),
		BytesOf([]byte{1, 1}),
		BytesOf([]byte{1, 1, 1}),
		BytesOf([]byte{2, 1, 1}),
	}

	for i := 0; i < len(bs)-1; i++ {
		tu.AssertEqual(t, bs[i].Compare(bs[i+1]), CmpLesser)
		tu.AssertEqual(t, bs[i].Compare(bs[i]), CmpEqual)
		tu.AssertEqual(t, bs[i+1].Compare(bs[i]), CmpGreater)
	}
}
