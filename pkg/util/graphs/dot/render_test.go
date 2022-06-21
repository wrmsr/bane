package dot

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestRender(*testing.T) {
	g := Graph{
		Stmts: []Stmt{
			NewNode("a").SetAttr("shape", "box"),
			NewNode("b").SetAttr("label", []any{[]any{"a", "b"}, []any{"c", "d"}}),
			NewEdge("a", "b"),
		},
	}

	var sb strings.Builder
	Render(bufio.NewWriter(&sb), g)
	fmt.Println(sb.String())
}
