package strings

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/wrmsr/bane/pkg/util/slices"
)

func ToCamel(s string) string {
	return strings.Join(slices.Map(cases.Title(language.Und).String, strings.Split(s, "_")), "")
}

func ToSnake(s string) string {
	var sb strings.Builder
	for i := 0; i < len(s); {
		r, n := utf8.DecodeRuneInString(s[i:])
		if unicode.IsUpper(r) {
			if i > 0 {
				sb.WriteRune('_')
			}
			sb.WriteRune(unicode.ToLower(r))
		} else {
			sb.WriteRune(r)
		}
		i += n
	}
	return sb.String()
}
