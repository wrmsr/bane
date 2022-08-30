package codecs

import (
	"fmt"
	"testing"

	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
)

func TestGzip(t *testing.T) {
	s := "foo"
	b := GzipCodec{}.Encode([]byte(s))
	fmt.Println(b)
	s2 := string(GzipCodec{}.Decode(b))
	fmt.Println(s2)
	tu.AssertEqual(t, s, s2)
}
