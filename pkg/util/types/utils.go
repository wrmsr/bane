package types

func Is[I any](v any) bool {
	_, ok := v.(I)
	return ok
}
