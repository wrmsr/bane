package bootstrap

import (
	eu "github.com/wrmsr/bane/pkg/util/errors"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

type Config struct {
	Env    EnvConfig
	Paths  PathsConfig
	Limits LimitsConfig
}

func Bootstrap(cfg Config) (fn bt.ErrFn, err error) {
	ds := eu.NewDeferStack()
	defer func() {
		if err != nil {
			err = eu.Append(err, ds.Call())
		}
	}()

	if err = EnvBootstrap(cfg.Env, ds); err != nil {
		return
	}
	if err = PathsBootstrap(cfg.Paths, ds); err != nil {
		return
	}
	if err = LimitsBootstrap(cfg.Limits, ds); err != nil {
		return
	}

	return ds.Call, nil
}
