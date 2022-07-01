package cache

import (
	"testing"

	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
)

func TestCache(t *testing.T) {
	c := NewCache(DefaultConfig())
	c.Put(100, 200)
	tu.AssertEqual(t, c.Get(100), 200)
}
