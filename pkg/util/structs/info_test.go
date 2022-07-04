package structs

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	ju "github.com/wrmsr/bane/pkg/util/json"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	"github.com/wrmsr/bane/pkg/util/slices"
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
	sm := getMapping(rfl.TypeOf[C]())
	fmt.Println(sm)

	sm2 := getSimpleFieldInfo(rfl.TypeOf[C](), nil, "")
	fmt.Println(sm2)
	fmt.Println(check.Must1(ju.MarshalIndentString(slices.Map(fieldInfoRepr, sm2), "", "  ")))

	st := NewStructInfoCache()
	for _, ty := range []any{
		//rfl.TypeOf[A](),
		//rfl.TypeOf[B](),
		rfl.TypeOf[C](),
	} {
		st.Info(ty)
	}

	//c := C{}
	//st.Info((*C)(nil)).Field("Z").SetValue(&c, int32(420))
	//tu.AssertEqual(t, c.Z, int32(420))
}
