// Copyright 2022 The Go Authors. All rights reserved. Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.
package format

import (
	"time"

	"github.com/wrmsr/bane/pkg/util/slog/buffer"
)

// This takes half the time of Time.AppendFormat.
func WriteTimeRFC3339Millis(buf *buffer.Buffer, t time.Time) {
	year, month, day := t.Date()
	buf.WritePosIntWidth(year, 4)
	buf.WriteB('-')
	buf.WritePosIntWidth(int(month), 2)
	buf.WriteB('-')
	buf.WritePosIntWidth(day, 2)
	buf.WriteB('T')
	hour, min, sec := t.Clock()
	buf.WritePosIntWidth(hour, 2)
	buf.WriteB(':')
	buf.WritePosIntWidth(min, 2)
	buf.WriteB(':')
	buf.WritePosIntWidth(sec, 2)
	ns := t.Nanosecond()
	buf.WriteB('.')
	buf.WritePosIntWidth(ns/1e6, 3)
	_, offsetSeconds := t.Zone()
	if offsetSeconds == 0 {
		buf.WriteB('Z')
	} else {
		offsetMinutes := offsetSeconds / 60
		if offsetMinutes < 0 {
			buf.WriteB('-')
			offsetMinutes = -offsetMinutes
		} else {
			buf.WriteB('+')
		}
		buf.WritePosIntWidth(offsetMinutes/60, 2)
		buf.WriteB(':')
		buf.WritePosIntWidth(offsetMinutes%60, 2)
	}
}
