package lisp

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/wrmsr/bane/core/check"
)

func TestParseSimple(t *testing.T) {
	s := check.Must1(os.ReadFile("test/simple.scm"))
	p := NewParser(bytes.NewReader(s))
	root := check.Must1(p.Parse())
	fmt.Println(root)
}

func TestParseMandelbrot(t *testing.T) {
	s := check.Must1(os.ReadFile("test/mandelbrot.scm"))
	p := NewParser(bytes.NewReader(s))
	root := check.Must1(p.Parse())
	fmt.Println(root)
}
