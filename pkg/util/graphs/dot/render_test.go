package dot

import (
	"fmt"
	"strings"
	"testing"

	iou "github.com/wrmsr/bane/pkg/util/io"
)

func TestRender(t *testing.T) {
	g := Graph{
		Stmts: []Stmt{
			NewNode("a").SetAttr("shape", "box"),
			NewNode("b").SetAttr("label", []any{[]any{"a", "b"}, []any{"c", "d"}}),
			NewEdge("a", "b"),
		},
	}

	var sb strings.Builder
	Render(iou.NewDiscardStringWriter(&sb), g)
	fmt.Println(sb.String())

	//tu.AssertNoErr(t, Open(context.Background(), sb.String()))
}
