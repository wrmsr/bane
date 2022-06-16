//go:build !nodev

package docker

import (
	"testing"

	bts "github.com/wrmsr/bane/pkg/dev/testing"
)

func TestMain(m *testing.M) {
	bts.Main(m, func() {})
}
