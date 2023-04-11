//go:build !nodev

package main

import (
	"github.com/wrmsr/bane/exp/util/bootstrap"
	"github.com/wrmsr/bane/pkg/util/check"
	log "github.com/wrmsr/bane/pkg/util/slog"
)

func main() {
	cl := check.Must1(bootstrap.Bootstrap(bootstrap.Config{}))
	defer log.OrFatal(cl)

	log.Info("hi")
}
