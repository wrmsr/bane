package slices

type Stack[T any] []T

func (s *Stack[T]) Top() *T {
	return &(*s)[len(*s)-1]
}

func (s *Stack[T]) Push(t T) {
	*s = append(*s, t)
}

func (s *Stack[T]) Pop() T {
	l := len(*s) - 1
	r := (*s)[l]
	*s = (*s)[:l]
	return r
}
