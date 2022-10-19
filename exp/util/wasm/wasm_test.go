package wasm

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/wrmsr/bane/pkg/util/slices"
)

//

type Part interface {
	isPart()

	Render(w io.Writer)
}

type part struct{}

func (p part) isPart() {}

type Atom struct {
	part
	s string
}

var _ Part = Atom{}

type Quote struct {
	part
	s string
}

var _ Part = Quote{}

type List struct {
	part
	ps []Part
}

var _ Part = List{}

func RenderPart(p Part) string {
	var bs strings.Builder
	p.Render(&bs)
	return bs.String()
}

//

func (p Atom) Render(w io.Writer) {
	_, _ = io.WriteString(w, p.s)
}

func (p Quote) Render(w io.Writer) {
	_, _ = w.Write([]byte{'"'})
	_, _ = io.WriteString(w, p.s)
	_, _ = w.Write([]byte{'"'})
}

func (p List) Render(w io.Writer) {
	_, _ = w.Write([]byte{'('})
	for i, c := range p.ps {
		if i > 0 {
			_, _ = w.Write([]byte{' '})
		}
		c.Render(w)
	}
	_, _ = w.Write([]byte{')'})
}

//

type Reader struct {
	r *bufio.Reader
}

func (r *Reader) readByte() byte {
	b, err := r.r.ReadByte()
	if err != nil {
		if err == io.EOF {
			return 0
		}
		panic(err)
	}
	return b
}

func (r *Reader) readString() string {
	var sb strings.Builder

l:
	for {
		switch b := r.readByte(); b {
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
			if sb.Len() > 0 {
				panic("misplaced quote")
			}
			for {
				switch b := r.readByte(); b {
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

func (r *Reader) readPart() Part {
	var stk slices.Stack[[]Part]
l:
	for {
		var p Part
		switch s := r.readString(); {
		case s == "":
			break

		case s == "(":
			stk.Push(nil)
			continue l

		case s == ")":
			if len(stk) < 1 {
				panic("mismatched parens")
			}
			if len(stk) < 2 {
				return List{ps: stk[0]}
			}
			p = List{ps: stk.Pop()}

		case s[0] == '"':
			if s[0] != '"' || s[len(s)-1] != '"' {
				panic("mismatched quotes")
			}
			p = Quote{s: s[1 : len(s)-1]}

		default:
			p = Atom{s: s}
		}

		if len(stk) > 0 {
			stk[len(stk)-1] = append(stk[len(stk)-1], p)
			continue
		}

		return p
	}
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

	r := Reader{r: bufio.NewReader(strings.NewReader(src))}
	for {
		p := r.readPart()
		if p == nil {
			break
		}
		fmt.Println(RenderPart(p))
	}
}
