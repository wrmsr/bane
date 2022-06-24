package builder

import (
	"fmt"
	"strings"
	"testing"

	iou "github.com/wrmsr/bane/pkg/util/io"
)

func TestBuilder(t *testing.T) {
	stmt := Select{
		Items: []SelectItem{
			NewSelectItem(NewIdentifier("foo")),
			NewSelectItem(NewLiteral("foo")),
		},
		From: NewTable(NewIdentifier("bar")),
		Where: And(
			Eq(
				NewIdentifier("baz"),
				NewLiteral("500"),
			),
			Ne(
				NewIdentifier("baz2"),
				NewLiteral("501"),
			),
		),
	}

	var sb strings.Builder
	Render(iou.NewDiscardStringWriter(&sb), stmt)
	fmt.Println(sb.String())
}
