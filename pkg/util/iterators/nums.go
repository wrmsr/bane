package iterators

import bt "github.com/wrmsr/bane/pkg/util/types"

//

type countIterator struct {
	n int
}

func Count() bt.Iterable[int] {
	return Factory(func() bt.Iterator[int] { return &countIterator{} }, nil)
}

var _ bt.Iterator[int] = &countIterator{}

func (i *countIterator) Iterate() bt.Iterator[int] {
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

func Range(start, stop, step int) bt.Iterable[int] {
	return Factory(func() bt.Iterator[int] {
		return &rangeIterator{
			n:    start,
			stop: stop,
			step: step,
		}
	}, nil)
}

func Range1(start, stop int) bt.Iterable[int] {
	return Range(start, stop, 1)
}

var _ bt.Iterator[int] = &rangeIterator{}

func (i *rangeIterator) Iterate() bt.Iterator[int] {
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
		panic(bt.IteratorExhaustedError{})
	}
	n := i.n
	i.n += i.step
	return n
}
