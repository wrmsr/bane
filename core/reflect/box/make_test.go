package box

import (
	"fmt"
	"testing"
)

type foo struct {
	X int
	S string
}

type bar struct {
	foos []foo
	f    float
}

func TestBox(t *testing.T) {
	b := BoxOf(420)
	fmt.Printf("%#v\n", b)

	b2 := BoxOf(foo{})
	fmt.Printf("%#v\n", b2)
}
