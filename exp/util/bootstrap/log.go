package bootstrap

import (
	stdlog "log"

	stdslog "golang.org/x/exp/slog"

	eu "github.com/wrmsr/bane/pkg/util/errors"
	"github.com/wrmsr/bane/pkg/util/slog"
	slogf "github.com/wrmsr/bane/pkg/util/slog/format"
)

type LogFormat int8

const (
	LogText LogFormat = iota
	LogJson
	LogNone
)

type LogConfig struct {
	Format LogFormat
	Level  slog.Level
}

func LogBootstrap(cfg LogConfig, ds *eu.DeferStack) error {
	var fo slogf.FormatterOpts

	fo.AddSource = true

	var f slog.Formatter
	switch cfg.Format {
	case LogText:
		f = slogf.NewTextFormatter(fo)
	case LogJson:
		f = slogf.NewJsonFormatter(fo)
	case LogNone:
		return nil
	default:
		panic(cfg.Format)
	}

	lv := cfg.Level

	o0 := stdlog.Default()
	o1 := stdslog.Default()

	w := o0.Writer()
	h := slog.NewFormatterHandler(f, w.Write, lv)
	l := slog.NewLogger(&h)
	n := stdslog.New(slog.AsStdHandler(&h))

	stdslog.SetDefault(n)
	slog.InstallOldWriter(o0, l, lv)
	ds.Defer(func() {
		o0.SetOutput(w)
		stdslog.SetDefault(o1)
	})

	return nil
}
