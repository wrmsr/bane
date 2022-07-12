package main

import (
	"context"
	"os"
	"os/signal"
	"strings"

	"github.com/chzyer/readline"

	"github.com/wrmsr/bane/pkg/util/log"
)

var interrupted = make(chan os.Signal, 1)

func Repl() {
	signal.Notify(interrupted, os.Interrupt)
	defer signal.Stop(interrupted)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		select {
		case <-interrupted:
			cancel()
		case <-ctx.Done():
		}
	}()

	rl, err := readline.NewEx(&readline.Config{
		Prompt: ">>> ",
	})
	if err != nil {
		panic(err)
	}
	defer log.OrError(rl.Close)

	var cmds []string
	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		cmds = append(cmds, line)
		if !strings.HasSuffix(line, ";") {
			rl.SetPrompt("... ")
			continue
		}
		cmd := strings.Join(cmds, " ")
		cmds = cmds[:0]
		rl.SetPrompt(">>> ")
		println(cmd)
	}
}

func main() {
	//Repl()
}
