//go:build !nodev

package dev

import (
	"os"
	"testing"

	"github.com/wrmsr/bane/pkg/util/log"
)

func DoTestMain(m *testing.M, init func()) {
	if init != nil {
		init()
	}

	_, closeInjector := SetupInjector()
	defer log.OrError(closeInjector)

	os.Exit(m.Run())
}
