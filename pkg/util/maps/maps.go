package maps

//

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

func MapKeys[KF, KT comparable, V any](fn func(k KF) KT, m map[KF]V) map[KT]V {
	r := make(map[KT]V, len(m))
	for kf, v := range m {
		r[fn(kf)] = v
	}
	return r
}

func MapValues[K comparable, VF, VT any](fn func(v VF) VT, m map[K]VF) map[K]VT {
	r := make(map[K]VT, len(m))
	for k, v := range m {
		r[k] = fn(v)
	}
	return r
}

func FilterKeys[K comparable, V any](fn func(k K) bool, m map[K]V) map[K]V {
	r := make(map[K]V)
	for k, v := range m {
		if fn(k) {
			r[k] = v
		}
	}
	return r
}

func FilterValues[K comparable, V any](fn func(v V) bool, m map[K]V) map[K]V {
	r := make(map[K]V)
	for k, v := range m {
		if fn(v) {
			r[k] = v
		}
	}
	return r
}

//

func Clone[K comparable, V any](m map[K]V) map[K]V {
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

//

func Get[K comparable, V any](m map[K]V, k K) V {
	return m[k]
}

func TryGet[K comparable, V any](m map[K]V, k K) (V, bool) {
	v, ok := m[k]
	return v, ok
}

func Contains[K comparable, V any](m map[K]V, k K) bool {
	_, ok := m[k]
	return ok
}

func Put[K comparable, V any](m map[K]V, k K, v V) {
	m[k] = v
}
