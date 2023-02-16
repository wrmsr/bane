package consteval

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	ju "github.com/wrmsr/bane/pkg/util/json"
	msh "github.com/wrmsr/bane/pkg/util/marshal"
)

func TestEval(t *testing.T) {
	s := &Scope{
		m: map[string]any{
			"foo": Basic{K: IntBasic, S: "420"},
			"bar": Ident{N: "foo"},
		},
	}

	ov := check.Must1(s.Reduce("foo"))
	fmt.Println(check.Must1(ju.MarshalPretty(check.Must1(msh.Marshal(&ov)))))
}
