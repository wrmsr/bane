package strings

import (
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
)

type nl = NestedList

func TestMatchedPairs(t *testing.T) {
	tu.AssertDeepEqual(t,
		check.Must1(ParseNestedList("0(12(34,56),78)09", '(', ')', ',')),
		[]nl{"0", []nl{"12", []nl{"34", "56"}, "78"}, "09"})

	var err error

	_, err = ParseNestedList("0((1)", '(', ')', ',')
	tu.AssertEqual(t, err != nil, true)

	_, err = ParseNestedList("0(1))", '(', ')', ',')
	tu.AssertEqual(t, err != nil, true)
}
