package main

import (
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/x/util/birdle"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	//

	var wf birdle.WordFilter

	//wf = birdle.WordLenFilter(5)

	var ws []string

	useBirds := true

	if useBirds {
		ws = birdle.PrepareWords(func() []string {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			return check.Must1(birdle.FetchWikiLinks(ctx, birdle.DefaultWikiBirdsPage))
		}(), wf)
	} else {
		ws = check.Must1(birdle.ReadWordsFile(birdle.DefaultWordsFilePath, wf))
	}

	//

	d := birdle.NewDictionary(ws)
	r := birdle.TermRenderer{}

	rd := bufio.NewReader(os.Stdin)

	runGame := func(g *birdle.Game) {
		fmt.Println(r.RenderGame(g))

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

	//

	g := check.Must1(birdle.NewGame(d.Random(), 6, d))

	runGame(g)
}
