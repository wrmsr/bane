package ssa

import (
	"fmt"
	"os"
	"testing"

	"golang.org/x/tools/go/analysis/passes/buildssa"

	"github.com/wrmsr/bane/exp/util/go/analysis/runner"
	"github.com/wrmsr/bane/pkg/util/check"
)

func TestSsa(t *testing.T) {
	analyzer := buildssa.Analyzer
	dir := check.Must1(os.Getwd())
	patterns := []string{"test"}

	pkgs := check.Must1(runner.LoadPackages(analyzer, dir, patterns...))
	results := runner.Run(analyzer, pkgs, runner.Options{})
	ssainfo := results[0].Result.(*buildssa.SSA)

	fmt.Println(ssainfo)
}
