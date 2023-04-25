package exprs

import (
	"fmt"
	"testing"
)

func TestExprs(t *testing.T) {
	fmt.Println(If{
		Then: []Expr{
			Nop{},
			Nop{},
		},
		Else: []Expr{
			Nop{},
		},
	}.Text().String())
}
