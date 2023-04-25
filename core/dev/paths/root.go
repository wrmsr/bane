//go:build !nodev

package paths

import (
	"os"

	"github.com/wrmsr/bane/core/check"
	osu "github.com/wrmsr/bane/core/os"
)

func FindProjectRoot() string {
	cwd := check.Must1(os.Getwd())
	return check.Must1(osu.FindParentDirWithFile(cwd, "go.mod"))
}

func CdToProjectRoot() {
	check.Must(os.Chdir(FindProjectRoot()))
}
