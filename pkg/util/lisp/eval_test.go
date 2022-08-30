package lisp

import (
	"fmt"
	"strings"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

func TestEvalSimple(t *testing.T) {
	src := `
(define (add1 x) (+ 1 x))
(display (add1 2))
    `

	p := check.Must1(NewParser(strings.NewReader(src)).Parse())

	prog := NewCompiler().Compile(p)
	fmt.Println(prog)

	ret := Evaluate(addIntrinsics(NewScope(nil)), prog)
	fmt.Println(ret)
}
