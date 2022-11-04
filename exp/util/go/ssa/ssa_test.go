package ssa

import (
	"fmt"
	"testing"

	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/packages"

	"github.com/wrmsr/bane/exp/util/go/analysis/runner"
	"github.com/wrmsr/bane/pkg/util/check"
)

func TestSsa(t *testing.T) {
	pkgs := check.Must1(packages.Load(
		&packages.Config{
			Mode: runner.PackageMode,
		},
		"github.com/wrmsr/bane/exp/util/go/ssa/test",
	))

	results := runner.Run(buildssa.Analyzer, pkgs, runner.Options{})

	ssainfo := results[0].Result.(*buildssa.SSA)

	fmt.Println(ssainfo)
}
