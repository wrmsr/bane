package wasm

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

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

	r := Reader{r: bufio.NewReader(strings.NewReader(src))}
	root := check.NotNil(r.ReadElement()).(List)
	check.Nil(r.ReadElement())

	fmt.Println(RenderElement(root))

	BuildModule(root)
}
