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

	var f slog.Formatter
	switch cfg.Format {
	case LogText:
		f = slogf.NewJsonFormatter(fo)
	case LogJson:
		f = slogf.NewJsonFormatter(fo)
	case LogNone:
		return nil
	default:
		panic(cfg.Format)
	}

	w := stdlog.Default().Writer()

	h := slog.NewFormatterHandler(
		f,
		w.Write,
		cfg.Level,
	)

	o := stdslog.Default()
	stdslog.SetDefault(stdslog.New(slog.AsStdHandler(&h)))
	ds.Defer(func() {
		stdslog.SetDefault(o)
	})

	return nil
}
