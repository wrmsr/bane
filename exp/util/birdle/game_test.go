package birdle

import (
	"fmt"
	"strings"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

const (
	termReset  = "\u001b[0m"
	termGreen  = "\u001b[32m"
	termYellow = "\u001b[33m"
)

func termRenderGuess(g guess) string {
	var sb strings.Builder
	for i, c := range g.word {
		m := g.marks[i]
		switch m {
		case Correct:
			sb.WriteString(termGreen)
		case Misplaced:
			sb.WriteString(termYellow)
		}
		sb.WriteRune(c)
		if m != Normal {
			sb.WriteString(termReset)
		}
	}
	return sb.String()
}

func TestBirdle(t *testing.T) {
	g := check.Must1(NewGame("FARTS", 3))

	check.Equal(check.Must1(g.Guess("CHIME")), false)
	fmt.Printf("%s\n", termRenderGuess(g.guesses[0]))
	check.Equal(check.Must1(g.Guess("FROCK")), false)
	fmt.Printf("%s\n", termRenderGuess(g.guesses[1]))
}
