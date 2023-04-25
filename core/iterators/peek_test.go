package iterators

import (
	"testing"

	tu "github.com/wrmsr/bane/core/dev/testing"
)

func TestPeek(t *testing.T) {
	pi := Peek(Range(0, 4, 1)).Iterate()

	tu.AssertEqual(t, pi.HasNext(), true)
	tu.AssertEqual(t, pi.Peek(), 0)
	tu.AssertEqual(t, pi.HasNext(), true)
	tu.AssertEqual(t, pi.Peek(), 0)
	tu.AssertEqual(t, pi.Next(), 0)

	tu.AssertEqual(t, pi.Peek(), 1)
	tu.AssertEqual(t, pi.Next(), 1)

	tu.AssertEqual(t, pi.HasNext(), true)
	tu.AssertEqual(t, pi.Peek(), 2)
	tu.AssertEqual(t, pi.Next(), 2)

	tu.AssertEqual(t, pi.HasNext(), true)
	tu.AssertEqual(t, pi.Peek(), 3)
	tu.AssertEqual(t, pi.Next(), 3)

	tu.AssertEqual(t, pi.HasNext(), false)
}
