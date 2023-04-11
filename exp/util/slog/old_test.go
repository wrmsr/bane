package slog

import (
	old "log"
	"testing"
)

func TestOld(t *testing.T) {
	//stdslog.SetDefault(stdslog.New(stdslog.HandlerOptions{AddSource: true}.NewJSONHandler(os.Stderr)))

	o := old.Default()
	InstallOldAdapter(o, Default(), LevelInfo)
	o.SetFlags(0)

	old.Println("hi")
}
