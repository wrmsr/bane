package impl

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/wrmsr/bane/pkg/util/dev/paths"
)

func TestImpl(t *testing.T) {
	cwd := filepath.Join(paths.FindProjectRoot(), "exp/util/annotate/dev/test")
	s := Run(cwd)
	fmt.Println(s)
}
