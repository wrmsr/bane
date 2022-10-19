package wasm

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/dev/paths"
)

func testWasm(src string) {
	r := Reader{r: bufio.NewReader(strings.NewReader(src))}
	root := check.NotNil(r.ReadElement()).(List)
	check.Nil(r.ReadElement())

	fmt.Println(RenderString(root))

	BuildModule(root)
}

func TestWasm(t *testing.T) {
	src := `
(module
  (memory 4096 4096
    (segment 1026 "\14\00")
  )
  (func $big_negative
    (local $temp f64)
    (block $block0
      (set_local $temp
        (f64.const -2147483648)
      )
      (set_local $temp
        (f64.const -2147483648)
      )
      (set_local $temp
        (f64.const -21474836480)
      )
      (set_local $temp
        (f64.const 0.039625)
      )
      (set_local $temp
        (f64.const -0.039625)
      )
    )
  )
)
`

	testWasm(src)
}

func TestWasmBig(t *testing.T) {
	src := string(check.Must1(os.ReadFile(filepath.Join(paths.FindProjectRoot(), "exp", "util", "wasm", "test", "boilerplate.wast"))))
	testWasm(src)
}
