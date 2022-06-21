//go:build !nodev

package testing

import (
	"testing"
	"time"
)

//

type Config struct {
	DefaultTimeout time.Duration
}

func DefaultConfig() Config {
	return Config{
		DefaultTimeout: time.Duration(10) * time.Second,
	}
}

//

func Main(t *testing.M, init func()) {
	init()
}
