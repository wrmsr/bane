package lisp

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

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
)

//

type Parser struct {
	pr *peekReader
}

func NewParser(r io.Reader) *Parser {
	return &Parser{
		pr: newPeekReader(r),
	}
}

//

var charMap = map[string]rune{
	"space":     ' ',
	"newline":   '\n',
	"backspace": '\b',
	"tab":       '\t',
	"linefeed":  '\n',
	"page":      '\f',
	"return":    '\r',
	"rubout":    0x7f,
}

func (pa *Parser) parseChar(ch string) Char {
	if charMap[ch] != 0 {
		return Char(charMap[ch])
	} else if chars := []rune(ch); len(chars) == 1 {
		return Char(chars[0])
	} else {
		panic(ParseError{fmt.Errorf("invalid character name #\\%s", ch)})
	}
}

func isAtomChar(ch rune) bool {
	return !(ch == '(' || ch == ')' || ch == '"' || isSpace(ch))
}

func (pa *Parser) parseSimple() Value {
	val := pa.pr.readWhile(func(ch rune) int {
		if isAtomChar(ch) {
			return 1
		}
		return 0
	})

	if val == "#t" {
		return Bool(true)
	} else if val == "#f" {
		return Bool(false)
	} else if strings.HasPrefix(val, `#\`) {
		return pa.parseChar(val[2:])
	} else if iv, err := strconv.ParseInt(val, 0, 64); err == nil {
		return Int(iv)
	} else if fv, err := strconv.ParseFloat(val, 64); err == nil {
		return Float(fv)
	} else if cv, err := strconv.ParseComplex(val, 128); err == nil {
		return Complex(cv)
	} else {
		return Atom(internedStrings.Get(val))
	}
}

func (pa *Parser) parseString() Value {
	if pa.pr.mustRead() != '"' {
		panic(UnexpectedCharErr)
	}

	var sb strings.Builder
	sb.WriteRune('"')
	pa.pr.readWhileInto(&sb, func(ch rune) int {
		if ch == '"' {
			return 0
		}
		if ch == '\\' {
			return 2
		}
		return 1
	})
	pa.pr.mustRead()
	sb.WriteRune('"')
	s := sb.String()

	ret, err := strconv.Unquote(s)
	if err != nil {
		panic(ParseError{err})
	} else {
		return String(ret)
	}
}

func (pa *Parser) parseValue() Value {
	pa.pr.skipSpace()
	ch := pa.pr.mustPeek()
	switch ch {
	case '\'':
		return MakeList(Atom("quote"), pa.parseValue())
	case ')':
		pa.pr.mustRead()
		return Atom(")")
	case '"':
		return pa.parseString()
	case '(':
		pa.pr.mustRead()
		return pa.parseList(false)
	default:
		return pa.parseSimple()
	}
}

func (pa *Parser) parseList(topLevel bool) *Cons {
	var p, q *Cons
	for {
		pa.pr.skipSpace()
		if pa.pr.isEof() {
			if !topLevel {
				panic(UnexpectedEofErr)
			}
			break
		}

		vv := pa.parseValue()
		if vv == Atom(")") {
			break
		}

		if vv != Atom(".") {
			AppendValue(&p, &q, vv)
			continue
		}

		if pa.pr.isEof() {
			panic(ParseError{errors.New("cdr expression expected")})
		}
		q.Cdr = pa.parseValue()
		if pa.pr.isEof() || pa.parseValue() != Atom(")") {
			panic(ParseError{errors.New("')' expected")})
		}
	}
	return p
}

//

func (pa *Parser) Parse() (ret *Cons, err error) {
	//defer func() {
	//	if r := recover(); r != nil {
	//		if r, ok := r.(ParseError); ok {
	//			err = eu.Append(err, r)
	//		} else {
	//			panic(r)
	//		}
	//	}
	//}()

	var body *Cons
	if !pa.pr.isEof() {
		body = pa.parseList(true)
		if !pa.pr.isEof() {
			err = eu.Append(err, UnexpectedCharErr)
		}
	}

	return &Cons{
		Atom("begin"),
		body,
	}, nil
}
