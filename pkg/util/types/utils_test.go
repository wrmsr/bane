package types

import (
	"testing"

	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
)

type FooI0 interface{ isFooI0() }

type FooI1 interface{ isFooI1() }

type FooS0 struct{}

func (f FooS0) isFooI0() {}

type FooS1 struct{}

func (f FooS1) isFooI0() {}
func (f FooS1) isFooI1() {}

type FooS2 struct{}

func (f FooS2) isFooI1() {}

func TestCanAssign(t *testing.T) {
	tu.AssertEqual(t, CanAssign[FooS0, FooI0](), true)
	tu.AssertEqual(t, CanAssign[FooS0, FooI1](), false)

	tu.AssertEqual(t, CanAssign[FooS1, FooI0](), true)
	tu.AssertEqual(t, CanAssign[FooS1, FooI1](), true)

	tu.AssertEqual(t, CanAssign[FooS2, FooI0](), false)
	tu.AssertEqual(t, CanAssign[FooS2, FooI1](), true)
}
