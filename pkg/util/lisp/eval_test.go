package lisp

import (
	"fmt"
	"strings"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

func TestEvalSimple(t *testing.T) {
	src := `
1
    `

	prog := NewCompiler().Compile(check.Must(NewParser(strings.NewReader(src)).Parse()))
	ret := Evaluate(NewScope(nil), prog)
	fmt.Println(ret)
}
