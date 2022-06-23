package strings

import (
	"fmt"
	"strings"
	"testing"

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
