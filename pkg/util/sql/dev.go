//go:build !nodev

package sql

import (
	"github.com/wrmsr/bane/pkg/util/dev"
	"github.com/wrmsr/bane/pkg/util/dev/docker"
	inj "github.com/wrmsr/bane/pkg/util/inject"
	opt "github.com/wrmsr/bane/pkg/util/optional"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	sqb "github.com/wrmsr/bane/pkg/util/sql/base"
	stru "github.com/wrmsr/bane/pkg/util/strings"
)

var _ = dev.Register(func() inj.Bindings {
	return inj.Bind(
		inj.As(inj.Tag(rfl.TypeOf[opt.Optional[sqb.Dsn]](), "mysql"), inj.Singleton(func(dsl *docker.ServiceLocator) opt.Optional[sqb.Dsn] {
			svc := dsl.Locate("bane-mysql")
			if svc == nil {
				return opt.None[sqb.Dsn]()
			}

			return opt.Just(sqb.Dsn{
				Host: svc.Host,
				Port: svc.Ports[3306],

				User: svc.Compose.Environment["MYSQL_USER"],
				Pass: stru.SecretOf(svc.Compose.Environment["MYSQL_PASS"]),
			})
		})),

		inj.As(inj.Tag(rfl.TypeOf[opt.Optional[sqb.Dsn]](), "postgres"), inj.Singleton(func(dsl *docker.ServiceLocator) opt.Optional[sqb.Dsn] {
			svc := dsl.Locate("bane-postgres")
			if svc == nil {
				return opt.None[sqb.Dsn]()
			}

			return opt.Just(sqb.Dsn{
				Host: svc.Host,
				Port: svc.Ports[5432],

				User: svc.Compose.Environment["POSTGRES_USER"],
				Pass: stru.SecretOf(svc.Compose.Environment["POSTGRES_PASSWORD"]),
			})
		})),
	)
})
