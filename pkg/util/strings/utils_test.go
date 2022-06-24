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
	tu.AssertDeepEqual(t, []string{"a", "b", "c"}, check.Must(ScanAllLines(strings.NewReader("a\nb\nc"), true)))
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
	tu.AssertDeepEqual(t, SplitFunc("abcdEfgHij", unicode.IsUpper), []string{"abcd", "Efg", "Hij"})
	tu.AssertDeepEqual(t, SplitFunc("AbcdEfgHij", unicode.IsUpper)
	tu.AssertDeepEqual(t, SplitFunc("AbcdEfgHiJ", unicode.IsUpper)
}
