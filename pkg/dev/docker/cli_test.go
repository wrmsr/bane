package docker

import (
	"context"
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/utils/check"
)

func TestCmd(t *testing.T) {
	ctx := context.Background()

	inspects := check.Must(CliInspectAll(ctx))

	fmt.Printf("%+v\n", inspects)
}
