package iterators

//

type countIterator struct {
	n int
}

func Count() Iterable[int] {
	return factory(func() Iterator[int] { return &countIterator{} })
}

var _ Iterator[int] = &countIterator{}

func (i *countIterator) Iterate() Iterator[int] {
	return i
}

func (i countIterator) HasNext() bool {
	return true
}

func (i *countIterator) Next() int {
	n := i.n
	i.n++
	return n
}

//

type rangeIterator struct {
	n, stop, step int
}

func Range(start, stop, step int) Iterable[int] {
	return factory(func() Iterator[int] {
		return &rangeIterator{
			n:    start,
			stop: stop,
			step: step,
		}
	})
}

var _ Iterator[int] = &rangeIterator{}

func (i *rangeIterator) Iterate() Iterator[int] {
	return i
}

func (i rangeIterator) HasNext() bool {
	if i.step >= 0 {
		return i.n < i.stop
	} else {
		return i.n > i.stop
	}
}

func (i *rangeIterator) Next() int {
	if !i.HasNext() {
		panic(IteratorExhaustedError{})
	}
	n := i.n
	i.n += i.step
	return n
}
