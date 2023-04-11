//go:build !nodev

package main

import (
	stdlog "log"

	stdslog "golang.org/x/exp/slog"

	"github.com/wrmsr/bane/exp/util/bootstrap"
	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/slog"
)

func main() {
	cl := check.Must1(bootstrap.Bootstrap(bootstrap.Config{}))
	defer slog.OrFatal(cl)

	slog.Info("hi")
	stdslog.Info("hi")
	stdlog.Println("hi")
}
