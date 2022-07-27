//go:build !nodev

package impl

import (
	"context"
	"io/ioutil"
	"os/exec"
	"time"

	"github.com/wrmsr/bane/pkg/util/check"
)

func FormatCode(ctx context.Context, s string) string {
	tf := check.Must1(ioutil.TempFile("", "def.go"))
	check.Must1(tf.WriteString(s))
	check.Must(tf.Close())

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "go", "run", "cmd/gofmt", "-s", "-w", tf.Name())
	check.Must(cmd.Run())

	return string(check.Must1(ioutil.ReadFile(tf.Name())))
}
