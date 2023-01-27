package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/wrmsr/bane/exp/util/birdle"
	"github.com/wrmsr/bane/pkg/util/check"
)

func main() {
	d := birdle.NewDictionary(check.Must1(birdle.ReadWordsFile(birdle.DefaultWordsFilePath, birdle.WordLenFilter(5))))
	g := check.Must1(birdle.NewGame(d.Random(), 5, d))
	r := birdle.TermRenderer{}

	rd := bufio.NewReader(os.Stdin)
	for {
		word := check.Must1(rd.ReadString('\n'))

		word = birdle.NormalizeWord(word)
		_, err := g.Guess(word)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(r.RenderGame(g))

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
