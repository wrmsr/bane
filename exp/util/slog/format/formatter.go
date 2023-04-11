// Copyright 2022 The Go Authors. All rights reserved. Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.
package format

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/slices"

	"github.com/wrmsr/bane/exp/util/slog"
	"github.com/wrmsr/bane/exp/util/slog/buffer"
)

//

type FormatterOpts struct {
	AddSource bool

	ReplaceAttr func(groups []string, a slog.Attr) slog.Attr
}

//

type Style interface {
	attrSep() string
	valueSep() string

	begin(s *state)
	end(s *state)

	openGroup(s *state, name string)
	closeGroup(s *state, name string)

	appendSource(s *state, file string, line int)
	appendString(s *state, str string)
	appendValue(s *state, v slog.Value)
	appendTime(s *state, t time.Time)
}

//

type formatter struct {
	opts FormatterOpts
	sty  Style

	mtx sync.Mutex

	preformattedAttrs []byte
	groupPrefix       string
	groups            []string
	nOpenGroups       int
}

var _ slog.Formatter = &formatter{}

func (f *formatter) clone() *formatter {
	return &formatter{
		opts: f.opts,

		preformattedAttrs: slices.Clip(f.preformattedAttrs),
		groupPrefix:       f.groupPrefix,
		groups:            slices.Clip(f.groups),
		nOpenGroups:       f.nOpenGroups,
	}
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

func (f *formatter) WithAttrs(as []slog.Attr) slog.Formatter {
	return f.withAttrs(as)
}

func (f *formatter) withGroup(name string) *formatter {
	if name == "" {
		return f
	}
	h2 := f.clone()
	h2.groups = append(h2.groups, name)
	return h2
}

func (f *formatter) WithGroup(name string) slog.Formatter {
	return f.WithGroup(name)
}

func (f *formatter) Format(r slog.Record, w slog.Writer) error {
	st := f.newHandleState(buffer.New(), true, "", nil)
	defer st.free()

	f.sty.begin(&st)

	// Built-in attributes. They are not in a group.
	stateGroups := st.groups
	st.groups = nil // So ReplaceAttrs sees no groups instead of the pre groups.
	rep := f.opts.ReplaceAttr

	// time
	if !r.Time.IsZero() {
		key := slog.TimeKey
		val := r.Time.Round(0) // strip monotonic to match Attr behavior
		if rep == nil {
			st.appendKey(key)
			st.appendTime(val)
		} else {
			st.appendAttr(slog.Time(key, val))
		}
	}

	// level
	key := slog.LevelKey
	val := r.Level
	if rep == nil {
		st.appendKey(key)
		st.appendString(val.String())
	} else {
		st.appendAttr(slog.Any(key, val))
	}

	// source
	if f.opts.AddSource {
		frame := slog.RecordFrame(r)
		if frame.File != "" {
			key := slog.SourceKey
			if rep == nil {
				st.appendKey(key)
				st.appendSource(frame.File, frame.Line)
			} else {
				buf := buffer.New()
				buf.WriteString(frame.File) // TODO: escape?
				buf.WriteByte(':')
				buf.WritePosInt(frame.Line)
				s := buf.String()
				buf.Free()
				st.appendAttr(slog.String(key, s))
			}
		}
	}

	key = slog.MessageKey
	msg := r.Message
	if rep == nil {
		st.appendKey(key)
		st.appendString(msg)
	} else {
		st.appendAttr(slog.String(key, msg))
	}

	st.groups = stateGroups // Restore groups passed to ReplaceAttrs.
	st.appendNonBuiltIns(r)
	st.buf.WriteByte('\n')

	f.mtx.Lock()
	defer f.mtx.Unlock()
	_, err := w(*st.buf)
	return err
}

//

func (s *state) appendNonBuiltIns(r slog.Record) {
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

// state holds state for a single call to commonHandler.handle. The initial value of sep determines whether to
// emit a separator before the next key, after which it stays true.
type state struct {
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

func (f *formatter) newHandleState(buf *buffer.Buffer, freeBuf bool, sep string, prefix *buffer.Buffer) state {
	s := state{
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

func (s *state) free() {
	if s.freeBuf {
		s.buf.Free()
	}
	if gs := s.groups; gs != nil {
		*gs = (*gs)[:0]
		groupPool.Put(gs)
	}
}

func (s *state) openGroups() {
	for _, n := range s.f.groups[s.f.nOpenGroups:] {
		s.openGroup(n)
	}
}

// Separator for group names and keys.
const keyComponentSep = '.'

// openGroup starts a new group of attributes with the given name.
func (s *state) openGroup(name string) {
	s.f.sty.openGroup(s, name)
	// Collect group names for ReplaceAttr.
	if s.groups != nil {
		*s.groups = append(*s.groups, name)
	}
}

// closeGroup ends the group with the given name.
func (s *state) closeGroup(name string) {
	s.f.sty.closeGroup(s, name)
	s.sep = s.f.sty.attrSep()
	if s.groups != nil {
		*s.groups = (*s.groups)[:len(*s.groups)-1]
	}
}

// appendAttr appends the Attr's key and value using app. It handles replacement and checking for an empty key. after
// replacement).
func (s *state) appendAttr(a slog.Attr) {
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

func (s *state) appendError(err error) {
	s.appendString(fmt.Sprintf("!ERROR:%v", err))
}

func (s *state) appendKey(key string) {
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

func (s *state) appendSource(file string, line int) {
	s.f.sty.appendSource(s, file, line)
}

func (s *state) appendString(str string) {
	s.f.sty.appendString(s, str)
}

func (s *state) appendValue(v slog.Value) {
	s.f.sty.appendValue(s, v)
}

func (s *state) appendTime(t time.Time) {
	s.f.sty.appendTime(s, t)
}
