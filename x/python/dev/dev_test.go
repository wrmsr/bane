//go:build !nodev

package dev

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/wrmsr/bane/core/check"
	"github.com/wrmsr/bane/core/dev/paths"
)

func TestVenvRun(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	cmd := exec.CommandContext(
		ctx,
		filepath.Join(paths.FindProjectRoot(), ".venv/bin/python"),
		"-c",
		`
import json, sys
vi = sys.version_info
print(json.dumps({
    a: getattr(vi, a)
    for a in [
        'major',
        'minor',
        'micro',
        'releaselevel',
        'serial',
    ]
}))
`,
	)

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	check.Must(cmd.Run())

	var m map[string]any
	check.Must(json.Unmarshal(outb.Bytes(), &m))
	fmt.Println(m)
}
