//go:build !nodev

package test

import (
	"fmt"
	"testing"
)

func TestFoo(t *testing.T) {
	foof("hi %s\n", "there")

	var f Foo
	f.init()
	fmt.Printf("%+v\n", f)
}
