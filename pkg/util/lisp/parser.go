package lisp

import (
	"errors"
	"io"
	"strconv"

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

type Parser struct {
	pr *peekReader
}

func NewParser(r io.Reader) *Parser {
	return &Parser{
		pr: newPeekReader(r),
	}
}

//

func isAtomChar(r rune) bool {
	return !(r == '(' || r == ')' || r == '"' || isSpace(r))
}

func (pa *Parser) parseSimple() Value {
	val := pa.pr.readWhile(func(ch rune) int {
		if isAtomChar(ch) {
			return 1
		}
		return 0
	})

	if iv, err := strconv.ParseInt(val, 0, 64); err == nil {
		return Int(iv)
	} else if fv, err := strconv.ParseFloat(val, 64); err == nil {
		return Float(fv)
	} else if cv, err := strconv.ParseComplex(val, 128); err == nil {
		return Complex(cv)
	} else {
		return Atom(internedStrings.Intern(val))
	}
}

func (pa *Parser) parseString() Value {
	if pa.pr.mustRead() != '"' {
		panic(UnexpectedCharErr)
	}

	s := pa.pr.readWhile(func(ch rune) int {
		if ch == '"' {
			return 0
		}
		if ch == '\\' {
			return 2
		}
		return 1
	})

	ret, err := strconv.Unquote(s)
	if err != nil {
		panic(ParseError{err})
	} else {
		return String(ret)
	}
}

func (pa *Parser) parseValue() Value {
	pa.pr.skipSpace()

	r := pa.pr.mustPeek()
	switch r {
	case 0:
		panic(IllegalCharErr)
	case '"':
		return pa.parseString()
	default:
		return pa.parseSimple()
	}

	//panic(ParseError{fmt.Errorf("unexpected character: %v", r)})

}

func (pa *Parser) parseList() *Cons {
	var p, q *Cons
	for !pa.pr.isEof() {
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
	defer func() {
		if r := recover(); r != nil {
			if r, ok := r.(ParseError); ok {
				err = eu.Append(err, r)
			} else {
				panic(r)
			}
		}
	}()

	var body *Cons
	if !pa.pr.isEof() {
		body = pa.parseList()
		if !pa.pr.isEof() {
			err = eu.Append(err, UnexpectedCharErr)
		}
	}

	return &Cons{
		Atom("begin"),
		body,
	}, nil
}
