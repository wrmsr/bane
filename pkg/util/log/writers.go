package log

import "log"

type Writer interface {
	Write(s string) error
}

type StdWriter struct{}

var _ Writer = StdWriter{}

func (w StdWriter) Write(s string) error {
	log.Print(s)
	return nil
}
