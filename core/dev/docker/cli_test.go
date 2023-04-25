//go:build !nodev

package docker

import (
	"context"
	"fmt"
	"testing"

	"github.com/wrmsr/bane/core/check"
)

func TestCmd(t *testing.T) {
	ctx := context.Background()

	inspects := check.Must1(CliInspectAll(ctx))

	fmt.Printf("%+v", inspects)
}
