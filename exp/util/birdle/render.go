package birdle

import "strings"

//

type Renderer interface {
	RenderGuess(Guess) string
}

//

const (
	termReset  = "\u001b[0m"
	termGreen  = "\u001b[32m"
	termYellow = "\u001b[33m"
)

type TermRenderer struct{}

var _ Renderer = TermRenderer{}

func (r TermRenderer) RenderGuess(guess Guess) string {
	var sb strings.Builder
	for i, c := range guess.word {
		m := guess.marks[i]
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
