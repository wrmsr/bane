/*
Copyright 2021 The ANTLR Project

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the
following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following
   disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following
   disclaimer in the documentation and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products
   derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES,
INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
// Copyright (c) 2012-2017 The ANTLR Project. All rights reserved. Use of this file is governed by the BSD 3-clause
// license that can be found in the LICENSE.txt file in the project root.

package antlr

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type DelimitingSplitter struct {
	input                  antlr.CharStream
	interpreter            antlr.ILexerATNSimulator
	virt                   antlr.Lexer // The most derived lexer implementation. Allows virtual method calls.
	tokenFactorySourcePair *antlr.TokenSourceCharStreamPair

	delimToken any
	delimiters [][]int

	factory antlr.TokenFactory

	token               antlr.Token
	channel             int
	tokenStartCharIndex int
	tokenStartLine      int
	tokenStartColumn    int
	text                string
	thetype             int
	mode                int

	hitEOF bool
	noSkip bool
}

func NewDelimitingSplitter(
	input antlr.CharStream,
	interpreter antlr.ILexerATNSimulator,
	virt antlr.Lexer,
	tokenFactorySourcePair *antlr.TokenSourceCharStreamPair,
) *DelimitingSplitter {
	return &DelimitingSplitter{
		input:                  input,
		interpreter:            interpreter,
		virt:                   virt,
		tokenFactorySourcePair: tokenFactorySourcePair,

		factory:             antlr.CommonTokenFactoryDEFAULT,
		tokenStartCharIndex: -1,
		tokenStartLine:      -1,
		tokenStartColumn:    -1,
		channel:             antlr.TokenDefaultChannel,
		thetype:             antlr.TokenInvalidType,
		mode:                antlr.LexerDefaultMode,
	}
}

func (l *DelimitingSplitter) NextToken() antlr.Token {
	tokenStartMarker := l.input.Mark()
	defer l.input.Release(tokenStartMarker)

	for {
		if l.hitEOF {
			l.emitEOF()
			return l.token
		}

		l.token = nil
		l.channel = antlr.TokenDefaultChannel
		l.tokenStartCharIndex = l.input.Index()
		l.tokenStartColumn = l.interpreter.GetCharPositionInLine()
		l.tokenStartLine = l.interpreter.GetLine()
		l.text = ""

		continueOuter := false
		for {
			l.thetype = antlr.TokenInvalidType
			ttype := antlr.LexerSkip

			ttype = l.safeMatch()

			/*
				for delimiter in l.delimiters:
					if l.match_delimiter(delimiter):
						ttype = l.delimiter_token
						break
				else:
					try:
						ttype = l.interp.match(input, l.mode)
					except LexerNoViableAltException as e:
						self.notifyListeners(e)  # report error
						self.recover(e)
			*/

			if l.input.LA(1) == antlr.TokenEOF {
				l.hitEOF = true
			}
			if l.thetype == antlr.TokenInvalidType {
				l.thetype = ttype
			}
			if l.thetype == antlr.LexerSkip {
				continueOuter = true
				break
			}
			if l.thetype != antlr.LexerMore {
				break
			}
		}

		if continueOuter {
			continue
		}
		if l.token == nil {
			l.virt.Emit()
		}
		return l.token
	}
}

func (l *DelimitingSplitter) notifyListeners(e antlr.RecognitionException) {
	start := l.tokenStartCharIndex
	stop := l.input.Index()
	text := l.input.GetTextFromInterval(antlr.NewInterval(start, stop))
	msg := "token recognition error at: '" + text + "'"
	//listener := l.getErrorListenerDispatch()
	//listener.SyntaxError(l, nil, l.TokenStartLine, l.TokenStartColumn, msg, e)
	_ = msg
}

func (l *DelimitingSplitter) safeMatch() (ret int) {
	defer func() {
		if e := recover(); e != nil {
			if re, ok := e.(antlr.RecognitionException); ok {
				l.notifyListeners(re) // Report error
				l.Recover(re)
				ret = antlr.LexerSkip // default
			}
		}
	}()

	return l.interpreter.Match(l.input, l.mode)
}

func (l *DelimitingSplitter) Recover(re antlr.RecognitionException) {
	if l.input.LA(1) != antlr.TokenEOF {
		if _, ok := re.(*antlr.LexerNoViableAltException); ok {
			// Skip a char and try again
			l.interpreter.Consume(l.input)
		} else {
			// TODO: Do we lose character or line position information?
			l.input.Consume()
		}
	}
}
func (l *DelimitingSplitter) emitEOF() antlr.Token {
	cpos := l.interpreter.GetCharPositionInLine()
	lpos := l.interpreter.GetLine()
	eof := l.factory.Create(l.tokenFactorySourcePair, antlr.TokenEOF, "", antlr.TokenDefaultChannel, l.input.Index(), l.input.Index()-1, lpos, cpos)
	l.token = eof
	return eof
}

func (l *DelimitingSplitter) matchDelimiter(delimiter []int) bool {
	for i, c := range delimiter {
		if l.input.LA(i+1) != c {
			return false
		}
	}
	l.input.Seek(l.input.Index() + len(delimiter))
	return true
}

/*
def split(self) -> ta.Tuple[ta.List[ta.Tuple[str, str]], str]:
	lst = []
	sb = io.StringIO()
	while True:
		token = self.nextToken()
		if token.type == antlr4.Token.EOF:
			break
		if token.type == l.delimiter_token:
			statement = sb.getvalue().strip()
			if statement:
				lst.append((statement, token.text))
			sb = io.StringIO()
		else:
			sb.write(token.text)
	partial = sb.getvalue()
	return lst, partial
*/
