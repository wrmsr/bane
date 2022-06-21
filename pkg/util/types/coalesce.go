package types

func Any[T comparable](vs ...T) T {
	var z T
	for _, v := range vs {
		if v != z {
			return v
		}
	}
	return z
}
