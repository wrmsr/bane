package birdle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	eu "github.com/wrmsr/bane/pkg/util/errors"
)

//

func NormalizeWord(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	return s
}

func CheckWord(s string) error {
	if s != NormalizeWord(s) {
		return fmt.Errorf("word not normalized")
	}
	if s == "" {
		return fmt.Errorf("word is empty")
	}

	var l rune
	for _, c := range s {
		if !(c >= 'A' && c <= 'Z') && c != ' ' {
			return fmt.Errorf("word contains invalid character: %v", c)
		}
		if l == ' ' && c == ' ' {
			return fmt.Errorf("word contains consecutive spaces")
		}
		l = c
	}

	return nil
}

//

func WordLenFilter(i int) func(string) bool {
	return func(s string) bool {
		return len(s) == i
	}
}

const DefaultWordsFilePath = "/usr/share/dict/words"

func ReadWordsFile(path string, filter func(string) bool) (s []string, err error) {
	var f *os.File
	f, err = os.Open(path)
	if err != nil {
		return nil, err
	}
	defer eu.AppendInvoke(&err, eu.Close(f))

	rd := bufio.NewReader(f)
	for {
		buf, _, rerr := rd.ReadLine()
		if rerr != nil {
			if rerr == io.EOF {
				break
			}
			err = rerr
			return
		}

		word := string(buf)
		word = NormalizeWord(word)
		if werr := CheckWord(word); werr != nil {
			continue
		}

		if filter != nil && !filter(word) {
			continue
		}

		s = append(s, word)
	}

	return
}

//

type Dictionary struct {
	s []string
	m map[string]struct{}
}

func NewDictionary(s []string) *Dictionary {
	m := make(map[string]struct{}, len(s))
	for _, w := range s {
		m[w] = struct{}{}
	}
	return &Dictionary{
		s: s,
		m: m,
	}
}

func (d *Dictionary) Len() int         { return len(d.s) }
func (d *Dictionary) Get(i int) string { return d.s[i] }

func (d *Dictionary) Contains(w string) bool {
	_, ok := d.m[w]
	return ok
}
