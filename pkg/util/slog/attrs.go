package slog

//

const ErrorKey = "error"

func E(e error) Attr {
	return Any(ErrorKey, e)
}

//

const PanicKey = "panic"

func P(v any) Attr {
	return Any(PanicKey, v)
}

//

const BadKey = "!BADKEY"

func ArgsToAttr(args []any) (Attr, []any) {
	switch x := args[0].(type) {
	case string:
		if len(args) == 1 {
			return String(BadKey, x), nil
		}
		a := Any(x, args[1])
		a.Value = a.Value.Resolve()
		return a, args[2:]

	case Attr:
		x.Value = x.Value.Resolve()
		return x, args[1:]

	default:
		return Any(BadKey, x), args[1:]
	}
}
