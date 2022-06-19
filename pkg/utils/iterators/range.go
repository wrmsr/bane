package iterators

type rangeIterator struct {
	n, stop, step int
}

func Range(start, stop, step int) Iterator[int] {
	return &rangeIterator{
		n:    start,
		stop: stop,
		step: step,
	}
}

var _ Iterator[int] = &rangeIterator{}

func (i *rangeIterator) Iterate() Iterator[int] {
	return i
}

func (i rangeIterator) HasNext() bool {
	return i.n < i.stop
}

func (i *rangeIterator) Next() int {
	if !i.HasNext() {
		panic(IteratorExhaustedError{})
	}
	n := i.n
	i.n += i.step
	return n
}
