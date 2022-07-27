//go:build !nodev

package impl

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/wrmsr/bane/pkg/util/dev/paths"
)

func TestImpl(t *testing.T) {
	cwd := filepath.Join(paths.FindProjectRoot(), "pkg/util/def/dev/test")
	s := Run(cwd)
	fmt.Println(s)
}
