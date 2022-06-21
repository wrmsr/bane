package dispatch

import (
	"fmt"
	"testing"

	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
)

//

type A struct{}
type B struct{}
type C struct {
	a A
	b B
}

//

type Foo struct {
	_Foo_do _Foo_do // TODO: codegen
}

//

func (f Foo) do_A(a A) string { return "Foo.do_A" }
func (f Foo) do_B(b B) string { return "Foo.do_B" }
func (f Foo) do_C(c C) string { return "Foo.do_C, " + f.do(c.a) + ", " + f.do(c.b) }

//

type Foo2 struct{ Foo }

func (f Foo2) do_B(b B) string { return "Foo2.do_B" }

//
// TODO: codegen

type _Foo_do interface {
	_do(v any) string
}

func (f Foo) do(v any) string {
	return f._Foo_do._do(v)
}

func (f *Foo) _init_do() {
	f._Foo_do = f
}

func (f Foo) _do(v any) string {
	switch v := v.(type) {
	case A:
		return f.do_A(v)
	case B:
		return f.do_B(v)
	case C:
		return f.do_C(v)
	}
	panic(fmt.Errorf("unhandled: %T", v))
}

func (f *Foo2) _init_do() {
	f.Foo._Foo_do = f
}

func (f Foo2) _do(v any) string {
	switch v := v.(type) {
	case A:
		return f.do_A(v)
	case B:
		return f.do_B(v)
	case C:
		return f.do_C(v)
	}
	panic(fmt.Errorf("unhandled: %T", v))
}

//

func TestDispatch(t *testing.T) {
	foo := Foo{}
	foo._init_do()
	tu.AssertEqual(t, foo.do(A{}), "Foo.do_A")
	tu.AssertEqual(t, foo.do(B{}), "Foo.do_B")
	tu.AssertEqual(t, foo.do(C{}), "Foo.do_C, Foo.do_A, Foo.do_B")

	foo2 := Foo2{}
	foo2._init_do()
	tu.AssertEqual(t, foo2.do(A{}), "Foo.do_A")
	tu.AssertEqual(t, foo2.do(B{}), "Foo2.do_B")
	tu.AssertEqual(t, foo2.do(C{}), "Foo.do_C, Foo.do_A, Foo2.do_B")
}
