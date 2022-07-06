package strings

import "testing"

func TestMatchedPairs(t *testing.T) {
	_, _ = ParseNestedList("0(12(34,56),78)", '(', ')', ',')
}
