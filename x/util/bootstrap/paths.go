package bootstrap

import (
	"os"

	eu "github.com/wrmsr/bane/pkg/util/errors"
)

type PathsConfig struct {
	Cwd string
}

func PathsBootstrap(cfg PathsConfig, ds *eu.DeferStack) (err error) {
	if cfg.Cwd != "" {
		if err = os.Chdir(cfg.Cwd); err != nil {
			return err
		}
	}
	return
}
