package containers

import (
	"testing"

	tu "github.com/wrmsr/bane/pkg/utils/dev/testing"
)

func TestSet(t *testing.T) {
	s := NewMutSetOf(10, 20, 30)
	tu.AssertEqual(t, s.Contains(10), true)
	tu.AssertEqual(t, s.Contains(11), false)

	s.Add(11)
	tu.AssertEqual(t, s.Contains(11), true)
}
