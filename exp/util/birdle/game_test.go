package birdle

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

func TestBirdle(t *testing.T) {
	d := NewDictionary(check.Must1(ReadWordsFile(DefaultWordsFilePath, WordLenFilter(5))))
	g := check.Must1(NewGame("CHIME", 3, d))
	r := TermRenderer{}

	check.Must1(g.Guess("CHEER"))
	fmt.Printf("%s\n", r.RenderGuess(g.guesses[0]))

	check.Must1(g.Guess("FROCK"))
	fmt.Printf("%s\n", r.RenderGuess(g.guesses[1]))
}
