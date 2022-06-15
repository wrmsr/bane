package log

type Line struct {
	Level   Level
	Message string
	Args    []Arg

	StackOffset int
}

type LineLogger interface {
	LogLine(line Line) error
}

type LineLoggerImpl struct {
	f Formatter
	w Writer
}

var _ LineLogger = LineLoggerImpl{}

func (l LineLoggerImpl) LogLine(line Line) error {
	s, err := l.f.FormatLine(line)
	if err != nil {
		return err
	}
	return l.w.Write(s)
}
