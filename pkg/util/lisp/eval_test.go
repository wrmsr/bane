package lisp

import (
	"fmt"
	"strings"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

func TestEvalSimple(t *testing.T) {
	src := `
(+ 1 2)
    `

	prog := NewCompiler().Compile(check.Must(NewParser(strings.NewReader(src)).Parse()))
	fmt.Println(prog)
	ret := Evaluate(addIntrinsics(NewScope(nil)), prog)
	fmt.Println(ret)
}
