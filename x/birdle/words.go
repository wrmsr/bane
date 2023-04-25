package birdle

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"

	eu "github.com/wrmsr/bane/core/errors"
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

type WordFilter func(string) bool

func WordLenFilter(i int) WordFilter {
	return func(s string) bool {
		return len(s) == i
	}
}

func PrepareWord(word string, filter WordFilter) string {
	word = NormalizeWord(word)
	if werr := CheckWord(word); werr != nil {
		return ""
	}
	if filter != nil && !filter(word) {
		return ""
	}
	return word
}

func PrepareWords(s []string, filter WordFilter) []string {
	var r []string
	for _, word := range s {
		if word = PrepareWord(word, filter); word != "" {
			r = append(r, word)
		}
	}
	return r
}

//

const DefaultWordsFilePath = "/usr/share/dict/words"

func ReadWordsFile(path string, filter WordFilter) (s []string, err error) {
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

		if word := PrepareWord(string(buf), filter); word != "" {
			s = append(s, word)
		}
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

func (d *Dictionary) Random() string {
	return d.s[rand.Intn(len(d.s))]
}

//

type Mark int8

const (
	Normal Mark = iota
	Correct
	Misplaced
)

//

type MarkedWord interface {
	Word() string
	Mark(int) Mark
}

type markedWord struct {
	word  string
	marks []Mark
}

func (mw markedWord) Word() string    { return mw.word }
func (mw markedWord) Mark(i int) Mark { return mw.marks[i] }
