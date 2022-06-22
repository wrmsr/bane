/*
Copyright (c) 2013 John Barton

MIT License

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
documentation files (the "Software"), to deal in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit
persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the
Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
package dotenv

import (
	"bufio"
	"errors"
	"io"
	"regexp"
	"strings"
)

type Parser struct {
	m map[string]string
}

func NewParser() *Parser {
	return &Parser{
		m: make(map[string]string),
	}
}

func isIgnoredLine(line string) bool {
	trimmedLine := strings.TrimSpace(line)
	return len(trimmedLine) == 0 || strings.HasPrefix(trimmedLine, "#")
}

func (p *Parser) Parse(r io.Reader) error {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	for _, fullLine := range lines {
		if !isIgnoredLine(fullLine) {
			key, value, err := p.parseLine(fullLine)
			if err != nil {
				return err
			}

			p.m[key] = value
		}
	}
	return nil
}

var exportRegex = regexp.MustCompile(`^\s*(?:export\s+)?(.*?)\s*$`)

func (p *Parser) parseLine(line string) (key string, value string, err error) {
	if len(line) == 0 {
		err = errors.New("zero length string")
		return
	}

	if strings.Contains(line, "#") {
		segmentsBetweenHashes := strings.Split(line, "#")
		quotesAreOpen := false
		var segmentsToKeep []string
		for _, segment := range segmentsBetweenHashes {
			if strings.Count(segment, "\"") == 1 || strings.Count(segment, "'") == 1 {
				if quotesAreOpen {
					quotesAreOpen = false
					segmentsToKeep = append(segmentsToKeep, segment)
				} else {
					quotesAreOpen = true
				}
			}

			if len(segmentsToKeep) == 0 || quotesAreOpen {
				segmentsToKeep = append(segmentsToKeep, segment)
			}
		}

		line = strings.Join(segmentsToKeep, "#")
	}

	firstEquals := strings.Index(line, "=")
	firstColon := strings.Index(line, ":")
	splitString := strings.SplitN(line, "=", 2)
	if firstColon != -1 && (firstColon < firstEquals || firstEquals == -1) {
		splitString = strings.SplitN(line, ":", 2)
	}

	if len(splitString) != 2 {
		err = errors.New("can't separate key from value")
		return
	}

	key = splitString[0]
	if strings.HasPrefix(key, "export") {
		key = strings.TrimPrefix(key, "export")
	}
	key = strings.TrimSpace(key)

	key = exportRegex.ReplaceAllString(splitString[0], "$1")

	value = p.parseValue(splitString[1])
	return
}

var (
	singleQuotesRegex  = regexp.MustCompile(`\A'(.*)'\z`)
	doubleQuotesRegex  = regexp.MustCompile(`\A"(.*)"\z`)
	escapeRegex        = regexp.MustCompile(`\\.`)
	unescapeCharsRegex = regexp.MustCompile(`\\([^$])`)
)

func (p *Parser) parseValue(value string) string {
	value = strings.Trim(value, " ")

	if len(value) > 1 {
		singleQuotes := singleQuotesRegex.FindStringSubmatch(value)
		doubleQuotes := doubleQuotesRegex.FindStringSubmatch(value)

		if singleQuotes != nil || doubleQuotes != nil {
			value = value[1 : len(value)-1]
		}

		if doubleQuotes != nil {
			value = escapeRegex.ReplaceAllStringFunc(value, func(match string) string {
				c := strings.TrimPrefix(match, `\`)
				switch c {
				case "n":
					return "\n"
				case "r":
					return "\r"
				default:
					return match
				}
			})
			value = unescapeCharsRegex.ReplaceAllString(value, "$1")
		}

		if singleQuotes == nil {
			value = p.expandVariables(value)
		}
	}

	return value
}

var expandVarRegex = regexp.MustCompile(`(\\)?(\$)(\()?\{?([A-Z0-9_]+)?\}?`)

func (p *Parser) expandVariables(v string) string {
	return expandVarRegex.ReplaceAllStringFunc(v, func(s string) string {
		submatch := expandVarRegex.FindStringSubmatch(s)

		if submatch == nil {
			return s
		}
		if submatch[1] == "\\" || submatch[2] == "(" {
			return submatch[0][1:]
		} else if submatch[4] != "" {
			return p.m[submatch[4]]
		}
		return s
	})
}
