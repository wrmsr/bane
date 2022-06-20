package strings

import (
	"strings"
	"testing"

	tu "github.com/wrmsr/bane/pkg/utils/dev/testing"
)

func TestIntern(t *testing.T) {
	sz := 256
	p := strings.Repeat("0123456789abcdef", sz)

	mk := func() string { return p + "!" }
	trn := Intern(mk())

	tu.AssertEqual(t, trn.Intern(mk()), mk())

	ct := 1_000_000
	a := make([]string, ct)
	for i := range a {
		a[i] = trn.Intern(mk())
		//a[i] = mk()
	}
}
