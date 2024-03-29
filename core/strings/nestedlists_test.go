package strings

import (
	"testing"

	"github.com/wrmsr/bane/core/check"
	tu "github.com/wrmsr/bane/core/dev/testing"
)

func TestNestedLists(t *testing.T) {
	tu.AssertDeepEqual(t,
		check.Must1(ParseNestedList("0(12(34,56),78)09", '(', ')', ',')),
		[]any{"0", []any{"12", []any{"34", "56"}, "78"}, "09"})

	var err error

	_, err = ParseNestedList("0((1)", '(', ')', ',')
	tu.AssertEqual(t, err != nil, true)

	_, err = ParseNestedList("0(1))", '(', ')', ',')
	tu.AssertEqual(t, err != nil, true)

	tu.AssertDeepEqual(t,
		check.Must1(ParseNestedList("Optional[github.com/wrmsr/bane/core/container.Map[int,github.com/wrmsr/bane/core/optional.Optional[string]]]", '[', ']', ',')),
		[]any{"Optional", []any{"github.com/wrmsr/bane/core/container.Map", []any{"int", "github.com/wrmsr/bane/core/optional.Optional", []any{"string"}}}})
}
