//go:build !nodev

package sql

import (
	"github.com/wrmsr/bane/pkg/util/dev"
	inj "github.com/wrmsr/bane/pkg/util/inject"
	sqb "github.com/wrmsr/bane/pkg/util/sql/base"
	stru "github.com/wrmsr/bane/pkg/util/strings"
)

var _ = dev.Register(func() inj.Bindings {
	return inj.Bind(
		inj.Singleton(func() sqb.Dsn {
			return sqb.Dsn{
				User: "foo",
				Pass: stru.SecretOf("bar"),
			}
		}),
	)
})
