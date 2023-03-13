package bootstrap

import (
	"os"
	"runtime"
	"time"

	eu "github.com/wrmsr/bane/pkg/util/errors"
	bt "github.com/wrmsr/bane/pkg/util/types"
	"github.com/wrmsr/bane/pkg/util/workers"
)

type DumpConfig struct {
	Interval bt.Optional[time.Duration]
}

func DumpBootstrap(cfg DumpConfig, ds *eu.DeferStack) error {
	if cfg.Interval.Present() {
		buf := make([]byte, 256*1024)
		fn := func() {
			n := runtime.Stack(buf, true)
			_, _ = os.Stdout.Write(buf[:n]) // FIXME:
		}

		w := workers.NewSleepWorker(func() bool {
			fn()
			return true
		}, workers.SleepWorkerConfig{
			Interval: cfg.Interval.Value(),
		})
		w.Start()
		ds.DeferErr(func() error {
			w.Stop()
			return nil
		})
	}

	return nil
}
