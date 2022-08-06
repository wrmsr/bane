package log

type DefaultLogger struct {
	impl Logger
}

var _ Logger = DefaultLogger{}

func (d DefaultLogger) Log(lvl Level, msg string, args ...Arg) {
	if d.impl != nil {
		d.impl.Log(lvl, msg, args...)
	}
}

func (d DefaultLogger) Debug(msg string, args ...Arg) {
	if d.impl != nil {
		d.impl.Debug(msg, args...)
	}
}

func (d DefaultLogger) Info(msg string, args ...Arg) {
	if d.impl != nil {
		d.impl.Info(msg, args...)
	}
}

func (d DefaultLogger) Warn(msg string, args ...Arg) {
	if d.impl != nil {
		d.impl.Warn(msg, args...)
	}
}

func (d DefaultLogger) Error(msg string, args ...Arg) {
	if d.impl != nil {
		d.impl.Error(msg, args...)
	}
}

func (d DefaultLogger) Panic(msg string, args ...Arg) {
	if d.impl != nil {
		d.impl.Panic(msg, args...)
	}
}

func (d DefaultLogger) Fatal(msg string, args ...Arg) {
	if d.impl != nil {
		d.impl.Fatal(msg, args...)
	}
}

func (d DefaultLogger) IfError(err error, args ...Arg) {
	if d.impl != nil {
		d.impl.IfError(err, args...)
	}
}

func (d DefaultLogger) IfPanic(err error, args ...Arg) {
	if d.impl != nil {
		d.impl.IfPanic(err, args...)
	}
}

func (d DefaultLogger) IfFatal(err error, args ...Arg) {
	if d.impl != nil {
		d.impl.IfFatal(err, args...)
	}
}

func (d DefaultLogger) OrError(fn func() error, args ...Arg) {
	if d.impl != nil {
		d.impl.OrError(fn, args...)
	}
}

func (d DefaultLogger) OrPanic(fn func() error, args ...Arg) {
	if d.impl != nil {
		d.impl.OrPanic(fn, args...)
	}
}

func (d DefaultLogger) OrFatal(fn func() error, args ...Arg) {
	if d.impl != nil {
		d.impl.OrFatal(fn, args...)
	}
}
