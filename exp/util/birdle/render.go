package birdle

import (
	"fmt"
	"strings"
)

//

type Renderer interface {
	RenderGuess(Guess) string
	RenderGame(*Game) string
}

//

const (
	termReset  = "\u001b[0m"
	termGreen  = "\u001b[32m"
	termYellow = "\u001b[33m"
)

type TermRenderer struct{}

var _ Renderer = TermRenderer{}

func (r TermRenderer) RenderGuess(ge Guess) string {
	var sb strings.Builder
	for i, c := range ge.word {
		m := ge.marks[i]
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

func (r TermRenderer) RenderGame(g *Game) string {
	var sb strings.Builder

	gm := g.GuessesMade()
	if gm > 0 {
		sb.WriteString(r.RenderGuess(g.GetGuess(gm - 1)))
	} else {
		l := len(g.Word())
		for i := 0; i < l; i++ {
			sb.WriteRune('_')
		}
	}
	sb.WriteRune(' ')

	sb.WriteRune('(')
	sb.WriteString(fmt.Sprintf("%d/%d", gm, g.GuessesLeft()))
	sb.WriteRune(' ')
	for c := 'A'; c <= 'Z'; c++ {
		if g.HasGuessedChar(c) {
			sb.WriteRune('_')
		} else {
			sb.WriteRune(c)
		}
	}
	sb.WriteRune(')')
	return sb.String()
}
