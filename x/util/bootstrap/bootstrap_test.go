package bootstrap

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/wrmsr/bane/pkg/util/check"
	eu "github.com/wrmsr/bane/pkg/util/errors"
	log "github.com/wrmsr/bane/pkg/util/slog"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

func TestBootstrap(t *testing.T) {
	cfg := Config{
		Env: EnvConfig{
			Entries: []EnvEntry{
				SetEnvEntry{"FOO", "BAR"},
				UnsetEnvEntry{"BAZ"},
			},
		},
		Limits: LimitsConfig{
			Limits: map[string]bt.Optional[uint64]{
				"nofile": bt.Just(uint64(1024)),
			},
			Reset: true,
		},
	}

	pe := func(k string) {
		fmt.Printf("%s = %s\n", k, os.Getenv(k))
	}
	pe("FOO")
	pe("BAR")
	shutdown := check.Must1(Bootstrap(cfg))
	pe("FOO")
	pe("BAR")
	check.Must(shutdown())
	pe("FOO")
	pe("BAR")
}

func TestDump(t *testing.T) {
	ds := eu.NewDeferStack()
	defer log.OrError(ds.Call)

	check.Must(DumpBootstrap(DumpConfig{Interval: bt.Just(1 * time.Second)}, ds))

	time.Sleep(3 * time.Second)
}
