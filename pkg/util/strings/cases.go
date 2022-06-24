package strings

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/wrmsr/bane/pkg/util/slices"
)

func ToCamel(s string) string {
	return strings.Join(slices.Map(cases.Title(language.Und).String, strings.Split(s, "_")), "")
}

func ToSnake(s string) string {
	//uppers := IndexAllFunc(s, unicode.IsUpper)
	var sb strings.Builder
	//for i := 1; i < len(uppers); i++ {
	//	sb.WriteString("_")
	//	sb.WriteString(s[])
	//}
	//return '_'.join([name[l:r].lower() for l, r in zip([None] + uppers, uppers + [None])]).strip('_')  # type: ignore
	return sb.String()
}
