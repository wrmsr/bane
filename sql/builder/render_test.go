package builder

import (
	"fmt"
	"strings"
	"testing"

	iou "github.com/wrmsr/bane/core/io"
)

func TestBuilder(t *testing.T) {
	n := Select{
		Items: []SelectItem{
			NewSelectItem(NewIdent("foo")),
			NewSelectItem(NewLiteral("foo")),
		},
		From: NewTable(NewIdent("bar")),
		Where: And(
			Eq(
				NewIdent("baz"),
				NewLiteral("500"),
			),
			Ne(
				NewIdent("baz2"),
				NewLiteral("501"),
			),
		),
	}

	var sb strings.Builder
	Render(iou.NewDiscardStringWriter(&sb), n)
	fmt.Println(sb.String())
}
