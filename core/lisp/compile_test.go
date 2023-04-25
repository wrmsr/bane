package lisp

import (
	"fmt"
	"strings"
	"testing"

	"github.com/wrmsr/bane/core/check"
)

func TestCompilerSimple(t *testing.T) {
	src := `
1
    `
	fmt.Println(NewCompiler().Compile(check.Must1(NewParser(strings.NewReader(src)).Parse())))
}
