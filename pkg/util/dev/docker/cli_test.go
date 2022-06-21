//go:build !nodev

package docker

import (
	"context"
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

func TestCmd(t *testing.T) {
	ctx := context.Background()

	inspects := check.Must(CliInspectAll(ctx))

	fmt.Printf("%+v", inspects)
}
