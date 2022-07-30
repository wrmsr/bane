package slices

func Find[T comparable](s []T, v T) (int, bool) {
	for i, c := range s {
		if c == v {
			return i, true
		}
	}
	return -1, false
}

func FindLast[T comparable](s []T, v T) (int, bool) {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == v {
			return i, true
		}
	}
	return -1, false
}

func FindFunc[T any](s []T, fn func(T) bool) (int, bool) {
	for i, c := range s {
		if fn(c) {
			return i, true
		}
	}
	return -1, false
}

func FindElemFunc[T any](s []T, fn func(T) bool) (T, bool) {
	for _, c := range s {
		if fn(c) {
			return c, true
		}
	}
	var z T
	return z, false
}

func FindFuncLast[T any](s []T, fn func(T) bool) (int, bool) {
	for i := len(s) - 1; i >= 0; i-- {
		if fn(s[i]) {
			return i, true
		}
	}
	return -1, false
}

func Any[T any](s []T, fn func(T) bool) bool {
	_, ok := FindFunc(s, fn)
	return ok
}

func All[T any](s []T, fn func(T) bool) bool {
	for _, c := range s {
		if !fn(c) {
			return false
		}
	}
	return true
}

func Contains[T comparable](s []T, v T) bool {
	_, ok := Find(s, v)
	return ok
}
