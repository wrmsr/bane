package tg

import (
	"fmt"
	"testing"
)

func TestConvArgs(t *testing.T) {
	ca := BuildConvArgs(
		Shape{1, 784, 69, 1},
		Shape{128, 784, 1, 1},
		ConvOpts{},
	)
	fmt.Printf("%+v\n", ca)
}
