package cfgfile

import (
	"bufio"
	"io"
)

const (
	sectionPatternTemplate = "" +
		"\\[" +
		"(?P<header>[^]]+)" +
		"\\]"

	optionPatternTemplate = "" +
		"(?P<option>.*?)" +
		// any number of space/tab, followed by any of the allowed delimiters, followed by any space/tab
		"\\s*(?P<vi>{delim})\\s*" +
		"(?P<value>.*)$"

	optionNoValuePatternTemplate = "" +
		"(?P<option>.*?)" +
		"\\s*(?:" +
		"(?P<vi>{delim})\\s*" + // optionally followed by any of the allowed delimiters, followed by any space/tab
		"(?P<value>.*))?$"

	defaultDelim = "=|:"
)

type Loader struct{}

func (l *Loader) read(r io.Reader) {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		sc.Text()
	}
}
