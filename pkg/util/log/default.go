package log

//

type DefaultNilLogger struct {
	impl Logger
}

var _ Logger = DefaultNilLogger{}

func (d DefaultNilLogger) Log(lvl Level, msg string, args ...Arg) {
	if d.impl != nil {
		d.impl.Log(lvl, msg, args...)
	}
}

func (d DefaultNilLogger) Debug(msg string, args ...Arg) {
	if d.impl != nil {
		d.impl.Debug(msg, args...)
	}
}

func (d DefaultNilLogger) Info(msg string, args ...Arg) {
	if d.impl != nil {
		d.impl.Info(msg, args...)
	}
}

func (d DefaultNilLogger) Warn(msg string, args ...Arg) {
	if d.impl != nil {
		d.impl.Warn(msg, args...)
	}
}

func (d DefaultNilLogger) Error(msg string, args ...Arg) {
	if d.impl != nil {
		d.impl.Error(msg, args...)
	}
}

func (d DefaultNilLogger) Panic(msg string, args ...Arg) {
	if d.impl != nil {
		d.impl.Panic(msg, args...)
	}
}

func (d DefaultNilLogger) Fatal(msg string, args ...Arg) {
	if d.impl != nil {
		d.impl.Fatal(msg, args...)
	}
}

func (d DefaultNilLogger) IfError(err error, args ...Arg) {
	if d.impl != nil {
		d.impl.IfError(err, args...)
	}
}

func (d DefaultNilLogger) IfPanic(err error, args ...Arg) {
	if d.impl != nil {
		d.impl.IfPanic(err, args...)
	}
}

func (d DefaultNilLogger) IfFatal(err error, args ...Arg) {
	if d.impl != nil {
		d.impl.IfFatal(err, args...)
	}
}

func (d DefaultNilLogger) OrError(fn func() error, args ...Arg) {
	if d.impl != nil {
		d.impl.OrError(fn, args...)
	}
}

func (d DefaultNilLogger) OrPanic(fn func() error, args ...Arg) {
	if d.impl != nil {
		d.impl.OrPanic(fn, args...)
	}
}

func (d DefaultNilLogger) OrFatal(fn func() error, args ...Arg) {
	if d.impl != nil {
		d.impl.OrFatal(fn, args...)
	}
}

//

type DefaultGlobalLogger struct {
	impl Logger
}

var _ Logger = DefaultGlobalLogger{}

func (d DefaultGlobalLogger) Log(lvl Level, msg string, args ...Arg) {
	if d.impl != nil {
		d.impl.Log(lvl, msg, args...)
	} else {
		Global().Log(lvl, msg, args...)
	}
}

func (d DefaultGlobalLogger) Debug(msg string, args ...Arg) {
	if d.impl != nil {
		d.impl.Debug(msg, args...)
	} else {
		Global().Debug(msg, args...)
	}
}

func (d DefaultGlobalLogger) Info(msg string, args ...Arg) {
	if d.impl != nil {
		d.impl.Info(msg, args...)
	} else {
		Global().Info(msg, args...)
	}
}

func (d DefaultGlobalLogger) Warn(msg string, args ...Arg) {
	if d.impl != nil {
		d.impl.Warn(msg, args...)
	} else {
		Global().Warn(msg, args...)
	}
}

func (d DefaultGlobalLogger) Error(msg string, args ...Arg) {
	if d.impl != nil {
		d.impl.Error(msg, args...)
	} else {
		Global().Error(msg, args...)
	}
}

func (d DefaultGlobalLogger) Panic(msg string, args ...Arg) {
	if d.impl != nil {
		d.impl.Panic(msg, args...)
	} else {
		Global().Panic(msg, args...)
	}
}

func (d DefaultGlobalLogger) Fatal(msg string, args ...Arg) {
	if d.impl != nil {
		d.impl.Fatal(msg, args...)
	} else {
		Global().Fatal(msg, args...)
	}
}

func (d DefaultGlobalLogger) IfError(err error, args ...Arg) {
	if d.impl != nil {
		d.impl.IfError(err, args...)
	} else {
		Global().IfError(err, args...)
	}
}

func (d DefaultGlobalLogger) IfPanic(err error, args ...Arg) {
	if d.impl != nil {
		d.impl.IfPanic(err, args...)
	} else {
		Global().IfPanic(err, args...)
	}
}

func (d DefaultGlobalLogger) IfFatal(err error, args ...Arg) {
	if d.impl != nil {
		d.impl.IfFatal(err, args...)
	} else {
		Global().IfFatal(err, args...)
	}
}

func (d DefaultGlobalLogger) OrError(fn func() error, args ...Arg) {
	if d.impl != nil {
		d.impl.OrError(fn, args...)
	} else {
		Global().OrError(fn, args...)
	}
}

func (d DefaultGlobalLogger) OrPanic(fn func() error, args ...Arg) {
	if d.impl != nil {
		d.impl.OrPanic(fn, args...)
	} else {
		Global().OrPanic(fn, args...)
	}
}

func (d DefaultGlobalLogger) OrFatal(fn func() error, args ...Arg) {
	if d.impl != nil {
		d.impl.OrFatal(fn, args...)
	} else {
		Global().OrFatal(fn, args...)
	}
}
