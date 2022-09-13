package dev

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/dev/paths"
	osu "github.com/wrmsr/bane/pkg/util/os"
)

const DatasetsRepo = "https://github.com/wrmsr/datasets"

func FetchDatasets() string {
	cd := filepath.Join(paths.FindProjectRoot(), ".cache")
	if !check.Must1(osu.Exists(cd)) {
		check.Must(os.Mkdir(cd, 0744))
	}

	dd := filepath.Join(cd, "datasets")
	if !check.Must1(osu.Exists(dd)) {
		ctx := context.Background()
		cmd := exec.CommandContext(ctx, "git", "clone", DatasetsRepo)
		cmd.Dir = cd
		check.Must(cmd.Run())
	}

	return dd
}
