package containers

import (
	"testing"

	tu "github.com/wrmsr/bane/pkg/utils/dev/testing"
)

func TestMap(t *testing.T) {
	m := NewMutMap[int, string]()
	m.Put(10, "ten")
	m.Put(20, "twenty")
	m.Put(30, "thirty")
	tu.AssertEqual(t, m.Contains(10), true)
	tu.AssertEqual(t, m.Get(10), "ten")
	tu.AssertEqual(t, m.Contains(11), false)

	m.Put(11, "eleven")
	tu.AssertEqual(t, m.Contains(11), true)
}
