package marshal

import (
	"fmt"
	"testing"

	stu "github.com/wrmsr/bane/pkg/util/structs"
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

func TestMarshal(t *testing.T) {
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

	em := &Marshalling{stu.NewStructInfoCache()}
	m := em.Explode(c)
	fmt.Printf("%+v\n", m)
}
