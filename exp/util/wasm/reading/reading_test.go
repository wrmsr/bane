package reading

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/dev/paths"
)

func TestReading(t *testing.T) {
	src := check.Must1(os.ReadFile(filepath.Join(paths.FindProjectRoot(), "exp", "util", "wasm", "test", "go", "main_tiny.wasm")))
	r := bytes.NewReader(src)
	r.ReadByte()
}
