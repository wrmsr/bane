//go:build !nodev

package docker

import (
	"fmt"
	"testing"
)

func TestServices(t *testing.T) {
	sl := NewServiceLocator(DefaultServiceLocatorConfig())
	fmt.Println(sl.Locate("bane-mysql"))
}
