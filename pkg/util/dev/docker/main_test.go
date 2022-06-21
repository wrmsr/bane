//go:build !nodev

package docker

import (
	"testing"

	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
)

func TestMain(m *testing.M) {
	tu.Main(m, func() {})
}
