// Copyright 2022 The Go Authors. All rights reserved. Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.
package handlers

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

type CommonOpts struct {
	AddSource bool

	Level slog.Leveler

	ReplaceAttr func(groups []string, a slog.Attr) slog.Attr
}

//

type commonHandlerImpl interface {
	attrSep() string
	valueSep() string

	beginHandle(s *handleState)
	endHandle(s *handleState)

	openGroup(s *handleState, name string)
	closeGroup(s *handleState, name string)

	appendSource(s *handleState, file string, line int)
	appendString(s *handleState, str string)
	appendValue(s *handleState, v slog.Value)
	appendTime(s *handleState, t time.Time)
}

//

type commonHandler struct {
	opts CommonOpts
	w    io.Writer
	i    commonHandlerImpl

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

func (h *commonHandler) enabled(l slog.Level) bool {
	minLevel := slog.LevelInfo
	if h.opts.Level != nil {
		minLevel = h.opts.Level.Level()
	}
	return l >= minLevel
}

func (h *commonHandler) withAttrs(as []slog.Attr) *commonHandler {
	h2 := h.clone()

	// Pre-format the attributes as an optimization.
	prefix := buffer.New()
	defer prefix.Free()
	prefix.WriteString(h.groupPrefix)

	state := h2.newHandleState((*buffer.Buffer)(&h2.preformattedAttrs), false, "", prefix)
	defer state.free()
	if len(h2.preformattedAttrs) > 0 {
		state.sep = h.i.attrSep()
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

func (h *commonHandler) handle(r slog.Record) error {
	state := h.newHandleState(buffer.New(), true, "", nil)
	defer state.free()

	h.i.beginHandle(&state)

	// Built-in attributes. They are not in a group.
	stateGroups := state.groups
	state.groups = nil // So ReplaceAttrs sees no groups instead of the pre groups.
	rep := h.opts.ReplaceAttr

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
	if h.opts.AddSource {
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

	h.mtx.Lock()
	defer h.mtx.Unlock()
	_, err := h.w.Write(*state.buf)
	return err
}

//

func (s *handleState) appendNonBuiltIns(r slog.Record) {
	// preformatted Attrs
	if len(s.h.preformattedAttrs) > 0 {
		s.buf.WriteString(s.sep)
		_, _ = s.buf.Write(s.h.preformattedAttrs)
		s.sep = s.h.i.attrSep()
	}

	// Attrs in Record -- unlike the built-in ones, they are in groups started from WithGroup.
	s.prefix = buffer.New()
	defer s.prefix.Free()
	s.prefix.WriteString(s.h.groupPrefix)
	s.openGroups()
	r.Attrs(func(a slog.Attr) {
		s.appendAttr(a)
	})

	s.h.i.endHandle(s)
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

// openGroup starts a new group of attributes with the given name.
func (s *handleState) openGroup(name string) {
	s.h.i.openGroup(s, name)
	// Collect group names for ReplaceAttr.
	if s.groups != nil {
		*s.groups = append(*s.groups, name)
	}
}

// closeGroup ends the group with the given name.
func (s *handleState) closeGroup(name string) {
	s.h.i.closeGroup(s, name)
	s.sep = s.h.i.attrSep()
	if s.groups != nil {
		*s.groups = (*s.groups)[:len(*s.groups)-1]
	}
}

// appendAttr appends the Attr's key and value using app.
// It handles replacement and checking for an empty key.
// after replacement).
func (s *handleState) appendAttr(a slog.Attr) {
	v := a.Value

	// Elide a non-group with an empty key.
	if a.Key == "" && v.Kind() != slog.KindGroup {
		return
	}

	if rep := s.h.opts.ReplaceAttr; rep != nil && v.Kind() != slog.KindGroup {
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
	s.buf.WriteString(s.h.i.valueSep())
	s.sep = s.h.i.attrSep()
}

func (s *handleState) appendSource(file string, line int) {
	s.h.i.appendSource(s, file, line)
}

func (s *handleState) appendString(str string) {
	s.h.i.appendString(s, str)
}

func (s *handleState) appendValue(v slog.Value) {
	s.h.i.appendValue(s, v)
}

func (s *handleState) appendTime(t time.Time) {
	s.h.i.appendTime(s, t)
}
