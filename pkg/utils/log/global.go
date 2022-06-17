package log

func global() Logger {
	return nil
}

func Log(lvl Level, msg string, args ...Arg) { global().Log(lvl, msg, args...) }

func Debug(msg string, args ...Arg) { global().Debug(msg, args...) }
func Info(msg string, args ...Arg)  { global().Info(msg, args...) }
func Warn(msg string, args ...Arg)  { global().Warn(msg, args...) }
func Error(msg string, args ...Arg) { global().Error(msg, args...) }
func Panic(msg string, args ...Arg) { global().Panic(msg, args...) }
func Fatal(msg string, args ...Arg) { global().Fatal(msg, args...) }

func OrError(fn func() error, args ...Arg) error { return global().OrError(fn, args...) }
func OrPanic(fn func() error, args ...Arg) error { return global().OrError(fn, args...) }
func OrFatal(fn func() error, args ...Arg) error { return global().OrError(fn, args...) }
