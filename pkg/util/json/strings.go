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
