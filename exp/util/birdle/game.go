//
/*
TODO:
 - bitmask chars
*/
package birdle

import (
	"fmt"
)

//

type State int8

const (
	Running State = iota
	Lost
	Won
)

func (s State) String() string {
	switch s {
	case Running:
		return "running"
	case Lost:
		return "lost"
	case Won:
		return "won"
	}
	panic(s)
}

//

type Mark int8

const (
	Normal Mark = iota
	Correct
	Misplaced
)

//

type guess struct {
	word  string
	marks []Mark
}

type Game struct {
	word string

	state   State
	guesses []guess

	wordChars   []rune
	wordCharSet map[rune]struct{}

	guessedCharSet map[rune]struct{}
}

func NewGame(word string, guessesAllowed int) (*Game, error) {
	if err := CheckWord(word); err != nil {
		return nil, err
	}
	if guessesAllowed < 1 {
		return nil, fmt.Errorf("invalid number of guesses: %v", guessesAllowed)
	}

	wcs := make([]rune, len(word))
	wcss := make(map[rune]struct{})
	for i, c := range word {
		wcs[i] = c
		wcss[c] = struct{}{}
	}

	return &Game{
		word: word,

		guesses: make([]guess, 0, guessesAllowed),

		wordChars:   wcs,
		wordCharSet: wcss,

		guessedCharSet: make(map[rune]struct{}),
	}, nil
}

func (g *Game) State() State { return g.state }

func (g *Game) GuessesMade() int    { return len(g.guesses) }
func (g *Game) GuessesAllowed() int { return cap(g.guesses) }
func (g *Game) GuessesLeft() int    { return cap(g.guesses) - len(g.guesses) }

func (g *Game) CheckGuess(word string) error {
	if err := CheckWord(word); err != nil {
		return err
	}
	if len(word) != len(g.word) {
		return fmt.Errorf("invalid guess length: guess %v != word %v", len(word), len(g.word))
	}
	return nil
}

func (g *Game) makeGuess(word string) guess {
	ge := guess{
		word:  word,
		marks: make([]Mark, len(word)),
	}
	for i, gc := range word {
		if gc == g.wordChars[i] {
			ge.marks[i] = Correct
		} else if _, ok := g.wordCharSet[gc]; ok {
			ge.marks[i] = Misplaced
		}
	}
	return ge
}

func (g *Game) Guess(word string) (bool, error) {
	if g.state != Running {
		return false, fmt.Errorf("invalid game state: %v", g.state)
	}
	if g.GuessesLeft() < 1 {
		panic(fmt.Errorf("no guesses left"))
	}

	if err := g.CheckGuess(word); err != nil {
		return false, err
	}

	ge := g.makeGuess(word)

	g.guesses = append(g.guesses, ge)
	for _, c := range word {
		g.guessedCharSet[c] = struct{}{}
	}

	if word == g.word {
		g.state = Won
		return true, nil
	}

	if g.GuessesLeft() < 1 {
		g.state = Lost
	}
	return false, nil
}
