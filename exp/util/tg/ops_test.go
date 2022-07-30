package tg

import (
	"testing"

	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
)

func TestOpTypes(t *testing.T) {
	tu.AssertEqual(t, SignOp.Type(), UnaryOpType)
	tu.AssertEqual(t, AddOp.Type(), BinaryOpType)
}
