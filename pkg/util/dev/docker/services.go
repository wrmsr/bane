//go:build !nodev

package docker

import (
	"context"
	"path/filepath"

	"github.com/wrmsr/bane/pkg/util/dev/paths"
	"github.com/wrmsr/bane/pkg/util/log"
	opt "github.com/wrmsr/bane/pkg/util/optional"
)

type ServiceLocator struct {
	pss opt.Optional[Pss]
	ins opt.Optional[Inspects]
	cmp opt.Optional[*ComposeConfig]
}

func (sl *ServiceLocator) Pss(ctx context.Context) Pss {
	return opt.SetIfAbsent(&sl.pss, func() Pss {
		pss, err := CliPs(ctx)
		if err != nil {
			log.Error("docker error", log.Err(err))
			return nil
		}
		return pss
	})
}

func (sl *ServiceLocator) Inspects(ctx context.Context) Inspects {
	return opt.SetIfAbsent(&sl.ins, func() Inspects {
		ins, err := CliInspectAll(ctx)
		if err != nil {
			log.Error("docker error", log.Err(err))
			return nil
		}
		return ins
	})
}

func (sl *ServiceLocator) ComposeConfig() *ComposeConfig {
	return opt.SetIfAbsent(&sl.cmp, func() *ComposeConfig {
		fp := filepath.Join(paths.FindProjectRoot(), "docker", "docker-compose.yml")
		cmp, err := ReadComposeConfig(fp)
		if err != nil {
			log.Error("docker error", log.Err(err))
			return nil
		}
		return cmp
	})
}
