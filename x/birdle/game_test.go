package birdle

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/core/check"
)

func TestBirdle(t *testing.T) {
	d := NewDictionary(check.Must1(ReadWordsFile(DefaultWordsFilePath, WordLenFilter(5))))
	g := check.Must1(NewGame("CHIME", 3, d))
	r := TermRenderer{}

	check.Must(g.Guess("CHEER"))
	fmt.Printf("%s\n", r.RenderGame(g))

	check.Must(g.Guess("FROCK"))
	fmt.Printf("%s\n", r.RenderGame(g))
}
