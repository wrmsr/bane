package check

func IgnoreErr(err error) {}

func IgnoreErr1[T any](v T, err error) T {
	return v
}

func IgnoreErr2[T, U any](v T, v2 U, err error) (T, U) {
	return v, v2
}

func IgnoreOk(bool bool) {}

func IgnoreOk1[T any](v T, ok bool) T {
	return v
}

func IgnoreOk2[T, U any](v T, v2 U, ok bool) (T, U) {
	return v, v2
}
