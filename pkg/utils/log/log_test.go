package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	l := LoggerImpl{
		BaseLoggerImpl{
			LineLoggerImpl{
				JsonFormatter{},
				StdWriter{},
			},
		},
	}

	l.Info("hi", A("foo", 100))
}
