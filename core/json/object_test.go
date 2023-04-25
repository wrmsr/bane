package json

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/wrmsr/bane/core/check"
	tu "github.com/wrmsr/bane/core/dev/testing"
	bt "github.com/wrmsr/bane/core/types"
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

func TestRawObject(t *testing.T) {
	kvs := []bt.Kv[int, string]{
		bt.KvOf(1, "one"),
		bt.KvOf(2, "two"),
		bt.KvOf(3, "three"),
		bt.KvOf(1, "one again"),
	}

	o := check.Must1(MakeRawObject(kvs))
	fmt.Println(o)

	kvs2 := check.Must1(UnmarshalRawObjectFields[int, string](o))
	fmt.Println(kvs2)

	tu.AssertDeepEqual(t, kvs, kvs2)
}
