//go:build !nodev

package test

import (
	"testing"

	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
)

func _intThing(x, y int) int {
	return x + y + 2
}

func InlineProxy[T any](v T) T {
	return v
}

var intThing = InlineProxy(_intThing)

func TestInline(t *testing.T) {
	tu.AssertEqual(t, Bar(2, 3), (2+3)*3)
	tu.AssertEqual(t, _def_inl_Bar(2, 3), (2+3)*3)

	tu.AssertEqual(t, Baz(InlineThing{x: 5}, 6), 12)
	tu.AssertEqual(t, _def_inl_Baz(InlineThing{x: 5}, 6), 12)

	tu.AssertEqual(t, intThing(2, 3), 2+3+2)
}
