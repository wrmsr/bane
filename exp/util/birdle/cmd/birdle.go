package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/wrmsr/bane/exp/util/birdle"
	"github.com/wrmsr/bane/pkg/util/check"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	d := birdle.NewDictionary(check.Must1(birdle.ReadWordsFile(birdle.DefaultWordsFilePath, birdle.WordLenFilter(5))))
	g := check.Must1(birdle.NewGame(d.Random(), 6, d))
	r := birdle.TermRenderer{}

	rd := bufio.NewReader(os.Stdin)
	for {
		word := check.Must1(rd.ReadString('\n'))

		word = birdle.NormalizeWord(word)
		if err := g.Guess(word); err != nil {
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
		fmt.Printf("You lost! The word was: %s\n", g.Word())
	case birdle.Won:
		fmt.Println("You won!")
	default:
		panic(st)
	}
}
