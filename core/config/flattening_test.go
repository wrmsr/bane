package config

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/core/check"
	tu "github.com/wrmsr/bane/core/dev/testing"
	ju "github.com/wrmsr/bane/core/json"
)

type m = map[string]any
type s = []any

func TestFlattening(t *testing.T) {
	o := m{
		"a": 1,
		"b": m{
			"c": 2,
		},
		"d": s{
			"e",
			m{
				"f": 3,
			},
		},
		"g": s{
			s{
				"a",
				"b",
			},
			s{
				"c",
				"d",
			},
		},
	}
	fmt.Println(check.Must1(ju.MarshalIndentString(o, "", "  ")))

	f := NewFlattening(DefaultFlatteningConfig())

	of := check.Must1(f.Flatten(o))
	fmt.Println(check.Must1(ju.MarshalIndentString(of, "", "  ")))

	uf := check.Must1(f.Unflatten(of))
	fmt.Println(check.Must1(ju.MarshalIndentString(uf, "", "  ")))

	tu.AssertDeepEqual(t, o, uf)
}
