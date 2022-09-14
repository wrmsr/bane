//go:build !nodev

package docker

import (
	"context"
	"net"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/dev"
	"github.com/wrmsr/bane/pkg/util/dev/paths"
	fnu "github.com/wrmsr/bane/pkg/util/funcs"
	inj "github.com/wrmsr/bane/pkg/util/inject"
	"github.com/wrmsr/bane/pkg/util/log"
	"github.com/wrmsr/bane/pkg/util/slices"
	stru "github.com/wrmsr/bane/pkg/util/strings"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

var _ = dev.Register(func() inj.Bindings {
	return inj.Bind(
		inj.Singleton{DefaultServiceLocatorConfig},
		inj.Singleton{NewServiceLocator},
	)
})

//

type ServiceLocatorConfig struct {
	Timeout      time.Duration
	ComposePaths []string
}

func DefaultServiceLocatorConfig() ServiceLocatorConfig {
	return ServiceLocatorConfig{
		Timeout: time.Duration(5) * time.Second,
		ComposePaths: []string{
			filepath.Join("docker", "docker-compose-ext.yml"),
			filepath.Join("docker", "docker-compose.yml"),
		},
	}
}

//

type Service struct {
	Name string

	Host  string
	Ports map[int]int

	Ps      *Ps
	Inspect *Inspect
	Compose *ComposeService
}

type ServiceLocator struct {
	cfg ServiceLocatorConfig

	pss bt.Optional[Pss]
	ins bt.Optional[Inspects]
	cmp bt.Optional[*ComposeConfig]

	m map[string]*Service
}

func NewServiceLocator(cfg ServiceLocatorConfig) *ServiceLocator {
	return &ServiceLocator{
		cfg: cfg,

		m: make(map[string]*Service),
	}
}

func (sl *ServiceLocator) Pss() Pss {
	return bt.SetIfAbsent(&sl.pss, func() Pss {
		ctx, cancel := context.WithTimeout(context.Background(), sl.cfg.Timeout)
		defer cancel()

		pss, err := CliPs(ctx)
		if err != nil {
			log.Error("docker error", log.E(err))
			return nil
		}

		return pss
	})
}

func (sl *ServiceLocator) Inspects() Inspects {
	return bt.SetIfAbsent(&sl.ins, func() Inspects {
		pss := sl.Pss()

		ctx, cancel := context.WithTimeout(context.Background(), sl.cfg.Timeout)
		defer cancel()

		ins, err := CliInspect(ctx, pss.Ids())
		if err != nil {
			log.Error("docker error", log.E(err))
			return nil
		}

		return ins
	})
}

func (sl *ServiceLocator) Compose() *ComposeConfig {
	return bt.SetIfAbsent(&sl.cmp, func() *ComposeConfig {
		fps := slices.Map(func(p string) string {
			return filepath.Join(paths.FindProjectRoot(), p)
		}, sl.cfg.ComposePaths)

		cmp, err := ReadComposeConfig(fps)
		if err != nil {
			log.Error("docker error", log.E(err))
			return nil
		}

		return cmp
	})
}

func (sl *ServiceLocator) Locate(name string) *Service {
	if s, ok := sl.m[name]; ok {
		return s
	}

	svc := &Service{
		Name: name,

		Ports: make(map[int]int),

		Compose: sl.Compose().Services[name],
	}

	if sl.Pss() != nil {
		svc.Ps = check.IgnoreOk1(slices.FindElemFunc(sl.Pss(), func(p *Ps) bool {
			return slices.Any(stru.TrimSpaceSplit(p.Names, ","), fnu.BindR1x1x1(MatchComposeName, name))
		}))
	}

	if sl.Inspects() != nil {
		svc.Inspect = check.IgnoreOk1(slices.FindElemFunc(sl.Inspects(), func(i *Inspect) bool {
			return MatchComposeName(strings.TrimLeft(i.Name, "/"), name)
		}))
	}

	canSee := func() bool {
		ctx, cancel := context.WithTimeout(context.Background(), sl.cfg.Timeout)
		defer cancel()

		ips, err := net.DefaultResolver.LookupHost(ctx, name)
		return err == nil && len(ips) > 0
	}()

	if canSee && svc.Compose != nil {
		svc.Host = name
		for _, p := range svc.Compose.Expose {
			pn := check.Must1(strconv.Atoi(p))
			svc.Ports[pn] = pn
		}

	} else if svc.Inspect != nil {
		svc.Host = "localhost"
		pbs := svc.Inspect.HostConfig.PortBindings
		if pbs != nil {
			for k, vs := range pbs {
				kps := strings.Split(k, "/")
				if len(kps) != 2 || kps[1] != "tcp" {
					continue
				}
				kp := check.Must1(strconv.Atoi(kps[0]))
				for _, v := range vs {
					if v.HostPort == "" {
						continue
					}
					vp := check.Must1(strconv.Atoi(v.HostPort))
					svc.Ports[kp] = vp
				}
			}
		}
	}

	return svc
}
