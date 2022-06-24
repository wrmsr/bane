package strings

import (
	"bufio"
	"io"
	"strings"

	fnu "github.com/wrmsr/bane/pkg/util/funcs"
	opt "github.com/wrmsr/bane/pkg/util/optional"
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

	n := opt.None[int]()
	for _, l := range lines {
		if l == "" {
			continue
		}

		c := NumLeadingSpaces(l)
		if n.Present() {
			n = opt.Just(fnu.Min(c, n.Value()))
		} else {
			n = opt.Just(c)
		}
	}

	if !n.Present() || n.Value() == 0 {
		return s
	}

	var sb strings.Builder
	for _, l := range lines {
		if l != "" {
			sb.WriteString(l[n.Value():])
		}
		sb.WriteRune('\n')
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
