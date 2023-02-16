package consteval

type Evaluator struct {
	scope map[string]Value
}

func (e *Evaluator) Eval(v Value) Value {
	switch v := v.(type) {
	case Basic:
		return v
	}
	panic(v)
}
