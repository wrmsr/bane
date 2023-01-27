package main

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

const (
	termReset  = "\u001b[0m"
	termGreen  = "\u001b[32m"
	termYellow = "\u001b[33m"
)



func TestBirdle(t *testing.T) {
	g := check.Must1(NewGame("FARTS", 3))

	check.Equal(check.Must1(g.Guess("CHIME")), false)
	check.Equal(check.Must1(g.Guess("FROCK")), false)

	fmt.Printf("hi %sgreen%s huh %syellow%s bye\n", termGreen, termReset, termYellow, termReset)
}
