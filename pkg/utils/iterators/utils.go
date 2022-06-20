package iterators

func Seq[T any](it Iterable[T]) []T {
	var r []T
	for i := it.Iterate(); i.HasNext(); {
		r = append(r, i.Next())
	}
	return r
}

func Take[T any](it Iterable[T], n int) []T {
	var r []T
	c := 0
	for i := it.Iterate(); i.HasNext(); {
		if c >= n {
			break
		}
		r = append(r, i.Next())
		c++
	}
	return r
}
