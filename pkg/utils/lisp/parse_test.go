package lisp

import (
	"fmt"
	"strings"
	"testing"

	"github.com/wrmsr/bane/pkg/utils/check"
)

func TestParse(t *testing.T) {
	p := NewParser(strings.NewReader("   1"))
	root := check.Must(p.Parse())
	fmt.Println(root)
}
