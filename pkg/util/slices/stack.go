package slices

type Stack[T any] struct {
	s []T
}

func (s *Stack[T]) Top() *T {
	return &s.s[len(s.s)-1]
}

func (s *Stack[T]) Push(t T) {
	s.s = append(s.s, t)
}

func (s *Stack[T]) Pop() T {
	l := len(s.s) - 1
	r := s.s[l]
	s.s = s.s[:l]
	return r
}

func (s *Stack[T]) Len() int {
	return len(s.s)
}
