package lisp

import (
	"fmt"
	"strings"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

func TestEvalSimple(t *testing.T) {
	src := `
(if (> (+ 1 2) 3) 4 3)
    `

	prog := NewCompiler().Compile(check.Must1(NewParser(strings.NewReader(src)).Parse()))
	fmt.Println(prog)
	ret := Evaluate(addIntrinsics(NewScope(nil)), prog)
	fmt.Println(ret)
}
