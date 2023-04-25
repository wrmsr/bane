package arm64

import (
	"testing"

	tu "github.com/wrmsr/bane/core/dev/testing"
)

func TestDa(t *testing.T) {
	for _, s := range []string{
		"mvn x1, x2, lsl #47",
	} {
		do_stmt(s)
	}
}

func TestSplitStmt(t *testing.T) {
	for _, tc := range []struct {
		s      string
		op     string
		params []string
	}{
		{
			"mvn x1, x2, lsl #47",
			"mvn",
			[]string{"x1", "x2", "lsl #47"},
		},
	} {
		op, params := splitstmt(tc.s)
		tu.AssertDeepEqual(t, op, tc.op)
		tu.AssertDeepEqual(t, params, tc.params)
	}
}
