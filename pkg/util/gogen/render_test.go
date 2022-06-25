package gogen

import (
	"fmt"
	"strings"
	"testing"

	iou "github.com/wrmsr/bane/pkg/util/io"
	opt "github.com/wrmsr/bane/pkg/util/optional"
)

func TestBuilder(t *testing.T) {
	n := NewFunc(
		NewIdent("foo"),
		nil,
		opt.None[Type](),
		NewBlock(

			NewIf(
				NewLit("foo"),
				NewBlock(
					NewExprStmt(NewLit("bar")),
				),
			),
		),
	)

	var sb strings.Builder
	Render(iou.NewIndentWriter(iou.NewDiscardStringWriter(&sb), "\t"), n)
	fmt.Println(sb.String())
}
