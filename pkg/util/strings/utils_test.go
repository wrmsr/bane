package strings

import (
	"fmt"
	"strings"
	"testing"
	"unicode"

	"github.com/wrmsr/bane/pkg/util/check"
	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
)

func TestScanAllLines(t *testing.T) {
	tu.AssertDeepEqual(t, []string{"a", "b", "c"}, check.Must1(ScanAllLines(strings.NewReader("a\nb\nc"), true)))
}

func TestDedent(t *testing.T) {
	s := `
    abc
    def
        ghi

    jkl
`
	fmt.Println(Dedent(s))
}

func TestSplitFunc(t *testing.T) {
	tu.AssertDeepEqual(t, SplitFunc(unicode.IsUpper, "abcdEfgHij"), []string{"abcd", "Efg", "Hij"})
	tu.AssertDeepEqual(t, SplitFunc(unicode.IsUpper, "AbcdEfgHij"), []string{"Abcd", "Efg", "Hij"})
	tu.AssertDeepEqual(t, SplitFunc(unicode.IsUpper, "AbcdEfgHiJ"), []string{"Abcd", "Efg", "Hi", "J"})
}
