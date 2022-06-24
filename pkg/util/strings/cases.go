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
	//uppers := slices.IndexMap()
	//uppers: ta.List[ta.Optional[int]] = [i for i, c in enumerate(name) if c.isupper()]
	//return '_'.join([name[l:r].lower() for l, r in zip([None] + uppers, uppers + [None])]).strip('_')  # type: ignore
	panic("nyi")
}
