//go:build !nodev

package debug

import (
	"context"
	"fmt"
	"testing"
)

func TestIsDebuggerAttached(t *testing.T) {
	fmt.Println(IsDebuggerAttached(context.Background()))
}
