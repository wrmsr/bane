package slices

func Reshape[T any](s []T, start, stop, step int) []T {
	if step == 0 {
		panic("zero step")
	}

	l := len(s)

	ci := func(i int) int {
		if i < 0 {
			i += l
			if i < 0 {
				i = 0
			}
		}
		if i > l {
			return l
		}
		return i
	}
	start = ci(start)
	stop = ci(stop)

	if step == 1 {
		return s[start:stop]
	}

	rnd := 0
	if step < 0 {
		rnd = step + 1
	} else {
		rnd = step - 1
	}
	rl := (stop - start + rnd) / step
	if rl < 1 {
		return nil
	}

	r := make([]T, rl)
	if step > 0 {
		for i, j := start, 0; i < stop; i, j = i+step, j+1 {
			r[j] = s[i]
		}
	} else {
		for i, j := start, 0; i > stop; i, j = i+step, j+1 {
			r[j] = s[i]
		}
	}
	return r
}
