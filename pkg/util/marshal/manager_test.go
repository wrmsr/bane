package marshal

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	ju "github.com/wrmsr/bane/pkg/util/json"
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

var testC = C{
	B: B{
		A: A{
			X: 100,
			Y: "two hundred",
		},
		Z: 300,
	},
	Z: 420,
}

func TestMarshal(t *testing.T) {
	em := NewManager()
	m := em.Marshal(testC)
	fmt.Printf("%+v\n", m)
	fmt.Println(check.Must1(ju.MarshalString(m)))
}
