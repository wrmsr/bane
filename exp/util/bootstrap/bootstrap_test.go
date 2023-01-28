package bootstrap

import (
	"fmt"
	"os"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

func TestBootstrap(t *testing.T) {
	cfg := Config{
		Env: EnvConfig{
			Entries: []EnvEntry{
				SetEnvEntry{"FOO", "BAR"},
				UnsetEnvEntry{"BAZ"},
			},
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
