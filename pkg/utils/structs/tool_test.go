package structs

import (
	"fmt"
	"testing"

	rfl "github.com/wrmsr/bane/pkg/utils/reflect"
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
	st := NewStructTool()
	for _, ty := range []any{
		rfl.Type[A](),
		rfl.Type[B](),
		rfl.Type[C](),
	} {
		st.Info(ty)
	}
}

func TestToMap(t *testing.T) {
	c := C{
		B: B{
			A: A{
				X: 100,
				Y: "two hundred",
			},
			Z: 300,
		},
		Z: 420,
	}

	m := NewStructTool().ToMap(c)
	fmt.Printf("%+v\n", m)
}
