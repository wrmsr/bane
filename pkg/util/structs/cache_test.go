package structs

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
	ju "github.com/wrmsr/bane/pkg/util/json"
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
		si := st.Info(ty)
		fmt.Println(check.Must1(ju.MarshalIndentString(structInfoRepr(si), "", "  ")))
	}

	c := C{}
	st.Info((*C)(nil)).Fields().ByFlat()["Z"].SetValue(&c, int32(420))
	tu.AssertEqual(t, c.Z, int32(420))
}
