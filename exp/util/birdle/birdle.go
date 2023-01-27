//
/*
TODO:
 - bitmask chars
*/
package main

import (
	"fmt"
	"strings"
)

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

type Game struct {
	word        string
	guessesLeft int

	state   State
	guesses []string
	chars   map[rune]struct{}
}

func NewGame(word string, numGuesses int) (*Game, error) {
	if err := CheckWord(word); err != nil {
		return nil, err
	}
	if numGuesses < 1 {
		return nil, fmt.Errorf("invalid number of guesses: %v", numGuesses)
	}

	return &Game{
		word:        word,
		guessesLeft: numGuesses,

		chars: make(map[rune]struct{}),
	}, nil
}

func NormalizeWord(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	return s
}

func CheckWord(s string) error {
	if s != NormalizeWord(s) {
		return fmt.Errorf("word not normalized")
	}
	var l rune
	for _, c := range s {
		if !(c >= 'A' && c <= 'Z') && c != ' ' {
			return fmt.Errorf("word contains invalid character: %v", c)
		}
		if l == ' ' && c == ' ' {
			return fmt.Errorf("word contains consecutive spaces")
		}
		l = c
	}
	return nil
}

func (g *Game) Guess(guess string) (bool, error) {
	if g.state != Running {
		return false, fmt.Errorf("invalid game state: %v", g.state)
	}
	if g.guessesLeft < 1 {
		panic(fmt.Errorf("no guesses left"))
	}

	guess = NormalizeWord(guess)
	if err := CheckWord(guess); err != nil {
		return false, err
	}
	if len(guess) != len(g.word) {
		return false, fmt.Errorf("invalid guess length: guess %v != word %v", len(guess), len(g.word))
	}

	g.guesses = append(g.guesses, guess)
	g.guessesLeft--

	if guess != g.word {
		if g.guessesLeft < 1 {
			g.state = Lost
		}
		return false, nil
	}

	g.state = Won
	return true, nil
}
