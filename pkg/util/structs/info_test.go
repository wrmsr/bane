package structs

import (
	"testing"

	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

type A struct {
	X int
	Y string
}

type B struct {
	A A
	Z int64
}

type C struct {
	B
	Z int32
}

func TestTool(t *testing.T) {
	st := NewStructInfoCache()
	for _, ty := range []any{
		rfl.TypeOf[A](),
		rfl.TypeOf[B](),
		rfl.TypeOf[C](),
	} {
		st.Info(ty)
	}

	c := C{}
	st.Info((*C)(nil)).Field("Z").SetValue(&c, int32(420))
	tu.AssertEqual(t, c.Z, int32(420))
}
