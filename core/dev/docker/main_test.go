//go:build !nodev

package docker

import (
	"testing"

	"github.com/wrmsr/bane/core/dev"
)

func TestMain(m *testing.M) {
	dev.DoTestMain(m, func() {})
}
