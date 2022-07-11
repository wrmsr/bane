package coalesce

func Cmp[T comparable](vs ...T) T {
	var z T
	for _, v := range vs {
		if v != z {
			return v
		}
	}
	return z
}

func CmpR[T comparable](vs ...T) T {
	var z T
	for i := len(vs) - 1; i >= 0; i-- {
		v := vs[i]
		if v != z {
			return v
		}
	}
	return z
}
