package json

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

func TestDecodeObject(t *testing.T) {
	for _, s := range []string{
		` {"a": "foo", "b": 123, "c": [{"d": "x"}, null]} `,
		`[["a", "b"], ["c", {"d": "e"}]]`,
		`null`,
	} {
		dec := json.NewDecoder(strings.NewReader(s))
		o := check.Must1(DecodeRawObject(dec))
		fmt.Println(o)
	}
}
