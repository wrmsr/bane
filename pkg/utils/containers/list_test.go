package containers

import (
	"testing"

	tu "github.com/wrmsr/bane/pkg/utils/dev/testing"
)

func TestList(t *testing.T) {
	l := NewMutListOf("zero", "one", "two", "three")
	tu.AssertEqual(t, l.Get(0), "zero")
	tu.AssertEqual(t, l.Get(1), "one")
	tu.AssertEqual(t, l.Get(2), "two")
	tu.AssertEqual(t, l.Get(3), "three")
	tu.AssertEqual(t, l.Len(), 4)

	l.Delete(2)
	tu.AssertEqual(t, l.Get(2), "three")
	tu.AssertEqual(t, l.Len(), 3)
}
