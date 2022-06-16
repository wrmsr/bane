package structs

import "testing"

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
		(*A)(nil),
		(*B)(nil),
		(*C)(nil),
	} {
		st.Info(ty)
	}
}
