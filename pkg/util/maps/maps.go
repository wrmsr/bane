package maps

func Keys[K comparable, V any](m map[K]V) []K {
	s := make([]K, len(m))
	i := 0
	for k := range m {
		s[i] = k
		i++
	}
	return s
}

func Values[K comparable, V any](m map[K]V) []V {
	s := make([]V, len(m))
	i := 0
	for _, v := range m {
		s[i] = v
		i++
	}
	return s
}

func Copy[K comparable, V any](m map[K]V) map[K]V {
	r := make(map[K]V, len(m))
	for k, v := range m {
		r[k] = v
	}
	return r
}

func Invert[K, V comparable](m map[K]V) map[V]K {
	r := make(map[V]K, len(m))
	for k, v := range m {
		r[v] = k
	}
	return r
}
