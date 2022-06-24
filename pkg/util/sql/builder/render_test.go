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
			{
				Expr: &Identifier{
					Name: "foo",
				},
			},
			{
				Expr: &Literal{
					String: "420",
				},
			},
		},
		From: &Table{
			Identifier: &Identifier{
				Name: "bar",
			},
		},
		Where: &InfixExpr{
			Op: AndOp,
			Args: []Expr{
				&InfixExpr{
					Op: EqOp,
					Args: []Expr{
						&Identifier{
							Name: "baz",
						},
						&Literal{
							String: "500",
						},
					},
				},
				&InfixExpr{
					Op: EqOp,
					Args: []Expr{
						&Identifier{
							Name: "baz2",
						},
						&Literal{
							String: "501",
						},
					},
				},
			},
		},
	}

	var sb strings.Builder
	Render(iou.NewDiscardStringWriter(&sb), stmt)
	fmt.Println(sb.String())
}
