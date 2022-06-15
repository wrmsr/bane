package check

func NotNil[T any](v T, s ...string) T {
	if v == nil {
		if len(s) > 0 {
			panic(s[0])
		}
		panic("cannot be nil")
	}
	return v
}
