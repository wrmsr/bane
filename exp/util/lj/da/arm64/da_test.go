package arm64

import (
	"regexp"
	"strings"
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

func splitStmt(s string) (string, []string) {
	rs := []rune(s)
	var parts []string
	var stk []rune
	l := 0
	for i, r := range rs {
		switch r {
		case '(':
			stk = append(stk, ')')
		case '[':
			stk = append(stk, '[')
		case '{':
			stk = append(stk, '{')
		case ')', ']', '}':
			if stk[len(stk)-1] != r {
				panic("unmatched")
			}
			stk = stk[:len(stk)-1]
		case ',':
			parts = append(parts, string(rs[l:i]))
			l = i + 1
		}
	}
	if len(stk) > 0 {
		panic("unclosed")
	}
	parts = append(parts, string(rs[l:]))

	m := regexp.MustCompile(`^\s*(\S+)(.*)$`).FindStringSubmatch(parts[0])
	op := m[1]
	if m[2] != "" {
		parts[0] = m[2]
	} else {
		parts = parts[1:]
	}

	for i, p := range parts {
		parts[i] = strings.TrimSpace(p)
	}

	return op, parts
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
		op, params := splitStmt(tc.s)
		tu.AssertDeepEqual(t, op, tc.op)
		tu.AssertDeepEqual(t, params, tc.params)
	}
}
