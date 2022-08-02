package types

import (
	"testing"

	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
)

func TestDefaultCmp(t *testing.T) {
	tu.AssertEqual(t, DefaultCmpImpl[int32]()(int32(100), int32(200)), CmpLesser)
	tu.AssertEqual(t, DefaultCmpImpl[string]()("100", "200"), CmpLesser)
}
