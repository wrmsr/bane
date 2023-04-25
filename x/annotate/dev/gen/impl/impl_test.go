//go:build !nodev

package impl

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/wrmsr/bane/core/dev/paths"
)

func TestImpl(t *testing.T) {
	cwd := filepath.Join(paths.FindProjectRoot(), "x/util/annotate/dev/test")
	s := Run(cwd)
	fmt.Println(s)
}
