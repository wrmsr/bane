package handlers

import "time"

type jsonCommonHandlerImpl struct{}

var _ commonHandlerImpl = jsonCommonHandlerImpl{}

func (i jsonCommonHandlerImpl) attrSep() string  { return "," }
func (i jsonCommonHandlerImpl) valueSep() string { return "=" }

func (i jsonCommonHandlerImpl) beginHandle(s *handleState) {
	s.buf.WriteByte('{')
}

func (i jsonCommonHandlerImpl) endHandle(s *handleState) {
	// Close all open groups.
	for range s.h.groups {
		s.buf.WriteByte('}')
	}
	// Close the top-level object.
	s.buf.WriteByte('}')
}

func (i jsonCommonHandlerImpl) openGroup(s *handleState, name string) {
	s.appendKey(name)
	s.buf.WriteByte('{')
	s.sep = ""
}

func (i jsonCommonHandlerImpl) closeGroup(s *handleState, name string) {
	s.buf.WriteByte('}')
}

func (i jsonCommonHandlerImpl) appendSource(s *handleState, file string, line int) {
	s.buf.WriteByte('"')
	*s.buf = appendEscapedJSONString(*s.buf, file)
	s.buf.WriteByte(':')
	s.buf.WritePosInt(line)
	s.buf.WriteByte('"')
}

func (i jsonCommonHandlerImpl) appendString(s *handleState, str string) {
	s.buf.WriteByte('"')
	*s.buf = appendEscapedJSONString(*s.buf, str)
	s.buf.WriteByte('"')
}

func (i jsonCommonHandlerImpl) appendValue(s *handleState, v Value) {
	var err error
	err = appendJSONValue(s, v)
	if err != nil {
		s.appendError(err)
	}
}

func (i jsonCommonHandlerImpl) appendTime(s *handleState, t time.Time) {
	appendJSONTime(s, t)
}
