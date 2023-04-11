package slog

import (
	"fmt"
	old "log"
	"os"
	"testing"

	stdslog "golang.org/x/exp/slog"

	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
)

func TestOldDetect(t *testing.T) {
	tu.AssertEqual(t, isSlogDefaultHandler(stdslog.Default().Handler()), true)

	h2 := stdslog.New(stdslog.HandlerOptions{AddSource: true}.NewJSONHandler(os.Stderr))
	tu.AssertEqual(t, isSlogDefaultHandler(h2.Handler()), false)
}

func TestOld(t *testing.T) {
	//stdslog.SetDefault(stdslog.New(stdslog.HandlerOptions{AddSource: true}.NewJSONHandler(os.Stderr)))

	o := old.Default()
	fmt.Printf("%#v\n", o)

	InstallOldAdapter(o, Default(), LevelInfo)
	o.SetFlags(0)

	old.Println("hi")
}
