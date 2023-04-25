//go:build !nodev

package sql

import (
	"github.com/wrmsr/bane/pkg/util/dev"
	"github.com/wrmsr/bane/pkg/util/dev/docker"
	inj "github.com/wrmsr/bane/pkg/util/inject"
	stru "github.com/wrmsr/bane/pkg/util/strings"
	bt "github.com/wrmsr/bane/pkg/util/types"
	sqb "github.com/wrmsr/bane/sql/base"
)

var _ = dev.Register(func() inj.Bindings {

	return inj.Bind(
		inj.As(inj.KeyOf[bt.Optional[sqb.Dsn]]{"mysql"}, inj.Singleton{func(
			dsl *docker.ServiceLocator,
		) bt.Optional[sqb.Dsn] {
			svc := dsl.Locate("bane-mysql")
			if svc == nil {
				return bt.None[sqb.Dsn]()
			}

			return bt.Just(sqb.Dsn{
				Host: svc.Host,
				Port: svc.Ports[3306],

				User: svc.Compose.Environment["MYSQL_USER"],
				Pass: stru.SecretOf(svc.Compose.Environment["MYSQL_PASSWORD"]),
			})
		}}),

		inj.As(inj.KeyOf[bt.Optional[sqb.Dsn]]{"postgres"}, inj.Singleton{func(
			dsl *docker.ServiceLocator,
		) bt.Optional[sqb.Dsn] {
			svc := dsl.Locate("bane-postgres")
			if svc == nil {
				return bt.None[sqb.Dsn]()
			}

			return bt.Just(sqb.Dsn{
				Host: svc.Host,
				Port: svc.Ports[5432],

				User: svc.Compose.Environment["POSTGRES_USER"],
				Pass: stru.SecretOf(svc.Compose.Environment["POSTGRES_PASSWORD"]),
			})
		}}),
	)
})
