package strings

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	fnu "github.com/wrmsr/bane/core/funcs"
	"github.com/wrmsr/bane/core/slices"
	bt "github.com/wrmsr/bane/core/types"
)

func ScanAllLines(r io.Reader, skipEmpty bool) ([]string, error) {
	scanner := bufio.NewScanner(r)
	var a []string
	for scanner.Scan() {
		s := scanner.Text()
		if skipEmpty && s == "" {
			continue
		}
		a = append(a, s)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return a, scanner.Err()
}

func NumLeadingSpaces(s string) int {
	i := 0
	for _, c := range s {
		if c != ' ' {
			return i
		}
		i++
	}
	return i
}

func Dedent(s string) string {
	lines := strings.Split(s, "\n")

	n := bt.None[int]()
	for _, l := range lines {
		if l == "" {
			continue
		}

		c := NumLeadingSpaces(l)
		if n.Present() {
			n = bt.Just(fnu.Min(c, n.Value()))
		} else {
			n = bt.Just(c)
		}
	}

	if !n.Present() || n.Value() == 0 {
		return s
	}

	var sb strings.Builder
	for i, l := range lines {
		if l != "" {
			sb.WriteString(l[n.Value():])
		}
		if i < len(l)-1 {
			sb.WriteRune('\n')
		}
	}
	return sb.String()
}

//

type StrFn = func(string) string

func Trimmer(cutset string) StrFn {
	return func(s string) string {
		return strings.Trim(s, cutset)
	}
}

func MapSplit(s string, fn StrFn, sep string) []string {
	var r []string
	for _, p := range strings.Split(fn(s), sep) {
		if p = fn(p); p != "" {
			r = append(r, p)
		}
	}
	return r
}

func TrimSplit(s, cutset, sep string) []string {
	return MapSplit(s, Trimmer(cutset), sep)
}

func TrimSpaceSplit(s, sep string) []string {
	return MapSplit(s, strings.TrimSpace, sep)
}

func MapCut(s string, fn StrFn, sep string) (before, after string, found bool) {
	b, a, found := strings.Cut(fn(s), sep)
	if !found {
		return
	}
	before = fn(b)
	after = fn(a)
	return
}

func TrimCut(s, cutset, sep string) (before, after string, found bool) {
	return MapCut(s, Trimmer(cutset), sep)
}

func TrimSpaceCut(s, sep string) (before, after string, found bool) {
	return MapCut(s, strings.TrimSpace, sep)
}

func LastCut(s, sep string) (before, after string, found bool) {
	if i := strings.LastIndex(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return "", s, false
}

func Append(l, r string) string {
	return l + r
}

func AppendNewline(l string) string {
	return l + "\n"
}

//

func IndexAllFunc(s string, f func(rune) bool) []int {
	var r []int
	for i, c := range []rune(s) {
		if f(c) {
			r = append(r, i)
		}
	}
	return r
}

func SplitFunc(f func(rune) bool, s string) []string {
	var ret []string
	rs := []rune(s)
	l := 0
	for i, r := range rs {
		if f(r) {
			if l != i {
				ret = append(ret, string(rs[l:i]))
			}
			l = i
		}
	}
	if l != len(rs) {
		ret = append(ret, string(rs[l:]))
	}
	return ret
}

func SplitAny(ds []rune, s string) []string {
	return SplitFunc(fnu.Bind1x1x1(slices.Contains[rune], ds), s)
}

//

func ToString(o any) string {
	return fmt.Sprintf("%s", o)
}

func ToValue(o any) string {
	return fmt.Sprintf("%v", o)
}
