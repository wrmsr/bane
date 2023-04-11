// Copyright 2022 The Go Authors. All rights reserved. Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.
package slog

import (
	"fmt"
	"io"
	"runtime"
	"strconv"
	"sync"
	"time"

	"golang.org/x/exp/slices"

	"github.com/wrmsr/bane/exp/util/slog/buffer"
)

type CommonOpts struct {
	AddSource bool

	Level Leveler

	ReplaceAttr func(groups []string, a Attr) Attr
}

type commonHandler struct {
	opts CommonOpts
	w    io.Writer

	mtx sync.Mutex

	preformattedAttrs []byte
	groupPrefix       string
	groups            []string
	nOpenGroups       int
}

func (h *commonHandler) clone() *commonHandler {
	return &commonHandler{
		w:    h.w,
		opts: h.opts,

		preformattedAttrs: slices.Clip(h.preformattedAttrs),
		groupPrefix:       h.groupPrefix,
		groups:            slices.Clip(h.groups),
		nOpenGroups:       h.nOpenGroups,
	}
}

func (h *commonHandler) enabled(l Level) bool {
	minLevel := LevelInfo
	if h.opts.Level != nil {
		minLevel = h.opts.Level.Level()
	}
	return l >= minLevel
}

func (h *commonHandler) withAttrs(as []Attr) *commonHandler {
	h2 := h.clone()

	// Pre-format the attributes as an optimization.
	prefix := buffer.New()
	defer prefix.Free()
	prefix.WriteString(h.groupPrefix)

	state := h2.newHandleState((*buffer.Buffer)(&h2.preformattedAttrs), false, "", prefix)
	defer state.free()
	if len(h2.preformattedAttrs) > 0 {
		state.sep = h.attrSep()
	}
	state.openGroups()
	for _, a := range as {
		state.appendAttr(a)
	}

	// Remember the new prefix for later keys.
	h2.groupPrefix = state.prefix.String()

	// Remember how many opened groups are in preformattedAttrs, so we don't open them again when we handle a Record.
	h2.nOpenGroups = len(h2.groups)

	return h2
}

func (h *commonHandler) withGroup(name string) *commonHandler {
	if name == "" {
		return h
	}
	h2 := h.clone()
	h2.groups = append(h2.groups, name)
	return h2
}

func RecordFrame(r Record) runtime.Frame {
	fs := runtime.CallersFrames([]uintptr{r.PC})
	f, _ := fs.Next()
	return f
}

func (h *commonHandler) handle(r Record) error {
	state := h.newHandleState(buffer.New(), true, "", nil)
	defer state.free()

	if h.json {
		state.buf.WriteByte('{')
	}

	// Built-in attributes. They are not in a group.
	stateGroups := state.groups
	state.groups = nil // So ReplaceAttrs sees no groups instead of the pre groups.
	rep := h.opts.ReplaceAttr

	// time
	if !r.Time.IsZero() {
		key := TimeKey
		val := r.Time.Round(0) // strip monotonic to match Attr behavior
		if rep == nil {
			state.appendKey(key)
			state.appendTime(val)
		} else {
			state.appendAttr(Time(key, val))
		}
	}

	// level
	key := LevelKey
	val := r.Level
	if rep == nil {
		state.appendKey(key)
		state.appendString(val.String())
	} else {
		state.appendAttr(Any(key, val))
	}

	// source
	if h.opts.AddSource {
		frame := RecordFrame(r)
		if frame.File != "" {
			key := SourceKey
			if rep == nil {
				state.appendKey(key)
				state.appendSource(frame.File, frame.Line)
			} else {
				buf := buffer.New()
				buf.WriteString(frame.File) // TODO: escape?
				buf.WriteByte(':')
				buf.WritePosInt(frame.Line)
				s := buf.String()
				buf.Free()
				state.appendAttr(String(key, s))
			}
		}
	}

	key = MessageKey
	msg := r.Message
	if rep == nil {
		state.appendKey(key)
		state.appendString(msg)
	} else {
		state.appendAttr(String(key, msg))
	}

	state.groups = stateGroups // Restore groups passed to ReplaceAttrs.
	state.appendNonBuiltIns(r)
	state.buf.WriteByte('\n')

	h.mtx.Lock()
	defer h.mtx.Unlock()
	_, err := h.w.Write(*state.buf)
	return err
}

func (s *handleState) appendNonBuiltIns(r Record) {
	// preformatted Attrs
	if len(s.h.preformattedAttrs) > 0 {
		s.buf.WriteString(s.sep)
		s.buf.Write(s.h.preformattedAttrs)
		s.sep = s.h.attrSep()
	}
	// Attrs in Record -- unlike the built-in ones, they are in groups started from WithGroup.
	s.prefix = buffer.New()
	defer s.prefix.Free()
	s.prefix.WriteString(s.h.groupPrefix)
	s.openGroups()
	r.Attrs(func(a Attr) {
		s.appendAttr(a)
	})
	if s.h.json {
		// Close all open groups.
		for range s.h.groups {
			s.buf.WriteByte('}')
		}
		// Close the top-level object.
		s.buf.WriteByte('}')
	}
}

// attrSep returns the separator between attributes.
func (h *commonHandler) attrSep() string {
	if h.json {
		return ","
	}
	return " "
}

// handleState holds state for a single call to commonHandler.handle. The initial value of sep determines whether to
// emit a separator before the next key, after which it stays true.
type handleState struct {
	h       *commonHandler
	buf     *buffer.Buffer
	freeBuf bool           // should buf be freed?
	sep     string         // separator to write before next key
	prefix  *buffer.Buffer // for text: key prefix
	groups  *[]string      // pool-allocated slice of active groups, for ReplaceAttr
}

var groupPool = sync.Pool{New: func() any {
	s := make([]string, 0, 10)
	return &s
}}

func (h *commonHandler) newHandleState(buf *buffer.Buffer, freeBuf bool, sep string, prefix *buffer.Buffer) handleState {
	s := handleState{
		h:       h,
		buf:     buf,
		freeBuf: freeBuf,
		sep:     sep,
		prefix:  prefix,
	}
	if h.opts.ReplaceAttr != nil {
		s.groups = groupPool.Get().(*[]string)
		*s.groups = append(*s.groups, h.groups[:h.nOpenGroups]...)
	}
	return s
}

func (s *handleState) free() {
	if s.freeBuf {
		s.buf.Free()
	}
	if gs := s.groups; gs != nil {
		*gs = (*gs)[:0]
		groupPool.Put(gs)
	}
}

func (s *handleState) openGroups() {
	for _, n := range s.h.groups[s.h.nOpenGroups:] {
		s.openGroup(n)
	}
}

// Separator for group names and keys.
const keyComponentSep = '.'

// openGroup starts a new group of attributes
// with the given name.
func (s *handleState) openGroup(name string) {
	if s.h.json {
		s.appendKey(name)
		s.buf.WriteByte('{')
		s.sep = ""
	} else {
		s.prefix.WriteString(name)
		s.prefix.WriteByte(keyComponentSep)
	}
	// Collect group names for ReplaceAttr.
	if s.groups != nil {
		*s.groups = append(*s.groups, name)
	}

}

// closeGroup ends the group with the given name.
func (s *handleState) closeGroup(name string) {
	if s.h.json {
		s.buf.WriteByte('}')
	} else {
		(*s.prefix) = (*s.prefix)[:len(*s.prefix)-len(name)-1 /* for keyComponentSep */]
	}
	s.sep = s.h.attrSep()
	if s.groups != nil {
		*s.groups = (*s.groups)[:len(*s.groups)-1]
	}
}

// appendAttr appends the Attr's key and value using app.
// It handles replacement and checking for an empty key.
// after replacement).
func (s *handleState) appendAttr(a Attr) {
	v := a.Value

	// Elide a non-group with an empty key.
	if a.Key == "" && v.Kind() != KindGroup {
		return
	}

	if rep := s.h.opts.ReplaceAttr; rep != nil && v.Kind() != KindGroup {
		var gs []string
		if s.groups != nil {
			gs = *s.groups
		}
		a = rep(gs, Attr{a.Key, v})
		if a.Key == "" {
			return
		}
		// Although all attributes in the Record are already resolved,
		// This one came from the user, so it may not have been.
		v = a.Value.Resolve()
	}

	if v.Kind() == KindGroup {
		attrs := v.Group()
		// Output only non-empty groups.
		if len(attrs) > 0 {
			// Inline a group with an empty key.
			if a.Key != "" {
				s.openGroup(a.Key)
			}
			for _, aa := range attrs {
				s.appendAttr(aa)
			}
			if a.Key != "" {
				s.closeGroup(a.Key)
			}
		}
	} else {
		s.appendKey(a.Key)
		s.appendValue(v)
	}
}

func (s *handleState) appendError(err error) {
	s.appendString(fmt.Sprintf("!ERROR:%v", err))
}

func (s *handleState) appendKey(key string) {
	s.buf.WriteString(s.sep)
	if s.prefix != nil {
		// TODO: optimize by avoiding allocation.
		s.appendString(string(*s.prefix) + key)
	} else {
		s.appendString(key)
	}
	if s.h.json {
		s.buf.WriteByte(':')
	} else {
		s.buf.WriteByte('=')
	}
	s.sep = s.h.attrSep()
}

func (s *handleState) appendSource(file string, line int) {
	if s.h.json {
		s.buf.WriteByte('"')
		*s.buf = appendEscapedJSONString(*s.buf, file)
		s.buf.WriteByte(':')
		s.buf.WritePosInt(line)
		s.buf.WriteByte('"')
	} else {
		// text
		if needsQuoting(file) {
			s.appendString(file + ":" + strconv.Itoa(line))
		} else {
			// common case: no quoting needed.
			s.appendString(file)
			s.buf.WriteByte(':')
			s.buf.WritePosInt(line)
		}
	}
}

func (s *handleState) appendString(str string) {
	if s.h.json {
		s.buf.WriteByte('"')
		*s.buf = appendEscapedJSONString(*s.buf, str)
		s.buf.WriteByte('"')
	} else {
		// text
		if needsQuoting(str) {
			*s.buf = strconv.AppendQuote(*s.buf, str)
		} else {
			s.buf.WriteString(str)
		}
	}
}

func (s *handleState) appendValue(v Value) {
	var err error
	if s.h.json {
		err = appendJSONValue(s, v)
	} else {
		err = appendTextValue(s, v)
	}
	if err != nil {
		s.appendError(err)
	}
}

func (s *handleState) appendTime(t time.Time) {
	if s.h.json {
		appendJSONTime(s, t)
	} else {
		writeTimeRFC3339Millis(s.buf, t)
	}
}

// This takes half the time of Time.AppendFormat.
func writeTimeRFC3339Millis(buf *buffer.Buffer, t time.Time) {
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
