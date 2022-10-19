package parsing

import (
	"bufio"
	"io"
	"strings"

	"github.com/wrmsr/bane/pkg/util/slices"
)

type Reader struct {
	r *bufio.Reader
}

func NewReader(r io.Reader) *Reader {
	return &Reader{r: bufio.NewReader(r)}
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

func (r *Reader) ReadElement() Element {
	var stk slices.Stack[[]Element]
l:
	for {
		var p Element
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
				return List{Ps: stk[0]}
			}
			p = List{Ps: stk.Pop()}

		case s[0] == '"':
			if s[0] != '"' || s[len(s)-1] != '"' {
				panic("mismatched quotes")
			}
			p = Quote{S: s[1 : len(s)-1]}

		default:
			p = Atom{S: s}
		}

		if len(stk) > 0 {
			stk[len(stk)-1] = append(stk[len(stk)-1], p)
			continue
		}

		return p
	}
}
