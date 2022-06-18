package strings

import (
	"strings"
	"testing"

	"github.com/wrmsr/bane/pkg/utils/check"
	tu "github.com/wrmsr/bane/pkg/utils/dev/testing"
)

func TestScanAllLines(t *testing.T) {
	tu.AssertDeepEqual(t, []string{"a", "b", "c"}, check.Must(ScanAllLines(strings.NewReader("a\nb\nc"), true)))
}
