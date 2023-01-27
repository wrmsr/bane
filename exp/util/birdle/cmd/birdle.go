package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/wrmsr/bane/exp/util/birdle"
	"github.com/wrmsr/bane/pkg/util/check"
)

func main() {
	g := check.Must1(birdle.NewGame("FARTS", 5))

	r := bufio.NewReader(os.Stdin)
	for {
		word := check.Must1(r.ReadString('\n'))

		word = birdle.NormalizeWord(word)
		if _, err := g.Guess(word); err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(word)

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
