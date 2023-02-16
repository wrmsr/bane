package consteval

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	ju "github.com/wrmsr/bane/pkg/util/json"
	msh "github.com/wrmsr/bane/pkg/util/marshal"
)

func TestEval(t *testing.T) {
	e := &Evaluator{
		scope: make(map[string]Value),
	}

	iv := Basic{K: IntBasic, S: "420"}
	ov := e.Eval(iv)
	fmt.Println(check.Must1(ju.MarshalPretty(check.Must1(msh.Marshal(&ov)))))
}
