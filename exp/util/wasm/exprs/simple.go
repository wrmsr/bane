package exprs

import wt "github.com/wrmsr/bane/exp/util/wasm/types"

//

type Const struct {
	expr
	S  string
	Ty wt.Type
}

var _ Expr = Const{}
