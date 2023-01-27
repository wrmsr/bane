package birdle

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

func TestBirdle(t *testing.T) {
	g := check.Must1(NewGame("FARTS", 3))
	r := TermRenderer{}

	check.Must1(g.Guess("CHIME"))
	fmt.Printf("%s\n", r.RenderGuess(g.guesses[0]))

	check.Must1(g.Guess("FROCK"))
	fmt.Printf("%s\n", r.RenderGuess(g.guesses[1]))
}
