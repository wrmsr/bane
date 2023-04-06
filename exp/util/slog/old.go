package slog

/*
func Log(lvl Level, msg string, args ...Arg) { Global().Log(lvl, msg, args...) }

func Debug(msg string, args ...Arg) { Global().Log(DebugLevel, msg, args...) }
func Info(msg string, args ...Arg)  { Global().Log(InfoLevel, msg, args...) }
func Warn(msg string, args ...Arg)  { Global().Log(WarnLevel, msg, args...) }
func Error(msg string, args ...Arg) { Global().Log(ErrorLevel, msg, args...) }

func Panic(msg any, args ...Arg) {
	if msg, ok := msg.(string); ok {
		Global().Log(PanicLevel, msg, args...)
		return
	}
	Global().Log(PanicLevel, fmt.Sprint(msg), args...)
}

func Fatal(msg any, args ...Arg) {
	if msg, ok := msg.(string); ok {
		Global().Log(FatalLevel, msg, args...)
		return
	}
	Global().Log(FatalLevel, fmt.Sprint(msg), args...)
}

func IfError(err error, args ...Arg) {
	if err != nil {
		Global().Log(ErrorLevel, "error", append(args, E(err))...)
	}
}

func IfPanic(err error, args ...Arg) {
	if err != nil {
		Global().Log(PanicLevel, "error", append(args, E(err))...)
	}
}

func IfFatal(err error, args ...Arg) {
	if err != nil {
		Global().Log(FatalLevel, "error", append(args, E(err))...)
	}
}

func OrError(fn func() error, args ...Arg) {
	if err := fn(); err != nil {
		Global().Log(ErrorLevel, "error", append(args, E(err))...)
	}
}

func OrPanic(fn func() error, args ...Arg) {
	if err := fn(); err != nil {
		Global().Log(PanicLevel, "error", append(args, E(err))...)
	}
}

func OrFatal(fn func() error, args ...Arg) {
	if err := fn(); err != nil {
		Global().Log(FatalLevel, "error", append(args, E(err))...)
	}
}

//

func Print(v ...any)                 { Global().Log(InfoLevel, fmt.Sprint(v...)) }
func Printf(format string, v ...any) { Global().Log(InfoLevel, fmt.Sprintf(format, v...)) }
func Println(v ...any)               { Global().Log(InfoLevel, fmt.Sprintln(v...)) }

func Fatalf(format string, v ...any) { Global().Log(FatalLevel, fmt.Sprintf(format, v...)) }
func Fatalln(v ...any)               { Global().Log(FatalLevel, fmt.Sprintln(v...)) }

func Panicf(format string, v ...any) { Global().Log(PanicLevel, fmt.Sprintf(format, v...)) }
func Panicln(v ...any)               { Global().Log(PanicLevel, fmt.Sprintln(v...)) }
*/
