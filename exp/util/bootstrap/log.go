package bootstrap

import (
	stdlog "log"

	stdslog "golang.org/x/exp/slog"

	eu "github.com/wrmsr/bane/pkg/util/errors"
	"github.com/wrmsr/bane/pkg/util/slog"
	slogf "github.com/wrmsr/bane/pkg/util/slog/format"
)

type LogConfig struct {
	Style string
}

func LogBootstrap(cfg LogConfig, ds *eu.DeferStack) error {
	//if cfg.NoDefault {
	//	return nil
	//}

	h := slog.NewFormatterHandler(
		slogf.NewJsonFormatter(slogf.FormatterOpts{}),
		stdlog.Default().Writer().Write,
		slog.LevelInfo,
	)

	stdslog.SetDefault(stdslog.New(slog.AsStdHandler(&h)))

	return nil
}
