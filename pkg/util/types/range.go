package types

//

type Range[T Rational] struct {
	Start, Stop, Step T
}

func RangeOf[T Rational](start, stop, step T) Range[T] {
	return Range[T]{Start: start, Stop: stop, Step: step}
}

func RangeTo[T Rational](stop T) Range[T] {
	return Range[T]{Stop: stop, Step: 1}
}

func (r Range[T]) Len() T {
	return (r.Stop - r.Start) / r.Step
}

//

type rangeIterator[T Rational] struct {
	r Range[T]
}

var _ Iterator[int] = &rangeIterator[int]{}

func (i *rangeIterator[T]) Iterate() Iterator[T] {
	return i
}

func (i rangeIterator[T]) HasNext() bool {
	if i.r.Step >= 0 {
		return i.r.Start < i.r.Stop
	} else {
		return i.r.Start > i.r.Stop
	}
}

func (i *rangeIterator[T]) Next() T {
	if !i.HasNext() {
		panic(IteratorExhaustedError{})
	}
	n := i.r.Start
	i.r.Start += i.r.Step
	return n
}

//

var _ Traversable[int] = Range[int]{}

func (r Range[T]) ForEach(fn func(T) bool) bool {
	for i := r.Start; i < r.Stop; i += r.Step {
		if !fn(i) {
			return false
		}
	}
	return true
}
