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

func (r Range[T]) StepOrOne() T {
	if r.Step == 0 {
		return 1
	}
	return r.Step
}

func (r Range[T]) Len() T {
	return (r.Stop - r.Start) / r.StepOrOne()
}

func (r Range[T]) HasNext() bool {
	if r.StepOrOne() >= 0 {
		return r.Start < r.Stop
	} else {
		return r.Start > r.Stop
	}
}

func (r Range[T]) Next() (Range[T], bool) {
	if !r.HasNext() {
		return Range[T]{}, false
	}
	r.Start += r.StepOrOne()
	return r, true
}

func (r Range[T]) Slice() []T {
	l := make([]T, 0, int((r.Stop-r.Start)/r.StepOrOne()))
	for ; r.Start < r.Stop; r.Start += r.StepOrOne() {
		l = append(l, r.Start)
	}
	return l
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
	return i.r.HasNext()
}

func (i *rangeIterator[T]) Next() T {
	c := i.r.Start
	n, ok := i.r.Next()
	if !ok {
		panic(IteratorExhaustedError{})
	}
	i.r = n
	return c
}

var _ Iterable[int] = Range[int]{}

func (r Range[T]) Iterate() Iterator[T] {
	return &rangeIterator[T]{r: r}
}

//

var _ AnyIterable = Range[int]{}

func (r Range[T]) AnyIterate() Iterator[any] {
	return &anyIterator[T]{it: r.Iterate()}
}

//

var _ Traversable[int] = Range[int]{}

func (r Range[T]) ForEach(fn func(T) bool) bool {
	for i := r.Start; i < r.Stop; i += r.StepOrOne() {
		if !fn(i) {
			return false
		}
	}
	return true
}
