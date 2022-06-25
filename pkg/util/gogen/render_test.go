package gogen

import (
	"fmt"
	"strings"
	"testing"

	iou "github.com/wrmsr/bane/pkg/util/io"
)

func TestBuilder(t *testing.T) {
	n := NewIf(
		NewLit("foo"),
		NewBlock(
			NewExprStmt(NewLit("bar")),
		),
	)

	var sb strings.Builder
	Render(iou.NewIndentWriter(iou.NewDiscardStringWriter(&sb), "\t"), n)
	fmt.Println(sb.String())
}
