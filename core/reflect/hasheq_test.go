package reflect

import (
	"testing"

	tu "github.com/wrmsr/bane/core/dev/testing"
)

func TestHashEq(t *testing.T) {
	ps0 := new(string)
	*ps0 = "foo"

	ps1 := new(string)
	*ps1 = *ps0

	tu.AssertEqual(t, PtrHashEq[string]().Eq(ps0, ps0), true)
	tu.AssertEqual(t, PtrHashEq[string]().Eq(ps0, ps1), false)
}
