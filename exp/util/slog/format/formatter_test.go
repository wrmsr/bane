package format_test

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"golang.org/x/exp/slog"

	slogf "github.com/wrmsr/bane/exp/util/slog/format"
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
