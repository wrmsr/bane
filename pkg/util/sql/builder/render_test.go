package builder

import (
	"fmt"
	"strings"
	"testing"

	iou "github.com/wrmsr/bane/pkg/util/io"
)

func TestBuilder(t *testing.T) {
	stmt := &Select{
		Items: []*SelectItem{
			NewSelectItem(NewIdentifier("foo"), nil),
			NewSelectItem(NewLiteral("foo"), nil),
		},
		From: NewTable(NewIdentifier("bar")),
		Where: NewInfixExpr(
			AndOp,
			NewInfixExpr(
				EqOp,
				NewIdentifier("baz"),
				NewLiteral("500"),
			),
			NewInfixExpr(
				NeOp,
				NewIdentifier("baz2"),
				NewLiteral("501"),
			),
		),
	}

	var sb strings.Builder
	Render(iou.NewDiscardStringWriter(&sb), stmt)
	fmt.Println(sb.String())
}
