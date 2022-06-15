package strings

import (
	"bufio"
	"io"
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
	return a, scanner.Err()
}
