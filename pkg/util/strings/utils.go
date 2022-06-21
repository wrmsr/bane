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
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return a, scanner.Err()
}
