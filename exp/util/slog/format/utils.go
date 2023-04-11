// Copyright 2022 The Go Authors. All rights reserved. Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.
package format

import (
	"time"

	"github.com/wrmsr/bane/exp/util/slog/buffer"
)

// This takes half the time of Time.AppendFormat.
func WriteTimeRFC3339Millis(buf *buffer.Buffer, t time.Time) {
	year, month, day := t.Date()
	buf.WritePosIntWidth(year, 4)
	buf.WriteByte('-')
	buf.WritePosIntWidth(int(month), 2)
	buf.WriteByte('-')
	buf.WritePosIntWidth(day, 2)
	buf.WriteByte('T')
	hour, min, sec := t.Clock()
	buf.WritePosIntWidth(hour, 2)
	buf.WriteByte(':')
	buf.WritePosIntWidth(min, 2)
	buf.WriteByte(':')
	buf.WritePosIntWidth(sec, 2)
	ns := t.Nanosecond()
	buf.WriteByte('.')
	buf.WritePosIntWidth(ns/1e6, 3)
	_, offsetSeconds := t.Zone()
	if offsetSeconds == 0 {
		buf.WriteByte('Z')
	} else {
		offsetMinutes := offsetSeconds / 60
		if offsetMinutes < 0 {
			buf.WriteByte('-')
			offsetMinutes = -offsetMinutes
		} else {
			buf.WriteByte('+')
		}
		buf.WritePosIntWidth(offsetMinutes/60, 2)
		buf.WriteByte(':')
		buf.WritePosIntWidth(offsetMinutes%60, 2)
	}
}
