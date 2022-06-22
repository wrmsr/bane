package lisp

import (
	"bufio"
	"io"
	"strings"

	ctr "github.com/wrmsr/bane/pkg/util/container"
	opt "github.com/wrmsr/bane/pkg/util/optional"
)

type peekReader struct {
	rd *bufio.Reader

	pch opt.Optional[rune]
}

func newPeekReader(r io.Reader) *peekReader {
	return &peekReader{
		rd: bufio.NewReader(r),
	}
}

//

func (pr *peekReader) peek() opt.Optional[rune] {
	if pr.pch.Present() {
		return pr.pch
	}

	och := pr.read()
	if !och.Present() {
		return opt.None[rune]()
	}

	pr.pch = och
	return pr.pch
}

func (pr *peekReader) mustPeek() rune {
	if och := pr.peek(); och.Present() {
		return och.Value()
	}
	panic(UnexpectedEofErr)
}

func (pr *peekReader) isEof() bool {
	och := pr.peek()
	return !och.Present()
}

func (pr *peekReader) read() opt.Optional[rune] {
	if pr.pch.Present() {
		och := pr.pch
		pr.pch = opt.None[rune]()
		return och
	}

	ch, _, e := pr.rd.ReadRune()
	if e != nil {
		if e == io.EOF {
			return opt.None[rune]()
		}
		panic(ParseError{e})
	}

	return opt.Just(ch)
}

func (pr *peekReader) mustRead() rune {
	if r := pr.read(); r.Present() {
		return r.Value()
	}
	panic(UnexpectedEofErr)
}

//

func (pr *peekReader) skipWhile(fn func(ch rune) bool) {
	for {
		och := pr.peek()
		if !och.Present() || !fn(och.Value()) {
			break
		}
		pr.read()
	}
}

var spaceRunes = ctr.NewSetOf(
	' ',
	'\f',
	'\n',
	'\r',
	'\t',
	'\v',
)

func isSpace(r rune) bool {
	return spaceRunes.Contains(r)
}

func (pr *peekReader) skipSpace() {
	pr.skipWhile(isSpace)
}

type RuneWriter interface {
	WriteRune(r rune) (int, error)
}

func (pr *peekReader) readWhileInto(w RuneWriter, fn func(ch rune) int) {
	for {
		och := pr.peek()
		if !och.Present() {
			break
		}

		r := fn(och.Value())
		if r < 1 {
			break
		}

		for ; r > 0; r-- {
			_, _ = w.WriteRune(pr.mustRead())
		}
	}
}

func (pr *peekReader) readWhile(fn func(ch rune) int) string {
	var sb strings.Builder
	pr.readWhileInto(&sb, fn)
	return sb.String()
}
