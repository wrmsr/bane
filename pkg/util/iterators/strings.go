package iterators

import (
	"unicode/utf8"

	bt "github.com/wrmsr/bane/pkg/util/types"
)

type stringIterator struct {
	s string
	n int
}

func OfString(s string) bt.Iterable[rune] {
	return Factory(func() bt.Iterator[rune] {
		return &stringIterator{s: s}
	}, s)
}

var _ bt.Iterator[rune] = &stringIterator{}

func (i *stringIterator) Iterate() bt.Iterator[rune] {
	return i
}

func (i *stringIterator) HasNext() bool {
	return i.n < len(i.s)
}

func (i *stringIterator) Next() rune {
	r, n := utf8.DecodeRuneInString(i.s[i.n:])
	i.n += n
	return r
}
