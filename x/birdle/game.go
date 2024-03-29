//
/*
TODO:
 - !! free chars - special chars are pre-filled
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

type Game struct {
	word string
	dict *Dictionary

	state   State
	guesses []markedWord

	wordChars   []rune
	wordCharSet map[rune]struct{}

	guessedCharSet map[rune]struct{}
}

func NewGame(word string, guessesAllowed int, dict *Dictionary) (*Game, error) {
	if err := CheckWord(word); err != nil {
		return nil, err
	}
	if !dict.Contains(word) {
		return nil, fmt.Errorf("word not in dictionary")
	}
	if guessesAllowed < 1 {
		return nil, fmt.Errorf("invalid number of guesses")
	}

	wcs := make([]rune, len(word))
	wcss := make(map[rune]struct{})
	for i, c := range word {
		wcs[i] = c
		wcss[c] = struct{}{}
	}

	return &Game{
		word: word,
		dict: dict,

		guesses: make([]markedWord, 0, guessesAllowed),

		wordChars:   wcs,
		wordCharSet: wcss,

		guessedCharSet: make(map[rune]struct{}),
	}, nil
}

func (g *Game) Word() string { return g.word }
func (g *Game) State() State { return g.state }

func (g *Game) GetGuess(i int) MarkedWord { return g.guesses[i] }
func (g *Game) GuessesMade() int          { return len(g.guesses) }
func (g *Game) GuessesAllowed() int       { return cap(g.guesses) }
func (g *Game) GuessesLeft() int          { return cap(g.guesses) - len(g.guesses) }

func (g *Game) HasGuessedChar(c rune) bool {
	_, ok := g.guessedCharSet[c]
	return ok
}

func (g *Game) CheckGuess(word string) error {
	if err := CheckWord(word); err != nil {
		return err
	}
	if len(word) != len(g.word) {
		return fmt.Errorf("invalid guess length: guess %v != word %v", len(word), len(g.word))
	}
	if !g.dict.Contains(word) {
		return fmt.Errorf("word not in dictionary")
	}
	return nil
}

func (g *Game) markGuess(word string) markedWord {
	mw := markedWord{
		word:  word,
		marks: make([]Mark, len(word)),
	}
	for i, gc := range word {
		if gc == g.wordChars[i] {
			mw.marks[i] = Correct
		} else if _, ok := g.wordCharSet[gc]; ok {
			mw.marks[i] = Misplaced
		}
	}
	return mw
}

func (g *Game) Guess(word string) error {
	if g.state != Running {
		return fmt.Errorf("invalid game state: %v", g.state)
	}
	if g.GuessesLeft() < 1 {
		panic(fmt.Errorf("no guesses left"))
	}

	if err := g.CheckGuess(word); err != nil {
		return err
	}

	g.guesses = append(g.guesses, g.markGuess(word))
	for _, c := range word {
		g.guessedCharSet[c] = struct{}{}
	}

	if word == g.word {
		g.state = Won
		return nil
	}

	if g.GuessesLeft() < 1 {
		g.state = Lost
	}
	return nil
}
