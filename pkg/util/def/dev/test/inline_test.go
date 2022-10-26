//go:build !nodev

package test

import (
	"testing"

	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
)

func TestInline(t *testing.T) {
	tu.AssertEqual(t, Bar(2, 3), (2+3)*3)
	tu.AssertEqual(t, _def_inl_Bar(2, 3), (2+3)*3)

	tu.AssertEqual(t, Baz(InlineThing{x: 5}, 6), 12)
	tu.AssertEqual(t, _def_inl_Baz(InlineThing{x: 5}, 6), 12)
}
