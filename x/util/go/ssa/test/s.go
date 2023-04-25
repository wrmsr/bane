package test

func S(o any) int {
	switch o := o.(type) {
	case int:
		return 0
	case string:
		return 1
	default:
		panic(o)
	}
}
