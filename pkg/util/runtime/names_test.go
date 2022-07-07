package runtime

import (
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
)

func TestParseGenericName(t *testing.T) {
	tn := "Optional[github.com/wrmsr/bane/pkg/util/container.Map[github.com/wrmsr/bane/pkg/util/optional.Optional[int],github.com/wrmsr/bane/pkg/util/optional.Optional[string]]]"
	pn := check.Must1(ParseGenericName(tn))
	tu.AssertEqual(t, pn.String(), tn)
}
