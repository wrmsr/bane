package runtime

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/wrmsr/bane/core/check"
	tu "github.com/wrmsr/bane/core/dev/testing"
)

func TestParseGenericName(t *testing.T) {
	tn := "Optional[github.com/wrmsr/bane/core/container.Map[github.com/wrmsr/bane/core/optional.Optional[int],github.com/wrmsr/bane/core/optional.Optional[string]]]"
	pn := check.Must1(ParseGenericName(tn))
	tu.AssertEqual(t, pn.String(), tn)

	fmt.Println(string(check.Must1(json.Marshal(pn))))
}
