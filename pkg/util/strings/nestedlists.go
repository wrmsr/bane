package strings

import (
	"errors"
	"unicode/utf8"

	fnu "github.com/wrmsr/bane/pkg/util/funcs"
	"github.com/wrmsr/bane/pkg/util/slices"
)

//

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

//

type NestedList = any // string | []NestedList

func ParseNestedList(s string, left, right, sep rune) (NestedList, error) {
	if s == "" {
		return "", nil
	}
	ctrl := []rune{left, right, sep}
	parts := slices.FlatMap(fnu.Bind1x1x1(SplitAny, ctrl), SplitAny(ctrl, s))

	var stk = Stack[[]NestedList]{}
	stk.Push(nil)

	add := func(o NestedList) error {
		if stk.Len() < 1 {
			return errors.New("too many right delimiters")
		}
		*stk.Top() = append(*stk.Top(), o)
		return nil
	}

	for i, c := range parts {
		u, x := utf8.DecodeRuneInString(c)
		if !slices.Contains(ctrl, u) {
			if i == 0 {
				if err := add(c); err != nil {
					return nil, err
				}
				continue
			}
			panic("unreachable")
		}
		c = c[x:]

		switch u {
		case left:
			stk.Push(nil)
		case right:
			l := stk.Pop()
			var a NestedList = l
			if len(l) == 1 {
				a = l[0]
			}
			if err := add(a); err != nil {
				return nil, err
			}
		case sep:
		default:
			panic("unreachable")
		}

		if c != "" {
			if err := add(c); err != nil {
				return nil, err
			}
		}
	}

	if stk.Len() != 1 {
		return nil, errors.New("too many left delimiters")
	}

	ret := stk.Pop()
	return ret, nil
}
