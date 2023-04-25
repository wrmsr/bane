package base

import stru "github.com/wrmsr/bane/core/strings"

type Dsn struct {
	Host string
	Port int

	User string
	Pass stru.Secret
}
