package lisp

import (
	"bufio"
	"io"

	ctr "github.com/wrmsr/bane/pkg/utils/containers"
)

//

type runeReader struct {
	rd *bufio.Reader

	p bool
	r rune
	e error
}

func newRuneReader(r io.Reader) runeReader {
	return runeReader{
		rd: bufio.NewReader(r),
	}
}

func (rd *runeReader) peekErr() (rune, error) {
	if rd.p {
		return rd.r, rd.e
	}

	rd.p = true
	rd.r, _, rd.e = rd.rd.ReadRune()
	return rd.r, rd.e
}

func (rd *runeReader) nextErr() (rune, error) {
	if rd.p {
		r, e := rd.r, rd.e
		rd.p = false
		rd.r, rd.e = 0, nil
		return r, e
	}

	r, _, e := rd.rd.ReadRune()
	return r, e
}

func (rd *runeReader) peek() rune {
	r, e := rd.peekErr()
	if e != nil {
		panic(e)
	}
	return r
}

func (rd *runeReader) next() rune {
	r, e := rd.nextErr()
	if e != nil {
		panic(e)
	}
	return r
}

func (rd *runeReader) more() bool {
	_, e := rd.peekErr()
	return e == nil
}

func (rd *runeReader) err() error {
	if !rd.p {
		return nil
	}
	return rd.e
}

//

type Parser struct {
	rd runeReader
}

func NewParser(r io.Reader) *Parser {
	return &Parser{
		rd: newRuneReader(r),
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

func (p *Parser) skipSpace() {
	for p.rd.more() && spaceRunes.Contains(p.rd.peek()) {
		p.rd.next()
	}
}

func (p *Parser) Parse() (*Cons, error) {
	p.skipSpace()

	c := &Cons{
		Int(10),
		Int(20),
	}

	if err := p.rd.err(); err != io.EOF {
		return c, err
	}
	return c, nil
}
