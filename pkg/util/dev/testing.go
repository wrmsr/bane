//go:build !nodev

package dev

import (
	"os"
	"testing"

	"github.com/wrmsr/bane/pkg/util/dev/paths"
	log "github.com/wrmsr/bane/pkg/util/slog"
)

func DoTestMain(m *testing.M, init func()) {
	paths.CdToProjectRoot()

	if init != nil {
		init()
	}

	_, closeInjector := SetupInjector()
	defer log.OrError(closeInjector)

	os.Exit(m.Run())
}
