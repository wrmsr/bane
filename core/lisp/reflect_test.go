package lisp

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/wrmsr/bane/core/check"
)

type Foo struct {
	Bar string
}

func (f Foo) Frob() {
	fmt.Println("frob")
}

func TestReflect(t *testing.T) {
	src := `
(display (gf x "Bar"))
(cm x "Frob")
`
	prog := NewCompiler().Compile(check.Must1(NewParser(strings.NewReader(src)).Parse()))
	sc := NewScope(nil)
	sc.Set("x", Reflect{v: reflect.ValueOf(Foo{Bar: "baz"})})
	Evaluate(addIntrinsics(sc), prog)
}
