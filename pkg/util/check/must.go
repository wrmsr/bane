package check

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func Must1[T any](v T, err error) T {
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

func Ok(ok bool) {
	if !ok {
		panic("must succeed")
	}
}

func Ok1[T any](v T, ok bool) T {
	if !ok {
		panic("must succeed")
	}
	return v
}

func Ok2[T, U any](v T, v2 U, ok bool) (T, U) {
	if !ok {
		panic("must succeed")
	}
	return v, v2
}
