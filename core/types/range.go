package types

import (
	"errors"
)

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
	step := r.StepOrOne()
	l := make([]T, 0, int((r.Stop-r.Start)/step))
	for ; r.Start < r.Stop; r.Start += step {
		l = append(l, r.Start)
	}
	return l
}

//

type RangeError[T Rational] struct {
	Range Range[T]
	Err   error
}

func (e RangeError[T]) Error() string {
	return e.Err.Error()
}

func (e RangeError[T]) Unwrap() error {
	return e.Err
}

func (r Range[T]) CheckNormal(n T) Range[T] {
	if r.Step == 0 {
		panic(RangeError[T]{Err: errors.New("step == 0"), Range: r})
	}
	if r.Step > 0 {
		if r.Start < 0 {
			panic(RangeError[T]{Err: errors.New("start < 0"), Range: r})
		}
		if r.Stop > n {
			panic(RangeError[T]{Err: errors.New("stop > n"), Range: r})
		}
		if r.Stop < r.Start {
			panic(RangeError[T]{Err: errors.New("stop < start"), Range: r})
		}
	} else {
		if r.Start > n {
			panic(RangeError[T]{Err: errors.New("start > n"), Range: r})
		}
		if r.Stop < 0 {
			panic(RangeError[T]{Err: errors.New("stop < 0"), Range: r})
		}
		if r.Stop > r.Start {
			panic(RangeError[T]{Err: errors.New("stop > start"), Range: r})
		}
	}
	return r
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
	step := r.StepOrOne()
	for i := r.Start; i < r.Stop; i += step {
		if !fn(i) {
			return false
		}
	}
	return true
}

//

func (r Range[T]) AsSlice() []T {
	return []T{
		r.Start,
		r.Stop,
		r.Step,
	}
}

//

type rangeAnySlice[T Rational] struct {
	r Range[T]
}

func (r Range[T]) AsAnySlice() AnySlice {
	return rangeAnySlice[T]{r: r}
}

var _ AnySlice = rangeAnySlice[int]{}

func (s rangeAnySlice[T]) Iterate() Iterator[any] {
	return IterateAnySlice(s)
}

func (s rangeAnySlice[T]) ForEach(fn func(v any) bool) bool {
	return fn(s.r.Start) && fn(s.r.Stop) && fn(s.r.Step)
}

func (s rangeAnySlice[T]) Len() int {
	return 3
}

func (s rangeAnySlice[T]) Get(i int) any {
	switch i {
	case 0:
		return s.r.Start
	case 1:
		return s.r.Stop
	case 2:
		return s.r.Step
	}
	panic(errors.New("index out of range"))
}
