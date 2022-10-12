package test

import (
	"testing"

	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
)

func TestInline(t *testing.T) {
	tu.AssertEqual(t, Bar(2, 3), (2+3)*3)
	//tu.AssertEqual(t, _def_inl_Bar(2, 3), (2+3)*3)
}
