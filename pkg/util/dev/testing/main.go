//go:build !nodev

package testing

import "testing"

func Main(t *testing.M, init func()) {
	init()
}
