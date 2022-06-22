package lisp

import (
	"fmt"
	"strings"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

func TestCompilerSimple(t *testing.T) {
	src := `
1
    `
	fmt.Println(NewCompiler().Compile(check.Must(NewParser(strings.NewReader(src)).Parse())))
}
