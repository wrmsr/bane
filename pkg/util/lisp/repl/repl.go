package main

import (
	"bufio"
	"fmt"
	"os"

	log "github.com/wrmsr/bane/pkg/util/slog"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		log.IfFatal(sc.Err())
		fmt.Println(sc.Text())
	}
}
