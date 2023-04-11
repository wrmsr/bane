// Copyright 2022 The Go Authors. All rights reserved. Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.
package format

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"

	ju "github.com/wrmsr/bane/pkg/util/json"
	"github.com/wrmsr/bane/pkg/util/slog"
	"github.com/wrmsr/bane/pkg/util/slog/buffer"
)

type jsonFormatterStyle struct{}

func NewJsonFormatter(opts FormatterOpts) slog.Formatter {
	return &formatter{
		opts: opts,
		sty:  jsonFormatterStyle{},
	}
}

var _ Style = jsonFormatterStyle{}

func (sty jsonFormatterStyle) attrSep() string  { return "," }
func (sty jsonFormatterStyle) valueSep() string { return "=" }

func (sty jsonFormatterStyle) begin(s *state) {
	s.buf.WriteB('{')
}

func (sty jsonFormatterStyle) end(s *state) {
	// Close all open groups.
	for range s.f.groups {
		s.buf.WriteB('}')
	}
	// Close the top-level object.
	s.buf.WriteB('}')
}

func (sty jsonFormatterStyle) openGroup(s *state, name string) {
	s.appendKey(name)
	s.buf.WriteB('{')
	s.sep = ""
}

func (sty jsonFormatterStyle) closeGroup(s *state, name string) {
	s.buf.WriteB('}')
}

func (sty jsonFormatterStyle) appendSource(s *state, file string, line int) {
	s.buf.WriteB('"')
	*s.buf = ju.AppendEncodedString(*s.buf, file, false)
	s.buf.WriteB(':')
	s.buf.WritePosInt(line)
	s.buf.WriteB('"')
}

func (sty jsonFormatterStyle) appendString(s *state, str string) {
	s.buf.WriteB('"')
	*s.buf = ju.AppendEncodedString(*s.buf, str, false)
	s.buf.WriteB('"')
}

func (sty jsonFormatterStyle) appendValue(s *state, v slog.Value) {
	var err error
	err = appendJSONValue(s, v)
	if err != nil {
		s.appendError(err)
	}
}

func (sty jsonFormatterStyle) appendTime(s *state, t time.Time) {
	appendJSONTime(s, t)
}

func appendJSONTime(s *state, t time.Time) {
	if y := t.Year(); y < 0 || y >= 10000 {
		s.appendError(errors.New("time.Time year outside of range [0,9999]"))
	}
	s.buf.WriteB('"')
	*s.buf = t.AppendFormat(*s.buf, time.RFC3339Nano)
	s.buf.WriteB('"')
}

func appendJSONValue(s *state, v slog.Value) error {
	switch v.Kind() {
	case slog.KindString:
		s.appendString(v.String())
	case slog.KindInt64:
		*s.buf = strconv.AppendInt(*s.buf, v.Int64(), 10)
	case slog.KindUint64:
		*s.buf = strconv.AppendUint(*s.buf, v.Uint64(), 10)
	case slog.KindFloat64:
		f := v.Float64()
		// json.Marshal fails on special floats, so handle them here.
		switch {
		case math.IsInf(f, 1):
			s.buf.WriteS(`"+Inf"`)
		case math.IsInf(f, -1):
			s.buf.WriteS(`"-Inf"`)
		case math.IsNaN(f):
			s.buf.WriteS(`"NaN"`)
		default:
			// json.Marshal is funny about floats; it doesn't always match strconv.AppendFloat. So just call it. That's
			// expensive, but floats are rare.
			if err := appendJSONMarshal(s.buf, f); err != nil {
				return err
			}
		}
	case slog.KindBool:
		*s.buf = strconv.AppendBool(*s.buf, v.Bool())
	case slog.KindDuration:
		// Do what json.Marshal does.
		*s.buf = strconv.AppendInt(*s.buf, int64(v.Duration()), 10)
	case slog.KindTime:
		s.appendTime(v.Time())
	case slog.KindAny:
		a := v.Any()
		if err, ok := a.(error); ok {
			s.appendString(err.Error())
		} else {
			return appendJSONMarshal(s.buf, a)
		}
	default:
		panic(fmt.Sprintf("bad kind: %d", v.Kind()))
	}
	return nil
}

func appendJSONMarshal(buf *buffer.Buffer, v any) error {
	var bb bytes.Buffer
	enc := json.NewEncoder(&bb)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(v); err != nil {
		return err
	}
	bs := bb.Bytes()
	_, _ = buf.Write(bs[:len(bs)-1])
	return nil
}
