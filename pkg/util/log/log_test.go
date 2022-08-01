package log

import (
	"testing"
)

type someStruct struct {
	i int
}

func (s *someStruct) foo() {
	Info("hi", A("s.i", s.i))
}

func TestLog(t *testing.T) {
	Info("hi")
	s := someStruct{i: 420}
	s.foo()

	l := newDefaultLogger()
	l.Info("hi")
}
