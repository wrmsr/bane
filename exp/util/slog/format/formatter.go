// Copyright 2022 The Go Authors. All rights reserved. Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.
package format

import (
	"fmt"
	"io"
	"sync"
	"time"

	"golang.org/x/exp/slices"

	"github.com/wrmsr/bane/exp/util/slog"
	"github.com/wrmsr/bane/exp/util/slog/buffer"
)

//

type FormatterOpts struct {
	AddSource bool

	Level slog.Leveler

	ReplaceAttr func(groups []string, a slog.Attr) slog.Attr
}

//

type formatterStyle interface {
	attrSep() string
	valueSep() string

	begin(s *formatState)
	end(s *formatState)

	openGroup(s *formatState, name string)
	closeGroup(s *formatState, name string)

	appendSource(s *formatState, file string, line int)
	appendString(s *formatState, str string)
	appendValue(s *formatState, v slog.Value)
	appendTime(s *formatState, t time.Time)
}

//

type formatter struct {
	opts FormatterOpts
	sty  formatterStyle

	mtx sync.Mutex

	preformattedAttrs []byte
	groupPrefix       string
	groups            []string
	nOpenGroups       int
}

func (f *formatter) clone() *formatter {
	return &formatter{
		opts: f.opts,

		preformattedAttrs: slices.Clip(f.preformattedAttrs),
		groupPrefix:       f.groupPrefix,
		groups:            slices.Clip(f.groups),
		nOpenGroups:       f.nOpenGroups,
	}
}

func (f *formatter) enabled(l slog.Level) bool {
	minLevel := slog.LevelInfo
	if f.opts.Level != nil {
		minLevel = f.opts.Level.Level()
	}
	return l >= minLevel
}

func (f *formatter) withAttrs(as []slog.Attr) *formatter {
	h2 := f.clone()

	// Pre-format the attributes as an optimization.
	prefix := buffer.New()
	defer prefix.Free()
	prefix.WriteString(f.groupPrefix)

	state := h2.newHandleState((*buffer.Buffer)(&h2.preformattedAttrs), false, "", prefix)
	defer state.free()
	if len(h2.preformattedAttrs) > 0 {
		state.sep = f.sty.attrSep()
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

func (f *formatter) withGroup(name string) *formatter {
	if name == "" {
		return f
	}
	h2 := f.clone()
	h2.groups = append(h2.groups, name)
	return h2
}

func (f *formatter) format(r slog.Record, w io.Writer) error {
	state := f.newHandleState(buffer.New(), true, "", nil)
	defer state.free()

	f.sty.begin(&state)

	// Built-in attributes. They are not in a group.
	stateGroups := state.groups
	state.groups = nil // So ReplaceAttrs sees no groups instead of the pre groups.
	rep := f.opts.ReplaceAttr

	// time
	if !r.Time.IsZero() {
		key := slog.TimeKey
		val := r.Time.Round(0) // strip monotonic to match Attr behavior
		if rep == nil {
			state.appendKey(key)
			state.appendTime(val)
		} else {
			state.appendAttr(slog.Time(key, val))
		}
	}

	// level
	key := slog.LevelKey
	val := r.Level
	if rep == nil {
		state.appendKey(key)
		state.appendString(val.String())
	} else {
		state.appendAttr(slog.Any(key, val))
	}

	// source
	if f.opts.AddSource {
		frame := slog.RecordFrame(r)
		if frame.File != "" {
			key := slog.SourceKey
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
				state.appendAttr(slog.String(key, s))
			}
		}
	}

	key = slog.MessageKey
	msg := r.Message
	if rep == nil {
		state.appendKey(key)
		state.appendString(msg)
	} else {
		state.appendAttr(slog.String(key, msg))
	}

	state.groups = stateGroups // Restore groups passed to ReplaceAttrs.
	state.appendNonBuiltIns(r)
	state.buf.WriteByte('\n')

	f.mtx.Lock()
	defer f.mtx.Unlock()
	_, err := w.Write(*state.buf)
	return err
}

//

func (s *formatState) appendNonBuiltIns(r slog.Record) {
	// preformatted Attrs
	if len(s.f.preformattedAttrs) > 0 {
		s.buf.WriteString(s.sep)
		_, _ = s.buf.Write(s.f.preformattedAttrs)
		s.sep = s.f.sty.attrSep()
	}

	// Attrs in Record -- unlike the built-in ones, they are in groups started from WithGroup.
	s.prefix = buffer.New()
	defer s.prefix.Free()
	s.prefix.WriteString(s.f.groupPrefix)
	s.openGroups()
	r.Attrs(func(a slog.Attr) {
		s.appendAttr(a)
	})

	s.f.sty.end(s)
}

// formatState holds state for a single call to commonHandler.handle. The initial value of sep determines whether to
// emit a separator before the next key, after which it stays true.
type formatState struct {
	f       *formatter
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

func (f *formatter) newHandleState(buf *buffer.Buffer, freeBuf bool, sep string, prefix *buffer.Buffer) formatState {
	s := formatState{
		f:       f,
		buf:     buf,
		freeBuf: freeBuf,
		sep:     sep,
		prefix:  prefix,
	}
	if f.opts.ReplaceAttr != nil {
		s.groups = groupPool.Get().(*[]string)
		*s.groups = append(*s.groups, f.groups[:f.nOpenGroups]...)
	}
	return s
}

func (s *formatState) free() {
	if s.freeBuf {
		s.buf.Free()
	}
	if gs := s.groups; gs != nil {
		*gs = (*gs)[:0]
		groupPool.Put(gs)
	}
}

func (s *formatState) openGroups() {
	for _, n := range s.f.groups[s.f.nOpenGroups:] {
		s.openGroup(n)
	}
}

// Separator for group names and keys.
const keyComponentSep = '.'

// openGroup starts a new group of attributes with the given name.
func (s *formatState) openGroup(name string) {
	s.f.sty.openGroup(s, name)
	// Collect group names for ReplaceAttr.
	if s.groups != nil {
		*s.groups = append(*s.groups, name)
	}
}

// closeGroup ends the group with the given name.
func (s *formatState) closeGroup(name string) {
	s.f.sty.closeGroup(s, name)
	s.sep = s.f.sty.attrSep()
	if s.groups != nil {
		*s.groups = (*s.groups)[:len(*s.groups)-1]
	}
}

// appendAttr appends the Attr's key and value using app. It handles replacement and checking for an empty key. after
// replacement).
func (s *formatState) appendAttr(a slog.Attr) {
	v := a.Value

	// Elide a non-group with an empty key.
	if a.Key == "" && v.Kind() != slog.KindGroup {
		return
	}

	if rep := s.f.opts.ReplaceAttr; rep != nil && v.Kind() != slog.KindGroup {
		var gs []string
		if s.groups != nil {
			gs = *s.groups
		}
		a = rep(gs, slog.Attr{a.Key, v})
		if a.Key == "" {
			return
		}
		// Although all attributes in the Record are already resolved,
		// This one came from the user, so it may not have been.
		v = a.Value.Resolve()
	}

	if v.Kind() == slog.KindGroup {
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

func (s *formatState) appendError(err error) {
	s.appendString(fmt.Sprintf("!ERROR:%v", err))
}

func (s *formatState) appendKey(key string) {
	s.buf.WriteString(s.sep)
	if s.prefix != nil {
		// TODO: optimize by avoiding allocation.
		s.appendString(string(*s.prefix) + key)
	} else {
		s.appendString(key)
	}
	s.buf.WriteString(s.f.sty.valueSep())
	s.sep = s.f.sty.attrSep()
}

func (s *formatState) appendSource(file string, line int) {
	s.f.sty.appendSource(s, file, line)
}

func (s *formatState) appendString(str string) {
	s.f.sty.appendString(s, str)
}

func (s *formatState) appendValue(v slog.Value) {
	s.f.sty.appendValue(s, v)
}

func (s *formatState) appendTime(t time.Time) {
	s.f.sty.appendTime(s, t)
}
