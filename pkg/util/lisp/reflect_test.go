package lisp

import (
	"reflect"
	"strings"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

func TestReflect(t *testing.T) {
	src := `(display x)`
	prog := NewCompiler().Compile(check.Must1(NewParser(strings.NewReader(src)).Parse()))
	sc := NewScope(nil)
	sc.Set("x", Reflect{v: reflect.ValueOf(420)})
	Evaluate(addIntrinsics(sc), prog)
}
