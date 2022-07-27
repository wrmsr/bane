/*
Copyright (c) 2009 The Go Authors. All rights reserved.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the
following conditions are met:

* Redistributions of source code must retain the above copyright notice, this list of conditions and the following
  disclaimer.
* Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following
  disclaimer in the documentation and/or other materials provided with the distribution.
* Neither the name of Google Inc. nor the names of its contributors may be used to endorse or promote products derived
  from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES,
INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
package json

import (
	"io"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"
)

func EncodeString(w io.Writer, s string, escapeHTML bool) error {
	if _, err := w.Write(_quoteBytes); err != nil {
		return err
	}

	start := 0
	for i := 0; i < len(s); {
		if b := s[i]; b < utf8.RuneSelf {
			if htmlSafeSet[b] || (!escapeHTML && safeSet[b]) {
				i++
				continue
			}
			if start < i {
				if _, err := w.Write([]byte(s[start:i])); err != nil {
					return err
				}
			}
			if _, err := w.Write(_backslashBytes); err != nil {
				return err
			}

			switch b {
			case '\\', '"':
				if _, err := w.Write([]byte{b}); err != nil {
					return err
				}
			case '\n':
				if _, err := w.Write([]byte{'n'}); err != nil {
					return err
				}
			case '\r':
				if _, err := w.Write([]byte{'r'}); err != nil {
					return err
				}
			case '\t':
				if _, err := w.Write([]byte{'t'}); err != nil {
					return err
				}
			default:
				// This encodes bytes < 0x20 except for \t, \n and \r. If escapeHTML is set, it also escapes <, >, and
				// & because they can lead to security holes when user-controlled strings are rendered into JSON and
				// served to some browsers.
				if _, err := w.Write([]byte(`u00`)); err != nil {
					return err
				}
				if _, err := w.Write([]byte{hex[b>>4], hex[b&0xF]}); err != nil {
					return err
				}
			}

			i++
			start = i
			continue
		}

		c, size := utf8.DecodeRuneInString(s[i:])
		if c == utf8.RuneError && size == 1 {
			if start < i {
				if _, err := w.Write([]byte(s[start:i])); err != nil {
					return err
				}
			}
			if _, err := w.Write([]byte(`\ufffd`)); err != nil {
				return err
			}

			i += size
			start = i
			continue
		}

		// U+2028 is LINE SEPARATOR.
		// U+2029 is PARAGRAPH SEPARATOR.
		// They are both technically valid characters in JSON strings, but don't work in JSONP, which has to be
		// evaluated as JavaScript, and can lead to security holes there. It is valid JSON to escape them, so we do so
		// unconditionally. See http://timelessrepo.com/json-isnt-a-javascript-subset for discussion.
		if c == '\u2028' || c == '\u2029' {
			if start < i {
				if _, err := w.Write([]byte(s[start:i])); err != nil {
					return err
				}
			}
			if _, err := w.Write([]byte(`\u202`)); err != nil {
				return err
			}
			if _, err := w.Write([]byte{hex[c&0xF]}); err != nil {
				return err
			}

			i += size
			start = i
			continue
		}

		i += size
	}

	if start < len(s) {
		if _, err := w.Write([]byte(s[start:])); err != nil {
			return err
		}
	}

	if _, err := w.Write(_quoteBytes); err != nil {
		return err
	}
	return nil
}

// NOTE: keep in sync with string above.
func EncodeStringBytes(w io.Writer, s []byte, escapeHTML bool) error {
	if _, err := w.Write(_quoteBytes); err != nil {
		return err
	}

	start := 0
	for i := 0; i < len(s); {
		if b := s[i]; b < utf8.RuneSelf {
			if htmlSafeSet[b] || (!escapeHTML && safeSet[b]) {
				i++
				continue
			}
			if start < i {
				if _, err := w.Write(s[start:i]); err != nil {
					return err
				}
			}
			if _, err := w.Write(_backslashBytes); err != nil {
				return err
			}

			switch b {
			case '\\', '"':
				if _, err := w.Write([]byte{b}); err != nil {
					return err
				}
			case '\n':
				if _, err := w.Write([]byte{'n'}); err != nil {
					return err
				}
			case '\r':
				if _, err := w.Write([]byte{'r'}); err != nil {
					return err
				}
			case '\t':
				if _, err := w.Write([]byte{'t'}); err != nil {
					return err
				}
			default:
				// This encodes bytes < 0x20 except for \t, \n and \r. If escapeHTML is set, it also escapes <, >, and &
				// because they can lead to security holes when user-controlled strings are rendered into JSON and
				// served to some browsers.
				if _, err := w.Write([]byte(`u00`)); err != nil {
					return err
				}
				if _, err := w.Write([]byte{hex[b>>4], hex[b&0xF]}); err != nil {
					return err
				}
			}

			i++
			start = i
			continue
		}

		c, size := utf8.DecodeRune(s[i:])
		if c == utf8.RuneError && size == 1 {
			if start < i {
				if _, err := w.Write(s[start:i]); err != nil {
					return err
				}
			}
			if _, err := w.Write([]byte(`\ufffd`)); err != nil {
				return err
			}
			i += size
			start = i
			continue
		}

		// U+2028 is LINE SEPARATOR.
		// U+2029 is PARAGRAPH SEPARATOR.
		// They are both technically valid characters in JSON strings, but don't work in JSONP, which has to be
		// evaluated as JavaScript, and can lead to security holes there. It is valid JSON to escape them, so we do so
		// unconditionally. See http://timelessrepo.com/json-isnt-a-javascript-subset for discussion.
		if c == '\u2028' || c == '\u2029' {
			if start < i {
				if _, err := w.Write(s[start:i]); err != nil {
					return err
				}
			}
			if _, err := w.Write([]byte(`\u202`)); err != nil {
				return err
			}
			if _, err := w.Write([]byte{hex[c&0xF]}); err != nil {
				return err
			}
			i += size
			start = i
			continue
		}

		i += size
	}

	if start < len(s) {
		if _, err := w.Write(s[start:]); err != nil {
			return err
		}
	}

	if _, err := w.Write(_quoteBytes); err != nil {
		return err
	}
	return nil
}

// getu4 decodes \uXXXX from the beginning of s, returning the hex value,
// or it returns -1.
func getu4(s []byte) rune {
	if len(s) < 6 || s[0] != '\\' || s[1] != 'u' {
		return -1
	}
	var r rune
	for _, c := range s[2:6] {
		switch {
		case '0' <= c && c <= '9':
			c = c - '0'
		case 'a' <= c && c <= 'f':
			c = c - 'a' + 10
		case 'A' <= c && c <= 'F':
			c = c - 'A' + 10
		default:
			return -1
		}
		r = r*16 + rune(c)
	}
	return r
}

// unquote converts a quoted JSON string literal s into an actual string t.
// The rules are different than for Go, so cannot use strconv.Unquote.
func Unquote(s []byte) (t string, ok bool) {
	s, ok = UnquoteBytes(s)
	t = string(s)
	return
}

func UnquoteBytes(s []byte) (t []byte, ok bool) {
	if len(s) < 2 || s[0] != '"' || s[len(s)-1] != '"' {
		return
	}
	s = s[1 : len(s)-1]

	// Check for unusual characters. If there are none,
	// then no unquoting is needed, so return a slice of the
	// original bytes.
	r := 0
	for r < len(s) {
		c := s[r]
		if c == '\\' || c == '"' || c < ' ' {
			break
		}
		if c < utf8.RuneSelf {
			r++
			continue
		}
		rr, size := utf8.DecodeRune(s[r:])
		if rr == utf8.RuneError && size == 1 {
			break
		}
		r += size
	}
	if r == len(s) {
		return s, true
	}

	b := make([]byte, len(s)+2*utf8.UTFMax)
	w := copy(b, s[0:r])
	for r < len(s) {
		// Out of room? Can only happen if s is full of
		// malformed UTF-8 and we're replacing each
		// byte with RuneError.
		if w >= len(b)-2*utf8.UTFMax {
			nb := make([]byte, (len(b)+utf8.UTFMax)*2)
			copy(nb, b[0:w])
			b = nb
		}
		switch c := s[r]; {
		case c == '\\':
			r++
			if r >= len(s) {
				return
			}
			switch s[r] {
			default:
				return
			case '"', '\\', '/', '\'':
				b[w] = s[r]
				r++
				w++
			case 'b':
				b[w] = '\b'
				r++
				w++
			case 'f':
				b[w] = '\f'
				r++
				w++
			case 'n':
				b[w] = '\n'
				r++
				w++
			case 'r':
				b[w] = '\r'
				r++
				w++
			case 't':
				b[w] = '\t'
				r++
				w++
			case 'u':
				r--
				rr := getu4(s[r:])
				if rr < 0 {
					return
				}
				r += 6
				if utf16.IsSurrogate(rr) {
					rr1 := getu4(s[r:])
					if dec := utf16.DecodeRune(rr, rr1); dec != unicode.ReplacementChar {
						// A valid pair; consume.
						r += 6
						w += utf8.EncodeRune(b[w:], dec)
						break
					}
					// Invalid surrogate; fall back to replacement rune.
					rr = unicode.ReplacementChar
				}
				w += utf8.EncodeRune(b[w:], rr)
			}

		// Quote, control characters are invalid.
		case c == '"', c < ' ':
			return

		// ASCII
		case c < utf8.RuneSelf:
			b[w] = c
			r++
			w++

		// Coerce to well-formed UTF-8.
		default:
			rr, size := utf8.DecodeRune(s[r:])
			r += size
			w += utf8.EncodeRune(b[w:], rr)
		}
	}
	return b[0:w], true
}
