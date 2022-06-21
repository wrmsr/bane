package lisp

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	ctr "github.com/wrmsr/bane/pkg/util/container"
	eu "github.com/wrmsr/bane/pkg/util/errors"
)

//

type ParseError struct {
	e error
}

func (e ParseError) Error() string {
	return e.e.Error()
}

func (e ParseError) Unwrap() error {
	return e.e
}

var (
	UnexpectedEofErr  = ParseError{errors.New("unexpected EOF")}
	UnexpectedCharErr = ParseError{errors.New("unexpected char")}
	IllegalCharErr    = ParseError{errors.New("illegal char")}
)

//

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

func isAtomChar(r rune) bool {
	return !(r == '(' || r == ')' || r == '"' || isSpace(r))
}

//

type Parser struct {
	rd *bufio.Reader

	pp bool
	pr rune
}

func NewParser(r io.Reader) *Parser {
	return &Parser{
		rd: bufio.NewReader(r),
	}
}

func (pa *Parser) peek() (rune, bool) {
	if pa.pp {
		return pa.pr, true
	}

	r, ok := pa.read()
	if !ok {
		return 0, false
	}

	pa.pp, pa.pr = true, r
	return r, true
}

func (pa *Parser) mustPeek() rune {
	if r, ok := pa.peek(); ok {
		return r
	}
	panic(UnexpectedEofErr)
}

func (pa *Parser) read() (rune, bool) {
	if pa.pp {
		r := pa.pr
		pa.pp = false
		return r, true
	}

	r, _, e := pa.rd.ReadRune()
	if e != nil {
		if e == io.EOF {
			return 0, false
		}
		panic(ParseError{e})
	}

	return r, true
}

func (pa *Parser) mustRead() rune {
	if r, ok := pa.read(); ok {
		return r
	}
	panic(UnexpectedEofErr)
}

func (pa *Parser) skipSpace() {
	for {
		r, ok := pa.peek()
		if !ok || !isSpace(r) {
			break
		}
		pa.read()
	}
}

func (pa *Parser) parseString() Value {
	if pa.mustRead() != '"' {
		panic(UnexpectedCharErr)
	}

	var sb strings.Builder
	for {
		c := pa.mustRead()
		if c == '"' {
			break
		}
		sb.WriteRune(c)
		if c == '\\' {
			sb.WriteRune(pa.mustRead())
		}

	}

	ret, err := strconv.Unquote(sb.String())
	if err != nil {
		panic(ParseError{err})
	} else {
		return String(ret)
	}
}

func (pa *Parser) parseValue() Value {
	pa.skipSpace()

	r := pa.mustPeek()
	switch r {
	case 0:
		panic(IllegalCharErr)
	case '"':
		return pa.parseString()
	}

	panic(ParseError{fmt.Errorf("unexpected character: %v", r)})

}

//func (pa *Parser) parseSimple() Value {
//	p := pa.p - 1
//	n := len(pa.s)
//
//	/* scan until the next space or EOF */
//	for pa.p < n && isAtomChar(pa.s[pa.p]) {
//		pa.p++
//	}
//
//	/* extract the token */
//	src := pa.s[p:pa.p]
//	val := string(src)
//
//	/* check for token types */
//	if iv, err := strconv.ParseInt(val, 0, 64); err == nil {
//		return Int(iv)
//	} else if fv, err := strconv.ParseFloat(val, 64); err == nil {
//		return Float(fv)
//	} else if cv, err := strconv.ParseComplex(val, 128); err == nil {
//		return Complex(cv)
//	} else {
//		return Atom(val)
//	}
//}

func (pa *Parser) Parse() (ret *Cons, err error) {
	defer func() {
		if r := recover(); r != nil {
			if r, ok := r.(ParseError); ok {
				err = eu.Append(err, r)
			} else {
				panic(r)
			}
		}
	}()

	ret = &Cons{
		pa.parseValue(),
		nil,
	}

	if _, ok := pa.peek(); ok {
		err = eu.Append(err, UnexpectedCharErr)
	}
	return
}
