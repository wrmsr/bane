package wasm

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

type Reader struct {
	r     io.Reader
	depth int
}

func (r *Reader) Next() bool {
	b := [1]byte{}
	n, err := r.r.Read(b[:])
	if err != nil {
		if err != io.EOF {
			panic(err)
		}
		return false
	}
	if n > 1 {
		panic(n)
	}
	fmt.Println(b[0])
	return true
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
    )
  )
)
`

	r := Reader{r: strings.NewReader(src)}
	for r.Next() {
	}
}
