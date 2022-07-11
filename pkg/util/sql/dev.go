//go:build !nodev

package sql

import (
	"github.com/wrmsr/bane/pkg/util/dev"
	inj "github.com/wrmsr/bane/pkg/util/inject"
)

var _ = dev.Register(func() inj.Bindings {
	return inj.Bind(
	//inj.Singleton(func () )
	)
})
