package strings

import (
	"strings"
	"testing"

	bts "github.com/wrmsr/bane/pkg/dev/testing"
	"github.com/wrmsr/bane/pkg/utils/check"
)

func TestScanAllLines(t *testing.T) {
	bts.AssertDeepEqual(t, []string{"a", "b", "c"}, check.Must(ScanAllLines(strings.NewReader("a\nb\nc"), true)))
}
