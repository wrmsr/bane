package check

func NoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func Must2[T, U any](v T, v2 U, err error) (T, U) {
	if err != nil {
		panic(err)
	}
	return v, v2
}
