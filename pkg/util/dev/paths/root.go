//go:build !nodev

package paths

import (
	"os"

	"github.com/wrmsr/bane/pkg/util/check"
	osu "github.com/wrmsr/bane/pkg/util/os"
)

func FindProjectRoot() string {
	cwd := check.Must1(os.Getwd())
	return check.Must1(osu.FindParentDirWithFile(cwd, "go.mod"))
}

func CdToProjectRoot() {
	check.Must(os.Chdir(FindProjectRoot()))
}
