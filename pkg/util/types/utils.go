package types

func Is[I any](v any) bool {
	_, ok := v.(I)
	return ok
}

func Zero[T any]() T {
	var z T
	return z
}
