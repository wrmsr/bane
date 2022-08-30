package lisp

import (
	"fmt"
	"strings"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

func TestEvalSimple(t *testing.T) {
	src := `
(define y 2)
(define (addy x) (+ y x))
(display (addy 2))
    `

	p := check.Must1(NewParser(strings.NewReader(src)).Parse())

	prog := NewCompiler().Compile(p)
	fmt.Println(prog)

	Evaluate(addIntrinsics(NewScope(nil)), prog)
}
