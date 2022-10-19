package strings

import (
	"errors"
	"unicode/utf8"

	fnu "github.com/wrmsr/bane/pkg/util/funcs"
	"github.com/wrmsr/bane/pkg/util/slices"
)

func ParseNestedList(s string, left, right, sep rune) ([]any, error) {
	if s == "" {
		return nil, nil
	}
	ctrl := []rune{left, right, sep}
	parts := slices.FlatMap(fnu.Bind1x1x1(SplitAny, ctrl), SplitAny(ctrl, s))

	var stk = slices.Stack[[]any]{}
	stk.Push(nil)

	add := func(o any) error {
		if len(stk) < 1 {
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
			if err := add(stk.Pop()); err != nil {
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

	if len(stk) != 1 {
		return nil, errors.New("too many left delimiters")
	}

	ret := stk.Pop()
	return ret, nil
}
