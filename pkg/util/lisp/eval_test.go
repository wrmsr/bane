package lisp

import (
	"fmt"
	"strings"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

func TestEvalSimple(t *testing.T) {
	/*
		; (define y 2)
		; (define (addy x) (+ y x))
		; (display (addy 2))

		; (define z 3)
		; (display (addy z))
		; (set! z 4)
		; (display (addy z))
	*/

	src := `(do ((x 1)) ((< 5 x)) (display x) (set! x (+ x 1)))`

	p := check.Must1(NewParser(strings.NewReader(src)).Parse())

	prog := NewCompiler().Compile(p)
	fmt.Println(prog)

	Evaluate(addIntrinsics(NewScope(nil)), prog)
}
