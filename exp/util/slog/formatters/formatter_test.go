package formatters

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"golang.org/x/exp/slog"
)

func TestHandlers(t *testing.T) {
	f := formatter{
		sty: jsonFormatterStyle{},
	}

	r := slog.NewRecord(
		time.Now(),
		slog.LevelInfo,
		"foo",
		0,
	)

	var b bytes.Buffer
	_ = f.format(r, &b)
	fmt.Println(b.String())
}
