//go:build !nodev

package main

import (
	"bufio"
	stdlog "log"
	"os"
	"time"

	stdslog "golang.org/x/exp/slog"

	"github.com/wrmsr/bane/x/bootstrap"

	"github.com/wrmsr/bane/core/check"
	"github.com/wrmsr/bane/core/slog"
	bt "github.com/wrmsr/bane/core/types"
)

func main() {
	cl := check.Must1(bootstrap.Bootstrap(bootstrap.Config{
		Dump: bootstrap.DumpConfig{
			Interval: bt.Just(10 * time.Second),
		},
	}))
	defer slog.OrFatal(cl)

	slog.Info("hi")
	stdslog.Info("hi")
	stdlog.Println("hi")

	bufio.NewScanner(os.Stdin).Scan()
}
