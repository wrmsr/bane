package types

func Is[I any](v any) bool {
	_, ok := v.(I)
	return ok
}

func Zero[T any]() T {
	var z T
	return z
}

func PtrTo[T any](v T) *T {
	return &v
}

func BoolNum[T Rational](b bool) T {
	if b {
		return 1
	}
	return 0
}
