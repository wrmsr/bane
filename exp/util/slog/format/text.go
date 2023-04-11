// Copyright 2022 The Go Authors. All rights reserved. Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.
package format

import (
	"encoding"
	"fmt"
	"reflect"
	"strconv"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/wrmsr/bane/exp/util/slog"
)

//

type textFormatterStyle struct{}

var _ formatterStyle = textFormatterStyle{}

func (i textFormatterStyle) attrSep() string  { return " " }
func (i textFormatterStyle) valueSep() string { return ":" }

func (i textFormatterStyle) begin(s *formatState) {}

func (i textFormatterStyle) end(s *formatState) {}

func (i textFormatterStyle) openGroup(s *formatState, name string) {
	s.prefix.WriteString(name)
	s.prefix.WriteByte(keyComponentSep)
}

func (i textFormatterStyle) closeGroup(s *formatState, name string) {
	*s.prefix = (*s.prefix)[:len(*s.prefix)-len(name)-1 /* for keyComponentSep */]
}

func (i textFormatterStyle) appendSource(s *formatState, file string, line int) {
	if needsQuoting(file) {
		s.appendString(file + ":" + strconv.Itoa(line))
	} else {
		// common case: no quoting needed.
		s.appendString(file)
		s.buf.WriteByte(':')
		s.buf.WritePosInt(line)
	}
}

func (i textFormatterStyle) appendString(s *formatState, str string) {
	if needsQuoting(str) {
		*s.buf = strconv.AppendQuote(*s.buf, str)
	} else {
		s.buf.WriteString(str)
	}
}

func (i textFormatterStyle) appendValue(s *formatState, v slog.Value) {
	var err error
	err = appendTextValue(s, v)
	if err != nil {
		s.appendError(err)
	}
}

func (i textFormatterStyle) appendTime(s *formatState, t time.Time) {
	WriteTimeRFC3339Millis(s.buf, t)
}

//

func appendTextValue(s *formatState, v slog.Value) error {
	switch v.Kind() {
	case slog.KindString:
		s.appendString(v.String())
	case slog.KindTime:
		s.appendTime(v.Time())
	case slog.KindAny:
		a := v.Any()
		if tm, ok := a.(encoding.TextMarshaler); ok {
			data, err := tm.MarshalText()
			if err != nil {
				return err
			}
			// TODO: avoid the conversion to string.
			s.appendString(string(data))
			return nil
		}
		if bs, ok := byteSlice(a); ok {
			// As of Go 1.19, this only allocates for strings longer than 32 bytes.
			s.buf.WriteString(strconv.Quote(string(bs)))
			return nil
		}
		s.appendString(fmt.Sprintf("%+v", v.Any()))
	default:
		*s.buf = appendValue(*s.buf, v)
	}
	return nil
}

func appendValue(dst []byte, v slog.Value) []byte {
	switch v.Kind() {
	case slog.KindString:
		return append(dst, v.String()...)
	case slog.KindInt64:
		return strconv.AppendInt(dst, v.Int64(), 10)
	case slog.KindUint64:
		return strconv.AppendUint(dst, v.Uint64(), 10)
	case slog.KindFloat64:
		return strconv.AppendFloat(dst, v.Float64(), 'g', -1, 64)
	case slog.KindBool:
		return strconv.AppendBool(dst, v.Bool())
	case slog.KindDuration:
		return append(dst, v.Duration().String()...)
	case slog.KindTime:
		return append(dst, v.Time().String()...)
	case slog.KindAny, slog.KindGroup, slog.KindLogValuer:
		return append(dst, fmt.Sprint(v.Any())...)
	default:
		panic(fmt.Sprintf("bad kind: %s", v.Kind()))
	}
}

// byteSlice returns its argument as a []byte if the argument's underlying type is []byte, along with a second return
// value of true. Otherwise it returns nil, false.
func byteSlice(a any) ([]byte, bool) {
	if bs, ok := a.([]byte); ok {
		return bs, true
	}
	// Like Printf's %s, we allow both the slice type and the byte element type to be named.
	t := reflect.TypeOf(a)
	if t != nil && t.Kind() == reflect.Slice && t.Elem().Kind() == reflect.Uint8 {
		return reflect.ValueOf(a).Bytes(), true
	}
	return nil, false
}

//

func needsQuoting(s string) bool {
	for i := 0; i < len(s); {
		b := s[i]
		if b < utf8.RuneSelf {
			if needsQuotingSet[b] {
				return true
			}
			i++
			continue
		}
		r, size := utf8.DecodeRuneInString(s[i:])
		if r == utf8.RuneError || unicode.IsSpace(r) || !unicode.IsPrint(r) {
			return true
		}
		i += size
	}
	return false
}

var needsQuotingSet = [utf8.RuneSelf]bool{
	'"': true,
	'=': true,
}

func init() {
	for i := 0; i < utf8.RuneSelf; i++ {
		r := rune(i)
		if unicode.IsSpace(r) || !unicode.IsPrint(r) {
			needsQuotingSet[i] = true
		}
	}
}
