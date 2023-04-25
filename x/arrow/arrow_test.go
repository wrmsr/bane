package arrow

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/apache/arrow/go/v10/arrow/flight"

	"github.com/wrmsr/bane/core/check"
	tu "github.com/wrmsr/bane/core/dev/testing"
	osu "github.com/wrmsr/bane/core/os"
)

func TestArrow(t *testing.T) {
	fp := filepath.Join(check.Must1(os.UserHomeDir()), ".cache/huggingface/datasets/squad/plain_text/1.0.0/d6ec3ceb99ca480ce37cdd35555d6cb2511d223b9150cce08a837ef62ffea453/squad-train.arrow")
	tu.AssertEqual(t, check.Must1(osu.Exists(fp)), true)

	_, _ = flight.NewRecordReader(nil)
}
