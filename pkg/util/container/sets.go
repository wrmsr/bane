package container

func Union[T comparable](ss ...Set[T]) Set[T] {
	r := NewMutSet[T](nil)
	for _, s := range ss {
		s.ForEach(func(v T) bool {
			r.Add(v)
			return true
		})
	}
	return r
}

func Intersection[T comparable](ss ...Set[T]) Set[T] {
	if len(ss) < 1 {
		return NewSet[T](nil)
	}
	cts := make(map[T]int)
	for _, s := range ss[1:] {
		s.ForEach(func(v T) bool {
			cts[v] = cts[v] + 1
			return true
		})
	}
	r := NewMutSet[T](nil)
	for v, ct := range cts {
		if ct == len(ss) {
			r.Add(v)
		}
	}
	return r
}

func Difference[T comparable](ss ...Set[T]) Set[T] {
	if len(ss) < 1 {
		return NewSet[T](nil)
	}
	r := NewMutSet[T](ss[0])
	for _, s := range ss[1:] {
		s.ForEach(func(v T) bool {
			r.Remove(v)
			return true
		})
	}
	return r
}
