package test

func M[K comparable, V any](m map[K]V) {
	i := 0
	for k := range m {
		if i%2 == 0 {
			delete(m, k)
		}
		i++
	}
}

func M_int_string(m map[int]string) {
	M(m)
}
