package log

//

type Arg struct {
	Name  string
	Value any
}

func A(name string, value any) Arg {
	return Arg{Name: name, Value: value}
}

func Err(err error) Arg {
	return Arg{Name: "error", Value: err}
}

//

type BaseLogger interface {
	Log(lvl Level, msg string, args ...Arg)
}

type BaseLoggerImpl struct {
	ll LineLogger
}

var _ BaseLogger = BaseLoggerImpl{}

func (l BaseLoggerImpl) Log(lvl Level, msg string, args ...Arg) {
	err := l.ll.LogLine(Line{
		Level:   lvl,
		Message: msg,
		Args:    args,
	})
	if err != nil {
		panic(err)
	}
}

//

type Logger interface {
	BaseLogger

	Debug(msg string, args ...Arg)
	Info(msg string, args ...Arg)
	Warn(msg string, args ...Arg)
	Error(msg string, args ...Arg)
	Panic(msg string, args ...Arg)
	Fatal(msg string, args ...Arg)

	OrError(fn func() error, args ...Arg) error
	OrPanic(fn func() error, args ...Arg) error
	OrFatal(fn func() error, args ...Arg) error
}

type LoggerImpl struct {
	BaseLogger
}

var _ Logger = LoggerImpl{}

func (l LoggerImpl) Debug(msg string, args ...Arg) { l.Log(DebugLevel, msg, args...) }
func (l LoggerImpl) Info(msg string, args ...Arg)  { l.Log(InfoLevel, msg, args...) }
func (l LoggerImpl) Warn(msg string, args ...Arg)  { l.Log(WarnLevel, msg, args...) }
func (l LoggerImpl) Error(msg string, args ...Arg) { l.Log(ErrorLevel, msg, args...) }
func (l LoggerImpl) Panic(msg string, args ...Arg) { l.Log(PanicLevel, msg, args...) }
func (l LoggerImpl) Fatal(msg string, args ...Arg) { l.Log(FatalLevel, msg, args...) }

func (l LoggerImpl) OrError(fn func() error, args ...Arg) error {
	err := fn()
	if err != nil {
		l.Log(ErrorLevel, "error", append(args, Err(err))...)
	}
	return err
}

func (l LoggerImpl) OrPanic(fn func() error, args ...Arg) error {
	err := fn()
	if err != nil {
		l.Log(ErrorLevel, "error", append(args, Err(err))...)
	}
	return err
}

func (l LoggerImpl) OrFatal(fn func() error, args ...Arg) error {
	err := fn()
	if err != nil {
		l.Log(ErrorLevel, "error", append(args, Err(err))...)
	}
	return err
}
