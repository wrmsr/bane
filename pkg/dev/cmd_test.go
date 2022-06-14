package dev

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"testing"

	tu "github.com/wrmsr/bane/pkg/dev/testing"
)

func TestCmd(t *testing.T) {
	cmd := exec.CommandContext(context.Background(), "docker", "ps", "--format", "{{.ID}}")

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	tu.AssertNoErr(t, cmd.Run())

	scanner := bufio.NewScanner(&outb)
	var cids []string
	for scanner.Scan() {
		s := scanner.Text()
		if s != "" {
			cids = append(cids, s)
		}
	}

	tu.AssertNoErr(t, scanner.Err())

	cmd = exec.CommandContext(context.Background(), "docker", 'inspect" ')

	fmt.Println(cids)
}
