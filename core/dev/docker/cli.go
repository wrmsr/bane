//go:build !nodev

package docker

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
)

type CliCmdError struct {
	Message string
}

func (e CliCmdError) Error() string {
	return fmt.Sprintf("docker cli cmd error: %s", e.Message)
}

func CliCmd(ctx context.Context, args ...string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, "docker", args...)

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	if err := cmd.Run(); err != nil {
		return nil, CliCmdError{errb.String()}
	}

	return outb.Bytes(), nil
}
