package format_test

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/wrmsr/bane/pkg/util/slog"
	slogf "github.com/wrmsr/bane/pkg/util/slog/format"
)

func TestHandlers(t *testing.T) {
	f := slogf.NewJsonFormatter(slogf.FormatterOpts{})

	r := slog.NewRecord(
		time.Now(),
		slog.LevelInfo,
		"foo",
		0,
	)

	var b bytes.Buffer
	_ = f.Format(r, (&b).Write)
	fmt.Println(b.String())
}
