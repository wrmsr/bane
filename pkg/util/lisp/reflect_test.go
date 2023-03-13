package lisp

import (
	"reflect"
	"strings"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

type Foo struct {
	Bar string
}

func TestReflect(t *testing.T) {
	src := `(display (gf x "Bar"))`
	prog := NewCompiler().Compile(check.Must1(NewParser(strings.NewReader(src)).Parse()))
	sc := NewScope(nil)
	sc.Set("x", Reflect{v: reflect.ValueOf(Foo{Bar: "baz"})})
	Evaluate(addIntrinsics(sc), prog)
}
