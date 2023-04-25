package base

import stru "github.com/wrmsr/bane/pkg/util/strings"

type Dsn struct {
	Host string
	Port int

	User string
	Pass stru.Secret
}
