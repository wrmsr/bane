package iterators

type countingIterator struct {
	n int
}

func Count() Iterator[int] {
	return countingIterator{}
}

func (i countingIterator) HasNext() bool {
	return true
}

func (i countingIterator) Next() int {
	n := i.n
	i.n++
	return n
}
