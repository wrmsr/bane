package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/wrmsr/bane/exp/util/birdle"
	"github.com/wrmsr/bane/pkg/util/check"
)

func main() {
	d := birdle.NewDictionary(check.Must1(birdle.ReadWordsFile(birdle.DefaultWordsFilePath)))
	g := check.Must1(birdle.NewGame("FARTS", 5, d))
	r := birdle.TermRenderer{}

	rd := bufio.NewReader(os.Stdin)
	for {
		word := check.Must1(rd.ReadString('\n'))

		word = birdle.NormalizeWord(word)
		guess, err := g.Guess(word)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(r.RenderGuess(guess))

		if g.State() != birdle.Running {
			break
		}
	}

	switch st := g.State(); st {
	case birdle.Lost:
		fmt.Println("You lost!")
	case birdle.Won:
		fmt.Println("You won!")
	default:
		panic(st)
	}
}
