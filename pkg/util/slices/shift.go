package slices

func ShiftAs[E any](s *[]any) (E, bool) {
	if len(*s) > 0 {
		if e, ok := (*s)[0].(E); ok {
			*s = (*s)[1:]
			return e, true
		}
	}
	var z E
	return z, false
}

func TryShiftAs[E any](s *[]any) E {
	e, _ := ShiftAs[E](s)
	return e
}
