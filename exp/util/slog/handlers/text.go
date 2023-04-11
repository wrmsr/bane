package handlers

import (
	"strconv"
	"time"
)

type textCommonHandlerImpl struct{}

var _ commonHandlerImpl = textCommonHandlerImpl{}

func (i textCommonHandlerImpl) attrSep() string  { return " " }
func (i textCommonHandlerImpl) valueSep() string { return ":" }

func (i textCommonHandlerImpl) beginHandle(s *handleState) {}

func (i textCommonHandlerImpl) endHandle(s *handleState) {}

func (i textCommonHandlerImpl) openGroup(s *handleState, name string) {
	s.prefix.WriteString(name)
	s.prefix.WriteByte(keyComponentSep)
}

func (i textCommonHandlerImpl) closeGroup(s *handleState, name string) {
	*s.prefix = (*s.prefix)[:len(*s.prefix)-len(name)-1 /* for keyComponentSep */]
}

func (i textCommonHandlerImpl) appendSource(s *handleState, file string, line int) {
	if needsQuoting(file) {
		s.appendString(file + ":" + strconv.Itoa(line))
	} else {
		// common case: no quoting needed.
		s.appendString(file)
		s.buf.WriteByte(':')
		s.buf.WritePosInt(line)
	}
}

func (i textCommonHandlerImpl) appendString(s *handleState, str string) {
	if needsQuoting(str) {
		*s.buf = strconv.AppendQuote(*s.buf, str)
	} else {
		s.buf.WriteString(str)
	}
}

func (i textCommonHandlerImpl) appendValue(s *handleState, v Value) {
	var err error
	err = appendTextValue(s, v)
	if err != nil {
		s.appendError(err)
	}
}

func (i textCommonHandlerImpl) appendTime(s *handleState, t time.Time) {
	writeTimeRFC3339Millis(s.buf, t)
}

