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
