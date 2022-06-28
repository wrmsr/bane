package lisp

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

func TestParseSimple(t *testing.T) {
	s := check.Must1(ioutil.ReadFile("test/simple.scm"))
	p := NewParser(bytes.NewReader(s))
	root := check.Must1(p.Parse())
	fmt.Println(root)
}

func TestParseMandelbrot(t *testing.T) {
	s := check.Must1(ioutil.ReadFile("test/mandelbrot.scm"))
	p := NewParser(bytes.NewReader(s))
	root := check.Must1(p.Parse())
	fmt.Println(root)
}
