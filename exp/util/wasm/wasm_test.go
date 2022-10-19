package wasm

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"testing"
)

//

type Part interface {
	isPart()
}

type part struct{}

func (p part) isPart() {}

type Atom struct {
	part
	s string
}

type String struct {
	part
	s string
}

type List struct {
	part
	ps []Part
}

//

type Reader struct {
	r *bufio.Reader
}

func (r *Reader) read() byte {
	b, err := r.r.ReadByte()
	if err != nil {
		if err == io.EOF {
			return 0
		}
		panic(err)
	}
	return b
}

func (r *Reader) next() string {
	var sb strings.Builder

l:
	for {
		b := r.read()
		switch b {
		case 0:
			break l

		case '(':
			if sb.Len() < 1 {
				return "("
			}
			break l

		case ')':
			if sb.Len() < 1 {
				return ")"
			}
			if err := r.r.UnreadByte(); err != nil {
				panic(err)
			}
			break l

		case '"':
			for {
				b := r.read()
				switch b {
				case 0:
					panic("mismatched quotes")

				case '"':
					break l

				default:
					sb.WriteByte(b)
				}
			}

		case ' ', '\n', '\r', '\t':
			if sb.Len() > 0 {
				break l
			}

		default:
			sb.WriteByte(b)
		}
	}

	return sb.String()
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

	r := Reader{r: bufio.NewReader(strings.NewReader(src))}
	for {
		s := r.next()
		if s == "" {
			break
		}
		fmt.Println(s)
	}
}
