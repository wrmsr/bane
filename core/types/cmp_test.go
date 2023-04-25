package types

import (
	"testing"

	tu "github.com/wrmsr/bane/core/dev/testing"
)

type myInt int

func TestDefaultCmp(t *testing.T) {
	tu.AssertEqual(t, DefaultCmpImpl[int32]()(int32(100), int32(200)), CmpLesser)
	tu.AssertEqual(t, DefaultCmpImpl[string]()("100", "200"), CmpLesser)

	// FIXME:
	//tu.AssertEqual(t, DefaultCmpImpl[myInt]()(myInt(100), myInt(200)), CmpLesser)
}
