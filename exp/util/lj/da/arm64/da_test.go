package arm64

import (
	"testing"

	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
)

func TestDa(t *testing.T) {
	//otp := regexp.MustCompile(`^(.+)_([0-9%*])$`)
	//
	//for k, _ := range map_op {
	//	m := otp.FindStringSubmatch(k)
	//	fmt.Printf("%s: %v\n", k, m[1])
	//}
	do_stmt("mvn x1, x2, lsl #47")
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
