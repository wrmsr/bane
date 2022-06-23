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
