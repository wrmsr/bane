//go:build !nodev

package docker

import (
	"testing"

	"github.com/wrmsr/bane/pkg/util/dev"
)

func TestMain(m *testing.M) {
	dev.DoTestMain(m, func() {})
}
