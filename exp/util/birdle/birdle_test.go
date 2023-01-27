package main

import (
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

func TestBirdle(t *testing.T) {
	g := check.Must1(NewGame("FARTS", 3))
	check.Equal(check.Must1(g.Guess("CHIME")), false)
}
