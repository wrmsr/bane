package reading

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/wrmsr/bane/core/check"
	"github.com/wrmsr/bane/core/dev/paths"
)

func TestReading(t *testing.T) {
	src := check.Must1(os.ReadFile(filepath.Join(paths.FindProjectRoot(), "exp", "util", "wasm", "test", "go", "main_tiny.wasm")))
	r := &ModuleReader{r: NewByteReader(src)}
	r.ReadModule()
}
